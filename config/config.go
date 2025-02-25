package config

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	krb5_config "github.com/oiweiwei/gokrb5.fork/v9/config"
	"github.com/oiweiwei/gokrb5.fork/v9/iana/etypeID"
	zerolog "github.com/rs/zerolog"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/smb2"

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"
	"github.com/oiweiwei/go-msrpc/ssp/krb5"
	"github.com/oiweiwei/go-msrpc/ssp/ntlm"

	"github.com/oiweiwei/go-msrpc/msrpc/well_known"

	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
)

type StringSlice []string

func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSlice) Set(value string) error {
	if value != "" {
		*s = append(*s, strings.Split(value, ",")...)
	}
	return nil
}

// Config struct represents the MSRPC configuration.
type Config struct {
	// Debug enables debug logging.
	Debug bool `json:"debug"`
	// The debug logger.
	Logger zerolog.Logger `json:"-"`
	// The server IP, or FQDN or binding string.
	Server string `json:"server"`
	// Server address.
	ServerAddress string `json:"server_address"`
	// The domain to connect to.
	Domain string `json:"domain"`
	// The username to connect with.
	Username string `json:"username"`
	// The workstation to connect from.
	Workstation string `json:"workstation"`
	// The timeout for the connection.
	Timeout time.Duration `json:"timeout"`

	// The credential configuration.
	Credential struct {
		// The password to use.
		Password string `json:"password"`
		// The NT hash to use.
		NTHash string `json:"nt_hash"`
		// The machine account password to use.
		MachineAccountPassword string `json:"machine_account_password"`
		// The machine account NT hash.
		MachineAccountNTHash string `json:"machine_account_nt_hash"`
		EncryptionKey        struct {
			KeyType  int    `json:"key_type"`
			KeyValue string `json:"key_value"`
		}
	} `json:"credential"`

	// The auth configuration.
	Auth struct {
		// The auth level to use. (none, connect, call, pkt, integrity, privacy)
		Level string `json:"level"`
		// The impersonation level to use (anonymous, identify, impersonate, delegate). (default is impersonate)
		Impersonation string `json:"impersonation"`
		// The auth type to use. (ntlm, krb5)
		Types StringSlice `json:"type"`
		// The target name to use.
		TargetName string `json:"target_name"`
		// The flag that indicates whether the SPNEGO should be used.
		SPNEGO bool `json:"spnego"`
		// The flag that indicates whether the header sign should be disabled.
		NoHeaderSign bool `json:"no_header_sign"`

		// The auth configuration for KRB5.
		KRB5 struct {
			// The path to the krb5.conf file.
			ConfigFile string `json:"config_file_path"`
			// The KDC server to connect to.
			KDCServer string `json:"kdc_server,omitempty"`
			// The admin server to connect to.
			AdminServer string `json:"admin_server,omitempty"`
			// The encryption types to use.
			EncryptionTypes StringSlice `json:"encryption_types"`
			// The path to the keytab file.
			Keytab string `json:"keytab_path"`
			// The path to the ccache file.
			CCache string `json:"ccache_path"`
			// The flag that indicates whether the 3-leg DCE authentication
			// should be used. (default true)
			DCEStyle bool `json:"dce_style"`
			// The flag that indicates whether the PAF-X FAST should be
			// disabled. (default true)
			DisablePAFXFAST bool `json:"disable_pafx_fast"`
			// The flag that indicates whether the mutual authentication
			// is required.
			MutualAuthn bool `json:"mutual_authn"`
		} `json:"auth_krb5_config"`

		// The auth configuration for NTLM.
		NTLM struct {
			// The flag that indicates whether the NTLMv1 should be used.
			// (default false)
			NTLMv1 bool `json:"ntlm_v1"`
			// The flag that indicates whether the Extended Session Security
			// should be used. (default true)
			NoESS bool `json:"no_ess"`
		} `json:"auth_ntlm_config"`
	} `json:"auth"`

	// The verification trailer configuration.
	Verify struct {
		// The flag that indicates whether the to include presentation
		// verification command.
		Presentation bool `json:"presentation"`
		// The flag that indicates whether the to include header2
		// verification command.
		Header2 bool `json:"header2"`
		// The flag that indicates whether the to include bit mask
		// verification command.
		BitMask bool `json:"bitmask"`
	} `json:"verify"`

	// The SMB2 configuration.
	SMB struct {
		// The port to connect to. (default is 445)
		Port int `json:"port"`
		// The flag that indicates whether the message should be signed.
		Sign bool `json:"sign"`
		// The flag that indicates whether the message should be encrypted.
		Seal bool `json:"seal"`
		// The SMB2/3 dialect to use.
		Dialect string `json:"dialect"`
	} `json:"smb"`

	// The Endpoint Mapper configuration.
	EPM struct {
		// The flag that indicates whether the EPM should be enabled.
		Enabled bool `json:"enabled"`
		// The auth level to use. (none, connect, call, pkt, integrity, privacy)
		// By default auth level is inherited from the Auth.Level.
		AuthLevel string `json:"auth_level"`
	} `json:"epm"`

	// The protocol to use. (ncacn_np or smb, ncacn_ip_tcp or tcp)
	Protocol string `json:"protocol"`

	// The transfer encoding to use (ndr20, ndr64)
	TrasnferEncoding string `json:"transfer_encoding"`

	// The flag that indicates whether credentials and mechanisms should be
	// included into connection options. If GlobalCredentials is true, then
	// credentials and mechanisms are not included into connection options.
	useGlobalCredentials bool

	// The flag that indicates whether machine account credentials should be
	// used.
	useMachineAccount bool

	// The flag that indicates whether Netlogon SSP should be used.
	useNetlogonSSP bool
}

