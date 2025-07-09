//go:build exclude

// rrp.go script dumps the windows registry keys on the remote machine (very slowly).
// NOTE that SMB cannot use security other than WithInsecure() and WithSeal().
package main

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/rrp/winreg/v1"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

var (
	cfg       = config.New().DisableEPM()
	keyName   string
	recurisve bool
	hive      string
)

func init() {
	config_flag.BindFlags(cfg, flag.CommandLine)
	flag.CommandLine.StringVar(&keyName, "key", "", "key to open")
	flag.CommandLine.BoolVar(&recurisve, "r", false, "recursively enumerate keys")
	flag.CommandLine.StringVar(&hive, "hive", "HKEY_CURRENT_USER", "registry hive to open (HKEY_LOCAL_MACHINE, HKEY_CURRENT_USER, HKEY_CLASSES_ROOT, HKEY_USERS, HKEY_CURRENT_CONFIG)")
}

func OpenHive(ctx context.Context, cli winreg.WinregClient, hive string) (*winreg.Key, error) {

	access := uint32(winreg.KeyQueryValue | winreg.KeyEnumerateSubKeys)

	switch hive {
	case winreg.KeyLocalMachine:
		resp, err := cli.OpenLocalMachine(ctx, &winreg.OpenLocalMachineRequest{
			DesiredAccess: access,
		})
		if err != nil {
			return nil, fmt.Errorf("open local machine: %v", err)
		}

		return resp.Key, nil
	case winreg.KeyCurrentUser:
		resp, err := cli.OpenCurrentUser(ctx, &winreg.OpenCurrentUserRequest{
			DesiredAccess: access,
		})
		if err != nil {
			return nil, fmt.Errorf("open current user: %v", err)
		}
		return resp.Key, nil
	case winreg.KeyClassesRoot:
		resp, err := cli.OpenClassesRoot(ctx, &winreg.OpenClassesRootRequest{
			DesiredAccess: access,
		})
		if err != nil {
			return nil, fmt.Errorf("open classes root: %v", err)
		}
		return resp.Key, nil
	case winreg.KeyUsers:
		resp, err := cli.OpenUsers(ctx, &winreg.OpenUsersRequest{
			DesiredAccess: access,
		})
		if err != nil {
			return nil, fmt.Errorf("open users: %v", err)
		}
		return resp.Key, nil
	case winreg.KeyCurrentConfig:
		resp, err := cli.OpenCurrentConfig(ctx, &winreg.OpenCurrentConfigRequest{
			DesiredAccess: access,
		})
		if err != nil {
			return nil, fmt.Errorf("open current config: %v", err)
		}
		return resp.Key, nil
	default:
		return nil, fmt.Errorf("unknown hive: %s", hive)
	}
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	hive, err := winreg.ParseRegistryHiveName(hive)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse hive", hive, err)
		os.Exit(1)
	}

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, cfg.ServerAddr(), cfg.DialOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cli, err := winreg.NewWinregClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cli", err)
		os.Exit(1)
	}

	key, err := OpenHive(ctx, cli, hive)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	kS := "\\\\" + hive

	if keyName != "" {

		keyName = strings.ReplaceAll(keyName, "/", "\\")

		sub, err := cli.BaseRegOpenKey(ctx, &winreg.BaseRegOpenKeyRequest{
			Key: key,
			SubKey: &winreg.UnicodeString{
				Buffer: keyName + "\x00",
			},
			DesiredAccess: winreg.KeyQueryValue | winreg.KeyEnumerateSubKeys,
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, "open key", keyName, err)
			os.Exit(1)
		}

		if err := Enumerate(ctx, cli, 0, kS+"\\"+keyName, sub.ResultKey); err != nil {
			fmt.Fprintln(os.Stderr, "key", err)
			os.Exit(1)
		}

	} else {
		if err := Enumerate(ctx, cli, 0, kS, key); err != nil {
			fmt.Fprintln(os.Stderr, "key", err)
			os.Exit(1)
		}
	}
}

func Enumerate(ctx context.Context, cli winreg.WinregClient, i int, kS string, h *winreg.Key) error {

	info, err := cli.BaseRegQueryInfoKey(ctx, &winreg.BaseRegQueryInfoKeyRequest{
		Key: h,
		ClassIn: &winreg.UnicodeString{
			MaximumLength: 1024,
		},
	})
	if err != nil {
		return fmt.Errorf("enumerate: query_info: %v", err)
	}

	fmt.Println(fmt.Sprintf("[%d]", i), kS, fmt.Sprintf("[subkeys=%d, values=%d]", info.SubKeysCount, info.ValuesCount))

	for i := 0; i < int(info.ValuesCount); i++ {

		value, err := cli.BaseRegEnumValue(ctx, &winreg.BaseRegEnumValueRequest{
			Key:   h,
			Index: uint32(i),
			ValueNameIn: &winreg.UnicodeString{
				MaximumLength: uint16(info.MaxValueNameLength) + 2,
			},
			DataLength: info.MaxValueLength,
		})

		if err != nil {
			return err
		}

		v, err := winreg.DecodeValue(value.Type, value.Data)
		if err != nil {
			return fmt.Errorf("enumerate: enum_values: %v", err)
		}

		switch t := fmt.Sprintf("{%d}", value.Type); value.Type {
		case winreg.RegDword, winreg.RegDwordBigEndian, winreg.RegQword, winreg.RegString, winreg.RegExpandString:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, v)
		case winreg.RegNone:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, "{NONE}")
		case winreg.RegMultistring:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t)
			for i, b := range v.([]string) {
				fmt.Println("\t\t", fmt.Sprintf("[%d]", i), b)
			}
		default:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, hex.EncodeToString(value.Data))
		}
	}

	if !recurisve {
		return nil
	}

	for i := 0; i < int(info.SubKeysCount); i++ {

		sub, err := cli.BaseRegEnumKey(ctx, &winreg.BaseRegEnumKeyRequest{
			Key:   h,
			Index: uint32(i),
			NameIn: &winreg.UnicodeString{
				MaximumLength: uint16(info.MaxSubKeyLength) + 2,
			},
			ClassIn: &winreg.UnicodeString{
				MaximumLength: uint16(info.MaxClassLength) + 2,
			},
			LastWriteTime: info.LastWriteTime,
		})

		if err != nil {
			return fmt.Errorf("enumerate: enum_keys: %v", err)
		}

		open, err := cli.BaseRegOpenKey(ctx, &winreg.BaseRegOpenKeyRequest{
			Key:           h,
			SubKey:        sub.NameOut,
			DesiredAccess: winreg.KeyQueryValue | winreg.KeyEnumerateSubKeys,
		})

		if err != nil {
			return fmt.Errorf("open: sub_key: %s: %v", sub.NameOut.Buffer, err)
		}

		Enumerate(ctx, cli, i, kS+"\\"+sub.NameOut.Buffer, open.ResultKey)
	}

	return nil
}
