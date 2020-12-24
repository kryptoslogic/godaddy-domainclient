# DomainAvailableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Whether or not the domain name is available | 
**Currency** | Pointer to **string** | Currency in which the &#x60;price&#x60; is listed. Only returned if tld is offered | [optional] [default to "USD"]
**Definitive** | **bool** | Whether or not the &#x60;available&#x60; answer has been definitively verified with the registry | 
**Domain** | **string** | Domain name | 
**Period** | Pointer to **int32** | Number of years included in the price. Only returned if tld is offered | [optional] 
**Price** | Pointer to **int32** | Price of the domain excluding taxes or fees. Only returned if tld is offered | [optional] 

## Methods

### NewDomainAvailableResponse

`func NewDomainAvailableResponse(available bool, definitive bool, domain string, ) *DomainAvailableResponse`

NewDomainAvailableResponse instantiates a new DomainAvailableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailableResponseWithDefaults

`func NewDomainAvailableResponseWithDefaults() *DomainAvailableResponse`

NewDomainAvailableResponseWithDefaults instantiates a new DomainAvailableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *DomainAvailableResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *DomainAvailableResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *DomainAvailableResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCurrency

`func (o *DomainAvailableResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DomainAvailableResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DomainAvailableResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DomainAvailableResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDefinitive

`func (o *DomainAvailableResponse) GetDefinitive() bool`

GetDefinitive returns the Definitive field if non-nil, zero value otherwise.

### GetDefinitiveOk

`func (o *DomainAvailableResponse) GetDefinitiveOk() (*bool, bool)`

GetDefinitiveOk returns a tuple with the Definitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitive

`func (o *DomainAvailableResponse) SetDefinitive(v bool)`

SetDefinitive sets Definitive field to given value.


### GetDomain

`func (o *DomainAvailableResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainAvailableResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainAvailableResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetPeriod

`func (o *DomainAvailableResponse) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainAvailableResponse) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainAvailableResponse) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainAvailableResponse) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrice

`func (o *DomainAvailableResponse) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *DomainAvailableResponse) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *DomainAvailableResponse) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *DomainAvailableResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


