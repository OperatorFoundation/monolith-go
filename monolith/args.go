package monolith

import "errors"

type Args struct {
	values []interface{}
	index int
}

func NewEmptyArgs() *Args {
	return &Args{
		values: make([]interface{}, 0),
		index: 0,
	}
}

func NewArgs(values []interface{}) *Args {
	return &Args{
		values: values,
		index:  0,
	}
}

func (args Args) Empty() bool {
	return len(args.values) <= 0
}

func (args Args) Pop() (interface{}, error) {
	if len(args.values) > 0 {
		value, rest := args.values[0], args.values[1:]
		args.values = rest
		return value, nil
	} else {
		return nil, errors.New("not enough args")
	}
}

func (args Args) PopInt() (int, error) {
	value, popError := args.Pop()
	if popError != nil {
		return 0, popError
	}

	if n, ok := value.(int); ok {
		return n, nil
	} else {
		return 0, errors.New("value was not an int")
	}
}

func (args Args) PopByte() (byte, error) {
	value, popError := args.Pop()
	if popError != nil {
		return 0, popError
	}

	if b, ok := value.(byte); ok {
		return b, nil
	}

	if n, ok := value.(int); ok {
		return byte(n), nil
	}

	return 0, errors.New("value was not an int")
}

func (args Args) Push(value interface{}) {
	args.values = append(args.values, value)
}