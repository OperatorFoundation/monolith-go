package monolith

import "github.com/deckarep/golang-set"

type Parseable interface {
	ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{})
}

func (description Description) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	resultBytes := bs
	resultArgs := args

	for _, part := range description.Parts {
		resultBytes, resultArgs = part.ArgsFromBytes(resultBytes, resultArgs)
	}

	return resultBytes, resultArgs
}

func (part BytesPart) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	resultBytes := bs
	resultArgs := args

	for index := 0; index < len(part.Items); index++ {
		resultBytes, resultArgs = part.Items[index].ArgsFromBytes(resultBytes, resultArgs)
	}

	return resultBytes, resultArgs
}

func (bt FixedByteType) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	if len(bs) == 0 {
		return bs, args
	}

	_, bs2 := bs[0], bs[1:]

	return bs2, args
}

func (bt EnumeratedByteType) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	if len(bs) == 0 {
		return bs, args
	}

	arg, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.Options))
	for index, option := range options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)
	if set.Contains(arg) {
		return bs2, append(args, arg)
	} else {
		return bs2, args
	}
}

func (bt RandomByteType) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	if len(bs) == 0 {
		return bs, args
	}

	_, bs2 := bs[0], bs[1:]

	return bs2, args
}

func (bt RandomEnumeratedByteType) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	if len(bs) == 0 {
		return bs, args
	}

	arg, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.RandomOptions))
	for index, option := range options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)
	if set.Contains(arg) {
		return bs2, append(args, arg)
	} else {
		return bs2, args
	}
}

