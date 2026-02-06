package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrRuntimeint32 = errors.New("Unsupported Error")

func (r *Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedValue := strconv.Quote(jsonValue)

	return []byte(quotedValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	unquotedJSONvalue, err := strconv.Unquote(string(jsonValue))

	if err != nil {
		return ErrRuntimeint32
	}
	parts := strings.Split(unquotedJSONvalue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrRuntimeint32
	}
	// Otherwise, parse the string containing the number into an int32. Again, if this
	// fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrRuntimeint32
	}
	// Convert the int32 to a Runtime type and assign this to the receiver. Note that we
	// use the * operator to deference the receiver (which is a pointer to a Runtime
	// type) in order to set the underlying value of the pointer.
	*r = Runtime(i)

	return nil
}
