# DomainContacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactAdmin** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactBilling** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactRegistrant** | [**Contact**](Contact.md) |  | 
**ContactTech** | Pointer to [**Contact**](Contact.md) |  | [optional] 

## Methods

### NewDomainContacts

`func NewDomainContacts(contactRegistrant Contact, ) *DomainContacts`

NewDomainContacts instantiates a new DomainContacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainContactsWithDefaults

`func NewDomainContactsWithDefaults() *DomainContacts`

NewDomainContactsWithDefaults instantiates a new DomainContacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactAdmin

`func (o *DomainContacts) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainContacts) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainContacts) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.

### HasContactAdmin

`func (o *DomainContacts) HasContactAdmin() bool`

HasContactAdmin returns a boolean if a field has been set.

### GetContactBilling

`func (o *DomainContacts) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainContacts) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainContacts) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.

### HasContactBilling

`func (o *DomainContacts) HasContactBilling() bool`

HasContactBilling returns a boolean if a field has been set.

### GetContactRegistrant

`func (o *DomainContacts) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainContacts) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainContacts) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.


### GetContactTech

`func (o *DomainContacts) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainContacts) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainContacts) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.

### HasContactTech

`func (o *DomainContacts) HasContactTech() bool`

HasContactTech returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


