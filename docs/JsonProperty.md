# JsonProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**Items** | Pointer to **map[string]interface{}** |  | [optional] 
**MaxItems** | Pointer to **int32** |  | [optional] 
**Maximum** | Pointer to **int32** |  | [optional] 
**MinItems** | Pointer to **int32** |  | [optional] 
**Minimum** | Pointer to **int32** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Required** | **bool** |  | 
**Type** | **string** |  | 

## Methods

### NewJsonProperty

`func NewJsonProperty(required bool, type_ string, ) *JsonProperty`

NewJsonProperty instantiates a new JsonProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonPropertyWithDefaults

`func NewJsonPropertyWithDefaults() *JsonProperty`

NewJsonPropertyWithDefaults instantiates a new JsonProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *JsonProperty) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *JsonProperty) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *JsonProperty) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *JsonProperty) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetFormat

`func (o *JsonProperty) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *JsonProperty) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *JsonProperty) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *JsonProperty) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetItems

`func (o *JsonProperty) GetItems() map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *JsonProperty) GetItemsOk() (*map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *JsonProperty) SetItems(v map[string]interface{})`

SetItems sets Items field to given value.

### HasItems

`func (o *JsonProperty) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaxItems

`func (o *JsonProperty) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *JsonProperty) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *JsonProperty) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *JsonProperty) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### GetMaximum

`func (o *JsonProperty) GetMaximum() int32`

GetMaximum returns the Maximum field if non-nil, zero value otherwise.

### GetMaximumOk

`func (o *JsonProperty) GetMaximumOk() (*int32, bool)`

GetMaximumOk returns a tuple with the Maximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximum

`func (o *JsonProperty) SetMaximum(v int32)`

SetMaximum sets Maximum field to given value.

### HasMaximum

`func (o *JsonProperty) HasMaximum() bool`

HasMaximum returns a boolean if a field has been set.

### GetMinItems

`func (o *JsonProperty) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *JsonProperty) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *JsonProperty) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *JsonProperty) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### GetMinimum

`func (o *JsonProperty) GetMinimum() int32`

GetMinimum returns the Minimum field if non-nil, zero value otherwise.

### GetMinimumOk

`func (o *JsonProperty) GetMinimumOk() (*int32, bool)`

GetMinimumOk returns a tuple with the Minimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimum

`func (o *JsonProperty) SetMinimum(v int32)`

SetMinimum sets Minimum field to given value.

### HasMinimum

`func (o *JsonProperty) HasMinimum() bool`

HasMinimum returns a boolean if a field has been set.

### GetPattern

`func (o *JsonProperty) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *JsonProperty) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *JsonProperty) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *JsonProperty) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRequired

`func (o *JsonProperty) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *JsonProperty) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *JsonProperty) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetType

`func (o *JsonProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonProperty) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


