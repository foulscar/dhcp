package dhcp

type OptionDataParameterRequestList struct {
	List []OptionCode
}

func (optD OptionDataParameterRequestList) Raw() []byte {
	data := make([]byte, len(optD.List))
	for i, optC := range optD.List {
		data[i] = byte(optC)
	}

	return data
}

func (optD OptionDataParameterRequestList) Add(optC OptionCode) {
        optD.List = append(optD.List, optC)
}

func MarshalOptionDataParameterRequestList(data []byte) (OptionData, error) {
	list := make([]OptionCode, len(data))
	for i, b := range data {
		list[i] = OptionCode(b)
	}
	return OptionDataParameterRequestList{List: list}, nil
}

func NewOptionParameterRequestList(optCodes ...OptionCode) Option {
        opt := Option{
                Code: OptionCodeParameterRequestList,
                Data: OptionDataParameterRequestList{
                        List: make([]OptionCode, 0),
                },
        }

        for _, optCode := range optCodes {
                opt.Data.(OptionDataParameterRequestList).Add(optCode)
        }

        return opt
}
