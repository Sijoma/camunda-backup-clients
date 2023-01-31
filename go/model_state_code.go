/*
Backup Management API

Management endpoint to query, take, and delete backups of Zeebe.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package camunda_backup_clients

import (
	"encoding/json"
	"fmt"
)

// StateCode The state of the backup.
type StateCode string

// List of StateCode
const (
	DOES_NOT_EXIST StateCode = "DOES_NOT_EXIST"
	IN_PROGRESS StateCode = "IN_PROGRESS"
	COMPLETED StateCode = "COMPLETED"
	FAILED StateCode = "FAILED"
	INCOMPLETE StateCode = "INCOMPLETE"
	INCOMPATIBLE StateCode = "INCOMPATIBLE"
)

// All allowed values of StateCode enum
var AllowedStateCodeEnumValues = []StateCode{
	"DOES_NOT_EXIST",
	"IN_PROGRESS",
	"COMPLETED",
	"FAILED",
	"INCOMPLETE",
	"INCOMPATIBLE",
}

func (v *StateCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StateCode(value)
	for _, existing := range AllowedStateCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StateCode", value)
}

// NewStateCodeFromValue returns a pointer to a valid StateCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStateCodeFromValue(v string) (*StateCode, error) {
	ev := StateCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StateCode: valid values are %v", v, AllowedStateCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StateCode) IsValid() bool {
	for _, existing := range AllowedStateCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StateCode value
func (v StateCode) Ptr() *StateCode {
	return &v
}

type NullableStateCode struct {
	value *StateCode
	isSet bool
}

func (v NullableStateCode) Get() *StateCode {
	return v.value
}

func (v *NullableStateCode) Set(val *StateCode) {
	v.value = val
	v.isSet = true
}

func (v NullableStateCode) IsSet() bool {
	return v.isSet
}

func (v *NullableStateCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateCode(val *StateCode) *NullableStateCode {
	return &NullableStateCode{value: val, isSet: true}
}

func (v NullableStateCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
