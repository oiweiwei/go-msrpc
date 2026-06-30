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

| Spec | Description | Package | Revision |
|------|-------------|---------|----------|
| [MS-ADTS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adts) | Active Directory Technical Specification: Claims | [msrpc/adts](./msrpc/adts) | 68.0 Major 5/25/2026 |
| [MS-BKRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bkrp) | BackupKey Remote Protocol | [msrpc/bkrp](./msrpc/bkrp) | 27.0 Major 4/23/2024 |
| [MS-BPAU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bpau) | BITS Peer-Caching: Peer Authentication Protocol | [msrpc/bpau](./msrpc/bpau) | 4.2 None 6/1/2017 |
| [MS-BRWSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-brwsa) | CIFS Browser Auxiliary Protocol | [msrpc/brwsa](./msrpc/brwsa) | 15.0 Major 4/23/2024 |
| [MS-CAPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-capr) | Central Access Policy ID Retrieval Protocol | [msrpc/capr](./msrpc/capr) | 9.0 Major 4/23/2024 |
| [MS-CMPO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmpo) | MSDTC Connection Manager: OleTx Transports Protocol | [msrpc/cmpo](./msrpc/cmpo) | 31.0 Major 7/29/2024 |
| [MS-CMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmrp) | Failover Cluster: Management API (ClusAPI) Protocol | [msrpc/cmrp](./msrpc/cmrp) | 44.0 Major 8/11/2025 |
| [MS-DFSNM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsnm) | DFS Namespace Management Protocol | [msrpc/dfsnm](./msrpc/dfsnm) | 33.0 Major 4/23/2024 |
| [MS-DHCPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dhcpm) | DHCP Server Management Protocol | [msrpc/dhcpm](./msrpc/dhcpm) | 36.0 None 9/16/2024 |
| [MS-DLTM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltm) | Distributed Link Tracking: Central Manager Protocol | [msrpc/dltm](./msrpc/dltm) | 9.1 None 6/1/2017 |
| [MS-DLTW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltw) | Distributed Link Tracking: Workstation Protocol | [msrpc/dltw](./msrpc/dltw) | 17.0 Major 4/23/2024 |
| [MS-DNSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dnsp) | DNS Server Management Protocol | [msrpc/dnsp](./msrpc/dnsp) | 40.0 Major 3/9/2026 |
| [MS-DRSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-drsr) | Directory Replication Service (DRS) Remote Protocol | [msrpc/drsr](./msrpc/drsr) | 46.0 Major 3/9/2026 |
| [MS-DSSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dssp) | Directory Services Setup Remote Protocol | [msrpc/dssp](./msrpc/dssp) | 16.0 Major 4/23/2024 |
| [MS-DTYP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp) | Windows Data Types | [msrpc/dtyp](./msrpc/dtyp) | 42.0 Major 11/19/2024 |
| [MS-EERR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-eerr) | ExtendedError Remote Data Structure | [msrpc/eerr](./msrpc/eerr) | 16.0 Major 4/23/2024 |
| [MS-EFSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-efsr) | Encrypting File System Remote (EFSRPC) Protocol | [msrpc/efsr](./msrpc/efsr) | 32.0 None 9/16/2024 |
| [MS-ERREF](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref) | Windows Error Codes | [msrpc/erref](./msrpc/erref) | 23.0 Major 11/19/2024 |
| [MS-EVEN6-BINXML](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6/c73573ae-1c90-43a2-a65f-ad7501155956) | EventLog BinXML encoding | [msrpc/binxml](./msrpc/binxml) | 25.0 Major 4/23/2024 |
| [MS-EVEN6](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6) | EventLog Remoting Protocol Version 6.0 | [msrpc/even6](./msrpc/even6) | 25.0 Major 4/23/2024 |
| [MS-EVEN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even) | EventLog Remoting Protocol | [msrpc/even](./msrpc/even) | 26.0 Major 4/23/2024 |
| [MS-FASP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fasp) | Firewall and Advanced Security Protocol | [msrpc/fasp](./msrpc/fasp) | 36.0 Major 6/8/2026 |
| [MS-FAX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fax) | Fax Server and Client Remote Protocol | [msrpc/fax](./msrpc/fax) | 29.0 None 4/23/2024 |
| [MS-FRS1](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs1) | File Replication Service (FRS) Remote Protocol | [msrpc/frs1](./msrpc/frs1) | 29.0 None 9/16/2024 |
| [MS-FSR2](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs2) | File Replication Service (FRS) Remote Protocol Version 2 | [msrpc/frs2](./msrpc/frs2) | 30.0 Major 4/23/2024 |
| [MS-ICPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-icpr) | ICertPassage Remote Protocol | [msrpc/icpr](./msrpc/icpr) | 24.0 Major 4/23/2024 |
| [MS-IRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-irp) | IIS Inetinfo Remote Protocol | [msrpc/irp](./msrpc/irp) | 13.0 Major 4/23/2024 |
| [MS-LREC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lrec) | Live Remote Event Capture (LREC) Protocol | [msrpc/lrec](./msrpc/lrec) | 5.0 Major 4/23/2024 |
| [MS-LSAD](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsad) | Local Security Authority (Domain Policy) Remote Protocol | [msrpc/lsad](./msrpc/lsad) | 49.0 Major 6/10/2025 |
| [MS-LSAT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsat) | Local Security Authority (Translation Methods) Remote Protocol | [msrpc/lsat](./msrpc/lsat) | 33.0 Major 5/1/2024 |
| [MS-MQDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqds) | MSMQ: Directory Service Protocol | [msrpc/mqds](./msrpc/mqds) | 24.1 None 6/1/2017 |
| [MS-MQMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmp) | MSMQ: Queue Manager Client Protocol | [msrpc/mqmp](./msrpc/mqmp) | 31.0 Major 4/23/2024 |
| [MS-MQMQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmq) | MSMQ: Data Structures | [msrpc/mqmq](./msrpc/mqmq) | 26.0 Major 4/23/2024 |
| [MS-MQMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmr) | MSMQ: Queue Manager Management Protocol | [msrpc/mqmr](./msrpc/mqmr) | 18.0 Major 4/23/2024 |
| [MS-MQQP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqqp) | MSMQ: Queue Manager to Queue Manager Protocol | [msrpc/mqqp](./msrpc/mqqp) | 25.0 Major 4/23/2024 |
| [MS-MQRR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqrr) | MSMQ: Queue Manager Remote Read Protocol | [msrpc/mqrr](./msrpc/mqrr) | 29.0 Major 11/21/2025 |
| [MS-MSRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-msrp) | Messenger Service Remote Protocol | [msrpc/msrp](./msrpc/msrp) | 10.3 None 6/1/2017 |
| [MS-NEGOEX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-negoex) | SPNEGO Extended Negotiation (NEGOEX) Security Mechanism | [msrpc/negoex](./msrpc/negoex) | 4.0 Major 4/23/2024 |
| [MS-NRPC-SECCHANNEL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc/fb50db72-7f71-478d-a180-12eb0ca3b36b) | Netlogon Secure Channel | [msrpc/nrpc](./msrpc/nrpc) | 49.0 Major 2/9/2026 |
| [MS-NRPC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc) | Netlogon Remote Protocol | [msrpc/nrpc](./msrpc/nrpc) | 49.0 Major 2/9/2026 |
| [MS-NSPI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nspi) | Name Service Provider Interface (NSPI) Protocol | [msrpc/nspi](./msrpc/nspi) | 16.0 Major 4/23/2024 |
| [MS-OXABREF](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxabref) | Address Book NSPI Referral Protocol | [msrpc/oxabref](./msrpc/oxabref) | 15.0 Major 5/20/2025 |
| [MS-OXCRPC](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcrpc) | Wire Format Protocol | [msrpc/oxcrpc](./msrpc/oxcrpc) | 25.0 Major 5/20/2025 |
| [MS-OXNSPI](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxnspi) | Exchange NSPI Protocol | [msrpc/nspi](./msrpc/nspi) | 15.0 Major 5/20/2025 |
| [MS-PAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pac) | Privilege Attribute Certificate Data Structure | [msrpc/pac](./msrpc/pac) | 26.0 Major 6/10/2024 |
| [MS-PAN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pan) | Print System Asynchronous Notification Protocol | [msrpc/pan](./msrpc/pan) | 19.0 Major 4/23/2024 |
| [MS-PAR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-par) | Print System Asynchronous Remote Protocol | [msrpc/par](./msrpc/par) | 19.0 Major 8/11/2025 |
| [MS-PCQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pcq) | Performance Counter Query Protocol | [msrpc/pcq](./msrpc/pcq) | 18.0 Major 4/23/2024 |
| [MS-RAA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raa) | Remote Authorization API Protocol | [msrpc/raa](./msrpc/raa) | 11.0 Major 4/23/2024 |
| [MS-RAIW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raiw) | Remote Administrative Interface: WINS | [msrpc/raiw](./msrpc/raiw) | 14.0 Major 4/23/2024 |
| [MS-RPCE-EPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpce) / [C706-EPM](https://pubs.opengroup.org/onlinepubs/9629399/apdxo.htm#tagcjh_35) | Endpoint Mapper | [msrpc/epm](./msrpc/epm) ||
| [MS-RPCL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpcl) | RPC Location Services Extensions | [msrpc/rpcl](./msrpc/rpcl) | 12.1 None 6/1/2017 |
| [MS-RPRN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn) | Print System Remote Protocol | [msrpc/rprn](./msrpc/rprn) | 39.0 8/11/2025 |
| [MS-RRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrp) | Windows Remote Registry Protocol | [msrpc/rrp](./msrpc/rrp) | 39.0 Major 10/21/2024 |
| [MS-RSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsp) | Remote Shutdown Protocol | [msrpc/rsp](./msrpc/rsp) | 12.0 Major 4/23/2024 |
| [MS-SAMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-samr) | Security Account Manager (SAM) Remote Protocol | [msrpc/samr](./msrpc/samr) | 51.1 Minor 4/27/2026 |
| [MS-SCMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmr) | Service Control Manager Remote Protocol | [msrpc/scmr](./msrpc/scmr) | 35.0 Major 8/11/2025 |
| [MS-SRVS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-srvs) | Server Service Remote Protocol | [msrpc/srvs](./msrpc/srvs) | 39.0 None 9/16/2024 |
| [MS-SSP](https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-ssp) | Single Sign-On Protocol | [msrpc/ssp](./msrpc/ssp) | 7.0 Major 10/5/2021 |
| [MS-SWN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-swn) | Service Witness Protocol | [msrpc/swn](./msrpc/swn) | 14.0 None 11/19/2024 |
| [MS-TRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-trp) | Telephony Remote Protocol | [msrpc/trp](./msrpc/trp) | 20.0 Major 4/23/2024 |
| [MS-TSCH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsch) | Task Scheduler Service Remoting Protocol | [msrpc/tsch](./msrpc/tsch) | 28.0 Major 4/23/2024 |
| [MS-TSGU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsgu) | Terminal Services Gateway Server Protocol | [msrpc/tsgu](./msrpc/tsgu) | 42.0 Major 4/23/2024 |
| [MS-TSTS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsts) | Terminal Services Runtime Interface Protocol | [msrpc/tsts](./msrpc/tsts) | 33.0 Major 11/21/2025 |
| [MS-W32T](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-w32t) | W32Time Remote Protocol | [msrpc/w32t](./msrpc/w32t) | 23.0 Major 4/23/2024 |
| [MS-WDSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wdsc) | Windows Deployment Services Control Protocol | [msrpc/wdsc](./msrpc/wdsc) | 9.0 Major 4/23/2024 |
| [MS-WKST](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wkst) | Workstation Service Remote Protocol | [msrpc/wkst](./msrpc/wkst) | 32.0 Major 4/23/2024 |

### DCOM Protocols

| Spec | Description | Package | Revision |
|------|-------------|---------|----------|
| [MC-CCFG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-ccfg) | Server Cluster: Configuration (ClusCfg) Protocol | [msrpc/dcom/ccfg](./msrpc/dcom/ccfg) | 1.2 None 1/17/2020 |
| [MC-IISA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-iisa) | IIS Application Host COM Protocol | [msrpc/dcom/iisa](./msrpc/dcom/iisa) | 15.0 Major 4/23/2024 |
| [MC-MQAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-mqac) | MSMQ: ActiveX Client Protocol | [msrpc/dcom/mqac](./msrpc/dcom/mqac) | 27.0 Major 4/23/2024 |
| [MS-ADTG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adtg) | Remote Data Services (RDS) Transport Protocol | [msrpc/dcom/adtg](./msrpc/dcom/adtg) | 19.1 Minor 5/30/2025 |
| [MS-COMA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-coma) | COM+ Remote Administration Protocol | [msrpc/dcom/coma](./msrpc/dcom/coma) | 15.0 Major 4/23/2024 |
| [MS-COMEV](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comev) | COM+ Event System Protocol | [msrpc/dcom/comev](./msrpc/dcom/comev) | 9.0 Major 9/9/2025 |
| [MS-COMT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comt) | COM+ Tracker Service Protocol | [msrpc/dcom/comt](./msrpc/dcom/comt) | 10.0 Major 4/23/2024 |
| [MS-COM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-com) | Component Object Model Plus (COM+) Protocol | [msrpc/dcom/com](./msrpc/dcom/com) | 12.0 Major 4/23/2024 |
| [MS-CSRA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csra) | Certificate Services Remote Administration Protocol | [msrpc/dcom/csra](./msrpc/dcom/csra) | 46.1 Minor 1/13/2026 |
| [MS-CSVP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csvp) | Failover Cluster: Setup and Validation Protocol (ClusPrep) | [msrpc/dcom/csvp](./msrpc/dcom/csvp) | 34.0 None 11/21/2025 |
| [MS-DCOM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dcom) | Distributed Component Object Model (DCOM) Remote Protocol | [msrpc/dcom](./msrpc/dcom) | 25.0 None 9/16/2024 |
| [MS-DFSRH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsrh) | DFS Replication Helper Protocol | [msrpc/dcom/dfsrh](./msrpc/dcom/dfsrh) | 17.0 Major 4/23/2024 |
| [MS-DMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dmrp) | Disk Management Remote Protocol | [msrpc/dcom/dmrp](./msrpc/dcom/dmrp) | 8.2 None 6/1/2017 |
| [MS-FSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fsrm) | File Server Resource Manager Protocol | [msrpc/dcom/fsrm](./msrpc/dcom/fsrm) | 36.0 Major 4/23/2024 |
| [MS-IISS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-iiss) | IIS ServiceControl Protocol | [msrpc/dcom/iiss](./msrpc/dcom/iiss) | 12.0 Major 4/23/2024 |
| [MS-IMSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-imsa) | IIS IMSAdminBaseW Remote Protocol | [msrpc/dcom/imsa](./msrpc/dcom/imsa) | 16.0 Major 4/23/2024 |
| [MS-IOI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ioi) | IManagedObject Interface Protocol | [msrpc/dcom/ioi](./msrpc/dcom/ioi) | 20.0 Major 3/13/2019 |
| [MS-OAUT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut) | OLE Automation Protocol | [msrpc/dcom/oaut](./msrpc/dcom/oaut) | 21.0 Major 4/23/2024 |
| [MS-OCSPA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ocspa) | Microsoft OCSP Administration Protocol | [msrpc/dcom/ocspa](./msrpc/dcom/ocspa) | 14.0 Major 8/11/2025 |
| [MS-PLA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pla) | Performance Logs and Alerts Protocol | [msrpc/dcom/pla](./msrpc/dcom/pla) | 26.0 Major 4/23/2024 |
| [MS-RAI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rai) | Remote Assistance Initiation Protocol | [msrpc/dcom/rai](./msrpc/dcom/rai) | 12.0 Major 4/23/2024 |
| [MS-RDPESC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpesc) | RDP: Smart Card Virtual Channel Extension | [msrpc/dcom/rdpesc](./msrpc/dcom/rdpesc) | 17.0 Major 4/23/2024 |
| [MS-RRASM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrasm) | RRAS Management Protocol | [msrpc/dcom/rrasm](./msrpc/dcom/rrasm) | 26.0 Major 4/23/2024 |
| [MS-RSMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsmp) | Removable Storage Manager (RSM) Remote Protocol | [msrpc/dcom/rsmp](./msrpc/dcom/rsmp) | 11.0 None 6/1/2017 |
| [MS-SCMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmp) | Shadow Copy Management Protocol | [msrpc/dcom/scmp](./msrpc/dcom/scmp) | 10.0 Major 4/23/2024 |
| [MS-TPMVSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tpmvsc) | TPM Virtual Smart Card Remote Protocol | [msrpc/dcom/tpmvsc](./msrpc/dcom/tpmvsc) | 8.0 Major 4/23/2024 |
| [MS-UAMG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-uamg) | Update Agent Management Protocol | [msrpc/dcom/uamg](./msrpc/dcom/uamg) | 15.0 Major 9/9/2025 |
| [MS-VDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-vds) | Virtual Disk Service (VDS) Protocol | [msrpc/dcom/vds](./msrpc/dcom/vds) | 31.0 None 9/16/2024 |
| [MS-WCCE](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wcce) | Windows Client Certificate Enrollment Protocol | [msrpc/dcom/wcce](./msrpc/dcom/wcce) | 53.0 Major 11/21/2025 |
| [MS-WMIO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmio) | WMI Encoding Version 1.0 Protocol | [msrpc/dcom/wmio](./msrpc/dcom/wmio) | 18.0 Major 4/23/2024 |
| [MS-WMI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmi) | Windows Management Instrumentation Remote Protocol | [msrpc/dcom/wmi](./msrpc/dcom/wmi) | 32.0 Major 4/23/2024 |
| [MS-WSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wsrm) | Windows System Resource Manager (WSRM) Protocol | [msrpc/dcom/wsrm](./msrpc/dcom/wsrm) | 14.2 None 6/1/2017 |

### Other

| Spec | Description | Package |
|------|-------------|---------|
| [MIMICOM](https://gist.github.com/gentilkiwi/e3d9c92b93ed4bb48f7956492c1d335a) | Mimikatz Command Interface | [msrpc/mimicom](./msrpc/mimicom) |

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
