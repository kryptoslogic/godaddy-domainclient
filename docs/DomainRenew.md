# DomainRenew

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | Pointer to **int32** | Number of years to extend the Domain. Must not exceed maximum for TLD. When omitted, defaults to &#x60;period&#x60; specified during original purchase | [optional] 

## Methods

### NewDomainRenew

`func NewDomainRenew() *DomainRenew`

NewDomainRenew instantiates a new DomainRenew object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRenewWithDefaults

`func NewDomainRenewWithDefaults() *DomainRenew`

NewDomainRenewWithDefaults instantiates a new DomainRenew object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *DomainRenew) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainRenew) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainRenew) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainRenew) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


