package monolith

//make stringsPart struct
//make it conform to monolith interfaces

type StringType interface {
	Validateable
	Parseable
	Countable
	StringFromArgs(args *Args, context *Context) (string, error)
}

type StringMessage struct {
	String string
}

func (s StringMessage) Bytes() []byte {
	return []byte(s.String)
}

type StringsPart struct {
	Items []StringType
}

type FixedStringType struct {
	String string
}

type VariableStringType struct {
	EndDelimiter byte
}
