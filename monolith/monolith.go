package monolith

type Monolith interface {
	Byteable
	Parseable
	Validateable
}

type Description struct {
	Parts []Monolith
}

type Byteable interface {
	BytesFromArgs(args []interface{}) ([]byte, []interface{})
}

type BytesPart struct {
	Items []Monolith
}

type ByteType interface {
	Bytes() []byte
}

type FixedByteType struct {
	Byte byte
}

type EnumeratedByteType struct {
	Options []byte
}

type RandomByteType struct {

}

type RandomEnumeratedByteType struct {
	RandomOptions []byte
}

type SemanticByteType struct {

}
