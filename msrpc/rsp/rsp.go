// The rsp package implements the RSP client protocol.
//
// # Introduction
//
// This document specifies the Remote Shutdown Protocol. The Remote Shutdown Protocol
// is a remote procedure call (RPC)-based protocol used to shut down or terminate shutdown
// on a remote computer.
//
// # Overview
//
// The Remote Shutdown Protocol is designed for shutting down a remote computer or for
// terminating the shutdown of a remote computer during the shutdown waiting period.
// Following are some of the examples of this protocol's applications:
//
// * Shut down a remote computer and display a message in the shutdown dialog box for
// 30 seconds.
//
// * Terminate a requested remote system shutdown during the shutdown waiting period.
//
// * Force applications to be closed, log off users, and shut down a remote computer.
//
// * Reboot a remote computer.
//
// In this document, the use of the terms client and server are in the protocol client
// and server context. This means that the client will initiate an RPC call and the
// server will respond.
//
// This is an RPC-based protocol. The protocol operation is stateless.
//
// This is a simple request-response protocol. For every method that the server receives,
// it executes the method and returns a completion. The client simply returns the completion
// status to the caller. This is a stateless protocol; each method call is independent
// of any previous method calls.
package rsp

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
	errors "github.com/oiweiwei/go-msrpc/dcerpc/errors"
	uuid "github.com/oiweiwei/go-msrpc/midl/uuid"
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
)

var (
	// import guard
	GoPackage = "rsp"
)

// UnicodeString structure represents REG_UNICODE_STRING RPC structure.
//
// This REG_UNICODE_STRING structure represents a counted string of Unicode (UTF-16)
// characters.
type UnicodeString struct {
	// Length:  The number of bytes actually used by the string. Because all UTF-16 characters
	// occupy 2 bytes, this MUST be an even number in the range [0...65534]. The behavior
	// for odd values is unspecified.
	Length uint16 `idl:"name:Length" json:"length"`
	// MaximumLength:  The number of bytes allocated for the string. This MUST be an even
	// number in the range [Length...65534].
	MaximumLength uint16 `idl:"name:MaximumLength" json:"maximum_length"`
	// Buffer:  The Unicode UTF-16 characters comprising the string described by the structure.
	// Note that counted strings might be terminated by a 0x0000 character, by convention;
	// if such a terminator is present, it SHOULD NOT count toward the Length (but MUST,
	// of course, be included in the MaximumLength).
	Buffer []uint16 `idl:"name:Buffer;size_is:((MaximumLength/2));length_is:((Length/2))" json:"buffer"`
}

func (o *UnicodeString) xxx_PreparePayload(ctx context.Context) error {
	if o.Buffer != nil && o.MaximumLength == 0 {
		o.MaximumLength = uint16((len(o.Buffer) * 2))
	}
	if o.Buffer != nil && o.Length == 0 {
		o.Length = uint16((len(o.Buffer) * 2))
	}
	if hook, ok := (interface{})(o).(interface{ AfterPreparePayload(context.Context) error }); ok {
		if err := hook.AfterPreparePayload(ctx); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) MarshalNDR(ctx context.Context, w ndr.Writer) error {
	if err := o.xxx_PreparePayload(ctx); err != nil {
		return err
	}
	if err := w.WriteAlign(7); err != nil {
		return err
	}
	if err := w.WriteData(o.Length); err != nil {
		return err
	}
	if err := w.WriteData(o.MaximumLength); err != nil {
		return err
	}
	if o.Buffer != nil || (o.MaximumLength/2) > 0 {
		_ptr_Buffer := ndr.MarshalNDRFunc(func(ctx context.Context, w ndr.Writer) error {
			dimSize1 := uint64((o.MaximumLength / 2))
			if err := w.WriteSize(dimSize1); err != nil {
				return err
			}
			sizeInfo := []uint64{
				dimSize1,
			}
			dimLength1 := uint64((o.Length / 2))
			if dimLength1 > sizeInfo[0] {
				dimLength1 = sizeInfo[0]
			} else {
				sizeInfo[0] = dimLength1
			}
			if err := w.WriteSize(0); err != nil {
				return err
			}
			if err := w.WriteSize(dimLength1); err != nil {
				return err
			}
			for i1 := range o.Buffer {
				i1 := i1
				if uint64(i1) >= sizeInfo[0] {
					break
				}
				if err := w.WriteData(o.Buffer[i1]); err != nil {
					return err
				}
			}
			for i1 := len(o.Buffer); uint64(i1) < sizeInfo[0]; i1++ {
				if err := w.WriteData(uint16(0)); err != nil {
					return err
				}
			}
			return nil
		})
		if err := w.WritePointer(&o.Buffer, _ptr_Buffer); err != nil {
			return err
		}
	} else {
		if err := w.WritePointer(nil); err != nil {
			return err
		}
	}
	return nil
}
func (o *UnicodeString) UnmarshalNDR(ctx context.Context, w ndr.Reader) error {
	if err := w.ReadAlign(7); err != nil {
		return err
	}
	if err := w.ReadData(&o.Length); err != nil {
		return err
	}
	if err := w.ReadData(&o.MaximumLength); err != nil {
		return err
	}
	_ptr_Buffer := ndr.UnmarshalNDRFunc(func(ctx context.Context, w ndr.Reader) error {
		sizeInfo := []uint64{
			0,
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		for sz1 := range sizeInfo {
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
			if err := w.ReadSize(&sizeInfo[sz1]); err != nil {
				return err
			}
		}
		if sizeInfo[0] > uint64(w.Len()) /* sanity-check */ {
			return fmt.Errorf("buffer overflow for size %d of array o.Buffer", sizeInfo[0])
		}
		o.Buffer = make([]uint16, sizeInfo[0])
		for i1 := range o.Buffer {
			i1 := i1
			if err := w.ReadData(&o.Buffer[i1]); err != nil {
				return err
			}
		}
		return nil
	})
	_s_Buffer := func(ptr interface{}) { o.Buffer = *ptr.(*[]uint16) }
	if err := w.ReadPointer(&o.Buffer, _s_Buffer, _ptr_Buffer); err != nil {
		return err
	}
	return nil
}
