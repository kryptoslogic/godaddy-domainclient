/*
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package godaddy

import (
	"encoding/json"
)

// ErrorFieldDomainContactsValidate struct for ErrorFieldDomainContactsValidate
type ErrorFieldDomainContactsValidate struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// An array of domain names the error is for. If tlds are specified in the request, `domains` will contain tlds. For example, if `domains` in request is [\"test1.com\", \"test2.uk\", \"net\"], and the field is invalid for com and net, then one of the `fields` in response will have [\"test1.com\", \"net\"] as `domains`
	Domains []string `json:"domains"`
	// Human-readable, English description of the problem with the contents of the field
	Message *string `json:"message,omitempty"`
	// 1) JSONPath referring to the field within the data containing an error<br/>or<br/>2) JSONPath referring to an object containing an error
	Path string `json:"path"`
	// JSONPath referring to the field on the object referenced by `path` containing an error
	PathRelated *string `json:"pathRelated,omitempty"`
}

// NewErrorFieldDomainContactsValidate instantiates a new ErrorFieldDomainContactsValidate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorFieldDomainContactsValidate(code string, domains []string, path string, ) *ErrorFieldDomainContactsValidate {
	this := ErrorFieldDomainContactsValidate{}
	this.Code = code
	this.Domains = domains
	this.Path = path
	return &this
}

// NewErrorFieldDomainContactsValidateWithDefaults instantiates a new ErrorFieldDomainContactsValidate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorFieldDomainContactsValidateWithDefaults() *ErrorFieldDomainContactsValidate {
	this := ErrorFieldDomainContactsValidate{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorFieldDomainContactsValidate) GetCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorFieldDomainContactsValidate) GetCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorFieldDomainContactsValidate) SetCode(v string) {
	o.Code = v
}

// GetDomains returns the Domains field value
func (o *ErrorFieldDomainContactsValidate) GetDomains() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value
// and a boolean to check if the value has been set.
func (o *ErrorFieldDomainContactsValidate) GetDomainsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Domains, true
}

// SetDomains sets field value
func (o *ErrorFieldDomainContactsValidate) SetDomains(v []string) {
	o.Domains = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorFieldDomainContactsValidate) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorFieldDomainContactsValidate) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorFieldDomainContactsValidate) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorFieldDomainContactsValidate) SetMessage(v string) {
	o.Message = &v
}

// GetPath returns the Path field value
func (o *ErrorFieldDomainContactsValidate) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ErrorFieldDomainContactsValidate) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ErrorFieldDomainContactsValidate) SetPath(v string) {
	o.Path = v
}

// GetPathRelated returns the PathRelated field value if set, zero value otherwise.
func (o *ErrorFieldDomainContactsValidate) GetPathRelated() string {
	if o == nil || o.PathRelated == nil {
		var ret string
		return ret
	}
	return *o.PathRelated
}

// GetPathRelatedOk returns a tuple with the PathRelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorFieldDomainContactsValidate) GetPathRelatedOk() (*string, bool) {
	if o == nil || o.PathRelated == nil {
		return nil, false
	}
	return o.PathRelated, true
}

// HasPathRelated returns a boolean if a field has been set.
func (o *ErrorFieldDomainContactsValidate) HasPathRelated() bool {
	if o != nil && o.PathRelated != nil {
		return true
	}

	return false
}

// SetPathRelated gets a reference to the given string and assigns it to the PathRelated field.
func (o *ErrorFieldDomainContactsValidate) SetPathRelated(v string) {
	o.PathRelated = &v
}

func (o ErrorFieldDomainContactsValidate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["domains"] = o.Domains
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.PathRelated != nil {
		toSerialize["pathRelated"] = o.PathRelated
	}
	return json.Marshal(toSerialize)
}

type NullableErrorFieldDomainContactsValidate struct {
	value *ErrorFieldDomainContactsValidate
	isSet bool
}

func (v NullableErrorFieldDomainContactsValidate) Get() *ErrorFieldDomainContactsValidate {
	return v.value
}

func (v *NullableErrorFieldDomainContactsValidate) Set(val *ErrorFieldDomainContactsValidate) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorFieldDomainContactsValidate) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorFieldDomainContactsValidate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorFieldDomainContactsValidate(val *ErrorFieldDomainContactsValidate) *NullableErrorFieldDomainContactsValidate {
	return &NullableErrorFieldDomainContactsValidate{value: val, isSet: true}
}

func (v NullableErrorFieldDomainContactsValidate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorFieldDomainContactsValidate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


