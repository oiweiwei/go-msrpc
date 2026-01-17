// The raa package implements the RAA client protocol.
//
// # Introduction
//
// This document specifies the Remote Authorization API Protocol. The Remote Authorization
// API Protocol is a Remote Procedure Call (RPC)-based protocol used to perform various
// authorization queries on remote computers.
//
// # Overview
//
// The Remote Authorization API (RAZA) protocol is designed to allow applications to
// simulate an access control decision that would be made when a given principal attempts
// to access a resource on a remote service that is protected with a given authorization
// policy. Because these are simulations, they can vary from the actual groups and/or
// claims in a user's token.
//
// For example, a user can log on with a password, or the user can log on using a smart
// card (with authentication assurance provisioned). Each type of logon will result
// in a different kind of impersonation token. Logging on using the password produces
// an impersonation token with a mapped group or claim; logging on using the smart card
// produces an impersonation token without a mapped group or claim.
//
// The following are some of the examples of this protocol's applications:
//
// * Simulate the groups and/or claims that a user would have if the user were to authenticate
// to a remote service.
//
// * Simulate a user's access to a hypothetical resource on a specific remote service
// that is protected with a given authorization policy.
//
// * Simulate how potential changes to the user's group or claim assignments can affect
// access to resources on the remote machine.
//
// The RAZA protocol defines client and server protocol roles.<1> A general description
// of message flow is as follows:
//
// *
//
// The RAZA client initiates a RAZA conversation by issuing a request to a RAZA server
// to initialize and maintain a resource manager object.
//
// *
//
// The RAZA server listens to an RPC endpoint ( c83d08f7-2128-4124-9674-3f5c23739ff9#gt_8be6a1fb-bc3c-4ee3-8018-c236f351222a
// ). When a client makes the preceding request, the RAZA server creates and maintains
// state for a resource manager object on behalf of the client.
//
// *
//
// The RAZA client can then request creation of a client context ( c83d08f7-2128-4124-9674-3f5c23739ff9#gt_519bff3c-1c9f-4d5a-aa88-a3c820a4ff3a
// ) for a user by specifying the user's security identifier (SID) ( c83d08f7-2128-4124-9674-3f5c23739ff9#gt_83f2020d-0804-4840-a5ac-e06439d50f8d
// ). After a client context is successfully created on the server, the RAZA client
// can examine the contents of the client context (for example, the group SIDs and claims
// within the client context) and/or modify the client context. Additionally, the RAZA
// client can perform an "AccessCheck" using the client context and a specified security
// descriptor ( c83d08f7-2128-4124-9674-3f5c23739ff9#gt_e5213722-75a9-44e7-b026-8e4833f0d350
// ).
//
// RAZA supports the following method calls to provide clients a way to simulate access
// control decisions.
//
// * AuthzrFreeContext
//
// * AuthzrInitializeContextFromSid
//
// * AuthzrInitializeCompoundContext
//
// * AuthzrAccessCheck
//
// * AuthzGetInformationFromContext
//
// * AuthzrModifyClaims
//
// * AuthzrModifySids
package raa

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf16"

	dcerpc "github.com/oiweiwei/go-msrpc/dcerpc"
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
)

var (
	// import guard
	GoPackage = "raa"
)
