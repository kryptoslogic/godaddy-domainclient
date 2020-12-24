# DomainTransferIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code from registrar for transferring a domain | 
**Consent** | [**Consent**](Consent.md) |  | 
**Period** | Pointer to **int32** | Can be more than 1 but no more than 10 years total including current registration length | [optional] 
**Privacy** | Pointer to **bool** | Whether or not privacy has been requested | [optional] [default to false]
**RenewAuto** | Pointer to **bool** | Whether or not the domain should be configured to automatically renew | [optional] [default to true]
**ContactAdmin** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactBilling** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactRegistrant** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactTech** | Pointer to [**Contact**](Contact.md) |  | [optional] 

## Methods

### NewDomainTransferIn

`func NewDomainTransferIn(authCode string, consent Consent, ) *DomainTransferIn`

NewDomainTransferIn instantiates a new DomainTransferIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainTransferInWithDefaults

`func NewDomainTransferInWithDefaults() *DomainTransferIn`

NewDomainTransferInWithDefaults instantiates a new DomainTransferIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *DomainTransferIn) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainTransferIn) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainTransferIn) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetConsent

`func (o *DomainTransferIn) GetConsent() Consent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *DomainTransferIn) GetConsentOk() (*Consent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *DomainTransferIn) SetConsent(v Consent)`

SetConsent sets Consent field to given value.


### GetPeriod

`func (o *DomainTransferIn) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *DomainTransferIn) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *DomainTransferIn) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *DomainTransferIn) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetPrivacy

`func (o *DomainTransferIn) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *DomainTransferIn) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *DomainTransferIn) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *DomainTransferIn) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetRenewAuto

`func (o *DomainTransferIn) GetRenewAuto() bool`

GetRenewAuto returns the RenewAuto field if non-nil, zero value otherwise.

### GetRenewAutoOk

`func (o *DomainTransferIn) GetRenewAutoOk() (*bool, bool)`

GetRenewAutoOk returns a tuple with the RenewAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuto

`func (o *DomainTransferIn) SetRenewAuto(v bool)`

SetRenewAuto sets RenewAuto field to given value.

### HasRenewAuto

`func (o *DomainTransferIn) HasRenewAuto() bool`

HasRenewAuto returns a boolean if a field has been set.

### GetContactAdmin

`func (o *DomainTransferIn) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainTransferIn) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainTransferIn) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.

### HasContactAdmin

`func (o *DomainTransferIn) HasContactAdmin() bool`

HasContactAdmin returns a boolean if a field has been set.

### GetContactBilling

`func (o *DomainTransferIn) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainTransferIn) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainTransferIn) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.

### HasContactBilling

`func (o *DomainTransferIn) HasContactBilling() bool`

HasContactBilling returns a boolean if a field has been set.

### GetContactRegistrant

`func (o *DomainTransferIn) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainTransferIn) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainTransferIn) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.

### HasContactRegistrant

`func (o *DomainTransferIn) HasContactRegistrant() bool`

HasContactRegistrant returns a boolean if a field has been set.

### GetContactTech

`func (o *DomainTransferIn) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainTransferIn) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainTransferIn) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.

### HasContactTech

`func (o *DomainTransferIn) HasContactTech() bool`

HasContactTech returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


