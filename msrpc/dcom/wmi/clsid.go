package wmi

import "github.com/oiweiwei/go-msrpc/msrpc/dcom"

var (
	// 4590F812-1D3A-11D0-891F-00AA004B2E24
	ClassObjectUnmarshalClassID = &dcom.ClassID{Data1: 0x4590F812, Data2: 0x1D3A, Data3: 0x11D0, Data4: []byte{0x89, 0x1F, 0x00, 0xAA, 0x00, 0x4B, 0x2E, 0x24}}

	// 674B6698-EE92-11D0-AD71-00C04FD8FDFF
	ContextUnmarshalClassID = &dcom.ClassID{Data1: 0x674B6698, Data2: 0xEE92, Data3: 0x11D0, Data4: []byte{0xAD, 0x71, 0x00, 0xC0, 0x4F, 0xD8, 0xFD, 0xFF}}
)