// New function returns a new configuration.
func New() *Config {

	cfg := &Config{}

	cfg.Timeout = 30 * time.Second

	cfg.Auth.Types = []string{"ntlm"}
	cfg.Auth.Level = "connect"
	cfg.Auth.Impersonation = "impersonate"
	cfg.Auth.SPNEGO = true

	cfg.EPM.Enabled = true
	cfg.EPM.AuthLevel = "none"

	cfg.Auth.KRB5.DCEStyle = true
	cfg.Auth.KRB5.DisablePAFXFAST = true

	cfg.TrasnferEncoding = "ndr20"

	return cfg
}

// UseGlobalCredentials function sets the flag that indicates whether
// credentials and mechanisms should be included into connection options.
// If GlobalCredentials is true, then credentials and mechanisms are not
// included into connection options.
// If GlobalCredentials is false, then credentials and mechanisms are
// included into connection options and credentials and mechanisms should
// be exported in code.
func (cfg *Config) UseGlobalCredentials() *Config {
	cfg.useGlobalCredentials = true
	return cfg
}

// UseMachineAccount function sets the flag that indicates whether machine
// account credentials only should be used.
func (cfg *Config) UseMachineAccount() *Config {
	cfg.useMachineAccount = true
	return cfg
}

// UseNetlogonSSP function sets the flag that indicates whether Netlogon SSP
// should be used.
func (cfg *Config) UseNetlogonSSP() *Config {
	cfg.useNetlogonSSP = true
	return cfg
}

// DisableEPM function disables the Endpoint Mapper.
func (cfg *Config) DisableEPM() *Config {
	cfg.EPM.Enabled = false
	return cfg
}

// NTLM function returns the NTLM configuration.
func (cfg *Config) NTLM() *ntlm.Config {

	ncfg := ntlm.NewConfig()

	if cfg.Auth.NTLM.NTLMv1 {
		ncfg.NTLMVersion = ntlm.NTLMv1
	} else {
		ncfg.NTLMVersion = ntlm.NTLMv2
	}

	if cfg.Auth.NTLM.NoESS {
		ncfg.NoESS = cfg.Auth.NTLM.NoESS
	}

	return ncfg
}

// KRB5 function returns the KRB5 configuration.
func (cfg *Config) KRB5() *krb5.Config {

	kcfg := krb5.NewConfig()

	kcfg.DCEStyle = cfg.Auth.KRB5.DCEStyle
	kcfg.DisablePAFXFAST = cfg.Auth.KRB5.DisablePAFXFAST

	if cfg.Auth.KRB5.MutualAuthn {
		kcfg.Flags = append(kcfg.Flags, int(gssapi.MutualAuthn))
	}

	if cfg.Auth.KRB5.CCache != "" {
		kcfg.CCachePath = cfg.Auth.KRB5.CCache
	}

	if cfg.Auth.KRB5.ConfigFile != "" {
		kcfg.KRB5ConfigPath = cfg.Auth.KRB5.ConfigFile
		return kcfg
	}

	if cfg.Auth.KRB5.KDCServer == "" {
		cfg.Auth.KRB5.KDCServer = cfg.Server
	}

	kcfg.KRB5Config, _ = cfg.GenKRB5Config()

	return kcfg
}

