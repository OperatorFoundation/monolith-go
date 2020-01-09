package monolith

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBytePart_MarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bp := BytesPart{Items:[]Monolith{FixedByteType{Byte: 0x0A}}}

	result, marshalError := MarshalTypedJSON(bp)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println("result", string(result))

	bp2, unmarshalError := UnmarshalTypedJSON(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bp, bp2) {
		t.Fail()
		return
	}
}

func TestFixedByteType_MarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bt := FixedByteType{Byte: 0x0A}

	result, marshalError := MarshalTypedJSON(bt)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println(string(result))

	bt2, unmarshalError := UnmarshalTypedJSON(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bt, bt2) {
		t.Fail()
		return
	}
}

func TestEnumeratedByteType_MarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bt := EnumeratedByteType{Options: []byte{0x00, 0x0A}}

	result, marshalError := MarshalTypedJSON(bt)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println(string(result))

	bt2, unmarshalError := UnmarshalTypedJSON(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bt, bt2) {
		t.Fail()
		return
	}
}

func TestRandomByteType_MarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bt := RandomByteType{}

	result, marshalError := MarshalTypedJSON(bt)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println(string(result))

	bt2, unmarshalError := UnmarshalTypedJSON(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bt, bt2) {
		t.Fail()
		return
	}
}

func TestRandomEnumeratedByteType_MarshalUnmarshal(t *testing.T) {
	InitializeRegistry()

	bt := RandomEnumeratedByteType{RandomOptions: []byte{0x00, 0x0A}}

	result, marshalError := MarshalTypedJSON(bt)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println(string(result))

	bt2, unmarshalError := UnmarshalTypedJSON(result)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(&bt, bt2) {
		fmt.Println(bt, bt2)
		t.Fail()
		return
	}
}
