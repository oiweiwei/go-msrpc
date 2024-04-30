# GSS-API

## Capabilities

```
capabilities (
    delegation BOOLEAN,
    mutual_authentication BOOLEAN,
    replay_detection BOOLEAN,
    sequencing BOOLEAN,
    anonymity BOOLEAN,
    confidentiality BOOLEAN /* optional */,
    itegrity BOOLEAN /* optional */,
)
```

## GSS Sec Context


## GSS Init Sec Context Call

Inputs:

```
(
    claimant_cred_handle CREDENTIAL HANDLE DEFAULT NULL /* use default */,
    input_context_handle CONTEXT HANDLE DEFAULT GSS_C_NO_CONTEXT /* not yet assigned */, 
    targ_name INTERNAL NAME,
    mech_type OBJECT IDENTIFIER DEFAULT NULL /* use default */, 
    req_capabilities CAPABILITY,
    lifetime_req INTEGER DEFAULT 0 /* use default */,
    chan_bindings OCTET STRING,
    input_token OCTET STRING DEFAULT NULL /* null or token from target */, 
)
```

Outputs:

```
(
    major_status INTEGER,
    minor_status INTEGER,
    output_context_handle CONTEXT HANDLE,
    mech_type OBJECT IDENTIFIER /* never NULL */,
    output_token OCTET STRING /* NULL or token to pass to context target */,
    capabilities CAPABILITIES /* replaces req_* flags */,
    prot_ready_state BOOLEAN,
    lifetime_rec INTEGER /* in seconds, or reserved value for INDEFINITE */,
)
```
