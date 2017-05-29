package UDPClient

import (
	. "github.com/lflxp/arm/udp"
)

func DhcpClient() {
	for i,tt := range optionsTests {
		p := NewPacket(BootRequest)
		for _,o := range tt.options {
			p.AddOption(o.Code,o.Value)
		}
		println(i,p)
		data := Broadcast{Net:"udp4",Port:67}
		data.GetConnect()
		data.Discover(p)
		defer data.Close()
	}
}


// oneOptionSlice is a test helper of []Option with a single
// Option.
var oneOptionSlice = []Option{
	Option{
		Code:  OptionSubnetMask,
		Value: []byte{255, 255, 255, 0},
	},
}

// twoOptionSlice is a test helper of []Option with two
// Option values.
var twoOptionsSlice = []Option{
	Option{
		Code:  OptionSubnetMask,
		Value: []byte{255, 255, 255, 0},
	},
	Option{
		Code:  OptionDomainNameServer,
		Value: []byte{8, 8, 8, 8},
	},
}

// optionsTests are tests used when applying and stripping Options
// from Packets.
var optionsTests = []struct {
	description string
	options     []Option
}{
	{
		description: "nil options",
		options:     nil,
	},
	{
		description: "empty options",
		options:     []Option{},
	},
	{
		description: "padding option",
		options: []Option{
			Option{
				Code: Pad,
			},
		},
	},
	{
		description: "one option",
		options:     oneOptionSlice,
	},
	{
		description: "two options",
		options:     twoOptionsSlice,
	},
	{
		description: "four options",
		options: []Option{
			Option{
				Code:  OptionSubnetMask,
				Value: []byte{255, 255, 255, 0},
			},
			Option{
				Code:  OptionDomainNameServer,
				Value: []byte{8, 8, 8, 8},
			},
			Option{
				Code:  OptionTimeServer,
				Value: []byte{127, 0, 0, 1},
			},
			Option{
				Code:  OptionMessage,
				Value: []byte{'h', 'e', 'l', 'l', 'o', 'w', 'o', 'r', 'l', 'd'},
			},
		},
	},
}

