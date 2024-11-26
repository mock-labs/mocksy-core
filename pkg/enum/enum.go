// Package enumvalidator provides TypeScript-like enum validation functionality for Go
package enumvalidator

import (
	"fmt"
	"reflect"
)

// EnumValidator provides validation functionality for enum-like values
type EnumValidator struct {
	values     map[interface{}]bool
	reverseMap map[interface{}]interface{} // For numeric types only
	enumType   reflect.Type
}

// NewEnumValidator creates a new validator from a map of values
func NewEnumValidator[K comparable, V any](enumMap map[K]V) *EnumValidator {
	values := make(map[interface{}]bool)
	reverseMap := make(map[interface{}]interface{})
	
	// Store all values for validation
	for k, v := range enumMap {
		values[k] = true
		
		// Create reverse mapping only for numeric types
		if reflect.TypeOf(k).Kind() >= reflect.Int && reflect.TypeOf(k).Kind() <= reflect.Float64 {
			reverseMap[v] = k
		}
	}
	
	return &EnumValidator{
		values:     values,
		reverseMap: reverseMap,
		enumType:   reflect.TypeOf((*K)(nil)).Elem(),
	}
}

// IsValid checks if a value is a valid enum value
func (e *EnumValidator) IsValid(value interface{}) bool {
	if !reflect.TypeOf(value).AssignableTo(e.enumType) {
		return false
	}
	return e.values[value]
}

// GetKey returns the key for a given value (only works for numeric enums)
func (e *EnumValidator) GetKey(value interface{}) (interface{}, error) {
	if key, exists := e.reverseMap[value]; exists {
		return key, nil
	}
	return nil, fmt.Errorf("value %v not found in enum", value)
}

// GetValues returns all valid enum values
func (e *EnumValidator) GetValues() []interface{} {
	values := make([]interface{}, 0, len(e.values))
	for v := range e.values {
		values = append(values, v)
	}
	return values
}

// MustBeValid panics if the value is not a valid enum value
func (e *EnumValidator) MustBeValid(value interface{}) {
	if !e.IsValid(value) {
		panic(fmt.Sprintf("invalid enum value: %v", value))
	}
}