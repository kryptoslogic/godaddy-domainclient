# DomainDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code for transferring the Domain | 
**ContactAdmin** | [**Contact**](Contact.md) |  | 
**ContactBilling** | [**Contact**](Contact.md) |  | 
**ContactRegistrant** | [**Contact**](Contact.md) |  | 
**ContactTech** | [**Contact**](Contact.md) |  | 
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
**NameServers** | **[]string** | Fully-qualified domain names for DNS servers | 
**Privacy** | **bool** | Whether or not the domain has privacy protection | 
**RenewAuto** | **bool** | Whether or not the domain is configured to automatically renew | 
**RenewDeadline** | **time.Time** | Date the domain must renew on | 
**Status** | **string** | Processing status of the domain&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;ACTIVE&lt;/strong&gt; - All is well&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;AWAITING*&lt;/strong&gt; - System is waiting for the end-user to complete an action&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CANCELLED*&lt;/strong&gt; - Domain has been cancelled, and may or may not be reclaimable&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;CONFISCATED&lt;/strong&gt; - Domain has been confiscated, usually for abuse, chargeback, or fraud&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;DISABLED*&lt;/strong&gt; - Domain has been disabled&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXCLUDED*&lt;/strong&gt; - Domain has been excluded from Firehose registration&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXPIRED*&lt;/strong&gt; - Domain has expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;FAILED*&lt;/strong&gt; - Domain has failed a required action, and the system is no longer retrying&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;HELD*&lt;/strong&gt; - Domain has been placed on hold, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;LOCKED*&lt;/strong&gt; - Domain has been locked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PARKED*&lt;/strong&gt; - Domain has been parked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PENDING*&lt;/strong&gt; - Domain is working its way through an automated workflow&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;RESERVED*&lt;/strong&gt; - Domain is reserved, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REVERTED&lt;/strong&gt; - Domain has been reverted, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;SUSPENDED*&lt;/strong&gt; - Domain has been suspended, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;TRANSFERRED*&lt;/strong&gt; - Domain has been transferred out&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNKNOWN&lt;/strong&gt; - Domain is in an unknown state&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNLOCKED*&lt;/strong&gt; - Domain has been unlocked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNPARKED*&lt;/strong&gt; - Domain has been unparked, and likely requires intervention from Support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UPDATED*&lt;/strong&gt; - Domain ownership has been transferred to another account&lt;/li&gt; &lt;/ul&gt; | 
**SubaccountId** | Pointer to **string** | Reseller subaccount shopperid who can manage the domain | [optional] 
**TransferProtected** | **bool** | Whether or not the domain is protected from transfer | 
**Verifications** | Pointer to [**VerificationsDomain**](VerificationsDomain.md) |  | [optional] 

## Methods

### NewDomainDetail

`func NewDomainDetail(authCode string, contactAdmin Contact, contactBilling Contact, contactRegistrant Contact, contactTech Contact, createdAt time.Time, domain string, domainId float64, expirationProtected bool, holdRegistrar bool, locked bool, nameServers []string, privacy bool, renewAuto bool, renewDeadline time.Time, status string, transferProtected bool, ) *DomainDetail`

NewDomainDetail instantiates a new DomainDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainDetailWithDefaults

`func NewDomainDetailWithDefaults() *DomainDetail`

NewDomainDetailWithDefaults instantiates a new DomainDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthCode

