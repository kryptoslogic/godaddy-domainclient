# DNSRecordCreateTypeName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** |  | 
**Port** | Pointer to **int32** | Service port (SRV only) | [optional] 
**Priority** | Pointer to **int32** | Record priority (MX and SRV only) | [optional] 
**Protocol** | Pointer to **string** | Service protocol (SRV only) | [optional] 
**Service** | Pointer to **string** | Service type (SRV only) | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** | Record weight (SRV only) | [optional] 

## Methods

### NewDNSRecordCreateTypeName

`func NewDNSRecordCreateTypeName(data string, ) *DNSRecordCreateTypeName`

NewDNSRecordCreateTypeName instantiates a new DNSRecordCreateTypeName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDNSRecordCreateTypeNameWithDefaults

`func NewDNSRecordCreateTypeNameWithDefaults() *DNSRecordCreateTypeName`

NewDNSRecordCreateTypeNameWithDefaults instantiates a new DNSRecordCreateTypeName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DNSRecordCreateTypeName) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DNSRecordCreateTypeName) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DNSRecordCreateTypeName) SetData(v string)`

SetData sets Data field to given value.


### GetPort

`func (o *DNSRecordCreateTypeName) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DNSRecordCreateTypeName) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DNSRecordCreateTypeName) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DNSRecordCreateTypeName) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPriority

`func (o *DNSRecordCreateTypeName) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *DNSRecordCreateTypeName) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *DNSRecordCreateTypeName) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *DNSRecordCreateTypeName) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProtocol

`func (o *DNSRecordCreateTypeName) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DNSRecordCreateTypeName) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DNSRecordCreateTypeName) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DNSRecordCreateTypeName) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *DNSRecordCreateTypeName) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *DNSRecordCreateTypeName) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *DNSRecordCreateTypeName) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *DNSRecordCreateTypeName) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTtl

`func (o *DNSRecordCreateTypeName) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DNSRecordCreateTypeName) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DNSRecordCreateTypeName) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DNSRecordCreateTypeName) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetWeight

`func (o *DNSRecordCreateTypeName) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DNSRecordCreateTypeName) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DNSRecordCreateTypeName) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DNSRecordCreateTypeName) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


