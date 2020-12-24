# DomainSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | Pointer to **string** | Authorization code for transferring the Domain | [optional] 
**ContactAdmin** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactBilling** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**ContactRegistrant** | [**Contact**](Contact.md) |  | 
**ContactTech** | Pointer to [**Contact**](Contact.md) |  | [optional] 
**CreatedAt** | **time.Time** | Date and time when this domain was created | 
**DeletedAt** | Pointer to **time.Time** | Date and time when this domain was deleted | [optional] 
**TransferAwayEligibleAt** | Pointer to **time.Time** | Date and time when this domain is eligible to transfer | [optional] 
**Domain** | **string** | Name of the domain | 
**DomainId** | **float64** | Unique identifier for this Domain | 
**ExpirationProtected** | **bool** | Whether or not the domain is protected from expiration | 
**Expires** | Pointer to **time.Time** | Date and time when this domain will expire | [optional] 
**ExposeWhois** | Pointer to **bool** | Whether or not the domain contact details should be shown in the WHOIS | [optional] 
**HoldRegistrar** | **bool** | Whether or not the domain is on-hold by the registrar | 
**Locked** | **bool** | Whether or not the domain is locked to prevent transfers | 
**NameServers** | Pointer to **[]string** | Fully-qualified domain names for DNS servers | [optional] 
**Privacy** | **bool** | Whether or not the domain has privacy protection | 
**RenewAuto** | **bool** | Whether or not the domain is configured to automatically renew | 
**RenewDeadline** | **time.Time** | Date the domain must renew on | 
**Renewable** | Pointer to **bool** | Whether or not the domain is eligble for renewal based on status | [optional] 
**Status** | **string** | Processing status of the domain&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;ACTIVE&lt;/strong&gt; - All is well&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;AWAITING*&lt;/strong&gt; - System is waiting for the end-user to complete an action&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CANCELLED*&lt;/strong&gt; - Domain has been cancelled, and may or may not be reclaimable&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CONFISCATED&lt;/strong&gt; - Domain has been confiscated, usually for abuse, chargeback, or fraud&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;DISABLED*&lt;/strong&gt; - Domain has been disabled&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXCLUDED*&lt;/strong&gt; - Domain has been excluded from Firehose registration&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXPIRED*&lt;/strong&gt; - Domain has expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;FAILED*&lt;/strong&gt; - Domain has failed a required action, and the system is no longer retrying&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;HELD*&lt;/strong&gt; - Domain has been placed on hold, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;LOCKED*&lt;/strong&gt; - Domain has been locked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PARKED*&lt;/strong&gt; - Domain has been parked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PENDING*&lt;/strong&gt; - Domain is working its way through an automated workflow&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;RESERVED*&lt;/strong&gt; - Domain is reserved, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REVERTED&lt;/strong&gt; - Domain has been reverted, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;SUSPENDED*&lt;/strong&gt; - Domain has been suspended, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;TRANSFERRED*&lt;/strong&gt; - Domain has been transferred out&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNKNOWN&lt;/strong&gt; - Domain is in an unknown state&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNLOCKED*&lt;/strong&gt; - Domain has been unlocked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNPARKED*&lt;/strong&gt; - Domain has been unparked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UPDATED*&lt;/strong&gt; - Domain ownership has been transferred to another account&lt;/li&gt; &lt;/ul&gt; | 
**TransferProtected** | **bool** | Whether or not the domain is protected from transfer | 

## Methods

### NewDomainSummary

`func NewDomainSummary(contactRegistrant Contact, createdAt time.Time, domain string, domainId float64, expirationProtected bool, holdRegistrar bool, locked bool, privacy bool, renewAuto bool, renewDeadline time.Time, status string, transferProtected bool, ) *DomainSummary`

NewDomainSummary instantiates a new DomainSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainSummaryWithDefaults

`func NewDomainSummaryWithDefaults() *DomainSummary`

NewDomainSummaryWithDefaults instantiates a new DomainSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *DomainSummary) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainSummary) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainSummary) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *DomainSummary) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetContactAdmin

`func (o *DomainSummary) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainSummary) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainSummary) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.

### HasContactAdmin

`func (o *DomainSummary) HasContactAdmin() bool`

HasContactAdmin returns a boolean if a field has been set.

### GetContactBilling

`func (o *DomainSummary) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainSummary) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainSummary) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.

### HasContactBilling

`func (o *DomainSummary) HasContactBilling() bool`

HasContactBilling returns a boolean if a field has been set.

### GetContactRegistrant

`func (o *DomainSummary) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainSummary) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainSummary) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.


### GetContactTech

`func (o *DomainSummary) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainSummary) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainSummary) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.

### HasContactTech

`func (o *DomainSummary) HasContactTech() bool`

HasContactTech returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DomainSummary) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainSummary) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainSummary) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *DomainSummary) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DomainSummary) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DomainSummary) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DomainSummary) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetTransferAwayEligibleAt

`func (o *DomainSummary) GetTransferAwayEligibleAt() time.Time`

GetTransferAwayEligibleAt returns the TransferAwayEligibleAt field if non-nil, zero value otherwise.

### GetTransferAwayEligibleAtOk

