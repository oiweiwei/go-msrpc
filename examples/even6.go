//go:build exclude

// even6.go runs various eventlog6 commands. dump will dump the event log,
// chan will list the log channels, publ will list the log publishers and meta will print
// all the templates for the given publisher.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-xmlfmt/xmlfmt"

	"github.com/oiweiwei/go-msrpc/ssp"
	"github.com/oiweiwei/go-msrpc/ssp/credential"
	"github.com/oiweiwei/go-msrpc/ssp/gssapi"

	"github.com/oiweiwei/go-msrpc/dcerpc"

	"github.com/oiweiwei/go-msrpc/msrpc/epm/epm/v3"
	"github.com/oiweiwei/go-msrpc/msrpc/even6/ieventservice/v1"

	"github.com/oiweiwei/go-msrpc/msrpc/binxml"

	_ "github.com/oiweiwei/go-msrpc/msrpc/erref/hresult"
	"github.com/oiweiwei/go-msrpc/msrpc/erref/win32"
)

func init() {
	// add credentials.
	gssapi.AddCredential(credential.NewFromPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))
	// add mechanism.
	gssapi.AddMechanism(ssp.SPNEGO)
	gssapi.AddMechanism(ssp.NTLM)
}

var (
	cmd string
	pbl string
	pth string
	qry string
)

/*
	structured-query:

	go run examples/even6.go -c dump  -query '<?xml version="1.0" encoding="UTF-8"?><QueryList><Query Id="1" Path="Application"><Select Path="Application">*[System/Provider/@Name="Microsoft-Windows-Security-SPP"]</Select></Query></QueryList>'

*/

func init() {
	flag.StringVar(&cmd, "c", "dump", "even6 command: dump | meta | chan | publ")
	flag.StringVar(&pbl, "publisher", "Microsoft-Windows-MSDTC 2", "publisher name")
	flag.StringVar(&pth, "path", "Application", "log path")
	flag.StringVar(&qry, "query", "*", "query")
	flag.Parse()
}

func main() {

	ctx := gssapi.NewSecurityContext(context.Background())

	cc, err := dcerpc.Dial(ctx, "ncacn_ip_tcp:"+os.Getenv("SERVER"),
		epm.EndpointMapper(ctx,
			net.JoinHostPort(os.Getenv("SERVER"), "135"),
			dcerpc.WithInsecure(),
		))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer cc.Close(ctx)

	cli, err := ieventservice.NewEventServiceClient(ctx, cc, dcerpc.WithSeal())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	switch cmd {
	case "dump":
		DumpLogs(ctx, cli)
	case "publ":
		Publishers(ctx, cli)
	case "meta":
		Metadata(ctx, cli)
	case "chan":
		Channels(ctx, cli)
	default:
		fmt.Fprintln(os.Stderr, "unkown command", cmd)
	}

}

var PublisherMetadata = []string{
	0:  "PublisherGUID:",
	1:  "ResourceFilePath:",
	2:  "ParameterFilePath:",
	3:  "MessageFilePath:",
	7:  "ChannelReferencePath:",
	8:  "ChannelReferenceIndex:",
	9:  "ChannelReferenceID:",
	10: "ChannelReferenceFlags:",
	11: "ChannelReferenceMessageID:",
	28: "",
}

var PublisherResourceMetadata = []string{
	4:  "HelpLink:",
	5:  "PublisherMessageID:",
	13: "LevelName:",
	14: "LevelValue:",
	15: "LevelMessageID:",
	17: "TaskName:",
	18: "TaskEventGUID:",
	19: "TaskValue:",
	20: "TaskMessageID:",
	22: "OpcodeName:",
	23: "OpcodeValue:",
	24: "OpcodeMessageID:",
	26: "KeywordName:",
	27: "KeywordValue:",
	28: "KeywordMessageID:",
}

var EventMetadata = []string{
	0: "EventID:",
	1: "Version:",
	2: "ChannelID:",
	3: "Level:",
	4: "Opcode:",
	5: "Task:",
	6: "Keyword:",
	7: "MessageID:",
	8: "EventDefinitionTemplate:",
}

func DataToString(data any) string {
	switch data := data.(type) {
	case *ieventservice.Variant_NullValue:
		return "NULL"
	case *ieventservice.Variant_GUIDValue:
		return data.GUIDValue.String()
	case *ieventservice.Variant_GUIDArray:
		r := ""
		for _, u := range data.GUIDArray.Array {
			r += "\n  - " + u.String()
		}
		return r
	case *ieventservice.Variant_StringValue:
		return data.StringValue
	case *ieventservice.Variant_StringArray:
		return "\n   - " + strings.Join(data.StringArray.Array, "\n   - ")
	case *ieventservice.Variant_Uint32Value:
		return strconv.FormatUint(uint64(data.Uint32Value), 10)
	case *ieventservice.Variant_Uint32Array:
		r := ""
		for _, u := range data.Uint32Array.Array {
			r += "\n   - " + strconv.FormatUint(uint64(u), 10)
		}
		return r
	case *ieventservice.Variant_Uint64Value:
		return "0x" + strconv.FormatUint(uint64(data.Uint64Value), 16)
	case *ieventservice.Variant_Uint64Array:
		r := ""
		for _, u := range data.Uint64Array.Array {
			r += "\n   - " + "0x" + strconv.FormatUint(uint64(u), 16)
		}
		return r
	}

	return "?"
}

