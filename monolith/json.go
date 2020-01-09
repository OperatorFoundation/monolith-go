package monolith

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

var registry map[string]reflect.Type

type Container struct {
	Typename string
	Value map[string]interface{}
}

func Register(i interface{}) {
	typevalue := reflect.TypeOf(i)
	name := typevalue.PkgPath()+"."+typevalue.Name()
	registry[name]=typevalue
}

func InitializeRegistry() {
	registry = make(map[string]reflect.Type)
	Register(BytesPart{})
	Register(FixedByteType{})
	Register(RandomByteType{})
	Register(EnumeratedByteType{})
	Register(RandomEnumeratedByteType{})
}

func MarshalTypedJSON(item interface{}) ([]byte, error) {
	switch item.(type) {
	case BytesPart:
		bp := item.(BytesPart)
		items := make([]interface{}, len(bp.Items))
		for index, item := range bp.Items {
			bytes, marshalError := GenericMarshalTypedJSON(item)
			if marshalError != nil {
				return nil, marshalError
			}

			var untyped map[string]interface{}
			unmarshalError := json.Unmarshal(bytes, &untyped)
			if unmarshalError != nil {
				return nil, unmarshalError
			}

			items[index] = untyped
		}
		container := map[string]interface{}{"Typename": "monolith.BytesPart", "Value": map[string]interface{}{"Items": items}}
		return json.Marshal(container)
	default:
		return GenericMarshalTypedJSON(item)
	}
}

func GenericMarshalTypedJSON(item interface{}) ([]byte, error) {
	result, marshalError := json.Marshal(item)
	if marshalError != nil {
		return nil, marshalError
	}

	var untyped map[string]interface{}
	unmarshalError := json.Unmarshal(result, &untyped)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	container := Container{
		Typename: reflect.TypeOf(item).String(),
		Value:    untyped,
	}

	fmt.Println(container)

	result, marshalError = json.Marshal(container)
	if marshalError != nil {
		return nil, marshalError
	}

	return result, nil
}

func UnmarshalTypedJSON(text []byte) (interface{}, error) {
	var container Container

	unmarshalError := json.Unmarshal(text, &container)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	switch container.Typename {
	case fullTypeName(BytesPart{}):
		return nil, errors.New("Bytes")
	default:
		return GenericUnmarshalTypedJSON(text)
	}
}

func fullTypeName(i interface{}) string {
	typevalue := reflect.TypeOf(i)
	return typevalue.PkgPath()+"."+typevalue.Name()
}

func GenericUnmarshalTypedJSON(text []byte) (interface{}, error) {
	var container Container

	unmarshalError := json.Unmarshal(text, &container)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	newtext, marshalError := json.Marshal(container.Value)
	if marshalError != nil {
		return nil, marshalError
	}

	typevalue, ok := registry[container.Typename]
	if !ok {
		fmt.Println(registry)
		return nil, errors.New("type not found")
	}

	btPtr := reflect.New(typevalue).Interface()
	unmarshalError = json.Unmarshal(newtext, btPtr)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return btPtr, nil
}

