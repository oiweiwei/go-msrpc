package well_known

import (
	"github.com/oiweiwei/go-msrpc/midl/uuid"
)

var (
	Null                                              = uuid.UUID{TimeLow: 0x0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0x0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x0}}
	MCCCFGIClusCfgAsyncEvictCleanup                   = uuid.UUID{TimeLow: 0x52c80b95, TimeMid: 0xc1ad, TimeHiAndVersion: 0x4240, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x89, Node: [6]uint8{0x72, 0xe9, 0xfa, 0x84, 0x2, 0x5e}}
	MCIISAIAppHostAdminManager                        = uuid.UUID{TimeLow: 0x9be77978, TimeMid: 0x73ed, TimeHiAndVersion: 0x4a9a, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0xfd, Node: [6]uint8{0x13, 0xf0, 0x9f, 0xec, 0x1b, 0x13}}
	MCIISAIAppHostChangeHandler                       = uuid.UUID{TimeLow: 0x9829352, TimeMid: 0x87c2, TimeHiAndVersion: 0x418d, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x79, Node: [6]uint8{0x41, 0x33, 0x96, 0x9a, 0x48, 0x9d}}
	MCIISAIAppHostChildElementCollection              = uuid.UUID{TimeLow: 0x8a90f5f, TimeMid: 0x702, TimeHiAndVersion: 0x48d6, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x5f, Node: [6]uint8{0x2, 0xa9, 0x88, 0x5a, 0x97, 0x68}}
	MCIISAIAppHostCollectionSchema                    = uuid.UUID{TimeLow: 0xde095db1, TimeMid: 0x5368, TimeHiAndVersion: 0x4d11, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0xf6, Node: [6]uint8{0xef, 0xef, 0x61, 0x9b, 0x7b, 0xcf}}
	MCIISAIAppHostConfigException                     = uuid.UUID{TimeLow: 0x4dfa1df3, TimeMid: 0x8900, TimeHiAndVersion: 0x4bc7, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xb5, Node: [6]uint8{0xd1, 0xa4, 0x58, 0xc5, 0x24, 0x10}}
	MCIISAIAppHostConfigFile                          = uuid.UUID{TimeLow: 0xada4e6fb, TimeMid: 0xe025, TimeHiAndVersion: 0x401e, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0xd0, Node: [6]uint8{0xc3, 0x13, 0x4a, 0x28, 0x1f, 0x7}}
	MCIISAIAppHostConfigLocationCollection            = uuid.UUID{TimeLow: 0x832a32f7, TimeMid: 0xb3ea, TimeHiAndVersion: 0x4b8c, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0x60, Node: [6]uint8{0x9a, 0x29, 0x23, 0x0, 0x11, 0x84}}
	MCIISAIAppHostConfigLocation                      = uuid.UUID{TimeLow: 0x370af178, TimeMid: 0x7758, TimeHiAndVersion: 0x4dad, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x46, Node: [6]uint8{0x73, 0x91, 0xf6, 0xe1, 0x85, 0x85}}
	MCIISAIAppHostConfigManager                       = uuid.UUID{TimeLow: 0x8f6d760f, TimeMid: 0xf0cb, TimeHiAndVersion: 0x4d69, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0xf6, Node: [6]uint8{0x84, 0x8b, 0x33, 0xe9, 0xbd, 0xc6}}
	MCIISAIAppHostConstantValueCollection             = uuid.UUID{TimeLow: 0x5b5a68e6, TimeMid: 0x8b9f, TimeHiAndVersion: 0x45e1, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x99, Node: [6]uint8{0xa9, 0x5f, 0xfc, 0xcd, 0xff, 0xff}}
	MCIISAIAppHostConstantValue                       = uuid.UUID{TimeLow: 0x716caf8, TimeMid: 0x7d05, TimeHiAndVersion: 0x4a46, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x99, Node: [6]uint8{0x77, 0x59, 0x4b, 0xe9, 0x13, 0x94}}
	MCIISAIAppHostElementCollection                   = uuid.UUID{TimeLow: 0xc8550bff, TimeMid: 0x5281, TimeHiAndVersion: 0x4b1e, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0x34, Node: [6]uint8{0x99, 0xb6, 0xfa, 0x38, 0x46, 0x4d}}
	MCIISAIAppHostElementSchemaCollection             = uuid.UUID{TimeLow: 0x344cdda, TimeMid: 0x151e, TimeHiAndVersion: 0x4cbf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0xda, Node: [6]uint8{0x66, 0xae, 0x61, 0xe9, 0x77, 0x54}}
	MCIISAIAppHostElementSchema                       = uuid.UUID{TimeLow: 0xef13d885, TimeMid: 0x642c, TimeHiAndVersion: 0x4709, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0xec, Node: [6]uint8{0xb8, 0x95, 0x61, 0xc6, 0xbc, 0x69}}
	MCIISAIAppHostElement                             = uuid.UUID{TimeLow: 0x64ff8ccc, TimeMid: 0xb287, TimeHiAndVersion: 0x4dae, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x8a, Node: [6]uint8{0xa7, 0x2c, 0xbf, 0x45, 0xf4, 0x53}}
	MCIISAIAppHostMappingExtension                    = uuid.UUID{TimeLow: 0x31a83ea0, TimeMid: 0xc0e4, TimeHiAndVersion: 0x4a2c, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x1, Node: [6]uint8{0x35, 0x3c, 0xc2, 0xa4, 0xc6, 0xa}}
	MCIISAIAppHostMethodCollection                    = uuid.UUID{TimeLow: 0xd6c7cd8f, TimeMid: 0xbb8d, TimeHiAndVersion: 0x4f96, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x91, Node: [6]uint8{0xd3, 0xa5, 0xf1, 0x32, 0x2, 0x69}}
	MCIISAIAppHostMethodInstance                      = uuid.UUID{TimeLow: 0xb80f3c42, TimeMid: 0x60e0, TimeHiAndVersion: 0x4ae0, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x7, Node: [6]uint8{0xf5, 0x28, 0x52, 0xd3, 0xdb, 0xed}}
	MCIISAIAppHostMethodSchema                        = uuid.UUID{TimeLow: 0x2d9915fb, TimeMid: 0x9d42, TimeHiAndVersion: 0x4328, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x82, Node: [6]uint8{0x1b, 0x46, 0x81, 0x9f, 0xab, 0x9e}}
	MCIISAIAppHostMethod                              = uuid.UUID{TimeLow: 0x7883ca1c, TimeMid: 0x1112, TimeHiAndVersion: 0x4447, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xc3, Node: [6]uint8{0x52, 0xfb, 0xeb, 0x38, 0x6, 0x9d}}
	MCIISAIAppHostPathMapper                          = uuid.UUID{TimeLow: 0xe7927575, TimeMid: 0x5cc3, TimeHiAndVersion: 0x403b, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2e, Node: [6]uint8{0x32, 0x8a, 0x6b, 0x90, 0x4b, 0xee}}
	MCIISAIAppHostPropertyCollection                  = uuid.UUID{TimeLow: 0x191775e, TimeMid: 0xbcff, TimeHiAndVersion: 0x445a, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xf4, Node: [6]uint8{0x3b, 0xdd, 0xa5, 0x4e, 0x28, 0x16}}
	MCIISAIAppHostPropertyException                   = uuid.UUID{TimeLow: 0xeafe4895, TimeMid: 0xa929, TimeHiAndVersion: 0x41ea, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x4d, Node: [6]uint8{0x61, 0x3e, 0x23, 0xf6, 0x2b, 0x71}}
	MCIISAIAppHostPropertySchemaCollection            = uuid.UUID{TimeLow: 0x8bed2c68, TimeMid: 0xa5fb, TimeHiAndVersion: 0x4b28, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x81, Node: [6]uint8{0xa0, 0xdc, 0x52, 0x67, 0x41, 0x9f}}
	MCIISAIAppHostPropertySchema                      = uuid.UUID{TimeLow: 0x450386db, TimeMid: 0x7409, TimeHiAndVersion: 0x4667, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x5e, Node: [6]uint8{0x38, 0x4d, 0xbb, 0xee, 0x2a, 0x9e}}
	MCIISAIAppHostProperty                            = uuid.UUID{TimeLow: 0xed35f7a1, TimeMid: 0x5024, TimeHiAndVersion: 0x4e7b, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0x4d, Node: [6]uint8{0x7, 0xdd, 0xaf, 0x4b, 0x52, 0x4d}}
	MCIISAIAppHostSectionDefinitionCollection         = uuid.UUID{TimeLow: 0xb7d381ee, TimeMid: 0x8860, TimeHiAndVersion: 0x47a1, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0xf4, Node: [6]uint8{0x1f, 0x33, 0xb2, 0xb1, 0xf3, 0x25}}
	MCIISAIAppHostSectionDefinition                   = uuid.UUID{TimeLow: 0xc5c04795, TimeMid: 0x321c, TimeHiAndVersion: 0x4014, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xd6, Node: [6]uint8{0xd4, 0x46, 0x58, 0x79, 0x93, 0x93}}
	MCIISAIAppHostSectionGroup                        = uuid.UUID{TimeLow: 0xdd8a158, TimeMid: 0xebe6, TimeHiAndVersion: 0x4008, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xd9, Node: [6]uint8{0xb7, 0xec, 0xc8, 0xf1, 0x10, 0x4b}}
	MCIISAIAppHostWritableAdminManager                = uuid.UUID{TimeLow: 0xfa7660f6, TimeMid: 0x7b3f, TimeHiAndVersion: 0x4237, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0xbf, Node: [6]uint8{0xed, 0xa, 0xd0, 0xdc, 0xbb, 0xd9}}
	MCMQACIMSMQApplication2                           = uuid.UUID{TimeLow: 0x12a30900, TimeMid: 0x7300, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xe6, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQApplication3                           = uuid.UUID{TimeLow: 0xeba96b1f, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQApplication                            = uuid.UUID{TimeLow: 0xd7d6e085, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQCollection                             = uuid.UUID{TimeLow: 0x188ac2f, TimeMid: 0xecb3, TimeHiAndVersion: 0x4173, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x79, Node: [6]uint8{0x63, 0x5c, 0xa2, 0x3, 0x9c, 0x72}}
	MCMQACIMSMQCoordinatedTransactionDispenser2       = uuid.UUID{TimeLow: 0xeba96b10, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQCoordinatedTransactionDispenser3       = uuid.UUID{TimeLow: 0xeba96b14, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQCoordinatedTransactionDispenser        = uuid.UUID{TimeLow: 0xd7d6e081, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQDestination                            = uuid.UUID{TimeLow: 0xeba96b16, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQEvent2                                 = uuid.UUID{TimeLow: 0xeba96b12, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQEvent3                                 = uuid.UUID{TimeLow: 0xeba96b1c, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQEvent                                  = uuid.UUID{TimeLow: 0xd7d6e077, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQManagement                             = uuid.UUID{TimeLow: 0xbe5f0241, TimeMid: 0xe489, TimeHiAndVersion: 0x4957, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xc4, Node: [6]uint8{0xa4, 0x52, 0xfc, 0xf3, 0xe2, 0x3e}}
	MCMQACIMSMQMessage2                               = uuid.UUID{TimeLow: 0xd9933be0, TimeMid: 0xa567, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xf3, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQMessage3                               = uuid.UUID{TimeLow: 0xeba96b1a, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQMessage4                               = uuid.UUID{TimeLow: 0xeba96b23, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQMessage                                = uuid.UUID{TimeLow: 0xd7d6e074, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQOutgoingQueueManagement                = uuid.UUID{TimeLow: 0x64c478fb, TimeMid: 0xf9b0, TimeHiAndVersion: 0x4695, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x7f, Node: [6]uint8{0x43, 0x9a, 0xc9, 0x43, 0x26, 0xd3}}
	MCMQACIMSMQPrivateDestination                     = uuid.UUID{TimeLow: 0xeba96b17, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQPrivateEvent                           = uuid.UUID{TimeLow: 0xd7ab3341, TimeMid: 0xc9d3, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x47, Node: [6]uint8{0x0, 0x80, 0xc7, 0xc5, 0xa2, 0xc0}}
	MCMQACIMSMQQuery2                                 = uuid.UUID{TimeLow: 0xeba96b0e, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQuery3                                 = uuid.UUID{TimeLow: 0xeba96b19, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQuery4                                 = uuid.UUID{TimeLow: 0xeba96b24, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQuery                                  = uuid.UUID{TimeLow: 0xd7d6e072, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQQueue2                                 = uuid.UUID{TimeLow: 0xef0574e0, TimeMid: 0x6d8, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueue3                                 = uuid.UUID{TimeLow: 0xeba96b1b, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueue4                                 = uuid.UUID{TimeLow: 0xeba96b20, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfo2                             = uuid.UUID{TimeLow: 0xfd174a80, TimeMid: 0x89cf, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xf2, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfo3                             = uuid.UUID{TimeLow: 0xeba96b1d, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfo4                             = uuid.UUID{TimeLow: 0xeba96b21, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfos2                            = uuid.UUID{TimeLow: 0xeba96b0f, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfos3                            = uuid.UUID{TimeLow: 0xeba96b1e, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfos4                            = uuid.UUID{TimeLow: 0xeba96b22, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQQueueInfos                             = uuid.UUID{TimeLow: 0xd7d6e07d, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQQueueInfo                              = uuid.UUID{TimeLow: 0xd7d6e07b, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQQueueManagement                        = uuid.UUID{TimeLow: 0x7fbe7759, TimeMid: 0x5760, TimeHiAndVersion: 0x444d, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0xa5, Node: [6]uint8{0x5e, 0x7a, 0xb9, 0xa8, 0x4c, 0xce}}
	MCMQACIMSMQQueue                                  = uuid.UUID{TimeLow: 0xd7d6e076, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQTransaction2                           = uuid.UUID{TimeLow: 0x2ce0c5b0, TimeMid: 0x6e67, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xe6, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQTransaction3                           = uuid.UUID{TimeLow: 0xeba96b13, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQTransactionDispenser2                  = uuid.UUID{TimeLow: 0xeba96b11, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQTransactionDispenser3                  = uuid.UUID{TimeLow: 0xeba96b15, TimeMid: 0x2168, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x8c, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x7, 0x4f, 0x6b}}
	MCMQACIMSMQTransactionDispenser                   = uuid.UUID{TimeLow: 0xd7d6e083, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACIMSMQTransaction                            = uuid.UUID{TimeLow: 0xd7d6e07f, TimeMid: 0xdccd, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0x0, 0x60, 0x97, 0xd, 0xeb, 0xae}}
	MCMQACITransaction                                = uuid.UUID{TimeLow: 0xfb15084, TimeMid: 0xaf41, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x2b, Node: [6]uint8{0x20, 0x4c, 0x4f, 0x4f, 0x50, 0x20}}
	MSADTGIDataFactory2                               = uuid.UUID{TimeLow: 0x70669eb, TimeMid: 0xb52f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0x70, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xbf, 0xb3}}
	MSADTGIDataFactory3                               = uuid.UUID{TimeLow: 0x4639db2a, TimeMid: 0xbfc5, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x18, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xbf, 0xb3}}
	MSADTGIDataFactory                                = uuid.UUID{TimeLow: 0xeac4842, TimeMid: 0x8763, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x43, Node: [6]uint8{0x0, 0xaa, 0x0, 0xa3, 0xf0, 0xd}}
	MSADTSClaims                                      = uuid.UUID{TimeLow: 0xbba9cb76, TimeMid: 0xeb0c, TimeHiAndVersion: 0x462c, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x1b, Node: [6]uint8{0x5d, 0x8c, 0x34, 0x41, 0x57, 0x1}}
	MSBKRPBackupKey                                   = uuid.UUID{TimeLow: 0x3dde7c30, TimeMid: 0x165d, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x8f, Node: [6]uint8{0x0, 0x80, 0x5f, 0x14, 0xdb, 0x40}}
	MSBPAUBitsPeerAuth                                = uuid.UUID{TimeLow: 0xe3d0d746, TimeMid: 0xd2af, TimeHiAndVersion: 0x40fd, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x7a, Node: [6]uint8{0xd, 0x70, 0x78, 0xbb, 0x70, 0x92}}
	MSBRWSABrowser                                    = uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x1, 0x28, 0x92, 0x2, 0x1, 0x62}}
	MSCAPRLsacap                                      = uuid.UUID{TimeLow: 0xafc07e2e, TimeMid: 0x311c, TimeHiAndVersion: 0x4435, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x8c, Node: [6]uint8{0xc4, 0x83, 0xff, 0xee, 0xc7, 0xc9}}
	MSCMPOIXnRemote                                   = uuid.UUID{TimeLow: 0x906b0ce0, TimeMid: 0xc70b, TimeHiAndVersion: 0x1067, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x17, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xda}}
	MSCMRPClusAPI                                     = uuid.UUID{TimeLow: 0xb97db8b2, TimeMid: 0x4c63, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xf6, Node: [6]uint8{0x8, 0x0, 0x2b, 0xe2, 0x3f, 0x2f}}
	MSCOMAIAlternateLaunch                            = uuid.UUID{TimeLow: 0x7f43b400, TimeMid: 0x1a0e, TimeHiAndVersion: 0x4d57, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xc9, Node: [6]uint8{0x6b, 0xc, 0x65, 0xf7, 0xa8, 0x89}}
	MSCOMAICapabilitySupport                          = uuid.UUID{TimeLow: 0x47cde9a1, TimeMid: 0xbf6, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x16, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb9, 0x98, 0x8e}}
	MSCOMAICatalog64BitSupport                        = uuid.UUID{TimeLow: 0x1d118904, TimeMid: 0x94b3, TimeHiAndVersion: 0x4a64, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xa6, Node: [6]uint8{0xed, 0x43, 0x26, 0x66, 0xa7, 0xb9}}
	MSCOMAICatalogSession                             = uuid.UUID{TimeLow: 0x182c40fa, TimeMid: 0x32e4, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x8b, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x23, 0x1c, 0x29}}
	MSCOMAICatalogTableInfo                           = uuid.UUID{TimeLow: 0xa8927a41, TimeMid: 0xd3ce, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x72, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	MSCOMAICatalogTableRead                           = uuid.UUID{TimeLow: 0xe3d6630, TimeMid: 0xb46b, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	MSCOMAICatalogTableWrite                          = uuid.UUID{TimeLow: 0xe3d6631, TimeMid: 0xb46b, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0x60, 0x8, 0xb0, 0xe5, 0xca}}
	MSCOMAICatalogUtils2                              = uuid.UUID{TimeLow: 0xc726744e, TimeMid: 0x5735, TimeHiAndVersion: 0x4f08, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x86, Node: [6]uint8{0xc5, 0x10, 0xee, 0x63, 0x8f, 0xb6}}
	MSCOMAICatalogUtils                               = uuid.UUID{TimeLow: 0x456129e2, TimeMid: 0x1078, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0x5f, 0xc7, 0x32, 0x4}}
	MSCOMAIContainerControl2                          = uuid.UUID{TimeLow: 0x6c935649, TimeMid: 0x30a6, TimeHiAndVersion: 0x4211, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x87, Node: [6]uint8{0xc4, 0xc8, 0x3e, 0x5f, 0xe1, 0xc7}}
	MSCOMAIContainerControl                           = uuid.UUID{TimeLow: 0x3f3b1b86, TimeMid: 0xdbbe, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xa6, Node: [6]uint8{0x0, 0x80, 0x5f, 0x85, 0xcf, 0xe3}}
	MSCOMAIExport2                                    = uuid.UUID{TimeLow: 0xf131ea3e, TimeMid: 0xb7be, TimeHiAndVersion: 0x480e, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xd, Node: [6]uint8{0x51, 0xcb, 0x27, 0x85, 0x77, 0x9e}}
	MSCOMAIExport                                     = uuid.UUID{TimeLow: 0xcfadac84, TimeMid: 0xe12c, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x4c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x99, 0xd, 0x54}}
	MSCOMAIImport2                                    = uuid.UUID{TimeLow: 0x1f7b1697, TimeMid: 0xecb2, TimeHiAndVersion: 0x4cbb, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0xe, Node: [6]uint8{0x75, 0xc4, 0x27, 0xf4, 0xa6, 0xf0}}
	MSCOMAIImport                                     = uuid.UUID{TimeLow: 0xc2be6970, TimeMid: 0xdf9e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x87, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	MSCOMAIRegister2                                  = uuid.UUID{TimeLow: 0x971668dc, TimeMid: 0xc3fe, TimeHiAndVersion: 0x4ea1, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x43, Node: [6]uint8{0xc, 0x72, 0x30, 0xf4, 0x94, 0xa1}}
	MSCOMAIRegister                                   = uuid.UUID{TimeLow: 0x8db2180e, TimeMid: 0xbd29, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x7e, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd7, 0xa9, 0x24}}
	MSCOMAIReplicationUtil                            = uuid.UUID{TimeLow: 0x98315903, TimeMid: 0x7be5, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0xc1, Node: [6]uint8{0x0, 0xa0, 0x24, 0x63, 0xd6, 0xe7}}
	MSCOMITransactionStream                           = uuid.UUID{TimeLow: 0x97199110, TimeMid: 0xdb2e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x51, Node: [6]uint8{0x0, 0x0, 0xf8, 0x5, 0xca, 0x53}}
	MSCOMEVIEnumEventObject                           = uuid.UUID{TimeLow: 0xf4a07d63, TimeMid: 0x2e25, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	MSCOMEVIEventClass2                               = uuid.UUID{TimeLow: 0xfb2b72a1, TimeMid: 0x7a68, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	MSCOMEVIEventClass3                               = uuid.UUID{TimeLow: 0x7fb7ea43, TimeMid: 0x2d76, TimeHiAndVersion: 0x4ea8, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xd9, Node: [6]uint8{0x3d, 0xec, 0xc2, 0x70, 0x29, 0x5e}}
	MSCOMEVIEventClass                                = uuid.UUID{TimeLow: 0xfb2b72a0, TimeMid: 0x7a68, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xf9, Node: [6]uint8{0x0, 0x80, 0xc7, 0xd7, 0x71, 0xbf}}
	MSCOMEVIEventObjectCollection                     = uuid.UUID{TimeLow: 0xf89ac270, TimeMid: 0xd4eb, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x82, Node: [6]uint8{0x0, 0x80, 0x5f, 0xc7, 0x92, 0x16}}
	MSCOMEVIEventSubscription2                        = uuid.UUID{TimeLow: 0x4a6b0e16, TimeMid: 0x2e38, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x65, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	MSCOMEVIEventSubscription3                        = uuid.UUID{TimeLow: 0xfbc1d17d, TimeMid: 0xc498, TimeHiAndVersion: 0x43a0, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0xaf, Node: [6]uint8{0x42, 0x3d, 0xdd, 0x53, 0xa, 0xf6}}
	MSCOMEVIEventSubscription                         = uuid.UUID{TimeLow: 0x4a6b0e15, TimeMid: 0x2e38, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x65, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	MSCOMEVIEventSystem2                              = uuid.UUID{TimeLow: 0x99cc098f, TimeMid: 0xa48a, TimeHiAndVersion: 0x4e9c, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x58, Node: [6]uint8{0x96, 0x5c, 0xa, 0xfc, 0x19, 0xd5}}
	MSCOMEVIEventSystemInitialize                     = uuid.UUID{TimeLow: 0xa0e8f27a, TimeMid: 0x888c, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x63, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb9, 0x26, 0xaf}}
	MSCOMEVIEventSystem                               = uuid.UUID{TimeLow: 0x4e14fb9f, TimeMid: 0x2e22, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xbb, 0xb3, 0x45}}
	MSCOMTIComTrackingInfoEvents                      = uuid.UUID{TimeLow: 0x4e6cdcc9, TimeMid: 0xfb25, TimeHiAndVersion: 0x4fd5, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0xc5, Node: [6]uint8{0xc9, 0xf4, 0xb6, 0x55, 0x9c, 0xec}}
	MSCOMTIGetTrackingData                            = uuid.UUID{TimeLow: 0xb60040e0, TimeMid: 0xbcf3, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x1d, Node: [6]uint8{0x0, 0x80, 0xc7, 0x29, 0x26, 0x4d}}
	MSCOMTIProcessDump                                = uuid.UUID{TimeLow: 0x23c9dd26, TimeMid: 0x2355, TimeHiAndVersion: 0x4fe2, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xde, Node: [6]uint8{0xf7, 0x79, 0xa2, 0x38, 0xad, 0xbd}}
	MSCSRAICertAdminD2                                = uuid.UUID{TimeLow: 0x7fe0d935, TimeMid: 0xdda6, TimeHiAndVersion: 0x443f, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd0, Node: [6]uint8{0x1c, 0xfb, 0x58, 0xfe, 0x41, 0xdd}}
	MSCSRAICertAdminD                                 = uuid.UUID{TimeLow: 0xd99e6e71, TimeMid: 0xfc88, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x98, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x3, 0x12, 0xf3}}
	MSCSVPIClusterCleanup                             = uuid.UUID{TimeLow: 0xd6105110, TimeMid: 0x8917, TimeHiAndVersion: 0x41a5, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x32, Node: [6]uint8{0x8e, 0xa, 0xa2, 0x93, 0x3d, 0xc9}}
	MSCSVPIClusterFirewall                            = uuid.UUID{TimeLow: 0xf1d6c29c, TimeMid: 0x8fbe, TimeHiAndVersion: 0x4691, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x24, Node: [6]uint8{0xf6, 0xd8, 0xde, 0xae, 0xaf, 0xc8}}
	MSCSVPIClusterLog                                 = uuid.UUID{TimeLow: 0x85923ca7, TimeMid: 0x1b6b, TimeHiAndVersion: 0x4e83, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xe4, Node: [6]uint8{0xf5, 0xba, 0x3b, 0xfb, 0xb8, 0xa3}}
	MSCSVPIClusterNetwork2                            = uuid.UUID{TimeLow: 0x2931c32c, TimeMid: 0xf731, TimeHiAndVersion: 0x4c56, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xeb, Node: [6]uint8{0x3d, 0x5f, 0x1c, 0x5e, 0x72, 0xbf}}
	MSCSVPIClusterSetup                               = uuid.UUID{TimeLow: 0x491260b5, TimeMid: 0x5c9, TimeHiAndVersion: 0x40d9, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xf2, Node: [6]uint8{0x1f, 0x7b, 0xda, 0xe0, 0x92, 0x7f}}
	MSCSVPIClusterStorage2                            = uuid.UUID{TimeLow: 0x12108a88, TimeMid: 0x6858, TimeHiAndVersion: 0x4467, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x2f, Node: [6]uint8{0xe6, 0xcf, 0x45, 0x68, 0xdf, 0xb6}}
	MSCSVPIClusterStorage3                            = uuid.UUID{TimeLow: 0x11942d87, TimeMid: 0xa1de, TimeHiAndVersion: 0x4e7f, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0xfb, Node: [6]uint8{0xa8, 0x40, 0xd9, 0xc5, 0x92, 0x8d}}
	MSCSVPIClusterUpdate                              = uuid.UUID{TimeLow: 0xe3c9b851, TimeMid: 0xc442, TimeHiAndVersion: 0x432b, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xc6, Node: [6]uint8{0xa7, 0xfa, 0xaf, 0xc0, 0x9d, 0x3b}}
	MSDCOMIActivation                                 = uuid.UUID{TimeLow: 0x4d9f4ab8, TimeMid: 0x7d1c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x1e, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6e, 0x7c, 0x57}}
	MSDCOMIObjectExporter                             = uuid.UUID{TimeLow: 0x99fcfec4, TimeMid: 0x5260, TimeHiAndVersion: 0x101b, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xcb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x21, 0x34, 0x7a}}
	MSDCOMIRemoteSCMActivator                         = uuid.UUID{TimeLow: 0x1a0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSDCOMIRemUnknown2                                = uuid.UUID{TimeLow: 0x143, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSDCOMIRemUnknown                                 = uuid.UUID{TimeLow: 0x131, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSDCOMIUnknown                                    = uuid.UUID{TimeLow: 0x0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSDFSNMNetDFS                                     = uuid.UUID{TimeLow: 0x4fc742e0, TimeMid: 0x4a10, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x73, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4a, 0xe6, 0x73}}
	MSDFSRHIADProxy2                                  = uuid.UUID{TimeLow: 0xc4b0c7d9, TimeMid: 0xabe0, TimeHiAndVersion: 0x4733, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xe1, Node: [6]uint8{0x9f, 0xde, 0xdf, 0x26, 0xc, 0x7a}}
	MSDFSRHIADProxy                                   = uuid.UUID{TimeLow: 0x4bb8ab1d, TimeMid: 0x9ef9, TimeHiAndVersion: 0x4100, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0xb6, Node: [6]uint8{0xdd, 0x4b, 0x4e, 0x41, 0x8b, 0x72}}
	MSDFSRHIServerHealthReport2                       = uuid.UUID{TimeLow: 0x20d15747, TimeMid: 0x6c48, TimeHiAndVersion: 0x4254, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x58, Node: [6]uint8{0x65, 0x3, 0x9f, 0xd8, 0xc6, 0x3c}}
	MSDFSRHIServerHealthReport                        = uuid.UUID{TimeLow: 0xe65e8028, TimeMid: 0x83e8, TimeHiAndVersion: 0x491b, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0xf7, Node: [6]uint8{0xaa, 0xf6, 0xbd, 0x51, 0xa0, 0xce}}
	MSDHCPMDhcpsrv2                                   = uuid.UUID{TimeLow: 0x5b821720, TimeMid: 0xf63b, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xd2, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc3, 0x24, 0xdb}}
	MSDHCPMDhcpsrv                                    = uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x46, 0xc3, 0xf8, 0x74, 0x53, 0x2d}}
	MSDLTMTrksvr                                      = uuid.UUID{TimeLow: 0x4da1c422, TimeMid: 0x943d, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xae, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc2, 0xaa, 0x3f}}
	MSDLTWTrkwks                                      = uuid.UUID{TimeLow: 0x300f3532, TimeMid: 0x38cc, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xf0, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6b, 0xa, 0xdd}}
	MSDMRPIDMNotify                                   = uuid.UUID{TimeLow: 0xd2d79df7, TimeMid: 0x3400, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x5f, 0xf5, 0x86}}
	MSDMRPIDMRemoteServer                             = uuid.UUID{TimeLow: 0x3a410f21, TimeMid: 0x553f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x5e, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x2c, 0x9d, 0x5d}}
	MSDMRPIVolumeClient2                              = uuid.UUID{TimeLow: 0x4bdafc52, TimeMid: 0xfe6a, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0xf8, Node: [6]uint8{0x0, 0x10, 0x5a, 0x11, 0x16, 0x4a}}
	MSDMRPIVolumeClient3                              = uuid.UUID{TimeLow: 0x135698d2, TimeMid: 0x3a37, TimeHiAndVersion: 0x4d26, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0xdf, Node: [6]uint8{0xe2, 0xbb, 0x6a, 0xe3, 0xac, 0x61}}
	MSDMRPIVolumeClient4                              = uuid.UUID{TimeLow: 0xdeb01010, TimeMid: 0x3a37, TimeHiAndVersion: 0x4d26, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0xdf, Node: [6]uint8{0xe2, 0xbb, 0x6a, 0xe3, 0xac, 0x61}}
	MSDMRPIVolumeClient                               = uuid.UUID{TimeLow: 0xd2d79df5, TimeMid: 0x3400, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x5f, 0xf5, 0x86}}
	MSDNSPDnsServer                                   = uuid.UUID{TimeLow: 0x50abc2a4, TimeMid: 0x574d, TimeHiAndVersion: 0x40b3, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x66, Node: [6]uint8{0xee, 0x4f, 0xd5, 0xfb, 0xa0, 0x76}}
	MSDRSRDrsuapi                                     = uuid.UUID{TimeLow: 0xe3514235, TimeMid: 0x4b06, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x4, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc2, 0xdc, 0xd2}}
	MSDRSRDsaop                                       = uuid.UUID{TimeLow: 0x7c44d7d4, TimeMid: 0x31d5, TimeHiAndVersion: 0x424c, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x5e, Node: [6]uint8{0x2b, 0x3e, 0x1f, 0x32, 0x3d, 0x22}}
	MSDSSPDssetup                                     = uuid.UUID{TimeLow: 0x3919286a, TimeMid: 0xb10c, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0xa8, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x2e, 0xf5}}
	MSEFSREfsrpc                                      = uuid.UUID{TimeLow: 0xdf1941c5, TimeMid: 0xfe89, TimeHiAndVersion: 0x4e79, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x10, Node: [6]uint8{0x46, 0x36, 0x57, 0xac, 0xf4, 0x4d}}
	MSEFSRLsarpc                                      = uuid.UUID{TimeLow: 0xc681d488, TimeMid: 0xd850, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x52, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0xf, 0x7e}}
	MSEVEN6Eventlog                                   = uuid.UUID{TimeLow: 0xf6beaff7, TimeMid: 0x1e19, TimeHiAndVersion: 0x4fbb, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x8f, Node: [6]uint8{0xb8, 0x9e, 0x20, 0x18, 0x33, 0x7c}}
	MSEVENEventlog                                    = uuid.UUID{TimeLow: 0x82273fdc, TimeMid: 0xe32a, TimeHiAndVersion: 0x18c3, ClockSeqHiAndReserved: 0x3f, ClockSeqLow: 0x78, Node: [6]uint8{0x82, 0x79, 0x29, 0xdc, 0x23, 0xea}}
	MSFASPRemoteFW                                    = uuid.UUID{TimeLow: 0x6b5bdd1e, TimeMid: 0x528c, TimeHiAndVersion: 0x422c, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0x8c, Node: [6]uint8{0xa4, 0x7, 0x9b, 0xe4, 0xfe, 0x48}}
	MSFAXFaxclient                                    = uuid.UUID{TimeLow: 0x6099fc12, TimeMid: 0x3eff, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xd0, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x1a, 0x4e}}
	MSFAXFax                                          = uuid.UUID{TimeLow: 0xea0a3165, TimeMid: 0x4834, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xf8, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x46, 0xcc}}
	MSFRS1Frsrpc                                      = uuid.UUID{TimeLow: 0xf5cc59b4, TimeMid: 0x4264, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x59, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2f, 0x84, 0x26}}
	MSFRS1NtFrsApi                                    = uuid.UUID{TimeLow: 0xd049b186, TimeMid: 0x814f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x3c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xb2, 0x32}}
	MSFRS2FrsTransport                                = uuid.UUID{TimeLow: 0x897e2e5f, TimeMid: 0x93f3, TimeHiAndVersion: 0x4376, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x9c, Node: [6]uint8{0xfd, 0x22, 0x77, 0x49, 0x5c, 0x27}}
	MSFRS2DfsrEndpoint                                = uuid.UUID{TimeLow: 0x5bc1ed07, TimeMid: 0xf5f5, TimeHiAndVersion: 0x485f, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xfd, Node: [6]uint8{0x6f, 0xd0, 0xac, 0xf9, 0xa2, 0x3c}}
	MSFSRMIFsrmActionCommand                          = uuid.UUID{TimeLow: 0x12937789, TimeMid: 0xe247, TimeHiAndVersion: 0x4917, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x20, Node: [6]uint8{0xf3, 0xee, 0x9c, 0x7e, 0xe7, 0x83}}
	MSFSRMIFsrmActionEmail2                           = uuid.UUID{TimeLow: 0x8276702f, TimeMid: 0x2532, TimeHiAndVersion: 0x4839, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0xbf, Node: [6]uint8{0x48, 0x72, 0x60, 0x9a, 0x2e, 0xa4}}
	MSFSRMIFsrmActionEmail                            = uuid.UUID{TimeLow: 0xd646567d, TimeMid: 0x26ae, TimeHiAndVersion: 0x4caa, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x84, Node: [6]uint8{0x4e, 0xa, 0xad, 0x20, 0x7f, 0xca}}
	MSFSRMIFsrmActionEventLog                         = uuid.UUID{TimeLow: 0x4c8f96c3, TimeMid: 0x5d94, TimeHiAndVersion: 0x4f37, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xf4, Node: [6]uint8{0xf5, 0x6a, 0xb4, 0x63, 0x54, 0x6f}}
	MSFSRMIFsrmActionReport                           = uuid.UUID{TimeLow: 0x2dbe63c4, TimeMid: 0xb340, TimeHiAndVersion: 0x48a0, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0xb0, Node: [6]uint8{0x15, 0x8e, 0x7, 0xfc, 0x56, 0x7e}}
	MSFSRMIFsrmAction                                 = uuid.UUID{TimeLow: 0x6cd6408a, TimeMid: 0xae60, TimeHiAndVersion: 0x463b, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0xf1, Node: [6]uint8{0xe1, 0x17, 0x53, 0x4d, 0x69, 0xdc}}
	MSFSRMIFsrmAutoApplyQuota                         = uuid.UUID{TimeLow: 0xf82e5729, TimeMid: 0x6aba, TimeHiAndVersion: 0x4740, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xc7, Node: [6]uint8{0xc7, 0xf5, 0x8f, 0x75, 0xfb, 0x7b}}
	MSFSRMIFsrmClassificationManager                  = uuid.UUID{TimeLow: 0xd2dc89da, TimeMid: 0xee91, TimeHiAndVersion: 0x48a0, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd8, Node: [6]uint8{0xcc, 0x72, 0xa5, 0x6f, 0x7d, 0x4}}
	MSFSRMIFsrmClassificationRule                     = uuid.UUID{TimeLow: 0xafc052c2, TimeMid: 0x5315, TimeHiAndVersion: 0x45ab, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x1b, Node: [6]uint8{0xc6, 0xdb, 0xe, 0x12, 0x1, 0x48}}
	MSFSRMIFsrmClassifierModuleDefinition             = uuid.UUID{TimeLow: 0xbb36ea26, TimeMid: 0x6318, TimeHiAndVersion: 0x4b8c, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x92, Node: [6]uint8{0xf7, 0x2d, 0xd6, 0x2, 0xe7, 0xa5}}
	MSFSRMIFsrmCollection                             = uuid.UUID{TimeLow: 0xf76fbf3b, TimeMid: 0x8ddd, TimeHiAndVersion: 0x4b42, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x5a, Node: [6]uint8{0xcb, 0x1c, 0x3f, 0xf1, 0xfe, 0xe8}}
	MSFSRMIFsrmCommittableCollection                  = uuid.UUID{TimeLow: 0x96deb3b5, TimeMid: 0x8b91, TimeHiAndVersion: 0x4a2a, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x93, Node: [6]uint8{0x80, 0xa3, 0x5d, 0x8a, 0xa8, 0x47}}
	MSFSRMIFsrmDerivedObjectsResult                   = uuid.UUID{TimeLow: 0x39322a2d, TimeMid: 0x38ee, TimeHiAndVersion: 0x4d0d, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x95, Node: [6]uint8{0x42, 0x1a, 0x80, 0x84, 0x9a, 0x82}}
	MSFSRMIFsrmFileGroupImported                      = uuid.UUID{TimeLow: 0xad55f10b, TimeMid: 0x5f11, TimeHiAndVersion: 0x4be7, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0xef, Node: [6]uint8{0xd9, 0xee, 0x2e, 0x47, 0xd, 0xed}}
	MSFSRMIFsrmFileGroupManager                       = uuid.UUID{TimeLow: 0x426677d5, TimeMid: 0x18c, TimeHiAndVersion: 0x485c, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x51, Node: [6]uint8{0x20, 0xb8, 0x6d, 0x0, 0xbd, 0xc4}}
	MSFSRMIFsrmFileGroup                              = uuid.UUID{TimeLow: 0x8dd04909, TimeMid: 0xe34, TimeHiAndVersion: 0x4d55, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0xaa, Node: [6]uint8{0x89, 0xe1, 0xf1, 0xa1, 0xbb, 0xb9}}
	MSFSRMIFsrmFileManagementJobManager               = uuid.UUID{TimeLow: 0xee321ecb, TimeMid: 0xd95e, TimeHiAndVersion: 0x48e9, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x7c, Node: [6]uint8{0xc7, 0x68, 0x5a, 0x1, 0x32, 0x35}}
	MSFSRMIFsrmFileManagementJob                      = uuid.UUID{TimeLow: 0x770687e, TimeMid: 0x9f36, TimeHiAndVersion: 0x4d6f, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x78, Node: [6]uint8{0x59, 0x9d, 0x18, 0x84, 0x61, 0xc9}}
	MSFSRMIFsrmFileScreenBase                         = uuid.UUID{TimeLow: 0xf3637e80, TimeMid: 0x5b22, TimeHiAndVersion: 0x4a2b, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x37, Node: [6]uint8{0xbb, 0xb6, 0x42, 0xb4, 0x1c, 0xfc}}
	MSFSRMIFsrmFileScreenException                    = uuid.UUID{TimeLow: 0xbee7ce02, TimeMid: 0xdf77, TimeHiAndVersion: 0x4515, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x89, Node: [6]uint8{0x78, 0xf0, 0x1c, 0x5a, 0xfc, 0x1a}}
	MSFSRMIFsrmFileScreenManager                      = uuid.UUID{TimeLow: 0xff4fa04e, TimeMid: 0x5a94, TimeHiAndVersion: 0x4bda, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xa0, Node: [6]uint8{0xd5, 0xb4, 0xd3, 0xc5, 0x2e, 0xba}}
	MSFSRMIFsrmFileScreenTemplateImported             = uuid.UUID{TimeLow: 0xe1010359, TimeMid: 0x3e5d, TimeHiAndVersion: 0x4ecd, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xe4, Node: [6]uint8{0xef, 0x48, 0x62, 0x2f, 0xdf, 0x30}}
	MSFSRMIFsrmFileScreenTemplateManager              = uuid.UUID{TimeLow: 0xcfe36cba, TimeMid: 0x1949, TimeHiAndVersion: 0x4e74, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x4f, Node: [6]uint8{0xf1, 0xd5, 0x80, 0xce, 0xaf, 0x13}}
	MSFSRMIFsrmFileScreenTemplate                     = uuid.UUID{TimeLow: 0x205bebf8, TimeMid: 0xdd93, TimeHiAndVersion: 0x452a, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0xa6, Node: [6]uint8{0x32, 0xb5, 0x66, 0xb3, 0x58, 0x28}}
	MSFSRMIFsrmFileScreen                             = uuid.UUID{TimeLow: 0x5f6325d3, TimeMid: 0xce88, TimeHiAndVersion: 0x4733, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xc1, Node: [6]uint8{0x2d, 0x6a, 0xef, 0xc5, 0xea, 0x7}}
	MSFSRMIFsrmMutableCollection                      = uuid.UUID{TimeLow: 0x1bb617b8, TimeMid: 0x3886, TimeHiAndVersion: 0x49dc, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0x82, Node: [6]uint8{0xa6, 0xc9, 0xf, 0xa3, 0x5d, 0xda}}
	MSFSRMIFsrmObject                                 = uuid.UUID{TimeLow: 0x22bcef93, TimeMid: 0x4a3f, TimeHiAndVersion: 0x4183, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0xf9, Node: [6]uint8{0x2f, 0x8b, 0x8a, 0x62, 0x8a, 0xee}}
	MSFSRMIFsrmPathMapper                             = uuid.UUID{TimeLow: 0x6f4dbfff, TimeMid: 0x6920, TimeHiAndVersion: 0x4821, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xc3, Node: [6]uint8{0xb7, 0xe9, 0x4c, 0x1f, 0xd6, 0xc}}
	MSFSRMIFsrmPipelineModuleDefinition               = uuid.UUID{TimeLow: 0x515c1277, TimeMid: 0x2c81, TimeHiAndVersion: 0x440e, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xcf, Node: [6]uint8{0x36, 0x79, 0x21, 0xed, 0x4f, 0x59}}
	MSFSRMIFsrmPropertyCondition                      = uuid.UUID{TimeLow: 0x326af66f, TimeMid: 0x2ac0, TimeHiAndVersion: 0x4f68, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x8c, Node: [6]uint8{0x47, 0x59, 0xf0, 0x54, 0xfa, 0x29}}
	MSFSRMIFsrmPropertyDefinition2                    = uuid.UUID{TimeLow: 0x47782152, TimeMid: 0xd16c, TimeHiAndVersion: 0x4229, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xe1, Node: [6]uint8{0xd, 0xdf, 0xe3, 0x8, 0xb9, 0xf6}}
	MSFSRMIFsrmPropertyDefinition                     = uuid.UUID{TimeLow: 0xede0150f, TimeMid: 0xe9a3, TimeHiAndVersion: 0x419c, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x7c, Node: [6]uint8{0x1, 0xfe, 0x5d, 0x24, 0xc5, 0xd3}}
	MSFSRMIFsrmPropertyDefinitionValue                = uuid.UUID{TimeLow: 0xe946d148, TimeMid: 0xbd67, TimeHiAndVersion: 0x4178, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x22, Node: [6]uint8{0x1c, 0x44, 0x92, 0x5e, 0xd7, 0x10}}
	MSFSRMIFsrmProperty                               = uuid.UUID{TimeLow: 0x4a73fee4, TimeMid: 0x4102, TimeHiAndVersion: 0x4fcc, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xfb, Node: [6]uint8{0x38, 0x61, 0x4f, 0x9e, 0xe7, 0x68}}
	MSFSRMIFsrmQuotaBase                              = uuid.UUID{TimeLow: 0x1568a795, TimeMid: 0x3924, TimeHiAndVersion: 0x4118, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x4b, Node: [6]uint8{0x68, 0xd8, 0xf0, 0xfa, 0x5d, 0xaf}}
	MSFSRMIFsrmQuotaManagerEx                         = uuid.UUID{TimeLow: 0x4846cb01, TimeMid: 0xd430, TimeHiAndVersion: 0x494f, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xb4, Node: [6]uint8{0xb1, 0x5, 0x49, 0x99, 0xfb, 0x9}}
	MSFSRMIFsrmQuotaManager                           = uuid.UUID{TimeLow: 0x8bb68c7d, TimeMid: 0x19d8, TimeHiAndVersion: 0x4ffb, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x9e, Node: [6]uint8{0xbe, 0x4f, 0xc1, 0x73, 0x40, 0x14}}
	MSFSRMIFsrmQuotaObject                            = uuid.UUID{TimeLow: 0x42dc3511, TimeMid: 0x61d5, TimeHiAndVersion: 0x48ae, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xdc, Node: [6]uint8{0x59, 0xfc, 0x0, 0xc0, 0xa8, 0xd6}}
	MSFSRMIFsrmQuotaTemplateImported                  = uuid.UUID{TimeLow: 0x9a2bf113, TimeMid: 0xa329, TimeHiAndVersion: 0x44cc, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x9a, Node: [6]uint8{0x5c, 0x0, 0xfc, 0xe8, 0xda, 0x40}}
	MSFSRMIFsrmQuotaTemplateManager                   = uuid.UUID{TimeLow: 0x4173ac41, TimeMid: 0x172d, TimeHiAndVersion: 0x4d52, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x3c, Node: [6]uint8{0xfd, 0xc7, 0xe4, 0x15, 0xf7, 0x17}}
	MSFSRMIFsrmQuotaTemplate                          = uuid.UUID{TimeLow: 0xa2efab31, TimeMid: 0x295e, TimeHiAndVersion: 0x46bb, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x76, Node: [6]uint8{0xe8, 0x6d, 0x58, 0xb5, 0x2e, 0x8b}}
	MSFSRMIFsrmQuota                                  = uuid.UUID{TimeLow: 0x377f739d, TimeMid: 0x9647, TimeHiAndVersion: 0x4b8e, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xd2, Node: [6]uint8{0x5f, 0xfc, 0xe6, 0xd7, 0x59, 0xcd}}
	MSFSRMIFsrmReportJob                              = uuid.UUID{TimeLow: 0x38e87280, TimeMid: 0x715c, TimeHiAndVersion: 0x4c7d, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x80, Node: [6]uint8{0xea, 0x16, 0x51, 0xa1, 0x9f, 0xef}}
	MSFSRMIFsrmReportManager                          = uuid.UUID{TimeLow: 0x27b899fe, TimeMid: 0x6ffa, TimeHiAndVersion: 0x4481, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x84, Node: [6]uint8{0xd3, 0xda, 0xad, 0xe8, 0xa0, 0x2b}}
	MSFSRMIFsrmReportScheduler                        = uuid.UUID{TimeLow: 0x6879caf9, TimeMid: 0x6617, TimeHiAndVersion: 0x4484, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x19, Node: [6]uint8{0x71, 0xc3, 0xd8, 0x64, 0x5f, 0x94}}
	MSFSRMIFsrmReport                                 = uuid.UUID{TimeLow: 0xd8cc81d9, TimeMid: 0x46b8, TimeHiAndVersion: 0x4fa4, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xa5, Node: [6]uint8{0x4a, 0xa9, 0xde, 0xc9, 0xb6, 0x38}}
	MSFSRMIFsrmRule                                   = uuid.UUID{TimeLow: 0xcb0df960, TimeMid: 0x16f5, TimeHiAndVersion: 0x4495, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x79, Node: [6]uint8{0x3f, 0x93, 0x60, 0xd8, 0x31, 0xdf}}
	MSFSRMIFsrmSetting                                = uuid.UUID{TimeLow: 0xf411d4fd, TimeMid: 0x14be, TimeHiAndVersion: 0x4260, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x40, Node: [6]uint8{0x3, 0xb7, 0xc9, 0x5e, 0x60, 0x8a}}
	MSFSRMIFsrmStorageModuleDefinition                = uuid.UUID{TimeLow: 0x15a81350, TimeMid: 0x497d, TimeHiAndVersion: 0x4aba, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xe9, Node: [6]uint8{0xd4, 0xdb, 0xcc, 0x55, 0x21, 0xfe}}
	MSFSRVPFileServerVssAgent                         = uuid.UUID{TimeLow: 0xa8e0653c, TimeMid: 0x2744, TimeHiAndVersion: 0x4389, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x1d, Node: [6]uint8{0x73, 0x73, 0xdf, 0x8b, 0x22, 0x92}}
	MSGKDIISDKey                                      = uuid.UUID{TimeLow: 0xb9785960, TimeMid: 0x524f, TimeHiAndVersion: 0x11df, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x6d, Node: [6]uint8{0x83, 0xdc, 0xde, 0xd7, 0x20, 0x85}}
	MSICPRICertPassage                                = uuid.UUID{TimeLow: 0x91ae6020, TimeMid: 0x9e3c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x7c, Node: [6]uint8{0x0, 0xaa, 0x0, 0xc0, 0x91, 0xbe}}
	MSIISSIIisServiceControl                          = uuid.UUID{TimeLow: 0xe8fb8620, TimeMid: 0x588f, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x61, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x79, 0xc5, 0xfe}}
	MSIMSAIIISApplicationAdmin                        = uuid.UUID{TimeLow: 0x7c4e1804, TimeMid: 0xe342, TimeHiAndVersion: 0x483d, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0x3e, Node: [6]uint8{0xa8, 0x50, 0xcf, 0xcc, 0x8d, 0x18}}
	MSIMSAIIISCertObj                                 = uuid.UUID{TimeLow: 0xbd0c73bc, TimeMid: 0x805b, TimeHiAndVersion: 0x4043, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0x30, Node: [6]uint8{0x9a, 0x28, 0xd6, 0x4d, 0xd7, 0xd2}}
	MSIMSAIMSAdminBase2W                              = uuid.UUID{TimeLow: 0x8298d101, TimeMid: 0xf992, TimeHiAndVersion: 0x43b7, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0xca, Node: [6]uint8{0x50, 0x52, 0xd8, 0x85, 0xb9, 0x95}}
	MSIMSAIMSAdminBase3W                              = uuid.UUID{TimeLow: 0xf612954d, TimeMid: 0x3b0b, TimeHiAndVersion: 0x4c56, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0x63, Node: [6]uint8{0x22, 0x7b, 0x7b, 0xe6, 0x24, 0xb4}}
	MSIMSAIMSAdminBaseW                               = uuid.UUID{TimeLow: 0x70b51430, TimeMid: 0xb6ca, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0xb9, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x22, 0xe7, 0x50}}
	MSIMSAIWamAdmin2                                  = uuid.UUID{TimeLow: 0x29822ab8, TimeMid: 0xf302, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x53, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	MSIMSAIWamAdmin                                   = uuid.UUID{TimeLow: 0x29822ab7, TimeMid: 0xf302, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x53, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x19, 0xc1}}
	MSIOIIManagedObject                               = uuid.UUID{TimeLow: 0xc3fcc19e, TimeMid: 0xa970, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x5a, Node: [6]uint8{0x0, 0xa0, 0xc9, 0xb7, 0xc9, 0xc4}}
	MSIOIIRemoteDispatch                              = uuid.UUID{TimeLow: 0x6619a740, TimeMid: 0x8154, TimeHiAndVersion: 0x43be, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x86, Node: [6]uint8{0x3, 0x19, 0x57, 0x8e, 0x2, 0xdb}}
	MSIOIIServicedComponentInfo                       = uuid.UUID{TimeLow: 0x8165b19e, TimeMid: 0x8d3a, TimeHiAndVersion: 0x4d0b, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xc8, Node: [6]uint8{0x97, 0xde, 0x31, 0xd, 0xb5, 0x83}}
	MSIRPInetinfo                                     = uuid.UUID{TimeLow: 0x82ad4280, TimeMid: 0x36b, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x2c, Node: [6]uint8{0x0, 0xaa, 0x0, 0x68, 0x87, 0xb0}}
	MSISTMIDirectoryEnum                              = uuid.UUID{TimeLow: 0x28bc8d5e, TimeMid: 0xca4b, TimeHiAndVersion: 0x4f54, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x3c, Node: [6]uint8{0xed, 0x96, 0x22, 0xd2, 0xb3, 0xac}}
	MSISTMIEnumCachedInitiator                        = uuid.UUID{TimeLow: 0xf093fe3d, TimeMid: 0x8131, TimeHiAndVersion: 0x4b73, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x42, Node: [6]uint8{0xef, 0x54, 0xc2, 0xb, 0x33, 0x7b}}
	MSISTMIEnumConnection                             = uuid.UUID{TimeLow: 0x6aea6b26, TimeMid: 0x680, TimeHiAndVersion: 0x411d, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x77, Node: [6]uint8{0xa1, 0x48, 0xdf, 0x30, 0x87, 0xd5}}
	MSISTMIEnumDisk2                                  = uuid.UUID{TimeLow: 0xa5ecfc73, TimeMid: 0x13, TimeHiAndVersion: 0x4a9e, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0x1c, Node: [6]uint8{0x59, 0xbf, 0x97, 0x35, 0xfd, 0xda}}
	MSISTMIEnumDisk                                   = uuid.UUID{TimeLow: 0xbb39e296, TimeMid: 0xad26, TimeHiAndVersion: 0x42c5, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x90, Node: [6]uint8{0x53, 0x25, 0x33, 0x3b, 0xb1, 0x1e}}
	MSISTMIEnumHost                                   = uuid.UUID{TimeLow: 0xe141fd54, TimeMid: 0xb79e, TimeHiAndVersion: 0x4938, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xbb, Node: [6]uint8{0xd5, 0x23, 0xc3, 0xd4, 0x9f, 0xf1}}
	MSISTMIEnumIDMethod                               = uuid.UUID{TimeLow: 0x345b026b, TimeMid: 0x5802, TimeHiAndVersion: 0x4e38, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0x75, Node: [6]uint8{0x79, 0x5e, 0x8, 0xb0, 0xb8, 0x3f}}
	MSISTMIEnumLMMountPoint                           = uuid.UUID{TimeLow: 0x100da538, TimeMid: 0x3f4a, TimeHiAndVersion: 0x45ab, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x52, Node: [6]uint8{0x70, 0x91, 0x48, 0x15, 0x27, 0x89}}
	MSISTMIEnumPortal                                 = uuid.UUID{TimeLow: 0x1995785d, TimeMid: 0x2a1e, TimeHiAndVersion: 0x492f, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x23, Node: [6]uint8{0xe6, 0x21, 0xea, 0xca, 0x39, 0xd9}}
	MSISTMIEnumResourceGroup                          = uuid.UUID{TimeLow: 0x640038f1, TimeMid: 0xd626, TimeHiAndVersion: 0x40d8, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x2b, Node: [6]uint8{0x9, 0x66, 0x6, 0x1, 0xd0, 0x45}}
	MSISTMIEnumSession2                               = uuid.UUID{TimeLow: 0x442931d5, TimeMid: 0xe522, TimeHiAndVersion: 0x4e64, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x81, Node: [6]uint8{0x74, 0xe9, 0x8a, 0x4e, 0x17, 0x48}}
	MSISTMIEnumSession                                = uuid.UUID{TimeLow: 0x40cc8569, TimeMid: 0x6d23, TimeHiAndVersion: 0x4005, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x58, Node: [6]uint8{0xe3, 0x7f, 0x8, 0xae, 0x19, 0x2b}}
	MSISTMIEnumSnapshot                               = uuid.UUID{TimeLow: 0x66c9b082, TimeMid: 0x7794, TimeHiAndVersion: 0x4948, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x9a, Node: [6]uint8{0xd8, 0xa5, 0xa6, 0x16, 0x37, 0x8f}}
	MSISTMIEnumSnsServer                              = uuid.UUID{TimeLow: 0xe2842c88, TimeMid: 0x7c3, TimeHiAndVersion: 0x4eb0, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0xa9, Node: [6]uint8{0xd3, 0xd9, 0x5e, 0x76, 0xfe, 0xf2}}
	MSISTMIEnumVolume2                                = uuid.UUID{TimeLow: 0x8c58f6b3, TimeMid: 0x4736, TimeHiAndVersion: 0x432a, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x1d, Node: [6]uint8{0x38, 0x9d, 0xe3, 0x50, 0x5c, 0x7c}}
	MSISTMIEnumVolume                                 = uuid.UUID{TimeLow: 0x81fe3594, TimeMid: 0x2495, TimeHiAndVersion: 0x4c91, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0xbb, Node: [6]uint8{0xeb, 0x57, 0x85, 0x61, 0x4e, 0xc7}}
	MSISTMIEnumWTDiskLunMapping                       = uuid.UUID{TimeLow: 0x1396de6f, TimeMid: 0xa794, TimeHiAndVersion: 0x4b11, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x3f, Node: [6]uint8{0x6b, 0x69, 0xa5, 0xb4, 0x7b, 0xae}}
	MSISTMIEnumWTDisk                                 = uuid.UUID{TimeLow: 0x56e65ea5, TimeMid: 0xcdff, TimeHiAndVersion: 0x4391, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x76, Node: [6]uint8{0x0, 0x6e, 0x42, 0xc2, 0xd7, 0x46}}
	MSISTMIHost2                                      = uuid.UUID{TimeLow: 0xb4fa8e86, TimeMid: 0x2517, TimeHiAndVersion: 0x4a88, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x67, Node: [6]uint8{0x75, 0x44, 0x72, 0x19, 0xee, 0xe4}}
	MSISTMIHost3                                      = uuid.UUID{TimeLow: 0xfe7f99f9, TimeMid: 0x1dfb, TimeHiAndVersion: 0x4afb, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x0, Node: [6]uint8{0x6a, 0x8d, 0xd0, 0xaa, 0xbf, 0x2c}}
	MSISTMIHostMgr2                                   = uuid.UUID{TimeLow: 0xdd6f0a28, TimeMid: 0x248f, TimeHiAndVersion: 0x4dd3, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0xe9, Node: [6]uint8{0x71, 0xae, 0xd8, 0xf6, 0x85, 0xc4}}
	MSISTMIHostMgr3                                   = uuid.UUID{TimeLow: 0x1454b97, TimeMid: 0xc6a5, TimeHiAndVersion: 0x4685, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0xa8, Node: [6]uint8{0x97, 0x79, 0xc8, 0x8a, 0xb9, 0x90}}
	MSISTMIHostMgr                                    = uuid.UUID{TimeLow: 0xb06a64e3, TimeMid: 0x814e, TimeHiAndVersion: 0x4ff9, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0xac, Node: [6]uint8{0x59, 0x7a, 0xd3, 0x25, 0x17, 0xc7}}
	MSISTMIHost                                       = uuid.UUID{TimeLow: 0xb0076fec, TimeMid: 0xa921, TimeHiAndVersion: 0x4034, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0xba, Node: [6]uint8{0x9, 0xb, 0xc6, 0xd0, 0x3b, 0xde}}
	MSISTMILocalDeviceMgr2                            = uuid.UUID{TimeLow: 0xb0d1ac4b, TimeMid: 0xf87a, TimeHiAndVersion: 0x49b2, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x8f, Node: [6]uint8{0xd4, 0x39, 0x24, 0x85, 0x75, 0xb2}}
	MSISTMILocalDeviceMgr3                            = uuid.UUID{TimeLow: 0xe645744b, TimeMid: 0xcae5, TimeHiAndVersion: 0x4712, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xaf, Node: [6]uint8{0x13, 0x5, 0x7f, 0x71, 0x95, 0xaf}}
	MSISTMILocalDeviceMgr                             = uuid.UUID{TimeLow: 0x8ad608a4, TimeMid: 0x6c16, TimeHiAndVersion: 0x4405, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x79, Node: [6]uint8{0xb2, 0x79, 0x10, 0xa6, 0x89, 0x95}}
	MSISTMISessionManager2                            = uuid.UUID{TimeLow: 0xc10a76d8, TimeMid: 0x1fe4, TimeHiAndVersion: 0x4c2f, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xd, Node: [6]uint8{0x66, 0x52, 0x65, 0x21, 0x52, 0x59}}
	MSISTMISessionManager                             = uuid.UUID{TimeLow: 0x8d7ae740, TimeMid: 0xb9c5, TimeHiAndVersion: 0x49fc, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x1e, Node: [6]uint8{0x89, 0x17, 0x19, 0x7, 0xcb, 0x86}}
	MSISTMISnapshotMgr                                = uuid.UUID{TimeLow: 0x4e65a71e, TimeMid: 0x4ede, TimeHiAndVersion: 0x4886, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x67, Node: [6]uint8{0x3c, 0x90, 0xa0, 0x8d, 0x1f, 0x29}}
	MSISTMISnapshot                                   = uuid.UUID{TimeLow: 0x883343f1, TimeMid: 0xceed, TimeHiAndVersion: 0x4e3a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x1b, Node: [6]uint8{0xf0, 0xda, 0xdf, 0xce, 0x28, 0x1e}}
	MSISTMISnsMgr                                     = uuid.UUID{TimeLow: 0x1b1c4d1c, TimeMid: 0xabc4, TimeHiAndVersion: 0x4d3a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x22, Node: [6]uint8{0x54, 0x7f, 0xba, 0x3a, 0xa8, 0xa0}}
	MSISTMIStatusNotify                               = uuid.UUID{TimeLow: 0x312cc019, TimeMid: 0xd5cd, TimeHiAndVersion: 0x4ca7, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x10, Node: [6]uint8{0x9e, 0xa, 0x66, 0x1f, 0x14, 0x7e}}
	MSISTMIWTDisk2                                    = uuid.UUID{TimeLow: 0x348a0821, TimeMid: 0x69bb, TimeHiAndVersion: 0x4889, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x1, Node: [6]uint8{0x6a, 0x9b, 0xde, 0x6f, 0xa7, 0x20}}
	MSISTMIWTDisk3                                    = uuid.UUID{TimeLow: 0x1822a95e, TimeMid: 0x1c2b, TimeHiAndVersion: 0x4d02, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x25, Node: [6]uint8{0xcc, 0x11, 0x6d, 0xd9, 0xdb, 0xde}}
	MSISTMIWTDiskMgr2                                 = uuid.UUID{TimeLow: 0x592381e5, TimeMid: 0x8d3c, TimeHiAndVersion: 0x42e9, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xde, Node: [6]uint8{0x4e, 0x77, 0xa1, 0xf7, 0x5a, 0xe4}}
	MSISTMIWTDiskMgr3                                 = uuid.UUID{TimeLow: 0xd6bd6d63, TimeMid: 0xe8cb, TimeHiAndVersion: 0x4905, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x34, Node: [6]uint8{0x8a, 0x27, 0x8c, 0x93, 0x19, 0x7a}}
	MSISTMIWTDiskMgr4                                 = uuid.UUID{TimeLow: 0x703e6b03, TimeMid: 0x7ad1, TimeHiAndVersion: 0x4ded, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xd, Node: [6]uint8{0xe9, 0x4, 0x96, 0xeb, 0xc5, 0xde}}
	MSISTMIWTDiskMgr                                  = uuid.UUID{TimeLow: 0x52ba97e7, TimeMid: 0x9364, TimeHiAndVersion: 0x4134, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0xcb, Node: [6]uint8{0xf8, 0x41, 0x52, 0x13, 0xbd, 0xd8}}
	MSISTMIWTDisk                                     = uuid.UUID{TimeLow: 0x866a78bc, TimeMid: 0xa2fb, TimeHiAndVersion: 0x4ac4, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0xd5, Node: [6]uint8{0xdb, 0x30, 0x41, 0xb4, 0xed, 0x75}}
	MSISTMIWTGeneral2                                 = uuid.UUID{TimeLow: 0x3c73848a, TimeMid: 0xa679, TimeHiAndVersion: 0x40c5, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x1, Node: [6]uint8{0xc9, 0x63, 0xe6, 0x7f, 0x99, 0x49}}
	MSISTMIWTGeneral                                  = uuid.UUID{TimeLow: 0xd71b2cae, TimeMid: 0x33e8, TimeHiAndVersion: 0x4567, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x96, Node: [6]uint8{0x3c, 0xcf, 0x31, 0x62, 0xb, 0xe2}}
	MSLRECNetEventForwarder                           = uuid.UUID{TimeLow: 0x22e5386d, TimeMid: 0x8b12, TimeHiAndVersion: 0x4bf0, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xec, Node: [6]uint8{0x6a, 0x1e, 0xa4, 0x19, 0xe3, 0x66}}
	MSLSADLsarpc                                      = uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}}
	MSLSATLsarpc                                      = uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}}
	MSMQDSDscomm2                                     = uuid.UUID{TimeLow: 0x708cca10, TimeMid: 0x9569, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xa5, Node: [6]uint8{0x0, 0x60, 0x97, 0x7d, 0x81, 0x18}}
	MSMQDSDscomm                                      = uuid.UUID{TimeLow: 0x77df7a80, TimeMid: 0xf298, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x58, Node: [6]uint8{0x0, 0xa0, 0x24, 0xc4, 0x80, 0xa8}}
	MSMQMPQmcomm2                                     = uuid.UUID{TimeLow: 0x76d12b80, TimeMid: 0x3467, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xff, Node: [6]uint8{0x0, 0x90, 0x27, 0x2f, 0x9e, 0xa3}}
	MSMQMPQmcomm                                      = uuid.UUID{TimeLow: 0xfdb3a030, TimeMid: 0x65f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x9b, Node: [6]uint8{0x0, 0xa0, 0x24, 0xea, 0x55, 0x25}}
	MSMQMRQmmgmt                                      = uuid.UUID{TimeLow: 0x41208ee0, TimeMid: 0xe970, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x9e, Node: [6]uint8{0x0, 0xe0, 0x2c, 0x6, 0x4c, 0x39}}
	MSMQQPQm2qm                                       = uuid.UUID{TimeLow: 0x1088a980, TimeMid: 0xeae5, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x9b, Node: [6]uint8{0x0, 0xa0, 0x24, 0x53, 0xc3, 0x37}}
	MSMQRRRemoteRead                                  = uuid.UUID{TimeLow: 0x1a9134dd, TimeMid: 0x7b39, TimeHiAndVersion: 0x45ba, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x88, Node: [6]uint8{0x44, 0xd0, 0x1c, 0xa4, 0x7f, 0x28}}
	MSMSRPMsgsvcsend                                  = uuid.UUID{TimeLow: 0x5a7b91f8, TimeMid: 0xff00, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0xb2, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0xe6, 0xfc}}
	MSMSRPMsgsvc                                      = uuid.UUID{TimeLow: 0x17fdd703, TimeMid: 0x1827, TimeHiAndVersion: 0x4e34, ClockSeqHiAndReserved: 0x79, ClockSeqLow: 0xd4, Node: [6]uint8{0x24, 0xa5, 0x5c, 0x53, 0xbb, 0x37}}
	MSNRPCLogon                                       = uuid.UUID{TimeLow: 0x12345678, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0xcf, 0xfb}}
	MSNSPINspi                                        = uuid.UUID{TimeLow: 0xf5cc5a18, TimeMid: 0x4264, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x59, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2f, 0x84, 0x26}}
	MSOAUTIDispatch                                   = uuid.UUID{TimeLow: 0x20400, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTIEnumVARIANT                                = uuid.UUID{TimeLow: 0x20404, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTITypeComp                                   = uuid.UUID{TimeLow: 0x20403, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTITypeInfo2                                  = uuid.UUID{TimeLow: 0x20412, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTITypeInfo                                   = uuid.UUID{TimeLow: 0x20401, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTITypeLib2                                   = uuid.UUID{TimeLow: 0x20411, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTITypeLib                                    = uuid.UUID{TimeLow: 0x20402, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOAUTIUnknown                                    = uuid.UUID{TimeLow: 0x0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0x0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSOCSPAIOCSPAdminD                                = uuid.UUID{TimeLow: 0x784b693d, TimeMid: 0x95f3, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x26, Node: [6]uint8{0x36, 0x5c, 0x9, 0x86, 0x59, 0xf2}}
	MSPANIRPCAsyncNotify                              = uuid.UUID{TimeLow: 0xb6edbfa, TimeMid: 0x4a24, TimeHiAndVersion: 0x4fc6, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x23, Node: [6]uint8{0x94, 0x2b, 0x1e, 0xca, 0x65, 0xd1}}
	MSPANIRPCRemoteObject                             = uuid.UUID{TimeLow: 0xae33069b, TimeMid: 0xa2a8, TimeHiAndVersion: 0x46ee, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x35, Node: [6]uint8{0xdd, 0xfd, 0x33, 0x9b, 0xe2, 0x81}}
	MSPARIRemoteWinspool                              = uuid.UUID{TimeLow: 0x76f03f96, TimeMid: 0xcdfd, TimeHiAndVersion: 0x44fc, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x2c, Node: [6]uint8{0x64, 0x95, 0xa, 0x0, 0x12, 0x9}}
	MSPCQPerflibV2                                    = uuid.UUID{TimeLow: 0xda5a86c5, TimeMid: 0x12c2, TimeHiAndVersion: 0x4943, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x30, Node: [6]uint8{0x7f, 0x74, 0xa8, 0x13, 0xd8, 0x53}}
	MSPLAIAlertDataCollector                          = uuid.UUID{TimeLow: 0x3837516, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIApiTracingDataCollector                     = uuid.UUID{TimeLow: 0x383751a, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIConfigurationDataCollector                  = uuid.UUID{TimeLow: 0x3837514, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIDataCollectorCollection                     = uuid.UUID{TimeLow: 0x3837502, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIDataCollectorSetCollection                  = uuid.UUID{TimeLow: 0x3837524, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIDataCollectorSet                            = uuid.UUID{TimeLow: 0x3837520, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIDataCollector                               = uuid.UUID{TimeLow: 0x38374ff, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIDataManager                                 = uuid.UUID{TimeLow: 0x3837541, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIFolderActionCollection                      = uuid.UUID{TimeLow: 0x3837544, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIFolderAction                                = uuid.UUID{TimeLow: 0x3837543, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIPerformanceCounterDataCollector             = uuid.UUID{TimeLow: 0x3837506, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIScheduleCollection                          = uuid.UUID{TimeLow: 0x383753d, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAISchedule                                    = uuid.UUID{TimeLow: 0x383753a, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAITraceDataCollector                          = uuid.UUID{TimeLow: 0x383750b, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAITraceDataProviderCollection                 = uuid.UUID{TimeLow: 0x3837510, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAITraceDataProvider                           = uuid.UUID{TimeLow: 0x3837512, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIValueMapItem                                = uuid.UUID{TimeLow: 0x3837533, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSPLAIValueMap                                    = uuid.UUID{TimeLow: 0x3837534, TimeMid: 0x98b, TimeHiAndVersion: 0x11d8, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x14, Node: [6]uint8{0x50, 0x50, 0x54, 0x50, 0x30, 0x30}}
	MSRAAAuthzr                                       = uuid.UUID{TimeLow: 0xb1c2170, TimeMid: 0x5732, TimeHiAndVersion: 0x4e0e, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0xd3, Node: [6]uint8{0xd9, 0xb1, 0x6f, 0x3b, 0x84, 0xd7}}
	MSRAAIgnoreCentralPolicy                          = uuid.UUID{TimeLow: 0x5fc860e0, TimeMid: 0x6f6e, TimeHiAndVersion: 0x4fc2, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0xcd, Node: [6]uint8{0x46, 0x32, 0x4f, 0x25, 0xe9, 0xb}}
	MSRAAUseCentralPolicy                             = uuid.UUID{TimeLow: 0x9a81c2bd, TimeMid: 0xa525, TimeHiAndVersion: 0x471d, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xed, Node: [6]uint8{0x49, 0x90, 0x7c, 0xb, 0x23, 0xda}}
	MSRAINPSIIASDataStoreComServer2                   = uuid.UUID{TimeLow: 0xc323be28, TimeMid: 0xe546, TimeHiAndVersion: 0x4c23, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0x1b, Node: [6]uint8{0xd6, 0xad, 0x8d, 0x8f, 0xac, 0x7b}}
	MSRAINPSIIASDataStoreComServer                    = uuid.UUID{TimeLow: 0x83e05bd5, TimeMid: 0xaec1, TimeHiAndVersion: 0x4e58, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x50, Node: [6]uint8{0xe8, 0x19, 0xc7, 0x29, 0x6f, 0x67}}
	MSRAIIPCHCollection                               = uuid.UUID{TimeLow: 0x833e4100, TimeMid: 0xaff7, TimeHiAndVersion: 0x4ac3, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc2, Node: [6]uint8{0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}
	MSRAIIPCHService                                  = uuid.UUID{TimeLow: 0x833e4200, TimeMid: 0xaff7, TimeHiAndVersion: 0x4ac3, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc2, Node: [6]uint8{0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}
	MSRAIIRASrv                                       = uuid.UUID{TimeLow: 0xf120a684, TimeMid: 0xb926, TimeHiAndVersion: 0x447f, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xf4, Node: [6]uint8{0xc9, 0x66, 0xcb, 0x78, 0x56, 0x48}}
	MSRAIISAFSession                                  = uuid.UUID{TimeLow: 0x833e41aa, TimeMid: 0xaff7, TimeHiAndVersion: 0x4ac3, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc2, Node: [6]uint8{0x9f, 0x24, 0xc1, 0x45, 0x7b, 0xce}}
	MSRAIWWinsif2                                     = uuid.UUID{TimeLow: 0x811109bf, TimeMid: 0xa4e1, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x54, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x1e, 0x9b, 0x45}}
	MSRAIWWinsif                                      = uuid.UUID{TimeLow: 0x45f52c28, TimeMid: 0x7f9f, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x2b, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2e, 0xfa, 0xbe}}
	MSRDPESCTypeScardPack                             = uuid.UUID{TimeLow: 0xa35af600, TimeMid: 0x9cf4, TimeHiAndVersion: 0x11cd, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x76, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2b, 0xd7, 0x11}}
	MSRPCEEndpointMapper                              = uuid.UUID{TimeLow: 0xe1af8308, TimeMid: 0x5d1f, TimeHiAndVersion: 0x11c9, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xa4, Node: [6]uint8{0x8, 0x0, 0x2b, 0x14, 0xa0, 0xfa}}
	MSRPCERemoteManagement                            = uuid.UUID{TimeLow: 0xafa8bd80, TimeMid: 0x7d8a, TimeHiAndVersion: 0x11c9, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0xf4, Node: [6]uint8{0x8, 0x0, 0x2b, 0x10, 0x29, 0x89}}
	MSRPCETransferNDR64                               = uuid.UUID{TimeLow: 0x71710533, TimeMid: 0xbeba, TimeHiAndVersion: 0x4937, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x19, Node: [6]uint8{0xb5, 0xdb, 0xef, 0x9c, 0xcc, 0x36}}
	MSRPCETransferNDR                                 = uuid.UUID{TimeLow: 0x8a885d04, TimeMid: 0x1ceb, TimeHiAndVersion: 0x11c9, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xe8, Node: [6]uint8{0x8, 0x0, 0x2b, 0x10, 0x48, 0x60}}
	MSRPCLLocToLoc                                    = uuid.UUID{TimeLow: 0xe33c0cc4, TimeMid: 0x482, TimeHiAndVersion: 0x101a, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0xc, Node: [6]uint8{0x2, 0x60, 0x8c, 0x6b, 0xa2, 0x18}}
	MSRPRNWinspool                                    = uuid.UUID{TimeLow: 0x12345678, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xab}}
	MSRRASMDIMSVC                                     = uuid.UUID{TimeLow: 0x8f09f000, TimeMid: 0xb7ed, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xd2, Node: [6]uint8{0x0, 0x0, 0x1a, 0x18, 0x1c, 0xad}}
	MSRRASMIRemoteICFICSConfig                        = uuid.UUID{TimeLow: 0x66a2db22, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	MSRRASMIRemoteIPV6Config                          = uuid.UUID{TimeLow: 0x6139d8a4, TimeMid: 0xe508, TimeHiAndVersion: 0x4ebb, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xc7, Node: [6]uint8{0xd7, 0xf2, 0x75, 0x14, 0x58, 0x97}}
	MSRRASMIRemoteNetworkConfig                       = uuid.UUID{TimeLow: 0x66a2db1b, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	MSRRASMIRemoteRouterRestart                       = uuid.UUID{TimeLow: 0x66a2db20, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	MSRRASMIRemoteSetDnsConfig                        = uuid.UUID{TimeLow: 0x66a2db21, TimeMid: 0xd706, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc9, 0xda, 0x4}}
	MSRRASMIRemoteSstpCertCheck                       = uuid.UUID{TimeLow: 0x5ff9bdf6, TimeMid: 0xbd91, TimeHiAndVersion: 0x4d8b, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0x14, Node: [6]uint8{0xd6, 0x31, 0x7a, 0xcc, 0x8d, 0xd8}}
	MSRRASMIRemoteStringIdConfig                      = uuid.UUID{TimeLow: 0x67e08fc2, TimeMid: 0x2984, TimeHiAndVersion: 0x4b62, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x2e, Node: [6]uint8{0xfc, 0x1a, 0xae, 0x64, 0xbb, 0xbb}}
	MSRRASMRASRPC                                     = uuid.UUID{TimeLow: 0x20610036, TimeMid: 0xfa22, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x23, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x11, 0xe5, 0xdf}}
	MSRRPWinreg                                       = uuid.UUID{TimeLow: 0x338cd001, TimeMid: 0x2244, TimeHiAndVersion: 0x31f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xaa, Node: [6]uint8{0x90, 0x0, 0x38, 0x0, 0x10, 0x3}}
	MSRSMPIClientSink                                 = uuid.UUID{TimeLow: 0x879c8bbe, TimeMid: 0x41b0, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x11, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0xbf, 0x70}}
	MSRSMPIMessenger                                  = uuid.UUID{TimeLow: 0x81e7188, TimeMid: 0xc080, TimeHiAndVersion: 0x4ff3, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0x38, Node: [6]uint8{0x29, 0xf6, 0x6d, 0x6c, 0xab, 0xfd}}
	MSRSMPINtmsLibraryControl1                        = uuid.UUID{TimeLow: 0x4e934f30, TimeMid: 0x341a, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	MSRSMPINtmsLibraryControl2                        = uuid.UUID{TimeLow: 0xdb90832f, TimeMid: 0x6910, TimeHiAndVersion: 0x4d46, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x5e, Node: [6]uint8{0x9f, 0xd6, 0xbf, 0xa7, 0x39, 0x3}}
	MSRSMPINtmsMediaServices1                         = uuid.UUID{TimeLow: 0xd02e4be0, TimeMid: 0x3419, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	MSRSMPINtmsNotifySink                             = uuid.UUID{TimeLow: 0xbb39332c, TimeMid: 0xbfee, TimeHiAndVersion: 0x4380, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x8a, Node: [6]uint8{0xba, 0xdc, 0x8a, 0xff, 0x5b, 0xb6}}
	MSRSMPINtmsObjectInfo1                            = uuid.UUID{TimeLow: 0x69ab7050, TimeMid: 0x3059, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xaf, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	MSRSMPINtmsObjectManagement1                      = uuid.UUID{TimeLow: 0xb057dc50, TimeMid: 0x3059, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xaf, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	MSRSMPINtmsObjectManagement2                      = uuid.UUID{TimeLow: 0x895a2c86, TimeMid: 0x270d, TimeHiAndVersion: 0x489d, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xc0, Node: [6]uint8{0xdc, 0x2a, 0x9b, 0x35, 0x28, 0xe}}
	MSRSMPINtmsObjectManagement3                      = uuid.UUID{TimeLow: 0x3bbed8d9, TimeMid: 0x2c9a, TimeHiAndVersion: 0x4b21, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x36, Node: [6]uint8{0xac, 0xb2, 0xf9, 0x95, 0xbe, 0x6c}}
	MSRSMPINtmsSession1                               = uuid.UUID{TimeLow: 0x8da03f40, TimeMid: 0x3419, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xb1, Node: [6]uint8{0x0, 0xa0, 0x24, 0xcb, 0x60, 0x19}}
	MSRSMPIRobustNtmsMediaServices1                   = uuid.UUID{TimeLow: 0x7d07f313, TimeMid: 0xa53f, TimeHiAndVersion: 0x459a, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x12, Node: [6]uint8{0x1, 0x2c, 0x15, 0xb1, 0x84, 0x6e}}
	MSRSMPIUnknown                                    = uuid.UUID{TimeLow: 0x0, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	MSRSPInitShutdown                                 = uuid.UUID{TimeLow: 0x894de0c0, TimeMid: 0xd55, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x22, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x21, 0xa1}}
	MSRSPWindowsShutdown                              = uuid.UUID{TimeLow: 0xd95afe70, TimeMid: 0xa6d5, TimeHiAndVersion: 0x4259, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2e, Node: [6]uint8{0x2c, 0x84, 0xda, 0x1d, 0xdb, 0xd}}
	MSRSPWinreg                                       = uuid.UUID{TimeLow: 0x338cd001, TimeMid: 0x2244, TimeHiAndVersion: 0x31f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xaa, Node: [6]uint8{0x90, 0x0, 0x38, 0x0, 0x10, 0x3}}
	MSSAMRSamr                                        = uuid.UUID{TimeLow: 0x12345778, TimeMid: 0x1234, TimeHiAndVersion: 0xabcd, ClockSeqHiAndReserved: 0xef, ClockSeqLow: 0x0, Node: [6]uint8{0x1, 0x23, 0x45, 0x67, 0x89, 0xac}}
	MSSCMPIVssDifferentialSoftwareSnapshotMgmt        = uuid.UUID{TimeLow: 0x214a0f28, TimeMid: 0xb737, TimeHiAndVersion: 0x4026, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x47, Node: [6]uint8{0x4f, 0x9e, 0x37, 0xd7, 0x95, 0x29}}
	MSSCMPIVssEnumMgmtObject                          = uuid.UUID{TimeLow: 0x1954e6b, TimeMid: 0x9254, TimeHiAndVersion: 0x4e6e, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x8c, Node: [6]uint8{0xc9, 0xe0, 0x5d, 0x0, 0x76, 0x96}}
	MSSCMPIVssEnumObject                              = uuid.UUID{TimeLow: 0xae1c7110, TimeMid: 0x2f60, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x39, Node: [6]uint8{0x0, 0xc0, 0x4f, 0x72, 0xd8, 0xe3}}
	MSSCMPIVssSnapshotMgmt                            = uuid.UUID{TimeLow: 0xfa7df749, TimeMid: 0x66e7, TimeHiAndVersion: 0x4986, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x7f, Node: [6]uint8{0xe2, 0xf0, 0x4a, 0xe5, 0x37, 0x72}}
	MSSCMRSvcctl                                      = uuid.UUID{TimeLow: 0x367abb81, TimeMid: 0x9844, TimeHiAndVersion: 0x35f1, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x32, Node: [6]uint8{0x98, 0xf0, 0x38, 0x0, 0x10, 0x3}}
	MSSRVSSrvsvc                                      = uuid.UUID{TimeLow: 0x4b324fc8, TimeMid: 0x1670, TimeHiAndVersion: 0x1d3, ClockSeqHiAndReserved: 0x12, ClockSeqLow: 0x78, Node: [6]uint8{0x5a, 0x47, 0xbf, 0x6e, 0xe1, 0x88}}
	MSSWNWitness                                      = uuid.UUID{TimeLow: 0xccd8c074, TimeMid: 0xd0e5, TimeHiAndVersion: 0x4a40, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xb4, Node: [6]uint8{0xd0, 0x74, 0xfa, 0xa6, 0xba, 0x28}}
	MSTPMVSCITpmVirtualSmartCardManager2              = uuid.UUID{TimeLow: 0xfdf8a2b9, TimeMid: 0x2de, TimeHiAndVersion: 0x47f4, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x26, Node: [6]uint8{0xaa, 0x85, 0xab, 0x5e, 0x52, 0x67}}
	MSTPMVSCITpmVirtualSmartCardManager3              = uuid.UUID{TimeLow: 0x3c745a97, TimeMid: 0xf375, TimeHiAndVersion: 0x4150, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x17, Node: [6]uint8{0x59, 0x50, 0xf6, 0x94, 0xc6, 0x99}}
	MSTPMVSCITpmVirtualSmartCardManagerStatusCallback = uuid.UUID{TimeLow: 0x1a1bb35f, TimeMid: 0xabb8, TimeHiAndVersion: 0x451c, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xae, Node: [6]uint8{0x33, 0xd9, 0x8f, 0x1b, 0xef, 0x4a}}
	MSTPMVSCITpmVirtualSmartCardManager               = uuid.UUID{TimeLow: 0x112b1dff, TimeMid: 0xd9dc, TimeHiAndVersion: 0x41f7, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x9f, Node: [6]uint8{0xd6, 0x7f, 0xee, 0x7c, 0xb5, 0x91}}
	MSTRPRemotesp                                     = uuid.UUID{TimeLow: 0x2f5f6521, TimeMid: 0xca47, TimeHiAndVersion: 0x1068, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x19, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xdb}}
	MSTRPTapsrv                                       = uuid.UUID{TimeLow: 0x2f5f6520, TimeMid: 0xca46, TimeHiAndVersion: 0x1067, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x19, Node: [6]uint8{0x0, 0xdd, 0x1, 0x6, 0x62, 0xda}}
	MSTSCHATSvc                                       = uuid.UUID{TimeLow: 0x1ff70682, TimeMid: 0xa51, TimeHiAndVersion: 0x30e8, ClockSeqHiAndReserved: 0x7, ClockSeqLow: 0x6d, Node: [6]uint8{0x74, 0xb, 0xe8, 0xce, 0xe9, 0x8b}}
	MSTSCHITaskSchedulerService                       = uuid.UUID{TimeLow: 0x86d35949, TimeMid: 0x83c9, TimeHiAndVersion: 0x4044, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x24, Node: [6]uint8{0xdb, 0x36, 0x32, 0x31, 0xfd, 0xc}}
	MSTSCHSASec                                       = uuid.UUID{TimeLow: 0x378e52b0, TimeMid: 0xc0a9, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x2d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x51, 0xe4, 0xf}}
	MSTSGUTsProxyRpcInterface                         = uuid.UUID{TimeLow: 0x44e265dd, TimeMid: 0x7daf, TimeHiAndVersion: 0x42cd, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x60, Node: [6]uint8{0x3c, 0xdb, 0x6e, 0x7a, 0x27, 0x29}}
	MSTSRAPIManageTelnetSessions                      = uuid.UUID{TimeLow: 0x34634fd, TimeMid: 0xba3f, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x6a, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x44, 0x13, 0x8c}}
	MSTSTSIcaApi                                      = uuid.UUID{TimeLow: 0x5ca4a760, TimeMid: 0xebb1, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x11, Node: [6]uint8{0x0, 0xa0, 0x24, 0x54, 0x20, 0xed}}
	MSTSTSRCMListener                                 = uuid.UUID{TimeLow: 0x497d95a6, TimeMid: 0x2d27, TimeHiAndVersion: 0x4bf5, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0xbd, Node: [6]uint8{0xa6, 0x4, 0x69, 0x57, 0x13, 0x3c}}
	MSTSTSRCMPublic                                   = uuid.UUID{TimeLow: 0xbde95fdf, TimeMid: 0xeee0, TimeHiAndVersion: 0x45de, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x12, Node: [6]uint8{0xe5, 0xa6, 0x1c, 0xd0, 0xd4, 0xfe}}
	MSTSTSSessEnvPublicRpc                            = uuid.UUID{TimeLow: 0x1257b580, TimeMid: 0xce2f, TimeHiAndVersion: 0x4109, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0xd6, Node: [6]uint8{0xa9, 0x45, 0x9d, 0xb, 0xf6, 0xbc}}
	MSTSTSTermSrvEnumeration                          = uuid.UUID{TimeLow: 0x88143fd0, TimeMid: 0xc28d, TimeHiAndVersion: 0x4b2b, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xef, Node: [6]uint8{0x8d, 0x88, 0x2f, 0x6a, 0x93, 0x90}}
	MSTSTSTermSrvNotification                         = uuid.UUID{TimeLow: 0x11899a43, TimeMid: 0x2b68, TimeHiAndVersion: 0x4a76, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xe3, Node: [6]uint8{0xa3, 0xd6, 0xad, 0x8c, 0x26, 0xce}}
	MSTSTSTermSrvSession                              = uuid.UUID{TimeLow: 0x484809d6, TimeMid: 0x4239, TimeHiAndVersion: 0x471b, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0xbc, Node: [6]uint8{0x61, 0xdf, 0x8c, 0x23, 0xac, 0x48}}
	MSTSTSTSVIPPublic                                 = uuid.UUID{TimeLow: 0x53b46b02, TimeMid: 0xc73b, TimeHiAndVersion: 0x4a3e, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0xee, Node: [6]uint8{0xb1, 0x6b, 0x80, 0x67, 0x2f, 0xc0}}
	MSUAMGIAutomaticUpdates2                          = uuid.UUID{TimeLow: 0x4a2f5c31, TimeMid: 0xcfd9, TimeHiAndVersion: 0x410e, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0xfb, Node: [6]uint8{0x29, 0xa6, 0x53, 0x97, 0x3a, 0xf}}
	MSUAMGIAutomaticUpdatesResults                    = uuid.UUID{TimeLow: 0xe7a4d634, TimeMid: 0x7942, TimeHiAndVersion: 0x4dd9, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x11, Node: [6]uint8{0x82, 0x22, 0x8b, 0xa3, 0x39, 0x1}}
	MSUAMGIAutomaticUpdates                           = uuid.UUID{TimeLow: 0x673425bf, TimeMid: 0xc082, TimeHiAndVersion: 0x4c7c, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xfd, Node: [6]uint8{0x56, 0x94, 0x64, 0xb8, 0xe0, 0xce}}
	MSUAMGICategoryCollection                         = uuid.UUID{TimeLow: 0x3a56bfb8, TimeMid: 0x576c, TimeHiAndVersion: 0x43f7, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x35, Node: [6]uint8{0xfe, 0x48, 0x38, 0xfd, 0x7e, 0x37}}
	MSUAMGICategory                                   = uuid.UUID{TimeLow: 0x81ddc1b8, TimeMid: 0x9d35, TimeHiAndVersion: 0x47a6, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x71, Node: [6]uint8{0x5b, 0x80, 0xf5, 0x19, 0x22, 0x3b}}
	MSUAMGIImageInformation                           = uuid.UUID{TimeLow: 0x7c907864, TimeMid: 0x346c, TimeHiAndVersion: 0x4aeb, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x3f, Node: [6]uint8{0x57, 0xda, 0x28, 0x9f, 0x96, 0x9f}}
	MSUAMGIInstallationBehavior                       = uuid.UUID{TimeLow: 0xd9a59339, TimeMid: 0xe245, TimeHiAndVersion: 0x4dbd, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x86, Node: [6]uint8{0x4d, 0x57, 0x63, 0xe3, 0x96, 0x24}}
	MSUAMGISearchJob_                                 = uuid.UUID{TimeLow: 0x7366ea16, TimeMid: 0x7a1a, TimeHiAndVersion: 0x4ea2, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x42, Node: [6]uint8{0x97, 0x3d, 0x3e, 0x9c, 0xd9, 0x9b}}
	MSUAMGISearchResult                               = uuid.UUID{TimeLow: 0xd40cff62, TimeMid: 0xe08c, TimeHiAndVersion: 0x4498, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x1a, Node: [6]uint8{0x1, 0xe2, 0x5f, 0xf, 0xd3, 0x3c}}
	MSUAMGIStringCollection                           = uuid.UUID{TimeLow: 0xeff90582, TimeMid: 0x2ddc, TimeHiAndVersion: 0x480f, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x6d, Node: [6]uint8{0x60, 0xf3, 0xfb, 0xc3, 0x62, 0xc3}}
	MSUAMGIUpdate2                                    = uuid.UUID{TimeLow: 0x144fe9b0, TimeMid: 0xd23d, TimeHiAndVersion: 0x4a8b, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x34, Node: [6]uint8{0xfb, 0x44, 0x57, 0x53, 0x3b, 0x7a}}
	MSUAMGIUpdate3                                    = uuid.UUID{TimeLow: 0x112eda6b, TimeMid: 0x95b3, TimeHiAndVersion: 0x476f, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x90, Node: [6]uint8{0xae, 0xe8, 0x2c, 0x6b, 0x81, 0x81}}
	MSUAMGIUpdate4                                    = uuid.UUID{TimeLow: 0x27e94b0d, TimeMid: 0x5139, TimeHiAndVersion: 0x49a2, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x61, Node: [6]uint8{0x93, 0x52, 0x2d, 0xc5, 0x46, 0x52}}
	MSUAMGIUpdate5                                    = uuid.UUID{TimeLow: 0xc1c2f21a, TimeMid: 0xd2f4, TimeHiAndVersion: 0x4902, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0xc6, Node: [6]uint8{0x8a, 0x8, 0x1c, 0x19, 0xa8, 0x90}}
	MSUAMGIUpdateCollection                           = uuid.UUID{TimeLow: 0x7f7438c, TimeMid: 0x7709, TimeHiAndVersion: 0x4ca5, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x18, Node: [6]uint8{0x91, 0x27, 0x92, 0x88, 0x13, 0x4e}}
	MSUAMGIUpdateDownloadContent2                     = uuid.UUID{TimeLow: 0xc97ad11b, TimeMid: 0xf257, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0x9f, Node: [6]uint8{0x37, 0x7f, 0x73, 0x3f, 0x6f, 0x68}}
	MSUAMGIUpdateDownloadContentCollection            = uuid.UUID{TimeLow: 0xbc5513c8, TimeMid: 0xb3b8, TimeHiAndVersion: 0x4bf7, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xd4, Node: [6]uint8{0x36, 0x1c, 0xd, 0x8c, 0x88, 0xba}}
	MSUAMGIUpdateDownloadContent                      = uuid.UUID{TimeLow: 0x54a2cb2d, TimeMid: 0x9a0c, TimeHiAndVersion: 0x48b6, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x50, Node: [6]uint8{0x9a, 0xbb, 0x69, 0xee, 0x2d, 0x2}}
	MSUAMGIUpdateExceptionCollectionOLE               = uuid.UUID{TimeLow: 0x503626a3, TimeMid: 0x8e14, TimeHiAndVersion: 0x4729, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x55, Node: [6]uint8{0xf, 0xe6, 0x64, 0xbd, 0x23, 0x21}}
	MSUAMGIUpdateExceptionCollection                  = uuid.UUID{TimeLow: 0xa37d00f5, TimeMid: 0x7bb0, TimeHiAndVersion: 0x4953, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x14, Node: [6]uint8{0xf9, 0xe9, 0x83, 0x26, 0xf2, 0xe8}}
	MSUAMGIUpdateException                            = uuid.UUID{TimeLow: 0xa376dd5e, TimeMid: 0x9d4, TimeHiAndVersion: 0x427f, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0x7c, Node: [6]uint8{0xfe, 0xd5, 0xb6, 0xe1, 0xc1, 0xd6}}
	MSUAMGIUpdateHistoryEntry2                        = uuid.UUID{TimeLow: 0xc2bfb780, TimeMid: 0x4539, TimeHiAndVersion: 0x4132, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x8c, Node: [6]uint8{0xa, 0x87, 0x72, 0x1, 0x3a, 0xb6}}
	MSUAMGIUpdateHistoryEntryCollection               = uuid.UUID{TimeLow: 0xa7f04f3c, TimeMid: 0xa290, TimeHiAndVersion: 0x435b, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xdf, Node: [6]uint8{0xa1, 0x16, 0xc3, 0x35, 0x7a, 0x5c}}
	MSUAMGIUpdateHistoryEntry                         = uuid.UUID{TimeLow: 0xbe56a644, TimeMid: 0xaf0e, TimeHiAndVersion: 0x4e0e, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x11, Node: [6]uint8{0xc1, 0xd8, 0xe6, 0x95, 0xcb, 0xff}}
	MSUAMGIUpdateIdentity                             = uuid.UUID{TimeLow: 0x46297823, TimeMid: 0x9940, TimeHiAndVersion: 0x4c09, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0xd9, Node: [6]uint8{0xcd, 0x3e, 0xa6, 0xd0, 0x59, 0x68}}
	MSUAMGIUpdateSearcher2                            = uuid.UUID{TimeLow: 0x4cbdcb2d, TimeMid: 0x1589, TimeHiAndVersion: 0x4beb, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0x1c, Node: [6]uint8{0x3e, 0x58, 0x2f, 0xf0, 0xad, 0xd0}}
	MSUAMGIUpdateSearcher3                            = uuid.UUID{TimeLow: 0x4c6895d, TimeMid: 0xeaf2, TimeHiAndVersion: 0x4034, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xf3, Node: [6]uint8{0x31, 0x1d, 0xe9, 0xbe, 0x41, 0x3a}}
	MSUAMGIUpdateSearcher                             = uuid.UUID{TimeLow: 0x8f45abf1, TimeMid: 0xf9ae, TimeHiAndVersion: 0x4b95, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0x33, Node: [6]uint8{0xf0, 0xf6, 0x6e, 0x50, 0x56, 0xea}}
	MSUAMGIUpdateService2                             = uuid.UUID{TimeLow: 0x1518b460, TimeMid: 0x6518, TimeHiAndVersion: 0x4172, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0xf, Node: [6]uint8{0xc7, 0x58, 0x83, 0xb2, 0x4c, 0xeb}}
	MSUAMGIUpdateServiceCollection                    = uuid.UUID{TimeLow: 0x9b0353aa, TimeMid: 0xe52, TimeHiAndVersion: 0x44ff, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0xb0, Node: [6]uint8{0x1f, 0x7f, 0xa0, 0x43, 0x7f, 0x88}}
	MSUAMGIUpdateServiceManager2                      = uuid.UUID{TimeLow: 0xbb8531d, TimeMid: 0x7e8d, TimeHiAndVersion: 0x424f, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x6c, Node: [6]uint8{0xa0, 0xb8, 0xf6, 0xa, 0x3e, 0x7b}}
	MSUAMGIUpdateServiceManager                       = uuid.UUID{TimeLow: 0x23857e3c, TimeMid: 0x2ba, TimeHiAndVersion: 0x44a3, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x23, Node: [6]uint8{0xb1, 0xc9, 0x0, 0x80, 0x5f, 0x37}}
	MSUAMGIUpdateServiceRegistration                  = uuid.UUID{TimeLow: 0xdde02280, TimeMid: 0x12b3, TimeHiAndVersion: 0x4e0b, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x7b, Node: [6]uint8{0x67, 0x47, 0xf6, 0xac, 0xb2, 0x86}}
	MSUAMGIUpdateService                              = uuid.UUID{TimeLow: 0x76b3b17e, TimeMid: 0xaed6, TimeHiAndVersion: 0x4da5, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xf0, Node: [6]uint8{0x83, 0x58, 0x7f, 0x81, 0xab, 0xe3}}
	MSUAMGIUpdateSession2                             = uuid.UUID{TimeLow: 0x91caf7b0, TimeMid: 0xeb23, TimeHiAndVersion: 0x49ed, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x37, Node: [6]uint8{0xc5, 0x2d, 0x81, 0x7f, 0x46, 0xf7}}
	MSUAMGIUpdateSession3                             = uuid.UUID{TimeLow: 0x918efd1e, TimeMid: 0xb5d8, TimeHiAndVersion: 0x4c90, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x40, Node: [6]uint8{0xae, 0xb9, 0xbd, 0xc5, 0x6f, 0x9d}}
	MSUAMGIUpdateSession                              = uuid.UUID{TimeLow: 0x816858a4, TimeMid: 0x260d, TimeHiAndVersion: 0x4260, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x3a, Node: [6]uint8{0x25, 0x85, 0xf1, 0xab, 0xc7, 0x6b}}
	MSUAMGIUpdate                                     = uuid.UUID{TimeLow: 0x6a92b07a, TimeMid: 0xd821, TimeHiAndVersion: 0x4682, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x23, Node: [6]uint8{0x5c, 0x80, 0x50, 0x22, 0xcc, 0x4d}}
	MSUAMGIWindowsDriverUpdate2                       = uuid.UUID{TimeLow: 0x615c4269, TimeMid: 0x7a48, TimeHiAndVersion: 0x43bd, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0xb7, Node: [6]uint8{0xbf, 0x6c, 0xa2, 0x7d, 0x6c, 0x3e}}
	MSUAMGIWindowsDriverUpdate3                       = uuid.UUID{TimeLow: 0x49ebd502, TimeMid: 0x4a96, TimeHiAndVersion: 0x41bd, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x3e, Node: [6]uint8{0x4c, 0x50, 0x57, 0xf4, 0x25, 0xc}}
	MSUAMGIWindowsDriverUpdate4                       = uuid.UUID{TimeLow: 0x4c6a2b, TimeMid: 0xc19, TimeHiAndVersion: 0x4c69, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x5c, Node: [6]uint8{0xa2, 0x69, 0xb2, 0x56, 0xd, 0xb9}}
	MSUAMGIWindowsDriverUpdate5                       = uuid.UUID{TimeLow: 0x70cf5c82, TimeMid: 0x8642, TimeHiAndVersion: 0x42bb, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xbc, Node: [6]uint8{0xc, 0xfd, 0x26, 0x3c, 0x6c, 0x4f}}
	MSUAMGIWindowsDriverUpdateEntryCollection         = uuid.UUID{TimeLow: 0xd521700, TimeMid: 0xa372, TimeHiAndVersion: 0x4bef, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x8b, Node: [6]uint8{0x3d, 0x0, 0xc1, 0xa, 0xde, 0xbd}}
	MSUAMGIWindowsDriverUpdateEntry                   = uuid.UUID{TimeLow: 0xed8bfe40, TimeMid: 0xa60b, TimeHiAndVersion: 0x42ea, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x52, Node: [6]uint8{0x81, 0x7d, 0xfc, 0xfa, 0x23, 0xec}}
	MSUAMGIWindowsDriverUpdate                        = uuid.UUID{TimeLow: 0xb383cd1a, TimeMid: 0x5ce9, TimeHiAndVersion: 0x4504, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x63, Node: [6]uint8{0x76, 0x4b, 0x12, 0x36, 0xf1, 0x91}}
	MSUAMGIWindowsUpdateAgentInfo                     = uuid.UUID{TimeLow: 0x85713fa1, TimeMid: 0x7796, TimeHiAndVersion: 0x4fa2, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x3b, Node: [6]uint8{0xe2, 0xd6, 0x12, 0x4d, 0xd3, 0x73}}
	MSVDSIEnumVdsObject                               = uuid.UUID{TimeLow: 0x118610b7, TimeMid: 0x8d94, TimeHiAndVersion: 0x4030, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0xb8, Node: [6]uint8{0x50, 0x8, 0x89, 0x78, 0x8e, 0x4e}}
	MSVDSIVdsAdvancedDisk2                            = uuid.UUID{TimeLow: 0x9723f420, TimeMid: 0x9355, TimeHiAndVersion: 0x42de, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x66, Node: [6]uint8{0xe3, 0x1b, 0xb1, 0x5b, 0xee, 0xac}}
	MSVDSIVdsAdvancedDisk3                            = uuid.UUID{TimeLow: 0x3858c0d5, TimeMid: 0xf35, TimeHiAndVersion: 0x4bf5, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x14, Node: [6]uint8{0x69, 0x87, 0x49, 0x63, 0xbc, 0x36}}
	MSVDSIVdsAdvancedDisk                             = uuid.UUID{TimeLow: 0x6e6f6b40, TimeMid: 0x977c, TimeHiAndVersion: 0x4069, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xdd, Node: [6]uint8{0xac, 0x71, 0x0, 0x59, 0xf8, 0xc0}}
	MSVDSIVdsAdviseSink                               = uuid.UUID{TimeLow: 0x8326cd1d, TimeMid: 0xcf59, TimeHiAndVersion: 0x4936, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x86, Node: [6]uint8{0x5e, 0xfc, 0x8, 0x79, 0x8e, 0x25}}
	MSVDSIVdsAsync                                    = uuid.UUID{TimeLow: 0xd5d23b6d, TimeMid: 0x5a55, TimeHiAndVersion: 0x4492, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x89, Node: [6]uint8{0x39, 0x7a, 0x3c, 0x2d, 0x2d, 0xbc}}
	MSVDSIVdsCreatePartitionEx                        = uuid.UUID{TimeLow: 0x9882f547, TimeMid: 0xcfc3, TimeHiAndVersion: 0x420b, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x50, Node: [6]uint8{0x0, 0xdf, 0xbe, 0xc5, 0x6, 0x62}}
	MSVDSIVdsDisk2                                    = uuid.UUID{TimeLow: 0x40f73c8b, TimeMid: 0x687d, TimeHiAndVersion: 0x4a13, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x96, Node: [6]uint8{0x3d, 0x7f, 0x2e, 0x68, 0x39, 0x36}}
	MSVDSIVdsDisk3                                    = uuid.UUID{TimeLow: 0x8f4b2f5d, TimeMid: 0xec15, TimeHiAndVersion: 0x4357, ClockSeqHiAndReserved: 0x99, ClockSeqLow: 0x2f, Node: [6]uint8{0x47, 0x3e, 0xf1, 0x9, 0x75, 0xb9}}
	MSVDSIVdsDiskOnline                               = uuid.UUID{TimeLow: 0x90681b1d, TimeMid: 0x6a7f, TimeHiAndVersion: 0x48e8, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0x61, Node: [6]uint8{0x31, 0xb7, 0xaa, 0x12, 0x53, 0x22}}
	MSVDSIVdsDiskPartitionMF2                         = uuid.UUID{TimeLow: 0x9cbe50ca, TimeMid: 0xf2d2, TimeHiAndVersion: 0x4bf4, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xe1, Node: [6]uint8{0x96, 0x89, 0x6b, 0x72, 0x96, 0x25}}
	MSVDSIVdsDiskPartitionMF                          = uuid.UUID{TimeLow: 0x538684e0, TimeMid: 0xba3d, TimeHiAndVersion: 0x4bc0, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xa9, Node: [6]uint8{0x16, 0x4a, 0xff, 0x85, 0xc2, 0xa9}}
	MSVDSIVdsDisk                                     = uuid.UUID{TimeLow: 0x7e5c822, TimeMid: 0xf00c, TimeHiAndVersion: 0x47a1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xce, Node: [6]uint8{0xb2, 0x44, 0xda, 0x56, 0xfd, 0x6}}
	MSVDSIVdsHbaPort                                  = uuid.UUID{TimeLow: 0x2abd757f, TimeMid: 0x2851, TimeHiAndVersion: 0x4997, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x13, Node: [6]uint8{0x47, 0xd2, 0xa8, 0x85, 0xd6, 0xca}}
	MSVDSIVdsHwProvider                               = uuid.UUID{TimeLow: 0xd99bdaae, TimeMid: 0xb13a, TimeHiAndVersion: 0x4178, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xdb, Node: [6]uint8{0xe2, 0x7f, 0x16, 0xb4, 0x60, 0x3e}}
	MSVDSIVdsIscsiInitiatorAdapter                    = uuid.UUID{TimeLow: 0xb07fedd4, TimeMid: 0x1682, TimeHiAndVersion: 0x4440, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0x89, Node: [6]uint8{0xa3, 0x9b, 0x55, 0x19, 0x4d, 0xc5}}
	MSVDSIVdsIscsiInitiatorPortal                     = uuid.UUID{TimeLow: 0x38a0a9ab, TimeMid: 0x7cc8, TimeHiAndVersion: 0x4693, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0x7, Node: [6]uint8{0x1f, 0x28, 0xbd, 0x3, 0xc3, 0xda}}
	MSVDSIVdsOpenDisk                                 = uuid.UUID{TimeLow: 0x75c8f324, TimeMid: 0xf715, TimeHiAndVersion: 0x4fe3, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x8e, Node: [6]uint8{0xf9, 0x1, 0x1b, 0x61, 0xa4, 0xa1}}
	MSVDSIVdsPack2                                    = uuid.UUID{TimeLow: 0x13b50bff, TimeMid: 0x290a, TimeHiAndVersion: 0x47dd, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x58, Node: [6]uint8{0xb7, 0xc5, 0x8d, 0xb1, 0xa7, 0x1a}}
	MSVDSIVdsPack                                     = uuid.UUID{TimeLow: 0x3b69d7f5, TimeMid: 0x9d94, TimeHiAndVersion: 0x4648, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xca, Node: [6]uint8{0x79, 0x93, 0x9b, 0xa2, 0x63, 0xbf}}
	MSVDSIVdsProvider                                 = uuid.UUID{TimeLow: 0x10c5e575, TimeMid: 0x7984, TimeHiAndVersion: 0x4e81, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0x6b, Node: [6]uint8{0x43, 0x1f, 0x5f, 0x92, 0xae, 0x42}}
	MSVDSIVdsRemovable                                = uuid.UUID{TimeLow: 0x316560b, TimeMid: 0x5db4, TimeHiAndVersion: 0x4ed9, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xb5, Node: [6]uint8{0x21, 0x34, 0x36, 0xdd, 0xc0, 0xd9}}
	MSVDSIVdsServiceHba                               = uuid.UUID{TimeLow: 0xac13689, TimeMid: 0x3134, TimeHiAndVersion: 0x47c6, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x7c, Node: [6]uint8{0x46, 0x69, 0x21, 0x68, 0x1, 0xbe}}
	MSVDSIVdsServiceInitialization                    = uuid.UUID{TimeLow: 0x4afc3636, TimeMid: 0xdb01, TimeHiAndVersion: 0x4052, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xc3, Node: [6]uint8{0x3, 0xbb, 0xcb, 0x8d, 0x3c, 0x69}}
	MSVDSIVdsServiceIscsi                             = uuid.UUID{TimeLow: 0x14fbe036, TimeMid: 0x3ed7, TimeHiAndVersion: 0x4e10, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xe9, Node: [6]uint8{0xa5, 0xff, 0x99, 0x1a, 0xff, 0x1}}
	MSVDSIVdsServiceLoader                            = uuid.UUID{TimeLow: 0xe0393303, TimeMid: 0x90d4, TimeHiAndVersion: 0x4a97, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x71, Node: [6]uint8{0xe9, 0xb6, 0x71, 0xee, 0x27, 0x29}}
	MSVDSIVdsServiceSAN                               = uuid.UUID{TimeLow: 0xfc5d23e8, TimeMid: 0xa88b, TimeHiAndVersion: 0x41a5, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0xe0, Node: [6]uint8{0x2d, 0x2f, 0x73, 0xc5, 0xa6, 0x30}}
	MSVDSIVdsServiceSw                                = uuid.UUID{TimeLow: 0x15fc031c, TimeMid: 0x652, TimeHiAndVersion: 0x4306, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xc3, Node: [6]uint8{0xf5, 0x58, 0xb8, 0xf8, 0x37, 0xe2}}
	MSVDSIVdsServiceUninstallDisk                     = uuid.UUID{TimeLow: 0xb6b22da8, TimeMid: 0xf903, TimeHiAndVersion: 0x4be7, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x92, Node: [6]uint8{0xc0, 0x9d, 0x87, 0x5a, 0xc9, 0xda}}
	MSVDSIVdsService                                  = uuid.UUID{TimeLow: 0x818a8ef, TimeMid: 0x9ba9, TimeHiAndVersion: 0x40d8, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xf9, Node: [6]uint8{0xe2, 0x28, 0x33, 0xcc, 0x77, 0x1e}}
	MSVDSIVdsSubSystemImportTarget                    = uuid.UUID{TimeLow: 0x83bfb87f, TimeMid: 0x43fb, TimeHiAndVersion: 0x4903, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xa6, Node: [6]uint8{0x12, 0x7f, 0x1, 0x2, 0x9e, 0xec}}
	MSVDSIVdsSwProvider                               = uuid.UUID{TimeLow: 0x9aa58360, TimeMid: 0xce33, TimeHiAndVersion: 0x4f92, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x58, Node: [6]uint8{0xed, 0x24, 0xb1, 0x44, 0x25, 0xb8}}
	MSVDSIVdsVDisk                                    = uuid.UUID{TimeLow: 0x1e062b84, TimeMid: 0xe5e6, TimeHiAndVersion: 0x4b4b, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x25, Node: [6]uint8{0x67, 0xb8, 0x1e, 0x8f, 0x13, 0xe8}}
	MSVDSIVdsVdProvider                               = uuid.UUID{TimeLow: 0xb481498c, TimeMid: 0x8354, TimeHiAndVersion: 0x45f9, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xa0, Node: [6]uint8{0xb, 0xdd, 0x28, 0x32, 0xa9, 0x1f}}
	MSVDSIVdsVolume2                                  = uuid.UUID{TimeLow: 0x72ae6713, TimeMid: 0xdcbb, TimeHiAndVersion: 0x4a03, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x6b, Node: [6]uint8{0x37, 0x1f, 0x6a, 0xc6, 0xb5, 0x3d}}
	MSVDSIVdsVolumeMF2                                = uuid.UUID{TimeLow: 0x4dbcee9a, TimeMid: 0x6343, TimeHiAndVersion: 0x4651, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x5f, Node: [6]uint8{0x5e, 0x75, 0xd7, 0x4d, 0x98, 0x3c}}
	MSVDSIVdsVolumeMF3                                = uuid.UUID{TimeLow: 0x6788faf9, TimeMid: 0x214e, TimeHiAndVersion: 0x4b85, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x59, Node: [6]uint8{0x26, 0x69, 0x53, 0x61, 0x6e, 0x9}}
	MSVDSIVdsVolumeMF                                 = uuid.UUID{TimeLow: 0xee2d5ded, TimeMid: 0x6236, TimeHiAndVersion: 0x4169, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x1d, Node: [6]uint8{0xb9, 0x77, 0x8c, 0xe0, 0x3d, 0xc6}}
	MSVDSIVdsVolumeOnline                             = uuid.UUID{TimeLow: 0x1be2275a, TimeMid: 0xb315, TimeHiAndVersion: 0x4f70, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x44, Node: [6]uint8{0x87, 0x9b, 0x3a, 0x2a, 0x53, 0xf2}}
	MSVDSIVdsVolumePlex                               = uuid.UUID{TimeLow: 0x4daa0135, TimeMid: 0xe1d1, TimeHiAndVersion: 0x40f1, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xa5, Node: [6]uint8{0x3c, 0xc1, 0xe5, 0x32, 0x21, 0xc3}}
	MSVDSIVdsVolumeShrink                             = uuid.UUID{TimeLow: 0xd68168c9, TimeMid: 0x82a2, TimeHiAndVersion: 0x4f85, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xe9, Node: [6]uint8{0x74, 0x70, 0x7c, 0x49, 0xa5, 0x8f}}
	MSVDSIVdsVolume                                   = uuid.UUID{TimeLow: 0x88306bb2, TimeMid: 0xe71f, TimeHiAndVersion: 0x478c, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0xa2, Node: [6]uint8{0x79, 0xda, 0x20, 0xa, 0xf, 0x11}}
	MSW32TW32Time                                     = uuid.UUID{TimeLow: 0x8fb6d884, TimeMid: 0x2388, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x35, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xda, 0x27, 0x95}}
	MSWCCEICertRequestD2                              = uuid.UUID{TimeLow: 0x5422fd3a, TimeMid: 0xd4b8, TimeHiAndVersion: 0x4cef, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x2e, Node: [6]uint8{0xe8, 0x7d, 0x4c, 0xa2, 0x2e, 0x90}}
	MSWCCEICertRequestD                               = uuid.UUID{TimeLow: 0xd99e6e70, TimeMid: 0xfc88, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x98, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x3, 0x12, 0xf3}}
	MSWDSCWdsRpcInterface                             = uuid.UUID{TimeLow: 0x1a927394, TimeMid: 0x352e, TimeHiAndVersion: 0x4553, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x3f, Node: [6]uint8{0x7c, 0xf4, 0xaa, 0xfc, 0xa6, 0x20}}
	MSWKSTWkssvc                                      = uuid.UUID{TimeLow: 0x6bffd098, TimeMid: 0xa112, TimeHiAndVersion: 0x3610, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x33, Node: [6]uint8{0x46, 0xc3, 0xf8, 0x7e, 0x34, 0x5a}}
	MSWMIIEnumWbemClassObject                         = uuid.UUID{TimeLow: 0x27947e1, TimeMid: 0xd731, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x57, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x1}}
	MSWMIIWbemBackupRestoreEx                         = uuid.UUID{TimeLow: 0xa359dec5, TimeMid: 0xe813, TimeHiAndVersion: 0x4834, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x2a, Node: [6]uint8{0xba, 0x7f, 0x1d, 0x77, 0x7d, 0x76}}
	MSWMIIWbemBackupRestore                           = uuid.UUID{TimeLow: 0xc49e32c7, TimeMid: 0xbc8b, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd4, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x83, 0x4}}
	MSWMIIWbemClassObject                             = uuid.UUID{TimeLow: 0xdc12a681, TimeMid: 0x737f, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x4d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4b, 0x2e, 0x24}}
	MSWMIIWbemContext                                 = uuid.UUID{TimeLow: 0x44aca674, TimeMid: 0xe8fc, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x7c, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	MSWMIIWbemFetchSmartEnum                          = uuid.UUID{TimeLow: 0x1c1c45ee, TimeMid: 0x4395, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	MSWMIIWbemLevel1Login                             = uuid.UUID{TimeLow: 0xf309ad18, TimeMid: 0xd86a, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x75, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x88, 0x20}}
	MSWMIIWbemLoginClientID                           = uuid.UUID{TimeLow: 0xd4781cd6, TimeMid: 0xe5d3, TimeHiAndVersion: 0x44df, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x94, Node: [6]uint8{0x93, 0xe, 0xfe, 0x48, 0xa8, 0x87}}
	MSWMIIWbemLoginHelper                             = uuid.UUID{TimeLow: 0x541679ab, TimeMid: 0x2e5f, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x4e, Node: [6]uint8{0x0, 0x10, 0x4b, 0xcc, 0x4b, 0x4a}}
	MSWMIIWbemObjectSink                              = uuid.UUID{TimeLow: 0x7c857801, TimeMid: 0x7381, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x4d, Node: [6]uint8{0x0, 0xaa, 0x0, 0x4b, 0x2e, 0x24}}
	MSWMIIWbemRefreshingServices                      = uuid.UUID{TimeLow: 0x2c9273e0, TimeMid: 0x1dc3, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x64, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	MSWMIIWbemRemoteRefresher                         = uuid.UUID{TimeLow: 0xf1e9c5b2, TimeMid: 0xf59b, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x62, Node: [6]uint8{0x0, 0x10, 0x5a, 0x1f, 0x81, 0x77}}
	MSWMIIWbemServices                                = uuid.UUID{TimeLow: 0x9556dc99, TimeMid: 0x828c, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0x7e, Node: [6]uint8{0x0, 0xaa, 0x0, 0x32, 0x40, 0xc7}}
	MSWMIIWbemWCOSmartEnum                            = uuid.UUID{TimeLow: 0x423ec01e, TimeMid: 0x2e35, TimeHiAndVersion: 0x11d2, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x4, Node: [6]uint8{0x0, 0x10, 0x4b, 0x70, 0x3e, 0xfd}}
	MSWSRMIResourceManager2                           = uuid.UUID{TimeLow: 0x2a3eb639, TimeMid: 0xd134, TimeHiAndVersion: 0x422d, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xd8, Node: [6]uint8{0xaa, 0xa1, 0xb5, 0x21, 0x62, 0x2}}
	MSWSRMIResourceManager                            = uuid.UUID{TimeLow: 0xc5cebee2, TimeMid: 0x9df5, TimeHiAndVersion: 0x4cdd, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x8c, Node: [6]uint8{0xc2, 0x47, 0x1b, 0xc1, 0x44, 0xb4}}
	MSWSRMIWRMAccounting                              = uuid.UUID{TimeLow: 0x4f7ca01c, TimeMid: 0xa9e5, TimeHiAndVersion: 0x45b6, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x42, Node: [6]uint8{0x23, 0x32, 0xa1, 0x33, 0x9c, 0x1d}}
	MSWSRMIWRMCalendar                                = uuid.UUID{TimeLow: 0x481e06cf, TimeMid: 0xab04, TimeHiAndVersion: 0x4498, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xfe, Node: [6]uint8{0x12, 0x4a, 0xa, 0x34, 0x29, 0x6d}}
	MSWSRMIWRMConfig                                  = uuid.UUID{TimeLow: 0x21546ae8, TimeMid: 0x4da5, TimeHiAndVersion: 0x445e, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x7f, Node: [6]uint8{0x62, 0x7f, 0xea, 0x39, 0xc5, 0xe8}}
	MSWSRMIWRMMachineGroup                            = uuid.UUID{TimeLow: 0x943991a5, TimeMid: 0xb3fe, TimeHiAndVersion: 0x41fa, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x96, Node: [6]uint8{0x7f, 0x7b, 0x65, 0x6e, 0xe3, 0x4b}}
	MSWSRMIWRMPolicy                                  = uuid.UUID{TimeLow: 0x59602eb6, TimeMid: 0x57b0, TimeHiAndVersion: 0x4fd8, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x4b, Node: [6]uint8{0xeb, 0xf0, 0x69, 0x71, 0xfe, 0x15}}
	MSWSRMIWRMProtocol                                = uuid.UUID{TimeLow: 0xf31931a9, TimeMid: 0x832d, TimeHiAndVersion: 0x481c, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0x3, Node: [6]uint8{0x88, 0x7a, 0xe, 0x6a, 0x79, 0xf0}}
	MSWSRMIWRMRemoteSessionMgmt                       = uuid.UUID{TimeLow: 0xfc910418, TimeMid: 0x55ca, TimeHiAndVersion: 0x45ef, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0x64, Node: [6]uint8{0x83, 0xd4, 0xce, 0x7d, 0x30, 0xe0}}
	MSWSRMIWRMResourceGroup                           = uuid.UUID{TimeLow: 0xbc681469, TimeMid: 0x9dd9, TimeHiAndVersion: 0x4bf4, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x3d, Node: [6]uint8{0x70, 0x9f, 0x69, 0xef, 0xe4, 0x31}}
	MSDLLAdvapi32                                     = uuid.UUID{TimeLow: 0x98fe2c90, TimeMid: 0xa542, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xef, Node: [6]uint8{0x0, 0xa0, 0xc9, 0x6, 0x29, 0x10}}
	MSDLLAppidsvc                                     = uuid.UUID{TimeLow: 0x8a7b5006, TimeMid: 0xcc13, TimeHiAndVersion: 0x11db, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0x5, Node: [6]uint8{0x0, 0x50, 0x56, 0xc0, 0x0, 0x8}}
	MSDLLAppinfo                                      = uuid.UUID{TimeLow: 0x201ef99a, TimeMid: 0x7fa0, TimeHiAndVersion: 0x444c, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x99, Node: [6]uint8{0x19, 0xba, 0x84, 0xf1, 0x2a, 0x1a}}
	MSDLLIAisAxISInfo                                 = uuid.UUID{TimeLow: 0x58e604e8, TimeMid: 0x9adb, TimeHiAndVersion: 0x4d2e, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0x64, Node: [6]uint8{0x3b, 0x6, 0x83, 0xfb, 0x14, 0x80}}
	MSDLLIAisCOMInfo                                  = uuid.UUID{TimeLow: 0x5f54ce7d, TimeMid: 0x5b79, TimeHiAndVersion: 0x4175, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x84, Node: [6]uint8{0xcb, 0x65, 0x31, 0x3a, 0xe, 0x98}}
	MSDLLIAisMSIInfo                                  = uuid.UUID{TimeLow: 0xfd7a0523, TimeMid: 0xdc70, TimeHiAndVersion: 0x43dd, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x2e, Node: [6]uint8{0x9c, 0x5e, 0xd4, 0x82, 0x25, 0xb1}}
	MSDLLAppmgmts                                     = uuid.UUID{TimeLow: 0x8c7daf44, TimeMid: 0xb6dc, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x4c, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6e, 0x7c, 0x57}}
	MSDLLAqueue                                       = uuid.UUID{TimeLow: 0xbfa951d1, TimeMid: 0x2f0e, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xd1, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x49, 0xa}}
	MSDLLAudiodgrpc                                   = uuid.UUID{TimeLow: 0xde3b9bc8, TimeMid: 0xbef7, TimeHiAndVersion: 0x4578, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0xde, Node: [6]uint8{0xf0, 0x89, 0x4, 0x84, 0x42, 0xdb}}
	MSDLLAudiosrv                                     = uuid.UUID{TimeLow: 0x3faf4738, TimeMid: 0x3a21, TimeHiAndVersion: 0x4307, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x6c, Node: [6]uint8{0xfd, 0xda, 0x9b, 0xb8, 0xc0, 0xd5}}
	MSDLLAudiorpc                                     = uuid.UUID{TimeLow: 0xc386ca3e, TimeMid: 0x9061, TimeHiAndVersion: 0x4a72, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x1e, Node: [6]uint8{0x49, 0x8d, 0x83, 0xbe, 0x18, 0x8f}}
	MSDLLWluirRpc                                     = uuid.UUID{TimeLow: 0x24019106, TimeMid: 0xa203, TimeHiAndVersion: 0x4642, ClockSeqHiAndReserved: 0xb8, ClockSeqLow: 0x8d, Node: [6]uint8{0x82, 0xda, 0xe9, 0x15, 0x89, 0x29}}
	MSDLLWluir1Rpc                                    = uuid.UUID{TimeLow: 0xb2507c30, TimeMid: 0xb126, TimeHiAndVersion: 0x494a, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xac, Node: [6]uint8{0xee, 0x32, 0xb6, 0xee, 0xb0, 0x39}}
	MSDLLAuthmgr32Rpc                                 = uuid.UUID{TimeLow: 0xfa4febc0, TimeMid: 0x4591, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0x95, ClockSeqLow: 0xe5, Node: [6]uint8{0x0, 0xaa, 0x0, 0x51, 0xe5, 0x10}}
	MSDLLBdesvc                                       = uuid.UUID{TimeLow: 0xae55c4c0, TimeMid: 0x64ce, TimeHiAndVersion: 0x11dd, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x8b, Node: [6]uint8{0x8, 0x0, 0x20, 0xc, 0x9a, 0x66}}
	MSDLLBferpc                                       = uuid.UUID{TimeLow: 0xdd490425, TimeMid: 0x5325, TimeHiAndVersion: 0x4565, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x74, Node: [6]uint8{0x7e, 0x27, 0xd6, 0xc0, 0x9c, 0x24}}
	MSDLLBthServRPCService                            = uuid.UUID{TimeLow: 0x2acb9d68, TimeMid: 0xb434, TimeHiAndVersion: 0x4b3e, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x66, Node: [6]uint8{0xe0, 0x6b, 0x4b, 0x3a, 0x84, 0xcb}}
	MSDLLCertPropSvc                                  = uuid.UUID{TimeLow: 0x30b044a5, TimeMid: 0xa225, TimeHiAndVersion: 0x43f0, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0xa4, Node: [6]uint8{0xe0, 0x60, 0xdf, 0x91, 0xf9, 0xc1}}
	MSDLLClusSvcJoinVersion                           = uuid.UUID{TimeLow: 0x6e17aaa0, TimeMid: 0x1a47, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0xbd, Node: [6]uint8{0x0, 0x0, 0xf8, 0x75, 0x29, 0x2e}}
	MSDLLClusSvc                                      = uuid.UUID{TimeLow: 0xe248d0b8, TimeMid: 0xbf15, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x5e, Node: [6]uint8{0x8, 0x0, 0x2b, 0xb4, 0x96, 0x49}}
	MSDLLClusSvcExtrocluster                          = uuid.UUID{TimeLow: 0xffe561b8, TimeMid: 0xbf15, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x5e, Node: [6]uint8{0x8, 0x0, 0x2b, 0xb4, 0x96, 0x49}}
	MSDLLCryptCertProtectSvc                          = uuid.UUID{TimeLow: 0xd72a7d4, TimeMid: 0x6148, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xaa, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xb6, 0x6e, 0xa0}}
	MSDLLCryptKSrSvc                                  = uuid.UUID{TimeLow: 0x68b58241, TimeMid: 0xc259, TimeHiAndVersion: 0x4f03, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xe5, Node: [6]uint8{0xa2, 0x65, 0x1d, 0xcb, 0xc9, 0x30}}
	MSDLLCryptKeyrSvc                                 = uuid.UUID{TimeLow: 0x8d0ffe72, TimeMid: 0xd252, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x8f, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xd9, 0x12, 0x6b}}
	MSDLLCryptCatSvc                                  = uuid.UUID{TimeLow: 0xf50aac00, TimeMid: 0xc7f3, TimeHiAndVersion: 0x428e, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x22, Node: [6]uint8{0xa6, 0xb7, 0x1b, 0xfb, 0x9d, 0x43}}
	MSDLLDhcpcsvc6                                    = uuid.UUID{TimeLow: 0x3c4728c5, TimeMid: 0xf0ab, TimeHiAndVersion: 0x448b, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xa1, Node: [6]uint8{0x6c, 0xe0, 0x1e, 0xb0, 0xa6, 0xd6}}
	MSDLLDhcpcsvc                                     = uuid.UUID{TimeLow: 0x3c4728c5, TimeMid: 0xf0ab, TimeHiAndVersion: 0x448b, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xa1, Node: [6]uint8{0x6c, 0xe0, 0x1e, 0xb0, 0xa6, 0xd5}}
	MSDLLDmadminrpc                                   = uuid.UUID{TimeLow: 0xd2d79dfa, TimeMid: 0x3400, TimeHiAndVersion: 0x11d0, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xb, Node: [6]uint8{0x0, 0xaa, 0x0, 0x5f, 0xf5, 0x86}}
	MSDLLDnsResolver                                  = uuid.UUID{TimeLow: 0x45776b01, TimeMid: 0x5956, TimeHiAndVersion: 0x4485, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0x80, Node: [6]uint8{0xf4, 0x28, 0xf7, 0xd6, 0x1, 0x29}}
	MSDLLDnsResolver2000                              = uuid.UUID{TimeLow: 0x65a93890, TimeMid: 0xfab9, TimeHiAndVersion: 0x43a3, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xa5, Node: [6]uint8{0x1e, 0x33, 0xa, 0xc2, 0x8f, 0x11}}
	MSDLLWinlan                                       = uuid.UUID{TimeLow: 0x1bddb2a6, TimeMid: 0xc0c3, TimeHiAndVersion: 0x41be, ClockSeqHiAndReserved: 0x87, ClockSeqLow: 0x3, Node: [6]uint8{0xdd, 0xbd, 0xf4, 0xf0, 0xe8, 0xa}}
	MSDLLEfskrpc                                      = uuid.UUID{TimeLow: 0x4eeb297, TimeMid: 0xcbf4, TimeHiAndVersion: 0x466b, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x2a, Node: [6]uint8{0xbf, 0xd6, 0xa2, 0xf1, 0xb, 0xba}}
	MSDLLEfsrpc                                       = uuid.UUID{TimeLow: 0xdf1941c5, TimeMid: 0xfe89, TimeHiAndVersion: 0x4e79, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x10, Node: [6]uint8{0x46, 0x36, 0x57, 0xac, 0xf4, 0x4d}}
	MSDLLEmdService                                   = uuid.UUID{TimeLow: 0x654976df, TimeMid: 0x1498, TimeHiAndVersion: 0x4056, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x5e, Node: [6]uint8{0xcb, 0x4e, 0x87, 0x58, 0x4b, 0xd8}}
	MSDLLGpsvc                                        = uuid.UUID{TimeLow: 0x2eb08e3e, TimeMid: 0x639f, TimeHiAndVersion: 0x4fba, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xb1, Node: [6]uint8{0x14, 0xf8, 0x78, 0x96, 0x10, 0x76}}
	MSDLLIasrpc                                       = uuid.UUID{TimeLow: 0x7d814569, TimeMid: 0x35b3, TimeHiAndVersion: 0x4850, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x32, Node: [6]uint8{0x83, 0x3, 0x5f, 0xce, 0xbf, 0x6e}}
	MSDLLIcardagtrpc                                  = uuid.UUID{TimeLow: 0x645e6c, TimeMid: 0xfc9f, TimeHiAndVersion: 0x4a0c, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x96, Node: [6]uint8{0xf0, 0xb, 0x66, 0x29, 0x77, 0x98}}
	MSDLLIertutil                                     = uuid.UUID{TimeLow: 0xf1ec59ab, TimeMid: 0x4ca9, TimeHiAndVersion: 0x4c30, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xd0, Node: [6]uint8{0x54, 0xef, 0x1d, 0xb4, 0x41, 0xb7}}
	MSDLLIdletask                                     = uuid.UUID{TimeLow: 0xa74ef1c, TimeMid: 0x41a4, TimeHiAndVersion: 0x4e06, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0xae, Node: [6]uint8{0xdc, 0x74, 0xfb, 0x1c, 0xdd, 0x53}}
	MSDLLISecureDesktop                               = uuid.UUID{TimeLow: 0x12e65dd8, TimeMid: 0x887f, TimeHiAndVersion: 0x41ef, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xbf, Node: [6]uint8{0x8d, 0x81, 0x6c, 0x42, 0xc2, 0xe7}}
	MSDLLWMsgKAPIs                                    = uuid.UUID{TimeLow: 0x76f226c3, TimeMid: 0xec14, TimeHiAndVersion: 0x4325, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x99, Node: [6]uint8{0x6a, 0x46, 0x34, 0x84, 0x18, 0xae}}
	MSDLLWMsgAPIs                                     = uuid.UUID{TimeLow: 0x76f226c3, TimeMid: 0xec14, TimeHiAndVersion: 0x4325, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x99, Node: [6]uint8{0x6a, 0x46, 0x34, 0x84, 0x18, 0xaf}}
	MSDLLIRPCSCLogon                                  = uuid.UUID{TimeLow: 0x95958c94, TimeMid: 0xa424, TimeHiAndVersion: 0x4055, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x2b, Node: [6]uint8{0xb7, 0xf4, 0xd5, 0xc4, 0x77, 0x70}}
	MSDLLNrpsrv                                       = uuid.UUID{TimeLow: 0x30adc50c, TimeMid: 0x5cbc, TimeHiAndVersion: 0x46ce, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0xe, Node: [6]uint8{0x91, 0x91, 0x47, 0x89, 0xe2, 0x3c}}
	MSDLLWinhttprpc                                   = uuid.UUID{TimeLow: 0x3473dd4d, TimeMid: 0x2e88, TimeHiAndVersion: 0x4006, ClockSeqHiAndReserved: 0x9c, ClockSeqLow: 0xba, Node: [6]uint8{0x22, 0x57, 0x9, 0x9, 0xdd, 0x10}}
	MSDLLIRPCAsyncNotifyChannel                       = uuid.UUID{TimeLow: 0x4a452661, TimeMid: 0x8290, TimeHiAndVersion: 0x4b36, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xbe, Node: [6]uint8{0x7f, 0x40, 0x93, 0xa9, 0x49, 0x78}}
	MSDLLIpTransition                                 = uuid.UUID{TimeLow: 0x552d076a, TimeMid: 0xcb29, TimeHiAndVersion: 0x4e44, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x6a, Node: [6]uint8{0xd1, 0x5e, 0x59, 0xe2, 0xc0, 0xaf}}
	MSDLLWinNsi                                       = uuid.UUID{TimeLow: 0x7ea70bcf, TimeMid: 0x48af, TimeHiAndVersion: 0x4f6a, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x68, Node: [6]uint8{0x6a, 0x44, 0x7, 0x54, 0xd5, 0xfa}}
	MSDLLXactSrvRPC                                   = uuid.UUID{TimeLow: 0x98716d03, TimeMid: 0x89ac, TimeHiAndVersion: 0x44c7, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x8c, Node: [6]uint8{0x28, 0x58, 0x24, 0xe5, 0x1c, 0x4a}}
	MSDLLIkeRpcIKE                                    = uuid.UUID{TimeLow: 0xa398e520, TimeMid: 0xd59a, TimeHiAndVersion: 0x4bdd, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x7a, Node: [6]uint8{0x3c, 0x1e, 0x3, 0x3, 0xa5, 0x11}}
	MSDLLIRPCWinlogonNotifications                    = uuid.UUID{TimeLow: 0xc9ac6db5, TimeMid: 0x82b7, TimeHiAndVersion: 0x4e55, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x8a, Node: [6]uint8{0xe4, 0x64, 0xed, 0x7b, 0x42, 0x77}}
	MSDLLIfssvc                                       = uuid.UUID{TimeLow: 0x46ea9280, TimeMid: 0x5bbf, TimeHiAndVersion: 0x445e, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x1d, Node: [6]uint8{0x41, 0xd0, 0xf6, 0xf, 0x50, 0x3a}}
	MSDLLImepadsm                                     = uuid.UUID{TimeLow: 0x52d9f704, TimeMid: 0xd3c6, TimeHiAndVersion: 0x4748, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x11, Node: [6]uint8{0x25, 0x50, 0x20, 0x9e, 0x80, 0xaf}}
	MSDLLImepadsm1                                    = uuid.UUID{TimeLow: 0x8f1acdc1, TimeMid: 0x754d, TimeHiAndVersion: 0x43eb, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0x29, Node: [6]uint8{0xaa, 0x16, 0x20, 0x92, 0x8e, 0x65}}
	MSDLLImjpdctrpc                                   = uuid.UUID{TimeLow: 0xcb407bbf, TimeMid: 0xc14f, TimeHiAndVersion: 0x4cd9, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x55, Node: [6]uint8{0xcb, 0xb0, 0x81, 0x46, 0x59, 0x8c}}
	MSDLLCpprovider                                   = uuid.UUID{TimeLow: 0x6f201a55, TimeMid: 0xa24d, TimeHiAndVersion: 0x495f, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc9, Node: [6]uint8{0x2f, 0x4f, 0xce, 0x34, 0xdf, 0x99}}
	MSDLLIphlpsvc                                     = uuid.UUID{TimeLow: 0x6f201a55, TimeMid: 0xa24d, TimeHiAndVersion: 0x495f, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xc9, Node: [6]uint8{0x2f, 0x4f, 0xce, 0x34, 0xdf, 0x98}}
	MSDLLIpnathlp                                     = uuid.UUID{TimeLow: 0xd674a233, TimeMid: 0x5829, TimeHiAndVersion: 0x49dd, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xf0, Node: [6]uint8{0x60, 0xcf, 0x9c, 0xeb, 0x71, 0x29}}
	MSDLLIpNatHlpRpcServer                            = uuid.UUID{TimeLow: 0xe64b9aee, TimeMid: 0xf372, TimeHiAndVersion: 0x4312, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x14, Node: [6]uint8{0x8f, 0x15, 0x2, 0xb5, 0xc8, 0xe3}}
	MSDLLIrftprpc                                     = uuid.UUID{TimeLow: 0xc821d64, TimeMid: 0xa3fc, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0x7a, Node: [6]uint8{0x0, 0x80, 0xc7, 0x5e, 0x4e, 0xc1}}
	MSDLLIrmon                                        = uuid.UUID{TimeLow: 0x7af5bbd0, TimeMid: 0x6063, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x2a, Node: [6]uint8{0x0, 0x80, 0xc7, 0x5e, 0x4e, 0xc1}}
	MSDLLIrmon1                                       = uuid.UUID{TimeLow: 0x209bb240, TimeMid: 0xb919, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xbb, ClockSeqLow: 0xb6, Node: [6]uint8{0x0, 0x80, 0xc7, 0x5e, 0x4e, 0xc1}}
	MSDLLIscsiexe                                     = uuid.UUID{TimeLow: 0xfd6bb951, TimeMid: 0xc830, TimeHiAndVersion: 0x4734, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x2c, Node: [6]uint8{0x18, 0xba, 0x6e, 0xc7, 0xab, 0x49}}
	MSDLLIsmserv_ip                                   = uuid.UUID{TimeLow: 0x130ceefb, TimeMid: 0xe466, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x8b, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x28, 0x83}}
	MSDLLIsmapi                                       = uuid.UUID{TimeLow: 0x68dcd486, TimeMid: 0x669e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xc, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc2, 0xdc, 0xd2}}
	MSDLLKdcrpc                                       = uuid.UUID{TimeLow: 0x2, TimeMid: 0x1, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x69}}
	MSDLLINCryptKeyIso                                = uuid.UUID{TimeLow: 0xb25a52bf, TimeMid: 0xe5dd, TimeHiAndVersion: 0x4f4a, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0xa6, Node: [6]uint8{0x8c, 0xa7, 0x27, 0x2a, 0xe, 0x86}}
	MSDLLLbservice                                    = uuid.UUID{TimeLow: 0x3357951c, TimeMid: 0xa1d1, TimeHiAndVersion: 0x47db, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x78, Node: [6]uint8{0xab, 0x94, 0x5d, 0x6, 0x3d, 0x3}}
	MSDLLLlsrpc                                       = uuid.UUID{TimeLow: 0x342cfd40, TimeMid: 0x3c6c, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0x93, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2e, 0x9c, 0x6d}}
	MSDLLLls_license                                  = uuid.UUID{TimeLow: 0x57674cd0, TimeMid: 0x5200, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0x97, Node: [6]uint8{0x8, 0x0, 0x2b, 0x2e, 0x9c, 0x6d}}
	MSDLLNsiC                                         = uuid.UUID{TimeLow: 0xd3fbb514, TimeMid: 0xe3b, TimeHiAndVersion: 0x11cb, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xad, Node: [6]uint8{0x8, 0x0, 0x2b, 0x1d, 0x29, 0xc3}}
	MSDLLNsiS                                         = uuid.UUID{TimeLow: 0xd6d70ef0, TimeMid: 0xe3b, TimeHiAndVersion: 0x11cb, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xc3, Node: [6]uint8{0x8, 0x0, 0x2b, 0x1d, 0x29, 0xc3}}
	MSDLLNsiM                                         = uuid.UUID{TimeLow: 0xd6d70ef0, TimeMid: 0xe3b, TimeHiAndVersion: 0x11cb, ClockSeqHiAndReserved: 0xac, ClockSeqLow: 0xc3, Node: [6]uint8{0x8, 0x0, 0x2b, 0x1d, 0x29, 0xc4}}
	MSDLLDsrolesvc                                    = uuid.UUID{TimeLow: 0x1cbcad78, TimeMid: 0xdf0b, TimeHiAndVersion: 0x4934, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x58, Node: [6]uint8{0x87, 0x83, 0x9e, 0xa5, 0x1, 0xc9}}
	MSDLLICryptProtect                                = uuid.UUID{TimeLow: 0x11220835, TimeMid: 0x5b26, TimeHiAndVersion: 0x4d94, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x86, Node: [6]uint8{0xc3, 0xe4, 0x75, 0xa8, 0x9, 0xde}}
	MSDLLPasswordRecovery                             = uuid.UUID{TimeLow: 0x5cbe92cb, TimeMid: 0xf4be, TimeHiAndVersion: 0x45c9, ClockSeqHiAndReserved: 0x9f, ClockSeqLow: 0xc9, Node: [6]uint8{0x33, 0xe7, 0x3e, 0x55, 0x7b, 0x20}}
	MSDLLSLspPrivateData                              = uuid.UUID{TimeLow: 0xace1c026, TimeMid: 0x8b3f, TimeHiAndVersion: 0x4711, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x18, Node: [6]uint8{0xf3, 0x45, 0xd1, 0x7f, 0x5b, 0xff}}
	MSDLLLsass                                        = uuid.UUID{TimeLow: 0xfb8a0729, TimeMid: 0x2d04, TimeHiAndVersion: 0x4658, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x93, Node: [6]uint8{0x27, 0xb4, 0xad, 0x55, 0x3f, 0xac}}
	MSDLLTermServLicensing                            = uuid.UUID{TimeLow: 0x12d4b7c8, TimeMid: 0x77d5, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x24, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x8, 0xd}}
	MSDLLHydraLsPipe                                  = uuid.UUID{TimeLow: 0x3d267954, TimeMid: 0xeeb7, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x4e, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xa3, 0x8, 0xd}}
	MSDLLTermSrvPrivate                               = uuid.UUID{TimeLow: 0x11f25515, TimeMid: 0xc879, TimeHiAndVersion: 0x400a, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x9e, Node: [6]uint8{0xb0, 0x74, 0xd5, 0xf0, 0x92, 0xfe}}
	MSDLLTermSrvAdmin                                 = uuid.UUID{TimeLow: 0x1e665584, TimeMid: 0x40fe, TimeHiAndVersion: 0x4450, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x6e, Node: [6]uint8{0x80, 0x23, 0x62, 0x39, 0x96, 0x94}}
	MSDLLMilcore                                      = uuid.UUID{TimeLow: 0x506c3b0e, TimeMid: 0x4bd1, TimeHiAndVersion: 0x4c56, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xc0, Node: [6]uint8{0x49, 0xa2, 0xe, 0xd4, 0xb5, 0x39}}
	MSDLLMpnotifyrpc                                  = uuid.UUID{TimeLow: 0x3ca78105, TimeMid: 0xa3a3, TimeHiAndVersion: 0x4a68, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x58, Node: [6]uint8{0x1a, 0x60, 0x6b, 0xab, 0x8f, 0xd6}}
	MSDLLACPBackgroundManagerPolicy                   = uuid.UUID{TimeLow: 0xfc48cd89, TimeMid: 0x98d6, TimeHiAndVersion: 0x4628, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x39, Node: [6]uint8{0x86, 0xf7, 0xa3, 0xe4, 0x16, 0x1a}}
	MSDLLAdhsvc                                       = uuid.UUID{TimeLow: 0xc49a5a70, TimeMid: 0x8a7f, TimeHiAndVersion: 0x4e70, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x16, Node: [6]uint8{0x1e, 0x8f, 0x1f, 0x19, 0x3e, 0xf1}}
	MSDLLAphostservice                                = uuid.UUID{TimeLow: 0xd2716e94, TimeMid: 0x25cb, TimeHiAndVersion: 0x4820, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x15, Node: [6]uint8{0x53, 0x78, 0x66, 0x57, 0x85, 0x62}}
	MSDLLAppinfopkg                                   = uuid.UUID{TimeLow: 0x497b57d, TimeMid: 0x2e66, TimeHiAndVersion: 0x424f, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0xc6, Node: [6]uint8{0x15, 0x7c, 0xd5, 0xd4, 0x17, 0x0}}
	MSDLLAudiosrvpbm                                  = uuid.UUID{TimeLow: 0x5fc2481b, TimeMid: 0xf8d7, TimeHiAndVersion: 0x466b, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x41, Node: [6]uint8{0xcc, 0x78, 0x6, 0xc7, 0x84, 0xa3}}
	MSDLLAudiosrvpbm1                                 = uuid.UUID{TimeLow: 0x910562c3, TimeMid: 0xebd9, TimeHiAndVersion: 0x46b9, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0xba, Node: [6]uint8{0x1d, 0x45, 0x84, 0x2a, 0xc, 0xeb}}
	MSDLLAudiosrvhrtf                                 = uuid.UUID{TimeLow: 0xc7ce3826, TimeMid: 0x891f, TimeHiAndVersion: 0x4376, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x61, Node: [6]uint8{0xc6, 0x3d, 0x24, 0x3, 0x14, 0x2c}}
	MSDLLAudiosrvwnf                                  = uuid.UUID{TimeLow: 0xcba4c918, TimeMid: 0xe55a, TimeHiAndVersion: 0x46ee, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0x62, Node: [6]uint8{0xca, 0xde, 0x15, 0x8e, 0x91, 0x65}}
	MSDLLAudiosrv1                                    = uuid.UUID{TimeLow: 0x7c69ac10, TimeMid: 0xfa12, TimeHiAndVersion: 0x4dbf, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xd9, Node: [6]uint8{0xc7, 0xf1, 0xe4, 0xf, 0x5d, 0xc5}}
	MSDLLAudiosrv2                                    = uuid.UUID{TimeLow: 0x47ac638a, TimeMid: 0x718f, TimeHiAndVersion: 0x49a0, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xc5, Node: [6]uint8{0x57, 0x4a, 0xc7, 0x7a, 0xcf, 0x4d}}
	MSDLLAudiosrv3                                    = uuid.UUID{TimeLow: 0x9704557, TimeMid: 0x82c0, TimeHiAndVersion: 0x416b, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xe4, Node: [6]uint8{0xc8, 0x5b, 0x8f, 0x78, 0x98, 0x3}}
	MSDLLRBiPtSrv                                     = uuid.UUID{TimeLow: 0xd3e2735, TimeMid: 0xcea0, TimeHiAndVersion: 0x4ecc, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0xe2, Node: [6]uint8{0x41, 0xa2, 0xd8, 0x1a, 0xed, 0x4e}}
	MSDLLRBiRtSrv                                     = uuid.UUID{TimeLow: 0x1b37ca91, TimeMid: 0x76b1, TimeHiAndVersion: 0x4f5e, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xc7, Node: [6]uint8{0x2a, 0xbf, 0xc6, 0x1f, 0x2b, 0xb0}}
	MSDLLSrvOdbPt                                     = uuid.UUID{TimeLow: 0x20c40295, TimeMid: 0x8dba, TimeHiAndVersion: 0x48e6, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0xbf, Node: [6]uint8{0x3e, 0x78, 0xef, 0x3b, 0xb1, 0x44}}
	MSDLLSrvOdb                                       = uuid.UUID{TimeLow: 0x2513bcbe, TimeMid: 0x6cd4, TimeHiAndVersion: 0x4348, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0x5e, Node: [6]uint8{0x7e, 0xfb, 0x3c, 0x33, 0x6d, 0xd3}}
	MSDLLRBiSrv                                       = uuid.UUID{TimeLow: 0x2d98a740, TimeMid: 0x581d, TimeHiAndVersion: 0x41b9, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xd, Node: [6]uint8{0xa8, 0x8b, 0x9d, 0x5c, 0xe9, 0x38}}
	MSDLLSrvOdbPriv                                   = uuid.UUID{TimeLow: 0x857fb1be, TimeMid: 0x84f, TimeHiAndVersion: 0x4fb5, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x9c, Node: [6]uint8{0x4b, 0x2c, 0x4b, 0xe5, 0xf0, 0xcf}}
	MSDLLRBiSrv1                                      = uuid.UUID{TimeLow: 0x8bfc3be1, TimeMid: 0x6def, TimeHiAndVersion: 0x4e2d, ClockSeqHiAndReserved: 0xaf, ClockSeqLow: 0x74, Node: [6]uint8{0x7c, 0x47, 0xcd, 0xa, 0xde, 0x4a}}
	MSDLLSrvOdbLb                                     = uuid.UUID{TimeLow: 0xb8cadbaf, TimeMid: 0xe84b, TimeHiAndVersion: 0x46b9, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xf2, Node: [6]uint8{0x6f, 0x71, 0xc0, 0x3f, 0x9e, 0x55}}
	MSDLLRBiSrvSignal                                 = uuid.UUID{TimeLow: 0xc605f9fb, TimeMid: 0xf0a3, TimeHiAndVersion: 0x4e2a, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x73, Node: [6]uint8{0x73, 0x56, 0xf, 0x8d, 0x9e, 0x3e}}
	MSDLLBri                                          = uuid.UUID{TimeLow: 0xd09bdeb5, TimeMid: 0x6171, TimeHiAndVersion: 0x4a34, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xe2, Node: [6]uint8{0x6, 0xfa, 0x82, 0x65, 0x25, 0x68}}
	MSDLLIChakraJIT                                   = uuid.UUID{TimeLow: 0xead694ed, TimeMid: 0x2243, TimeHiAndVersion: 0x44cb, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0xdc, Node: [6]uint8{0x85, 0xd3, 0xba, 0x93, 0x4d, 0xab}}
	MSDLLCoremessaging                                = uuid.UUID{TimeLow: 0xdf4df73a, TimeMid: 0xc52d, TimeHiAndVersion: 0x4e3a, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x3, Node: [6]uint8{0x84, 0x37, 0xfd, 0xf8, 0x30, 0x2a}}
	MSDLLCrypttpmeksvc                                = uuid.UUID{TimeLow: 0x2579ff35, TimeMid: 0xab0, TimeHiAndVersion: 0x4e5a, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xfa, Node: [6]uint8{0x1d, 0x88, 0xc4, 0xe0, 0xcb, 0x92}}
	MSDLLWarpJITSvc                                   = uuid.UUID{TimeLow: 0x5a0ce74d, TimeMid: 0xf9cf, TimeHiAndVersion: 0x4dea, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xc1, Node: [6]uint8{0x2d, 0x5f, 0xe4, 0xc8, 0x9d, 0x51}}
	MSDLLDab                                          = uuid.UUID{TimeLow: 0x7419cf08, TimeMid: 0x91a7, TimeHiAndVersion: 0x4afd, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x5e, Node: [6]uint8{0x1d, 0xd7, 0x6d, 0xe0, 0x94, 0xfd}}
	MSDLLDaschallenge                                 = uuid.UUID{TimeLow: 0x2e7d4935, TimeMid: 0x59d2, TimeHiAndVersion: 0x4312, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xc8, Node: [6]uint8{0x41, 0x90, 0xa, 0xa5, 0x49, 0x5f}}
	MSDLLDasaepstore                                  = uuid.UUID{TimeLow: 0x5b665b9a, TimeMid: 0xa086, TimeHiAndVersion: 0x4e26, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x24, Node: [6]uint8{0x96, 0xab, 0x5, 0xb, 0xe, 0xc3}}
	MSDLLDas                                          = uuid.UUID{TimeLow: 0x850cee52, TimeMid: 0x3038, TimeHiAndVersion: 0x4277, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0xb4, Node: [6]uint8{0xe0, 0x5d, 0xb8, 0xb2, 0xc3, 0x5c}}
	MSDLLDasinbassoc                                  = uuid.UUID{TimeLow: 0xa1d4eae7, TimeMid: 0x39f8, TimeHiAndVersion: 0x4bca, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x72, Node: [6]uint8{0x83, 0x27, 0x67, 0xf5, 0x8, 0x2a}}
	MSDLLDasrpcdev                                    = uuid.UUID{TimeLow: 0xbd84cd86, TimeMid: 0x9825, TimeHiAndVersion: 0x4376, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x3d, Node: [6]uint8{0x33, 0x4c, 0x54, 0x3f, 0x89, 0xb1}}
	MSDLLDashost                                      = uuid.UUID{TimeLow: 0xeeee008d, TimeMid: 0x5c99, TimeHiAndVersion: 0x4e4b, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x1b, Node: [6]uint8{0x54, 0x7a, 0x26, 0xe8, 0xab, 0xd0}}
	MSDLLUtcApi                                       = uuid.UUID{TimeLow: 0x4c9dbf19, TimeMid: 0xd39e, TimeHiAndVersion: 0x4bb9, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xee, Node: [6]uint8{0x8f, 0x71, 0x79, 0xb2, 0x2, 0x83}}
	MSDLLUtcTelemetryOptInApi                         = uuid.UUID{TimeLow: 0x95095ec8, TimeMid: 0x32ea, TimeHiAndVersion: 0x4eb0, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xe2, Node: [6]uint8{0x4, 0x1f, 0x97, 0xb3, 0x61, 0x68}}
	MSDLLUtcWerHelperApi                              = uuid.UUID{TimeLow: 0x98cd761e, TimeMid: 0xe77d, TimeHiAndVersion: 0x41c8, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xc0, Node: [6]uint8{0xf, 0xb7, 0x56, 0xd9, 0xe, 0xc2}}
	MSDLLUtcEventTranscriptApi                        = uuid.UUID{TimeLow: 0xd22895ef, TimeMid: 0xaff4, TimeHiAndVersion: 0x42c5, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0xb2, Node: [6]uint8{0xb1, 0x44, 0x66, 0xd3, 0x4a, 0xb4}}
	MSDLLUtcTenantApi                                 = uuid.UUID{TimeLow: 0xe38f5360, TimeMid: 0x8572, TimeHiAndVersion: 0x473e, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0x96, Node: [6]uint8{0x1b, 0x46, 0x87, 0x3b, 0xee, 0xab}}
	MSDLLUtcApi1                                      = uuid.UUID{TimeLow: 0xfd8be72b, TimeMid: 0xa9cd, TimeHiAndVersion: 0x4b2c, ClockSeqHiAndReserved: 0xa9, ClockSeqLow: 0xca, Node: [6]uint8{0x4d, 0xed, 0x24, 0x2f, 0xbe, 0x4d}}
	MSDLLDpapisrv                                     = uuid.UUID{TimeLow: 0x7f1317a8, TimeMid: 0x4dea, TimeHiAndVersion: 0x4fa2, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0x51, Node: [6]uint8{0xdf, 0x55, 0x16, 0xff, 0x88, 0x79}}
	MSDLLDssvc                                        = uuid.UUID{TimeLow: 0xbf4dc912, TimeMid: 0xe52f, TimeHiAndVersion: 0x4904, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0xbe, Node: [6]uint8{0x93, 0x17, 0xc1, 0xbd, 0xd4, 0x97}}
	MSDLLDusmsvc                                      = uuid.UUID{TimeLow: 0xc27f3c08, TimeMid: 0x92ba, TimeHiAndVersion: 0x478c, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x46, Node: [6]uint8{0xb4, 0x19, 0xc4, 0xce, 0xf0, 0xe2}}
	MSDLLEeprov                                       = uuid.UUID{TimeLow: 0x714dc5c4, TimeMid: 0xc5f6, TimeHiAndVersion: 0x466a, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x37, Node: [6]uint8{0xa5, 0x73, 0xc9, 0x58, 0x3, 0x1e}}
	MSDLLPsmSrvAppHost                                = uuid.UUID{TimeLow: 0x361ae94, TimeMid: 0x316, TimeHiAndVersion: 0x4c6c, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0xd8, Node: [6]uint8{0xc5, 0x94, 0x37, 0x58, 0x0, 0xe2}}
	MSDLLPsmSrvApp                                    = uuid.UUID{TimeLow: 0x85b0334, TimeMid: 0xe454, TimeHiAndVersion: 0x4d91, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x8c, Node: [6]uint8{0x41, 0x34, 0xf9, 0xe7, 0x93, 0xf3}}
	MSDLLPsmSrvInfo                                   = uuid.UUID{TimeLow: 0x3b338d89, TimeMid: 0x6cfa, TimeHiAndVersion: 0x44b8, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x7e, Node: [6]uint8{0x53, 0x15, 0x31, 0xbc, 0x99, 0x92}}
	MSDLLPsmSrvActivation                             = uuid.UUID{TimeLow: 0x4bec6bb8, TimeMid: 0xb5c2, TimeHiAndVersion: 0x4b6f, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0xc1, Node: [6]uint8{0x5d, 0xa5, 0xcf, 0x92, 0xd0, 0xd9}}
	MSDLLPsmSrvTimer                                  = uuid.UUID{TimeLow: 0x5824833b, TimeMid: 0x3c1a, TimeHiAndVersion: 0x4ad2, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xfd, Node: [6]uint8{0xc3, 0x1d, 0x19, 0xe2, 0x3e, 0xd2}}
	MSDLLPsmSrvQuiesce                                = uuid.UUID{TimeLow: 0x8782d3b9, TimeMid: 0xebbd, TimeHiAndVersion: 0x4644, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xd8, Node: [6]uint8{0xe8, 0x72, 0x53, 0x81, 0x91, 0x9b}}
	MSDLLPsmSrv                                       = uuid.UUID{TimeLow: 0xbdaa0970, TimeMid: 0x413b, TimeHiAndVersion: 0x4a3e, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x5d, Node: [6]uint8{0xf6, 0xdc, 0x9d, 0x7e, 0x7, 0x60}}
	MSDLLWscsvc                                       = uuid.UUID{TimeLow: 0x6bba54a, TimeMid: 0xbe05, TimeHiAndVersion: 0x49f9, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xa0, Node: [6]uint8{0x30, 0xf7, 0x90, 0x26, 0x10, 0x23}}
	MSDLLPcaSvc                                       = uuid.UUID{TimeLow: 0x767a036, TimeMid: 0xd22, TimeHiAndVersion: 0x48aa, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x69, Node: [6]uint8{0xb6, 0x19, 0x48, 0xf, 0x38, 0xcb}}
	MSDLLRmgrSrv                                      = uuid.UUID{TimeLow: 0x2c7fd9ce, TimeMid: 0xe706, TimeHiAndVersion: 0x4b40, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x12, Node: [6]uint8{0x95, 0x31, 0x7, 0xef, 0x9b, 0xb0}}
	MSDLLRmtSrv                                       = uuid.UUID{TimeLow: 0xc521facf, TimeMid: 0x9a9, TimeHiAndVersion: 0x42c5, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0x55, Node: [6]uint8{0x72, 0x38, 0x85, 0x95, 0xcb, 0xf0}}
	MSDLLHamRpcSrv                                    = uuid.UUID{TimeLow: 0x82a3471, TimeMid: 0x31b6, TimeHiAndVersion: 0x422a, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x31, Node: [6]uint8{0xa5, 0x44, 0x1, 0x96, 0xc, 0x62}}
	MSDLLHamRpcSrvSess                                = uuid.UUID{TimeLow: 0xd47017b, TimeMid: 0xb33b, TimeHiAndVersion: 0x46ad, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x18, Node: [6]uint8{0xfe, 0x96, 0x45, 0x6c, 0x50, 0x78}}
	MSDLLHamRpcSrvHostId                              = uuid.UUID{TimeLow: 0xff1f646, TimeMid: 0x13bb, TimeHiAndVersion: 0x400a, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x50, Node: [6]uint8{0x9a, 0x78, 0xf2, 0xb7, 0xa8, 0x5a}}
	MSDLLRmGameModeRSrv                               = uuid.UUID{TimeLow: 0x178d84be, TimeMid: 0x9291, TimeHiAndVersion: 0x4994, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0xc6, Node: [6]uint8{0x3f, 0x90, 0x9a, 0xca, 0x5a, 0x3}}
	MSDLLRmCoreRpcSrv                                 = uuid.UUID{TimeLow: 0x1832bcf6, TimeMid: 0xcab8, TimeHiAndVersion: 0x41d4, ClockSeqHiAndReserved: 0x85, ClockSeqLow: 0xd2, Node: [6]uint8{0xc9, 0x41, 0x7, 0x64, 0xf7, 0x5a}}
	MSDLLHamRpcSrvState                               = uuid.UUID{TimeLow: 0x4ed8abcc, TimeMid: 0xf1e2, TimeHiAndVersion: 0x438b, ClockSeqHiAndReserved: 0x98, ClockSeqLow: 0x1f, Node: [6]uint8{0xbb, 0xe, 0x8a, 0xbc, 0x1, 0xc}}
	MSDLLHamRpcSrvFullTrust                           = uuid.UUID{TimeLow: 0x6982a06e, TimeMid: 0x5fe2, TimeHiAndVersion: 0x46b1, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x9c, Node: [6]uint8{0xa2, 0xc5, 0x45, 0xbf, 0xa0, 0x69}}
	MSDLLHamRpcSrvServicing                           = uuid.UUID{TimeLow: 0x95406f0b, TimeMid: 0xb239, TimeHiAndVersion: 0x4318, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xbb, Node: [6]uint8{0xce, 0xa3, 0xa4, 0x6f, 0xf0, 0xdc}}
	MSDLLCrmRpcSrv                                    = uuid.UUID{TimeLow: 0xdd59071b, TimeMid: 0x3215, TimeHiAndVersion: 0x4c59, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x81, Node: [6]uint8{0x97, 0x2e, 0xda, 0xdc, 0xf, 0x6a}}
	MSDLLHamRpcSrvActivity                            = uuid.UUID{TimeLow: 0xe53d94ca, TimeMid: 0x7464, TimeHiAndVersion: 0x4839, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x44, Node: [6]uint8{0x9, 0xa2, 0xfb, 0x8b, 0x3a, 0xe5}}
	MSDLLHamRpcSrvDebug                               = uuid.UUID{TimeLow: 0xfae436b0, TimeMid: 0xb864, TimeHiAndVersion: 0x4a87, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0xda, Node: [6]uint8{0x29, 0x85, 0x47, 0xcd, 0x82, 0xf2}}
	MSDLLRmGameModeSrv                                = uuid.UUID{TimeLow: 0x4dace966, TimeMid: 0xa243, TimeHiAndVersion: 0x4450, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x3f, Node: [6]uint8{0x9b, 0x7b, 0xcb, 0x53, 0x15, 0xb8}}
	MSDLLActiveSyncServer                             = uuid.UUID{TimeLow: 0xc53aa2e, TimeMid: 0xfb1c, TimeHiAndVersion: 0x49c5, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0xb6, Node: [6]uint8{0xc5, 0x4f, 0x8e, 0x58, 0x57, 0xcd}}
	MSDLLAccountsMgmtRpc                              = uuid.UUID{TimeLow: 0x923c9623, TimeMid: 0xdb7f, TimeHiAndVersion: 0x4b34, ClockSeqHiAndReserved: 0x9e, ClockSeqLow: 0x6d, Node: [6]uint8{0xe8, 0x65, 0x80, 0xf8, 0xca, 0x2a}}
	MSDLLIUserMarshalData                             = uuid.UUID{TimeLow: 0xd3c7f20, TimeMid: 0x1c8d, TimeHiAndVersion: 0x4654, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xb3, Node: [6]uint8{0x51, 0x56, 0x3b, 0x29, 0x8b, 0xda}}
	MSDLLUsermgr                                      = uuid.UUID{TimeLow: 0xb18fbab6, TimeMid: 0x56f8, TimeHiAndVersion: 0x4702, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xe0, Node: [6]uint8{0x41, 0x5, 0x32, 0x93, 0xa8, 0x69}}
	MSDLLFmMuxSrvNotificationCoreUI                   = uuid.UUID{TimeLow: 0xfc77b1a, TimeMid: 0x95d8, TimeHiAndVersion: 0x4a2e, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0xc0, Node: [6]uint8{0xcf, 0xf5, 0x42, 0x37, 0x46, 0x2b}}
	MSDLLFmMuxSrvCoreUI                               = uuid.UUID{TimeLow: 0x8ec21e98, TimeMid: 0xb5ce, TimeHiAndVersion: 0x4916, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xd6, Node: [6]uint8{0x44, 0x9f, 0xa4, 0x28, 0xa0, 0x7}}
	MSDLLFmMuxSrv                                     = uuid.UUID{TimeLow: 0xb1ef227e, TimeMid: 0xdfa5, TimeHiAndVersion: 0x421e, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0xbb, Node: [6]uint8{0x67, 0xa6, 0xa1, 0x29, 0xc4, 0x96}}
	MSDLLIdSegSrv                                     = uuid.UUID{TimeLow: 0x1a0d010f, TimeMid: 0x1c33, TimeHiAndVersion: 0x432c, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0xf5, Node: [6]uint8{0x8c, 0xf4, 0xe8, 0x5, 0x30, 0x99}}
	MSDLLTSSDFarmRpc                                  = uuid.UUID{TimeLow: 0x29770a8f, TimeMid: 0x829b, TimeHiAndVersion: 0x4158, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xa2, Node: [6]uint8{0x78, 0xcd, 0x48, 0x85, 0x1, 0xf7}}
	MSDLLTSSDRpc                                      = uuid.UUID{TimeLow: 0xb12fd546, TimeMid: 0xc875, TimeHiAndVersion: 0x4b41, ClockSeqHiAndReserved: 0x97, ClockSeqLow: 0xd8, Node: [6]uint8{0x95, 0x4, 0x87, 0x66, 0x22, 0x2}}
	MSDLLNgcHandlerRpc                                = uuid.UUID{TimeLow: 0x2d24ff0b, TimeMid: 0x1bab, TimeHiAndVersion: 0x404c, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0xfd, Node: [6]uint8{0x42, 0xc8, 0x55, 0x77, 0xbf, 0x68}}
	MSDLLProxyMgr                                     = uuid.UUID{TimeLow: 0x2e6035b2, TimeMid: 0xe8f1, TimeHiAndVersion: 0x41a7, ClockSeqHiAndReserved: 0xa0, ClockSeqLow: 0x44, Node: [6]uint8{0x65, 0x6b, 0x43, 0x9c, 0x4c, 0x34}}
	MSDLLRpcSrvProxyMgr                               = uuid.UUID{TimeLow: 0xc36be077, TimeMid: 0xe14b, TimeHiAndVersion: 0x4fe9, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0xbc, Node: [6]uint8{0xe8, 0x56, 0xef, 0x4f, 0x4, 0x8b}}
	MSDLLTaskScheduler                                = uuid.UUID{TimeLow: 0x33d84484, TimeMid: 0x3626, TimeHiAndVersion: 0x47ee, ClockSeqHiAndReserved: 0x8c, ClockSeqLow: 0x6f, Node: [6]uint8{0xe7, 0xe9, 0x8b, 0x11, 0x3b, 0xe1}}
	MSDLLPhoneRpc                                     = uuid.UUID{TimeLow: 0x3573e5f2, TimeMid: 0xcfe7, TimeHiAndVersion: 0x4a79, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x5f, Node: [6]uint8{0xfe, 0x7c, 0x68, 0x82, 0x37, 0x38}}
	MSDLLPhoneRpc1                                    = uuid.UUID{TimeLow: 0x68227ae7, TimeMid: 0x9a32, TimeHiAndVersion: 0x45b0, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x72, Node: [6]uint8{0xbf, 0x96, 0x19, 0x96, 0x58, 0x38}}
	MSDLLPhoneRpc2                                    = uuid.UUID{TimeLow: 0x8be456ec, TimeMid: 0x9244, TimeHiAndVersion: 0x4d10, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0xe8, Node: [6]uint8{0x1d, 0xdf, 0x1b, 0xaa, 0x9a, 0xde}}
	MSDLLPhoneRpc3                                    = uuid.UUID{TimeLow: 0xae58b386, TimeMid: 0xc914, TimeHiAndVersion: 0x4a73, ClockSeqHiAndReserved: 0xba, ClockSeqLow: 0x5c, Node: [6]uint8{0x2c, 0x3e, 0x27, 0x49, 0xe4, 0x78}}
	MSDLLPhoneRpc4                                    = uuid.UUID{TimeLow: 0xf26e2372, TimeMid: 0xd601, TimeHiAndVersion: 0x44f0, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xb8, Node: [6]uint8{0x25, 0x91, 0xd2, 0xaf, 0x2f, 0x82}}
	MSDLLI_pSchRpc                                    = uuid.UUID{TimeLow: 0x3a9ef155, TimeMid: 0x691d, TimeHiAndVersion: 0x4449, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x5, Node: [6]uint8{0x9, 0xad, 0x57, 0x3, 0x18, 0x23}}
	MSDLLPimIMService                                 = uuid.UUID{TimeLow: 0x43890c94, TimeMid: 0xbfd7, TimeHiAndVersion: 0x4655, ClockSeqHiAndReserved: 0xad, ClockSeqLow: 0x6a, Node: [6]uint8{0xb4, 0xa6, 0x83, 0x97, 0xcd, 0xcb}}
	MSDLLSsdpsrv                                      = uuid.UUID{TimeLow: 0x4b112204, TimeMid: 0xe19, TimeHiAndVersion: 0x11d3, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x2b, Node: [6]uint8{0x0, 0x0, 0xf8, 0x1f, 0xeb, 0x9f}}
	MSDLLWinBioCredentialManager                      = uuid.UUID{TimeLow: 0x4be96a0f, TimeMid: 0x9f52, TimeHiAndVersion: 0x4729, ClockSeqHiAndReserved: 0xa5, ClockSeqLow: 0x1d, Node: [6]uint8{0xc7, 0x6, 0x10, 0xf1, 0x18, 0xb0}}
	MSDLLWinBioServer                                 = uuid.UUID{TimeLow: 0xc0e9671e, TimeMid: 0x33c6, TimeHiAndVersion: 0x4438, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x64, Node: [6]uint8{0x56, 0xb2, 0xe1, 0xb1, 0xc7, 0xb4}}
	MSDLLColorManagementRpcServer                     = uuid.UUID{TimeLow: 0x509bc7ae, TimeMid: 0x77be, TimeHiAndVersion: 0x4ee8, ClockSeqHiAndReserved: 0xb0, ClockSeqLow: 0x7c, Node: [6]uint8{0xd, 0x9, 0x6b, 0xb4, 0x43, 0x45}}
	MSDLLNlasvc                                       = uuid.UUID{TimeLow: 0x4c8d0bef, TimeMid: 0xd7f1, TimeHiAndVersion: 0x49f0, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0x2, Node: [6]uint8{0xca, 0xa0, 0x5f, 0x58, 0xd1, 0x14}}
	MSDLLTokenBindingApi                              = uuid.UUID{TimeLow: 0x51a227ae, TimeMid: 0x825b, TimeHiAndVersion: 0x41f2, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0xa9, Node: [6]uint8{0x1a, 0xc9, 0x55, 0x7a, 0x10, 0x18}}
	MSDLLSymmetricPopKeyApi                           = uuid.UUID{TimeLow: 0x8fb74744, TimeMid: 0xb2ff, TimeHiAndVersion: 0x4c00, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0xd, Node: [6]uint8{0x9e, 0xf9, 0xa1, 0x91, 0xfe, 0x1b}}
	MSDLLNcbRpcSrv                                    = uuid.UUID{TimeLow: 0x5222821f, TimeMid: 0xd5e2, TimeHiAndVersion: 0x4885, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0xf1, Node: [6]uint8{0x5f, 0x61, 0x85, 0xa0, 0xec, 0x41}}
	MSDLLNcbKapi                                      = uuid.UUID{TimeLow: 0x880fd55e, TimeMid: 0x43b9, TimeHiAndVersion: 0x11e0, ClockSeqHiAndReserved: 0xb1, ClockSeqLow: 0xa8, Node: [6]uint8{0xcf, 0x4e, 0xdf, 0xd7, 0x20, 0x85}}
	MSDLLNcbRpcSrv1                                   = uuid.UUID{TimeLow: 0xe40f7b57, TimeMid: 0x7a25, TimeHiAndVersion: 0x4cd3, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0x35, Node: [6]uint8{0x7f, 0x7d, 0x3d, 0xf9, 0xd1, 0x6b}}
	MSDLLStorSvc                                      = uuid.UUID{TimeLow: 0x54b4c689, TimeMid: 0x969a, TimeHiAndVersion: 0x476f, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0xc2, Node: [6]uint8{0x99, 0x8, 0x85, 0xe9, 0xf5, 0x62}}
	MSDLLStorSvc1                                     = uuid.UUID{TimeLow: 0xbe7f785e, TimeMid: 0xe3a, TimeHiAndVersion: 0x4ab7, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xde, Node: [6]uint8{0x7e, 0x46, 0xe4, 0x43, 0xbe, 0x29}}
	MSDLLGcsSrv                                       = uuid.UUID{TimeLow: 0x88abcbc3, TimeMid: 0x34ea, TimeHiAndVersion: 0x76ae, ClockSeqHiAndReserved: 0x82, ClockSeqLow: 0x15, Node: [6]uint8{0x76, 0x75, 0x20, 0x65, 0x5a, 0x23}}
	MSDLLClipSvc                                      = uuid.UUID{TimeLow: 0x64d1d045, TimeMid: 0xf675, TimeHiAndVersion: 0x460b, ClockSeqHiAndReserved: 0x8a, ClockSeqLow: 0x94, Node: [6]uint8{0x57, 0x2, 0x46, 0xb3, 0x6d, 0xab}}
	MSDLLRpSrv                                        = uuid.UUID{TimeLow: 0x55e6b932, TimeMid: 0x1979, TimeHiAndVersion: 0x45d6, ClockSeqHiAndReserved: 0x90, ClockSeqLow: 0xc5, Node: [6]uint8{0x7f, 0x62, 0x70, 0x72, 0x41, 0x12}}
	MSDLLRpSrv1                                       = uuid.UUID{TimeLow: 0x76c217bc, TimeMid: 0xc8b4, TimeHiAndVersion: 0x4201, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x45, Node: [6]uint8{0x37, 0x3a, 0xd9, 0x3, 0x2b, 0x1a}}
	MSDLLRasCustomRpc                                 = uuid.UUID{TimeLow: 0x650a7e26, TimeMid: 0xeab8, TimeHiAndVersion: 0x5533, ClockSeqHiAndReserved: 0xce, ClockSeqLow: 0x43, Node: [6]uint8{0x9c, 0x1d, 0xfc, 0xe1, 0x15, 0x11}}
	MSDLLCSebiPublisher                               = uuid.UUID{TimeLow: 0x1377d115, TimeMid: 0x98fd, TimeHiAndVersion: 0x4034, ClockSeqHiAndReserved: 0xb5, ClockSeqLow: 0x74, Node: [6]uint8{0x11, 0x11, 0x56, 0xca, 0x23, 0x9c}}
	MSDLLNgcSecureBioHandlerRpc                       = uuid.UUID{TimeLow: 0x7642249b, TimeMid: 0x84c2, TimeHiAndVersion: 0x4404, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xeb, Node: [6]uint8{0x1e, 0xa, 0x24, 0x58, 0x83, 0x9a}}
	MSDLLCSebiEvent                                   = uuid.UUID{TimeLow: 0x697dcda9, TimeMid: 0x3ba9, TimeHiAndVersion: 0x4eb2, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0x47, Node: [6]uint8{0xe1, 0x1f, 0x19, 0x1, 0xb0, 0xd2}}
	MSDLLSebiEvent                                    = uuid.UUID{TimeLow: 0x9b008953, TimeMid: 0xf195, TimeHiAndVersion: 0x4bf9, ClockSeqHiAndReserved: 0xbd, ClockSeqLow: 0xe0, Node: [6]uint8{0x44, 0x71, 0x97, 0x1e, 0x58, 0xed}}
	MSDLLSgrmBroker                                   = uuid.UUID{TimeLow: 0x7a20fcec, TimeMid: 0xdec4, TimeHiAndVersion: 0x4c59, ClockSeqHiAndReserved: 0xbe, ClockSeqLow: 0x57, Node: [6]uint8{0x21, 0x2e, 0x8f, 0x65, 0xd3, 0xde}}
	MSDLLSetOperatorRpc                               = uuid.UUID{TimeLow: 0x7aeb6705, TimeMid: 0x3ae6, TimeHiAndVersion: 0x471a, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x2d, Node: [6]uint8{0xf3, 0x9c, 0x10, 0x9e, 0xdc, 0x12}}
	MSDLLWcmSvcRpc                                    = uuid.UUID{TimeLow: 0xabfb6ca3, TimeMid: 0xc5e, TimeHiAndVersion: 0x4734, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0x85, Node: [6]uint8{0xa, 0xee, 0x72, 0xfe, 0x8d, 0x1c}}
	MSDLLWcmSvcRpc1                                   = uuid.UUID{TimeLow: 0xb37f900a, TimeMid: 0xeae4, TimeHiAndVersion: 0x4304, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xab, Node: [6]uint8{0x12, 0xbb, 0x66, 0x8c, 0x1, 0x88}}
	MSDLLWcmSvcProxyInfoRpc                           = uuid.UUID{TimeLow: 0xc2d1b5dd, TimeMid: 0xfa81, TimeHiAndVersion: 0x4460, ClockSeqHiAndReserved: 0x9d, ClockSeqLow: 0xd6, Node: [6]uint8{0xe7, 0x65, 0x8b, 0x85, 0x45, 0x4b}}
	MSDLLWcmSvcMgrRpc                                 = uuid.UUID{TimeLow: 0xe7f76134, TimeMid: 0x9ef5, TimeHiAndVersion: 0x4949, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0xd6, Node: [6]uint8{0x33, 0x68, 0xcc, 0x9, 0x88, 0xf3}}
	MSDLLWcmSvcCapRpc                                 = uuid.UUID{TimeLow: 0xf44e62af, TimeMid: 0xdab1, TimeHiAndVersion: 0x44c2, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0x13, Node: [6]uint8{0x4, 0x9a, 0x9d, 0xe4, 0x17, 0xd6}}
	MSDLLDfsDsSvc                                     = uuid.UUID{TimeLow: 0x7f1343fe, TimeMid: 0x50a9, TimeHiAndVersion: 0x4927, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0x78, Node: [6]uint8{0xc, 0x58, 0x59, 0x51, 0x7b, 0xac}}
	MSDLLFwNetIsolationRpc                            = uuid.UUID{TimeLow: 0xf47433c3, TimeMid: 0x3e9d, TimeHiAndVersion: 0x4157, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xd4, Node: [6]uint8{0x83, 0xaa, 0x1f, 0x5c, 0x2d, 0x4c}}
	MSDLLFwRpc1                                       = uuid.UUID{TimeLow: 0x2fb92682, TimeMid: 0x6599, TimeHiAndVersion: 0x42dc, ClockSeqHiAndReserved: 0xae, ClockSeqLow: 0x13, Node: [6]uint8{0xbd, 0x2c, 0xa8, 0x9b, 0xd1, 0x1c}}
	MSDLLFwRpc                                        = uuid.UUID{TimeLow: 0x7f9d11bf, TimeMid: 0x7fb9, TimeHiAndVersion: 0x436b, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0x12, Node: [6]uint8{0xb2, 0xd5, 0xc, 0x5d, 0x4c, 0x3}}
	MSDLLNetSetupShim                                 = uuid.UUID{TimeLow: 0xa111f1c6, TimeMid: 0x5923, TimeHiAndVersion: 0x47c0, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x68, Node: [6]uint8{0xd0, 0xba, 0xfb, 0x57, 0x79, 0x1}}
	MSDLLNetSetupSvc                                  = uuid.UUID{TimeLow: 0xa111f1c5, TimeMid: 0x5923, TimeHiAndVersion: 0x47c0, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0x68, Node: [6]uint8{0xd0, 0xba, 0xfb, 0x57, 0x79, 0x1}}
	MSDLLDeviceCredentialMgrRpc                       = uuid.UUID{TimeLow: 0x8337aebc, TimeMid: 0x5564, TimeHiAndVersion: 0x46fd, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x41, Node: [6]uint8{0x77, 0x98, 0xf1, 0x8d, 0x2e, 0x4b}}
	MSDLLVscRpc                                       = uuid.UUID{TimeLow: 0x8bef2320, TimeMid: 0xf308, TimeHiAndVersion: 0x4720, ClockSeqHiAndReserved: 0xb9, ClockSeqLow: 0x13, Node: [6]uint8{0x1, 0x29, 0xce, 0xcf, 0xa6, 0xb9}}
	MSDLLNgcLocalAccountVaultRpc                      = uuid.UUID{TimeLow: 0x9cbc9d3a, TimeMid: 0x7586, TimeHiAndVersion: 0x4814, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x70, Node: [6]uint8{0x18, 0x73, 0x7d, 0xcb, 0xe5, 0x23}}
	MSDLLPhoneProviders                               = uuid.UUID{TimeLow: 0xa3bae3f7, TimeMid: 0xbf97, TimeHiAndVersion: 0x49fb, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x8d, Node: [6]uint8{0x2a, 0x5e, 0x86, 0x57, 0xb4, 0x36}}
	MSDLLLicenseManagerSvc                            = uuid.UUID{TimeLow: 0xa4b8d482, TimeMid: 0x80ce, TimeHiAndVersion: 0x40d6, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x4d, Node: [6]uint8{0xb2, 0x2a, 0x1, 0xa4, 0x4f, 0xe7}}
	MSDLLTbiRpc                                       = uuid.UUID{TimeLow: 0xa500d4c6, TimeMid: 0xdd1, TimeHiAndVersion: 0x4543, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0xc, Node: [6]uint8{0xd5, 0xf9, 0x34, 0x86, 0xea, 0xf8}}
	MSDLLDeManagementRpcServer                        = uuid.UUID{TimeLow: 0xa69816f5, TimeMid: 0x83b6, TimeHiAndVersion: 0x4d48, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x33, Node: [6]uint8{0x6, 0x7f, 0x99, 0xf5, 0xf2, 0xd3}}
	MSDLLDeoRpcServer                                 = uuid.UUID{TimeLow: 0xb54e9aa3, TimeMid: 0xcf29, TimeHiAndVersion: 0x4f21, ClockSeqHiAndReserved: 0xa8, ClockSeqLow: 0xea, Node: [6]uint8{0x98, 0xc5, 0x85, 0xc, 0xe2, 0x96}}
	MSDLLPfRpcServerSuperfetch                        = uuid.UUID{TimeLow: 0xb58aa02e, TimeMid: 0x2884, TimeHiAndVersion: 0x4e97, ClockSeqHiAndReserved: 0x81, ClockSeqLow: 0x76, Node: [6]uint8{0x4e, 0xe0, 0x6d, 0x79, 0x41, 0x84}}
	MSDLLWcnpRpc                                      = uuid.UUID{TimeLow: 0xc100beab, TimeMid: 0xd33a, TimeHiAndVersion: 0x4a4b, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x23, Node: [6]uint8{0xbb, 0xef, 0x46, 0x63, 0xd0, 0x17}}
	MSDLLWcnTransportRpc                              = uuid.UUID{TimeLow: 0xc100beac, TimeMid: 0xd33a, TimeHiAndVersion: 0x4a4b, ClockSeqHiAndReserved: 0xbf, ClockSeqLow: 0x23, Node: [6]uint8{0xbb, 0xef, 0x46, 0x63, 0xd0, 0x17}}
	MSDLLNgcMgmt                                      = uuid.UUID{TimeLow: 0xc225e799, TimeMid: 0x29de, TimeHiAndVersion: 0x42af, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x5, Node: [6]uint8{0x1e, 0x21, 0x27, 0xcc, 0x5, 0x6e}}
	MSDLLUdmSvc                                       = uuid.UUID{TimeLow: 0xc8ba73d2, TimeMid: 0x3d55, TimeHiAndVersion: 0x429c, ClockSeqHiAndReserved: 0x8e, ClockSeqLow: 0x9a, Node: [6]uint8{0xc4, 0x4f, 0x0, 0x6f, 0x69, 0xfc}}
	MSDLLUtcAdminApi                                  = uuid.UUID{TimeLow: 0xd4051bde, TimeMid: 0x9cdd, TimeHiAndVersion: 0x4910, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x93, Node: [6]uint8{0x4a, 0xa8, 0x5e, 0xc3, 0xc4, 0x82}}
	MSDLLNgcMgmt1                                     = uuid.UUID{TimeLow: 0xd9844ed9, TimeMid: 0xf72a, TimeHiAndVersion: 0x4745, ClockSeqHiAndReserved: 0xa4, ClockSeqLow: 0xa1, Node: [6]uint8{0xee, 0x71, 0xf9, 0x50, 0x78, 0x1d}}
	MSDLLGidsHandlerRpc                               = uuid.UUID{TimeLow: 0xe6f89680, TimeMid: 0xfc98, TimeHiAndVersion: 0x11e3, ClockSeqHiAndReserved: 0x80, ClockSeqLow: 0xd4, Node: [6]uint8{0x10, 0x60, 0x4b, 0x68, 0x1c, 0xfa}}
	MSDLLHcsRpc                                       = uuid.UUID{TimeLow: 0xe7a216af, TimeMid: 0x1ec1, TimeHiAndVersion: 0x447f, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x3f, Node: [6]uint8{0xa8, 0x72, 0x78, 0xdb, 0x56, 0x4d}}
	MSDLLUSSvc                                        = uuid.UUID{TimeLow: 0xe8748f69, TimeMid: 0xa2a4, TimeHiAndVersion: 0x40df, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x66, Node: [6]uint8{0x62, 0xdb, 0xeb, 0x69, 0x6e, 0x26}}
	MSDLLClusterConnRpc                               = uuid.UUID{TimeLow: 0xf2c9b409, TimeMid: 0xc1c9, TimeHiAndVersion: 0x4100, ClockSeqHiAndReserved: 0x86, ClockSeqLow: 0x39, Node: [6]uint8{0xd8, 0xab, 0x14, 0x86, 0x69, 0x4a}}
	MSDLLWitnessNodesRpc                              = uuid.UUID{TimeLow: 0xeb081a0d, TimeMid: 0x10ee, TimeHiAndVersion: 0x478a, ClockSeqHiAndReserved: 0xa1, ClockSeqLow: 0xdd, Node: [6]uint8{0x50, 0x99, 0x52, 0x83, 0xe7, 0xa8}}
	MSDLLLogonWluirRpc                                = uuid.UUID{TimeLow: 0xf3f09ffd, TimeMid: 0xfbcf, TimeHiAndVersion: 0x4291, ClockSeqHiAndReserved: 0x94, ClockSeqLow: 0x4d, Node: [6]uint8{0x70, 0xad, 0x6e, 0xe, 0x73, 0xbb}}
	MSDLLDeviceCredentialRpc                          = uuid.UUID{TimeLow: 0xfd6b7e61, TimeMid: 0x2bed, TimeHiAndVersion: 0x4d48, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x67, Node: [6]uint8{0xd7, 0x46, 0xfe, 0x44, 0x9f, 0xed}}
	MSDLLNgcRpc                                       = uuid.UUID{TimeLow: 0x30034843, TimeMid: 0x29d, TimeHiAndVersion: 0x46ec, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0xff, Node: [6]uint8{0x5d, 0x12, 0x98, 0x7f, 0x85, 0xc4}}
	MSDLLSECOMNRpc                                    = uuid.UUID{TimeLow: 0xa00e7603, TimeMid: 0x27b5, TimeHiAndVersion: 0x4a1a, ClockSeqHiAndReserved: 0x84, ClockSeqLow: 0x52, Node: [6]uint8{0xd0, 0x1, 0xf4, 0x11, 0x88, 0xa9}}
	MSDLLIhvmsmapi                                    = uuid.UUID{TimeLow: 0xc3f42c6e, TimeMid: 0xd4cc, TimeHiAndVersion: 0x4e5a, ClockSeqHiAndReserved: 0x93, ClockSeqLow: 0x8b, Node: [6]uint8{0x9c, 0x5e, 0x8a, 0x5d, 0x8c, 0x2e}}
	MSDLLWiFiUserRpc                                  = uuid.UUID{TimeLow: 0xaf7fead8, TimeMid: 0xc34a, TimeHiAndVersion: 0x461f, ClockSeqHiAndReserved: 0x88, ClockSeqLow: 0x94, Node: [6]uint8{0x6d, 0x6f, 0xe, 0x5e, 0xdd, 0xcd}}
	MSDLLWlanSvcLowPriv                               = uuid.UUID{TimeLow: 0x266f33b4, TimeMid: 0xc7c1, TimeHiAndVersion: 0x4bd1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x52, Node: [6]uint8{0xdd, 0xb8, 0xf2, 0x21, 0x4e, 0xb0}}
	MSDLLWlanSvc                                      = uuid.UUID{TimeLow: 0x266f33b4, TimeMid: 0xc7c1, TimeHiAndVersion: 0x4bd1, ClockSeqHiAndReserved: 0x8f, ClockSeqLow: 0x52, Node: [6]uint8{0xdd, 0xb8, 0xf2, 0x21, 0x4e, 0xa9}}
	MSDLLWDiagRpc                                     = uuid.UUID{TimeLow: 0x25952c5d, TimeMid: 0x7976, TimeHiAndVersion: 0x4aa1, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xcb, Node: [6]uint8{0xc3, 0x5f, 0x7a, 0xe7, 0x9d, 0x1b}}
	MSDLLRpcDot11ExtIhv                               = uuid.UUID{TimeLow: 0x98e96949, TimeMid: 0xbc59, TimeHiAndVersion: 0x47f1, ClockSeqHiAndReserved: 0x92, ClockSeqLow: 0xd1, Node: [6]uint8{0x8c, 0x25, 0xb4, 0x6f, 0x85, 0xc7}}
	MSDLLSignalRpc                                    = uuid.UUID{TimeLow: 0x7b91411, TimeMid: 0x14fe, TimeHiAndVersion: 0x4856, ClockSeqHiAndReserved: 0x89, ClockSeqLow: 0x82, Node: [6]uint8{0xc3, 0xfa, 0x75, 0x81, 0xd7, 0x84}}
	MSDLLLiveIdSvcNotifyRpc                           = uuid.UUID{TimeLow: 0x572e35b4, TimeMid: 0x1344, TimeHiAndVersion: 0x4565, ClockSeqHiAndReserved: 0x96, ClockSeqLow: 0xa1, Node: [6]uint8{0xf5, 0xdf, 0x3b, 0xfa, 0x89, 0xbb}}
	MSDLLLiveIdSvcRpc                                 = uuid.UUID{TimeLow: 0xcc105610, TimeMid: 0xda03, TimeHiAndVersion: 0x467e, ClockSeqHiAndReserved: 0xbc, ClockSeqLow: 0x73, Node: [6]uint8{0x5b, 0x9e, 0x29, 0x37, 0x45, 0x8d}}
	MSDLLOnlineProviderCertRpc                        = uuid.UUID{TimeLow: 0xfaf2447b, TimeMid: 0xb348, TimeHiAndVersion: 0x4feb, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0xbe, Node: [6]uint8{0xbe, 0xee, 0x5b, 0x7f, 0x77, 0x78}}
	MSDLLUnimdmRpc                                    = uuid.UUID{TimeLow: 0x2f5f6521, TimeMid: 0xcb55, TimeHiAndVersion: 0x1059, ClockSeqHiAndReserved: 0xb4, ClockSeqLow: 0x46, Node: [6]uint8{0x0, 0xdf, 0xb, 0xce, 0x31, 0xdb}}
	MSDLLColorAdapterRpcServer                        = uuid.UUID{TimeLow: 0x3cc5a774, TimeMid: 0x9689, TimeHiAndVersion: 0x4395, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0xd9, Node: [6]uint8{0x2d, 0x51, 0xd6, 0x9a, 0xdc, 0xf}}
	MSDLLSCardCacheRpc                                = uuid.UUID{TimeLow: 0x8833d1d0, TimeMid: 0x965f, TimeHiAndVersion: 0x4216, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0xe9, Node: [6]uint8{0xfb, 0xe5, 0x8c, 0xad, 0x31, 0x0}}
	MSDLLSCardRpc                                     = uuid.UUID{TimeLow: 0xc6b5235a, TimeMid: 0xe413, TimeHiAndVersion: 0x481d, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0xc8, Node: [6]uint8{0x31, 0x68, 0x1b, 0x1f, 0xaa, 0xf5}}
	MSDLLBthEvtBrRpc                                  = uuid.UUID{TimeLow: 0x5dea026d, TimeMid: 0xf999, TimeHiAndVersion: 0x40b1, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x34, Node: [6]uint8{0x21, 0x64, 0xfd, 0x8, 0x67, 0x83}}
	MSDLLRpcBluetoothGatt                             = uuid.UUID{TimeLow: 0xa533b58, TimeMid: 0xed9, TimeHiAndVersion: 0x4085, ClockSeqHiAndReserved: 0xb6, ClockSeqLow: 0xe8, Node: [6]uint8{0x95, 0x79, 0x5e, 0x14, 0x79, 0x72}}
	MSDLLAppXSetTrustLabelRpc                         = uuid.UUID{TimeLow: 0xae2dc901, TimeMid: 0x312d, TimeHiAndVersion: 0x41df, ClockSeqHiAndReserved: 0x8b, ClockSeqLow: 0x79, Node: [6]uint8{0xe8, 0x35, 0xe6, 0x3d, 0xb8, 0x74}}
	MSDLLAppXSetTrustLabel1Rpc                        = uuid.UUID{TimeLow: 0xff9fd3c4, TimeMid: 0x742e, TimeHiAndVersion: 0x45e0, ClockSeqHiAndReserved: 0x91, ClockSeqLow: 0xdd, Node: [6]uint8{0x2f, 0x5b, 0xc6, 0x32, 0xa1, 0xdf}}
	MSDLLVoIpRpc                                      = uuid.UUID{TimeLow: 0x847c6e33, TimeMid: 0x7372, TimeHiAndVersion: 0x4a11, ClockSeqHiAndReserved: 0x9b, ClockSeqLow: 0x86, Node: [6]uint8{0xf5, 0xa4, 0xaf, 0x44, 0x6d, 0xd8}}
	MSDLLWFDSConMgrRpc                                = uuid.UUID{TimeLow: 0x9fa6aff6, TimeMid: 0xe0ad, TimeHiAndVersion: 0x48df, ClockSeqHiAndReserved: 0xa6, ClockSeqLow: 0xb4, Node: [6]uint8{0xe6, 0x4b, 0xe6, 0x6a, 0x17, 0xa1}}
	MSDLLFrameServerRpc                               = uuid.UUID{TimeLow: 0x6ddfc7d1, TimeMid: 0x7fca, TimeHiAndVersion: 0x44eb, ClockSeqHiAndReserved: 0xa2, ClockSeqLow: 0x79, Node: [6]uint8{0xe9, 0x98, 0x8f, 0x4d, 0xb3, 0x2b}}
	MSDLLWwanRpc                                      = uuid.UUID{TimeLow: 0xb4cb7611, TimeMid: 0xad0b, TimeHiAndVersion: 0x4c2d, ClockSeqHiAndReserved: 0xb3, ClockSeqLow: 0x5f, Node: [6]uint8{0xff, 0xe4, 0x57, 0x85, 0xc7, 0x9}}
	MSDLLWwan2Rpc                                     = uuid.UUID{TimeLow: 0x9c56d792, TimeMid: 0x591, TimeHiAndVersion: 0x4431, ClockSeqHiAndReserved: 0x8d, ClockSeqLow: 0x1f, Node: [6]uint8{0x68, 0x1b, 0xfd, 0x80, 0xe4, 0xc0}}
	MSDLLWwanMgmtRpc                                  = uuid.UUID{TimeLow: 0xde2daf3b, TimeMid: 0x5c16, TimeHiAndVersion: 0x4613, ClockSeqHiAndReserved: 0xb2, ClockSeqLow: 0x4, Node: [6]uint8{0xd8, 0x10, 0xee, 0x62, 0x9d, 0x9e}}
	MSDLLIMpService                                   = uuid.UUID{TimeLow: 0xc503f532, TimeMid: 0x443a, TimeHiAndVersion: 0x4c69, ClockSeqHiAndReserved: 0x83, ClockSeqLow: 0x0, Node: [6]uint8{0xcc, 0xd1, 0xfb, 0xdb, 0x38, 0x39}}
	RPCSSIMachineActivatorControl                     = uuid.UUID{TimeLow: 0xc6f3ee72, TimeMid: 0xce7e, TimeHiAndVersion: 0x11d1, ClockSeqHiAndReserved: 0xb7, ClockSeqLow: 0x1e, Node: [6]uint8{0x0, 0xc0, 0x4f, 0xc3, 0x11, 0x1a}}
	RPCSSISCMLocalActivator                           = uuid.UUID{TimeLow: 0x136, TimeMid: 0x0, TimeHiAndVersion: 0x0, ClockSeqHiAndReserved: 0xc0, ClockSeqLow: 0x0, Node: [6]uint8{0x0, 0x0, 0x0, 0x0, 0x0, 0x46}}
	RPCSSISCM                                         = uuid.UUID{TimeLow: 0x412f241e, TimeMid: 0xc12a, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0xff, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6e, 0x7a, 0x17}}
	RPCSSIROT                                         = uuid.UUID{TimeLow: 0xb9e79e60, TimeMid: 0x3d52, TimeHiAndVersion: 0x11ce, ClockSeqHiAndReserved: 0xaa, ClockSeqLow: 0xa1, Node: [6]uint8{0x0, 0x0, 0x69, 0x1, 0x29, 0x3f}}
	RPCSSILocalObjectExporter                         = uuid.UUID{TimeLow: 0xe60c73e6, TimeMid: 0x88f9, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0x9a, ClockSeqLow: 0xf1, Node: [6]uint8{0x0, 0x20, 0xaf, 0x6e, 0x72, 0xf4}}
	RPCSSLocalepmp                                    = uuid.UUID{TimeLow: 0xb0a6584, TimeMid: 0x9e0f, TimeHiAndVersion: 0x11cf, ClockSeqHiAndReserved: 0xa3, ClockSeqLow: 0xcf, Node: [6]uint8{0x0, 0x80, 0x5f, 0x68, 0xcb, 0x1b}}
	RPCSSDbgidl                                       = uuid.UUID{TimeLow: 0x1d55b526, TimeMid: 0xc137, TimeHiAndVersion: 0x46c5, ClockSeqHiAndReserved: 0xab, ClockSeqLow: 0x79, Node: [6]uint8{0x63, 0x8f, 0x2a, 0x68, 0xe8, 0x69}}
	RPCSSFwidl                                        = uuid.UUID{TimeLow: 0x64fe0b7f, TimeMid: 0x9ef5, TimeHiAndVersion: 0x4553, ClockSeqHiAndReserved: 0xa7, ClockSeqLow: 0xdb, Node: [6]uint8{0x9a, 0x19, 0x75, 0x77, 0x75, 0x54}}
)

type UUID uuid.UUID

func (u UUID) Describe() string {
	switch (uuid.UUID)(u) {
	case Null:
		return ": Null: Null"
	case MCCCFGIClusCfgAsyncEvictCleanup:
		return "MC-CCFG: Server Cluster: Configuration (ClusCfg) Protocol: IClusCfgAsyncEvictCleanup"
	case MCIISAIAppHostAdminManager:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostAdminManager"
	case MCIISAIAppHostChangeHandler:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostChangeHandler"
	case MCIISAIAppHostChildElementCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostChildElementCollection"
	case MCIISAIAppHostCollectionSchema:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostCollectionSchema"
	case MCIISAIAppHostConfigException:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConfigException"
	case MCIISAIAppHostConfigFile:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConfigFile"
	case MCIISAIAppHostConfigLocationCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConfigLocationCollection"
	case MCIISAIAppHostConfigLocation:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConfigLocation"
	case MCIISAIAppHostConfigManager:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConfigManager"
	case MCIISAIAppHostConstantValueCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConstantValueCollection"
	case MCIISAIAppHostConstantValue:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostConstantValue"
	case MCIISAIAppHostElementCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostElementCollection"
	case MCIISAIAppHostElementSchemaCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostElementSchemaCollection"
	case MCIISAIAppHostElementSchema:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostElementSchema"
	case MCIISAIAppHostElement:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostElement"
	case MCIISAIAppHostMappingExtension:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostMappingExtension"
	case MCIISAIAppHostMethodCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostMethodCollection"
	case MCIISAIAppHostMethodInstance:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostMethodInstance"
	case MCIISAIAppHostMethodSchema:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostMethodSchema"
	case MCIISAIAppHostMethod:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostMethod"
	case MCIISAIAppHostPathMapper:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostPathMapper"
	case MCIISAIAppHostPropertyCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostPropertyCollection"
	case MCIISAIAppHostPropertyException:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostPropertyException"
	case MCIISAIAppHostPropertySchemaCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostPropertySchemaCollection"
	case MCIISAIAppHostPropertySchema:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostPropertySchema"
	case MCIISAIAppHostProperty:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostProperty"
	case MCIISAIAppHostSectionDefinitionCollection:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostSectionDefinitionCollection"
	case MCIISAIAppHostSectionDefinition:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostSectionDefinition"
	case MCIISAIAppHostSectionGroup:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostSectionGroup"
	case MCIISAIAppHostWritableAdminManager:
		return "MC-IISA: Internet Information Services (IIS) Application Host COM Protocol: IAppHostWritableAdminManager"
	case MCMQACIMSMQApplication2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQApplication2"
	case MCMQACIMSMQApplication3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQApplication3"
	case MCMQACIMSMQApplication:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQApplication"
	case MCMQACIMSMQCollection:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQCollection"
	case MCMQACIMSMQCoordinatedTransactionDispenser2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQCoordinatedTransactionDispenser2"
	case MCMQACIMSMQCoordinatedTransactionDispenser3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQCoordinatedTransactionDispenser3"
	case MCMQACIMSMQCoordinatedTransactionDispenser:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQCoordinatedTransactionDispenser"
	case MCMQACIMSMQDestination:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQDestination"
	case MCMQACIMSMQEvent2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQEvent2"
	case MCMQACIMSMQEvent3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQEvent3"
	case MCMQACIMSMQEvent:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQEvent"
	case MCMQACIMSMQManagement:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQManagement"
	case MCMQACIMSMQMessage2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQMessage2"
	case MCMQACIMSMQMessage3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQMessage3"
	case MCMQACIMSMQMessage4:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQMessage4"
	case MCMQACIMSMQMessage:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQMessage"
	case MCMQACIMSMQOutgoingQueueManagement:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQOutgoingQueueManagement"
	case MCMQACIMSMQPrivateDestination:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQPrivateDestination"
	case MCMQACIMSMQPrivateEvent:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQPrivateEvent"
	case MCMQACIMSMQQuery2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQuery2"
	case MCMQACIMSMQQuery3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQuery3"
	case MCMQACIMSMQQuery4:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQuery4"
	case MCMQACIMSMQQuery:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQuery"
	case MCMQACIMSMQQueue2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueue2"
	case MCMQACIMSMQQueue3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueue3"
	case MCMQACIMSMQQueue4:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueue4"
	case MCMQACIMSMQQueueInfo2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfo2"
	case MCMQACIMSMQQueueInfo3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfo3"
	case MCMQACIMSMQQueueInfo4:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfo4"
	case MCMQACIMSMQQueueInfos2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfos2"
	case MCMQACIMSMQQueueInfos3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfos3"
	case MCMQACIMSMQQueueInfos4:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfos4"
	case MCMQACIMSMQQueueInfos:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfos"
	case MCMQACIMSMQQueueInfo:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueInfo"
	case MCMQACIMSMQQueueManagement:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueueManagement"
	case MCMQACIMSMQQueue:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQQueue"
	case MCMQACIMSMQTransaction2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransaction2"
	case MCMQACIMSMQTransaction3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransaction3"
	case MCMQACIMSMQTransactionDispenser2:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransactionDispenser2"
	case MCMQACIMSMQTransactionDispenser3:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransactionDispenser3"
	case MCMQACIMSMQTransactionDispenser:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransactionDispenser"
	case MCMQACIMSMQTransaction:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: IMSMQTransaction"
	case MCMQACITransaction:
		return "MC-MQAC: Message Queuing (MSMQ): ActiveX Client Protocol: ITransaction"
	case MSADTGIDataFactory2:
		return "MS-ADTG: Remote Data Services (RDS) Transport Protocol: IDataFactory2"
	case MSADTGIDataFactory3:
		return "MS-ADTG: Remote Data Services (RDS) Transport Protocol: IDataFactory3"
	case MSADTGIDataFactory:
		return "MS-ADTG: Remote Data Services (RDS) Transport Protocol: IDataFactory"
	case MSADTSClaims:
		return "MS-ADTS: Active Directory Technical Specification: Claims"
	case MSBKRPBackupKey:
		return "MS-BKRP: BackupKey Remote Protocol: BackupKey"
	case MSBPAUBitsPeerAuth:
		return "MS-BPAU: Background Intelligent Transfer Service (BITS) Peer-Caching: Peer Authentication Protocol: BitsPeerAuth"
	case MSBRWSABrowser:
		return "MS-BRWSA: Common Internet File System (CIFS) Browser Auxiliary Protocol: Browser"
	case MSCAPRLsacap:
		return "MS-CAPR: Central Access Policy Identifier (ID) Retrieval Protocol: lsacap"
	case MSCMPOIXnRemote:
		return "MS-CMPO: MSDTC Connection Manager: OleTx Transports Protocol: IXnRemote"
	case MSCMRPClusAPI:
		return "MS-CMRP: Failover Cluster: Management API (ClusAPI) Protocol: ClusAPI"
	case MSCOMAIAlternateLaunch:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IAlternateLaunch"
	case MSCOMAICapabilitySupport:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICapabilitySupport"
	case MSCOMAICatalog64BitSupport:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalog64BitSupport"
	case MSCOMAICatalogSession:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogSession"
	case MSCOMAICatalogTableInfo:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogTableInfo"
	case MSCOMAICatalogTableRead:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogTableRead"
	case MSCOMAICatalogTableWrite:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogTableWrite"
	case MSCOMAICatalogUtils2:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogUtils2"
	case MSCOMAICatalogUtils:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: ICatalogUtils"
	case MSCOMAIContainerControl2:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IContainerControl2"
	case MSCOMAIContainerControl:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IContainerControl"
	case MSCOMAIExport2:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IExport2"
	case MSCOMAIExport:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IExport"
	case MSCOMAIImport2:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IImport2"
	case MSCOMAIImport:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IImport"
	case MSCOMAIRegister2:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IRegister2"
	case MSCOMAIRegister:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IRegister"
	case MSCOMAIReplicationUtil:
		return "MS-COMA: Component Object Model Plus (COM+) Remote Administration Protocol: IReplicationUtil"
	case MSCOMITransactionStream:
		return "MS-COM: Component Object Model Plus (COM+) Protocol: ITransactionStream"
	case MSCOMEVIEnumEventObject:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEnumEventObject"
	case MSCOMEVIEventClass2:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventClass2"
	case MSCOMEVIEventClass3:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventClass3"
	case MSCOMEVIEventClass:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventClass"
	case MSCOMEVIEventObjectCollection:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventObjectCollection"
	case MSCOMEVIEventSubscription2:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSubscription2"
	case MSCOMEVIEventSubscription3:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSubscription3"
	case MSCOMEVIEventSubscription:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSubscription"
	case MSCOMEVIEventSystem2:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSystem2"
	case MSCOMEVIEventSystemInitialize:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSystemInitialize"
	case MSCOMEVIEventSystem:
		return "MS-COMEV: Component Object Model Plus (COM+) Event System Protocol: IEventSystem"
	case MSCOMTIComTrackingInfoEvents:
		return "MS-COMT: Component Object Model Plus (COM+) Tracker Service Protocol: IComTrackingInfoEvents"
	case MSCOMTIGetTrackingData:
		return "MS-COMT: Component Object Model Plus (COM+) Tracker Service Protocol: IGetTrackingData"
	case MSCOMTIProcessDump:
		return "MS-COMT: Component Object Model Plus (COM+) Tracker Service Protocol: IProcessDump"
	case MSCSRAICertAdminD2:
		return "MS-CSRA: Certificate Services Remote Administration Protocol: ICertAdminD2"
	case MSCSRAICertAdminD:
		return "MS-CSRA: Certificate Services Remote Administration Protocol: ICertAdminD"
	case MSCSVPIClusterCleanup:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterCleanup"
	case MSCSVPIClusterFirewall:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterFirewall"
	case MSCSVPIClusterLog:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterLog"
	case MSCSVPIClusterNetwork2:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterNetwork2"
	case MSCSVPIClusterSetup:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterSetup"
	case MSCSVPIClusterStorage2:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterStorage2"
	case MSCSVPIClusterStorage3:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterStorage3"
	case MSCSVPIClusterUpdate:
		return "MS-CSVP: Failover Cluster: Setup and Validation Protocol (ClusPrep): IClusterUpdate"
	case MSDCOMIActivation:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IActivation"
	case MSDCOMIObjectExporter:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IObjectExporter"
	case MSDCOMIRemoteSCMActivator:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IRemoteSCMActivator"
	case MSDCOMIRemUnknown2:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IRemUnknown2"
	case MSDCOMIRemUnknown:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IRemUnknown"
	case MSDCOMIUnknown:
		return "MS-DCOM: Distributed Component Object Model (DCOM) Remote Protocol: IUnknown"
	case MSDFSNMNetDFS:
		return "MS-DFSNM: Distributed File System (DFS): Namespace Management Protocol: NetDFS"
	case MSDFSRHIADProxy2:
		return "MS-DFSRH: DFS Replication Helper Protocol: IADProxy2"
	case MSDFSRHIADProxy:
		return "MS-DFSRH: DFS Replication Helper Protocol: IADProxy"
	case MSDFSRHIServerHealthReport2:
		return "MS-DFSRH: DFS Replication Helper Protocol: IServerHealthReport2"
	case MSDFSRHIServerHealthReport:
		return "MS-DFSRH: DFS Replication Helper Protocol: IServerHealthReport"
	case MSDHCPMDhcpsrv2:
		return "MS-DHCPM: Microsoft Dynamic Host Configuration Protocol (DHCP) Server Management Protocol: dhcpsrv2"
	case MSDHCPMDhcpsrv:
		return "MS-DHCPM: Microsoft Dynamic Host Configuration Protocol (DHCP) Server Management Protocol: dhcpsrv"
	case MSDLTMTrksvr:
		return "MS-DLTM: Distributed Link Tracking: Central Manager Protocol: trksvr"
	case MSDLTWTrkwks:
		return "MS-DLTW: Distributed Link Tracking: Workstation Protocol: trkwks"
	case MSDMRPIDMNotify:
		return "MS-DMRP: Disk Management Remote Protocol: IDMNotify"
	case MSDMRPIDMRemoteServer:
		return "MS-DMRP: Disk Management Remote Protocol: IDMRemoteServer"
	case MSDMRPIVolumeClient2:
		return "MS-DMRP: Disk Management Remote Protocol: IVolumeClient2"
	case MSDMRPIVolumeClient3:
		return "MS-DMRP: Disk Management Remote Protocol: IVolumeClient3"
	case MSDMRPIVolumeClient4:
		return "MS-DMRP: Disk Management Remote Protocol: IVolumeClient4"
	case MSDMRPIVolumeClient:
		return "MS-DMRP: Disk Management Remote Protocol: IVolumeClient"
	case MSDNSPDnsServer:
		return "MS-DNSP: Domain Name Service (DNS) Server Management Protocol: DnsServer"
	case MSDRSRDrsuapi:
		return "MS-DRSR: Directory Replication Service (DRS) Remote Protocol: drsuapi"
	case MSDRSRDsaop:
		return "MS-DRSR: Directory Replication Service (DRS) Remote Protocol: dsaop"
	case MSDSSPDssetup:
		return "MS-DSSP: Directory Services Setup Remote Protocol: dssetup"
	case MSEFSREfsrpc:
		return "MS-EFSR: Encrypting File System Remote (EFSRPC) Protocol: efsrpc"
	case MSEFSRLsarpc:
		return "MS-EFSR: Encrypting File System Remote (EFSRPC) Protocol: lsarpc"
	case MSEVEN6Eventlog:
		return "MS-EVEN6: EventLog Remoting Protocol Version 6.0: Eventlog"
	case MSEVENEventlog:
		return "MS-EVEN: EventLog Remoting Protocol: eventlog"
	case MSFASPRemoteFW:
		return "MS-FASP: Firewall and Advanced Security Protocol: RemoteFW"
	case MSFAXFaxclient:
		return "MS-FAX: Fax Server and Client Remote Protocol: faxclient"
	case MSFAXFax:
		return "MS-FAX: Fax Server and Client Remote Protocol: fax"
	case MSFRS1Frsrpc:
		return "MS-FRS1: File Replication Service Protocol: frsrpc"
	case MSFRS1NtFrsApi:
		return "MS-FRS1: File Replication Service Protocol: NtFrsApi"
	case MSFRS2FrsTransport:
		return "MS-FRS2: Distributed File System Replication Protocol: FrsTransport"
	case MSFRS2DfsrEndpoint:
		return "MS-FRS2: Distributed File System Replication Protocol: DfsrEndpoint"
	case MSFSRMIFsrmActionCommand:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmActionCommand"
	case MSFSRMIFsrmActionEmail2:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmActionEmail2"
	case MSFSRMIFsrmActionEmail:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmActionEmail"
	case MSFSRMIFsrmActionEventLog:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmActionEventLog"
	case MSFSRMIFsrmActionReport:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmActionReport"
	case MSFSRMIFsrmAction:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmAction"
	case MSFSRMIFsrmAutoApplyQuota:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmAutoApplyQuota"
	case MSFSRMIFsrmClassificationManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmClassificationManager"
	case MSFSRMIFsrmClassificationRule:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmClassificationRule"
	case MSFSRMIFsrmClassifierModuleDefinition:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmClassifierModuleDefinition"
	case MSFSRMIFsrmCollection:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmCollection"
	case MSFSRMIFsrmCommittableCollection:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmCommittableCollection"
	case MSFSRMIFsrmDerivedObjectsResult:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmDerivedObjectsResult"
	case MSFSRMIFsrmFileGroupImported:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileGroupImported"
	case MSFSRMIFsrmFileGroupManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileGroupManager"
	case MSFSRMIFsrmFileGroup:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileGroup"
	case MSFSRMIFsrmFileManagementJobManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileManagementJobManager"
	case MSFSRMIFsrmFileManagementJob:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileManagementJob"
	case MSFSRMIFsrmFileScreenBase:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenBase"
	case MSFSRMIFsrmFileScreenException:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenException"
	case MSFSRMIFsrmFileScreenManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenManager"
	case MSFSRMIFsrmFileScreenTemplateImported:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenTemplateImported"
	case MSFSRMIFsrmFileScreenTemplateManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenTemplateManager"
	case MSFSRMIFsrmFileScreenTemplate:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreenTemplate"
	case MSFSRMIFsrmFileScreen:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmFileScreen"
	case MSFSRMIFsrmMutableCollection:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmMutableCollection"
	case MSFSRMIFsrmObject:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmObject"
	case MSFSRMIFsrmPathMapper:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPathMapper"
	case MSFSRMIFsrmPipelineModuleDefinition:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPipelineModuleDefinition"
	case MSFSRMIFsrmPropertyCondition:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPropertyCondition"
	case MSFSRMIFsrmPropertyDefinition2:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPropertyDefinition2"
	case MSFSRMIFsrmPropertyDefinition:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPropertyDefinition"
	case MSFSRMIFsrmPropertyDefinitionValue:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmPropertyDefinitionValue"
	case MSFSRMIFsrmProperty:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmProperty"
	case MSFSRMIFsrmQuotaBase:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaBase"
	case MSFSRMIFsrmQuotaManagerEx:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaManagerEx"
	case MSFSRMIFsrmQuotaManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaManager"
	case MSFSRMIFsrmQuotaObject:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaObject"
	case MSFSRMIFsrmQuotaTemplateImported:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaTemplateImported"
	case MSFSRMIFsrmQuotaTemplateManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaTemplateManager"
	case MSFSRMIFsrmQuotaTemplate:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuotaTemplate"
	case MSFSRMIFsrmQuota:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmQuota"
	case MSFSRMIFsrmReportJob:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmReportJob"
	case MSFSRMIFsrmReportManager:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmReportManager"
	case MSFSRMIFsrmReportScheduler:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmReportScheduler"
	case MSFSRMIFsrmReport:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmReport"
	case MSFSRMIFsrmRule:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmRule"
	case MSFSRMIFsrmSetting:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmSetting"
	case MSFSRMIFsrmStorageModuleDefinition:
		return "MS-FSRM: File Server Resource Manager Protocol: IFsrmStorageModuleDefinition"
	case MSFSRVPFileServerVssAgent:
		return "MS-FSRVP: File Server Remote VSS Protocol: FileServerVssAgent"
	case MSGKDIISDKey:
		return "MS-GKDI: Group Key Distribution Protocol: ISDKey"
	case MSICPRICertPassage:
		return "MS-ICPR: ICertPassage Remote Protocol: ICertPassage"
	case MSIISSIIisServiceControl:
		return "MS-IISS: Internet Information Services (IIS) ServiceControl: IIisServiceControl"
	case MSIMSAIIISApplicationAdmin:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IIISApplicationAdmin"
	case MSIMSAIIISCertObj:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IIISCertObj"
	case MSIMSAIMSAdminBase2W:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IMSAdminBase2W"
	case MSIMSAIMSAdminBase3W:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IMSAdminBase3W"
	case MSIMSAIMSAdminBaseW:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IMSAdminBaseW"
	case MSIMSAIWamAdmin2:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IWamAdmin2"
	case MSIMSAIWamAdmin:
		return "MS-IMSA: Internet Information Services (IIS) IMSAdminBaseW Remote Protocol: IWamAdmin"
	case MSIOIIManagedObject:
		return "MS-IOI: IManagedObject Interface Protocol: IManagedObject"
	case MSIOIIRemoteDispatch:
		return "MS-IOI: IManagedObject Interface Protocol: IRemoteDispatch"
	case MSIOIIServicedComponentInfo:
		return "MS-IOI: IManagedObject Interface Protocol: IServicedComponentInfo"
	case MSIRPInetinfo:
		return "MS-IRP: Internet Information Services (IIS) Inetinfo Remote: inetinfo"
	case MSISTMIDirectoryEnum:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IDirectoryEnum"
	case MSISTMIEnumCachedInitiator:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumCachedInitiator"
	case MSISTMIEnumConnection:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumConnection"
	case MSISTMIEnumDisk2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumDisk2"
	case MSISTMIEnumDisk:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumDisk"
	case MSISTMIEnumHost:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumHost"
	case MSISTMIEnumIDMethod:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumIDMethod"
	case MSISTMIEnumLMMountPoint:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumLMMountPoint"
	case MSISTMIEnumPortal:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumPortal"
	case MSISTMIEnumResourceGroup:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumResourceGroup"
	case MSISTMIEnumSession2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumSession2"
	case MSISTMIEnumSession:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumSession"
	case MSISTMIEnumSnapshot:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumSnapshot"
	case MSISTMIEnumSnsServer:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumSnsServer"
	case MSISTMIEnumVolume2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumVolume2"
	case MSISTMIEnumVolume:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumVolume"
	case MSISTMIEnumWTDiskLunMapping:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumWTDiskLunMapping"
	case MSISTMIEnumWTDisk:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IEnumWTDisk"
	case MSISTMIHost2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHost2"
	case MSISTMIHost3:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHost3"
	case MSISTMIHostMgr2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHostMgr2"
	case MSISTMIHostMgr3:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHostMgr3"
	case MSISTMIHostMgr:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHostMgr"
	case MSISTMIHost:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IHost"
	case MSISTMILocalDeviceMgr2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ILocalDeviceMgr2"
	case MSISTMILocalDeviceMgr3:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ILocalDeviceMgr3"
	case MSISTMILocalDeviceMgr:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ILocalDeviceMgr"
	case MSISTMISessionManager2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ISessionManager2"
	case MSISTMISessionManager:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ISessionManager"
	case MSISTMISnapshotMgr:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ISnapshotMgr"
	case MSISTMISnapshot:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ISnapshot"
	case MSISTMISnsMgr:
		return "MS-ISTM: iSCSI Software Target Management Protocol: ISnsMgr"
	case MSISTMIStatusNotify:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IStatusNotify"
	case MSISTMIWTDisk2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDisk2"
	case MSISTMIWTDisk3:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDisk3"
	case MSISTMIWTDiskMgr2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDiskMgr2"
	case MSISTMIWTDiskMgr3:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDiskMgr3"
	case MSISTMIWTDiskMgr4:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDiskMgr4"
	case MSISTMIWTDiskMgr:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDiskMgr"
	case MSISTMIWTDisk:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTDisk"
	case MSISTMIWTGeneral2:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTGeneral2"
	case MSISTMIWTGeneral:
		return "MS-ISTM: iSCSI Software Target Management Protocol: IWTGeneral"
	case MSLRECNetEventForwarder:
		return "MS-LREC: Live Remote Event Capture (LREC) Protocol: NetEventForwarder"
	case MSLSADLsarpc:
		return "MS-LSAD: Local Security Authority (Domain Policy) Remote Protocol: lsarpc"
	case MSLSATLsarpc:
		return "MS-LSAT: Local Security Authority (Translation Methods) Remote Protocol: lsarpc"
	case MSMQDSDscomm2:
		return "MS-MQDS: Message Queuing (MSMQ): Directory Service Protocol: dscomm2"
	case MSMQDSDscomm:
		return "MS-MQDS: Message Queuing (MSMQ): Directory Service Protocol: dscomm"
	case MSMQMPQmcomm2:
		return "MS-MQMP: Message Queuing (MSMQ): Queue Manager Client Protocol: qmcomm2"
	case MSMQMPQmcomm:
		return "MS-MQMP: Message Queuing (MSMQ): Queue Manager Client Protocol: qmcomm"
	case MSMQMRQmmgmt:
		return "MS-MQMR: Message Queuing (MSMQ): Queue Manager Management Protocol: qmmgmt"
	case MSMQQPQm2qm:
		return "MS-MQQP: Message Queuing (MSMQ): Queue Manager to Queue Manager Protocol: qm2qm"
	case MSMQRRRemoteRead:
		return "MS-MQRR: Message Queuing (MSMQ): Queue Manager Remote Read Protocol: RemoteRead"
	case MSMSRPMsgsvcsend:
		return "MS-MSRP: Messenger Service Remote Protocol: msgsvcsend"
	case MSMSRPMsgsvc:
		return "MS-MSRP: Messenger Service Remote Protocol: msgsvc"
	case MSNRPCLogon:
		return "MS-NRPC: Netlogon Remote Protocol: logon"
	case MSNSPINspi:
		return "MS-NSPI: Name Service Provider Interface (NSPI) Protocol: nspi"
	case MSOAUTIDispatch:
		return "MS-OAUT: OLE Automation Protocol: IDispatch"
	case MSOAUTIEnumVARIANT:
		return "MS-OAUT: OLE Automation Protocol: IEnumVARIANT"
	case MSOAUTITypeComp:
		return "MS-OAUT: OLE Automation Protocol: ITypeComp"
	case MSOAUTITypeInfo2:
		return "MS-OAUT: OLE Automation Protocol: ITypeInfo2"
	case MSOAUTITypeInfo:
		return "MS-OAUT: OLE Automation Protocol: ITypeInfo"
	case MSOAUTITypeLib2:
		return "MS-OAUT: OLE Automation Protocol: ITypeLib2"
	case MSOAUTITypeLib:
		return "MS-OAUT: OLE Automation Protocol: ITypeLib"
	case MSOAUTIUnknown:
		return "MS-OAUT: OLE Automation Protocol: IUnknown"
	case MSOCSPAIOCSPAdminD:
		return "MS-OCSPA: Microsoft OCSP Administration Protocol: IOCSPAdminD"
	case MSPANIRPCAsyncNotify:
		return "MS-PAN: Print System Asynchronous Notification Protocol: IRPCAsyncNotify"
	case MSPANIRPCRemoteObject:
		return "MS-PAN: Print System Asynchronous Notification Protocol: IRPCRemoteObject"
	case MSPARIRemoteWinspool:
		return "MS-PAR: Print System Asynchronous Remote Protocol: IRemoteWinspool"
	case MSPCQPerflibV2:
		return "MS-PCQ: Performance Counter Query Protocol: PerflibV2"
	case MSPLAIAlertDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: IAlertDataCollector"
	case MSPLAIApiTracingDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: IApiTracingDataCollector"
	case MSPLAIConfigurationDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: IConfigurationDataCollector"
	case MSPLAIDataCollectorCollection:
		return "MS-PLA: Performance Logs and Alerts Protocol: IDataCollectorCollection"
	case MSPLAIDataCollectorSetCollection:
		return "MS-PLA: Performance Logs and Alerts Protocol: IDataCollectorSetCollection"
	case MSPLAIDataCollectorSet:
		return "MS-PLA: Performance Logs and Alerts Protocol: IDataCollectorSet"
	case MSPLAIDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: IDataCollector"
	case MSPLAIDataManager:
		return "MS-PLA: Performance Logs and Alerts Protocol: IDataManager"
	case MSPLAIFolderActionCollection:
		return "MS-PLA: Performance Logs and Alerts Protocol: IFolderActionCollection"
	case MSPLAIFolderAction:
		return "MS-PLA: Performance Logs and Alerts Protocol: IFolderAction"
	case MSPLAIPerformanceCounterDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: IPerformanceCounterDataCollector"
	case MSPLAIScheduleCollection:
		return "MS-PLA: Performance Logs and Alerts Protocol: IScheduleCollection"
	case MSPLAISchedule:
		return "MS-PLA: Performance Logs and Alerts Protocol: ISchedule"
	case MSPLAITraceDataCollector:
		return "MS-PLA: Performance Logs and Alerts Protocol: ITraceDataCollector"
	case MSPLAITraceDataProviderCollection:
		return "MS-PLA: Performance Logs and Alerts Protocol: ITraceDataProviderCollection"
	case MSPLAITraceDataProvider:
		return "MS-PLA: Performance Logs and Alerts Protocol: ITraceDataProvider"
	case MSPLAIValueMapItem:
		return "MS-PLA: Performance Logs and Alerts Protocol: IValueMapItem"
	case MSPLAIValueMap:
		return "MS-PLA: Performance Logs and Alerts Protocol: IValueMap"
	case MSRAAAuthzr:
		return "MS-RAA: Remote Authorization API Protocol: authzr"
	case MSRAAIgnoreCentralPolicy:
		return "MS-RAA: Ignore Central Policy Object: IgnoreCentralPolicy"
	case MSRAAUseCentralPolicy:
		return "MS-RAA: Take Central Policy Into Account Object: UseCentralPolicy"
	case MSRAINPSIIASDataStoreComServer2:
		return "MS-RAINPS: Remote Administrative Interface: Network Policy Server (NPS) Protocol: IIASDataStoreComServer2"
	case MSRAINPSIIASDataStoreComServer:
		return "MS-RAINPS: Remote Administrative Interface: Network Policy Server (NPS) Protocol: IIASDataStoreComServer"
	case MSRAIIPCHCollection:
		return "MS-RAI: Remote Assistance Initiation Protocol: IPCHCollection"
	case MSRAIIPCHService:
		return "MS-RAI: Remote Assistance Initiation Protocol: IPCHService"
	case MSRAIIRASrv:
		return "MS-RAI: Remote Assistance Initiation Protocol: IRASrv"
	case MSRAIISAFSession:
		return "MS-RAI: Remote Assistance Initiation Protocol: ISAFSession"
	case MSRAIWWinsif2:
		return "MS-RAIW: Remote Administrative Interface: WINS: winsif2"
	case MSRAIWWinsif:
		return "MS-RAIW: Remote Administrative Interface: WINS: winsif"
	case MSRDPESCTypeScardPack:
		return "MS-RDPESC: Remote Desktop Protocol: Smart Card Virtual Channel Extension: TypeScardPack"
	case MSRPCEEndpointMapper:
		return "MS-RPCE: Endpoint Mapper Protocol: EndpointMapper"
	case MSRPCERemoteManagement:
		return "MS-RPCE: Remote Management Protocol: RemoteManagement"
	case MSRPCETransferNDR64:
		return "MS-RPCE: Transfer Syntax NDR64: TransferNDR64"
	case MSRPCETransferNDR:
		return "MS-RPCE: Transfer Syntax NDR: TransferNDR"
	case MSRPCLLocToLoc:
		return "MS-RPCL: Remote Procedure Call Location Services Extensions: LocToLoc"
	case MSRPRNWinspool:
		return "MS-RPRN: Print System Remote Protocol: winspool"
	case MSRRASMDIMSVC:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: DIMSVC"
	case MSRRASMIRemoteICFICSConfig:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteICFICSConfig"
	case MSRRASMIRemoteIPV6Config:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteIPV6Config"
	case MSRRASMIRemoteNetworkConfig:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteNetworkConfig"
	case MSRRASMIRemoteRouterRestart:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteRouterRestart"
	case MSRRASMIRemoteSetDnsConfig:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteSetDnsConfig"
	case MSRRASMIRemoteSstpCertCheck:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteSstpCertCheck"
	case MSRRASMIRemoteStringIdConfig:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: IRemoteStringIdConfig"
	case MSRRASMRASRPC:
		return "MS-RRASM: Routing and Remote Access Server (RRAS) Management Protocol: RASRPC"
	case MSRRPWinreg:
		return "MS-RRP: Windows Remote Registry Protocol: winreg"
	case MSRSMPIClientSink:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: IClientSink"
	case MSRSMPIMessenger:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: IMessenger"
	case MSRSMPINtmsLibraryControl1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsLibraryControl1"
	case MSRSMPINtmsLibraryControl2:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsLibraryControl2"
	case MSRSMPINtmsMediaServices1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsMediaServices1"
	case MSRSMPINtmsNotifySink:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsNotifySink"
	case MSRSMPINtmsObjectInfo1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsObjectInfo1"
	case MSRSMPINtmsObjectManagement1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsObjectManagement1"
	case MSRSMPINtmsObjectManagement2:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsObjectManagement2"
	case MSRSMPINtmsObjectManagement3:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsObjectManagement3"
	case MSRSMPINtmsSession1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: INtmsSession1"
	case MSRSMPIRobustNtmsMediaServices1:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: IRobustNtmsMediaServices1"
	case MSRSMPIUnknown:
		return "MS-RSMP: Removable Storage Manager (RSM) Remote Protocol: IUnknown"
	case MSRSPInitShutdown:
		return "MS-RSP: Remote Shutdown Protocol: InitShutdown"
	case MSRSPWindowsShutdown:
		return "MS-RSP: Remote Shutdown Protocol: WindowsShutdown"
	case MSRSPWinreg:
		return "MS-RSP: Remote Shutdown Protocol: winreg"
	case MSSAMRSamr:
		return "MS-SAMR: Security Account Manager (SAM) Remote Protocol (Client-to-Server): samr"
	case MSSCMPIVssDifferentialSoftwareSnapshotMgmt:
		return "MS-SCMP: Shadow Copy Management Protocol: IVssDifferentialSoftwareSnapshotMgmt"
	case MSSCMPIVssEnumMgmtObject:
		return "MS-SCMP: Shadow Copy Management Protocol: IVssEnumMgmtObject"
	case MSSCMPIVssEnumObject:
		return "MS-SCMP: Shadow Copy Management Protocol: IVssEnumObject"
	case MSSCMPIVssSnapshotMgmt:
		return "MS-SCMP: Shadow Copy Management Protocol: IVssSnapshotMgmt"
	case MSSCMRSvcctl:
		return "MS-SCMR: Service Control Manager Remote Protocol: svcctl"
	case MSSRVSSrvsvc:
		return "MS-SRVS: Server Service Remote Protocol: srvsvc"
	case MSSWNWitness:
		return "MS-SWN: Service Witness Protocol: Witness"
	case MSTPMVSCITpmVirtualSmartCardManager2:
		return "MS-TPMVSC: Trusted Platform Module (TPM) Virtual Smart Card Management Protocol: ITpmVirtualSmartCardManager2"
	case MSTPMVSCITpmVirtualSmartCardManager3:
		return "MS-TPMVSC: Trusted Platform Module (TPM) Virtual Smart Card Management Protocol: ITpmVirtualSmartCardManager3"
	case MSTPMVSCITpmVirtualSmartCardManagerStatusCallback:
		return "MS-TPMVSC: Trusted Platform Module (TPM) Virtual Smart Card Management Protocol: ITpmVirtualSmartCardManagerStatusCallback"
	case MSTPMVSCITpmVirtualSmartCardManager:
		return "MS-TPMVSC: Trusted Platform Module (TPM) Virtual Smart Card Management Protocol: ITpmVirtualSmartCardManager"
	case MSTRPRemotesp:
		return "MS-TRP: Telephony Remote Protocol: remotesp"
	case MSTRPTapsrv:
		return "MS-TRP: Telephony Remote Protocol: tapsrv"
	case MSTSCHATSvc:
		return "MS-TSCH: Task Scheduler Service Remoting Protocol: ATSvc"
	case MSTSCHITaskSchedulerService:
		return "MS-TSCH: Task Scheduler Service Remoting Protocol: ITaskSchedulerService"
	case MSTSCHSASec:
		return "MS-TSCH: Task Scheduler Service Remoting Protocol: SASec"
	case MSTSGUTsProxyRpcInterface:
		return "MS-TSGU: Terminal Services Gateway Server Protocol: TsProxyRpcInterface"
	case MSTSRAPIManageTelnetSessions:
		return "MS-TSRAP: Telnet Server Remote Administration Protocol: IManageTelnetSessions"
	case MSTSTSIcaApi:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: IcaApi"
	case MSTSTSRCMListener:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: RCMListener"
	case MSTSTSRCMPublic:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: RCMPublic"
	case MSTSTSSessEnvPublicRpc:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: SessEnvPublicRpc"
	case MSTSTSTermSrvEnumeration:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: TermSrvEnumeration"
	case MSTSTSTermSrvNotification:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: TermSrvNotification"
	case MSTSTSTermSrvSession:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: TermSrvSession"
	case MSTSTSTSVIPPublic:
		return "MS-TSTS: Terminal Services Terminal Server Runtime Interface Protocol: TSVIPPublic"
	case MSUAMGIAutomaticUpdates2:
		return "MS-UAMG: Update Agent Management Protocol: IAutomaticUpdates2"
	case MSUAMGIAutomaticUpdatesResults:
		return "MS-UAMG: Update Agent Management Protocol: IAutomaticUpdatesResults"
	case MSUAMGIAutomaticUpdates:
		return "MS-UAMG: Update Agent Management Protocol: IAutomaticUpdates"
	case MSUAMGICategoryCollection:
		return "MS-UAMG: Update Agent Management Protocol: ICategoryCollection"
	case MSUAMGICategory:
		return "MS-UAMG: Update Agent Management Protocol: ICategory"
	case MSUAMGIImageInformation:
		return "MS-UAMG: Update Agent Management Protocol: IImageInformation"
	case MSUAMGIInstallationBehavior:
		return "MS-UAMG: Update Agent Management Protocol: IInstallationBehavior"
	case MSUAMGISearchJob_:
		return "MS-UAMG: Update Agent Management Protocol: ISearchJob_"
	case MSUAMGISearchResult:
		return "MS-UAMG: Update Agent Management Protocol: ISearchResult"
	case MSUAMGIStringCollection:
		return "MS-UAMG: Update Agent Management Protocol: IStringCollection"
	case MSUAMGIUpdate2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdate2"
	case MSUAMGIUpdate3:
		return "MS-UAMG: Update Agent Management Protocol: IUpdate3"
	case MSUAMGIUpdate4:
		return "MS-UAMG: Update Agent Management Protocol: IUpdate4"
	case MSUAMGIUpdate5:
		return "MS-UAMG: Update Agent Management Protocol: IUpdate5"
	case MSUAMGIUpdateCollection:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateCollection"
	case MSUAMGIUpdateDownloadContent2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateDownloadContent2"
	case MSUAMGIUpdateDownloadContentCollection:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateDownloadContentCollection"
	case MSUAMGIUpdateDownloadContent:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateDownloadContent"
	case MSUAMGIUpdateExceptionCollectionOLE:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateExceptionCollectionOLE"
	case MSUAMGIUpdateExceptionCollection:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateExceptionCollection"
	case MSUAMGIUpdateException:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateException"
	case MSUAMGIUpdateHistoryEntry2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateHistoryEntry2"
	case MSUAMGIUpdateHistoryEntryCollection:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateHistoryEntryCollection"
	case MSUAMGIUpdateHistoryEntry:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateHistoryEntry"
	case MSUAMGIUpdateIdentity:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateIdentity"
	case MSUAMGIUpdateSearcher2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSearcher2"
	case MSUAMGIUpdateSearcher3:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSearcher3"
	case MSUAMGIUpdateSearcher:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSearcher"
	case MSUAMGIUpdateService2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateService2"
	case MSUAMGIUpdateServiceCollection:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateServiceCollection"
	case MSUAMGIUpdateServiceManager2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateServiceManager2"
	case MSUAMGIUpdateServiceManager:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateServiceManager"
	case MSUAMGIUpdateServiceRegistration:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateServiceRegistration"
	case MSUAMGIUpdateService:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateService"
	case MSUAMGIUpdateSession2:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSession2"
	case MSUAMGIUpdateSession3:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSession3"
	case MSUAMGIUpdateSession:
		return "MS-UAMG: Update Agent Management Protocol: IUpdateSession"
	case MSUAMGIUpdate:
		return "MS-UAMG: Update Agent Management Protocol: IUpdate"
	case MSUAMGIWindowsDriverUpdate2:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdate2"
	case MSUAMGIWindowsDriverUpdate3:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdate3"
	case MSUAMGIWindowsDriverUpdate4:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdate4"
	case MSUAMGIWindowsDriverUpdate5:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdate5"
	case MSUAMGIWindowsDriverUpdateEntryCollection:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdateEntryCollection"
	case MSUAMGIWindowsDriverUpdateEntry:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdateEntry"
	case MSUAMGIWindowsDriverUpdate:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsDriverUpdate"
	case MSUAMGIWindowsUpdateAgentInfo:
		return "MS-UAMG: Update Agent Management Protocol: IWindowsUpdateAgentInfo"
	case MSVDSIEnumVdsObject:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IEnumVdsObject"
	case MSVDSIVdsAdvancedDisk2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsAdvancedDisk2"
	case MSVDSIVdsAdvancedDisk3:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsAdvancedDisk3"
	case MSVDSIVdsAdvancedDisk:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsAdvancedDisk"
	case MSVDSIVdsAdviseSink:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsAdviseSink"
	case MSVDSIVdsAsync:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsAsync"
	case MSVDSIVdsCreatePartitionEx:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsCreatePartitionEx"
	case MSVDSIVdsDisk2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDisk2"
	case MSVDSIVdsDisk3:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDisk3"
	case MSVDSIVdsDiskOnline:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDiskOnline"
	case MSVDSIVdsDiskPartitionMF2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDiskPartitionMF2"
	case MSVDSIVdsDiskPartitionMF:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDiskPartitionMF"
	case MSVDSIVdsDisk:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsDisk"
	case MSVDSIVdsHbaPort:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsHbaPort"
	case MSVDSIVdsHwProvider:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsHwProvider"
	case MSVDSIVdsIscsiInitiatorAdapter:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsIscsiInitiatorAdapter"
	case MSVDSIVdsIscsiInitiatorPortal:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsIscsiInitiatorPortal"
	case MSVDSIVdsOpenDisk:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsOpenDisk"
	case MSVDSIVdsPack2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsPack2"
	case MSVDSIVdsPack:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsPack"
	case MSVDSIVdsProvider:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsProvider"
	case MSVDSIVdsRemovable:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsRemovable"
	case MSVDSIVdsServiceHba:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceHba"
	case MSVDSIVdsServiceInitialization:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceInitialization"
	case MSVDSIVdsServiceIscsi:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceIscsi"
	case MSVDSIVdsServiceLoader:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceLoader"
	case MSVDSIVdsServiceSAN:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceSAN"
	case MSVDSIVdsServiceSw:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceSw"
	case MSVDSIVdsServiceUninstallDisk:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsServiceUninstallDisk"
	case MSVDSIVdsService:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsService"
	case MSVDSIVdsSubSystemImportTarget:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsSubSystemImportTarget"
	case MSVDSIVdsSwProvider:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsSwProvider"
	case MSVDSIVdsVDisk:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVDisk"
	case MSVDSIVdsVdProvider:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVdProvider"
	case MSVDSIVdsVolume2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolume2"
	case MSVDSIVdsVolumeMF2:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumeMF2"
	case MSVDSIVdsVolumeMF3:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumeMF3"
	case MSVDSIVdsVolumeMF:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumeMF"
	case MSVDSIVdsVolumeOnline:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumeOnline"
	case MSVDSIVdsVolumePlex:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumePlex"
	case MSVDSIVdsVolumeShrink:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolumeShrink"
	case MSVDSIVdsVolume:
		return "MS-VDS: Virtual Disk Service (VDS) Protocol: IVdsVolume"
	case MSW32TW32Time:
		return "MS-W32T: W32Time Remote Protocol: W32Time"
	case MSWCCEICertRequestD2:
		return "MS-WCCE: Windows Client Certificate Enrollment Protocol: ICertRequestD2"
	case MSWCCEICertRequestD:
		return "MS-WCCE: Windows Client Certificate Enrollment Protocol: ICertRequestD"
	case MSWDSCWdsRpcInterface:
		return "MS-WDSC: Windows Deployment Services Control Protocol: WdsRpcInterface"
	case MSWKSTWkssvc:
		return "MS-WKST: Workstation Service Remote Protocol: wkssvc"
	case MSWMIIEnumWbemClassObject:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IEnumWbemClassObject"
	case MSWMIIWbemBackupRestoreEx:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemBackupRestoreEx"
	case MSWMIIWbemBackupRestore:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemBackupRestore"
	case MSWMIIWbemClassObject:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemClassObject"
	case MSWMIIWbemContext:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemContext"
	case MSWMIIWbemFetchSmartEnum:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemFetchSmartEnum"
	case MSWMIIWbemLevel1Login:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemLevel1Login"
	case MSWMIIWbemLoginClientID:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemLoginClientID"
	case MSWMIIWbemLoginHelper:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemLoginHelper"
	case MSWMIIWbemObjectSink:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemObjectSink"
	case MSWMIIWbemRefreshingServices:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemRefreshingServices"
	case MSWMIIWbemRemoteRefresher:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemRemoteRefresher"
	case MSWMIIWbemServices:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemServices"
	case MSWMIIWbemWCOSmartEnum:
		return "MS-WMI: Windows Management Instrumentation Remote Protocol: IWbemWCOSmartEnum"
	case MSWSRMIResourceManager2:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IResourceManager2"
	case MSWSRMIResourceManager:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IResourceManager"
	case MSWSRMIWRMAccounting:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMAccounting"
	case MSWSRMIWRMCalendar:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMCalendar"
	case MSWSRMIWRMConfig:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMConfig"
	case MSWSRMIWRMMachineGroup:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMMachineGroup"
	case MSWSRMIWRMPolicy:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMPolicy"
	case MSWSRMIWRMProtocol:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMProtocol"
	case MSWSRMIWRMRemoteSessionMgmt:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMRemoteSessionMgmt"
	case MSWSRMIWRMResourceGroup:
		return "MS-WSRM: Windows System Resource Manager (WSRM) Protocol: IWRMResourceGroup"
	case MSDLLAdvapi32:
		return "MS-DLL: advapi32.dll: advapi32"
	case MSDLLAppidsvc:
		return "MS-DLL: appidsvc.dll: appidsvc"
	case MSDLLAppinfo:
		return "MS-DLL: appinfo.dll: Application Information Service: appinfo"
	case MSDLLIAisAxISInfo:
		return "MS-DLL: appinfo.dll: Application AxIS Information Service: IAisAxISInfo"
	case MSDLLIAisCOMInfo:
		return "MS-DLL: appinfo.dll: Application COM Information Service: IAisCOMInfo"
	case MSDLLIAisMSIInfo:
		return "MS-DLL: appinfo.dll: Application MSI Information Service: IAisMSIInfo"
	case MSDLLAppmgmts:
		return "MS-DLL: appmgmts.dll: Application Management service: appmgmts"
	case MSDLLAqueue:
		return "MS-DLL: aqueue.dll: aqueue"
	case MSDLLAudiodgrpc:
		return "MS-DLL: audiodg.exe: Windows Audio Device Graph: audiodgrpc"
	case MSDLLAudiosrv:
		return "MS-DLL: audiosrv.dll: Windows Audio Service audiosrv: audiosrv"
	case MSDLLAudiorpc:
		return "MS-DLL: audiosrv.dll: Windows Audio Service Audiorpc: audiorpc"
	case MSDLLWluirRpc:
		return "MS-DLL: authui.dll: Windows Authentication UI: WluirRpc"
	case MSDLLWluir1Rpc:
		return "MS-DLL: authui.dll: Windows Authentication UI: Wluir1Rpc"
	case MSDLLAuthmgr32Rpc:
		return "MS-DLL: autmgr32.exe: Windows Automation Manager: Authmgr32Rpc"
	case MSDLLBdesvc:
		return "MS-DLL: bdesvc.dll: BitLocker Drive Encryption Service: bdesvc"
	case MSDLLBferpc:
		return "MS-DLL: bfe.dll: Base Filtering Engine: bferpc"
	case MSDLLBthServRPCService:
		return "MS-DLL: bthserv.dll: Bluetooth Support Service: BthServRPCService"
	case MSDLLCertPropSvc:
		return "MS-DLL: certprop.dll: Certificate Propagation Service: CertPropSvc"
	case MSDLLClusSvcJoinVersion:
		return "MS-DLL: clussvc.exe: Microsoft JoinVersion Interface: ClusSvcJoinVersion"
	case MSDLLClusSvc:
		return "MS-DLL: clussvc.exe: Microsoft Cluster Server Service: ClusSvc"
	case MSDLLClusSvcExtrocluster:
		return "MS-DLL: clussvc.exe: Microsoft Extrocluster Interface: ClusSvcExtrocluster"
	case MSDLLCryptCertProtectSvc:
		return "MS-DLL: cryptsvc.dll: Cryptographic Services: CryptCertProtectSvc"
	case MSDLLCryptKSrSvc:
		return "MS-DLL: cryptsvc.dll: Cryptographic Services: CryptKSrSvc"
	case MSDLLCryptKeyrSvc:
		return "MS-DLL: cryptsvc.dll: Cryptographic Services: CryptKeyrSvc"
	case MSDLLCryptCatSvc:
		return "MS-DLL: cryptsvc.dll: Cryptographic Services: CryptCatSvc"
	case MSDLLDhcpcsvc6:
		return "MS-DLL: dhcpcsvc6.dll: DHCPv6 Client Service: dhcpcsvc6"
	case MSDLLDhcpcsvc:
		return "MS-DLL: dhcpcsvc.dll: DHCP Client Service: dhcpcsvc"
	case MSDLLDmadminrpc:
		return "MS-DLL: dmadmin.exe: Logical Disk Manager service process: dmadminrpc"
	case MSDLLDnsResolver:
		return "MS-DLL: dnsrslvr.dll: DNS Client Service (Windows XP): DnsResolver"
	case MSDLLDnsResolver2000:
		return "MS-DLL: dnsrslvr.dll: DNS Client Service (Windows 2000): DnsResolver2000"
	case MSDLLWinlan:
		return "MS-DLL: dot3svc.dll: Wired AutoConfig Service: winlan"
	case MSDLLEfskrpc:
		return "MS-DLL: efssvc.dll: Encrypting File System (EFS) Service: EFSK: efskrpc"
	case MSDLLEfsrpc:
		return "MS-DLL: efssvc.dll: Encrypting File System (EFS) Service: efsrpc"
	case MSDLLEmdService:
		return "MS-DLL: emdmgmt.dll: ReadyBoost Service: EmdService"
	case MSDLLGpsvc:
		return "MS-DLL: gpsvc.dll: Group Policy Service: gpsvc"
	case MSDLLIasrpc:
		return "MS-DLL: ias.dll: Internet Authentication Service: iasrpc"
	case MSDLLIcardagtrpc:
		return "MS-DLL: icardagt.exe: Windows CardSpace User Interface Agent: icardagtrpc"
	case MSDLLIertutil:
		return "MS-DLL: iertutil.dll: Microsoft Internet Explorer Runtime Utilities: iertutil"
	case MSDLLIdletask:
		return "MS-DLL: schedsvc.dll: Task Scheduler Engine/Task Scheduler Service: Taskplaner idletask: idletask"
	case MSDLLISecureDesktop:
		return "MS-DLL: winlogon.exe: Windows Logon Application: ISecureDesktop"
	case MSDLLWMsgKAPIs:
		return "MS-DLL: winlogon.exe: Windows Logon Application: WMsgKAPIs"
	case MSDLLWMsgAPIs:
		return "MS-DLL: winlogon.exe: Windows Logon Application: WMsgAPIs"
	case MSDLLIRPCSCLogon:
		return "MS-DLL: winlogon.exe: Windows Logon Application: IRPCSCLogon"
	case MSDLLNrpsrv:
		return "MS-DLL: nrpsrv.dll: Name Resolution Proxy (NRP) RPC interface: nrpsrv"
	case MSDLLWinhttprpc:
		return "MS-DLL: winhttp.dll: Windows HTTP Services: winhttprpc"
	case MSDLLIRPCAsyncNotifyChannel:
		return "MS-DLL: spoolsv.exe: Spooler SubSystem App: IRPCAsyncNotifyChannel"
	case MSDLLIpTransition:
		return "MS-DLL: iphlpsvc.dll: Service that offers IPv6 connectivity over an IPv4 network: IpTransition"
	case MSDLLWinNsi:
		return "MS-DLL: nsisvc.dll: Network Store Interface RPC server: WinNsi"
	case MSDLLXactSrvRPC:
		return "MS-DLL: srvsvc.dll: Server Service: XactSrvRPC"
	case MSDLLIkeRpcIKE:
		return "MS-DLL: ikeext.dll: IKE extension: IkeRpcIKE"
	case MSDLLIRPCWinlogonNotifications:
		return "MS-DLL: webcint.dll: Web DAV Service: IRPCWinlogonNotifications"
	case MSDLLIfssvc:
		return "MS-DLL: ifssvc.exe: FS Web Agent Authentication Service: ifssvc"
	case MSDLLImepadsm:
		return "MS-DLL: IMEPADSM.DLL: Microsoft IME 2012: imepadsm"
	case MSDLLImepadsm1:
		return "MS-DLL: IMEPADSM.DLL: Microsoft IME 2012: imepadsm1"
	case MSDLLImjpdctrpc:
		return "MS-DLL: IMJPDCT.EXE: Microsoft IME: imjpdctrpc"
	case MSDLLCpprovider:
		return "MS-DLL: IPHLPAPI.DLL: Connectivity Helper API: cpprovider"
	case MSDLLIphlpsvc:
		return "MS-DLL: iphlpsvc.dll: IP Helper: iphlpsvc"
	case MSDLLIpnathlp:
		return "MS-DLL: ipnathlp.dll: Microsoft NAT Helper Components: ipnathlp"
	case MSDLLIpNatHlpRpcServer:
		return "MS-DLL: ipnathlp.dll: Microsoft NAT Helper Components: IpNatHlpRpcServer"
	case MSDLLIrftprpc:
		return "MS-DLL: irftp.exe: Wireless Link File Transfer Application: irftprpc"
	case MSDLLIrmon:
		return "MS-DLL: irmon.dll: Infrared Monitor: irmon"
	case MSDLLIrmon1:
		return "MS-DLL: irmon.dll: Infrared Monitor: irmon1"
	case MSDLLIscsiexe:
		return "MS-DLL: iscsiexe.dll: iSCSI Discovery service: iscsiexe"
	case MSDLLIsmserv_ip:
		return "MS-DLL: ismip.dll: Active Directory ISM IP Transport: ismserv_ip"
	case MSDLLIsmapi:
		return "MS-DLL: ismserv.exe: Inter-site Messaging service: ismapi"
	case MSDLLKdcrpc:
		return "MS-DLL: kdcsvc.dll: KDC RPC: kdcrpc"
	case MSDLLINCryptKeyIso:
		return "MS-DLL: keyiso.dll: CNG Key Isolation Service: INCryptKeyIso"
	case MSDLLLbservice:
		return "MS-DLL: LBService.dll: RPC/HTTP Load Balancing Coordinator: lbservice"
	case MSDLLLlsrpc:
		return "MS-DLL: llssrv.exe: Licensing Logging Service: llsrpc"
	case MSDLLLls_license:
		return "MS-DLL: llssrv.exe: License Logging Service: lls_license"
	case MSDLLNsiC:
		return "MS-DLL: locator.exe: RPC locator service: Network Store Interface Client: NsiC"
	case MSDLLNsiS:
		return "MS-DLL: locator.exe: RPC locator service: Network Store Interface Service: NsiS"
	case MSDLLNsiM:
		return "MS-DLL: locator.exe: RPC locator service: Network Store Interface Manager: NsiM"
	case MSDLLDsrolesvc:
		return "MS-DLL: lsasrv.dll: Microsoft Active Directory DSROLE Service: dsrolesvc"
	case MSDLLICryptProtect:
		return "MS-DLL: lsasrv.dll: Protected Storage Service: ICryptProtect"
	case MSDLLPasswordRecovery:
		return "MS-DLL: lsasrv.dll: Protected Storage Service: PasswordRecovery"
	case MSDLLSLspPrivateData:
		return "MS-DLL: lsasrv.dll: Protected Storage Service: SLspPrivateData"
	case MSDLLLsass:
		return "MS-DLL: lsass.exe: Local Security Authority Subsystem Service: lsass"
	case MSDLLTermServLicensing:
		return "MS-DLL: lserver.dll: Terminal Server License Server: TermServLicensing"
	case MSDLLHydraLsPipe:
		return "MS-DLL: lserver.dll: Terminal Server License Server: HydraLsPipe"
	case MSDLLTermSrvPrivate:
		return "MS-DLL: lsm.dll: Terminal Services Terminal Server Runtime Interface: TermSrvPrivate"
	case MSDLLTermSrvAdmin:
		return "MS-DLL: lsm.dll: Terminal Services Terminal Server Runtime Interface: TermSrvAdmin"
	case MSDLLMilcore:
		return "MS-DLL: milcore.dll: MIL Core Library: milcore"
	case MSDLLMpnotifyrpc:
		return "MS-DLL: mpnotify.exe: Windows NT Multiple Provider Notification Application: mpnotifyrpc"
	case MSDLLACPBackgroundManagerPolicy:
		return "MS-DLL: ACPBackgroundManagerPolicy.dll: ACP Background Manager Policy: ACPBackgroundManagerPolicy"
	case MSDLLAdhsvc:
		return "MS-DLL: adhsvc.dll: Annotation Proxy Manager client server: adhsvc"
	case MSDLLAphostservice:
		return "MS-DLL: aphostservice.dll: Accounts Host Service: aphostservice"
	case MSDLLAppinfopkg:
		return "MS-DLL: appinfo.dll: Application Information Service: appinfopkg"
	case MSDLLAudiosrvpbm:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: PBM: audiosrvpbm"
	case MSDLLAudiosrvpbm1:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: PBM: audiosrvpbm1"
	case MSDLLAudiosrvhrtf:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: HRTF Info: audiosrvhrtf"
	case MSDLLAudiosrvwnf:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: Device Graph WNF State: audiosrvwnf"
	case MSDLLAudiosrv1:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: audiosrv1"
	case MSDLLAudiosrv2:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: audiosrv2"
	case MSDLLAudiosrv3:
		return "MS-DLL: audiosrv.dll: Windows Audio Service: audiosrv3"
	case MSDLLRBiPtSrv:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: RBiPtSrv"
	case MSDLLRBiRtSrv:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: RBiRtSrv"
	case MSDLLSrvOdbPt:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: SrvOdbPt"
	case MSDLLSrvOdb:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: SrvOdb"
	case MSDLLRBiSrv:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: RBiSrv"
	case MSDLLSrvOdbPriv:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: SrvOdbPriv"
	case MSDLLRBiSrv1:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: RBiSrv1"
	case MSDLLSrvOdbLb:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: SrvOdbLb"
	case MSDLLRBiSrvSignal:
		return "MS-DLL: bisrv.dll: Background Tasks Infrastructure Service: RBiSrvSignal"
	case MSDLLBri:
		return "MS-DLL: BrokerLib.dll: Broker Base Library: Bri"
	case MSDLLIChakraJIT:
		return "MS-DLL: chakra.dll: ChakraCore Just-in-time (JIT) Compilation: IChakraJIT"
	case MSDLLCoremessaging:
		return "MS-DLL: coremessaging.dll: Microsoft CoreMessaging: coremessaging"
	case MSDLLCrypttpmeksvc:
		return "MS-DLL: crypttpmeksvc.dll: Cryptographic TPM Endorsement Key Services: crypttpmeksvc"
	case MSDLLWarpJITSvc:
		return "MS-DLL: d3d10warp.dll: Direct3D 10 Rasterizer: WarpJITSvc"
	case MSDLLDab:
		return "MS-DLL: DAB.dll: Desktop Activity Broker: dab"
	case MSDLLDaschallenge:
		return "MS-DLL: das.dll: Device Association Service: Challenge: daschallenge"
	case MSDLLDasaepstore:
		return "MS-DLL: das.dll: Device Association Service: AEP Store: dasaepstore"
	case MSDLLDas:
		return "MS-DLL: das.dll: Device Association Service: das"
	case MSDLLDasinbassoc:
		return "MS-DLL: das.dll: Device Association Service: Inbound Association: dasinbassoc"
	case MSDLLDasrpcdev:
		return "MS-DLL: das.dll: Device Association Service: RPC Dev: dasrpcdev"
	case MSDLLDashost:
		return "MS-DLL: dashost.exe: Device Association Framework Provider Host: dashost"
	case MSDLLUtcApi:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcApi"
	case MSDLLUtcTelemetryOptInApi:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcTelemetryOptInApi"
	case MSDLLUtcWerHelperApi:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcWerHelperApi"
	case MSDLLUtcEventTranscriptApi:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcEventTranscriptApi"
	case MSDLLUtcTenantApi:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcTenantApi"
	case MSDLLUtcApi1:
		return "MS-DLL: diagtrack.dll: Microsoft Windows Diagnostics Tracking: UtcApi1"
	case MSDLLDpapisrv:
		return "MS-DLL: dpapisrv.dll: DPAPI Server: dpapisrv"
	case MSDLLDssvc:
		return "MS-DLL: dssvc.dll: Data Sharing Service NT Service: dssvc"
	case MSDLLDusmsvc:
		return "MS-DLL: dusmsvc.dll: Data Usage (DusmSvc) Service: dusmsvc"
	case MSDLLEeprov:
		return "MS-DLL: eeprov.dll: Energy Estimator SRUM provider: eeprov"
	case MSDLLPsmSrvAppHost:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvAppHost"
	case MSDLLPsmSrvApp:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvApp"
	case MSDLLPsmSrvInfo:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvInfo"
	case MSDLLPsmSrvActivation:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvActivation"
	case MSDLLPsmSrvTimer:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvTimer"
	case MSDLLPsmSrvQuiesce:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrvQuiesce"
	case MSDLLPsmSrv:
		return "MS-DLL: psmsrv.dll: Process State Manager (PSM) Service: PsmSrv"
	case MSDLLWscsvc:
		return "MS-DLL: wscsvc.dll: Windows Security Center Service: wscsvc"
	case MSDLLPcaSvc:
		return "MS-DLL: pcasvc.dll: Program Compatibility Assistant Service: PcaSvc"
	case MSDLLRmgrSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: RmgrSrv"
	case MSDLLRmtSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: RmtSrv"
	case MSDLLHamRpcSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrv"
	case MSDLLHamRpcSrvSess:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvSess"
	case MSDLLHamRpcSrvHostId:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvHostId"
	case MSDLLRmGameModeRSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: RmGameModeRSrv"
	case MSDLLRmCoreRpcSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: RmCoreRpcSrv"
	case MSDLLHamRpcSrvState:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvState"
	case MSDLLHamRpcSrvFullTrust:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvFullTrust"
	case MSDLLHamRpcSrvServicing:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvServicing"
	case MSDLLCrmRpcSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: CrmRpcSrv"
	case MSDLLHamRpcSrvActivity:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvActivity"
	case MSDLLHamRpcSrvDebug:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: HamRpcSrvDebug"
	case MSDLLRmGameModeSrv:
		return "MS-DLL: psmserviceexthost.dll: Resource Manager PSM Service Extension: RmGameModeSrv"
	case MSDLLActiveSyncServer:
		return "MS-DLL: SyncController.dll: SyncController for managing sync of mail, contacts, calendar: ActiveSyncServer"
	case MSDLLAccountsMgmtRpc:
		return "MS-DLL: SyncController.dll: SyncController for managing sync of mail, contacts, calendar: AccountsMgmtRpc"
	case MSDLLIUserMarshalData:
		return "MS-DLL: usermgr.dll: Runtime components required for multi-user interaction: IUserMarshalData"
	case MSDLLUsermgr:
		return "MS-DLL: usermgr.dll: Runtime components required for multi-user interaction: usermgr"
	case MSDLLFmMuxSrvNotificationCoreUI:
		return "MS-DLL: modernexecserver.dll: Modern Execution Server: FmMuxSrvNotificationCoreUI"
	case MSDLLFmMuxSrvCoreUI:
		return "MS-DLL: modernexecserver.dll: Modern Execution Server: FmMuxSrvCoreUI"
	case MSDLLFmMuxSrv:
		return "MS-DLL: modernexecserver.dll: Modern Execution Server: FmMuxSrv"
	case MSDLLIdSegSrv:
		return "MS-DLL: srvsvc.dll: Server Service: IdSegSrv"
	case MSDLLTSSDFarmRpc:
		return "MS-DLL: sessenv.dll: Remote Desktop Configuration service: TSSDFarmRpc"
	case MSDLLTSSDRpc:
		return "MS-DLL: sessenv.dll: Remote Desktop Configuration service: TSSDRpc"
	case MSDLLNgcHandlerRpc:
		return "MS-DLL: NgcCtnrSvc.dll: Microsoft Passport Container Service: NgcHandlerRpc"
	case MSDLLProxyMgr:
		return "MS-DLL: httpprxm.dll: Proxy Manager: ProxyMgr"
	case MSDLLRpcSrvProxyMgr:
		return "MS-DLL: httpprxm.dll: Proxy Manager: RpcSrvProxyMgr"
	case MSDLLTaskScheduler:
		return "MS-DLL: WPTaskScheduler.dll: WP Task Scheduler: TaskScheduler"
	case MSDLLPhoneRpc:
		return "MS-DLL: phoneservice.dll: phone calls and other telephony related functionality: PhoneRpc"
	case MSDLLPhoneRpc1:
		return "MS-DLL: phoneservice.dll: phone calls and other telephony related functionality: PhoneRpc1"
	case MSDLLPhoneRpc2:
		return "MS-DLL: phoneservice.dll: phone calls and other telephony related functionality: PhoneRpc2"
	case MSDLLPhoneRpc3:
		return "MS-DLL: phoneservice.dll: phone calls and other telephony related functionality: PhoneRpc3"
	case MSDLLPhoneRpc4:
		return "MS-DLL: phoneservice.dll: phone calls and other telephony related functionality: PhoneRpc4"
	case MSDLLI_pSchRpc:
		return "MS-DLL: schedsvc.dll: Task Scheduler Service: I_pSchRpc"
	case MSDLLPimIMService:
		return "MS-DLL: pimindexmaintenance.dll: contacts indexing and other user data related tasks: PimIMService"
	case MSDLLSsdpsrv:
		return "MS-DLL: ssdpsrv.dll: SSDP service: ssdpsrv"
	case MSDLLWinBioCredentialManager:
		return "MS-DLL: wbiosrvc.dll: Windows Biometric Service: WinBioCredentialManager"
	case MSDLLWinBioServer:
		return "MS-DLL: wbiosrvc.dll: Windows Biometric Service: WinBioServer"
	case MSDLLColorManagementRpcServer:
		return "MS-DLL: DispBroker.Desktop.dll: Display Policy Service: ColorManagementRpcServer"
	case MSDLLNlasvc:
		return "MS-DLL: nlasvc.dll: Network Location Awareness 2: nlasvc"
	case MSDLLTokenBindingApi:
		return "MS-DLL: keyiso.dll: CNG Key Isolation Service: TokenBindingApi"
	case MSDLLSymmetricPopKeyApi:
		return "MS-DLL: keyiso.dll: CNG Key Isolation Service: SymmetricPopKeyApi"
	case MSDLLNcbRpcSrv:
		return "MS-DLL: ncbservice.dll: Network Connection Broker: NcbRpcSrv"
	case MSDLLNcbKapi:
		return "MS-DLL: ncbservice.dll: Network Connection Broker: NcbKapi"
	case MSDLLNcbRpcSrv1:
		return "MS-DLL: ncbservice.dll: Network Connection Broker: NcbRpcSrv1"
	case MSDLLStorSvc:
		return "MS-DLL: storsvc.dll: Storage Service: StorSvc"
	case MSDLLStorSvc1:
		return "MS-DLL: storsvc.dll: Storage Service: StorSvc1"
	case MSDLLGcsSrv:
		return "MS-DLL: resourcepolicyserver.dll: Resource Policy RM Service Extension: GcsSrv"
	case MSDLLClipSvc:
		return "MS-DLL: ClipSVC.dll: Client License Service: ClipSvc"
	case MSDLLRpSrv:
		return "MS-DLL: resourcepolicyserver.dll: Resource Policy RM Service Extension: RpSrv"
	case MSDLLRpSrv1:
		return "MS-DLL: resourcepolicyserver.dll: Resource Policy RM Service Extension: RpSrv1"
	case MSDLLRasCustomRpc:
		return "MS-DLL: rascustom.dll: Custom Protocol Engine: RasCustomRpc"
	case MSDLLCSebiPublisher:
		return "MS-DLL: systemeventsbrokerserver.dll: System Events Broker: CSebiPublisher"
	case MSDLLNgcSecureBioHandlerRpc:
		return "MS-DLL: NgcCtnrSvc.dll: Microsoft Passport Container Service: NgcSecureBioHandlerRpc"
	case MSDLLCSebiEvent:
		return "MS-DLL: systemeventsbrokerserver.dll: System Events Broker: CSebiEvent"
	case MSDLLSebiEvent:
		return "MS-DLL: systemeventsbrokerserver.dll: System Events Broker: SebiEvent"
	case MSDLLSgrmBroker:
		return "MS-DLL: SgrmBroker.exe: System Guard Runtime Monitor Broker Service: SgrmBroker"
	case MSDLLSetOperatorRpc:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: SetOperatorRpc"
	case MSDLLWcmSvcRpc:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: WcmSvcRpc"
	case MSDLLWcmSvcRpc1:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: WcmSvcRpc1"
	case MSDLLWcmSvcProxyInfoRpc:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: WcmSvcProxyInfoRpc"
	case MSDLLWcmSvcMgrRpc:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: WcmSvcMgrRpc"
	case MSDLLWcmSvcCapRpc:
		return "MS-DLL: wcmsvc.dll: Windows Connection Manager Service: WcmSvcCapRpc"
	case MSDLLDfsDsSvc:
		return "MS-DLL: wkssvc.dll: Workstation Service: DfsDsSvc"
	case MSDLLFwNetIsolationRpc:
		return "MS-DLL: mpssvc.dll: Microsoft Protection Service: FwNetIsolationRpc"
	case MSDLLFwRpc1:
		return "MS-DLL: MPSSVC.dll: Microsoft Protection Service: FwRpc1"
	case MSDLLFwRpc:
		return "MS-DLL: MPSSVC.dll: Microsoft Protection Service: FwRpc"
	case MSDLLNetSetupShim:
		return "MS-DLL: NetSetupShim.dll: Network Configuration API: NetSetupShim"
	case MSDLLNetSetupSvc:
		return "MS-DLL: NetSetupSvc.dll: Network Setup Service: NetSetupSvc"
	case MSDLLDeviceCredentialMgrRpc:
		return "MS-DLL: ngcsvc.dll: Microsoft Passport Service: DeviceCredentialMgrRpc"
	case MSDLLVscRpc:
		return "MS-DLL: NgcCtnrSvc.dll: Microsoft Passport Container Service: VscRpc"
	case MSDLLNgcLocalAccountVaultRpc:
		return "MS-DLL: ngcsvc.dll: Microsoft Passport Service: NgcLocalAccountVaultRpc"
	case MSDLLPhoneProviders:
		return "MS-DLL: PhoneProviders.dll: Phonespecific Component Provider for Windows Telephony Stack: PhoneProviders"
	case MSDLLLicenseManagerSvc:
		return "MS-DLL: licensemanagersvc.dll: Windows License Manager: LicenseManagerSvc"
	case MSDLLTbiRpc:
		return "MS-DLL: timebrokerserver.dll: Time Event Broker: TbiRpc"
	case MSDLLDeManagementRpcServer:
		return "MS-DLL: Microsoft.Graphics.Display.DisplayEnhancementService.dll: Microsoft Graphics Display Enhancement Service: DeManagementRpcServer"
	case MSDLLDeoRpcServer:
		return "MS-DLL: Microsoft.Graphics.Display.DisplayEnhancementService.dll: Microsoft Graphics Display Enhancement Service: DeoRpcServer"
	case MSDLLPfRpcServerSuperfetch:
		return "MS-DLL: sysmain.dll: Superfetch Service Host: PfRpcServerSuperfetch"
	case MSDLLWcnpRpc:
		return "MS-DLL: wcncsvc.dll: Windows Connect Now - Config Registrar Service: WcnpRpc"
	case MSDLLWcnTransportRpc:
		return "MS-DLL: wcncsvc.dll: Windows Connect Now - Config Registrar Service: WcnTransportRpc"
	case MSDLLNgcMgmt:
		return "MS-DLL: ngcsvc.dll: Microsoft Passport Service: NgcMgmt"
	case MSDLLUdmSvc:
		return "MS-DLL: userdataservice.dll: The endpoint for 3rd party APIs to read/write user data: UdmSvc"
	case MSDLLUtcAdminApi:
		return "MS-DLL: tellib.dll: Microsoft Telemetry Library: UtcAdminApi"
	case MSDLLNgcMgmt1:
		return "MS-DLL: ngcsvc.dll: Microsoft Passport Service: NgcMgmt1"
	case MSDLLGidsHandlerRpc:
		return "MS-DLL: NgcCtnrGidsHandler.dll: Microsoft Passport Container GIDS Handler: GidsHandlerRpc"
	case MSDLLHcsRpc:
		return "MS-DLL: vmcompute.exe: VM Compute: HcsRpc"
	case MSDLLUSSvc:
		return "MS-DLL: unistore.dll: Unified Store: USSvc"
	case MSDLLClusterConnRpc:
		return "MS-DLL: wkssvc.dll: Workstation Service: ClusterConnRpc"
	case MSDLLWitnessNodesRpc:
		return "MS-DLL: wkssvc.dll: Workstation Service: WitnessNodesRpc"
	case MSDLLLogonWluirRpc:
		return "MS-DLL: logoncontroller.dll: Logon UX Controller: LogonWluirRpc"
	case MSDLLDeviceCredentialRpc:
		return "MS-DLL: ngcsvc.dll: Microsoft Passport Service: DeviceCredentialRpc"
	case MSDLLNgcRpc:
		return "MS-DLL: NgcCtnrSvc.dll: Microsoft Passport Container Service: NgcRpc"
	case MSDLLSECOMNRpc:
		return "MS-DLL: SECOMN.DLL: Sound Research Corp: SECOMNRpc"
	case MSDLLIhvmsmapi:
		return "MS-DLL: wlanmsm.dll: Windows Wireless LAN 802.11 MSM: ihvmsmapi"
	case MSDLLWiFiUserRpc:
		return "MS-DLL: wifinetworkmanager.dll: Wireless Network Manager: WiFiUserRpc"
	case MSDLLWlanSvcLowPriv:
		return "MS-DLL: wlansvc.dll: Windows WLAN AutoConfig Service: WlanSvcLowPriv"
	case MSDLLWlanSvc:
		return "MS-DLL: wlansvc.dll: Windows WLAN AutoConfig Service: WlanSvc"
	case MSDLLWDiagRpc:
		return "MS-DLL: wlansvc.dll: Windows WLAN AutoConfig Service: WDiagRpc"
	case MSDLLRpcDot11ExtIhv:
		return "MS-DLL: wlanext.exe: Windows Wireless LAN 802.11 Extensibility Framework: RpcDot11ExtIhv"
	case MSDLLSignalRpc:
		return "MS-DLL: PackageStateRoaming.dll: Package State Roaming: SignalRpc"
	case MSDLLLiveIdSvcNotifyRpc:
		return "MS-DLL: wlidsvc.dll: Microsoft Account Service : LiveIdSvcNotifyRpc"
	case MSDLLLiveIdSvcRpc:
		return "MS-DLL: wlidsvc.dll: Microsoft Account Service : LiveIdSvcRpc"
	case MSDLLOnlineProviderCertRpc:
		return "MS-DLL: wlidsvc.dll: Microsoft Account Service : OnlineProviderCertRpc"
	case MSDLLUnimdmRpc:
		return "MS-DLL: unimdm.tsp: Universal Modem Driver TSP: UnimdmRpc"
	case MSDLLColorAdapterRpcServer:
		return "MS-DLL: Microsoft.Graphics.Display.DisplayEnhancementService.dll: Microsoft Graphics Display Enhancement Service: ColorAdapterRpcServer"
	case MSDLLSCardCacheRpc:
		return "MS-DLL: SCardSvr.dll: Smart Card Resource Management Server: SCardCacheRpc"
	case MSDLLSCardRpc:
		return "MS-DLL: SCardSvr.dll: Smart Card Resource Management Server: SCardRpc"
	case MSDLLBthEvtBrRpc:
		return "MS-DLL: Windows.Internal.Bluetooth.dll: Windows Internal Bluetooth: BthEvtBrRpc"
	case MSDLLRpcBluetoothGatt:
		return "MS-DLL: Microsoft.Bluetooth.Service.dll: Microsoft Bluetooth Service: RpcBluetoothGatt"
	case MSDLLAppXSetTrustLabelRpc:
		return "MS-DLL: AppXDeploymentServer.dll: AppX Deployment Server: AppXSetTrustLabelRpc"
	case MSDLLAppXSetTrustLabel1Rpc:
		return "MS-DLL: AppXDeploymentServer.dll: AppX Deployment Server: AppXSetTrustLabel1Rpc"
	case MSDLLVoIpRpc:
		return "MS-DLL: VoipRT.dll: VoIP Runtime: VoIpRpc"
	case MSDLLWFDSConMgrRpc:
		return "MS-DLL: WFDSConMgrSvc.dll: WiFi Direct Services Connection Manager Service: WFDSConMgrRpc"
	case MSDLLFrameServerRpc:
		return "MS-DLL: FrameServer.dll: Windows Camera Frame Server: FrameServerRpc"
	case MSDLLWwanRpc:
		return "MS-DLL: wwansvc.dll: WWAN Auto Config Service: WwanRpc"
	case MSDLLWwan2Rpc:
		return "MS-DLL: wwansvc.dll: WWAN Auto Config Service: Wwan2Rpc"
	case MSDLLWwanMgmtRpc:
		return "MS-DLL: wwansvc.dll: WWAN Auto Config Service: WwanMgmtRpc"
	case MSDLLIMpService:
		return "MS-DLL: mpsvc.dll: Windows Defender Service: IMpService"
	case RPCSSIMachineActivatorControl:
		return "RPCSS: Machine Activator Control: IMachineActivatorControl"
	case RPCSSISCMLocalActivator:
		return "RPCSS: The ISCM Local Activactor: ISCMLocalActivator"
	case RPCSSISCM:
		return "RPCSS: The ISCM RPC interface: ISCM"
	case RPCSSIROT:
		return "RPCSS: Running Object Table Interface: IROT"
	case RPCSSILocalObjectExporter:
		return "RPCSS: Local Object Exporter Service: ILocalObjectExporter"
	case RPCSSLocalepmp:
		return "RPCSS: Local Endpoint Mapper: localepmp"
	case RPCSSDbgidl:
		return "RPCSS: Dbgidl: Dbgidl"
	case RPCSSFwidl:
		return "RPCSS: Fwidl: Fwidl"
	}
	return ""
}

func (u UUID) WellKnownEndpoint() []string {
	switch (uuid.UUID)(u) {
	case MSBKRPBackupKey:
		return []string{"ncacn_np:protected_storage", "ncacn_np:ntsvcs"}
	case MSBRWSABrowser:
		return []string{"ncacn_np:browser"}
	case MSCAPRLsacap:
		return []string{"ncacn_np:lsarpc"}
	case MSDCOMIActivation:
		return []string{"ncacn_ip_tcp:135"}
	case MSDCOMIObjectExporter:
		return []string{"ncacn_ip_tcp:135"}
	case MSDCOMIRemoteSCMActivator:
		return []string{"ncacn_ip_tcp:135"}
	case MSDFSNMNetDFS:
		return []string{"ncacn_np:NETDFS"}
	case MSDHCPMDhcpsrv2:
		return []string{"ncacn_np:DHCPSERVER"}
	case MSDHCPMDhcpsrv:
		return []string{"ncacn_np:DHCPSERVER"}
	case MSDLTWTrkwks:
		return []string{"ncacn_np:ntsvcs", "ncacn_np:trkwks"}
	case MSDNSPDnsServer:
		return []string{"ncacn_np:DNSSERVER"}
	case MSDSSPDssetup:
		return []string{"ncacn_np:lsarpc"}
	case MSEFSREfsrpc:
		return []string{"ncacn_np:efsrpc"}
	case MSEFSRLsarpc:
		return []string{"ncacn_np:lsarpc"}
	case MSEVENEventlog:
		return []string{"ncacn_np:eventlog"}
	case MSFAXFax:
		return []string{"ncacn_np:SHAREDFAX"}
	case MSFSRVPFileServerVssAgent:
		return []string{"ncacn_np:FssagentRpc"}
	case MSICPRICertPassage:
		return []string{"ncacn_np:cert"}
	case MSIRPInetinfo:
		return []string{"ncacn_np:inetinfo"}
	case MSLSADLsarpc:
		return []string{"ncacn_np:lsarpc"}
	case MSLSATLsarpc:
		return []string{"ncacn_np:lsarpc"}
	case MSMSRPMsgsvc:
		return []string{"ncacn_np:MSGSVC"}
	case MSNRPCLogon:
		return []string{"ncacn_np:NETLOGON"}
	case MSPCQPerflibV2:
		return []string{"ncacn_np:winreg"}
	case MSRAIWWinsif2:
		return []string{"ncacn_np:WinsPipe"}
	case MSRAIWWinsif:
		return []string{"ncacn_np:WinsPipe"}
	case MSRPCEEndpointMapper:
		return []string{"ncacn_ip_tcp:135", "ncacn_np:epmapper"}
	case MSRPCERemoteManagement:
		return []string{"ncacn_ip_tcp:135"}
	case MSRPCLLocToLoc:
		return []string{"ncacn_np:Locator"}
	case MSRPRNWinspool:
		return []string{"ncacn_np:spoolss"}
	case MSRRASMRASRPC:
		return []string{"ncacn_np:ROUTER"}
	case MSRRPWinreg:
		return []string{"ncacn_np:winreg"}
	case MSRSPInitShutdown:
		return []string{"ncacn_np:InitShutdown"}
	case MSRSPWindowsShutdown:
		return []string{"ncacn_np:Shutdown"}
	case MSRSPWinreg:
		return []string{"ncacn_np:winreg"}
	case MSSAMRSamr:
		return []string{"ncacn_np:samr"}
	case MSSCMRSvcctl:
		return []string{"ncacn_np:svcctl"}
	case MSSRVSSrvsvc:
		return []string{"ncacn_np:srvsvc"}
	case MSTRPTapsrv:
		return []string{"ncacn_np:tapsrv"}
	case MSTSCHATSvc:
		return []string{"ncacn_np:atsvc"}
	case MSTSTSIcaApi:
		return []string{"ncacn_np:Ctx_WinStation_API_service"}
	case MSTSTSRCMListener:
		return []string{"ncacn_np:TermSrv_API_service"}
	case MSTSTSRCMPublic:
		return []string{"ncacn_np:TermSrv_API_service"}
	case MSTSTSSessEnvPublicRpc:
		return []string{"ncacn_np:SessEnvPublicRpc"}
	case MSTSTSTermSrvEnumeration:
		return []string{"ncacn_np:LSM_API_service", "ncacn_np:UNIFIED_API_service"}
	case MSTSTSTermSrvNotification:
		return []string{"ncacn_np:LSM_API_service", "ncacn_np:UNIFIED_API_service"}
	case MSTSTSTermSrvSession:
		return []string{"ncacn_np:LSM_API_service", "ncacn_np:UNIFIED_API_service"}
	case MSTSTSTSVIPPublic:
		return []string{"ncacn_np:TSVIP_Service"}
	case MSW32TW32Time:
		return []string{"ncacn_np:W32TIME", "ncacn_np:W32TIME_ALT"}
	case MSWKSTWkssvc:
		return []string{"ncacn_np:wkssvc", "ncacn_np:samr"}
	case MSDLLLlsrpc:
		return []string{"ncacn_np:llsrpc"}
	case MSDLLLls_license:
		return []string{"ncacn_np:llsrpc"}
	case MSDLLNsiC:
		return []string{"ncacn_np:locator"}
	case MSDLLNsiS:
		return []string{"ncacn_np:locator"}
	case MSDLLNsiM:
		return []string{"ncacn_np:locator"}
	case RPCSSIROT:
		return []string{"ncacn_ip_tcp:135"}
	case RPCSSILocalObjectExporter:
		return []string{"ncacn_ip_tcp:135"}
	}
	return nil
}
func (u UUID) Name() string {
	switch (uuid.UUID)(u) {
	case Null:
		return "Null"
	case MCCCFGIClusCfgAsyncEvictCleanup:
		return "IClusCfgAsyncEvictCleanup"
	case MCIISAIAppHostAdminManager:
		return "IAppHostAdminManager"
	case MCIISAIAppHostChangeHandler:
		return "IAppHostChangeHandler"
	case MCIISAIAppHostChildElementCollection:
		return "IAppHostChildElementCollection"
	case MCIISAIAppHostCollectionSchema:
		return "IAppHostCollectionSchema"
	case MCIISAIAppHostConfigException:
		return "IAppHostConfigException"
	case MCIISAIAppHostConfigFile:
		return "IAppHostConfigFile"
	case MCIISAIAppHostConfigLocationCollection:
		return "IAppHostConfigLocationCollection"
	case MCIISAIAppHostConfigLocation:
		return "IAppHostConfigLocation"
	case MCIISAIAppHostConfigManager:
		return "IAppHostConfigManager"
	case MCIISAIAppHostConstantValueCollection:
		return "IAppHostConstantValueCollection"
	case MCIISAIAppHostConstantValue:
		return "IAppHostConstantValue"
	case MCIISAIAppHostElementCollection:
		return "IAppHostElementCollection"
	case MCIISAIAppHostElementSchemaCollection:
		return "IAppHostElementSchemaCollection"
	case MCIISAIAppHostElementSchema:
		return "IAppHostElementSchema"
	case MCIISAIAppHostElement:
		return "IAppHostElement"
	case MCIISAIAppHostMappingExtension:
		return "IAppHostMappingExtension"
	case MCIISAIAppHostMethodCollection:
		return "IAppHostMethodCollection"
	case MCIISAIAppHostMethodInstance:
		return "IAppHostMethodInstance"
	case MCIISAIAppHostMethodSchema:
		return "IAppHostMethodSchema"
	case MCIISAIAppHostMethod:
		return "IAppHostMethod"
	case MCIISAIAppHostPathMapper:
		return "IAppHostPathMapper"
	case MCIISAIAppHostPropertyCollection:
		return "IAppHostPropertyCollection"
	case MCIISAIAppHostPropertyException:
		return "IAppHostPropertyException"
	case MCIISAIAppHostPropertySchemaCollection:
		return "IAppHostPropertySchemaCollection"
	case MCIISAIAppHostPropertySchema:
		return "IAppHostPropertySchema"
	case MCIISAIAppHostProperty:
		return "IAppHostProperty"
	case MCIISAIAppHostSectionDefinitionCollection:
		return "IAppHostSectionDefinitionCollection"
	case MCIISAIAppHostSectionDefinition:
		return "IAppHostSectionDefinition"
	case MCIISAIAppHostSectionGroup:
		return "IAppHostSectionGroup"
	case MCIISAIAppHostWritableAdminManager:
		return "IAppHostWritableAdminManager"
	case MCMQACIMSMQApplication2:
		return "IMSMQApplication2"
	case MCMQACIMSMQApplication3:
		return "IMSMQApplication3"
	case MCMQACIMSMQApplication:
		return "IMSMQApplication"
	case MCMQACIMSMQCollection:
		return "IMSMQCollection"
	case MCMQACIMSMQCoordinatedTransactionDispenser2:
		return "IMSMQCoordinatedTransactionDispenser2"
	case MCMQACIMSMQCoordinatedTransactionDispenser3:
		return "IMSMQCoordinatedTransactionDispenser3"
	case MCMQACIMSMQCoordinatedTransactionDispenser:
		return "IMSMQCoordinatedTransactionDispenser"
	case MCMQACIMSMQDestination:
		return "IMSMQDestination"
	case MCMQACIMSMQEvent2:
		return "IMSMQEvent2"
	case MCMQACIMSMQEvent3:
		return "IMSMQEvent3"
	case MCMQACIMSMQEvent:
		return "IMSMQEvent"
	case MCMQACIMSMQManagement:
		return "IMSMQManagement"
	case MCMQACIMSMQMessage2:
		return "IMSMQMessage2"
	case MCMQACIMSMQMessage3:
		return "IMSMQMessage3"
	case MCMQACIMSMQMessage4:
		return "IMSMQMessage4"
	case MCMQACIMSMQMessage:
		return "IMSMQMessage"
	case MCMQACIMSMQOutgoingQueueManagement:
		return "IMSMQOutgoingQueueManagement"
	case MCMQACIMSMQPrivateDestination:
		return "IMSMQPrivateDestination"
	case MCMQACIMSMQPrivateEvent:
		return "IMSMQPrivateEvent"
	case MCMQACIMSMQQuery2:
		return "IMSMQQuery2"
	case MCMQACIMSMQQuery3:
		return "IMSMQQuery3"
	case MCMQACIMSMQQuery4:
		return "IMSMQQuery4"
	case MCMQACIMSMQQuery:
		return "IMSMQQuery"
	case MCMQACIMSMQQueue2:
		return "IMSMQQueue2"
	case MCMQACIMSMQQueue3:
		return "IMSMQQueue3"
	case MCMQACIMSMQQueue4:
		return "IMSMQQueue4"
	case MCMQACIMSMQQueueInfo2:
		return "IMSMQQueueInfo2"
	case MCMQACIMSMQQueueInfo3:
		return "IMSMQQueueInfo3"
	case MCMQACIMSMQQueueInfo4:
		return "IMSMQQueueInfo4"
	case MCMQACIMSMQQueueInfos2:
		return "IMSMQQueueInfos2"
	case MCMQACIMSMQQueueInfos3:
		return "IMSMQQueueInfos3"
	case MCMQACIMSMQQueueInfos4:
		return "IMSMQQueueInfos4"
	case MCMQACIMSMQQueueInfos:
		return "IMSMQQueueInfos"
	case MCMQACIMSMQQueueInfo:
		return "IMSMQQueueInfo"
	case MCMQACIMSMQQueueManagement:
		return "IMSMQQueueManagement"
	case MCMQACIMSMQQueue:
		return "IMSMQQueue"
	case MCMQACIMSMQTransaction2:
		return "IMSMQTransaction2"
	case MCMQACIMSMQTransaction3:
		return "IMSMQTransaction3"
	case MCMQACIMSMQTransactionDispenser2:
		return "IMSMQTransactionDispenser2"
	case MCMQACIMSMQTransactionDispenser3:
		return "IMSMQTransactionDispenser3"
	case MCMQACIMSMQTransactionDispenser:
		return "IMSMQTransactionDispenser"
	case MCMQACIMSMQTransaction:
		return "IMSMQTransaction"
	case MCMQACITransaction:
		return "ITransaction"
	case MSADTGIDataFactory2:
		return "IDataFactory2"
	case MSADTGIDataFactory3:
		return "IDataFactory3"
	case MSADTGIDataFactory:
		return "IDataFactory"
	case MSADTSClaims:
		return "Claims"
	case MSBKRPBackupKey:
		return "BackupKey"
	case MSBPAUBitsPeerAuth:
		return "BitsPeerAuth"
	case MSBRWSABrowser:
		return "Browser"
	case MSCAPRLsacap:
		return "lsacap"
	case MSCMPOIXnRemote:
		return "IXnRemote"
	case MSCMRPClusAPI:
		return "ClusAPI"
	case MSCOMAIAlternateLaunch:
		return "IAlternateLaunch"
	case MSCOMAICapabilitySupport:
		return "ICapabilitySupport"
	case MSCOMAICatalog64BitSupport:
		return "ICatalog64BitSupport"
	case MSCOMAICatalogSession:
		return "ICatalogSession"
	case MSCOMAICatalogTableInfo:
		return "ICatalogTableInfo"
	case MSCOMAICatalogTableRead:
		return "ICatalogTableRead"
	case MSCOMAICatalogTableWrite:
		return "ICatalogTableWrite"
	case MSCOMAICatalogUtils2:
		return "ICatalogUtils2"
	case MSCOMAICatalogUtils:
		return "ICatalogUtils"
	case MSCOMAIContainerControl2:
		return "IContainerControl2"
	case MSCOMAIContainerControl:
		return "IContainerControl"
	case MSCOMAIExport2:
		return "IExport2"
	case MSCOMAIExport:
		return "IExport"
	case MSCOMAIImport2:
		return "IImport2"
	case MSCOMAIImport:
		return "IImport"
	case MSCOMAIRegister2:
		return "IRegister2"
	case MSCOMAIRegister:
		return "IRegister"
	case MSCOMAIReplicationUtil:
		return "IReplicationUtil"
	case MSCOMITransactionStream:
		return "ITransactionStream"
	case MSCOMEVIEnumEventObject:
		return "IEnumEventObject"
	case MSCOMEVIEventClass2:
		return "IEventClass2"
	case MSCOMEVIEventClass3:
		return "IEventClass3"
	case MSCOMEVIEventClass:
		return "IEventClass"
	case MSCOMEVIEventObjectCollection:
		return "IEventObjectCollection"
	case MSCOMEVIEventSubscription2:
		return "IEventSubscription2"
	case MSCOMEVIEventSubscription3:
		return "IEventSubscription3"
	case MSCOMEVIEventSubscription:
		return "IEventSubscription"
	case MSCOMEVIEventSystem2:
		return "IEventSystem2"
	case MSCOMEVIEventSystemInitialize:
		return "IEventSystemInitialize"
	case MSCOMEVIEventSystem:
		return "IEventSystem"
	case MSCOMTIComTrackingInfoEvents:
		return "IComTrackingInfoEvents"
	case MSCOMTIGetTrackingData:
		return "IGetTrackingData"
	case MSCOMTIProcessDump:
		return "IProcessDump"
	case MSCSRAICertAdminD2:
		return "ICertAdminD2"
	case MSCSRAICertAdminD:
		return "ICertAdminD"
	case MSCSVPIClusterCleanup:
		return "IClusterCleanup"
	case MSCSVPIClusterFirewall:
		return "IClusterFirewall"
	case MSCSVPIClusterLog:
		return "IClusterLog"
	case MSCSVPIClusterNetwork2:
		return "IClusterNetwork2"
	case MSCSVPIClusterSetup:
		return "IClusterSetup"
	case MSCSVPIClusterStorage2:
		return "IClusterStorage2"
	case MSCSVPIClusterStorage3:
		return "IClusterStorage3"
	case MSCSVPIClusterUpdate:
		return "IClusterUpdate"
	case MSDCOMIActivation:
		return "IActivation"
	case MSDCOMIObjectExporter:
		return "IObjectExporter"
	case MSDCOMIRemoteSCMActivator:
		return "IRemoteSCMActivator"
	case MSDCOMIRemUnknown2:
		return "IRemUnknown2"
	case MSDCOMIRemUnknown:
		return "IRemUnknown"
	case MSDCOMIUnknown:
		return "IUnknown"
	case MSDFSNMNetDFS:
		return "NetDFS"
	case MSDFSRHIADProxy2:
		return "IADProxy2"
	case MSDFSRHIADProxy:
		return "IADProxy"
	case MSDFSRHIServerHealthReport2:
		return "IServerHealthReport2"
	case MSDFSRHIServerHealthReport:
		return "IServerHealthReport"
	case MSDHCPMDhcpsrv2:
		return "dhcpsrv2"
	case MSDHCPMDhcpsrv:
		return "dhcpsrv"
	case MSDLTMTrksvr:
		return "trksvr"
	case MSDLTWTrkwks:
		return "trkwks"
	case MSDMRPIDMNotify:
		return "IDMNotify"
	case MSDMRPIDMRemoteServer:
		return "IDMRemoteServer"
	case MSDMRPIVolumeClient2:
		return "IVolumeClient2"
	case MSDMRPIVolumeClient3:
		return "IVolumeClient3"
	case MSDMRPIVolumeClient4:
		return "IVolumeClient4"
	case MSDMRPIVolumeClient:
		return "IVolumeClient"
	case MSDNSPDnsServer:
		return "DnsServer"
	case MSDRSRDrsuapi:
		return "drsuapi"
	case MSDRSRDsaop:
		return "dsaop"
	case MSDSSPDssetup:
		return "dssetup"
	case MSEFSREfsrpc:
		return "efsrpc"
	case MSEFSRLsarpc:
		return "lsarpc"
	case MSEVEN6Eventlog:
		return "Eventlog"
	case MSEVENEventlog:
		return "eventlog"
	case MSFASPRemoteFW:
		return "RemoteFW"
	case MSFAXFaxclient:
		return "faxclient"
	case MSFAXFax:
		return "fax"
	case MSFRS1Frsrpc:
		return "frsrpc"
	case MSFRS1NtFrsApi:
		return "NtFrsApi"
	case MSFRS2FrsTransport:
		return "FrsTransport"
	case MSFRS2DfsrEndpoint:
		return "DfsrEndpoint"
	case MSFSRMIFsrmActionCommand:
		return "IFsrmActionCommand"
	case MSFSRMIFsrmActionEmail2:
		return "IFsrmActionEmail2"
	case MSFSRMIFsrmActionEmail:
		return "IFsrmActionEmail"
	case MSFSRMIFsrmActionEventLog:
		return "IFsrmActionEventLog"
	case MSFSRMIFsrmActionReport:
		return "IFsrmActionReport"
	case MSFSRMIFsrmAction:
		return "IFsrmAction"
	case MSFSRMIFsrmAutoApplyQuota:
		return "IFsrmAutoApplyQuota"
	case MSFSRMIFsrmClassificationManager:
		return "IFsrmClassificationManager"
	case MSFSRMIFsrmClassificationRule:
		return "IFsrmClassificationRule"
	case MSFSRMIFsrmClassifierModuleDefinition:
		return "IFsrmClassifierModuleDefinition"
	case MSFSRMIFsrmCollection:
		return "IFsrmCollection"
	case MSFSRMIFsrmCommittableCollection:
		return "IFsrmCommittableCollection"
	case MSFSRMIFsrmDerivedObjectsResult:
		return "IFsrmDerivedObjectsResult"
	case MSFSRMIFsrmFileGroupImported:
		return "IFsrmFileGroupImported"
	case MSFSRMIFsrmFileGroupManager:
		return "IFsrmFileGroupManager"
	case MSFSRMIFsrmFileGroup:
		return "IFsrmFileGroup"
	case MSFSRMIFsrmFileManagementJobManager:
		return "IFsrmFileManagementJobManager"
	case MSFSRMIFsrmFileManagementJob:
		return "IFsrmFileManagementJob"
	case MSFSRMIFsrmFileScreenBase:
		return "IFsrmFileScreenBase"
	case MSFSRMIFsrmFileScreenException:
		return "IFsrmFileScreenException"
	case MSFSRMIFsrmFileScreenManager:
		return "IFsrmFileScreenManager"
	case MSFSRMIFsrmFileScreenTemplateImported:
		return "IFsrmFileScreenTemplateImported"
	case MSFSRMIFsrmFileScreenTemplateManager:
		return "IFsrmFileScreenTemplateManager"
	case MSFSRMIFsrmFileScreenTemplate:
		return "IFsrmFileScreenTemplate"
	case MSFSRMIFsrmFileScreen:
		return "IFsrmFileScreen"
	case MSFSRMIFsrmMutableCollection:
		return "IFsrmMutableCollection"
	case MSFSRMIFsrmObject:
		return "IFsrmObject"
	case MSFSRMIFsrmPathMapper:
		return "IFsrmPathMapper"
	case MSFSRMIFsrmPipelineModuleDefinition:
		return "IFsrmPipelineModuleDefinition"
	case MSFSRMIFsrmPropertyCondition:
		return "IFsrmPropertyCondition"
	case MSFSRMIFsrmPropertyDefinition2:
		return "IFsrmPropertyDefinition2"
	case MSFSRMIFsrmPropertyDefinition:
		return "IFsrmPropertyDefinition"
	case MSFSRMIFsrmPropertyDefinitionValue:
		return "IFsrmPropertyDefinitionValue"
	case MSFSRMIFsrmProperty:
		return "IFsrmProperty"
	case MSFSRMIFsrmQuotaBase:
		return "IFsrmQuotaBase"
	case MSFSRMIFsrmQuotaManagerEx:
		return "IFsrmQuotaManagerEx"
	case MSFSRMIFsrmQuotaManager:
		return "IFsrmQuotaManager"
	case MSFSRMIFsrmQuotaObject:
		return "IFsrmQuotaObject"
	case MSFSRMIFsrmQuotaTemplateImported:
		return "IFsrmQuotaTemplateImported"
	case MSFSRMIFsrmQuotaTemplateManager:
		return "IFsrmQuotaTemplateManager"
	case MSFSRMIFsrmQuotaTemplate:
		return "IFsrmQuotaTemplate"
	case MSFSRMIFsrmQuota:
		return "IFsrmQuota"
	case MSFSRMIFsrmReportJob:
		return "IFsrmReportJob"
	case MSFSRMIFsrmReportManager:
		return "IFsrmReportManager"
	case MSFSRMIFsrmReportScheduler:
		return "IFsrmReportScheduler"
	case MSFSRMIFsrmReport:
		return "IFsrmReport"
	case MSFSRMIFsrmRule:
		return "IFsrmRule"
	case MSFSRMIFsrmSetting:
		return "IFsrmSetting"
	case MSFSRMIFsrmStorageModuleDefinition:
		return "IFsrmStorageModuleDefinition"
	case MSFSRVPFileServerVssAgent:
		return "FileServerVssAgent"
	case MSGKDIISDKey:
		return "ISDKey"
	case MSICPRICertPassage:
		return "ICertPassage"
	case MSIISSIIisServiceControl:
		return "IIisServiceControl"
	case MSIMSAIIISApplicationAdmin:
		return "IIISApplicationAdmin"
	case MSIMSAIIISCertObj:
		return "IIISCertObj"
	case MSIMSAIMSAdminBase2W:
		return "IMSAdminBase2W"
	case MSIMSAIMSAdminBase3W:
		return "IMSAdminBase3W"
	case MSIMSAIMSAdminBaseW:
		return "IMSAdminBaseW"
	case MSIMSAIWamAdmin2:
		return "IWamAdmin2"
	case MSIMSAIWamAdmin:
		return "IWamAdmin"
	case MSIOIIManagedObject:
		return "IManagedObject"
	case MSIOIIRemoteDispatch:
		return "IRemoteDispatch"
	case MSIOIIServicedComponentInfo:
		return "IServicedComponentInfo"
	case MSIRPInetinfo:
		return "inetinfo"
	case MSISTMIDirectoryEnum:
		return "IDirectoryEnum"
	case MSISTMIEnumCachedInitiator:
		return "IEnumCachedInitiator"
	case MSISTMIEnumConnection:
		return "IEnumConnection"
	case MSISTMIEnumDisk2:
		return "IEnumDisk2"
	case MSISTMIEnumDisk:
		return "IEnumDisk"
	case MSISTMIEnumHost:
		return "IEnumHost"
	case MSISTMIEnumIDMethod:
		return "IEnumIDMethod"
	case MSISTMIEnumLMMountPoint:
		return "IEnumLMMountPoint"
	case MSISTMIEnumPortal:
		return "IEnumPortal"
	case MSISTMIEnumResourceGroup:
		return "IEnumResourceGroup"
	case MSISTMIEnumSession2:
		return "IEnumSession2"
	case MSISTMIEnumSession:
		return "IEnumSession"
	case MSISTMIEnumSnapshot:
		return "IEnumSnapshot"
	case MSISTMIEnumSnsServer:
		return "IEnumSnsServer"
	case MSISTMIEnumVolume2:
		return "IEnumVolume2"
	case MSISTMIEnumVolume:
		return "IEnumVolume"
	case MSISTMIEnumWTDiskLunMapping:
		return "IEnumWTDiskLunMapping"
	case MSISTMIEnumWTDisk:
		return "IEnumWTDisk"
	case MSISTMIHost2:
		return "IHost2"
	case MSISTMIHost3:
		return "IHost3"
	case MSISTMIHostMgr2:
		return "IHostMgr2"
	case MSISTMIHostMgr3:
		return "IHostMgr3"
	case MSISTMIHostMgr:
		return "IHostMgr"
	case MSISTMIHost:
		return "IHost"
	case MSISTMILocalDeviceMgr2:
		return "ILocalDeviceMgr2"
	case MSISTMILocalDeviceMgr3:
		return "ILocalDeviceMgr3"
	case MSISTMILocalDeviceMgr:
		return "ILocalDeviceMgr"
	case MSISTMISessionManager2:
		return "ISessionManager2"
	case MSISTMISessionManager:
		return "ISessionManager"
	case MSISTMISnapshotMgr:
		return "ISnapshotMgr"
	case MSISTMISnapshot:
		return "ISnapshot"
	case MSISTMISnsMgr:
		return "ISnsMgr"
	case MSISTMIStatusNotify:
		return "IStatusNotify"
	case MSISTMIWTDisk2:
		return "IWTDisk2"
	case MSISTMIWTDisk3:
		return "IWTDisk3"
	case MSISTMIWTDiskMgr2:
		return "IWTDiskMgr2"
	case MSISTMIWTDiskMgr3:
		return "IWTDiskMgr3"
	case MSISTMIWTDiskMgr4:
		return "IWTDiskMgr4"
	case MSISTMIWTDiskMgr:
		return "IWTDiskMgr"
	case MSISTMIWTDisk:
		return "IWTDisk"
	case MSISTMIWTGeneral2:
		return "IWTGeneral2"
	case MSISTMIWTGeneral:
		return "IWTGeneral"
	case MSLRECNetEventForwarder:
		return "NetEventForwarder"
	case MSLSADLsarpc:
		return "lsarpc"
	case MSLSATLsarpc:
		return "lsarpc"
	case MSMQDSDscomm2:
		return "dscomm2"
	case MSMQDSDscomm:
		return "dscomm"
	case MSMQMPQmcomm2:
		return "qmcomm2"
	case MSMQMPQmcomm:
		return "qmcomm"
	case MSMQMRQmmgmt:
		return "qmmgmt"
	case MSMQQPQm2qm:
		return "qm2qm"
	case MSMQRRRemoteRead:
		return "RemoteRead"
	case MSMSRPMsgsvcsend:
		return "msgsvcsend"
	case MSMSRPMsgsvc:
		return "msgsvc"
	case MSNRPCLogon:
		return "logon"
	case MSNSPINspi:
		return "nspi"
	case MSOAUTIDispatch:
		return "IDispatch"
	case MSOAUTIEnumVARIANT:
		return "IEnumVARIANT"
	case MSOAUTITypeComp:
		return "ITypeComp"
	case MSOAUTITypeInfo2:
		return "ITypeInfo2"
	case MSOAUTITypeInfo:
		return "ITypeInfo"
	case MSOAUTITypeLib2:
		return "ITypeLib2"
	case MSOAUTITypeLib:
		return "ITypeLib"
	case MSOAUTIUnknown:
		return "IUnknown"
	case MSOCSPAIOCSPAdminD:
		return "IOCSPAdminD"
	case MSPANIRPCAsyncNotify:
		return "IRPCAsyncNotify"
	case MSPANIRPCRemoteObject:
		return "IRPCRemoteObject"
	case MSPARIRemoteWinspool:
		return "IRemoteWinspool"
	case MSPCQPerflibV2:
		return "PerflibV2"
	case MSPLAIAlertDataCollector:
		return "IAlertDataCollector"
	case MSPLAIApiTracingDataCollector:
		return "IApiTracingDataCollector"
	case MSPLAIConfigurationDataCollector:
		return "IConfigurationDataCollector"
	case MSPLAIDataCollectorCollection:
		return "IDataCollectorCollection"
	case MSPLAIDataCollectorSetCollection:
		return "IDataCollectorSetCollection"
	case MSPLAIDataCollectorSet:
		return "IDataCollectorSet"
	case MSPLAIDataCollector:
		return "IDataCollector"
	case MSPLAIDataManager:
		return "IDataManager"
	case MSPLAIFolderActionCollection:
		return "IFolderActionCollection"
	case MSPLAIFolderAction:
		return "IFolderAction"
	case MSPLAIPerformanceCounterDataCollector:
		return "IPerformanceCounterDataCollector"
	case MSPLAIScheduleCollection:
		return "IScheduleCollection"
	case MSPLAISchedule:
		return "ISchedule"
	case MSPLAITraceDataCollector:
		return "ITraceDataCollector"
	case MSPLAITraceDataProviderCollection:
		return "ITraceDataProviderCollection"
	case MSPLAITraceDataProvider:
		return "ITraceDataProvider"
	case MSPLAIValueMapItem:
		return "IValueMapItem"
	case MSPLAIValueMap:
		return "IValueMap"
	case MSRAAAuthzr:
		return "authzr"
	case MSRAAIgnoreCentralPolicy:
		return "IgnoreCentralPolicy"
	case MSRAAUseCentralPolicy:
		return "UseCentralPolicy"
	case MSRAINPSIIASDataStoreComServer2:
		return "IIASDataStoreComServer2"
	case MSRAINPSIIASDataStoreComServer:
		return "IIASDataStoreComServer"
	case MSRAIIPCHCollection:
		return "IPCHCollection"
	case MSRAIIPCHService:
		return "IPCHService"
	case MSRAIIRASrv:
		return "IRASrv"
	case MSRAIISAFSession:
		return "ISAFSession"
	case MSRAIWWinsif2:
		return "winsif2"
	case MSRAIWWinsif:
		return "winsif"
	case MSRDPESCTypeScardPack:
		return "TypeScardPack"
	case MSRPCEEndpointMapper:
		return "EndpointMapper"
	case MSRPCERemoteManagement:
		return "RemoteManagement"
	case MSRPCETransferNDR64:
		return "TransferNDR64"
	case MSRPCETransferNDR:
		return "TransferNDR"
	case MSRPCLLocToLoc:
		return "LocToLoc"
	case MSRPRNWinspool:
		return "winspool"
	case MSRRASMDIMSVC:
		return "DIMSVC"
	case MSRRASMIRemoteICFICSConfig:
		return "IRemoteICFICSConfig"
	case MSRRASMIRemoteIPV6Config:
		return "IRemoteIPV6Config"
	case MSRRASMIRemoteNetworkConfig:
		return "IRemoteNetworkConfig"
	case MSRRASMIRemoteRouterRestart:
		return "IRemoteRouterRestart"
	case MSRRASMIRemoteSetDnsConfig:
		return "IRemoteSetDnsConfig"
	case MSRRASMIRemoteSstpCertCheck:
		return "IRemoteSstpCertCheck"
	case MSRRASMIRemoteStringIdConfig:
		return "IRemoteStringIdConfig"
	case MSRRASMRASRPC:
		return "RASRPC"
	case MSRRPWinreg:
		return "winreg"
	case MSRSMPIClientSink:
		return "IClientSink"
	case MSRSMPIMessenger:
		return "IMessenger"
	case MSRSMPINtmsLibraryControl1:
		return "INtmsLibraryControl1"
	case MSRSMPINtmsLibraryControl2:
		return "INtmsLibraryControl2"
	case MSRSMPINtmsMediaServices1:
		return "INtmsMediaServices1"
	case MSRSMPINtmsNotifySink:
		return "INtmsNotifySink"
	case MSRSMPINtmsObjectInfo1:
		return "INtmsObjectInfo1"
	case MSRSMPINtmsObjectManagement1:
		return "INtmsObjectManagement1"
	case MSRSMPINtmsObjectManagement2:
		return "INtmsObjectManagement2"
	case MSRSMPINtmsObjectManagement3:
		return "INtmsObjectManagement3"
	case MSRSMPINtmsSession1:
		return "INtmsSession1"
	case MSRSMPIRobustNtmsMediaServices1:
		return "IRobustNtmsMediaServices1"
	case MSRSMPIUnknown:
		return "IUnknown"
	case MSRSPInitShutdown:
		return "InitShutdown"
	case MSRSPWindowsShutdown:
		return "WindowsShutdown"
	case MSRSPWinreg:
		return "winreg"
	case MSSAMRSamr:
		return "samr"
	case MSSCMPIVssDifferentialSoftwareSnapshotMgmt:
		return "IVssDifferentialSoftwareSnapshotMgmt"
	case MSSCMPIVssEnumMgmtObject:
		return "IVssEnumMgmtObject"
	case MSSCMPIVssEnumObject:
		return "IVssEnumObject"
	case MSSCMPIVssSnapshotMgmt:
		return "IVssSnapshotMgmt"
	case MSSCMRSvcctl:
		return "svcctl"
	case MSSRVSSrvsvc:
		return "srvsvc"
	case MSSWNWitness:
		return "Witness"
	case MSTPMVSCITpmVirtualSmartCardManager2:
		return "ITpmVirtualSmartCardManager2"
	case MSTPMVSCITpmVirtualSmartCardManager3:
		return "ITpmVirtualSmartCardManager3"
	case MSTPMVSCITpmVirtualSmartCardManagerStatusCallback:
		return "ITpmVirtualSmartCardManagerStatusCallback"
	case MSTPMVSCITpmVirtualSmartCardManager:
		return "ITpmVirtualSmartCardManager"
	case MSTRPRemotesp:
		return "remotesp"
	case MSTRPTapsrv:
		return "tapsrv"
	case MSTSCHATSvc:
		return "ATSvc"
	case MSTSCHITaskSchedulerService:
		return "ITaskSchedulerService"
	case MSTSCHSASec:
		return "SASec"
	case MSTSGUTsProxyRpcInterface:
		return "TsProxyRpcInterface"
	case MSTSRAPIManageTelnetSessions:
		return "IManageTelnetSessions"
	case MSTSTSIcaApi:
		return "IcaApi"
	case MSTSTSRCMListener:
		return "RCMListener"
	case MSTSTSRCMPublic:
		return "RCMPublic"
	case MSTSTSSessEnvPublicRpc:
		return "SessEnvPublicRpc"
	case MSTSTSTermSrvEnumeration:
		return "TermSrvEnumeration"
	case MSTSTSTermSrvNotification:
		return "TermSrvNotification"
	case MSTSTSTermSrvSession:
		return "TermSrvSession"
	case MSTSTSTSVIPPublic:
		return "TSVIPPublic"
	case MSUAMGIAutomaticUpdates2:
		return "IAutomaticUpdates2"
	case MSUAMGIAutomaticUpdatesResults:
		return "IAutomaticUpdatesResults"
	case MSUAMGIAutomaticUpdates:
		return "IAutomaticUpdates"
	case MSUAMGICategoryCollection:
		return "ICategoryCollection"
	case MSUAMGICategory:
		return "ICategory"
	case MSUAMGIImageInformation:
		return "IImageInformation"
	case MSUAMGIInstallationBehavior:
		return "IInstallationBehavior"
	case MSUAMGISearchJob_:
		return "ISearchJob_"
	case MSUAMGISearchResult:
		return "ISearchResult"
	case MSUAMGIStringCollection:
		return "IStringCollection"
	case MSUAMGIUpdate2:
		return "IUpdate2"
	case MSUAMGIUpdate3:
		return "IUpdate3"
	case MSUAMGIUpdate4:
		return "IUpdate4"
	case MSUAMGIUpdate5:
		return "IUpdate5"
	case MSUAMGIUpdateCollection:
		return "IUpdateCollection"
	case MSUAMGIUpdateDownloadContent2:
		return "IUpdateDownloadContent2"
	case MSUAMGIUpdateDownloadContentCollection:
		return "IUpdateDownloadContentCollection"
	case MSUAMGIUpdateDownloadContent:
		return "IUpdateDownloadContent"
	case MSUAMGIUpdateExceptionCollectionOLE:
		return "IUpdateExceptionCollectionOLE"
	case MSUAMGIUpdateExceptionCollection:
		return "IUpdateExceptionCollection"
	case MSUAMGIUpdateException:
		return "IUpdateException"
	case MSUAMGIUpdateHistoryEntry2:
		return "IUpdateHistoryEntry2"
	case MSUAMGIUpdateHistoryEntryCollection:
		return "IUpdateHistoryEntryCollection"
	case MSUAMGIUpdateHistoryEntry:
		return "IUpdateHistoryEntry"
	case MSUAMGIUpdateIdentity:
		return "IUpdateIdentity"
	case MSUAMGIUpdateSearcher2:
		return "IUpdateSearcher2"
	case MSUAMGIUpdateSearcher3:
		return "IUpdateSearcher3"
	case MSUAMGIUpdateSearcher:
		return "IUpdateSearcher"
	case MSUAMGIUpdateService2:
		return "IUpdateService2"
	case MSUAMGIUpdateServiceCollection:
		return "IUpdateServiceCollection"
	case MSUAMGIUpdateServiceManager2:
		return "IUpdateServiceManager2"
	case MSUAMGIUpdateServiceManager:
		return "IUpdateServiceManager"
	case MSUAMGIUpdateServiceRegistration:
		return "IUpdateServiceRegistration"
	case MSUAMGIUpdateService:
		return "IUpdateService"
	case MSUAMGIUpdateSession2:
		return "IUpdateSession2"
	case MSUAMGIUpdateSession3:
		return "IUpdateSession3"
	case MSUAMGIUpdateSession:
		return "IUpdateSession"
	case MSUAMGIUpdate:
		return "IUpdate"
	case MSUAMGIWindowsDriverUpdate2:
		return "IWindowsDriverUpdate2"
	case MSUAMGIWindowsDriverUpdate3:
		return "IWindowsDriverUpdate3"
	case MSUAMGIWindowsDriverUpdate4:
		return "IWindowsDriverUpdate4"
	case MSUAMGIWindowsDriverUpdate5:
		return "IWindowsDriverUpdate5"
	case MSUAMGIWindowsDriverUpdateEntryCollection:
		return "IWindowsDriverUpdateEntryCollection"
	case MSUAMGIWindowsDriverUpdateEntry:
		return "IWindowsDriverUpdateEntry"
	case MSUAMGIWindowsDriverUpdate:
		return "IWindowsDriverUpdate"
	case MSUAMGIWindowsUpdateAgentInfo:
		return "IWindowsUpdateAgentInfo"
	case MSVDSIEnumVdsObject:
		return "IEnumVdsObject"
	case MSVDSIVdsAdvancedDisk2:
		return "IVdsAdvancedDisk2"
	case MSVDSIVdsAdvancedDisk3:
		return "IVdsAdvancedDisk3"
	case MSVDSIVdsAdvancedDisk:
		return "IVdsAdvancedDisk"
	case MSVDSIVdsAdviseSink:
		return "IVdsAdviseSink"
	case MSVDSIVdsAsync:
		return "IVdsAsync"
	case MSVDSIVdsCreatePartitionEx:
		return "IVdsCreatePartitionEx"
	case MSVDSIVdsDisk2:
		return "IVdsDisk2"
	case MSVDSIVdsDisk3:
		return "IVdsDisk3"
	case MSVDSIVdsDiskOnline:
		return "IVdsDiskOnline"
	case MSVDSIVdsDiskPartitionMF2:
		return "IVdsDiskPartitionMF2"
	case MSVDSIVdsDiskPartitionMF:
		return "IVdsDiskPartitionMF"
	case MSVDSIVdsDisk:
		return "IVdsDisk"
	case MSVDSIVdsHbaPort:
		return "IVdsHbaPort"
	case MSVDSIVdsHwProvider:
		return "IVdsHwProvider"
	case MSVDSIVdsIscsiInitiatorAdapter:
		return "IVdsIscsiInitiatorAdapter"
	case MSVDSIVdsIscsiInitiatorPortal:
		return "IVdsIscsiInitiatorPortal"
	case MSVDSIVdsOpenDisk:
		return "IVdsOpenDisk"
	case MSVDSIVdsPack2:
		return "IVdsPack2"
	case MSVDSIVdsPack:
		return "IVdsPack"
	case MSVDSIVdsProvider:
		return "IVdsProvider"
	case MSVDSIVdsRemovable:
		return "IVdsRemovable"
	case MSVDSIVdsServiceHba:
		return "IVdsServiceHba"
	case MSVDSIVdsServiceInitialization:
		return "IVdsServiceInitialization"
	case MSVDSIVdsServiceIscsi:
		return "IVdsServiceIscsi"
	case MSVDSIVdsServiceLoader:
		return "IVdsServiceLoader"
	case MSVDSIVdsServiceSAN:
		return "IVdsServiceSAN"
	case MSVDSIVdsServiceSw:
		return "IVdsServiceSw"
	case MSVDSIVdsServiceUninstallDisk:
		return "IVdsServiceUninstallDisk"
	case MSVDSIVdsService:
		return "IVdsService"
	case MSVDSIVdsSubSystemImportTarget:
		return "IVdsSubSystemImportTarget"
	case MSVDSIVdsSwProvider:
		return "IVdsSwProvider"
	case MSVDSIVdsVDisk:
		return "IVdsVDisk"
	case MSVDSIVdsVdProvider:
		return "IVdsVdProvider"
	case MSVDSIVdsVolume2:
		return "IVdsVolume2"
	case MSVDSIVdsVolumeMF2:
		return "IVdsVolumeMF2"
	case MSVDSIVdsVolumeMF3:
		return "IVdsVolumeMF3"
	case MSVDSIVdsVolumeMF:
		return "IVdsVolumeMF"
	case MSVDSIVdsVolumeOnline:
		return "IVdsVolumeOnline"
	case MSVDSIVdsVolumePlex:
		return "IVdsVolumePlex"
	case MSVDSIVdsVolumeShrink:
		return "IVdsVolumeShrink"
	case MSVDSIVdsVolume:
		return "IVdsVolume"
	case MSW32TW32Time:
		return "W32Time"
	case MSWCCEICertRequestD2:
		return "ICertRequestD2"
	case MSWCCEICertRequestD:
		return "ICertRequestD"
	case MSWDSCWdsRpcInterface:
		return "WdsRpcInterface"
	case MSWKSTWkssvc:
		return "wkssvc"
	case MSWMIIEnumWbemClassObject:
		return "IEnumWbemClassObject"
	case MSWMIIWbemBackupRestoreEx:
		return "IWbemBackupRestoreEx"
	case MSWMIIWbemBackupRestore:
		return "IWbemBackupRestore"
	case MSWMIIWbemClassObject:
		return "IWbemClassObject"
	case MSWMIIWbemContext:
		return "IWbemContext"
	case MSWMIIWbemFetchSmartEnum:
		return "IWbemFetchSmartEnum"
	case MSWMIIWbemLevel1Login:
		return "IWbemLevel1Login"
	case MSWMIIWbemLoginClientID:
		return "IWbemLoginClientID"
	case MSWMIIWbemLoginHelper:
		return "IWbemLoginHelper"
	case MSWMIIWbemObjectSink:
		return "IWbemObjectSink"
	case MSWMIIWbemRefreshingServices:
		return "IWbemRefreshingServices"
	case MSWMIIWbemRemoteRefresher:
		return "IWbemRemoteRefresher"
	case MSWMIIWbemServices:
		return "IWbemServices"
	case MSWMIIWbemWCOSmartEnum:
		return "IWbemWCOSmartEnum"
	case MSWSRMIResourceManager2:
		return "IResourceManager2"
	case MSWSRMIResourceManager:
		return "IResourceManager"
	case MSWSRMIWRMAccounting:
		return "IWRMAccounting"
	case MSWSRMIWRMCalendar:
		return "IWRMCalendar"
	case MSWSRMIWRMConfig:
		return "IWRMConfig"
	case MSWSRMIWRMMachineGroup:
		return "IWRMMachineGroup"
	case MSWSRMIWRMPolicy:
		return "IWRMPolicy"
	case MSWSRMIWRMProtocol:
		return "IWRMProtocol"
	case MSWSRMIWRMRemoteSessionMgmt:
		return "IWRMRemoteSessionMgmt"
	case MSWSRMIWRMResourceGroup:
		return "IWRMResourceGroup"
	case MSDLLAdvapi32:
		return "advapi32"
	case MSDLLAppidsvc:
		return "appidsvc"
	case MSDLLAppinfo:
		return "appinfo"
	case MSDLLIAisAxISInfo:
		return "IAisAxISInfo"
	case MSDLLIAisCOMInfo:
		return "IAisCOMInfo"
	case MSDLLIAisMSIInfo:
		return "IAisMSIInfo"
	case MSDLLAppmgmts:
		return "appmgmts"
	case MSDLLAqueue:
		return "aqueue"
	case MSDLLAudiodgrpc:
		return "audiodgrpc"
	case MSDLLAudiosrv:
		return "audiosrv"
	case MSDLLAudiorpc:
		return "audiorpc"
	case MSDLLWluirRpc:
		return "WluirRpc"
	case MSDLLWluir1Rpc:
		return "Wluir1Rpc"
	case MSDLLAuthmgr32Rpc:
		return "Authmgr32Rpc"
	case MSDLLBdesvc:
		return "bdesvc"
	case MSDLLBferpc:
		return "bferpc"
	case MSDLLBthServRPCService:
		return "BthServRPCService"
	case MSDLLCertPropSvc:
		return "CertPropSvc"
	case MSDLLClusSvcJoinVersion:
		return "ClusSvcJoinVersion"
	case MSDLLClusSvc:
		return "ClusSvc"
	case MSDLLClusSvcExtrocluster:
		return "ClusSvcExtrocluster"
	case MSDLLCryptCertProtectSvc:
		return "CryptCertProtectSvc"
	case MSDLLCryptKSrSvc:
		return "CryptKSrSvc"
	case MSDLLCryptKeyrSvc:
		return "CryptKeyrSvc"
	case MSDLLCryptCatSvc:
		return "CryptCatSvc"
	case MSDLLDhcpcsvc6:
		return "dhcpcsvc6"
	case MSDLLDhcpcsvc:
		return "dhcpcsvc"
	case MSDLLDmadminrpc:
		return "dmadminrpc"
	case MSDLLDnsResolver:
		return "DnsResolver"
	case MSDLLDnsResolver2000:
		return "DnsResolver2000"
	case MSDLLWinlan:
		return "winlan"
	case MSDLLEfskrpc:
		return "efskrpc"
	case MSDLLEfsrpc:
		return "efsrpc"
	case MSDLLEmdService:
		return "EmdService"
	case MSDLLGpsvc:
		return "gpsvc"
	case MSDLLIasrpc:
		return "iasrpc"
	case MSDLLIcardagtrpc:
		return "icardagtrpc"
	case MSDLLIertutil:
		return "iertutil"
	case MSDLLIdletask:
		return "idletask"
	case MSDLLISecureDesktop:
		return "ISecureDesktop"
	case MSDLLWMsgKAPIs:
		return "WMsgKAPIs"
	case MSDLLWMsgAPIs:
		return "WMsgAPIs"
	case MSDLLIRPCSCLogon:
		return "IRPCSCLogon"
	case MSDLLNrpsrv:
		return "nrpsrv"
	case MSDLLWinhttprpc:
		return "winhttprpc"
	case MSDLLIRPCAsyncNotifyChannel:
		return "IRPCAsyncNotifyChannel"
	case MSDLLIpTransition:
		return "IpTransition"
	case MSDLLWinNsi:
		return "WinNsi"
	case MSDLLXactSrvRPC:
		return "XactSrvRPC"
	case MSDLLIkeRpcIKE:
		return "IkeRpcIKE"
	case MSDLLIRPCWinlogonNotifications:
		return "IRPCWinlogonNotifications"
	case MSDLLIfssvc:
		return "ifssvc"
	case MSDLLImepadsm:
		return "imepadsm"
	case MSDLLImepadsm1:
		return "imepadsm1"
	case MSDLLImjpdctrpc:
		return "imjpdctrpc"
	case MSDLLCpprovider:
		return "cpprovider"
	case MSDLLIphlpsvc:
		return "iphlpsvc"
	case MSDLLIpnathlp:
		return "ipnathlp"
	case MSDLLIpNatHlpRpcServer:
		return "IpNatHlpRpcServer"
	case MSDLLIrftprpc:
		return "irftprpc"
	case MSDLLIrmon:
		return "irmon"
	case MSDLLIrmon1:
		return "irmon1"
	case MSDLLIscsiexe:
		return "iscsiexe"
	case MSDLLIsmserv_ip:
		return "ismserv_ip"
	case MSDLLIsmapi:
		return "ismapi"
	case MSDLLKdcrpc:
		return "kdcrpc"
	case MSDLLINCryptKeyIso:
		return "INCryptKeyIso"
	case MSDLLLbservice:
		return "lbservice"
	case MSDLLLlsrpc:
		return "llsrpc"
	case MSDLLLls_license:
		return "lls_license"
	case MSDLLNsiC:
		return "NsiC"
	case MSDLLNsiS:
		return "NsiS"
	case MSDLLNsiM:
		return "NsiM"
	case MSDLLDsrolesvc:
		return "dsrolesvc"
	case MSDLLICryptProtect:
		return "ICryptProtect"
	case MSDLLPasswordRecovery:
		return "PasswordRecovery"
	case MSDLLSLspPrivateData:
		return "SLspPrivateData"
	case MSDLLLsass:
		return "lsass"
	case MSDLLTermServLicensing:
		return "TermServLicensing"
	case MSDLLHydraLsPipe:
		return "HydraLsPipe"
	case MSDLLTermSrvPrivate:
		return "TermSrvPrivate"
	case MSDLLTermSrvAdmin:
		return "TermSrvAdmin"
	case MSDLLMilcore:
		return "milcore"
	case MSDLLMpnotifyrpc:
		return "mpnotifyrpc"
	case MSDLLACPBackgroundManagerPolicy:
		return "ACPBackgroundManagerPolicy"
	case MSDLLAdhsvc:
		return "adhsvc"
	case MSDLLAphostservice:
		return "aphostservice"
	case MSDLLAppinfopkg:
		return "appinfopkg"
	case MSDLLAudiosrvpbm:
		return "audiosrvpbm"
	case MSDLLAudiosrvpbm1:
		return "audiosrvpbm1"
	case MSDLLAudiosrvhrtf:
		return "audiosrvhrtf"
	case MSDLLAudiosrvwnf:
		return "audiosrvwnf"
	case MSDLLAudiosrv1:
		return "audiosrv1"
	case MSDLLAudiosrv2:
		return "audiosrv2"
	case MSDLLAudiosrv3:
		return "audiosrv3"
	case MSDLLRBiPtSrv:
		return "RBiPtSrv"
	case MSDLLRBiRtSrv:
		return "RBiRtSrv"
	case MSDLLSrvOdbPt:
		return "SrvOdbPt"
	case MSDLLSrvOdb:
		return "SrvOdb"
	case MSDLLRBiSrv:
		return "RBiSrv"
	case MSDLLSrvOdbPriv:
		return "SrvOdbPriv"
	case MSDLLRBiSrv1:
		return "RBiSrv1"
	case MSDLLSrvOdbLb:
		return "SrvOdbLb"
	case MSDLLRBiSrvSignal:
		return "RBiSrvSignal"
	case MSDLLBri:
		return "Bri"
	case MSDLLIChakraJIT:
		return "IChakraJIT"
	case MSDLLCoremessaging:
		return "coremessaging"
	case MSDLLCrypttpmeksvc:
		return "crypttpmeksvc"
	case MSDLLWarpJITSvc:
		return "WarpJITSvc"
	case MSDLLDab:
		return "dab"
	case MSDLLDaschallenge:
		return "daschallenge"
	case MSDLLDasaepstore:
		return "dasaepstore"
	case MSDLLDas:
		return "das"
	case MSDLLDasinbassoc:
		return "dasinbassoc"
	case MSDLLDasrpcdev:
		return "dasrpcdev"
	case MSDLLDashost:
		return "dashost"
	case MSDLLUtcApi:
		return "UtcApi"
	case MSDLLUtcTelemetryOptInApi:
		return "UtcTelemetryOptInApi"
	case MSDLLUtcWerHelperApi:
		return "UtcWerHelperApi"
	case MSDLLUtcEventTranscriptApi:
		return "UtcEventTranscriptApi"
	case MSDLLUtcTenantApi:
		return "UtcTenantApi"
	case MSDLLUtcApi1:
		return "UtcApi1"
	case MSDLLDpapisrv:
		return "dpapisrv"
	case MSDLLDssvc:
		return "dssvc"
	case MSDLLDusmsvc:
		return "dusmsvc"
	case MSDLLEeprov:
		return "eeprov"
	case MSDLLPsmSrvAppHost:
		return "PsmSrvAppHost"
	case MSDLLPsmSrvApp:
		return "PsmSrvApp"
	case MSDLLPsmSrvInfo:
		return "PsmSrvInfo"
	case MSDLLPsmSrvActivation:
		return "PsmSrvActivation"
	case MSDLLPsmSrvTimer:
		return "PsmSrvTimer"
	case MSDLLPsmSrvQuiesce:
		return "PsmSrvQuiesce"
	case MSDLLPsmSrv:
		return "PsmSrv"
	case MSDLLWscsvc:
		return "wscsvc"
	case MSDLLPcaSvc:
		return "PcaSvc"
	case MSDLLRmgrSrv:
		return "RmgrSrv"
	case MSDLLRmtSrv:
		return "RmtSrv"
	case MSDLLHamRpcSrv:
		return "HamRpcSrv"
	case MSDLLHamRpcSrvSess:
		return "HamRpcSrvSess"
	case MSDLLHamRpcSrvHostId:
		return "HamRpcSrvHostId"
	case MSDLLRmGameModeRSrv:
		return "RmGameModeRSrv"
	case MSDLLRmCoreRpcSrv:
		return "RmCoreRpcSrv"
	case MSDLLHamRpcSrvState:
		return "HamRpcSrvState"
	case MSDLLHamRpcSrvFullTrust:
		return "HamRpcSrvFullTrust"
	case MSDLLHamRpcSrvServicing:
		return "HamRpcSrvServicing"
	case MSDLLCrmRpcSrv:
		return "CrmRpcSrv"
	case MSDLLHamRpcSrvActivity:
		return "HamRpcSrvActivity"
	case MSDLLHamRpcSrvDebug:
		return "HamRpcSrvDebug"
	case MSDLLRmGameModeSrv:
		return "RmGameModeSrv"
	case MSDLLActiveSyncServer:
		return "ActiveSyncServer"
	case MSDLLAccountsMgmtRpc:
		return "AccountsMgmtRpc"
	case MSDLLIUserMarshalData:
		return "IUserMarshalData"
	case MSDLLUsermgr:
		return "usermgr"
	case MSDLLFmMuxSrvNotificationCoreUI:
		return "FmMuxSrvNotificationCoreUI"
	case MSDLLFmMuxSrvCoreUI:
		return "FmMuxSrvCoreUI"
	case MSDLLFmMuxSrv:
		return "FmMuxSrv"
	case MSDLLIdSegSrv:
		return "IdSegSrv"
	case MSDLLTSSDFarmRpc:
		return "TSSDFarmRpc"
	case MSDLLTSSDRpc:
		return "TSSDRpc"
	case MSDLLNgcHandlerRpc:
		return "NgcHandlerRpc"
	case MSDLLProxyMgr:
		return "ProxyMgr"
	case MSDLLRpcSrvProxyMgr:
		return "RpcSrvProxyMgr"
	case MSDLLTaskScheduler:
		return "TaskScheduler"
	case MSDLLPhoneRpc:
		return "PhoneRpc"
	case MSDLLPhoneRpc1:
		return "PhoneRpc1"
	case MSDLLPhoneRpc2:
		return "PhoneRpc2"
	case MSDLLPhoneRpc3:
		return "PhoneRpc3"
	case MSDLLPhoneRpc4:
		return "PhoneRpc4"
	case MSDLLI_pSchRpc:
		return "I_pSchRpc"
	case MSDLLPimIMService:
		return "PimIMService"
	case MSDLLSsdpsrv:
		return "ssdpsrv"
	case MSDLLWinBioCredentialManager:
		return "WinBioCredentialManager"
	case MSDLLWinBioServer:
		return "WinBioServer"
	case MSDLLColorManagementRpcServer:
		return "ColorManagementRpcServer"
	case MSDLLNlasvc:
		return "nlasvc"
	case MSDLLTokenBindingApi:
		return "TokenBindingApi"
	case MSDLLSymmetricPopKeyApi:
		return "SymmetricPopKeyApi"
	case MSDLLNcbRpcSrv:
		return "NcbRpcSrv"
	case MSDLLNcbKapi:
		return "NcbKapi"
	case MSDLLNcbRpcSrv1:
		return "NcbRpcSrv1"
	case MSDLLStorSvc:
		return "StorSvc"
	case MSDLLStorSvc1:
		return "StorSvc1"
	case MSDLLGcsSrv:
		return "GcsSrv"
	case MSDLLClipSvc:
		return "ClipSvc"
	case MSDLLRpSrv:
		return "RpSrv"
	case MSDLLRpSrv1:
		return "RpSrv1"
	case MSDLLRasCustomRpc:
		return "RasCustomRpc"
	case MSDLLCSebiPublisher:
		return "CSebiPublisher"
	case MSDLLNgcSecureBioHandlerRpc:
		return "NgcSecureBioHandlerRpc"
	case MSDLLCSebiEvent:
		return "CSebiEvent"
	case MSDLLSebiEvent:
		return "SebiEvent"
	case MSDLLSgrmBroker:
		return "SgrmBroker"
	case MSDLLSetOperatorRpc:
		return "SetOperatorRpc"
	case MSDLLWcmSvcRpc:
		return "WcmSvcRpc"
	case MSDLLWcmSvcRpc1:
		return "WcmSvcRpc1"
	case MSDLLWcmSvcProxyInfoRpc:
		return "WcmSvcProxyInfoRpc"
	case MSDLLWcmSvcMgrRpc:
		return "WcmSvcMgrRpc"
	case MSDLLWcmSvcCapRpc:
		return "WcmSvcCapRpc"
	case MSDLLDfsDsSvc:
		return "DfsDsSvc"
	case MSDLLFwNetIsolationRpc:
		return "FwNetIsolationRpc"
	case MSDLLFwRpc1:
		return "FwRpc1"
	case MSDLLFwRpc:
		return "FwRpc"
	case MSDLLNetSetupShim:
		return "NetSetupShim"
	case MSDLLNetSetupSvc:
		return "NetSetupSvc"
	case MSDLLDeviceCredentialMgrRpc:
		return "DeviceCredentialMgrRpc"
	case MSDLLVscRpc:
		return "VscRpc"
	case MSDLLNgcLocalAccountVaultRpc:
		return "NgcLocalAccountVaultRpc"
	case MSDLLPhoneProviders:
		return "PhoneProviders"
	case MSDLLLicenseManagerSvc:
		return "LicenseManagerSvc"
	case MSDLLTbiRpc:
		return "TbiRpc"
	case MSDLLDeManagementRpcServer:
		return "DeManagementRpcServer"
	case MSDLLDeoRpcServer:
		return "DeoRpcServer"
	case MSDLLPfRpcServerSuperfetch:
		return "PfRpcServerSuperfetch"
	case MSDLLWcnpRpc:
		return "WcnpRpc"
	case MSDLLWcnTransportRpc:
		return "WcnTransportRpc"
	case MSDLLNgcMgmt:
		return "NgcMgmt"
	case MSDLLUdmSvc:
		return "UdmSvc"
	case MSDLLUtcAdminApi:
		return "UtcAdminApi"
	case MSDLLNgcMgmt1:
		return "NgcMgmt1"
	case MSDLLGidsHandlerRpc:
		return "GidsHandlerRpc"
	case MSDLLHcsRpc:
		return "HcsRpc"
	case MSDLLUSSvc:
		return "USSvc"
	case MSDLLClusterConnRpc:
		return "ClusterConnRpc"
	case MSDLLWitnessNodesRpc:
		return "WitnessNodesRpc"
	case MSDLLLogonWluirRpc:
		return "LogonWluirRpc"
	case MSDLLDeviceCredentialRpc:
		return "DeviceCredentialRpc"
	case MSDLLNgcRpc:
		return "NgcRpc"
	case MSDLLSECOMNRpc:
		return "SECOMNRpc"
	case MSDLLIhvmsmapi:
		return "ihvmsmapi"
	case MSDLLWiFiUserRpc:
		return "WiFiUserRpc"
	case MSDLLWlanSvcLowPriv:
		return "WlanSvcLowPriv"
	case MSDLLWlanSvc:
		return "WlanSvc"
	case MSDLLWDiagRpc:
		return "WDiagRpc"
	case MSDLLRpcDot11ExtIhv:
		return "RpcDot11ExtIhv"
	case MSDLLSignalRpc:
		return "SignalRpc"
	case MSDLLLiveIdSvcNotifyRpc:
		return "LiveIdSvcNotifyRpc"
	case MSDLLLiveIdSvcRpc:
		return "LiveIdSvcRpc"
	case MSDLLOnlineProviderCertRpc:
		return "OnlineProviderCertRpc"
	case MSDLLUnimdmRpc:
		return "UnimdmRpc"
	case MSDLLColorAdapterRpcServer:
		return "ColorAdapterRpcServer"
	case MSDLLSCardCacheRpc:
		return "SCardCacheRpc"
	case MSDLLSCardRpc:
		return "SCardRpc"
	case MSDLLBthEvtBrRpc:
		return "BthEvtBrRpc"
	case MSDLLRpcBluetoothGatt:
		return "RpcBluetoothGatt"
	case MSDLLAppXSetTrustLabelRpc:
		return "AppXSetTrustLabelRpc"
	case MSDLLAppXSetTrustLabel1Rpc:
		return "AppXSetTrustLabel1Rpc"
	case MSDLLVoIpRpc:
		return "VoIpRpc"
	case MSDLLWFDSConMgrRpc:
		return "WFDSConMgrRpc"
	case MSDLLFrameServerRpc:
		return "FrameServerRpc"
	case MSDLLWwanRpc:
		return "WwanRpc"
	case MSDLLWwan2Rpc:
		return "Wwan2Rpc"
	case MSDLLWwanMgmtRpc:
		return "WwanMgmtRpc"
	case MSDLLIMpService:
		return "IMpService"
	case RPCSSIMachineActivatorControl:
		return "IMachineActivatorControl"
	case RPCSSISCMLocalActivator:
		return "ISCMLocalActivator"
	case RPCSSISCM:
		return "ISCM"
	case RPCSSIROT:
		return "IROT"
	case RPCSSILocalObjectExporter:
		return "ILocalObjectExporter"
	case RPCSSLocalepmp:
		return "localepmp"
	case RPCSSDbgidl:
		return "Dbgidl"
	case RPCSSFwidl:
		return "Fwidl"
	}
	return ""
}
