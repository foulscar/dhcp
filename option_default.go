package dhcp

import "strconv"

type OptionDataDefault struct {
	Data []byte
}

func (optD OptionDataDefault) String() string {
	return strconv.Itoa(len(optD.Data)) + " bytes"
}

func (optD OptionDataDefault) IsValid() bool { return true }

func (optD OptionDataDefault) Marshal() ([]byte, error) { return optD.Data, nil }

func UnmarshalOptionDefault(data []byte) (OptionData, error) {
	return OptionDataDefault{Data: data}, nil
}
