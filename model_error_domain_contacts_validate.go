/*
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package godaddy

import (
	"encoding/json"
)

// ErrorDomainContactsValidate struct for ErrorDomainContactsValidate
type ErrorDomainContactsValidate struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// List of the specific fields, and the errors found with their contents
	Fields *[]ErrorFieldDomainContactsValidate `json:"fields,omitempty"`
	// Human-readable, English description of the error
	Message *string `json:"message,omitempty"`
	// Stack trace indicating where the error occurred.<br/>NOTE: This attribute <strong>MAY</strong> be included for Development and Test environments. However, it <strong>MUST NOT</strong> be exposed from OTE nor Production systems
	Stack *[]string `json:"stack,omitempty"`
}

// NewErrorDomainContactsValidate instantiates a new ErrorDomainContactsValidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDomainContactsValidate(code string, ) *ErrorDomainContactsValidate {
	this := ErrorDomainContactsValidate{}
	this.Code = code
	return &this
}

// NewErrorDomainContactsValidateWithDefaults instantiates a new ErrorDomainContactsValidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDomainContactsValidateWithDefaults() *ErrorDomainContactsValidate {
	this := ErrorDomainContactsValidate{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorDomainContactsValidate) GetCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorDomainContactsValidate) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorDomainContactsValidate) SetCode(v string) {
	o.Code = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *ErrorDomainContactsValidate) GetFields() []ErrorFieldDomainContactsValidate {
	if o == nil || o.Fields == nil {
		var ret []ErrorFieldDomainContactsValidate
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDomainContactsValidate) GetFieldsOk() (*[]ErrorFieldDomainContactsValidate, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *ErrorDomainContactsValidate) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []ErrorFieldDomainContactsValidate and assigns it to the Fields field.
func (o *ErrorDomainContactsValidate) SetFields(v []ErrorFieldDomainContactsValidate) {
	o.Fields = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorDomainContactsValidate) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDomainContactsValidate) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorDomainContactsValidate) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorDomainContactsValidate) SetMessage(v string) {
	o.Message = &v
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *ErrorDomainContactsValidate) GetStack() []string {
	if o == nil || o.Stack == nil {
		var ret []string
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorDomainContactsValidate) GetStackOk() (*[]string, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *ErrorDomainContactsValidate) HasStack() bool {
	if o != nil && o.Stack != nil {
		return true
	}

	return false
}

// SetStack gets a reference to the given []string and assigns it to the Stack field.
func (o *ErrorDomainContactsValidate) SetStack(v []string) {
	o.Stack = &v
}

func (o ErrorDomainContactsValidate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	return json.Marshal(toSerialize)
}

type NullableErrorDomainContactsValidate struct {
	value *ErrorDomainContactsValidate
	isSet bool
}

func (v NullableErrorDomainContactsValidate) Get() *ErrorDomainContactsValidate {
	return v.value
}

func (v *NullableErrorDomainContactsValidate) Set(val *ErrorDomainContactsValidate) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDomainContactsValidate) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDomainContactsValidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDomainContactsValidate(val *ErrorDomainContactsValidate) *NullableErrorDomainContactsValidate {
	return &NullableErrorDomainContactsValidate{value: val, isSet: true}
}

func (v NullableErrorDomainContactsValidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDomainContactsValidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


