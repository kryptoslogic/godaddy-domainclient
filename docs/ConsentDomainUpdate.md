# ConsentDomainUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreedAt** | **string** | Timestamp indicating when the end-user consented to these agreements | 
**AgreedBy** | **string** | Originating client IP address of the end-user&#39;s computer when they consented to the agreements | 
**AgreementKeys** | **[]string** | Unique identifiers of the agreements to which the end-user has agreed, as required by the elements being updated&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;EXPOSE_WHOIS&lt;/strong&gt; - Required when the exposeWhois field is updated to true&lt;/li&gt;&lt;/ul&gt; | 

## Methods

### NewConsentDomainUpdate

`func NewConsentDomainUpdate(agreedAt string, agreedBy string, agreementKeys []string, ) *ConsentDomainUpdate`

NewConsentDomainUpdate instantiates a new ConsentDomainUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentDomainUpdateWithDefaults

`func NewConsentDomainUpdateWithDefaults() *ConsentDomainUpdate`

NewConsentDomainUpdateWithDefaults instantiates a new ConsentDomainUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreedAt

`func (o *ConsentDomainUpdate) GetAgreedAt() string`

GetAgreedAt returns the AgreedAt field if non-nil, zero value otherwise.

### GetAgreedAtOk

`func (o *ConsentDomainUpdate) GetAgreedAtOk() (*string, bool)`

GetAgreedAtOk returns a tuple with the AgreedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreedAt

`func (o *ConsentDomainUpdate) SetAgreedAt(v string)`

SetAgreedAt sets AgreedAt field to given value.


### GetAgreedBy

`func (o *ConsentDomainUpdate) GetAgreedBy() string`

GetAgreedBy returns the AgreedBy field if non-nil, zero value otherwise.

### GetAgreedByOk

`func (o *ConsentDomainUpdate) GetAgreedByOk() (*string, bool)`

GetAgreedByOk returns a tuple with the AgreedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreedBy

`func (o *ConsentDomainUpdate) SetAgreedBy(v string)`

SetAgreedBy sets AgreedBy field to given value.


### GetAgreementKeys

`func (o *ConsentDomainUpdate) GetAgreementKeys() []string`

GetAgreementKeys returns the AgreementKeys field if non-nil, zero value otherwise.

### GetAgreementKeysOk

`func (o *ConsentDomainUpdate) GetAgreementKeysOk() (*[]string, bool)`

GetAgreementKeysOk returns a tuple with the AgreementKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementKeys

`func (o *ConsentDomainUpdate) SetAgreementKeys(v []string)`

SetAgreementKeys sets AgreementKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


