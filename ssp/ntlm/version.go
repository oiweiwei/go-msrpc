package ntlm

const (
	WindowsMajorVersion5  = 0x05
	WindowsMajorVersion6  = 0x06
	WindowsMajorVersion10 = 0x0a

	WindowsMinorVersion0 = 0x00
	WindowsMinorVersion1 = 0x01
	WindowsMinorVersion2 = 0x02
	WindowsMinorVersion3 = 0x03
)

type ProductVersion struct {
	// An 8-bit unsigned integer that SHOULD contain the major version number
	// of the operating system in use.
	Major uint8
	// An 8-bit unsigned integer that SHOULD contain the minor version number
	// of the operating system in use.
	Minor uint8
}

var (
	// Windows XP ServicePack 2.
	WindowsXPSP2 = ProductVersion{WindowsMajorVersion5, WindowsMinorVersion1}
	// Windows Vista.
	WindowsVista = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion0}
	// Windows 7.
	Windows7 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion1}
	// Windows 8.
	Windows8 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion2}
	// Windows 8.1
	Windows8_1 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion3}
	// Windows 10.
	Windows10 = ProductVersion{WindowsMajorVersion10, WindowsMinorVersion0}
	// Windows Server 2003.
	WindowsServer2003 = ProductVersion{WindowsMajorVersion5, WindowsMinorVersion2}
	// Windows Server 2008.
	WindowsServer2008 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion0}
	// Windows Server 2008 R2.
	WindowsServer2008R2 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion1}
	// WindowsServer 2012.
	WindowsServer2012 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion2}
	// Windows Server 2012 R2.
	WindowsServer2012R2 = ProductVersion{WindowsMajorVersion6, WindowsMinorVersion3}
	// Windows Server 2016/2019.
	WindowsServer2016_2019 = ProductVersion{WindowsMajorVersion10, WindowsMinorVersion0}
)

// An 8-bit unsigned integer that contains a value indicating the
// revision of the NTLMSSP in use.
type Revision uint8

const (
	// NTLMSSP_REVISION_W2K3. Version 15 of the NTLMSSP is in use.
	RevisionCurrent Revision = 0x0f
)

type Version struct {
	// The product version.
	ProductVersion ProductVersion
	// A 16-bit unsigned integer that contains the build number of the
	// operating system in use. This field SHOULD be set to a 16-bit
	// quantity that identifies the operating system build number.
	ProductBuild uint16
	// A 24-bit data area that SHOULD be set to zero and MUST be
	// ignored by the recipient.
	_ [3]byte
	// An 8-bit unsigned integer that contains a value indicating
	// the current revision of the NTLMSSP in use. This field
	// SHOULD contain the following value:
	// NTLMSSP_REVISION_W2K3: 0x0F (Version 15 of the NTLMSSP is in use).
	Revision Revision
}

// IsZero function returns `true` if version is empty.
func (v Version) IsZero() bool {
	return v == Version{}
}

var DefaultVersion = Version{ProductVersion: Windows10, ProductBuild: 19044, Revision: RevisionCurrent}
