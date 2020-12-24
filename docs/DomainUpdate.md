# DomainUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locked** | Pointer to **bool** | Whether or not the domain should be locked to prevent transfers | [optional] 
**NameServers** | Pointer to **[]map[string]interface{}** | Fully-qualified domain names for Name Servers to associate with the domain | [optional] 
**RenewAuto** | Pointer to **bool** | Whether or not the domain should be configured to automatically renew | [optional] 
**SubaccountId** | Pointer to **string** | Reseller subaccount shopperid who can manage the domain | [optional] 
**ExposeWhois** | Pointer to **bool** | Whether or not the domain contact details should be shown in the WHOIS | [optional] 
**Consent** | Pointer to [**ConsentDomainUpdate**](ConsentDomainUpdate.md) |  | [optional] 

## Methods

### NewDomainUpdate

`func NewDomainUpdate() *DomainUpdate`

NewDomainUpdate instantiates a new DomainUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUpdateWithDefaults

`func NewDomainUpdateWithDefaults() *DomainUpdate`

NewDomainUpdateWithDefaults instantiates a new DomainUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocked

`func (o *DomainUpdate) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DomainUpdate) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DomainUpdate) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *DomainUpdate) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetNameServers

`func (o *DomainUpdate) GetNameServers() []map[string]interface{}`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DomainUpdate) GetNameServersOk() (*[]map[string]interface{}, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DomainUpdate) SetNameServers(v []map[string]interface{})`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *DomainUpdate) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetRenewAuto

`func (o *DomainUpdate) GetRenewAuto() bool`

GetRenewAuto returns the RenewAuto field if non-nil, zero value otherwise.

### GetRenewAutoOk

`func (o *DomainUpdate) GetRenewAutoOk() (*bool, bool)`

GetRenewAutoOk returns a tuple with the RenewAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuto

`func (o *DomainUpdate) SetRenewAuto(v bool)`

SetRenewAuto sets RenewAuto field to given value.

### HasRenewAuto

`func (o *DomainUpdate) HasRenewAuto() bool`

HasRenewAuto returns a boolean if a field has been set.

### GetSubaccountId

`func (o *DomainUpdate) GetSubaccountId() string`

GetSubaccountId returns the SubaccountId field if non-nil, zero value otherwise.

### GetSubaccountIdOk

`func (o *DomainUpdate) GetSubaccountIdOk() (*string, bool)`

GetSubaccountIdOk returns a tuple with the SubaccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccountId

`func (o *DomainUpdate) SetSubaccountId(v string)`

SetSubaccountId sets SubaccountId field to given value.

### HasSubaccountId

`func (o *DomainUpdate) HasSubaccountId() bool`

HasSubaccountId returns a boolean if a field has been set.

### GetExposeWhois

`func (o *DomainUpdate) GetExposeWhois() bool`

GetExposeWhois returns the ExposeWhois field if non-nil, zero value otherwise.

### GetExposeWhoisOk

`func (o *DomainUpdate) GetExposeWhoisOk() (*bool, bool)`

GetExposeWhoisOk returns a tuple with the ExposeWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeWhois

`func (o *DomainUpdate) SetExposeWhois(v bool)`

SetExposeWhois sets ExposeWhois field to given value.

### HasExposeWhois

`func (o *DomainUpdate) HasExposeWhois() bool`

HasExposeWhois returns a boolean if a field has been set.

### GetConsent

`func (o *DomainUpdate) GetConsent() ConsentDomainUpdate`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *DomainUpdate) GetConsentOk() (*ConsentDomainUpdate, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *DomainUpdate) SetConsent(v ConsentDomainUpdate)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *DomainUpdate) HasConsent() bool`

HasConsent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


