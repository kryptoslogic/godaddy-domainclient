# JsonSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Models** | **map[string]interface{}** |  | 
**Properties** | **map[string]interface{}** |  | 
**Required** | **[]string** |  | 

## Methods

### NewJsonSchema

`func NewJsonSchema(id string, models map[string]interface{}, properties map[string]interface{}, required []string, ) *JsonSchema`

NewJsonSchema instantiates a new JsonSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonSchemaWithDefaults

`func NewJsonSchemaWithDefaults() *JsonSchema`

NewJsonSchemaWithDefaults instantiates a new JsonSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JsonSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JsonSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JsonSchema) SetId(v string)`

SetId sets Id field to given value.


### GetModels

`func (o *JsonSchema) GetModels() map[string]interface{}`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *JsonSchema) GetModelsOk() (*map[string]interface{}, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *JsonSchema) SetModels(v map[string]interface{})`

SetModels sets Models field to given value.


### GetProperties

`func (o *JsonSchema) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *JsonSchema) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *JsonSchema) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.


### GetRequired

`func (o *JsonSchema) GetRequired() []string`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *JsonSchema) GetRequiredOk() (*[]string, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *JsonSchema) SetRequired(v []string)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


