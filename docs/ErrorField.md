# ErrorField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | 
**Message** | Pointer to **string** | Human-readable, English description of the problem with the contents of the field | [optional] 
**Path** | **string** | &lt;ul&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field containing an error&lt;/li&gt; &lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;OR&lt;/strong&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field that refers to an object containing an error, with more detail in &#x60;pathRelated&#x60;&lt;/li&gt; &lt;/ul&gt; | 
**PathRelated** | Pointer to **string** | JSONPath referring to a field containing an error, which is referenced by &#x60;path&#x60; | [optional] 

## Methods

### NewErrorField

`func NewErrorField(code string, path string, ) *ErrorField`

NewErrorField instantiates a new ErrorField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorFieldWithDefaults

`func NewErrorFieldWithDefaults() *ErrorField`

NewErrorFieldWithDefaults instantiates a new ErrorField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorField) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorField) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorField) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorField) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorField) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorField) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorField) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPath

`func (o *ErrorField) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ErrorField) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ErrorField) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathRelated

`func (o *ErrorField) GetPathRelated() string`

GetPathRelated returns the PathRelated field if non-nil, zero value otherwise.

### GetPathRelatedOk

`func (o *ErrorField) GetPathRelatedOk() (*string, bool)`

GetPathRelatedOk returns a tuple with the PathRelated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathRelated

`func (o *ErrorField) SetPathRelated(v string)`

SetPathRelated sets PathRelated field to given value.

### HasPathRelated

`func (o *ErrorField) HasPathRelated() bool`

HasPathRelated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