// Mechanisms function returns the set of mechanisms.
// If GlobalCredentials is true, then mechanisms are not included into
// connection options. You should manually call gssapi.AddMechanism
// function to add mechanisms into global security context.
func (cfg *Config) Mechanisms() []gssapi.MechanismFactory {

	mechanisms := []gssapi.MechanismFactory{}

	if cfg.Auth.SPNEGO {
		mechanisms = append(mechanisms, ssp.SPNEGO)
	}

	for _, typ := range cfg.Auth.Types {
		switch typ {
		case "ntlm":
			mechanisms = append(mechanisms, gssapi.WithDefaultConfig(ssp.NTLM, cfg.NTLM()))
		case "krb5":
			mechanisms = append(mechanisms, gssapi.WithDefaultConfig(ssp.KRB5, cfg.KRB5()))
		}
	}

	return mechanisms
}

// MachineAccountCredentials function returns the set of machine account
// credentials.
func (cfg *Config) MachineAccountCredentials() []credential.Credential {

	creds := []credential.Credential{}

	if cfg.Credential.MachineAccountPassword != "" {
		creds = append(creds, credential.NewFromPassword(cfg.Username, cfg.Credential.MachineAccountPassword,
			credential.Workstation(cfg.Workstation)))
	}

	if cfg.Credential.MachineAccountNTHash != "" {
		creds = append(creds, credential.NewFromNTHash(cfg.Username, cfg.Credential.MachineAccountNTHash,
			credential.Workstation(cfg.Workstation)))
	}

	if cfg.Credential.EncryptionKey.KeyValue != "" {
		creds = append(creds, credential.NewFromEncryptionKey(cfg.Username, cfg.Credential.EncryptionKey.KeyType,
			cfg.Credential.EncryptionKey.KeyValue, credential.Workstation(cfg.Workstation)))
	}

	return creds
}

// Credentials function returns the set of credentials.
// If GlobalCredentials is true, then credentials are not included into
// connection options. You should manually call dcerpc.WithCredentials or
// gssapi.AddCredential function to add credentials into global security
// context.
func (cfg *Config) Credentials() []credential.Credential {

	creds := []credential.Credential{}

	if cfg.Credential.Password != "" {
		creds = append(creds, credential.NewFromPassword(cfg.Username, cfg.Credential.Password,
			credential.Workstation(cfg.Workstation)))
	}

	if cfg.Credential.NTHash != "" {
		creds = append(creds, credential.NewFromNTHash(cfg.Username, cfg.Credential.NTHash,
			credential.Workstation(cfg.Workstation)))
	}

	if cfg.Auth.KRB5.Keytab != "" {
		creds = append(creds, credential.NewFromKeytabFile(cfg.Username, cfg.Auth.KRB5.Keytab))
	}

	creds = append(creds, cfg.MachineAccountCredentials()...)

	if len(creds) == 0 {
		creds = append(creds, credential.Anonymous())
	}

	return creds
}

