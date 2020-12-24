# VerificationRealName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the real name verification&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;APPROVED&lt;/strong&gt; - All is well&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;PENDING&lt;/strong&gt; - Real name verification is working its way through the workflow&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_DOCUMENT_OUTDATED&lt;/strong&gt; - Local government verification shows there is a newer version of your document.  Upload the latest version of the document and retry real name verification&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_EXPIRED_BUSINESS_LICENSE&lt;/strong&gt; - Business license is expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_EXPIRED_ORGANIZATION_CODE&lt;/strong&gt; - Organization code certificate number has expired&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_ILLEGIBLE_DOCUMENT_NAME&lt;/strong&gt; - There isn’t a clear name on your uploaded document, please upload a different document to retry real name verification&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_ILLEGIBLE_IDENTIFICATION&lt;/strong&gt; - Registrant identification is not clear.  Upload a better image to retry&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INCOMPLETE_IDENTIFICATION&lt;/strong&gt; - Registrant identification is incomplete&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INCOMPLETE_REGISTRATION_LETTER&lt;/strong&gt; - Registration letter is incomplete&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INCONSISTENT_IDENTITY_CARD&lt;/strong&gt; - Provided identity card is inconsistent with the identity card on record&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INCONSISTENT_ORGANIZATION_CODE&lt;/strong&gt; - Provided organization information is inconsistent with the results obtained using the submitted organization code&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INCONSISTENT_REGISTRANT_NAME&lt;/strong&gt; - Name on the registrant identification does not match the name in the system&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INVALID_BUSINESS_LICENSE_OR_ORGANIZATION_CODE&lt;/strong&gt; - Your document contains an invalid business license or organization code certificate number&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_INVALID_DOCUMENT&lt;/strong&gt; - Document is invalid.  Please upload another document to retry real name verification&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_BUSINESS_ID&lt;/strong&gt; - Business id does not match the business id in the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_BUSINESS_NAME&lt;/strong&gt; - Business name does not match the business name in the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_DOCUMENT_ID&lt;/strong&gt; - Document id does not match the id in the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_DOCUMENT_NAME&lt;/strong&gt; - Document name does not match the name in the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_DOCUMENT_TYPE&lt;/strong&gt; - Document type does not match the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_REGISTRANT_INFO&lt;/strong&gt; - The information provided for the registrant does not match the document&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_REGISTRANT_LOCALITY&lt;/strong&gt; - Registrant region is overseas, but a local identity document was provided&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_MISMATCH_REGISTRANT_NAME&lt;/strong&gt; - Registrant name has been changed, so the request must be resubmitted&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_UNABLE_TO_OPEN&lt;/strong&gt; - Registrant identification could not be opened.  Please upload the document again to retry real name verification&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_UNABLE_TO_VERIFY&lt;/strong&gt; - Unable to initiate verification.  Please upload the document again to retry real name verification&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REJECTED_UNKNOWN_ERROR&lt;/strong&gt; - Document was rejected due to an unknown error. For more information, contact customer support&lt;/li&gt; &lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;UNABLE_TO_RETRIEVE_STATUS&lt;/strong&gt; - Unable to retrieve status for the real name verification process.  Retry, if this status persists, contact customer support&lt;/li&gt; &lt;/ul&gt; | 

## Methods

### NewVerificationRealName

`func NewVerificationRealName(status string, ) *VerificationRealName`

NewVerificationRealName instantiates a new VerificationRealName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRealNameWithDefaults

`func NewVerificationRealNameWithDefaults() *VerificationRealName`

NewVerificationRealNameWithDefaults instantiates a new VerificationRealName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VerificationRealName) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerificationRealName) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerificationRealName) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


