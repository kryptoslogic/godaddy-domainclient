# DomainPurchase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consent** | [**Consent**](Consent.md) |  | 
**ContactAdmin** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactBilling** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactRegistrant** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactTech** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Domain** | **string** | For internationalized domain names with non-ascii characters, the domain name is converted to punycode before format and pattern validation rules are checked | 
**NameServers** | Pointer to **[]string** |  | [optional] 
**Period** | Pointer to **int32** |  | [optional] 
**Privacy** | Pointer to **bool** |  | [optional] [default to false]
**RenewAuto** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewDomainPurchase

`func NewDomainPurchase(consent Consent, domain string, ) *DomainPurchase`

NewDomainPurchase instantiates a new DomainPurchase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPurchaseWithDefaults

`func NewDomainPurchaseWithDefaults() *DomainPurchase`

NewDomainPurchaseWithDefaults instantiates a new DomainPurchase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsent

`func (o *DomainPurchase) GetConsent() Consent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *DomainPurchase) GetConsentOk() (*Consent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *DomainPurchase) SetConsent(v Consent)`

SetConsent sets Consent field to given value.


### GetContactAdmin

`func (o *DomainPurchase) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainPurchase) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainPurchase) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.

### HasContactAdmin

`func (o *DomainPurchase) HasContactAdmin() bool`

HasContactAdmin returns a boolean if a field has been set.

### GetContactBilling

`func (o *DomainPurchase) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainPurchase) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainPurchase) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.

### HasContactBilling

`func (o *DomainPurchase) HasContactBilling() bool`

HasContactBilling returns a boolean if a field has been set.

### GetContactRegistrant

`func (o *DomainPurchase) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainPurchase) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainPurchase) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.

### HasContactRegistrant

`func (o *DomainPurchase) HasContactRegistrant() bool`

HasContactRegistrant returns a boolean if a field has been set.

### GetContactTech

`func (o *DomainPurchase) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainPurchase) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainPurchase) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.

### HasContactTech

`func (o *DomainPurchase) HasContactTech() bool`

HasContactTech returns a boolean if a field has been set.

### GetDomain

`func (o *DomainPurchase) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainPurchase) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainPurchase) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetNameServers

`func (o *DomainPurchase) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DomainPurchase) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DomainPurchase) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *DomainPurchase) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetPeriod

`func (o *DomainPurchase) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainPurchase) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainPurchase) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainPurchase) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrivacy

`func (o *DomainPurchase) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *DomainPurchase) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *DomainPurchase) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *DomainPurchase) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRenewAuto

`func (o *DomainPurchase) GetRenewAuto() bool`

GetRenewAuto returns the RenewAuto field if non-nil, zero value otherwise.

### GetRenewAutoOk

`func (o *DomainPurchase) GetRenewAutoOk() (*bool, bool)`

GetRenewAutoOk returns a tuple with the RenewAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuto

`func (o *DomainPurchase) SetRenewAuto(v bool)`

SetRenewAuto sets RenewAuto field to given value.

### HasRenewAuto

`func (o *DomainPurchase) HasRenewAuto() bool`

HasRenewAuto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