// ClientOptions function returns the set of client options.
func (cfg *Config) ClientOptions(ctx context.Context) []dcerpc.Option {

	options := []dcerpc.Option{}

	switch cfg.TrasnferEncoding {
	case "ndr20":
		options = append(options, dcerpc.WithNDR20())
	case "ndr64":
		options = append(options, dcerpc.WithNDR64())
	}

	switch cfg.Auth.Level {
	case "none":
		options = append(options, dcerpc.WithInsecure())
	case "connect":
		options = append(options, dcerpc.WithConnect())
	case "call":
		options = append(options, dcerpc.WithSecurityLevel(dcerpc.AuthLevelCall))
	case "pkt":
		options = append(options, dcerpc.WithSecurityLevel(dcerpc.AuthLevelPkt))
	case "integrity":
		options = append(options, dcerpc.WithSign())
	case "privacy":
		options = append(options, dcerpc.WithSeal())
	}

	if !cfg.useGlobalCredentials {

		if cfg.Auth.SPNEGO {
			options = append(options, dcerpc.WithMechanism(ssp.SPNEGO))
		}

		for _, typ := range cfg.Auth.Types {
			switch typ {
			case "ntlm":
				options = append(options, dcerpc.WithMechanism(ssp.NTLM, cfg.NTLM()))
			case "krb5":
				options = append(options, dcerpc.WithMechanism(ssp.KRB5, cfg.KRB5()))
			}
		}

		if cfg.useNetlogonSSP {
			options = append(options, dcerpc.WithMechanism(ssp.Netlogon))
		}
	}

	if cfg.Auth.TargetName != "" {
		options = append(options, dcerpc.WithTargetName(cfg.Auth.TargetName))
	}

	if cfg.Verify.BitMask {
		options = append(options, dcerpc.WithVerifyBitMask(true))
	}

	if cfg.Verify.Header2 {
		options = append(options, dcerpc.WithVerifyHeader2(true))
	}

	if cfg.Verify.Presentation {
		options = append(options, dcerpc.WithVerifyPresentation(true))
	}

	if !cfg.useGlobalCredentials {
		if cfg.useMachineAccount {
			for _, cred := range cfg.MachineAccountCredentials() {
				options = append(options, dcerpc.WithCredentials(cred))
			}
		} else {
			for _, cred := range cfg.Credentials() {
				options = append(options, dcerpc.WithCredentials(cred))
			}
		}
	}

	return options

}

// epmOptions ...
func (cfg *Config) EPMOptions(ctx context.Context) []dcerpc.Option {

	options := []dcerpc.Option{}

	authLevel := cfg.EPM.AuthLevel
	if authLevel == "" {
		authLevel = cfg.Auth.Level
	}

	switch authLevel {
	case "none":
		options = append(options, dcerpc.WithInsecure())
	case "connect":
		options = append(options, dcerpc.WithConnect())
	case "call":
		options = append(options, dcerpc.WithSecurityLevel(dcerpc.AuthLevelCall))
	case "pkt":
		options = append(options, dcerpc.WithSecurityLevel(dcerpc.AuthLevelPkt))
	case "integrity":
		options = append(options, dcerpc.WithSign())
	case "privacy":
		options = append(options, dcerpc.WithSeal())
	}

	return append(append(cfg.ClientOptions(ctx), cfg.getDialOptions()...), options...)
}

func (cfg *Config) Log() {

	log := zerolog.New(os.Stderr)

	if cfg.Debug {
		b, err := json.Marshal(cfg)
		if err != nil {
			log.Error().Err(err).Msg("config error")
		} else {
			log.Debug().RawJSON("config", b).Msg("config")
		}
	}
}

func (cfg *Config) getDialOptions() []dcerpc.Option {

	options := []dcerpc.Option{}

	if cfg.Debug {
		options = append(options, dcerpc.WithLogger(zerolog.New(os.Stderr)))
	}

	if cfg.Timeout != 0 {
		options = append(options, dcerpc.WithTimeout(cfg.Timeout))
	}

	if cfg.SMB.Port != 0 {
		options = append(options, dcerpc.WithSMBPort(cfg.SMB.Port))
	}

	if cfg.Auth.NoHeaderSign {
		options = append(options, dcerpc.NoHeaderSign())
	}

	if dialer := cfg.SMBDialerOptions(); len(dialer) > 0 {
		options = append(options, dcerpc.WithSMBDialer(smb2.NewDialer(dialer...)))
	}

	if !cfg.useGlobalCredentials {
		if cfg.useMachineAccount {
			for _, cred := range cfg.MachineAccountCredentials() {
				options = append(options, dcerpc.WithCredentials(cred))
			}
		} else {
			for _, cred := range cfg.Credentials() {
				options = append(options, dcerpc.WithCredentials(cred))
			}
		}
	}

	return options
}

