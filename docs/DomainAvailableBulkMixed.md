# DomainAvailableBulkMixed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | [**[]DomainAvailableResponse**](DomainAvailableResponse.md) | Domain available response array | 
**Errors** | Pointer to [**[]DomainAvailableError**](DomainAvailableError.md) | Errors encountered while performing a domain available check | [optional] 

## Methods

### NewDomainAvailableBulkMixed

`func NewDomainAvailableBulkMixed(domains []DomainAvailableResponse, ) *DomainAvailableBulkMixed`

NewDomainAvailableBulkMixed instantiates a new DomainAvailableBulkMixed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainAvailableBulkMixedWithDefaults

`func NewDomainAvailableBulkMixedWithDefaults() *DomainAvailableBulkMixed`

NewDomainAvailableBulkMixedWithDefaults instantiates a new DomainAvailableBulkMixed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *DomainAvailableBulkMixed) GetDomains() []DomainAvailableResponse`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DomainAvailableBulkMixed) GetDomainsOk() (*[]DomainAvailableResponse, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DomainAvailableBulkMixed) SetDomains(v []DomainAvailableResponse)`

SetDomains sets Domains field to given value.


### GetErrors

`func (o *DomainAvailableBulkMixed) GetErrors() []DomainAvailableError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DomainAvailableBulkMixed) GetErrorsOk() (*[]DomainAvailableError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DomainAvailableBulkMixed) SetErrors(v []DomainAvailableError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DomainAvailableBulkMixed) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


