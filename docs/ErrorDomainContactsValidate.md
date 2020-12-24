# ErrorDomainContactsValidate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | 
**Fields** | Pointer to [**[]ErrorFieldDomainContactsValidate**](ErrorFieldDomainContactsValidate.md) | List of the specific fields, and the errors found with their contents | [optional] 
**Message** | Pointer to **string** | Human-readable, English description of the error | [optional] 
**Stack** | Pointer to **[]string** | Stack trace indicating where the error occurred.&lt;br/&gt;NOTE: This attribute &lt;strong&gt;MAY&lt;/strong&gt; be included for Development and Test environments. However, it &lt;strong&gt;MUST NOT&lt;/strong&gt; be exposed from OTE nor Production systems | [optional] 

## Methods

### NewErrorDomainContactsValidate

`func NewErrorDomainContactsValidate(code string, ) *ErrorDomainContactsValidate`

NewErrorDomainContactsValidate instantiates a new ErrorDomainContactsValidate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDomainContactsValidateWithDefaults

`func NewErrorDomainContactsValidateWithDefaults() *ErrorDomainContactsValidate`

NewErrorDomainContactsValidateWithDefaults instantiates a new ErrorDomainContactsValidate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDomainContactsValidate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDomainContactsValidate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDomainContactsValidate) SetCode(v string)`

SetCode sets Code field to given value.


### GetFields

`func (o *ErrorDomainContactsValidate) GetFields() []ErrorFieldDomainContactsValidate`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorDomainContactsValidate) GetFieldsOk() (*[]ErrorFieldDomainContactsValidate, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorDomainContactsValidate) SetFields(v []ErrorFieldDomainContactsValidate)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorDomainContactsValidate) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorDomainContactsValidate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDomainContactsValidate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDomainContactsValidate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorDomainContactsValidate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStack

`func (o *ErrorDomainContactsValidate) GetStack() []string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ErrorDomainContactsValidate) GetStackOk() (*[]string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ErrorDomainContactsValidate) SetStack(v []string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ErrorDomainContactsValidate) HasStack() bool`

HasStack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


