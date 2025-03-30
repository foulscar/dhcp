package dhcp

import "strings"

// ErrorExt implements the error interface.
// I added this because Message/Options errors became very verbose and I
// wanted a better way of storing/reading the data.
// See ErrorExt.JSON()
type ErrorExt struct {
	// Main contains the main/root error message
	Main string
	// Details provides a tree of errors.
	// Can be nil
	Details []ErrorExt
}

// Error returns a string with a typical error message
func (e ErrorExt) Error() string {
	sb := strings.Builder{}
	sb.WriteString(e.Main)

	if e.Details == nil || len(e.Details) < 1 {
		return sb.String()
	}

	sb.WriteString(": [")
	for i, err := range e.Details {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Error())
	}
	sb.WriteString("]")

	return sb.String()
}

// JSON returns a string in JSON format
func (e ErrorExt) JSON() string {
	sb := strings.Builder{}

	sb.WriteString("{")

	sb.WriteString(`"error":` + `"` + e.Main + `",`)
	sb.WriteString(`"details":[`)
	if e.Details != nil && len(e.Details) > 0 {
		for i, err := range e.Details {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(err.JSON())
		}
	}
	sb.WriteString("]")

	sb.WriteString("}")

	return sb.String()
}

// Add adds details to an ErrorExt
func (e ErrorExt) Add(details ...ErrorExt) {
	e.Details = append(e.Details, details...)
}

// NewErrorExt constructs an ErrorExt with optional details
func NewErrorExt(mainMsg string, details ...ErrorExt) ErrorExt {
	return ErrorExt{
		Main:    mainMsg,
		Details: details,
	}
}
