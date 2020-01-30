package monolith

type TimedPart struct {
	milliseconds uint
	Items []ByteType
}

type TimedMessage struct {
	milliseconds uint
	bytes []byte
}

func (message TimedMessage) Bytes() []byte {
	return message.bytes
}

func (part TimedPart) MessageFromArgs(args []interface{}) (Message, []interface{}) {
	var b []byte
	var result []byte

	for index := 0; index < len(part.Items); index++ {
		b, args = part.Items[index].BytesFromArgs(args)
		result = append(result, b...)
	}

	m := TimedMessage{
		milliseconds: part.milliseconds,
		bytes:        result,
	}

	return m, args
}

func (part TimedPart) ArgsFromBytes(bs []byte, args []interface{}) ([]byte, []interface{}) {
	resultBytes := bs
	resultArgs := args

	for index := 0; index < len(part.Items); index++ {
		resultBytes, resultArgs = part.Items[index].ArgsFromBytes(resultBytes, resultArgs)
	}

	return resultBytes, resultArgs
}

func (part TimedPart) Validate(bs []byte) ([]byte, Validity) {
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
