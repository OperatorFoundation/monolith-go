package monolith

import mapset "github.com/deckarep/golang-set"

type Description struct {
	Parts []Byteable
}

func (desc Description) Bytes() []byte {
	result := make([]byte, 0)

	for _, part := range desc.Parts {
		result = append(result, part.Bytes()...)
	}

	return result
}

type Byteable interface {
	Bytes() []byte
}

type BytesPart struct {
	Items []Byteable
}

func (part BytesPart) Bytes() []byte {
	result := make([]byte, 0)

	for index := 0; index < len(part.Items); index++ {
		result = append(result, part.Items[index].Bytes()...)
	}

	return result
}

type ByteType interface {
	Bytes() []byte
}

type FixedByteType struct {
	Byte byte
}

func (bt FixedByteType) Bytes() []byte {
	result := make([]byte, 0)
	result = append(result, bt.Byte)
	return result
}

type EnumeratedByteType struct {
	Options []byte
}

func (bt EnumeratedByteType) Bytes() []byte {
	return
}

type RandomByteType struct {

}

func (bt RandomByteType) Bytes() []byte {
	return nil
}

type SemanticByteType struct {

}

func (bt SemanticByteType) Bytes() []byte {
	return nil
}