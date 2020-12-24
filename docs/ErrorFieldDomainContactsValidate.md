# ErrorFieldDomainContactsValidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | 
**Domains** | **[]string** | An array of domain names the error is for. If tlds are specified in the request, &#x60;domains&#x60; will contain tlds. For example, if &#x60;domains&#x60; in request is [\&quot;test1.com\&quot;, \&quot;test2.uk\&quot;, \&quot;net\&quot;], and the field is invalid for com and net, then one of the &#x60;fields&#x60; in response will have [\&quot;test1.com\&quot;, \&quot;net\&quot;] as &#x60;domains&#x60; | 
**Message** | Pointer to **string** | Human-readable, English description of the problem with the contents of the field | [optional] 
**Path** | **string** | 1) JSONPath referring to the field within the data containing an error&lt;br/&gt;or&lt;br/&gt;2) JSONPath referring to an object containing an error | 
**PathRelated** | Pointer to **string** | JSONPath referring to the field on the object referenced by &#x60;path&#x60; containing an error | [optional] 

## Methods

### NewErrorFieldDomainContactsValidate

`func NewErrorFieldDomainContactsValidate(code string, domains []string, path string, ) *ErrorFieldDomainContactsValidate`

NewErrorFieldDomainContactsValidate instantiates a new ErrorFieldDomainContactsValidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorFieldDomainContactsValidateWithDefaults

`func NewErrorFieldDomainContactsValidateWithDefaults() *ErrorFieldDomainContactsValidate`

NewErrorFieldDomainContactsValidateWithDefaults instantiates a new ErrorFieldDomainContactsValidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorFieldDomainContactsValidate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorFieldDomainContactsValidate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorFieldDomainContactsValidate) SetCode(v string)`

SetCode sets Code field to given value.


### GetDomains

`func (o *ErrorFieldDomainContactsValidate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *ErrorFieldDomainContactsValidate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *ErrorFieldDomainContactsValidate) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetMessage

`func (o *ErrorFieldDomainContactsValidate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorFieldDomainContactsValidate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorFieldDomainContactsValidate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorFieldDomainContactsValidate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPath

`func (o *ErrorFieldDomainContactsValidate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ErrorFieldDomainContactsValidate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ErrorFieldDomainContactsValidate) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathRelated

`func (o *ErrorFieldDomainContactsValidate) GetPathRelated() string`

GetPathRelated returns the PathRelated field if non-nil, zero value otherwise.

### GetPathRelatedOk

`func (o *ErrorFieldDomainContactsValidate) GetPathRelatedOk() (*string, bool)`

GetPathRelatedOk returns a tuple with the PathRelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRelated

`func (o *ErrorFieldDomainContactsValidate) SetPathRelated(v string)`

SetPathRelated sets PathRelated field to given value.

### HasPathRelated

`func (o *ErrorFieldDomainContactsValidate) HasPathRelated() bool`

HasPathRelated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


