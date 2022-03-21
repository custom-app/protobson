package protobson

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strconv"
	"strings"
)

const (
	fieldPrefix = "pb_field_"
)

// FieldNamer is used to build field name in document from its field descriptor
type FieldNamer interface {
	// FieldDescriptorToFieldName is function to get document field name from proto field descriptor
	FieldDescriptorToFieldName(fd protoreflect.FieldDescriptor) string
	// FieldNameToFieldDescriptor is function to get proto field descriptor from document field name
	FieldNameToFieldDescriptor(fd protoreflect.FieldDescriptors, name string) (protoreflect.FieldDescriptor, error)
}

// WithFieldNamer - returns options with FieldNamer instance
func WithFieldNamer(namer FieldNamer) Option {
	return func(codec *protobufCodec) {
		codec.fieldNamer = namer
	}
}

// WithFieldNamerByNumber makes FieldNamer Option with FieldNamerByNumber instance
func WithFieldNamerByNumber() Option {
	return WithFieldNamer(&FieldNamerByNumber{})
}

// WithFieldNamerByName makes FieldNamer Option with FieldNamerByName instance
func WithFieldNamerByName() Option {
	return WithFieldNamer(&FieldNamerByName{})
}

// WithFieldNamerByJsonName makes FieldNamer Option with FieldNamerByJsonName instance
func WithFieldNamerByJsonName() Option {
	return WithFieldNamer(&FieldNamerByJsonName{})
}

// FieldNamerByNumber makes field names based on field tag number. With this implementation fields in proto spec can
// be renamed and documents still can be decoded without loss of any data
type FieldNamerByNumber struct {
}

func (f *FieldNamerByNumber) FieldDescriptorToFieldName(fd protoreflect.FieldDescriptor) string {
	return fmt.Sprintf("%v%v", fieldPrefix, fd.Number())
}

func (f *FieldNamerByNumber) FieldNameToFieldDescriptor(fd protoreflect.FieldDescriptors,
	name string) (protoreflect.FieldDescriptor, error) {
	if !strings.HasPrefix(name, fieldPrefix) {
		return nil, nil
	}
	numString := strings.TrimPrefix(name, fieldPrefix)
	num, err := strconv.Atoi(numString)
	if err != nil {
		return nil, err
	}
	return fd.ByNumber(protoreflect.FieldNumber(num)), nil
}

// FieldNamerByName makes field name based on field name in proto spec
type FieldNamerByName struct {
}

func (f *FieldNamerByName) FieldDescriptorToFieldName(fd protoreflect.FieldDescriptor) string {
	return string(fd.Name())
}

func (f *FieldNamerByName) FieldNameToFieldDescriptor(fd protoreflect.FieldDescriptors,
	name string) (protoreflect.FieldDescriptor, error) {
	return fd.ByName(protoreflect.Name(name)), nil
}

// FieldNamerByJsonName makes field name based on json name of field
type FieldNamerByJsonName struct {
}

func (f *FieldNamerByJsonName) FieldDescriptorToFieldName(fd protoreflect.FieldDescriptor) string {
	return fd.JSONName()
}

func (f *FieldNamerByJsonName) FieldNameToFieldDescriptor(fd protoreflect.FieldDescriptors,
	name string) (protoreflect.FieldDescriptor, error) {
	return fd.ByJSONName(name), nil
}
