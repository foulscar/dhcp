package dhcp

import (
	"net"
	"reflect"
	"testing"
)

func TestMarshalOptions(t *testing.T) {
	testOpts := make(Options)
	testOpts[OptionCodeSubnetMask] = Option{
		Code: OptionCodeSubnetMask,
		Data: OptionDataSubnetMask{Mask: net.IPMask([]byte{0xff, 0xff, 0xff, 0x00})},
	}

	testData := []byte{
		byte(OptionCodePad),
		byte(OptionCodeSubnetMask), byte(4),
		0xff, 0xff, 0xff, 0x00,
		byte(OptionCodeEnd),
	}

	opts, _ := MarshalOptions(testData)
	if !reflect.DeepEqual(opts, testOpts) {
		t.Fatalf("test failed")
	}
}
