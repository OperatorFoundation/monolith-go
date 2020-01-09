package monolith

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
)

func TestDescription_GobMarshalUnmarshal(t *testing.T) {
	InitializeGobRegistry()

	parts := make([]Monolith, 0)
	part := BytesPart{
		Items: []Monolith{FixedByteType{Byte: 0x0A}},
	}
	parts = append(parts, part)
	desc := Description{parts}

	args := make([]interface{}, 0)

	instance := Instance{
		Desc: desc,
		Args: args,
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	marshalError := encoder.Encode(instance)
	if marshalError != nil {
		t.Fail()
		return
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(buffer.Bytes()))
	fmt.Println("result", encoded)

	decoder := gob.NewDecoder(&buffer)

	var i2 Instance
	unmarshalError := decoder.Decode(&i2)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(instance, i2) {
		fmt.Println("ne", instance, i2)
		fmt.Println(instance.Args, i2.Args, reflect.DeepEqual(instance.Args, i2.Args))
		t.Fail()
		return
	}
}

func TestBytePart_GobMarshalUnmarshal(t *testing.T) {
	InitializeGobRegistry()

	bp := BytesPart{Items:[]Monolith{FixedByteType{Byte: 0x0A}}}

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	marshalError := encoder.Encode(bp)
	if marshalError != nil {
		t.Fail()
		return
	}

	fmt.Println("result", string(result.Bytes()))

	decoder := gob.NewDecoder(&result)

	var bp2 BytesPart
	unmarshalError := decoder.Decode(&bp2)
	if unmarshalError != nil {
		t.Fail()
		return
	}

	if !reflect.DeepEqual(bp, bp2) {
		fmt.Println("ne", bp, bp2)
		t.Fail()
		return
	}
}