func Metadata(ctx context.Context, cli ieventservice.EventServiceClient) {

	meta, err := cli.GetPublisherMetadata(ctx, &ieventservice.GetPublisherMetadataRequest{
		PublisherID: pbl,
		Locale:      1033,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "publisher_metadata", err)
		return
	}

	for i, variant := range meta.PublisherMetadataProperties.Properties {
		if PublisherMetadata[i] == "" {
			continue
		}
		fmt.Println(PublisherMetadata[i], DataToString((any)(variant.Variant.Value)))
	}

	for _, i := range []uint32{4, 5, 12, 16, 21, 25} {

		rmeta, err := cli.GetPublisherResourceMetadata(ctx, &ieventservice.GetPublisherResourceMetadataRequest{
			Handle:     meta.PublisherMetadata,
			PropertyID: uint32(i),
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, "publisher_resource_metadata", i, err)
			continue
		}

		for i, variant := range rmeta.PublisherMetadataProperties.Properties {
			if PublisherResourceMetadata[i] == "" || variant.Type == ieventservice.VariantTypeNull {
				continue
			}
			fmt.Println(PublisherResourceMetadata[i], DataToString((any)(variant.Variant.Value)))
		}
	}

	enum, err := cli.GetEventMetadataEnum(ctx, &ieventservice.GetEventMetadataEnumRequest{
		PublisherMetadata: meta.PublisherMetadata,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "event_metadata_enum", err)
		return
	}

	event, err := cli.GetNextEventMetadata(ctx, &ieventservice.GetNextEventMetadataRequest{
		EventMetadataEnum: enum.EventMetadataEnum,
		RequestedLength:   100,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "next_event_metadata", err)
		return
	}

	for i, vl := range event.EventMetadataInstances {
		fmt.Println(" ------------------------------- ")
		fmt.Printf("[%d] \n", i)
		fmt.Println(" ------------------------------- ")
		for j, variant := range vl.Properties {
			fmt.Println(EventMetadata[j], DataToString((any)(variant.Variant.Value)))
		}

		e, err := cli.MessageRender(ctx, &ieventservice.MessageRenderRequest{
			PublisherConfigObject: meta.PublisherMetadata,
			SizeEventID:           1,
			Flags:                 0x00000008,
			MessageID:             vl.Properties[7].Variant.GetValue().(uint32),
			Values: &ieventservice.VariantList{
				Properties: []*ieventservice.Variant{},
			},
			MaxSizeString: 16384,
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, "message_render", err)
			continue
		}

		if e.Error != nil && e.Error.Error != 0 {
			fmt.Println(" ------------------------------- ")
			fmt.Println("ERROR:", win32.FromCode(e.Error.Error), e.Error.SubError, e.Error.SubErrorParam)
			fmt.Println(" ------------------------------- ")
		}

		fmt.Println(" ------------------------------- ")
		fmt.Println(string(e.String))
		fmt.Println(" ------------------------------- ")
	}

}

func Channels(ctx context.Context, cli ieventservice.EventServiceClient) {

	chans, err := cli.GetChannelList(ctx, &ieventservice.GetChannelListRequest{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "channels_list", err)
		return
	}

	for i, chn := range chans.ChannelPaths {
		fmt.Println(fmt.Sprintf("[%d] %q", i, chn))
	}
}

func Publishers(ctx context.Context, cli ieventservice.EventServiceClient) {

	publs, err := cli.GetPublisherList(ctx, &ieventservice.GetPublisherListRequest{})
	if err != nil {
		fmt.Fprintln(os.Stderr, "publisher_list", err)
		return
	}

	for i, publ := range publs.PublisherIDs {
		fmt.Println(fmt.Sprintf("[%d] %q", i, publ))
	}
}

func DumpLogs(ctx context.Context, cli ieventservice.EventServiceClient) {

	query, err := cli.RegisterLogQuery(ctx, &ieventservice.RegisterLogQueryRequest{
		Path:  pth,
		Query: qry,
		Flags: 0x00000101,
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "register_log_query:", err)
		return
	}

	defer func() {
		if _, err := cli.Close(ctx, &ieventservice.CloseRequest{
			Handle: query.Handle.ContextHandle(),
		}); err != nil {
			fmt.Fprintln(os.Stderr, "close_log_handle:", err)
			return
		}
	}()

	defer func() {
		if _, err := cli.Close(ctx, &ieventservice.CloseRequest{
			Handle: query.OperationControl.ContextHandle(),
		}); err != nil {
			fmt.Fprintln(os.Stderr, "close_control_handle:", err)
			return
		}
	}()

	for {

		events, err := cli.QueryNext(ctx, &ieventservice.QueryNextRequest{
			LogQuery:               query.Handle,
			RequestedRecordsLength: 100,
			TimeoutEnd:             3000,
			Flags:                  0,
		})

		if err != nil {
			fmt.Fprintln(os.Stderr, "query_next:", err)
			return
		}

		messages := make([][]byte, len(events.EventDataIndices))

		for i, offset := range events.EventDataIndices {
			messages[i] = events.ResultBuffer[offset : offset+events.EventDataSizes[i]]
		}

		for _, m := range messages {
			rs, err := binxml.Unmarshal(m)
			if err != nil {
				fmt.Fprintln(os.Stderr, "parse_binxml:", err)
				return
			}
			fmt.Println(xmlfmt.FormatXML(binxml.NewRenderer().Render(rs.Document), "", "  "))
		}

		if len(events.EventDataIndices) != 100 {
			break
		}
	}
}
