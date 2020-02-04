package monolith

import "github.com/deckarep/golang-set"

type Validity int

const (
	Valid Validity = iota
	Invalid
	Incomplete
)

func (v Validity) String() string {
	return [...]string{"Valid", "Invalid", "Incomplete"}[v]
}

type Validateable interface {
	Validate(bs []byte) ([]byte, Validity)
}

func (description Description) Validate(bs []byte) ([]byte, Validity) {
	bs2 := bs
	valid := Valid
	for _, part := range description.Parts {
		bs2, valid = part.Validate(bs2)
		switch valid {
		case Valid:
			continue
		case Invalid:
			return bs2, Invalid
		case Incomplete:
			return bs2, Incomplete
		default:
			return bs2, Invalid
		}
	}

	return bs2, Valid
}

func (part BytesPart) Validate(bs []byte) ([]byte, Validity) {
	bs2 := bs
	valid := Valid

	for index := 0; index < len(part.Items); index++ {
		bs2, valid = part.Items[index].Validate(bs2)
		switch valid {
		case Valid:
			continue
		case Invalid:
			return bs2, Invalid
		case Incomplete:
			return bs2, Incomplete
		default:
		}
	}

	return bs2, Valid
}

func (bt FixedByteType) Validate(bs []byte) ([]byte, Validity) {
	if len(bs) == 0 {
		return bs, Incomplete
	}

	b, bs2 := bs[0], bs[1:]

	if b == bt.Byte {
		return bs2, Valid
	} else {
		return bs2, Invalid
	}
}

func (bt EnumeratedByteType) Validate(bs []byte) ([]byte, Validity) {
	if len(bs) == 0 {
		return bs, Incomplete
	}

	b, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.Options))
	for index, option := range bt.Options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)
	if set.Contains(b) {
		return bs2, Valid
	} else {
		return bs2, Invalid
	}
}

func (bt RandomByteType) Validate(bs []byte) ([]byte, Validity) {
	if len(bs) == 0 {
		return bs, Incomplete
	}

	_, bs2 := bs[0], bs[1:]

	return bs2, Valid
}

func (bt RandomEnumeratedByteType) Validate(bs []byte) ([]byte, Validity) {
	if len(bs) == 0 {
		return bs, Incomplete
	}

	b, bs2 := bs[0], bs[1:]

	options := make([]interface{}, len(bt.RandomOptions))
	for index, option := range options {
		options[index] = option
	}

	set := mapset.NewSetFromSlice(options)

	if set.Contains(b) {
		return bs2, Valid
	} else {
		return bs2, Invalid
	}
}

