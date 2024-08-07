package socks

// Socks4AHeader represents the header of SOCKS4A protocol
type Socks4AHeader struct {
	Version  uint8   // Version number (always 4)
	Command  uint8   // Command code
	Port     uint16  // Destination port
	IPAddr   [4]byte // IPv4 address
	UserID   []byte  // User ID
	NullByte uint8   // Null byte (always 0)
	Domain   []byte  // Domain name
}

// Socks5Header represents the header of SOCKS5 protocol
type Socks5Header struct {
	Version     uint8  // Version number (always 5)
	Command     uint8  // Command code
	Reserved    uint8  // Reserved (must be 0x00)
	AddressType uint8  // Address type
	Destination string // Destination address (IPv4, IPv6, or domain name)
	Port        uint16 // Destination port
}

func (s5 *Socks5Header) Check() bool {
	if s5.Version != 0x05 {
		return false
	}
	if s5.Command < 0x01 || s5.Command > 0x03 {
		return false
	}
	if s5.Reserved != 0 {
		return false
	}
	if s5.AddressType != 0x01 && s5.AddressType != 0x03 && s5.AddressType != 0x04 {
		return false
	}
	return true

}

func (s4 *Socks4AHeader) Check() bool {
	if s4.Version != 0x04 {
		return false
	}
	if s4.Command != 0x01 && s4.Command != 0x02 {
		return false
	}
	if s4.NullByte != 0x00 {
		return false
	}
	return true
}
