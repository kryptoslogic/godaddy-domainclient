# DNSRecordCreateType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** |  | 
**Name** | **string** |  | 
**Port** | Pointer to **int32** | Service port (SRV only) | [optional] 
**Priority** | Pointer to **int32** | Record priority (MX and SRV only) | [optional] 
**Protocol** | Pointer to **string** | Service protocol (SRV only) | [optional] 
**Service** | Pointer to **string** | Service type (SRV only) | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** | Record weight (SRV only) | [optional] 

## Methods

### NewDNSRecordCreateType

`func NewDNSRecordCreateType(data string, name string, ) *DNSRecordCreateType`

NewDNSRecordCreateType instantiates a new DNSRecordCreateType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordCreateTypeWithDefaults

`func NewDNSRecordCreateTypeWithDefaults() *DNSRecordCreateType`

NewDNSRecordCreateTypeWithDefaults instantiates a new DNSRecordCreateType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DNSRecordCreateType) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DNSRecordCreateType) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DNSRecordCreateType) SetData(v string)`

SetData sets Data field to given value.


### GetName

`func (o *DNSRecordCreateType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DNSRecordCreateType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DNSRecordCreateType) SetName(v string)`

SetName sets Name field to given value.


### GetPort

`func (o *DNSRecordCreateType) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecordCreateType) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecordCreateType) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSRecordCreateType) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *DNSRecordCreateType) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordCreateType) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordCreateType) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordCreateType) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *DNSRecordCreateType) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DNSRecordCreateType) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DNSRecordCreateType) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DNSRecordCreateType) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *DNSRecordCreateType) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DNSRecordCreateType) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DNSRecordCreateType) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DNSRecordCreateType) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTtl

`func (o *DNSRecordCreateType) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordCreateType) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordCreateType) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecordCreateType) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetWeight

`func (o *DNSRecordCreateType) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecordCreateType) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecordCreateType) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DNSRecordCreateType) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


