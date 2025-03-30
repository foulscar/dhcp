package dhcp

import "strings"

// ErrorExt implements the error interface.
// I added this because Message/Options errors became very verbose and I
// wanted a better way of storing/reading the data.
// You as the user do NOT have to use this, but you can if you want to :)
// See ErrorExt.JSON()
type ErrorExt struct {
	// Main contains the main/root error message
	Main string
	// Reasons provides a list of reasons (ErrorExt) on why the error occurred.
	Reasons []ErrorExt
}

// Error returns a string with a typical error message
func (e ErrorExt) Error() string {
	sb := strings.Builder{}
	sb.WriteString(e.Main)

	if !e.HasReasons() {
		return sb.String()
	}

	sb.WriteString(": [")
	for i, err := range e.Reasons {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(err.Error())
	}
	sb.WriteString("]")

	return sb.String()
}

// String just calls ErrorExt.Error()
func (e ErrorExt) String() string {
	return e.Error()
}

// JSON returns a string in JSON format
func (e ErrorExt) JSON() string {
	sb := strings.Builder{}

	sb.WriteString("{")

	sb.WriteString(`"error":` + `"` + e.Main + `"`)
	if e.HasReasons() {
		sb.WriteString(`,"reasons":[`)
		for i, err := range e.Reasons {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(err.JSON())
		}
		sb.WriteString("]")
	}

	sb.WriteString("}")

	return sb.String()
}

// Add adds reasons to an ErrorExt with args as multiple *ErrorExt, error, or string arguments
func (e ErrorExt) Add(args ...any) {
	for _, arg := range args {
		err := handleErrorExtContextArg(arg)
		if err == nil {
			continue
		}
		e.Reasons = append(e.Reasons, *err)
	}
}

// HasReasons returns true if the ErrorExt contains reasons on why the error occurred
func (e ErrorExt) HasReasons() bool {
	if e.Reasons == nil || len(e.Reasons) == 0 {
		return false
	}
	return true
}

// NewErrorExt constructs an ErrorExt with args as multiple *ErrorExt, error or string arguments.
// The first argument is what determines the main message (can also provide details if using *ErrorExt)
// Returns nil if no arguments are given or the determined main message is an empty string
func NewErrorExt(args ...any) *ErrorExt {
	if len(args) == 0 {
		return nil
	}

	var mainMsg string
	newDetails := make([]ErrorExt, 0)

	for i, arg := range args {
		err := handleErrorExtContextArg(arg)

		if i == 0 {
			mainMsg = err.Main
			newDetails = append(newDetails, err.Reasons...)
			continue
		}
		if err == nil {
			continue
		}

		newDetails = append(newDetails, *err)
	}

	if mainMsg == "" {
		return nil
	}

	return &ErrorExt{
		Main:    mainMsg,
		Reasons: newDetails,
	}
}

func handleErrorExtContextArg[T any](arg T) *ErrorExt {
	switch v := any(arg).(type) {
	case *ErrorExt:
		return v
	case error:
		return &ErrorExt{Main: v.Error(), Reasons: make([]ErrorExt, 0)}
	case string:
		return &ErrorExt{Main: v, Reasons: make([]ErrorExt, 0)}
	}

	return nil
}
