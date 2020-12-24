/*
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package godaddy

import (
	"encoding/json"
)

// JsonDataType struct for JsonDataType
type JsonDataType struct {
	Format *string `json:"format,omitempty"`
	Pattern *string `json:"pattern,omitempty"`
	Type string `json:"type"`
}

// NewJsonDataType instantiates a new JsonDataType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJsonDataType(type_ string, ) *JsonDataType {
	this := JsonDataType{}
	this.Type = type_
	return &this
}

// NewJsonDataTypeWithDefaults instantiates a new JsonDataType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJsonDataTypeWithDefaults() *JsonDataType {
	this := JsonDataType{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *JsonDataType) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDataType) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *JsonDataType) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *JsonDataType) SetFormat(v string) {
	o.Format = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *JsonDataType) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonDataType) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *JsonDataType) HasPattern() bool {
	if o != nil && o.Pattern != nil {
		return true
	}

	return false
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *JsonDataType) SetPattern(v string) {
	o.Pattern = &v
}

// GetType returns the Type field value
func (o *JsonDataType) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JsonDataType) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JsonDataType) SetType(v string) {
	o.Type = v
}

func (o JsonDataType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableJsonDataType struct {
	value *JsonDataType
	isSet bool
}

func (v NullableJsonDataType) Get() *JsonDataType {
	return v.value
}

func (v *NullableJsonDataType) Set(val *JsonDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableJsonDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableJsonDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJsonDataType(val *JsonDataType) *NullableJsonDataType {
	return &NullableJsonDataType{value: val, isSet: true}
}

func (v NullableJsonDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJsonDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


