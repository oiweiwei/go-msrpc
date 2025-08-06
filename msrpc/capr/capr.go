// The capr package implements the CAPR client protocol.
//
// # Introduction
//
// The Central Access Policy Identifier (ID) Retrieval Protocol enables an administrative
// tool to query the Central Access Policies (CAPs) configured on a remote computer.
//
// # Overview
//
// The Central Access Policy ID Retrieval (CAPR) Protocol is designed to allow an administrative
// tool running on one computer to remotely query the set of central access control
// policies configured on another computer.
//
// Central access policy objects are created in Active Directory using administrative
// authorization tools. Selected central access policy objects are deployed to other
// computers using Group Policy: Central Access Policies Extension (CAPE, described
// in [MS-GPCAP]). Other administrative tools can then use CAPR to determine which central
// policy objects have been deployed to a given remote computer.
//
// Within CAPE and CAPR, central access policies are represented by Central Access Policy
// IDs (CAPIDs). A CAPID is simply the SID of a central access policy object within
// Active Directory.
//
// The typical use scenario is as follows. An administrative interface tool uses CAPR
// to obtain the CAPIDs of one or more central access policy objects. The tool then
// uses these CAPIDs with CAPE and Lightweight Directory Access Protocol (LDAP): The
// Protocol, specified in [RFC4511], to obtain detailed information about the policies.
// That data can then be presented to the user and manipulated in whatever manner is
// appropriate to the administrative interface tool, such as to perform authorization
// tasks.
//
// This protocol defines one RPC call, LsarGetAvailableCAPIDs, for client applications
// to use. See section 3.1.4.1 for details of this call's use.
package capr

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
	lsarpc "github.com/oiweiwei/go-msrpc/msrpc/lsat/lsarpc/v0"
	ndr "github.com/oiweiwei/go-msrpc/ndr"
)

var (
	_ = context.Background
	_ = fmt.Errorf
	_ = utf16.Encode
	_ = strings.TrimPrefix
	_ = ndr.ZeroString
	_ = (*uuid.UUID)(nil)
	_ = (*dcerpc.SyntaxID)(nil)
	_ = (*errors.Error)(nil)
	_ = lsarpc.GoPackage
)

var (
	// import guard
	GoPackage = "capr"
)

// WrappedCapidSet structure represents LSAPR_WRAPPED_CAPID_SET RPC structure.
//
// The LSAPR_WRAPPED_CAPID_SET structure is a container for an array of LSAPR_SID_INFORMATION
// structures.
type WrappedCapidSet struct {
	// Entries: The number of elements in the SidInfo array.
	Entries uint32 `idl:"name:Entries" json:"entries"`
	// SidInfo: A pointer to an array of LSAPR_SID_INFORMATION structures, as defined in
	// [MS-LSAT] section 2.2.17.
	SIDInfo []*lsarpc.SIDInformation `idl:"name:SidInfo;size_is:(Entries)" json:"sid_info"`
}

func (o *WrappedCapidSet) xxx_PreparePayload(ctx context.Context) error {
	if err := ndr.BeforePreparePayload(ctx, o); err != nil {
		return err
	}
	if o.SIDInfo != nil && o.Entries == 0 {
		o.Entries = uint32(len(o.SIDInfo))
	}
	if err := ndr.AfterPreparePayload(ctx, o); err != nil {
		return err
	}
	return nil
}
func (o *WrappedCapidSet) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(9); err != nil {
		return err
	}
	if err := w.WriteData(o.Entries); err != nil {
		return err
	}
	if o.SIDInfo != nil || o.Entries > 0 {
		_ptr_SidInfo := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64(o.Entries)
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			for i1 := range o.SIDInfo {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if o.SIDInfo[i1] != nil {
					if err := o.SIDInfo[i1].MarshalNDR(ctx, w); err != nil {
						return err
					}
				} else {
					if err := (&lsarpc.SIDInformation{}).MarshalNDR(ctx, w); err != nil {
						return err
					}
				}
			}
			for i1 := len(o.SIDInfo); uint64(i1) < sizeInfo[0]; i1++ {
				if err := (&lsarpc.SIDInformation{}).MarshalNDR(ctx, w); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.SIDInfo, _ptr_SidInfo); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *WrappedCapidSet) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(9); err != nil {
		return err
	}
	if err := w.ReadData(&o.Entries); err != nil {
		return err
	}
	_ptr_SidInfo := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		// XXX: for opaque unmarshaling
		if o.Entries > 0 && sizeInfo[0] == 0 {
			sizeInfo[0] = uint64(o.Entries)
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.SIDInfo", sizeInfo[0])
		}
		o.SIDInfo = make([]*lsarpc.SIDInformation, sizeInfo[0])
		for i1 := range o.SIDInfo {
			i1 := i1
			if o.SIDInfo[i1] == nil {
				o.SIDInfo[i1] = &lsarpc.SIDInformation{}
			}
			if err := o.SIDInfo[i1].UnmarshalNDR(ctx, w); err != nil {
				return err
			}
		}
		return nil
	})
	_s_SidInfo := func(ptr interface{}) { o.SIDInfo = *ptr.(*[]*lsarpc.SIDInformation) }
	if err := w.ReadPointer(&o.SIDInfo, _s_SidInfo, _ptr_SidInfo); err != nil {
		return err
	}
	return nil
}
