package monolith

import "github.com/deckarep/golang-set"

type Validateable interface {
	Validate(bs []byte) ([]byte, bool)
}

func (desc Description) Validate(bs []byte) ([]byte, bool) {
	bs2 := bs
	valid := true
	for _, part := range desc.Parts {
		bs2, valid = part.Validate(bs2)
		if !valid {
			return bs2, false
		}
	}

	return bs2, true
}

func (part BytesPart) Validate(bs []byte) ([]byte, bool) {
	bs2 := bs
	valid := true

	for index := 0; index < len(part.Items); index++ {
		bs2, valid = part.Items[index].Validate(bs2)
		if !valid {
			return bs2, false
		}
	}

	return bs2, true
}

func (bt FixedByteType) Validate(bs []byte) ([]byte, bool) {
	if len(bs) == 0 {
		return bs, false
	}

	b, bs2 := bs[0], bs[1:]

	return bs2, b == bt.Byte
}

func (bt EnumeratedByteType) Validate(bs []byte) ([]byte, bool) {
	if len(bs) == 0 {
		return bs, false
	}

	b, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.Options))
	for index, option := range options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)
	return bs2, set.Contains(b)
}

func (bt RandomByteType) Validate(bs []byte) ([]byte, bool) {
	if len(bs) == 0 {
		return bs, false
	}

	_, bs2 := bs[0], bs[1:]

	return bs2, true
}

func (bt RandomEnumeratedByteType) Validate(bs []byte) ([]byte, bool) {
	if len(bs) == 0 {
		return bs, false
	}

	b, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.RandomOptions))
	for index, option := range options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)
	return bs2, set.Contains(b)
}