// ServerAddr returns the server address (with optional binding).
func (cfg *Config) ServerAddr() string {

	switch cfg.Protocol {
	case "ncacn_np", "smb":
		return "ncacn_np:" + cfg.ServerAddress
	case "ncacn_ip_tcp", "tcp":
		return "ncacn_ip_tcp:" + cfg.ServerAddress
	}

	return cfg.Server
}

// DialOptions function returns the set of connection options.
func (cfg *Config) DialOptions(ctx context.Context) []dcerpc.Option {
	if cfg.EPM.Enabled {
		return append(cfg.getDialOptions(), epm.EndpointMapper(ctx, cfg.ServerAddr(), cfg.EPMOptions(ctx)...))
	} else {
		return append(cfg.getDialOptions(), well_known.EndpointMapper())
	}
}

// SMBDialerOptions function returns the set of SMB dialer options.
func (cfg *Config) SMBDialerOptions() []smb2.DialerOption {

	options := []smb2.DialerOption{}

	gssOptions := []gssapi.ContextOption{}

	if !cfg.useGlobalCredentials {

		if cfg.Auth.SPNEGO {
			gssOptions = append(gssOptions, gssapi.WithMechanismFactory(ssp.SPNEGO))
		}

		for _, typ := range cfg.Auth.Types {
			switch typ {
			case "ntlm":
				gssOptions = append(gssOptions, gssapi.WithMechanismFactory(ssp.NTLM, cfg.NTLM()))
			case "krb5":
				gssOptions = append(gssOptions, gssapi.WithMechanismFactory(ssp.KRB5, cfg.KRB5()))
			}
		}

	}

	if cfg.Auth.TargetName != "" {
		gssOptions = append(gssOptions, gssapi.WithTargetName(cfg.Auth.TargetName))
	}

	if !cfg.useGlobalCredentials {
		for _, cred := range cfg.Credentials() {
			gssOptions = append(gssOptions, gssapi.WithCredential(cred))
		}
	}

	options = append(options, smb2.WithSecurity(gssOptions...))

	if cfg.SMB.Sign {
		options = append(options, smb2.WithSign())
	}

	if cfg.SMB.Seal {
		options = append(options, smb2.WithSeal())
	}

	if cfg.SMB.Dialect != "" {
		switch strings.ReplaceAll(strings.TrimPrefix(cfg.SMB.Dialect, "smb"), ".", "") {
		case "202":
			options = append(options, smb2.WithDialect(smb2.SMB202))
		case "210":
			options = append(options, smb2.WithDialect(smb2.SMB210))
		case "300":
			options = append(options, smb2.WithDialect(smb2.SMB300))
		case "302":
			options = append(options, smb2.WithDialect(smb2.SMB302))
		case "311":
			options = append(options, smb2.WithDialect(smb2.SMB311))
		}
	}

	return options
}

// ValidateAuthLevel function validates the auth level.
func ValidateAuthLevel(level string) error {
	switch level {
	case "none", "connect", "call", "pkt", "integrity", "privacy":
	default:
		return fmt.Errorf("invalid auth level: %s", level)
	}
	return nil
}

func ValidateTransferEncoding(encoding string) error {
	switch encoding {
	case "ndr20", "ndr64":
	default:
		return fmt.Errorf("invalid transfer encoding: %s", encoding)
	}
	return nil
}

func ValidateImpersonation(level string) error {

	switch level {
	case "anonymous", "identify", "impersonate", "delegate":
	default:
		return fmt.Errorf("invalid impersonation level: %s", level)
	}

	return nil
}

