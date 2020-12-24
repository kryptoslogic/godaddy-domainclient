# JsonDataType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewJsonDataType

`func NewJsonDataType(type_ string, ) *JsonDataType`

NewJsonDataType instantiates a new JsonDataType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonDataTypeWithDefaults

`func NewJsonDataTypeWithDefaults() *JsonDataType`

NewJsonDataTypeWithDefaults instantiates a new JsonDataType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *JsonDataType) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *JsonDataType) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *JsonDataType) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *JsonDataType) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetPattern

`func (o *JsonDataType) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *JsonDataType) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *JsonDataType) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *JsonDataType) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetType

`func (o *JsonDataType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonDataType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonDataType) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


