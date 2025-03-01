package dhcp

type OptionDataDefault struct {
	Data []byte
}

func (optD OptionDataDefault) Raw() []byte {
	return optD.Data
}
