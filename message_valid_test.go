package dhcp

import "testing"

func TestIsEncodedMessage(t *testing.T) {
	dataShort := make([]byte, 299)
	dataNoCookie := make([]byte, 300)
	dataWithCookie := make([]byte, 300)
	copy(dataWithCookie[236:300], MagicCookie)

	var tests = []struct {
		name string
		data []byte
		want bool
	}{
		{"Data with len < 300 cannot be a message", dataShort, false},
		{"Data with no MagicCookie cannot be a message", dataNoCookie, false},
		{"Data with len >= 300 and with MagicCookie could be a message", dataWithCookie, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			answer := IsEncodedMessage(tt.data)
			if answer != tt.want {
				t.Errorf("got %t, want %t", answer, tt.want)
			}
		})
	}
}