`func (o *DomainDetail) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *DomainDetail) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *DomainDetail) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.


### GetContactAdmin

`func (o *DomainDetail) GetContactAdmin() Contact`

GetContactAdmin returns the ContactAdmin field if non-nil, zero value otherwise.

### GetContactAdminOk

`func (o *DomainDetail) GetContactAdminOk() (*Contact, bool)`

GetContactAdminOk returns a tuple with the ContactAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactAdmin

`func (o *DomainDetail) SetContactAdmin(v Contact)`

SetContactAdmin sets ContactAdmin field to given value.


### GetContactBilling

`func (o *DomainDetail) GetContactBilling() Contact`

GetContactBilling returns the ContactBilling field if non-nil, zero value otherwise.

### GetContactBillingOk

`func (o *DomainDetail) GetContactBillingOk() (*Contact, bool)`

GetContactBillingOk returns a tuple with the ContactBilling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBilling

`func (o *DomainDetail) SetContactBilling(v Contact)`

SetContactBilling sets ContactBilling field to given value.


### GetContactRegistrant

`func (o *DomainDetail) GetContactRegistrant() Contact`

GetContactRegistrant returns the ContactRegistrant field if non-nil, zero value otherwise.

### GetContactRegistrantOk

`func (o *DomainDetail) GetContactRegistrantOk() (*Contact, bool)`

GetContactRegistrantOk returns a tuple with the ContactRegistrant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactRegistrant

`func (o *DomainDetail) SetContactRegistrant(v Contact)`

SetContactRegistrant sets ContactRegistrant field to given value.


### GetContactTech

`func (o *DomainDetail) GetContactTech() Contact`

GetContactTech returns the ContactTech field if non-nil, zero value otherwise.

### GetContactTechOk

`func (o *DomainDetail) GetContactTechOk() (*Contact, bool)`

GetContactTechOk returns a tuple with the ContactTech field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactTech

`func (o *DomainDetail) SetContactTech(v Contact)`

SetContactTech sets ContactTech field to given value.


### GetCreatedAt

`func (o *DomainDetail) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DomainDetail) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DomainDetail) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeletedAt

`func (o *DomainDetail) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DomainDetail) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DomainDetail) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DomainDetail) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetTransferAwayEligibleAt

`func (o *DomainDetail) GetTransferAwayEligibleAt() time.Time`

GetTransferAwayEligibleAt returns the TransferAwayEligibleAt field if non-nil, zero value otherwise.

### GetTransferAwayEligibleAtOk

`func (o *DomainDetail) GetTransferAwayEligibleAtOk() (*time.Time, bool)`

GetTransferAwayEligibleAtOk returns a tuple with the TransferAwayEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferAwayEligibleAt

`func (o *DomainDetail) SetTransferAwayEligibleAt(v time.Time)`

SetTransferAwayEligibleAt sets TransferAwayEligibleAt field to given value.

### HasTransferAwayEligibleAt

`func (o *DomainDetail) HasTransferAwayEligibleAt() bool`

HasTransferAwayEligibleAt returns a boolean if a field has been set.

### GetDomain

`func (o *DomainDetail) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainDetail) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainDetail) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetDomainId

`func (o *DomainDetail) GetDomainId() float64`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DomainDetail) GetDomainIdOk() (*float64, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DomainDetail) SetDomainId(v float64)`

SetDomainId sets DomainId field to given value.


### GetExpirationProtected

`func (o *DomainDetail) GetExpirationProtected() bool`

GetExpirationProtected returns the ExpirationProtected field if non-nil, zero value otherwise.

### GetExpirationProtectedOk

`func (o *DomainDetail) GetExpirationProtectedOk() (*bool, bool)`

GetExpirationProtectedOk returns a tuple with the ExpirationProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationProtected

`func (o *DomainDetail) SetExpirationProtected(v bool)`

SetExpirationProtected sets ExpirationProtected field to given value.


### GetExpires

`func (o *DomainDetail) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *DomainDetail) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *DomainDetail) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *DomainDetail) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetExposeWhois

`func (o *DomainDetail) GetExposeWhois() bool`

GetExposeWhois returns the ExposeWhois field if non-nil, zero value otherwise.

### GetExposeWhoisOk

`func (o *DomainDetail) GetExposeWhoisOk() (*bool, bool)`

GetExposeWhoisOk returns a tuple with the ExposeWhois field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeWhois

`func (o *DomainDetail) SetExposeWhois(v bool)`

SetExposeWhois sets ExposeWhois field to given value.

### HasExposeWhois

