package protobson_test

import (
	"github.com/custom-app/protobson"
	pbtest "github.com/custom-app/protobson/test"
	"testing"
)

func TestFieldNamerByNumber_FieldDescriptorToFieldName(t *testing.T) {
	fieldName := (&protobson.FieldNamerByNumber{}).FieldDescriptorToFieldName(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields().ByName("string_field"))
	if fieldName != "pb_field_1" {
		t.Errorf("field namer by number descriptor to field name: wrong result: %s!=pb_field_1", fieldName)
	}
}

func TestFieldNamerByNumber_FieldNameToFieldDescriptor(t *testing.T) {
	fieldD, err := (&protobson.FieldNamerByNumber{}).FieldNameToFieldDescriptor(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields(), "pb_field_2")
	if err != nil {
		t.Errorf("field namer by number field name to descriptor: err not nil: %v", err)
		return
	}
	if fieldD == nil {
		t.Errorf("field namer by number field name to descriptor: field descriptor is nil")
		return
	}
	if fieldD.Name() != "int32_field" {
		t.Errorf("field namer by number field name to descriptor: wrong result: %s!=int32_field", fieldD.Name())
		return
	}
}

func TestFieldNamerByName_FieldDescriptorToFieldName(t *testing.T) {
	fieldName := (&protobson.FieldNamerByName{}).FieldDescriptorToFieldName(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields().ByName("string_field"))
	if fieldName != "string_field" {
		t.Errorf("field namer by name descriptor to field name: wrong result: %s!=string_field", fieldName)
	}
}

func TestFieldNamerByName_FieldNameToFieldDescriptor(t *testing.T) {
	fieldD, err := (&protobson.FieldNamerByName{}).FieldNameToFieldDescriptor(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields(), "int32_field")
	if err != nil {
		t.Errorf("field namer by name field name to descriptor: err not nil: %v", err)
		return
	}
	if fieldD == nil {
		t.Errorf("field namer by name field name to descriptor: field descriptor is nil")
		return
	}
	if fieldD.Name() != "int32_field" {
		t.Errorf("field namer by name field name to descriptor: wrong result: %s!=int32_field", fieldD.Name())
		return
	}
}

func TestFieldNamerByJsonName_FieldDescriptorToFieldName(t *testing.T) {
	fieldName := (&protobson.FieldNamerByJsonName{}).FieldDescriptorToFieldName(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields().ByNumber(1))
	if fieldName != "stringField" {
		t.Errorf("field namer by json name descriptor to field name: wrong result: %s!=stringField", fieldName)
	}
}

func TestFieldNamerByJsonName_FieldNameToFieldDescriptor(t *testing.T) {
	fieldD, err := (&protobson.FieldNamerByJsonName{}).FieldNameToFieldDescriptor(
		(&pbtest.SimpleMessage{}).ProtoReflect().Descriptor().Fields(), "int32Field")
	if err != nil {
		t.Errorf("field namer by json name field name to descriptor: err not nil: %v", err)
		return
	}
	if fieldD == nil {
		t.Errorf("field namer by json name field name to descriptor: field descriptor is nil")
		return
	}
	if fieldD.Name() != "int32_field" {
		t.Errorf("field namer by json name field name to descriptor: wrong result: %s!=int32_field", fieldD.Name())
		return
	}
}
