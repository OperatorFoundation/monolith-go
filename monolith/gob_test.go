package monolith

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytePart_GobMarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bp := BytesPart{Items:[]Monolith{FixedByteType{Byte: 0x0A}}}

	result, marshalError := gob.Marshal(bp)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println("result", string(result))

	bp2, unmarshalError := gob.Unmarshal(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bp, bp2) {
		t.Fail()
		return
	}
}
