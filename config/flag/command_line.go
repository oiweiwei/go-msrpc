package command_line

import (
	"flag"
	"os"

	"github.com/oiweiwei/go-msrpc/config"
)

func BindFlags(c *config.Config, flagSet *flag.FlagSet) {

	flagSet.BoolVar(&c.Debug, "debug", c.Debug, "enable debug output")

	flagSet.StringVar(&c.Server, "server", c.Server, "server to connect to")
	flagSet.StringVar(&c.Domain, "domain", c.Domain, "domain to authenticate to")
	flagSet.StringVar(&c.Username, "username", c.Username, "username to authenticate as")
	flagSet.StringVar(&c.Workstation, "workstation", c.Workstation, "workstation to authenticate from")

	flagSet.DurationVar(&c.Timeout, "timeout", c.Timeout, "timeout")

	flagSet.StringVar(&c.TrasnferEncoding, "transfer-encoding", c.TrasnferEncoding, "transfer encoding: ndr20, ndr64")

	flagSet.StringVar(&c.Credential.Password, "password", c.Credential.Password, "password to authenticate with")
	flagSet.StringVar(&c.Credential.NTHash, "nthash", c.Credential.NTHash, "NT hash to authenticate with")
	flagSet.StringVar(&c.Credential.MachineAccountPassword, "machine-account-password", c.Credential.MachineAccountPassword, "machine account password to authenticate with")
	flagSet.StringVar(&c.Credential.MachineAccountNTHash, "machine-account-nthash", c.Credential.MachineAccountNTHash, "machine account NT hash to authenticate with")
	flagSet.IntVar(&c.Credential.EncryptionKey.KeyType, "encryption-key-type", c.Credential.EncryptionKey.KeyType, "encryption key type: rc4 (23), aes128 (17), aes256 (18)")
	flagSet.StringVar(&c.Credential.EncryptionKey.KeyValue, "encryption-key-value", c.Credential.EncryptionKey.KeyValue, "encryption key")

	flagSet.StringVar(&c.Auth.Level, "auth-level", c.Auth.Level, "authentication level: none, connect, call, pkt, integrity, privacy")
	flagSet.Var(&c.Auth.Types, "auth-type", "authentication type: ntlm, krb5")
	flagSet.StringVar(&c.Auth.TargetName, "target-name", c.Auth.TargetName, "target name")
	flagSet.BoolVar(&c.Auth.SPNEGO, "auth-spnego", c.Auth.SPNEGO, "use spnego")
	flagSet.StringVar(&c.Auth.Impersonation, "impersonation", c.Auth.Impersonation, "impersonation level: anonymous, identify, impersonate, delegate")
	flagSet.BoolVar(&c.Auth.NoHeaderSign, "auth-no-header-sign", false, "do not sign headers")
	flagSet.StringVar(&c.Auth.KRB5.ConfigFile, "krb5-config-file", c.Auth.KRB5.ConfigFile, "path to krb5.conf")
	flagSet.StringVar(&c.Auth.KRB5.KDCServer, "krb5-kdc-server", c.Auth.KRB5.KDCServer, "KDC server to authenticate to")
	flagSet.StringVar(&c.Auth.KRB5.AdminServer, "krb5-admin-server", c.Auth.KRB5.AdminServer, "admin server to authenticate to")
	flagSet.StringVar(&c.Auth.KRB5.Keytab, "krb5-keytab-path", c.Auth.KRB5.Keytab, "path to keytab")
	flagSet.StringVar(&c.Auth.KRB5.CCache, "krb5-ccache-path", c.Auth.KRB5.CCache, "path to ccache")
	flagSet.Var(&c.Auth.KRB5.EncryptionTypes, "krb5-encryption-types", "encryption types to use: aes256-cts-hmac-sha1-96, aes128-cts-hmac-sha1-96, arcfour-hmac-md5")
	flagSet.BoolVar(&c.Auth.KRB5.DCEStyle, "krb5-dce-style", c.Auth.KRB5.DCEStyle, "use DCE style")
	flagSet.BoolVar(&c.Auth.KRB5.DisablePAFXFAST, "krb5-disable-pafx-fast", c.Auth.KRB5.DisablePAFXFAST, "disable PA-FX-FAST")
	flagSet.BoolVar(&c.Auth.KRB5.MutualAuthn, "krb5-mutual-authn", c.Auth.KRB5.MutualAuthn, "use mutual authentication")
	flagSet.BoolVar(&c.Auth.KRB5.AnyServiceClassSPN, "krb5-any-service-class-spn", c.Auth.KRB5.AnyServiceClassSPN, "use any service class SPN")

	flagSet.BoolVar(&c.Auth.NTLM.NTLMv1, "ntlm-v1", c.Auth.NTLM.NTLMv1, "use NTLMv1")
	flagSet.BoolVar(&c.Auth.NTLM.NoESS, "ntlm-no-ess", c.Auth.NTLM.NoESS, "use no extended session security")

	flagSet.BoolVar(&c.Verify.Presentation, "verify-presentation", false, "verify presentation")
	flagSet.BoolVar(&c.Verify.Header2, "verify-header2", false, "verify header2")
	flagSet.BoolVar(&c.Verify.BitMask, "verify-bitmask", false, "verify bitmask")

	flagSet.IntVar(&c.SMB.Port, "smb-port", 445, "SMB port")
	flagSet.BoolVar(&c.SMB.Sign, "smb-sign", false, "SMB signing")
	flagSet.BoolVar(&c.SMB.Seal, "smb-seal", false, "SMB sealing")
	flagSet.StringVar(&c.SMB.Dialect, "smb-dialect", c.SMB.Dialect, "SMB dialect: 2.0.2 (202), 2.1.0 (210), 3.0.0 (300), 3.0.2 (302), 3.1.1 (311)")

	flagSet.BoolVar(&c.EPM.Enabled, "epm", c.EPM.Enabled, "use endpoint mapper")
	flagSet.StringVar(&c.EPM.AuthLevel, "epm-auth-level", c.EPM.AuthLevel, "endpoint mapper authentication level: none, connect, call, pkt, integrity, privacy")

	flagSet.StringVar(&c.Protocol, "protocol", c.Protocol, "protocol to use, ncacn_np (smb), ncacn_ip_tcp (tcp)")
}

func ParseAndValidate(cfg *config.Config, flagSet *flag.FlagSet) error {

	flagSet.Parse(os.Args[1:])

	if cfg.Server == "" && flagSet.NArg() > 0 {
		cfg.Server = flagSet.Arg(0)
		flagSet.Parse(flagSet.Args()[1:])
	}

	return cfg.Validate()
}
