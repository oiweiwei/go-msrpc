package dcom

// IID_IActivation: RPC interface UUID for IActivation. {4d9f4ab8-7d1c-11cf-861e-0020af6e7c57}
var ActivationIID = &IID{Data1: 0x4d9f4ab8, Data2: 0x7d1c, Data3: 0x11cf, Data4: []byte{0x86, 0x1e, 0x00, 0x20, 0xaf, 0x6e, 0x7c, 0x57}}

// IID_IActivationPropertiesIn: The value of the iid field of the pActProperties OBJREF structure {000001A2-0000-0000-C000-000000000046}
var ActivationPropertiesInIID = &IID{Data1: 0x000001A2, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// IID_IActivationPropertiesOut: The value of the iid field of the ppActProperties OBJREF structure {000001A3-0000-0000-C000-000000000046}
var ActivationPropertiesOutIID = &IID{Data1: 0x000001A3, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// IID_IContext: The value of the iid field of the Context structure. {000001c0-0000-0000-C000-000000000046}
var ContextIID = &IID{Data1: 0x000001c0, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}

// IID_IObjectExporter: RPC interface UUID for IObjectExporter. {99fcfec4-5260-101b-bbcb-00aa0021347a}
var ObjectExporterIID = &IID{Data1: 0x99fcfec4, Data2: 0x5260, Data3: 0x101b, Data4: []byte{0xbb, 0xcb, 0x00, 0xaa, 0x00, 0x21, 0x34, 0x7a}}

// IID_IRemoteSCMActivator: RPC interface UUID for IRemoteSCMActivator. {000001A0-0000-0000-C000-000000000046}
var RemoteSCMActivatorIID = &IID{Data1: 0x000001A0, Data2: 0x0000, Data3: 0x0000, Data4: []byte{0xc0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
