package dcom

import (
	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

func (o *IPID) MarshalJSON() ([]byte, error) { return o.GUID().MarshalJSON() }

func (o *IID) MarshalJSON() ([]byte, error) { return o.GUID().MarshalJSON() }

func (o *IID) String() string { return o.GUID().String() }

func (o *ClassID) MarshalJSON() ([]byte, error) { return o.GUID().MarshalJSON() }

func (o *ClassID) String() string { return o.GUID().String() }

func (o *IPID) UUID() *uuid.UUID { return o.GUID().UUID() }
