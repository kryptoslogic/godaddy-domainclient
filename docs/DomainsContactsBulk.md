# DomainsContactsBulk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactAdmin** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactBilling** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactPresence** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactRegistrant** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactTech** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**Domains** | **[]string** | An array of domain names to be validated against. Alternatively, you can specify the extracted tlds. However, full domain names are required if the tld is &#x60;uk&#x60; | 
**EntityType** | Pointer to **string** | Canadian Presence Requirement (CA) | [optional] 

## Methods

### NewDomainsContactsBulk

`func NewDomainsContactsBulk(domains []string, ) *DomainsContactsBulk`

NewDomainsContactsBulk instantiates a new DomainsContactsBulk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainsContactsBulkWithDefaults

`func NewDomainsContactsBulkWithDefaults() *DomainsContactsBulk`

NewDomainsContactsBulkWithDefaults instantiates a new DomainsContactsBulk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactAdmin

`func (o *DomainsContactsBulk) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainsContactsBulk) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainsContactsBulk) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.

### HasContactAdmin

`func (o *DomainsContactsBulk) HasContactAdmin() bool`

HasContactAdmin returns a boolean if a field has been set.

### GetContactBilling

`func (o *DomainsContactsBulk) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainsContactsBulk) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainsContactsBulk) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.

### HasContactBilling

`func (o *DomainsContactsBulk) HasContactBilling() bool`

HasContactBilling returns a boolean if a field has been set.

### GetContactPresence

`func (o *DomainsContactsBulk) GetContactPresence() Contact`

GetContactPresence returns the ContactPresence field if non-nil, zero value otherwise.

### GetContactPresenceOk

`func (o *DomainsContactsBulk) GetContactPresenceOk() (*Contact, bool)`

GetContactPresenceOk returns a tuple with the ContactPresence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPresence

`func (o *DomainsContactsBulk) SetContactPresence(v Contact)`

SetContactPresence sets ContactPresence field to given value.

### HasContactPresence

`func (o *DomainsContactsBulk) HasContactPresence() bool`

HasContactPresence returns a boolean if a field has been set.

### GetContactRegistrant

`func (o *DomainsContactsBulk) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainsContactsBulk) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainsContactsBulk) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.

### HasContactRegistrant

`func (o *DomainsContactsBulk) HasContactRegistrant() bool`

HasContactRegistrant returns a boolean if a field has been set.

### GetContactTech

`func (o *DomainsContactsBulk) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainsContactsBulk) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainsContactsBulk) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.

### HasContactTech

`func (o *DomainsContactsBulk) HasContactTech() bool`

HasContactTech returns a boolean if a field has been set.

### GetDomains

`func (o *DomainsContactsBulk) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DomainsContactsBulk) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DomainsContactsBulk) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetEntityType

`func (o *DomainsContactsBulk) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DomainsContactsBulk) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DomainsContactsBulk) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DomainsContactsBulk) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


