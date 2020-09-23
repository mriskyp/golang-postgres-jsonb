package jsonb

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type JSONB []byte

const (
	defaultInt        = 0
	defaultNullString = "null"
)

/**
first you need to add type JSONB in your struct
then get value of query row from postgres struct scan
then get value of jsonb from this builder
*/

//IsNull  check if jsonb is null
func (j JSONB) IsNull() bool {
	checker := (len(j) == defaultInt || string(j) == defaultNullString)
	return checker
}

//IsValid  check if jsonb is valid format
func (j JSONB) IsValid(value interface{}) bool {
	_, ok := value.(JSONB)
	if !ok {
		return false
	}
	return true
}

//GetValue  is a func to get value of jsonb
func (j JSONB) GetValue() (driver.Value, error) {
	if j.IsNull() {
		return nil, nil
	}
	return string(j), nil
}

//Scan  is a func to scan jsonb from interface
func (j *JSONB) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	isValid := j.IsValid(value)
	if !isValid {
		return errors.New("data type is not jsonb")
	}
	return nil
}

// MarshalJSONB is a func returns JSONB encoding of m.
func (j JSONB) MarshalJSONB() ([]byte, error) {
	if j == nil {
		return []byte(defaultNullString), nil
	}
	return j, nil
}

// UnmarshalJSONB is a func return jsonb unmarshal as string
func (j *JSONB) UnmarshalJSONB(data []byte) (string, error) {
	var jsonbStringify string
	if j == nil {
		return jsonbStringify, errors.New("error Unmarshal JSON Raw on nil pointer")
	}
	err := json.Unmarshal(data, &jsonbStringify)
	if err != nil {
		errUnmarshal := fmt.Sprintf("error Unmarshal JSON Raw %s", err.Error())
		return jsonbStringify, errors.New(errUnmarshal)
	}
	return jsonbStringify, nil
}