func (cfg *Config) ParseServerAddr() error {

	if cfg.Server == "" {
		return fmt.Errorf("server is required")
	}

	cfg.ServerAddress = cfg.Server

	ip, hostName, binding, err := dcerpc.ParseServerAddr(cfg.Server)
	if err != nil {
		return fmt.Errorf("server address: %w", err)
	}

	if hostName != "" && cfg.Auth.TargetName == "" {
		cfg.Auth.TargetName = "host" + "/" + hostName
	}

	if binding != nil {

		if binding.ProtocolSequence != 0 {
			cfg.Protocol = ""
		}

		cfg.Server = binding.String()

		if ip != nil {
			cfg.ServerAddress = ip.String()
		} else if binding.NetworkAddress != "" {
			cfg.ServerAddress = binding.NetworkAddress
		} else if hostName != "" {
			cfg.ServerAddress = hostName
		} else if binding.ComputerName != "" {
			cfg.ServerAddress = binding.ComputerName
		}

		if cfg.Auth.TargetName == "" {
			cfg.Auth.TargetName = "host" + "/" + cfg.ServerAddress
		}

		if binding.Username != "" {
			cfg.Username = binding.Username
		}

		if binding.Password != "" {
			if strings.HasPrefix(binding.Password, "hash:") {
				cfg.Credential.NTHash = strings.TrimPrefix(binding.Password, "hash:")
			} else if strings.HasPrefix(binding.Password, "key:") {
				part := strings.Split(strings.TrimPrefix(binding.Password, "key:"), ":")
				if len(part) == 2 {
					cfg.Credential.EncryptionKey.KeyValue = part[1]
					if itype, err := strconv.Atoi(part[0]); err == nil {
						cfg.Credential.EncryptionKey.KeyType = itype
					} else {
						itype, ok := etypeID.ETypesByName[part[0]]
						if !ok {
							return fmt.Errorf("invalid encryption key type: %s", part[0])
						}
						cfg.Credential.EncryptionKey.KeyType = int(itype)
					}
				} else {
					return fmt.Errorf("invalid encryption key: %s", binding.Password)
				}
			} else {
				cfg.Credential.Password = binding.Password
			}
		}

		if cfg.Auth.KRB5.KDCServer == "" {
			cfg.Auth.KRB5.KDCServer = cfg.ServerAddress
		}

		if cfg.Auth.KRB5.AdminServer == "" {
			cfg.Auth.KRB5.AdminServer = cfg.ServerAddress
		}

		extras := make(map[string]bool)

		for _, extra := range binding.Extra {
			extras[extra] = true
		}

		if extras["krb5"] || extras["ntlm"] {
			// clear the auth type.
			cfg.Auth.SPNEGO = false
			cfg.Auth.Types = []string{}
		}

		for _, extra := range binding.Extra {
			switch extra {
			// transfer encoding.
			case "ndr20":
				cfg.TrasnferEncoding = "ndr20"
			case "ndr64":
				cfg.TrasnferEncoding = "ndr64"
			// auth type keywords.
			case "spnego":
				cfg.Auth.SPNEGO = true
			case "krb5":
				cfg.Auth.Types = append(cfg.Auth.Types, "krb5")
			case "ntlm":
				cfg.Auth.Types = append(cfg.Auth.Types, "ntlm")
			// auth level keywords.
			case "connect":
				cfg.Auth.Level = "connect"
			case "call":
				cfg.Auth.Level = "call"
			case "pkt":
				cfg.Auth.Level = "pkt"
			case "integrity":
				cfg.Auth.Level = "integrity"
			case "privacy":
				cfg.Auth.Level = "privacy"
			// impersonation keywords.
			case "anonymous":
				cfg.Auth.Impersonation = "anonymous"
			case "identify":
				cfg.Auth.Impersonation = "identify"
			case "impersonate":
				cfg.Auth.Impersonation = "impersonate"
			case "delegate":
				cfg.Auth.Impersonation = "delegate"
			// verify keywords.
			case "presentation":
				cfg.Verify.Presentation = true
			case "header2":
				cfg.Verify.Header2 = true
			case "bitmask":
				cfg.Verify.BitMask = true
			// smb2 keywords.
			case "smb2_sign":
				cfg.SMB.Sign = true
			case "smb2_seal":
				cfg.SMB.Seal = true
			}
		}
	}

	return nil
}

