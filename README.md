# go-msrpc

[![Go Reference](https://pkg.go.dev/badge/github.com/oiweiwei/go-msrpc.svg)](https://pkg.go.dev/github.com/oiweiwei/go-msrpc)
[![Go Report Card](https://goreportcard.com/badge/github.com/oiweiwei/go-msrpc)](https://goreportcard.com/report/github.com/oiweiwei/go-msrpc)

MS-RPC/DCOM client library for Go. Implements the [Microsoft Extension](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-rpce/290c38b1-92fe-4229-91e6-4fc376610c15) of [C706: DCE/RPC 1.1](http://www.dcerpc.org/documentation/) and includes ready-to-use generated stubs for all major Windows RPC and DCOM protocols: [Netlogon](./msrpc/nrpc), [Windows Registry](./msrpc/rrp), [Eventlog](./msrpc/even6), [WMI](./msrpc/dcom/wmi) ([query](./examples/wmic.go), [exec](./examples/wmiexec.go)), [DCOM/OXID](./msrpc/dcom), and [many more](#generated-stubs).

## Installation

```sh
go get github.com/oiweiwei/go-msrpc
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"

    "github.com/oiweiwei/go-msrpc/dcerpc"
    "github.com/oiweiwei/go-msrpc/msrpc/rrp/winreg/v1"
    "github.com/oiweiwei/go-msrpc/ssp"
    "github.com/oiweiwei/go-msrpc/ssp/credential"
    "github.com/oiweiwei/go-msrpc/ssp/gssapi"

    _ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

func main() {
    ctx := gssapi.NewSecurityContext(context.Background(),
        gssapi.WithCredential(credential.NewFromPassword("DOMAIN\\User", "password")),
        gssapi.WithMechanismFactory(ssp.SPNEGO),
        gssapi.WithMechanismFactory(ssp.NTLM),
    )

    // connect to server
    cc, err := dcerpc.Dial(ctx, "192.168.1.1", dcerpc.WithEndpoint("ncacn_np:[winreg]"))
    if err != nil {
        panic(err)
    }

    // create protocol client (MS-RRP)
    cli, err := winreg.NewWinregClient(ctx, cc, dcerpc.WithSeal())
    if err != nil {
        panic(err)
    }

    // make RPC call
    resp, err := cli.OpenLocalMachine(ctx, &winreg.OpenLocalMachineRequest{
        DesiredAccess: winreg.KeyRead,
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("HKLM handle: %v\n", resp.Key)
}
```



See [examples/samples_with_config](./examples/samples_with_config) and the [msrpc](./msrpc/doc.go) package documentation.

```sh
# Using string binding format
go run examples/samples_with_config/dnsp.go Administrator%P@ssw0rd@ncacn_ip_tcp:dc01.msad.local[privacy,spnego,krb5]

go run examples/samples_with_config/wmic.go Administrator%P@ssw0rd@ncacn_ip_tcp:dc01.msad.local[privacy,spnego,krb5] \
    --query "SELECT * FROM Win32_ComputerSystem"

# Using explicit flags
go run examples/samples_with_config/dnsp.go \
    --username=Administrator \
    --domain=MSAD.LOCAL \
    --password=P@ssw0rd \
    --auth-level=privacy \
    --auth-spnego \
    --auth-type=krb5 \
    --server=dc01.msad.local
```

Older examples in [examples/](./examples) use environment variables instead:

| Variable | Description | Example |
|----------|-------------|---------|
| `USERNAME` | Domain\Username | `"MSAD2.COM\User"` |
| `PASSWORD` | Password | `"password"` |
| `PASSWORD_MD4` | NT hash (generate with [nt_hash.go](./examples/helpers/nthash/nt_hash.go)) | `"f077ca4b7d73486a45e75dcdd74cd5bd"` |
| `WORKSTATION` | Workstation name | `"Ubuntu"` |
| `SERVER` | Server FQDN or IP | `"192.168.0.22"` |
| `SERVER_NAME` | Server NetBIOS name | `"WIN2019"` |
| `SERVER_HOST` | Server FQDN | `"my-server.win2019.com"` |
| `SAM_USERNAME` | Machine account name (see [netlogon_sec_channel.go](./examples/netlogon_sec_channel.go)) | `"COMPUTER$"` |
| `SAM_PASSWORD` | Machine account password | `"password"` |
| `SAM_WORKSTATION` | Machine account workstation | `"COMPUTER"` |
| `TARGET` | Kerberos SPN | `"host/my-server.win2019.com"` |
| `KRB5_CONFIG` | Kerberos config path | `"/path/to/krb5.conf"` |

> **Tip:** [RedTeamPentesting/adauth](https://github.com/RedTeamPentesting/adauth) integrates well with this library and provides a convenient way to handle Active Directory authentication (Kerberos, NTLM, pass-the-hash, PKINIT) from command-line tools.

## Code Generation

The IDL parser and Go code generator live at [github.com/oiweiwei/midl-gen-go](https://github.com/oiweiwei/midl-gen-go) and are consumed here via the published Docker image `ghcr.io/oiweiwei/midl-gen-go`.

Regenerate all stubs:

```sh
make all
```

Regenerate a single protocol (e.g. `nrpc.go`):

```sh
make nrpc.go
```

Under the hood each target runs:

```sh
docker run --rm \
  -v $(pwd):/work \
  -u $(id -u):$(id -g) \
  ghcr.io/oiweiwei/midl-gen-go generate \
  --pkg "github.com/oiweiwei/go-msrpc/msrpc/" \
  -I /work/idl \
  -I /work/idl/h \
  --output /work/msrpc/ \
  --msdn-openspecs-indexer-file /work/msdn/index.yaml \
  --msdn-openspecs-extra-file /work/msdn/extra.yaml \
  --msdn-openspecs-cache-dir /work/msdn/.cache/ \
  /work/idl/<protocol>.idl
```

To add a new IDL file, place it in `idl/` (or `idl/dcom/` for DCOM), then add a corresponding target to the `all` rule in the Makefile.

## Features

### DCE/RPC v5 Client (`dcerpc` package)

- Transfer Syntax: NDR 2.0 and NDR64
- Transports: Named Pipe (SMB2/3) and TCP
- Connection multiplexing: multiple clients over a single connection
- Multiple connections per association group, with shared context handles
- Verification trailer support
- Kerberos, Netlogon, NTLM, SPNEGO authentication
- Endpoint mapper and string binding support
- Basic DCOM support
- Eventlog BinXML parser
- WMIO object marshaler/unmarshaler

### MS-RPCE Extensions

- Security Context Multiplexing
- Bind-time Feature Negotiation
- Header Signing
- NDR64

### GSS-API / SSP (`ssp` package)

GSS-API interface definitions live in `ssp/gssapi`. The `ssp` package implements the following security providers:

- **Kerberos** (via [jcmturner/gokrb5 fork](https://github.com/oiweiwei/gokrb5.fork/tree/master/v9)):
  - Encryption: RC4-HMAC, DES-CBC-MD5, DES-CBC-CRC, AES128-CTS-HMAC-SHA1, AES256-CTS-HMAC-SHA1
  - DCE-style AP Request/Reply
  - Mutual and non-mutual authentication
  - Wrap/GetMic-Ex methods
- **NTLM**: NTLMv1 and NTLMv2
- **Netlogon**: RC4-HMAC and AES-SHA2
- **SPNEGO**: MechListMIC and NegTokenInit2

### SMB2 Client

Based on the [hirochachacha/go-smb2 fork](https://github.com/oiweiwei/go-smb2.fork), with the following additions:

- Force-encryption support
- Kerberos/NTLM integration via `ssp/gssapi`
- Fix for `NT_STATUS_PENDING`
- Keying material export (Application Key, Session Key)

## Generated Stubs

### RPC Protocols

| Spec | Description | Package |
|------|-------------|---------|
| [MS-ADTS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adts) | Active Directory Technical Specification: Claims | [msrpc/adts](./msrpc/adts) |
| [MS-BKRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bkrp) | BackupKey Remote Protocol | [msrpc/bkrp](./msrpc/bkrp) |
| [MS-BPAU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bpau) | BITS Peer-Caching: Peer Authentication Protocol | [msrpc/bpau](./msrpc/bpau) |
| [MS-BRWSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-brwsa) | CIFS Browser Auxiliary Protocol | [msrpc/brwsa](./msrpc/brwsa) |
| [MS-CAPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-capr) | Central Access Policy ID Retrieval Protocol | [msrpc/capr](./msrpc/capr) |
| [MS-CMPO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmpo) | MSDTC Connection Manager: OleTx Transports Protocol | [msrpc/cmpo](./msrpc/cmpo) |
| [MS-CMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmrp) | Failover Cluster: Management API (ClusAPI) Protocol | [msrpc/cmrp](./msrpc/cmrp) |
| [MS-DFSNM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsnm) | DFS Namespace Management Protocol | [msrpc/dfsnm](./msrpc/dfsnm) |
| [MS-DHCPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dhcpm) | DHCP Server Management Protocol | [msrpc/dhcpm](./msrpc/dhcpm) |
| [MS-DLTM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltm) | Distributed Link Tracking: Central Manager Protocol | [msrpc/dltm](./msrpc/dltm) |
| [MS-DLTW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltw) | Distributed Link Tracking: Workstation Protocol | [msrpc/dltw](./msrpc/dltw) |
| [MS-DNSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dnsp) | DNS Server Management Protocol | [msrpc/dnsp](./msrpc/dnsp) |
| [MS-DRSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-drsr) | Directory Replication Service (DRS) Remote Protocol | [msrpc/drsr](./msrpc/drsr) |
| [MS-DSSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dssp) | Directory Services Setup Remote Protocol | [msrpc/dssp](./msrpc/dssp) |
| [MS-DTYP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp) | Windows Data Types | [msrpc/dtyp](./msrpc/dtyp) |
| [MS-EERR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-eerr) | ExtendedError Remote Data Structure | [msrpc/eerr](./msrpc/eerr) |
| [MS-EFSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-efsr) | Encrypting File System Remote (EFSRPC) Protocol | [msrpc/efsr](./msrpc/efsr) |
| [MS-ERREF](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref) | Windows Error Codes | [msrpc/erref](./msrpc/erref) |
| [MS-EVEN6-BINXML](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6/c73573ae-1c90-43a2-a65f-ad7501155956) | EventLog BinXML encoding | [msrpc/binxml](./msrpc/binxml) |
| [MS-EVEN6](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6) | EventLog Remoting Protocol Version 6.0 | [msrpc/even6](./msrpc/even6) |
| [MS-EVEN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even) | EventLog Remoting Protocol | [msrpc/even](./msrpc/even) |
| [MS-FASP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fasp) | Firewall and Advanced Security Protocol | [msrpc/fasp](./msrpc/fasp) |
| [MS-FAX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fax) | Fax Server and Client Remote Protocol | [msrpc/fax](./msrpc/fax) |
| [MS-FRS1](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs1) | File Replication Service (FRS) Remote Protocol | [msrpc/frs1](./msrpc/frs1) |
| [MS-FSR2](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs2) | File Replication Service (FRS) Remote Protocol Version 2 | [msrpc/frs2](./msrpc/frs2) |
| [MS-ICPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-icpr) | ICertPassage Remote Protocol | [msrpc/icpr](./msrpc/icpr) |
| [MS-IRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-irp) | IIS Inetinfo Remote Protocol | [msrpc/irp](./msrpc/irp) |
| [MS-LREC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lrec) | Live Remote Event Capture (LREC) Protocol | [msrpc/lrec](./msrpc/lrec) |
| [MS-LSAD](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsad) | Local Security Authority (Domain Policy) Remote Protocol | [msrpc/lsad](./msrpc/lsad) |
| [MS-LSAT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsat) | Local Security Authority (Translation Methods) Remote Protocol | [msrpc/lsat](./msrpc/lsat) |
| [MS-MQDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqds) | MSMQ: Directory Service Protocol | [msrpc/mqds](./msrpc/mqds) |
| [MS-MQMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmp) | MSMQ: Queue Manager Client Protocol | [msrpc/mqmp](./msrpc/mqmp) |
| [MS-MQMQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmq) | MSMQ: Data Structures | [msrpc/mqmq](./msrpc/mqmq) |
| [MS-MQMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmr) | MSMQ: Queue Manager Management Protocol | [msrpc/mqmr](./msrpc/mqmr) |
| [MS-MQQP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqqp) | MSMQ: Queue Manager to Queue Manager Protocol | [msrpc/mqqp](./msrpc/mqqp) |
| [MS-MQRR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqrr) | MSMQ: Queue Manager Remote Read Protocol | [msrpc/mqrr](./msrpc/mqrr) |
| [MS-MSRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-msrp) | Messenger Service Remote Protocol | [msrpc/msrp](./msrpc/msrp) |
| [MS-NEGOEX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-negoex) | SPNEGO Extended Negotiation (NEGOEX) Security Mechanism | [msrpc/negoex](./msrpc/negoex) |
| [MS-NRPC-SECCHANNEL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc/fb50db72-7f71-478d-a180-12eb0ca3b36b) | Netlogon Secure Channel | [msrpc/nrpc](./msrpc/nrpc) |
| [MS-NRPC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc) | Netlogon Remote Protocol | [msrpc/nrpc](./msrpc/nrpc) |
| [MS-NSPI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nspi) | Name Service Provider Interface (NSPI) Protocol | [msrpc/nspi](./msrpc/nspi) |
| [MS-OXABREF](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxabref) | Address Book NSPI Referral Protocol | - |
| [MS-OXCRPC](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcrpc) | Wire Format Protocol | [msrpc/oxcrpc](./msrpc/oxcrpc) |
| [MS-OXNSPI](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxnspi) | Exchange NSPI Protocol | [msrpc/nspi](./msrpc/nspi) |
| [MS-PAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pac) | Privilege Attribute Certificate Data Structure | [msrpc/pac](./msrpc/pac) |
| [MS-PAN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pan) | Print System Asynchronous Notification Protocol | [msrpc/pan](./msrpc/pan) |
| [MS-PAR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-par) | Print System Asynchronous Remote Protocol | [msrpc/par](./msrpc/par) |
| [MS-PCQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pcq) | Performance Counter Query Protocol | [msrpc/pcq](./msrpc/pcq) |
| [MS-RAA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raa) | Remote Authorization API Protocol | [msrpc/raa](./msrpc/raa) |
| [MS-RAIW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raiw) | Remote Administrative Interface: WINS | [msrpc/raiw](./msrpc/raiw) |
| [MS-RPCE-EPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpce) / [C706-EPM](https://pubs.opengroup.org/onlinepubs/9629399/apdxo.htm#tagcjh_35) | Endpoint Mapper | [msrpc/epm](./msrpc/epm) |
| [MS-RPCL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpcl) | RPC Location Services Extensions | [msrpc/rpcl](./msrpc/rpcl) |
| [MS-RPRN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn) | Print System Remote Protocol | [msrpc/rprn](./msrpc/rprn) |
| [MS-RRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrp) | Windows Remote Registry Protocol | [msrpc/rrp](./msrpc/rrp) |
| [MS-RSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsp) | Remote Shutdown Protocol | [msrpc/rsp](./msrpc/rsp) |
| [MS-SAMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-samr) | Security Account Manager (SAM) Remote Protocol | [msrpc/samr](./msrpc/samr) |
| [MS-SCH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsch) | Task Scheduler Service Remoting Protocol | [msrpc/sch](./msrpc/sch) |
| [MS-SCMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmr) | Service Control Manager Remote Protocol | [msrpc/scmr](./msrpc/scmr) |
| [MS-SRVS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-srvs) | Server Service Remote Protocol | [msrpc/srvs](./msrpc/srvs) |
| [MS-SSP](https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-ssp) | Single Sign-On Protocol | [msrpc/ssp](./msrpc/ssp) |
| [MS-SWN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-swn) | Service Witness Protocol | [msrpc/swn](./msrpc/swn) |
| [MS-TRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-trp) | Telephony Remote Protocol | [msrpc/trp](./msrpc/trp) |
| [MS-TSCH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsch) | Task Scheduler Service Remoting Protocol | [msrpc/tsch](./msrpc/tsch) |
| [MS-TSGU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsgu) | Terminal Services Gateway Server Protocol | [msrpc/tsgu](./msrpc/tsgu) |
| [MS-TSTS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsts) | Terminal Services Runtime Interface Protocol | [msrpc/tsts](./msrpc/tsts) |
| [MS-W32T](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-w32t) | W32Time Remote Protocol | [msrpc/w32t](./msrpc/w32t) |
| [MS-WDSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wdsc) | Windows Deployment Services Control Protocol | [msrpc/wdsc](./msrpc/wdsc) |
| [MS-WKST](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wkst) | Workstation Service Remote Protocol | [msrpc/wkst](./msrpc/wkst) |

### DCOM Protocols

| Spec | Description | Package |
|------|-------------|---------|
| [MC-CCFG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-ccfg) | Server Cluster: Configuration (ClusCfg) Protocol | [msrpc/dcom/ccfg](./msrpc/dcom/ccfg) |
| [MC-IISA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-iisa) | IIS Application Host COM Protocol | [msrpc/dcom/iisa](./msrpc/dcom/iisa) |
| [MC-MQAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-mqac) | MSMQ: ActiveX Client Protocol | [msrpc/dcom/mqac](./msrpc/dcom/mqac) |
| [MS-ADTG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adtg) | Remote Data Services (RDS) Transport Protocol | [msrpc/dcom/adtg](./msrpc/dcom/adtg) |
| [MS-COMA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-coma) | COM+ Remote Administration Protocol | [msrpc/dcom/coma](./msrpc/dcom/coma) |
| [MS-COMEV](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comev) | COM+ Event System Protocol | [msrpc/dcom/comev](./msrpc/dcom/comev) |
| [MS-COMT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comt) | COM+ Tracker Service Protocol | [msrpc/dcom/comt](./msrpc/dcom/comt) |
| [MS-COM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-com) | Component Object Model Plus (COM+) Protocol | [msrpc/dcom/com](./msrpc/dcom/com) |
| [MS-CSRA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csra) | Certificate Services Remote Administration Protocol | [msrpc/dcom/csra](./msrpc/dcom/csra) |
| [MS-CSVP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csvp) | Failover Cluster: Setup and Validation Protocol (ClusPrep) | [msrpc/dcom/csvp](./msrpc/dcom/csvp) |
| [MS-DCOM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dcom) | Distributed Component Object Model (DCOM) Remote Protocol | [msrpc/dcom](./msrpc/dcom) |
| [MS-DFSRH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsrh) | DFS Replication Helper Protocol | [msrpc/dcom/dfsrh](./msrpc/dcom/dfsrh) |
| [MS-DMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dmrp) | Disk Management Remote Protocol | [msrpc/dcom/dmrp](./msrpc/dcom/dmrp) |
| [MS-FSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fsrm) | File Server Resource Manager Protocol | [msrpc/dcom/fsrm](./msrpc/dcom/fsrm) |
| [MS-IISS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-iiss) | IIS ServiceControl Protocol | [msrpc/dcom/iiss](./msrpc/dcom/iiss) |
| [MS-IMSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-imsa) | IIS IMSAdminBaseW Remote Protocol | [msrpc/dcom/imsa](./msrpc/dcom/imsa) |
| [MS-IOI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ioi) | IManagedObject Interface Protocol | [msrpc/dcom/ioi](./msrpc/dcom/ioi) |
| [MS-OAUT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut) | OLE Automation Protocol | [msrpc/dcom/oaut](./msrpc/dcom/oaut) |
| [MS-OCSPA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ocspa) | Microsoft OCSP Administration Protocol | [msrpc/dcom/ocspa](./msrpc/dcom/ocspa) |
| [MS-PLA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pla) | Performance Logs and Alerts Protocol | [msrpc/dcom/pla](./msrpc/dcom/pla) |
| [MS-RAI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rai) | Remote Assistance Initiation Protocol | [msrpc/dcom/rai](./msrpc/dcom/rai) |
| [MS-RDPESC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpesc) | RDP: Smart Card Virtual Channel Extension | [msrpc/dcom/rdpesc](./msrpc/dcom/rdpesc) |
| [MS-RRASM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrasm) | RRAS Management Protocol | [msrpc/dcom/rrasm](./msrpc/dcom/rrasm) |
| [MS-RSMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsmp) | Removable Storage Manager (RSM) Remote Protocol | [msrpc/dcom/rsmp](./msrpc/dcom/rsmp) |
| [MS-SCMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmp) | Shadow Copy Management Protocol | [msrpc/dcom/scmp](./msrpc/dcom/scmp) |
| [MS-TPMVSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tpmvsc) | TPM Virtual Smart Card Remote Protocol | [msrpc/dcom/tpmvsc](./msrpc/dcom/tpmvsc) |
| [MS-UAMG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-uamg) | Update Agent Management Protocol | [msrpc/dcom/uamg](./msrpc/dcom/uamg) |
| [MS-VDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-vds) | Virtual Disk Service (VDS) Protocol | [msrpc/dcom/vds](./msrpc/dcom/vds) |
| [MS-WCCE](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wcce) | Windows Client Certificate Enrollment Protocol | [msrpc/dcom/wcce](./msrpc/dcom/wcce) |
| [MS-WMIO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmio) | WMI Encoding Version 1.0 Protocol | [msrpc/dcom/wmio](./msrpc/dcom/wmio) |
| [MS-WMI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmi) | Windows Management Instrumentation Remote Protocol | [msrpc/dcom/wmi](./msrpc/dcom/wmi) |
| [MS-WSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wsrm) | Windows System Resource Manager (WSRM) Protocol | [msrpc/dcom/wsrm](./msrpc/dcom/wsrm) |

### Other

| Spec | Description | Package |
|------|-------------|---------|
| [MIMICOM](https://gist.github.com/gentilkiwi/e3d9c92b93ed4bb48f7956492c1d335a) | Mimikatz COM Interface | [msrpc/mimicom](./msrpc/mimicom) |

### Documentation

Generated code includes documentation pulled from the MSDN portal. Accuracy may vary due to inconsistencies in the upstream HTML source.

## Open Questions

- Why does IObjectExporter not support NDR64?
- Why does the server return indistinguishable pointers for NDR64?
- Why does SMB2 not support certain auth levels (e.g. Winreg supports only Insecure and Privacy)?

## References

- [Samba](https://www.samba.org/)
- [Impacket](https://github.com/fortra/impacket)
- [go-smb2](https://github.com/hirochachacha/go-smb2)
- [gokrb5](https://github.com/jcmturner/gokrb5)

## Contributing

Open an issue before submitting a PR. The project is still maturing and there are likely undiscovered bugs.
