//go:build exclude

package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/oiweiwei/go-msrpc/dcerpc"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	config "github.com/oiweiwei/go-msrpc/config"
	config_flag "github.com/oiweiwei/go-msrpc/config/flag"

	"github.com/oiweiwei/go-msrpc/msrpc/dtyp"
	"github.com/oiweiwei/go-msrpc/msrpc/samr/samr/v1"

	"github.com/oiweiwei/go-msrpc/msrpc/erref/ntstatus"
)

var (
	cfg = config.New()
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
		return
	}

	cli, err := samr.NewSamrClient(ctx, cc, cfg.ClientOptions(ctx)...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	serverH, err := cli.Connect(ctx, &samr.ConnectRequest{
		DesiredAccess: dtyp.AccessMaskGenericRead | dtyp.AccessMaskGenericExecute | dtyp.AccessMaskAccessSystemSecurity,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "connect:", err)
		return
	}

	domains, err := ListDomains(ctx, cli, serverH.Server)
	if err != nil {
		fmt.Fprintln(os.Stderr, "list domains:", err)
		return
	}

	for _, domainName := range domains {

		domain, err := cli.LookupDomainInSAMServer(ctx, &samr.LookupDomainInSAMServerRequest{
			Server: serverH.Server,
			Name:   &dtyp.UnicodeString{Buffer: domainName},
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, "lookup domain:", err)
			return
		}

		fmt.Println(domainName, domain.DomainID)

		domainH, err := cli.OpenDomain(ctx, &samr.OpenDomainRequest{
			Server:        serverH.Server,
			DesiredAccess: dtyp.AccessMaskGenericRead | dtyp.AccessMaskGenericExecute | dtyp.AccessMaskAccessSystemSecurity,
			DomainID:      domain.DomainID,
		})
		if err != nil {
			fmt.Fprintln(os.Stderr, "open domain:", err)
			return
		}

		groups, err := ListGroups(ctx, cli, domainH.Domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, "list groups:", err)
			return
		}

		fmt.Println("----------------------------------------")
		fmt.Println("GROUPS:")
		fmt.Println("----------------------------------------")
		for _, group := range groups {
			fmt.Println("\t", group.Name.Buffer, "|", domain.DomainID.AddRelativeID(group.RelativeID))
		}

		aliases, err := ListAliases(ctx, cli, domainH.Domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, "list aliases:", err)
			return
		}

		fmt.Println("----------------------------------------")
		fmt.Println("ALIASES:")
		fmt.Println("----------------------------------------")
		for _, alias := range aliases {
			fmt.Println("\t", alias.Name.Buffer, "|", domain.DomainID.AddRelativeID(alias.RelativeID))
		}

		users, err := ListUsers(ctx, cli, domainH.Domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, "list users:", err)
			return
		}
		fmt.Println("----------------------------------------")
		fmt.Println("USERS:")
		fmt.Println("----------------------------------------")
		for _, user := range users {
			fmt.Println("\t", user.Name.Buffer, "|", domain.DomainID.AddRelativeID(user.RelativeID))
		}

		fmt.Println("========================================")
	}

}

func OldLargeIntegerToFiletime(li *samr.OldLargeInteger) *dtyp.Filetime {
	return &dtyp.Filetime{
		LowDateTime:  li.LowPart,
		HighDateTime: uint32(li.HighPart),
	}
}

func OldLargeIntegerToTime(li *samr.OldLargeInteger) string {

	ft := OldLargeIntegerToFiletime(li)

	if ft.IsNever() || ft.IsZero() {
		return "never"
	}

	return ft.AsTime().Format(time.RFC3339)
}

func ListDomains(ctx context.Context, cli samr.SamrClient, handle *samr.Handle) ([]string, error) {

	var domains []string

	for enum := uint32(0); ; {
		doms, err := cli.EnumerateDomainsInSAMServer(ctx, &samr.EnumerateDomainsInSAMServerRequest{
			Server:             handle,
			EnumerationContext: enum,
		})
		if err != nil {
			if !errors.Is(err, ntstatus.StatusMoreEntries) {
				return nil, err
			}
		}

		for _, dom := range doms.Buffer.Buffer {
			domains = append(domains, dom.Name.Buffer)
		}

		if enum = doms.EnumerationContext; enum == 0 || doms.CountReturned == 0 {
			break
		}
	}

	return domains, nil
}

func ListGroups(ctx context.Context, cli samr.SamrClient, handle *samr.Handle) ([]*samr.RIDEnumeration, error) {

	var groups []*samr.RIDEnumeration

	for enum := uint32(0); ; {

		grps, err := cli.EnumerateGroupsInDomain(ctx, &samr.EnumerateGroupsInDomainRequest{
			Domain:             handle,
			EnumerationContext: enum,
		})
		if err != nil {
			if !errors.Is(err, ntstatus.StatusMoreEntries) {
				return nil, err
			}
		}

		groups = append(groups, grps.Buffer.Buffer...)

		if enum = grps.EnumerationContext; grps.CountReturned == 0 || enum == 0 {
			break
		}
	}

	return groups, nil
}

func ListAliases(ctx context.Context, cli samr.SamrClient, handle *samr.Handle) ([]*samr.RIDEnumeration, error) {

	var aliases []*samr.RIDEnumeration

	for enum := uint32(0); ; {

		alses, err := cli.EnumerateAliasesInDomain(ctx, &samr.EnumerateAliasesInDomainRequest{
			Domain:             handle,
			EnumerationContext: enum,
		})
		if err != nil {
			if !errors.Is(err, ntstatus.StatusMoreEntries) {
				return nil, err
			}
		}

		aliases = append(aliases, alses.Buffer.Buffer...)

		if enum = alses.EnumerationContext; alses.CountReturned == 0 || enum == 0 {
			break
		}
	}

	return aliases, nil
}

func ListUsers(ctx context.Context, cli samr.SamrClient, handle *samr.Handle) ([]*samr.RIDEnumeration, error) {

	var users []*samr.RIDEnumeration

	for enum := uint32(0); ; {

		usrs, err := cli.EnumerateUsersInDomain(ctx, &samr.EnumerateUsersInDomainRequest{
			Domain:             handle,
			EnumerationContext: enum,
		})
		if err != nil {
			if !errors.Is(err, ntstatus.StatusMoreEntries) {
				return nil, err
			}
		}

		users = append(users, usrs.Buffer.Buffer...)

		if enum = usrs.EnumerationContext; usrs.CountReturned == 0 || enum == 0 {
			break
		}
	}

	return users, nil
}
