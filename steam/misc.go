package steam

const (
	headerStr       = "\xFF\xFF\xFF\xFF"
	nullByteStr     = "\x00\x00\x00\x00\x00\x00\x00\x00"
	maxHosts        = 2500 // max# hosts to retrieve; cannot be larger than 6930 per steam
	maxPacketSize   = 1400 // specified by steam protocol
	QueryTimeout    = 3    // sec; connect, read & write timeout
	QueryRetryCount = 3    // # of times to re-request rules, players, info on failure
)

var (
	testServerInfoData = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x49, 0x11, 0x71, 0x6C, 0x2E, 0x73, 0x79, 0x6E,
		0x63, 0x6F, 0x72, 0x65, 0x2E, 0x6F, 0x72, 0x67, 0x20, 0x2D, 0x20, 0x55,
		0x53, 0x20, 0x43, 0x45, 0x4E, 0x54, 0x52, 0x41, 0x4C, 0x20, 0x23, 0x31,
		0x00, 0x74, 0x68, 0x75, 0x6E, 0x64, 0x65, 0x72, 0x73, 0x74, 0x72, 0x75,
		0x63, 0x6B, 0x00, 0x62, 0x61, 0x73, 0x65, 0x71, 0x33, 0x00, 0x43, 0x6C,
		0x61, 0x6E, 0x20, 0x41, 0x72, 0x65, 0x6E, 0x61, 0x00, 0x00, 0x00, 0x02,
		0x10, 0x00, 0x64, 0x6C, 0x00, 0x01, 0x31, 0x30, 0x36, 0x33, 0x00, 0xB1,
		0x38, 0x6D, 0x02, 0xF8, 0xC1, 0x4D, 0x7B, 0x17, 0x40, 0x01, 0x63, 0x6C,
		0x61, 0x6E, 0x61, 0x72, 0x65, 0x6E, 0x61, 0x2C, 0x73, 0x79, 0x6E, 0x63,
		0x6F, 0x72, 0x65, 0x2C, 0x74, 0x65, 0x78, 0x61, 0x73, 0x00, 0x48, 0x4F,
		0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	testPlayerInfoData = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x44, 0x0A, 0x00, 0x46, 0x6F, 0x72, 0x6D, 0x6F,
		0x73, 0x69, 0x74, 0x79, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4B, 0x00, 0x2A,
		0x45, 0x00, 0x54, 0x68, 0x72, 0x65, 0x65, 0x77, 0x61, 0x79, 0x63, 0x72,
		0x61, 0x70, 0x73, 0x00, 0x00, 0x00, 0x00, 0x00, 0x26, 0x99, 0x1F, 0x45,
		0x00, 0x68, 0x61, 0x78, 0x20, 0x5F, 0x63, 0x6C, 0x61, 0x79, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x6C, 0x85, 0x1F, 0x45, 0x00, 0x7A, 0x75, 0x72, 0x6E,
		0x2D, 0x61, 0x66, 0x6B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x14, 0xE1, 0xD8,
		0x44, 0x00, 0x2E, 0x6D, 0x69, 0x6E, 0x74, 0x6F, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x4C, 0xD7, 0xB3, 0x44, 0x00, 0x6A, 0x61, 0x73, 0x33, 0x39, 0x30,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x33, 0x61, 0x89, 0x44, 0x00, 0x63, 0x6A,
		0x75, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1C, 0x61, 0x80, 0x44, 0x00, 0x76,
		0x72, 0x65, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x96, 0x79, 0x69, 0x43,
		0x00, 0x73, 0x6C, 0x65, 0x65, 0x70, 0x79, 0x64, 0x6F, 0x67, 0x31, 0x34,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0xCF, 0x48, 0x43, 0x00, 0x73, 0x77,
		0x69, 0x73, 0x73, 0x79, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0C, 0x3D, 0x42,
		0x43, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	testRuleInfoData = []byte{
		0xFF, 0xFF, 0xFF, 0xFF, 0x45, 0x2B, 0x00, 0x62, 0x6F, 0x74, 0x5F, 0x6D,
		0x69, 0x6E, 0x70, 0x6C, 0x61, 0x79, 0x65, 0x72, 0x73, 0x00, 0x34, 0x00,
		0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x6C, 0x69, 0x6D, 0x69, 0x74,
		0x00, 0x38, 0x00, 0x64, 0x6D, 0x66, 0x6C, 0x61, 0x67, 0x73, 0x00, 0x30,
		0x00, 0x66, 0x72, 0x61, 0x67, 0x6C, 0x69, 0x6D, 0x69, 0x74, 0x00, 0x30,
		0x00, 0x67, 0x5F, 0x61, 0x64, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
		0x53, 0x63, 0x6F, 0x72, 0x65, 0x42, 0x6F, 0x6E, 0x75, 0x73, 0x00, 0x33,
		0x00, 0x67, 0x5F, 0x61, 0x64, 0x45, 0x6C, 0x69, 0x6D, 0x53, 0x63, 0x6F,
		0x72, 0x65, 0x42, 0x6F, 0x6E, 0x75, 0x73, 0x00, 0x32, 0x00, 0x67, 0x5F,
		0x61, 0x64, 0x54, 0x6F, 0x75, 0x63, 0x68, 0x53, 0x63, 0x6F, 0x72, 0x65,
		0x42, 0x6F, 0x6E, 0x75, 0x73, 0x00, 0x31, 0x00, 0x67, 0x5F, 0x62, 0x6C,
		0x75, 0x65, 0x53, 0x63, 0x6F, 0x72, 0x65, 0x00, 0x30, 0x00, 0x67, 0x5F,
		0x63, 0x75, 0x73, 0x74, 0x6F, 0x6D, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6E,
		0x67, 0x73, 0x00, 0x30, 0x00, 0x67, 0x5F, 0x66, 0x61, 0x63, 0x74, 0x6F,
		0x72, 0x79, 0x00, 0x64, 0x75, 0x65, 0x6C, 0x00, 0x67, 0x5F, 0x66, 0x61,
		0x63, 0x74, 0x6F, 0x72, 0x79, 0x54, 0x69, 0x74, 0x6C, 0x65, 0x00, 0x44,
		0x75, 0x65, 0x6C, 0x00, 0x67, 0x5F, 0x66, 0x72, 0x65, 0x65, 0x7A, 0x65,
		0x52, 0x6F, 0x75, 0x6E, 0x64, 0x44, 0x65, 0x6C, 0x61, 0x79, 0x00, 0x34,
		0x30, 0x30, 0x30, 0x00, 0x67, 0x5F, 0x67, 0x61, 0x6D, 0x65, 0x53, 0x74,
		0x61, 0x74, 0x65, 0x00, 0x49, 0x4E, 0x5F, 0x50, 0x52, 0x4F, 0x47, 0x52,
		0x45, 0x53, 0x53, 0x00, 0x67, 0x5F, 0x67, 0x61, 0x6D, 0x65, 0x74, 0x79,
		0x70, 0x65, 0x00, 0x31, 0x00, 0x67, 0x5F, 0x67, 0x72, 0x61, 0x76, 0x69,
		0x74, 0x79, 0x00, 0x38, 0x30, 0x30, 0x00, 0x67, 0x5F, 0x69, 0x6E, 0x73,
		0x74, 0x61, 0x47, 0x69, 0x62, 0x00, 0x30, 0x00, 0x67, 0x5F, 0x69, 0x74,
		0x65, 0x6D, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x00, 0x31, 0x35, 0x00,
		0x67, 0x5F, 0x69, 0x74, 0x65, 0x6D, 0x54, 0x69, 0x6D, 0x65, 0x72, 0x73,
		0x00, 0x30, 0x00, 0x67, 0x5F, 0x6C, 0x65, 0x76, 0x65, 0x6C, 0x53, 0x74,
		0x61, 0x72, 0x74, 0x54, 0x69, 0x6D, 0x65, 0x00, 0x31, 0x34, 0x34, 0x35,
		0x38, 0x39, 0x37, 0x33, 0x33, 0x35, 0x00, 0x67, 0x5F, 0x6C, 0x6F, 0x61,
		0x64, 0x6F, 0x75, 0x74, 0x00, 0x30, 0x00, 0x67, 0x5F, 0x6E, 0x65, 0x65,
		0x64, 0x70, 0x61, 0x73, 0x73, 0x00, 0x30, 0x00, 0x67, 0x5F, 0x6F, 0x76,
		0x65, 0x72, 0x74, 0x69, 0x6D, 0x65, 0x00, 0x31, 0x32, 0x30, 0x00, 0x67,
		0x5F, 0x71, 0x75, 0x61, 0x64, 0x44, 0x61, 0x6D, 0x61, 0x67, 0x65, 0x46,
		0x61, 0x63, 0x74, 0x6F, 0x72, 0x00, 0x33, 0x00, 0x67, 0x5F, 0x72, 0x65,
		0x64, 0x53, 0x63, 0x6F, 0x72, 0x65, 0x00, 0x30, 0x00, 0x67, 0x5F, 0x72,
		0x6F, 0x75, 0x6E, 0x64, 0x57, 0x61, 0x72, 0x6D, 0x75, 0x70, 0x44, 0x65,
		0x6C, 0x61, 0x79, 0x00, 0x31, 0x30, 0x30, 0x30, 0x30, 0x00, 0x67, 0x5F,
		0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6E, 0x67, 0x48, 0x65, 0x61, 0x6C,
		0x74, 0x68, 0x00, 0x31, 0x30, 0x30, 0x00, 0x67, 0x5F, 0x74, 0x65, 0x61,
		0x6D, 0x46, 0x6F, 0x72, 0x63, 0x65, 0x42, 0x61, 0x6C, 0x61, 0x6E, 0x63,
		0x65, 0x00, 0x31, 0x00, 0x67, 0x5F, 0x74, 0x65, 0x61, 0x6D, 0x53, 0x69,
		0x7A, 0x65, 0x4D, 0x69, 0x6E, 0x00, 0x31, 0x00, 0x67, 0x5F, 0x74, 0x69,
		0x6D, 0x65, 0x6F, 0x75, 0x74, 0x43, 0x6F, 0x75, 0x6E, 0x74, 0x00, 0x30,
		0x00, 0x67, 0x5F, 0x76, 0x6F, 0x74, 0x65, 0x46, 0x6C, 0x61, 0x67, 0x73,
		0x00, 0x31, 0x33, 0x33, 0x32, 0x30, 0x00, 0x67, 0x5F, 0x77, 0x65, 0x61,
		0x70, 0x6F, 0x6E, 0x72, 0x65, 0x73, 0x70, 0x61, 0x77, 0x6E, 0x00, 0x35,
		0x00, 0x6D, 0x61, 0x70, 0x6E, 0x61, 0x6D, 0x65, 0x00, 0x75, 0x73, 0x65,
		0x61, 0x6E, 0x64, 0x61, 0x62, 0x75, 0x73, 0x65, 0x00, 0x6D, 0x65, 0x72,
		0x63, 0x79, 0x6C, 0x69, 0x6D, 0x69, 0x74, 0x00, 0x30, 0x00, 0x70, 0x72,
		0x6F, 0x74, 0x6F, 0x63, 0x6F, 0x6C, 0x00, 0x39, 0x31, 0x00, 0x72, 0x6F,
		0x75, 0x6E, 0x64, 0x6C, 0x69, 0x6D, 0x69, 0x74, 0x00, 0x31, 0x30, 0x00,
		0x72, 0x6F, 0x75, 0x6E, 0x64, 0x74, 0x69, 0x6D, 0x65, 0x6C, 0x69, 0x6D,
		0x69, 0x74, 0x00, 0x31, 0x38, 0x30, 0x00, 0x73, 0x63, 0x6F, 0x72, 0x65,
		0x6C, 0x69, 0x6D, 0x69, 0x74, 0x00, 0x31, 0x35, 0x30, 0x00, 0x73, 0x76,
		0x5F, 0x68, 0x6F, 0x73, 0x74, 0x6E, 0x61, 0x6D, 0x65, 0x00, 0x4E, 0x41,
		0x4B, 0x45, 0x44, 0x20, 0x2D, 0x20, 0x44, 0x75, 0x65, 0x6C, 0x20, 0x23,
		0x32, 0x20, 0x2D, 0x20, 0x46, 0x72, 0x61, 0x6E, 0x6B, 0x66, 0x75, 0x72,
		0x74, 0x00, 0x73, 0x76, 0x5F, 0x6D, 0x61, 0x78, 0x63, 0x6C, 0x69, 0x65,
		0x6E, 0x74, 0x73, 0x00, 0x31, 0x38, 0x00, 0x73, 0x76, 0x5F, 0x70, 0x72,
		0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6C, 0x69, 0x65, 0x6E, 0x74, 0x73,
		0x00, 0x34, 0x00, 0x74, 0x65, 0x61, 0x6D, 0x73, 0x69, 0x7A, 0x65, 0x00,
		0x30, 0x00, 0x74, 0x69, 0x6D, 0x65, 0x6C, 0x69, 0x6D, 0x69, 0x74, 0x00,
		0x31, 0x30, 0x00, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6F, 0x6E, 0x00, 0x31,
		0x30, 0x36, 0x33, 0x20, 0x6C, 0x69, 0x6E, 0x75, 0x78, 0x2D, 0x78, 0x36,
		0x34, 0x20, 0x4F, 0x63, 0x74, 0x20, 0x32, 0x33, 0x20, 0x32, 0x30, 0x31,
		0x35, 0x20, 0x31, 0x37, 0x3A, 0x30, 0x34, 0x3A, 0x31, 0x31, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
)