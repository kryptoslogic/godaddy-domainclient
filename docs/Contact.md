# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressMailing** | [**Address**](Address.md) |  | 
**Email** | **string** |  | 
**Fax** | Pointer to **string** |  | [optional] 
**JobTitle** | Pointer to **string** |  | [optional] 
**NameFirst** | **string** |  | 
**NameLast** | **string** |  | 
**NameMiddle** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Phone** | **string** |  | 

## Methods

### NewContact

`func NewContact(addressMailing Address, email string, nameFirst string, nameLast string, phone string, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressMailing

`func (o *Contact) GetAddressMailing() Address`

GetAddressMailing returns the AddressMailing field if non-nil, zero value otherwise.

### GetAddressMailingOk

`func (o *Contact) GetAddressMailingOk() (*Address, bool)`

GetAddressMailingOk returns a tuple with the AddressMailing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMailing

`func (o *Contact) SetAddressMailing(v Address)`

SetAddressMailing sets AddressMailing field to given value.


### GetEmail

`func (o *Contact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Contact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Contact) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFax

`func (o *Contact) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *Contact) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *Contact) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *Contact) HasFax() bool`

HasFax returns a boolean if a field has been set.

### GetJobTitle

`func (o *Contact) GetJobTitle() string`

GetJobTitle returns the JobTitle field if non-nil, zero value otherwise.

### GetJobTitleOk

`func (o *Contact) GetJobTitleOk() (*string, bool)`

GetJobTitleOk returns a tuple with the JobTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTitle

`func (o *Contact) SetJobTitle(v string)`

SetJobTitle sets JobTitle field to given value.

### HasJobTitle

`func (o *Contact) HasJobTitle() bool`

HasJobTitle returns a boolean if a field has been set.

### GetNameFirst

`func (o *Contact) GetNameFirst() string`

GetNameFirst returns the NameFirst field if non-nil, zero value otherwise.

### GetNameFirstOk

`func (o *Contact) GetNameFirstOk() (*string, bool)`

GetNameFirstOk returns a tuple with the NameFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameFirst

`func (o *Contact) SetNameFirst(v string)`

SetNameFirst sets NameFirst field to given value.


### GetNameLast

`func (o *Contact) GetNameLast() string`

GetNameLast returns the NameLast field if non-nil, zero value otherwise.

### GetNameLastOk

`func (o *Contact) GetNameLastOk() (*string, bool)`

GetNameLastOk returns a tuple with the NameLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameLast

`func (o *Contact) SetNameLast(v string)`

SetNameLast sets NameLast field to given value.


### GetNameMiddle

`func (o *Contact) GetNameMiddle() string`

GetNameMiddle returns the NameMiddle field if non-nil, zero value otherwise.

### GetNameMiddleOk

`func (o *Contact) GetNameMiddleOk() (*string, bool)`

GetNameMiddleOk returns a tuple with the NameMiddle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameMiddle

`func (o *Contact) SetNameMiddle(v string)`

SetNameMiddle sets NameMiddle field to given value.

### HasNameMiddle

`func (o *Contact) HasNameMiddle() bool`

HasNameMiddle returns a boolean if a field has been set.

### GetOrganization

`func (o *Contact) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Contact) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Contact) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Contact) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetPhone

`func (o *Contact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Contact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Contact) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