// Validate function validates the configuration.
func (cfg *Config) Validate() error {

	if err := cfg.ParseServerAddr(); err != nil {
		return err
	}

	if cfg.Debug {
		cfg.Logger = zerolog.New(os.Stderr).With().Str("source", "config").Logger()
	}

	for _, typ := range cfg.Auth.Types {
		switch typ {
		case "ntlm", "krb5":
		default:
			return fmt.Errorf("invalid auth type: %s", typ)
		}
	}

	if err := ValidateTransferEncoding(cfg.TrasnferEncoding); err != nil {
		return err
	}

	if err := ValidateAuthLevel(cfg.Auth.Level); err != nil {
		return err
	}

	if err := ValidateImpersonation(cfg.Auth.Impersonation); err != nil {
		return err
	}

	if cfg.EPM.Enabled {
		if cfg.EPM.AuthLevel != "" {
			if err := ValidateAuthLevel(cfg.EPM.AuthLevel); err != nil {
				return fmt.Errorf("epm: %w", err)
			}
		}
	}

	for _, typ := range cfg.Auth.Types {
		if typ == "krb5" && cfg.Auth.TargetName != "" && !strings.HasPrefix(cfg.Auth.TargetName, "host/") {
			cfg.Auth.TargetName = "host" + "/" + cfg.Auth.TargetName
		}
	}

	if cfg.Domain == "" {
		if cfg.Domain = credential.DomainName(cfg.Username); cfg.Domain == "" {

			addr := cfg.ServerAddress

			if net.ParseIP(addr) != nil {
				if addr = strings.TrimPrefix(cfg.Auth.TargetName, "host/"); net.ParseIP(addr) != nil {
					addr = ""
				}
			}

			if addr != "" {
				spl := strings.Split(strings.TrimRight(addr, "."), ".")
				if len(spl) > 1 {
					cfg.Domain = strings.ToUpper(strings.Join(spl[len(spl)-2:], "."))
				}
			}
		}
	}

	if cfg.Domain == "" {
		return fmt.Errorf("domain is required")
	}

	for _, typ := range cfg.Auth.Types {
		if typ == "krb5" {
			if len(cfg.Auth.KRB5.EncryptionTypes) == 0 {
				cfg.Auth.KRB5.EncryptionTypes = []string{"aes128-cts-hmac-sha1-96", "aes256-cts-hmac-sha1-96", "arcfour-hmac-md5"}
			}
			if cfg.Auth.KRB5.ConfigFile == "" {
				if _, err := cfg.GenKRB5Config(); err != nil {
					return err
				}
			}
		}
	}

	if cfg.Username != "" && credential.DomainName(cfg.Username) == "" {
		cfg.Username = cfg.Domain + "\\" + cfg.Username
	}

	if strings.HasSuffix(cfg.Username, "$") && cfg.Workstation == "" {
		cfg.Workstation = strings.TrimPrefix(strings.TrimSuffix(cfg.Username, "$"), credential.DomainName(cfg.Username)+"\\")
	}

	if cfg.Auth.TargetName == "" {
		cfg.Auth.TargetName = cfg.Server
	}

	if !strings.Contains(cfg.Auth.TargetName, "/") {
		cfg.Auth.TargetName = "host" + "/" + cfg.Auth.TargetName
	}

	if cfg.Protocol != "" {
		switch cfg.Protocol {
		case "ncacn_np", "ncacn_ip_tcp", "smb", "tcp":
		default:
			return fmt.Errorf("invalid protocol: %s", cfg.Protocol)
		}
	}

	cfg.Log()

	return nil
}

func (cfg *Config) GenKRB5Config() (*krb5_config.Config, error) {

	tmp := &strings.Builder{}

	if err := KRB5ConfigT.Execute(tmp, cfg); err != nil {
		return nil, fmt.Errorf("krb5 config error: execute template: %v", err)
	}

	krb5Config, err := krb5_config.NewFromString(tmp.String())
	if err != nil {
		return nil, fmt.Errorf("krb5 config error: parse config: %v", err)
	}

	return krb5Config, nil
}

var KRB5ConfigT = template.Must(template.New("krb5-config").Parse(`
[realms]
{{ .Domain }} = {
{{- if .Auth.KRB5.KDCServer }}
	kdc = {{ .Auth.KRB5.KDCServer }}
{{- end }}
{{- if .Auth.KRB5.AdminServer }}
	admin_server = {{ .Auth.KRB5.AdminServer }}
{{- end }}
}

[libdefaults]
allow_weak_crypto = true
default_realm = {{ .Domain }}
default_tkt_enctypes = {{ range $encType := .Auth.KRB5.EncryptionTypes }}{{ $encType }} {{ end }}
default_tgs_enctypes = {{ range $encType := .Auth.KRB5.EncryptionTypes }}{{ $encType }} {{ end }}
permitted_enctypes = {{ range $encType := .Auth.KRB5.EncryptionTypes }}{{ $encType }} {{ end }}
`))
