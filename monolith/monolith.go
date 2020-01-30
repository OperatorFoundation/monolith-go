package monolith

type Monolith interface {
	Messageable
	Parseable
	Validateable
}

type Description struct {
	Parts []Monolith
}

type Byteable interface {
	Bytes() []byte
}

type Messageable interface {
	MessageFromArgs(args []interface{}) (Message, []interface{})
}

type BytesPart struct {
	Items []ByteType
}

type ByteType interface {
	Validateable
	Parseable
	BytesFromArgs(args []interface{}) ([]byte, []interface{})
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

type Message interface {
	Byteable
}

type BytesMessage struct {
	bytes []byte
}

func (message BytesMessage) Bytes() []byte {
	return message.bytes
}
