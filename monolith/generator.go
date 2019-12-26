package monolith

import mapset "github.com/deckarep/golang-set"
import "math/rand"

type Instance struct {
	Desc Description
	Args []interface{}
}

func (i Instance) Bytes() []byte {
	b, _ := i.Desc.BytesFromArgs(i.Args)
	return b
}

func (desc Description) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	var b []byte
	result := make([]byte, 0)

	for _, part := range desc.Parts {
		b, args = part.BytesFromArgs(args)
		result = append(result, b...)
	}

	return result, args
}

func (part BytesPart) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	var b []byte
	result := make([]byte, 0)

	for index := 0; index < len(part.Items); index++ {
		b, args = part.Items[index].BytesFromArgs(args)
		result = append(result, b...)
	}

	return result, args
}

func (bt FixedByteType) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	result := make([]byte, 0)
	result = append(result, bt.Byte)
	return result, args
}

func (bt EnumeratedByteType) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	var arg interface{}
	var b byte

	arg, args = args[0], args[1:]
	switch arg.(type) {
	case uint8:
		b = arg.(byte)
	case int:
		b = byte(arg.(int))
	default:
		return nil, args
	}

	options := make([]interface{}, len(bt.Options))
	for index, option := range bt.Options {
		options[index] = option
	}
	set := mapset.NewSetFromSlice(options)
	if set.Contains(b) {
		result := make([]byte, 0)
		result = append(result, b)
		return result, args
	} else {
		return nil, args
	}
}

func (bt RandomByteType) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	bs := make([]byte, 1)
	rand.Read(bs)

	result := make([]byte, 0)
	result = append(result, bs[0])
	return result, args
}

func (bt RandomEnumeratedByteType) BytesFromArgs(args []interface{}) ([]byte, []interface{}) {
	index := rand.Intn(len(bt.RandomOptions))
	b := bt.RandomOptions[index]

	result := make([]byte, 0)
	result = append(result, b)
	return result, args
}