package dhcp

import "strconv"

// OptionDataDefault exists for combatibility.
// It will simply hold raw data, nothing more
type OptionDataDefault struct {
	Data []byte
}

// String returns "x bytes", where x is the number of bytes stored in optD.Data.
// OptionDataDefault exists for combatibility and is not verbose
func (optD OptionDataDefault) String() string {
	return strconv.Itoa(len(optD.Data)) + " bytes"
}

// IsValid always returns true.
// OptionDataDefault exists for combatibility and will not check if your data is valid
func (optD OptionDataDefault) IsValid() bool { return true }

// Marshal returns optD.Data and will never return an error.
// OptionDataDefault exists for combatibility and will not check if your data is valid
func (optD OptionDataDefault) Marshal() ([]byte, error) { return optD.Data, nil }

// UnmarshalOptionDefault constructs OptionDataDefault from your data and will never return an error.
// OptionDataDefault exists for combatibility and will not check if your data is valid
func UnmarshalOptionDefault(data []byte) (OptionData, error) {
	return OptionDataDefault{Data: data}, nil
}