`func (o *DomainSummary) GetTransferAwayEligibleAtOk() (*time.Time, bool)`

GetTransferAwayEligibleAtOk returns a tuple with the TransferAwayEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAwayEligibleAt

`func (o *DomainSummary) SetTransferAwayEligibleAt(v time.Time)`

SetTransferAwayEligibleAt sets TransferAwayEligibleAt field to given value.

### HasTransferAwayEligibleAt

`func (o *DomainSummary) HasTransferAwayEligibleAt() bool`

HasTransferAwayEligibleAt returns a boolean if a field has been set.

### GetDomain

`func (o *DomainSummary) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainSummary) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainSummary) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainId

`func (o *DomainSummary) GetDomainId() float64`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainSummary) GetDomainIdOk() (*float64, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainSummary) SetDomainId(v float64)`

SetDomainId sets DomainId field to given value.


### GetExpirationProtected

`func (o *DomainSummary) GetExpirationProtected() bool`

GetExpirationProtected returns the ExpirationProtected field if non-nil, zero value otherwise.

### GetExpirationProtectedOk

`func (o *DomainSummary) GetExpirationProtectedOk() (*bool, bool)`

GetExpirationProtectedOk returns a tuple with the ExpirationProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationProtected

`func (o *DomainSummary) SetExpirationProtected(v bool)`

SetExpirationProtected sets ExpirationProtected field to given value.


### GetExpires

`func (o *DomainSummary) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DomainSummary) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DomainSummary) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *DomainSummary) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExposeWhois

`func (o *DomainSummary) GetExposeWhois() bool`

GetExposeWhois returns the ExposeWhois field if non-nil, zero value otherwise.

### GetExposeWhoisOk

`func (o *DomainSummary) GetExposeWhoisOk() (*bool, bool)`

GetExposeWhoisOk returns a tuple with the ExposeWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeWhois

`func (o *DomainSummary) SetExposeWhois(v bool)`

SetExposeWhois sets ExposeWhois field to given value.

### HasExposeWhois

`func (o *DomainSummary) HasExposeWhois() bool`

HasExposeWhois returns a boolean if a field has been set.

### GetHoldRegistrar

`func (o *DomainSummary) GetHoldRegistrar() bool`

GetHoldRegistrar returns the HoldRegistrar field if non-nil, zero value otherwise.

### GetHoldRegistrarOk

`func (o *DomainSummary) GetHoldRegistrarOk() (*bool, bool)`

GetHoldRegistrarOk returns a tuple with the HoldRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldRegistrar

`func (o *DomainSummary) SetHoldRegistrar(v bool)`

SetHoldRegistrar sets HoldRegistrar field to given value.


### GetLocked

`func (o *DomainSummary) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DomainSummary) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DomainSummary) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetNameServers

`func (o *DomainSummary) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DomainSummary) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DomainSummary) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.

### HasNameServers

`func (o *DomainSummary) HasNameServers() bool`

HasNameServers returns a boolean if a field has been set.

### GetPrivacy

`func (o *DomainSummary) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *DomainSummary) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *DomainSummary) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.


### GetRenewAuto

`func (o *DomainSummary) GetRenewAuto() bool`

GetRenewAuto returns the RenewAuto field if non-nil, zero value otherwise.

### GetRenewAutoOk

`func (o *DomainSummary) GetRenewAutoOk() (*bool, bool)`

GetRenewAutoOk returns a tuple with the RenewAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuto

`func (o *DomainSummary) SetRenewAuto(v bool)`

SetRenewAuto sets RenewAuto field to given value.


### GetRenewDeadline

`func (o *DomainSummary) GetRenewDeadline() time.Time`

GetRenewDeadline returns the RenewDeadline field if non-nil, zero value otherwise.

### GetRenewDeadlineOk

`func (o *DomainSummary) GetRenewDeadlineOk() (*time.Time, bool)`

GetRenewDeadlineOk returns a tuple with the RenewDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDeadline

`func (o *DomainSummary) SetRenewDeadline(v time.Time)`

SetRenewDeadline sets RenewDeadline field to given value.


### GetRenewable

`func (o *DomainSummary) GetRenewable() bool`

GetRenewable returns the Renewable field if non-nil, zero value otherwise.

### GetRenewableOk

`func (o *DomainSummary) GetRenewableOk() (*bool, bool)`

GetRenewableOk returns a tuple with the Renewable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewable

`func (o *DomainSummary) SetRenewable(v bool)`

SetRenewable sets Renewable field to given value.

### HasRenewable

`func (o *DomainSummary) HasRenewable() bool`

HasRenewable returns a boolean if a field has been set.

### GetStatus

`func (o *DomainSummary) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainSummary) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainSummary) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTransferProtected

`func (o *DomainSummary) GetTransferProtected() bool`

GetTransferProtected returns the TransferProtected field if non-nil, zero value otherwise.

### GetTransferProtectedOk

`func (o *DomainSummary) GetTransferProtectedOk() (*bool, bool)`

GetTransferProtectedOk returns a tuple with the TransferProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferProtected

`func (o *DomainSummary) SetTransferProtected(v bool)`

SetTransferProtected sets TransferProtected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


