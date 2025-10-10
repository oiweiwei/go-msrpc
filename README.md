# MS-RPC IDL Parser/Codegen for Go / MS-RPC/DCOM Client

The IDL parser for the [Microsoft Extension](https://docs.microsoft.com/en-us/openspecs/windows_protocols/ms-rpce/290c38b1-92fe-4229-91e6-4fc376610c15) of the [C706: DCE/RPC 1.1](http://www.dcerpc.org/documentation/).

The client stub generator for many [MSRPC](./msrpc) / [DCOM](./msrpc/dcom) services including (but not limited to - see complete list below)
[Netlogon](./msrpc/nrpc), [Windows Registry](./msrpc/rrp), [Eventlog](./msrpc/even6), [DCOM (OXID resolver)](./msrpc/dcom), [WMI](./msrpc/dcom/wmi) ([query](./examples/wmic.go) and [method exec](./examples/wmiexec.go)) support.

## Usage

### Codegen

For codegeneration, run `make all` to regenerate all sources, or `make nrpc.go` to regenerate specific IDL.

To onboard new file (example.idl), add it to corresponding folder (`idl/`, or for DCOM `idl/dcom/`), update Makefile `all` target with `all.go` file.

### Examples

See [examples/samples_with_config](./examples/samples_with_config) and [msrpc](./msrpc/doc.go) package documentation.

```sh
# run using string binding extension.
go run examples/samples_with_config/dnsp.go Administrator%P@ssw0rd@ncacn_ip_tcp:dc01.msad.local[privacy,spnego,krb5]

go run examples/samples_with_config/wmic.go Administrator%P@ssw0rd@ncacn_ip_tcp:dc01.msad.local[privacy,spnego,krb5] \
    --query "SELECT * FROM Win32_ComputerSystem"

# same as above, but using command-line args
go run examples/samples_with_config/dnsp.go \
    --username=Administrator \
    --domain=MSAD.LOCAL \
    --password=P@ssw0rd \
    --auth-level=privacy \
    --auth-spnego \
    --auth-type=krb5 \
    --server=dc01.msad.local
```

### Examples (Old)

See [examples](./examples) and [dcerpc](./dcerpc/doc.go) package documentation.

Examples rely on following environment variables:

| Name | Description | Example |
| ---- | ----------- | ------- |
| **USERNAME** | The Domain\Username | `"MSAD2.COM\User"` |
| **PASSWORD** | The password | `"password"` |
| **PASSWORD_MD4** | The password hash (use [go run examples/helpers/nt_hash.go -d $PASSWORD](./examples/helpers/nt_hash.go) to generate the hash) | `"f077ca4b7d73486a45e75dcdd74cd5bd"` |
| **WORKSTATION** | The workstation name | `"Ubuntu"` |
| **SERVER** | The server FQDN or IP | `"192.168.0.22"` |
| **SERVER_NAME** | The server NetBIOS name | `"WIN2019"` |
| **SERVER_HOST** | The server FQDN | `"my-server.win2019.com"` |
| **SAM_USERNAME** | The machine account name (see [examples/netlogon_sec_channel.go](./examples/netlogon_sec_channel.go)) | `"COMPUTER$"` |
| **SAM_PASSWORD** | The machine account password (see [examples/netlogon_sec_channel.go](./examples/netlogon_sec_channel.go)) | `"password"` |
| **SAM_WORKSTATION** | The machine account workstation name | `"COMPUTER"` |
| **TARGET** | The target name (SPN) for kerberos. | `"host/my-server.win2019.com"` |
| **KRB5_CONFIG** | The kerberos config path. | `"/path/to/krb5.conf"` |

## Features

### Connection-oriented DCE/RPC v5 client implementation

The library implements the CO RPC v5 (`dcerpc` package) with following features:

 * Transfer Syntax NDR2.0 and NDR64

 * CO transport over Named Pipe (SMB2/3) and TCP.

 * Connection Multiplexing: multiple clients over single connection

 * Multiple Connection per Association Group: ability to use context handles
   from one connection on another, flexibility in arranging the
   clients-per-connection-per-association

 * Verification Trailer: ability to add verification trailer to the request
   payload

 * Kerberos, Netlogon, NTLM, SPNEGO Authentication

 * Endpoint mapper / string binding support

 * DCOM basic support

 * Eventlog BinXML parser

 * WMIO object unmarshaler / marshaler.

### MS-RPCE Extensions

The library implements some of the extensions defined in MS-RPCE document:

 * Security Context Multiplexing: ability to create multiple security contexts
   over the same logical connection.

 * Bind-time Feature Negotiation: (actually not a feature).

 * Header Signing: (legacy thing)

 * NDR64

### GSS-API / SSP Client Side

The library contains the GSS-API interface definitions. (`ssp/gssapi`)

The library contains the `ssp` package which has an implementation for the
various security service providers, like Kerberos, NTLM, Netlogon (Secure Channel),
SPNEGO.

The kerberos implementation is based on the [jcmturner/gokrb5 fork](https://github.com/oiweiwei/gokrb5.fork/tree/master/v9).
Any changes or feature requests should be addressed there.

 * GSSAPI interface implementation including Wrap/GetMic-Ex-methods defined in Microsoft
   documentation

 * Kerberos:

    * Supported Encryption Types:

        * RC4-HMAC

        * DES-CBC-MD5

        * DES-CBC-CRC

        * AES128-CTS-HMAC-SHA1

        * AES256-CTS-HMAC-SHA1

    * DCE Style AP Request and AP Reply

    * Mutual and Non-mutual Authn

 * NTLM

    * Supported Versions: NTLMv1, NTLMv2

 * Netlogon:

    * Supported Encryption Types:

        * RC4-HMAC

        * AES-SHA2

 * SPNEGO:

    * Supported Mech List MIC

    * Supported NegTokenInit2

### SMB2 Client

The SMB2 client implementation is based on the [hirochachacha/go-smb2 fork](https://github.com/oiweiwei/go-smb2.fork).
Any changes or feature requests should be addressed there.

The set of changes includes:

  * SMB2 Force-Encryption Support

  * Integration with ssp/gssapi for Kerberos/NTLM authentication.

  * Fix for `NT_STATUS_PENDING` error

  * Keying material export (Application Key, Session Key)

## Generated Stubs

| Code | Description | Package |
| ------- | -------- | ------- |
| [MS-ADTS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adts) | Active Directory Technical Specification: Claims | [github.com/oiweiwei/go-msrpc/msrpc/adts](./msrpc/adts) |
| [MS-BKRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bkrp) | BackupKey Remote Protocol | [github.com/oiweiwei/msrpc/bkrp](./msrpc/bkrp) |
| [MS-BPAU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-bpau) | Background Intelligent Transfer Service (BITS) Peer-Caching: Peer Authentication Protocol | [github.com/oiweiwei/msrpc/bpau](./msrpc/bpau) |
| [MS-BRWSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-brwsa) | Common Internet File System (CIFS) Browser Auxiliary Protocol | [github.com/oiweiwei/msrpc/brwsa](./msrpc/brwsa) |
| [MS-CAPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-capr) | Central Access Policy Identifier (ID) Retrieval Protocol | [github.com/oiweiwei/msrpc/capr](./msrpc/capr) |
| [MS-CMPO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmpo) | MSDTC Connection Manager: OleTx Transports Protocol | [github.com/oiweiwei/msrpc/cmpo](./msrpc/cmpo) |
| [MS-CMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-cmrp) | Failover Cluster: Management API (ClusAPI) Protocol | [github.com/oiweiwei/msrpc/cmrp](./msrpc/cmrp) |
| [MS-DFSNM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsnm) | Distributed File System (DFS): Namespace Management Protocol | [github.com/oiweiwei/msrpc/dfsnm](./msrpc/dfsnm) |
| [MS-DHCPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dhcpm) | Microsoft Dynamic Host Configuration Protocol (DHCP) Server Management Protocol | [github.com/oiweiwei/msrpc/dhcpm](./msrpc/dhcpm) |
| [MS-DLTM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltm) | Distributed Link Tracking: Central Manager Protocol | [github.com/oiweiwei/msrpc/dltm](./msrpc/dltm) |
| [MS-DLTW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dltw) | Distributed Link Tracking: Workstation Protocol | [github.com/oiweiwei/msrpc/dltw](./msrpc/dltw) |
| [MS-DNSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dnsp) | Domain Name Service (DNS) Server Management ProtocolDomain Name Service (DNS) Server Management Protocol | [github.com/oiweiwei/msrpc/dnsp](./msrpc/dnsp) |
| [MS-DRSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-drsr) | Directory Replication Service (DRS) Remote Protocol | [github.com/oiweiwei/msrpc/drsr](./msrpc/drsr) |
| [MS-DSSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dssp) | Directory Services Setup Remote Protocol | [github.com/oiweiwei/msrpc/dssp](./msrpc/dssp) |
| [MS-DTYP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dtyp) | Windows Data Types | [github.com/oiweiwei/msrpc/dtyp](./msrpc/dtyp) |
| [MS-EERR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-eerr) | ExtendedError Remote Data Structure | [github.com/oiweiwei/msrpc/eerr](./msrpc/eerr) |
| [MS-EFSR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-efsr) | Encrypting File System Remote (EFSRPC) Protocol | [github.com/oiweiwei/msrpc/efsr](./msrpc/efsr) |
| [MS-ERREF](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-erref) | Windows Error Codes | [github.com/oiweiwei/msrpc/erref](./msrpc/erref) |
| [MS-EVEN6-BINXML](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6/c73573ae-1c90-43a2-a65f-ad7501155956) | BinXml encodes an XML document so that the original XML text can be correctly reproduced from the encoding. | [github.com/oiweiwei/msrpc/binxml](./msrpc/binxml) |
| [MS-EVEN6](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even6) | EventLog Remoting Protocol Version 6.0 | [github.com/oiweiwei/msrpc/even6](./msrpc/even6) |
| [MS-EVEN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-even) | EventLog Remoting Protocol | [github.com/oiweiwei/msrpc/even](./msrpc/even) |
| [MS-FASP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fasp) | Firewall and Advanced Security Protocol | [github.com/oiweiwei/msrpc/fasp](./msrpc/fasp) |
| [MS-FAX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fax) | Fax Server and Client Remote Protocol | [github.com/oiweiwei/msrpc/fax](./msrpc/fax) |
| [MS-FRS1](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs1) | File Replication Service (FRS) Remote Protocol | [github.com/oiweiwei/msrpc/frs1](./msrpc/frs1) |
| [MS-FSR2](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-frs2) | File Replication Service (FRS) Remote Protocol Version 2 | [github.com/oiweiwei/msrpc/frs2](./msrpc/frs2) |
| [MS-ICPR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-icpr) | ICertPassage Remote Protocol | [github.com/oiweiwei/msrpc/icpr](./msrpc/icpr) |
| [MS-IRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-irp) | Internet Information Services (IIS) Inetinfo Remote Protocol | [github.com/oiweiwei/msrpc/irp](./msrpc/irp) |
| [MS-LREC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lrec) | Live Remote Event Capture (LREC) Protocol | [github.com/oiweiwei/msrpc/lrec](./msrpc/lrec) |
| [MS-LSAD](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsad) | Local Security Authority (Domain Policy) Remote Protocol | [github.com/oiweiwei/msrpc/lsad](./msrpc/lsad) |
| [MS-LSAT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-lsat) | Local Security Authority (Translation Methods) Remote Protocol | [github.com/oiweiwei/msrpc/lsat](./msrpc/lsat) |
| [MS-MQDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqds) | Message Queuing (MSMQ): Directory Service Protocol | [github.com/oiweiwei/msrpc/mqds](./msrpc/mqds) |
| [MS-MQMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmp) | Message Queuing (MSMQ): Queue Manager Client Protocol | [github.com/oiweiwei/msrpc/mqmp](./msrpc/mqmp) |
| [MS-MQMQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmq) | Message Queuing (MSMQ): Data Structures | [github.com/oiweiwei/msrpc/mqmq](./msrpc/mqmq) |
| [MS-MQMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqmr) | Message Queuing (MSMQ): Queue Manager Management Protocol | [github.com/oiweiwei/msrpc/mqmr](./msrpc/mqmr) |
| [MS-MQQP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqqp) | Message Queuing (MSMQ): Queue Manager to Queue Manager Protocol | [github.com/oiweiwei/msrpc/mqqp](./msrpc/mqqp) |
| [MS-MQRR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-mqrr) | Message Queuing (MSMQ): Queue Manager Remote Read Protocol | [github.com/oiweiwei/msrpc/mqrr](./msrpc/mqrr) |
| [MS-MSRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-msrp) | Messenger Service Remote Protocol | [github.com/oiweiwei/msrpc/msrp](./msrpc/msrp) |
| [MS-NEGOEX](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-negoex) | SPNEGO Extended Negotiation (NEGOEX) Security Mechanism | [github.com/oiweiwei/msrpc/negoex](./msrpc/negoex) |
| [MS-NRPC-SECCHANNEL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc/fb50db72-7f71-478d-a180-12eb0ca3b36b) | Secure Channel Establishment and Maintenance | [github.com/oiweiwei/msrpc/nrpc](./msrpc/nrpc) |
| [MS-NRPC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nrpc) | Netlogon Remote Protocol | [github.com/oiweiwei/msrpc/nrpc](./msrpc/nrpc) |
| [MS-NSPI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-nspi) | Name Service Provider Interface (NSPI) Protocol | [github.com/oiweiwei/msrpc/nspi](./msrpc/nspi) |
| [MS-OXABREF](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxabref) | Address Book Name Service Provider Interface (NSPI) Referral Protocol |
| [MS-OXCRPC](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxcrpc) | Wire Format Protocol | [github.com/oiweiwei/msrpc/oxcrpc](./msrpc/oxcrpc) |
| [MS-OXNSPI](https://learn.microsoft.com/en-us/openspecs/exchange_server_protocols/ms-oxnspi) | Name Service Provider Interface (NSPI) Protocol | [github.com/oiweiwei/msrpc/nspi](./msrpc/nspi) |
| [MS-PAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pac) | Privilege Attribute Certificate Data Structure | [github.com/oiweiwei/msrpc/pac](./msrpc/pac) |
| [MS-PAN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pan) | Print System Asynchronous Notification Protocol | [github.com/oiweiwei/msrpc/pan](./msrpc/pan) |
| [MS-PAR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-par) | Print System Asynchronous Remote Protocol | [github.com/oiweiwei/msrpc/par](./msrpc/par) |
| [MS-PCQ](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pcq) | Performance Counter Query Protocol | [github.com/oiweiwei/msrpc/pcq](./msrpc/pcq) |
| [MS-RAA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raa) | Remote Authorization API Protocol | [github.com/oiweiwei/msrpc/raa](./msrpc/raa) |
| [MS-RAIW](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-raiw) | Remote Administrative Interface: WINS | [github.com/oiweiwei/msrpc/raiw](./msrpc/raiw) |
| [MS-RPCE-EPM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpce) [C706-EPM](https://pubs.opengroup.org/onlinepubs/9629399/apdxo.htm#tagcjh_35) | Endpoint Mapper | [github.com/oiweiwei/msrpc/epm](./msrpc/epm) |
| [MS-RPCL](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rpcl) | Remote Procedure Call Location Services Extensions | [github.com/oiweiwei/msrpc/rpcl](./msrpc/rpcl) |
| [MS-RPRN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rprn) | Print System Remote Protocol | [github.com/oiweiwei/msrpc/rprn](./msrpc/rprn) |
| [MS-RRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrp) | Windows Remote Registry Protocol | [github.com/oiweiwei/msrpc/rrp](./msrpc/rrp) |
| [MS-RSP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsp) | Remote Shutdown Protocol | [github.com/oiweiwei/msrpc/rsp](./msrpc/rsp) |
| [MS-SAMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-samr) | Security Account Manager (SAM) Remote Protocol (Client-to-Server) | [github.com/oiweiwei/msrpc/samr](./msrpc/samr) |
| [MS-SCH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsch) | Task Scheduler Service Remoting Protocol | [github.com/oiweiwei/msrpc/sch](./msrpc/sch) |
| [MS-SCMR](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmr) | Service Control Manager Remote Protocol | [github.com/oiweiwei/msrpc/scmr](./msrpc/scmr) |
| [MS-SRVS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-srvs) | Server Service Remote Protocol | [github.com/oiweiwei/msrpc/srvs](./msrpc/srvs) |
| [MS-SSP](https://learn.microsoft.com/en-us/openspecs/sharepoint_protocols/ms-ssp) | Single Sign-On Protocol | [github.com/oiweiwei/msrpc/ssp](./msrpc/ssp) |
| [MS-SWN](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-swn) | Service Witness Protocol | [github.com/oiweiwei/msrpc/swn](./msrpc/swn) |
| [MS-TRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-trp) | Telephony Remote Protocol | [github.com/oiweiwei/msrpc/trp](./msrpc/trp) |
| [MS-TSCH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsch) | Task Scheduler Service Remoting Protocol | [github.com/oiweiwei/msrpc/tsch](./msrpc/tsch) |
| [MS-TSGU](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tsgu) | Terminal Services Gateway Server Protocol | [github.com/oiweiwei/msrpc/tsgu](./msrpc/tsgu) |
| [MS-W32T](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-w32t) | W32Time Remote Protocol | [github.com/oiweiwei/msrpc/w32t](./msrpc/w32t) |
| [MS-WDSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wdsc) | Windows Deployment Services Control Protocol | [github.com/oiweiwei/msrpc/wdsc](./msrpc/wdsc) |
| [MS-WKST](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wkst) | Workstation Service Remote Protocol | [github.com/oiweiwei/msrpc/wkst](./msrpc/wkst) |

### Generated DCOM Stubs

| Code | Description | Package |
| ------- | -------- | ------- |
| [MC-CCFG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-ccfg) | Server Cluster: Configuration (ClusCfg) Protocol | [github.com/oiweiwei/msrpc/ccfg](./msrpc/dcom/ccfg) |
| [MC-IISA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-iisa) | Internet Information Services (IIS) Application Host COM Protocol | [github.com/oiweiwei/msrpc/iisa](./msrpc/dcom/iisa) |
| [MC-MQAC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/mc-mqac) | Message Queuing (MSMQ): ActiveX Client Protocol | [github.com/oiweiwei/msrpc/mqac](./msrpc/dcom/mqac) |
| [MS-ADTG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adtg) | Remote Data Services (RDS) Transport Protocol | [github.com/oiweiwei/msrpc/adtg](./msrpc/dcom/adtg) |
| [MS-COMA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-coma) | Component Object Model Plus (COM+) Remote Administration Protocol | [github.com/oiweiwei/msrpc/coma](./msrpc/dcom/coma) |
| [MS-COMEV](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comev) | Component Object Model Plus (COM+) Event System Protocol | [github.com/oiweiwei/msrpc/comev](./msrpc/dcom/comev) |
| [MS-COMT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-comt) | Component Object Model Plus (COM+) Tracker Service Protocol | [github.com/oiweiwei/msrpc/comt](./msrpc/dcom/comt) |
| [MS-COM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-com) | Component Object Model Plus (COM+) Protocol | [github.com/oiweiwei/msrpc/com](./msrpc/dcom/com) |
| [MS-CSRA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csra) | Certificate Services Remote Administration Protocol | [github.com/oiweiwei/msrpc/csra](./msrpc/dcom/csra) |
| [MS-CSVP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-csvp) | Failover Cluster: Setup and Validation Protocol (ClusPrep) | [github.com/oiweiwei/msrpc/csvp](./msrpc/dcom/csvp) |
| [MS-DCOM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dcom) | Distributed Component Object Model (DCOM) Remote Protocol | [github.com/oiweiwei/msrpc/dcom](./msrpc/dcom) |
| [MS-DFSRH](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dfsrh) | DFS Replication Helper Protocol | [github.com/oiweiwei/msrpc/dfsrh](./msrpc/dcom/dfsrh) |
| [MS-DMRP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-dmrp) | Disk Management Remote Protocol | [github.com/oiweiwei/msrpc/dmrp](./msrpc/dcom/dmrp) |
| [MS-FSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-fsrm) | File Server Resource Manager Protocol | [github.com/oiweiwei/msrpc/fsrm](./msrpc/dcom/fsrm) |
| [MS-IISS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-iiss) | Internet Information Services (IIS) ServiceControl Protocol | [github.com/oiweiwei/msrpc/iiss](./msrpc/dcom/iiss) |
| [MS-IMSA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-imsa) | Internet Information Services (IIS) IMSAdminBaseW Remote Protocol | [github.com/oiweiwei/msrpc/imsa](./msrpc/dcom/imsa) |
| [MS-IOI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ioi) | IManagedObject Interface Protocol | [github.com/oiweiwei/msrpc/ioi](./msrpc/dcom/ioi) |
| [MS-OAUT](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-oaut) | OLE Automation Protocol | [github.com/oiweiwei/msrpc/oaut](./msrpc/dcom/oaut) |
| [MS-OCSPA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-ocspa) | Microsoft OCSP Administration Protocol | [github.com/oiweiwei/msrpc/ocspa](./msrpc/dcom/ocspa) |
| [MS-PLA](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-pla) | Performance Logs and Alerts Protocol | [github.com/oiweiwei/msrpc/pla](./msrpc/dcom/pla) |
| [MS-RAI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rai) | Remote Assistance Initiation Protocol | [github.com/oiweiwei/msrpc/rai](./msrpc/dcom/rai) |
| [MS-RDPESC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rdpesc) | Remote Desktop Protocol: Smart Card Virtual Channel Extension | [github.com/oiweiwei/msrpc/rdpesc](./msrpc/dcom/rdpesc) |
| [MS-RRASM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rrasm) | Routing and Remote Access Server (RRAS) Management Protocol | [github.com/oiweiwei/msrpc/rrasm](./msrpc/dcom/rrasm) |
| [MS-RSMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-rsmp) | Removable Storage Manager (RSM) Remote Protocol | [github.com/oiweiwei/msrpc/rsmp](./msrpc/dcom/rsmp) |
| [MS-SCMP](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-scmp) | Shadow Copy Management Protocol | [github.com/oiweiwei/msrpc/scmp](./msrpc/dcom/scmp) |
| [MS-TPMVSC](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-tpmvsc) | Trusted Platform Module (TPM) Virtual Smart Card Remote Protocol | [github.com/oiweiwei/msrpc/tpmvsc](./msrpc/dcom/tpmvsc) |
| [MS-UAMG](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-uamg) | Update Agent Management Protocol | [github.com/oiweiwei/msrpc/uamg](./msrpc/dcom/uamg) |
| [MS-VDS](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-vds) | Virtual Disk Service (VDS) Protocol | [github.com/oiweiwei/msrpc/vds](./msrpc/dcom/vds) |
| [MS-WCCE](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wcce) | Windows Client Certificate Enrollment Protocol | [github.com/oiweiwei/msrpc/wcce](./msrpc/dcom/wcce) |
| [MS-WMIO](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmio) | Windows Management Instrumentation Encoding Version 1.0 Protocol | [github.com/oiweiwei/msrpc/wmio](./msrpc/dcom/wmio) |
| [MS-WMI](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wmi) | Windows Management Instrumentation Remote Protocol | [github.com/oiweiwei/msrpc/wmi](./msrpc/dcom/wmi) |
| [MS-WSRM](https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-wsrm) | Windows System Resource Manager (WSRM) Protocol | [github.com/oiweiwei/msrpc/wsrm](./msrpc/dcom/wsrm) |

### Documentation

The codegen package also generates the documentation for the generated code
pulled from the MSDN portal. (it can be quite inaccurate with determining
general comment boundaries vs actual field descriptions, so inaccurate can
be an HTML on MSDN side).

### Naming

The `codegen/go_names` contains the ad-hoc naming engine, which sometimes
quite sucks (so does the overall naming convention in IDL documents, seriously,
how much time the average microsoft developer saves by writing `para` instead of `param`),
but for most of the situations, provide a way to generate the
names that comply with golang naming convention and give more intuition behind
this or that field.

### Generated Stubs

## MIDL Implementation Limitations

 * L.0001: `#define` statements are applicable only for constant declaration;

 * L.0002: `cpp_quote` contents are limited only for constant declaration;

 * L.0005: `int const` declaration is not supported.

 * L.0006: `wchar_t`, `status_error_t` are predefined.

# TODO

 * Testing (I don't have much time)

 * Handle reserved arguments/structure fields used for `switch_is` and `size_is` statements.

 * Derive the type from field name, like `^f[A-Z]` -> `boolean`.

 * Callbacks Support / Server-Side Support

 * Static strings

 * Investigate: Association Group ID is not shared across several named pipe connections. (each NP requires dedicated connection).

 * Convenient way to combine SPNEGO and NTLM/KRB5 within connection option.

# Open Questions

 * Why IObjectExporter does not support NDR64?

 * Why server returns indistinguishable pointers for NDR64?

 * Why SMB2 does not support certain auth levels (ie Winreg supports only Insecure and Privacy)?

# References

Without these projects, it would be absolutely impossible to implement go-msrpc.

 * [Samba](https://www.samba.org/)

 * [Impacket](https://github.com/fortra/impacket)

 * [go-smb2](https://github.com/hirochachacha/go-smb2)

 * [go-krb5](https://github.com/jcmturner/gokrb5)

# Collaboration

Don't hesitate to raise an issues (and only then raise a PR), the project is quite raw,
and I don't have much time, so, a lot of errors and issues are yet to discover.
