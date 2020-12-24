/*
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package godaddy

import (
	"encoding/json"
)

// TldSummary struct for TldSummary
type TldSummary struct {
	// Name of the top-level domain
	Name string `json:"name"`
	// Type of the top-level domain
	Type string `json:"type"`
}

// NewTldSummary instantiates a new TldSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTldSummary(name string, type_ string, ) *TldSummary {
	this := TldSummary{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewTldSummaryWithDefaults instantiates a new TldSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTldSummaryWithDefaults() *TldSummary {
	this := TldSummary{}
	var type_ string = "GENERIC"
	this.Type = type_
	return &this
}

// GetName returns the Name field value
func (o *TldSummary) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TldSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TldSummary) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *TldSummary) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TldSummary) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TldSummary) SetType(v string) {
	o.Type = v
}

func (o TldSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTldSummary struct {
	value *TldSummary
	isSet bool
}

func (v NullableTldSummary) Get() *TldSummary {
	return v.value
}

func (v *NullableTldSummary) Set(val *TldSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTldSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTldSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTldSummary(val *TldSummary) *NullableTldSummary {
	return &NullableTldSummary{value: val, isSet: true}
}

func (v NullableTldSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTldSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


