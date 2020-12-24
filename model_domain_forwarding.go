/*
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package godaddy

import (
	"encoding/json"
)

// DomainForwarding struct for DomainForwarding
type DomainForwarding struct {
	// The fqdn (domain or sub domain) to forward (ex somedomain.com or sub.somedomain.com)
	Fqdn string `json:"fqdn"`
	// The type of fowarding to implement<br/><ul><li><strong style='margin-left: 12px;'>MASKED</strong> - Prevents the forwarded domain or subdomain URL from displaying in the browser's address bar.</li><li><strong style='margin-left: 12px;'>REDIRECT_PERMANENT*</strong> - Redirects to the url you specified in the forwardTo field using a `301 Moved Permanently` HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.</li><li><strong style='margin-left: 12px;'>REDIRECT_TEMPORARY</strong> - Redirects to the url you specified in the forwardTo field using a `302 Found` HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.</li></ul>
	Type string `json:"type"`
	// Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/)
	Url string `json:"url"`
	Mask *DomainForwardingMask `json:"mask,omitempty"`
}

// NewDomainForwarding instantiates a new DomainForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainForwarding(fqdn string, type_ string, url string, ) *DomainForwarding {
	this := DomainForwarding{}
	this.Fqdn = fqdn
	this.Type = type_
	this.Url = url
	return &this
}

// NewDomainForwardingWithDefaults instantiates a new DomainForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainForwardingWithDefaults() *DomainForwarding {
	this := DomainForwarding{}
	var type_ string = "REDIRECT_PERMANENT"
	this.Type = type_
	return &this
}

// GetFqdn returns the Fqdn field value
func (o *DomainForwarding) GetFqdn() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *DomainForwarding) GetFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *DomainForwarding) SetFqdn(v string) {
	o.Fqdn = v
}

// GetType returns the Type field value
func (o *DomainForwarding) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DomainForwarding) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DomainForwarding) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value
func (o *DomainForwarding) GetUrl() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DomainForwarding) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DomainForwarding) SetUrl(v string) {
	o.Url = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *DomainForwarding) GetMask() DomainForwardingMask {
	if o == nil || o.Mask == nil {
		var ret DomainForwardingMask
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainForwarding) GetMaskOk() (*DomainForwardingMask, bool) {
	if o == nil || o.Mask == nil {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *DomainForwarding) HasMask() bool {
	if o != nil && o.Mask != nil {
		return true
	}

	return false
}

// SetMask gets a reference to the given DomainForwardingMask and assigns it to the Mask field.
func (o *DomainForwarding) SetMask(v DomainForwardingMask) {
	o.Mask = &v
}

func (o DomainForwarding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fqdn"] = o.Fqdn
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if o.Mask != nil {
		toSerialize["mask"] = o.Mask
	}
	return json.Marshal(toSerialize)
}

type NullableDomainForwarding struct {
	value *DomainForwarding
	isSet bool
}

func (v NullableDomainForwarding) Get() *DomainForwarding {
	return v.value
}

func (v *NullableDomainForwarding) Set(val *DomainForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainForwarding(val *DomainForwarding) *NullableDomainForwarding {
	return &NullableDomainForwarding{value: val, isSet: true}
}

func (v NullableDomainForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