`func (o *DomainDetail) HasExposeWhois() bool`

HasExposeWhois returns a boolean if a field has been set.

### GetHoldRegistrar

`func (o *DomainDetail) GetHoldRegistrar() bool`

GetHoldRegistrar returns the HoldRegistrar field if non-nil, zero value otherwise.

### GetHoldRegistrarOk

`func (o *DomainDetail) GetHoldRegistrarOk() (*bool, bool)`

GetHoldRegistrarOk returns a tuple with the HoldRegistrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldRegistrar

`func (o *DomainDetail) SetHoldRegistrar(v bool)`

SetHoldRegistrar sets HoldRegistrar field to given value.


### GetLocked

`func (o *DomainDetail) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *DomainDetail) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *DomainDetail) SetLocked(v bool)`

SetLocked sets Locked field to given value.


### GetNameServers

`func (o *DomainDetail) GetNameServers() []string`

GetNameServers returns the NameServers field if non-nil, zero value otherwise.

### GetNameServersOk

`func (o *DomainDetail) GetNameServersOk() (*[]string, bool)`

GetNameServersOk returns a tuple with the NameServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameServers

`func (o *DomainDetail) SetNameServers(v []string)`

SetNameServers sets NameServers field to given value.


### GetPrivacy

`func (o *DomainDetail) GetPrivacy() bool`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *DomainDetail) GetPrivacyOk() (*bool, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *DomainDetail) SetPrivacy(v bool)`

SetPrivacy sets Privacy field to given value.


### GetRenewAuto

`func (o *DomainDetail) GetRenewAuto() bool`

GetRenewAuto returns the RenewAuto field if non-nil, zero value otherwise.

### GetRenewAutoOk

`func (o *DomainDetail) GetRenewAutoOk() (*bool, bool)`

GetRenewAutoOk returns a tuple with the RenewAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewAuto

`func (o *DomainDetail) SetRenewAuto(v bool)`

SetRenewAuto sets RenewAuto field to given value.


### GetRenewDeadline

`func (o *DomainDetail) GetRenewDeadline() time.Time`

GetRenewDeadline returns the RenewDeadline field if non-nil, zero value otherwise.

### GetRenewDeadlineOk

`func (o *DomainDetail) GetRenewDeadlineOk() (*time.Time, bool)`

GetRenewDeadlineOk returns a tuple with the RenewDeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewDeadline

`func (o *DomainDetail) SetRenewDeadline(v time.Time)`

SetRenewDeadline sets RenewDeadline field to given value.


### GetStatus

`func (o *DomainDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DomainDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DomainDetail) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSubaccountId

`func (o *DomainDetail) GetSubaccountId() string`

GetSubaccountId returns the SubaccountId field if non-nil, zero value otherwise.

### GetSubaccountIdOk

`func (o *DomainDetail) GetSubaccountIdOk() (*string, bool)`

GetSubaccountIdOk returns a tuple with the SubaccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubaccountId

`func (o *DomainDetail) SetSubaccountId(v string)`

SetSubaccountId sets SubaccountId field to given value.

### HasSubaccountId

`func (o *DomainDetail) HasSubaccountId() bool`

HasSubaccountId returns a boolean if a field has been set.

### GetTransferProtected

`func (o *DomainDetail) GetTransferProtected() bool`

GetTransferProtected returns the TransferProtected field if non-nil, zero value otherwise.

### GetTransferProtectedOk

`func (o *DomainDetail) GetTransferProtectedOk() (*bool, bool)`

GetTransferProtectedOk returns a tuple with the TransferProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferProtected

`func (o *DomainDetail) SetTransferProtected(v bool)`

SetTransferProtected sets TransferProtected field to given value.


### GetVerifications

`func (o *DomainDetail) GetVerifications() VerificationsDomain`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *DomainDetail) GetVerificationsOk() (*VerificationsDomain, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *DomainDetail) SetVerifications(v VerificationsDomain)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *DomainDetail) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


