# LegalAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgreementKey** | **string** | Unique identifier for the legal agreement | 
**Content** | **string** | Contents of the legal agreement, suitable for embedding | 
**Title** | **string** | Title of the legal agreement | 
**Url** | Pointer to **string** | URL to a page containing the legal agreement | [optional] 

## Methods

### NewLegalAgreement

`func NewLegalAgreement(agreementKey string, content string, title string, ) *LegalAgreement`

NewLegalAgreement instantiates a new LegalAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLegalAgreementWithDefaults

`func NewLegalAgreementWithDefaults() *LegalAgreement`

NewLegalAgreementWithDefaults instantiates a new LegalAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgreementKey

`func (o *LegalAgreement) GetAgreementKey() string`

GetAgreementKey returns the AgreementKey field if non-nil, zero value otherwise.

### GetAgreementKeyOk

`func (o *LegalAgreement) GetAgreementKeyOk() (*string, bool)`

GetAgreementKeyOk returns a tuple with the AgreementKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementKey

`func (o *LegalAgreement) SetAgreementKey(v string)`

SetAgreementKey sets AgreementKey field to given value.


### GetContent

`func (o *LegalAgreement) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *LegalAgreement) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *LegalAgreement) SetContent(v string)`

SetContent sets Content field to given value.


### GetTitle

`func (o *LegalAgreement) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LegalAgreement) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LegalAgreement) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *LegalAgreement) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LegalAgreement) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LegalAgreement) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LegalAgreement) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


