# ErrorLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | 
**Fields** | Pointer to [**[]ErrorField**](ErrorField.md) | List of the specific fields, and the errors found with their contents | [optional] 
**Message** | Pointer to **string** | Human-readable, English description of the error | [optional] 
**RetryAfterSec** | **int32** | Number of seconds to wait before attempting a similar request | 

## Methods

### NewErrorLimit

`func NewErrorLimit(code string, retryAfterSec int32, ) *ErrorLimit`

NewErrorLimit instantiates a new ErrorLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorLimitWithDefaults

`func NewErrorLimitWithDefaults() *ErrorLimit`

NewErrorLimitWithDefaults instantiates a new ErrorLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorLimit) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorLimit) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorLimit) SetCode(v string)`

SetCode sets Code field to given value.


### GetFields

`func (o *ErrorLimit) GetFields() []ErrorField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ErrorLimit) GetFieldsOk() (*[]ErrorField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ErrorLimit) SetFields(v []ErrorField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ErrorLimit) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorLimit) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorLimit) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorLimit) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorLimit) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRetryAfterSec

`func (o *ErrorLimit) GetRetryAfterSec() int32`

GetRetryAfterSec returns the RetryAfterSec field if non-nil, zero value otherwise.

### GetRetryAfterSecOk

`func (o *ErrorLimit) GetRetryAfterSecOk() (*int32, bool)`

GetRetryAfterSecOk returns a tuple with the RetryAfterSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSec

`func (o *ErrorLimit) SetRetryAfterSec(v int32)`

SetRetryAfterSec sets RetryAfterSec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


