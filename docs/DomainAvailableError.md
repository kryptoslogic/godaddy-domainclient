# DomainAvailableError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | 
**Domain** | **string** | Domain name | 
**Message** | Pointer to **string** | Human-readable, English description of the error | [optional] 
**Path** | **string** | &lt;ul&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field containing an error&lt;/li&gt; &lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;OR&lt;/strong&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field that refers to an object containing an error, with more detail in &#x60;pathRelated&#x60;&lt;/li&gt; &lt;/ul&gt; | 
**Status** | **int32** | HTTP status code that would return for a single check | 

## Methods

### NewDomainAvailableError

`func NewDomainAvailableError(code string, domain string, path string, status int32, ) *DomainAvailableError`

NewDomainAvailableError instantiates a new DomainAvailableError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailableErrorWithDefaults

`func NewDomainAvailableErrorWithDefaults() *DomainAvailableError`

NewDomainAvailableErrorWithDefaults instantiates a new DomainAvailableError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *DomainAvailableError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *DomainAvailableError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *DomainAvailableError) SetCode(v string)`

SetCode sets Code field to given value.


### GetDomain

`func (o *DomainAvailableError) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainAvailableError) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainAvailableError) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetMessage

`func (o *DomainAvailableError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DomainAvailableError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DomainAvailableError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DomainAvailableError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPath

`func (o *DomainAvailableError) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DomainAvailableError) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DomainAvailableError) SetPath(v string)`

SetPath sets Path field to given value.


### GetStatus

`func (o *DomainAvailableError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainAvailableError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainAvailableError) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


