//go:build exclude

// rrp.go script dumps the windows registry keys on the remote machine (very slowly).
// NOTE that SMB cannot use security other than WithInsecure() and WithSeal().
package main

import (
	"context"
	"encoding/binary"
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
	cfg = config.New().DisableEPM()
)

func init() {
	config_flag.BindFlags(cfg, flag.CommandLine)
}

func main() {

	if err := config_flag.ParseAndValidate(cfg, flag.CommandLine); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
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

	key, err := cli.OpenCurrentUser(ctx, &winreg.OpenCurrentUserRequest{
		DesiredAccess: winreg.KeyQueryValue | winreg.KeyEnumerateSubKeys,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "key", err)
		os.Exit(1)
	}

	k := "\\\\HCLASS_CURRENT_USER"

	if err := RecursiveEnumerateKeys(ctx, cli, 0, k, key.Key); err != nil {
		fmt.Fprintln(os.Stderr, "key", err)
		os.Exit(1)
	}

	kkey, err := cli.BaseRegOpenKey(ctx, &winreg.BaseRegOpenKeyRequest{
		Key: key.Key,
		SubKey: &winreg.UnicodeString{
			Buffer: "Software\\Classes\\Extensions\\ContractId\u0000",
		},
		DesiredAccess: winreg.KeyQueryValue | winreg.KeyEnumerateSubKeys,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "key", err)
		os.Exit(1)
	}

	if err := RecursiveEnumerateKeys(ctx, cli, 0, k+"\\Software\\Classes\\Extensions\\ContractId", kkey.ResultKey); err != nil {
		fmt.Fprintln(os.Stderr, "key", err)
		os.Exit(1)
	}
}

func RecursiveEnumerateKeys(ctx context.Context, cli winreg.WinregClient, i int, n string, h *winreg.Key) error {

	info, err := cli.BaseRegQueryInfoKey(ctx, &winreg.BaseRegQueryInfoKeyRequest{
		Key: h,
		ClassIn: &winreg.UnicodeString{
			MaximumLength: 1024,
		},
	})
	if err != nil {
		return fmt.Errorf("enumerate: query_info: %v", err)
	}

	fmt.Println(fmt.Sprintf("[%d]", i), n, fmt.Sprintf("[subkeys=%d, values=%d]", info.SubKeysCount, info.ValuesCount))

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

		switch t := fmt.Sprintf("{%d}", value.Type); value.Type {
		case winreg.RegDword:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, binary.LittleEndian.Uint32(value.Data))
		case winreg.RegDwordBigEndian:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, binary.BigEndian.Uint32(value.Data))
		case winreg.RegMultistring:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t)
			for i, b := range strings.Split(string(value.Data), string([]byte{0})) {
				fmt.Println("\t\t", fmt.Sprintf("[%d]", i), b)
			}
		case winreg.RegQword:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, binary.LittleEndian.Uint64(value.Data))
		case winreg.RegNone:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, "{NONE}")
		case winreg.RegString, winreg.RegExpandString:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, string(value.Data))
		default:
			fmt.Println("\t", value.ValueNameOut.Buffer, "=", t, hex.EncodeToString(value.Data))
		}
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
			return fmt.Errorf("open: sub_key: %v", err)
		}

		RecursiveEnumerateKeys(ctx, cli, i, n+"\\"+sub.NameOut.Buffer, open.ResultKey)
	}

	return nil
}
