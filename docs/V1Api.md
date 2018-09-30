# \V1Api

All URIs are relative to *https://api.ote-godaddy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Available**](V1Api.md#Available) | **Get** /v1/domains/available | Determine whether or not the specified domain is available for purchase
[**AvailableBulk**](V1Api.md#AvailableBulk) | **Post** /v1/domains/available | Determine whether or not the specified domains are available for purchase
[**Cancel**](V1Api.md#Cancel) | **Delete** /v1/domains/{domain} | Cancel a purchased domain
[**CancelPrivacy**](V1Api.md#CancelPrivacy) | **Delete** /v1/domains/{domain}/privacy | Submit a privacy cancellation request for the given domain
[**ContactsValidate**](V1Api.md#ContactsValidate) | **Post** /v1/domains/contacts/validate | Validate the request body using the Domain Contact Validation Schema for specified domains.
[**Get**](V1Api.md#Get) | **Get** /v1/domains/{domain} | Retrieve details for the specified Domain
[**GetAgreement**](V1Api.md#GetAgreement) | **Get** /v1/domains/agreements | Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons
[**List**](V1Api.md#List) | **Get** /v1/domains | Retrieve a list of Domains for the specified Shopper
[**Purchase**](V1Api.md#Purchase) | **Post** /v1/domains/purchase | Purchase and register the specified Domain
[**PurchasePrivacy**](V1Api.md#PurchasePrivacy) | **Post** /v1/domains/{domain}/privacy/purchase | Purchase privacy for a specified domain
[**RecordAdd**](V1Api.md#RecordAdd) | **Patch** /v1/domains/{domain}/records | Add the specified DNS Records to the specified Domain
[**RecordGet**](V1Api.md#RecordGet) | **Get** /v1/domains/{domain}/records/{type}/{name} | Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name
[**RecordReplace**](V1Api.md#RecordReplace) | **Put** /v1/domains/{domain}/records | Replace all DNS Records for the specified Domain
[**RecordReplaceType**](V1Api.md#RecordReplaceType) | **Put** /v1/domains/{domain}/records/{type} | Replace all DNS Records for the specified Domain with the specified Type
[**RecordReplaceTypeName**](V1Api.md#RecordReplaceTypeName) | **Put** /v1/domains/{domain}/records/{type}/{name} | Replace all DNS Records for the specified Domain with the specified Type and Name
[**Renew**](V1Api.md#Renew) | **Post** /v1/domains/{domain}/renew | Renew the specified Domain
[**Schema**](V1Api.md#Schema) | **Get** /v1/domains/purchase/schema/{tld} | Retrieve the schema to be submitted when registering a Domain for the specified TLD
[**Suggest**](V1Api.md#Suggest) | **Get** /v1/domains/suggest | Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper&#39;s purchase history
[**Tlds**](V1Api.md#Tlds) | **Get** /v1/domains/tlds | Retrieves a list of TLDs supported and enabled for sale
[**TransferIn**](V1Api.md#TransferIn) | **Post** /v1/domains/{domain}/transfer | Purchase and start or restart transfer process
[**Update**](V1Api.md#Update) | **Patch** /v1/domains/{domain} | Update details for the specified Domain
[**UpdateContacts**](V1Api.md#UpdateContacts) | **Patch** /v1/domains/{domain}/contacts | Update domain
[**Validate**](V1Api.md#Validate) | **Post** /v1/domains/purchase/validate | Validate the request body using the Domain Purchase Schema for the specified TLD
[**VerifyEmail**](V1Api.md#VerifyEmail) | **Post** /v1/domains/{domain}/verifyRegistrantEmail | Re-send Contact E-mail Verification for specified Domain


# **Available**
> DomainAvailableResponse Available(ctx, domain, optional)
Determine whether or not the specified domain is available for purchase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain name whose availability is to be checked | 
 **optional** | ***AvailableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvailableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkType** | **optional.String**| Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [default to FAST]
 **forTransfer** | **optional.Bool**| Whether or not to include domains available for transfer. If set to True, checkType is ignored | [default to false]

### Return type

[**DomainAvailableResponse**](DomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AvailableBulk**
> DomainAvailableBulk AvailableBulk(ctx, domains, optional)
Determine whether or not the specified domains are available for purchase

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domains** | **[]string**| Domain names for which to check availability | 
 **optional** | ***AvailableBulkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvailableBulkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **checkType** | **optional.String**| Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [default to FAST]

### Return type

[**DomainAvailableBulk**](DomainAvailableBulk.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Cancel**
> Cancel(ctx, domain)
Cancel a purchased domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain to cancel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelPrivacy**
> CancelPrivacy(ctx, domain, optional)
Submit a privacy cancellation request for the given domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose privacy is to be cancelled | 
 **optional** | ***CancelPrivacyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CancelPrivacyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **optional.String**| Shopper ID of the owner of the domain | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContactsValidate**
> ContactsValidate(ctx, body, optional)
Validate the request body using the Domain Contact Validation Schema for specified domains.

All contacts specified in request will be validated against all domains specifed in \"domains\". As an alternative, you can also pass in tlds, with the exception of `uk`, which requires full domain names

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainsContactsBulk**](DomainsContactsBulk.md)| An instance document expected for domains contacts validation | 
 **optional** | ***ContactsValidateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContactsValidateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPrivateLabelId** | **optional.Int32**| PrivateLabelId to operate as, if different from JWT | [default to 1]
 **marketId** | **optional.String**| MarketId in which the request is being made, and for which responses should be localized | [default to en-US]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Get**
> DomainDetail Get(ctx, domain, optional)
Retrieve details for the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain name whose details are to be retrieved | 
 **optional** | ***GetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **optional.String**| Shopper ID expected to own the specified domain | 

### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAgreement**
> []LegalAgreement GetAgreement(ctx, tlds, privacy, optional)
Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tlds** | [**[]string**](string.md)| list of TLDs whose legal agreements are to be retrieved | 
  **privacy** | **bool**| Whether or not privacy has been requested | 
 **optional** | ***GetAgreementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAgreementOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xMarketId** | **optional.String**| Unique identifier of the Market used to retrieve/translate Legal Agreements | [default to en-US]
 **forTransfer** | **optional.Bool**| Whether or not domain tranfer has been requested | 

### Return type

[**[]LegalAgreement**](LegalAgreement.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **List**
> []DomainSummary List(ctx, optional)
Retrieve a list of Domains for the specified Shopper

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **optional.String**| Shopper ID whose domains are to be retrieved | 
 **statuses** | [**optional.Interface of []string**](string.md)| Only include results with &#x60;status&#x60; value in the specified set | 
 **statusGroups** | [**optional.Interface of []string**](string.md)| Only include results with &#x60;status&#x60; value in any of the specified groups | 
 **limit** | **optional.Int32**| Maximum number of domains to return | 
 **marker** | **optional.String**| Marker Domain to use as the offset in results | 
 **includes** | [**optional.Interface of []string**](string.md)| Optional details to be included in the response | 
 **modifiedDate** | **optional.String**| Only include results that have been modified since the specified date | 

### Return type

[**[]DomainSummary**](DomainSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Purchase**
> DomainPurchaseResponse Purchase(ctx, body, optional)
Purchase and register the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainPurchase**](DomainPurchase.md)| An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 
 **optional** | ***PurchaseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchaseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **optional.String**| The Shopper for whom the domain should be purchased | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PurchasePrivacy**
> DomainPurchaseResponse PurchasePrivacy(ctx, domain, body, optional)
Purchase privacy for a specified domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain for which to purchase privacy | 
  **body** | [**PrivacyPurchase**](PrivacyPurchase.md)| Options for purchasing privacy | 
 **optional** | ***PurchasePrivacyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PurchasePrivacyOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| Shopper ID of the owner of the domain | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordAdd**
> RecordAdd(ctx, domain, records, optional)
Add the specified DNS Records to the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose DNS Records are to be augmented | 
  **records** | [**ArrayOfDnsRecord**](ArrayOfDnsRecord.md)| DNS Records to add to whatever currently exists | 
 **optional** | ***RecordAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordGet**
> []DnsRecord RecordGet(ctx, domain, type_, name, optional)
Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose DNS Records are to be retrieved | 
  **type_** | **string**| DNS Record Type for which DNS Records are to be retrieved | 
  **name** | **string**| DNS Record Name for which DNS Records are to be retrieved | 
 **optional** | ***RecordGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xShopperId** | **optional.String**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 
 **offset** | **optional.Int32**| Number of results to skip for pagination | 
 **limit** | **optional.Int32**| Maximum number of items to return | 

### Return type

[**[]DnsRecord**](DNSRecord.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplace**
> RecordReplace(ctx, domain, records, optional)
Replace all DNS Records for the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose DNS Records are to be replaced | 
  **records** | [**[]DnsRecord**](DNSRecord.md)| DNS Records to replace whatever currently exists | 
 **optional** | ***RecordReplaceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordReplaceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplaceType**
> RecordReplaceType(ctx, domain, type_, records, optional)
Replace all DNS Records for the specified Domain with the specified Type

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose DNS Records are to be replaced | 
  **type_** | **string**| DNS Record Type for which DNS Records are to be replaced | 
  **records** | [**[]DnsRecordCreateType**](DNSRecordCreateType.md)| DNS Records to replace whatever currently exists | 
 **optional** | ***RecordReplaceTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordReplaceTypeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xShopperId** | **optional.String**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RecordReplaceTypeName**
> RecordReplaceTypeName(ctx, domain, type_, name, records, optional)
Replace all DNS Records for the specified Domain with the specified Type and Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose DNS Records are to be replaced | 
  **type_** | **string**| DNS Record Type for which DNS Records are to be replaced | 
  **name** | **string**| DNS Record Name for which DNS Records are to be replaced | 
  **records** | [**[]DnsRecordCreateTypeName**](DNSRecordCreateTypeName.md)| DNS Records to replace whatever currently exists | 
 **optional** | ***RecordReplaceTypeNameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordReplaceTypeNameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xShopperId** | **optional.String**| Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Renew**
> DomainPurchaseResponse Renew(ctx, domain, optional)
Renew the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain to renew | 
 **optional** | ***RenewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RenewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **optional.String**| Shopper for whom Domain is to be renewed. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 
 **body** | [**optional.Interface of DomainRenew**](DomainRenew.md)| Options for renewing existing Domain | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Schema**
> JsonSchema Schema(ctx, tld)
Retrieve the schema to be submitted when registering a Domain for the specified TLD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tld** | **string**| The Top-Level Domain whose schema should be retrieved | 

### Return type

[**JsonSchema**](JsonSchema.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Suggest**
> []DomainSuggestion Suggest(ctx, optional)
Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper's purchase history

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SuggestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SuggestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **optional.String**| Shopper ID for which the suggestions are being generated | 
 **query** | **optional.String**| Domain name or set of keywords for which alternative domain names will be suggested | 
 **country** | **optional.String**| Two-letter ISO country code to be used as a hint for target region&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.iso.org/iso/country_codes.htm\&quot;&gt;more&lt;/a&gt; | 
 **city** | **optional.String**| Name of city to be used as a hint for target region | 
 **sources** | [**optional.Interface of []string**](string.md)| Sources to be queried&lt;br/&gt;&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong&gt;CC_TLD&lt;/strong&gt; - Varies the TLD using Country Codes&lt;/li&gt; &lt;li&gt;&lt;strong&gt;EXTENSION&lt;/strong&gt; - Varies the TLD&lt;/li&gt; &lt;li&gt;&lt;strong&gt;KEYWORD_SPIN&lt;/strong&gt; - Identifies keywords and then rotates each one&lt;/li&gt; &lt;li&gt;&lt;strong&gt;PREMIUM&lt;/strong&gt; - Includes variations with premium prices&lt;/li&gt;&lt;/ul&gt; | 
 **tlds** | [**optional.Interface of []string**](string.md)| Top-level domains to be included in suggestions&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.godaddy.com/tlds/gtld.aspx#domain_search_form\&quot;&gt;more&lt;/a&gt; | 
 **lengthMax** | **optional.Int32**| Maximum length of second-level domain | 
 **lengthMin** | **optional.Int32**| Minimum length of second-level domain | 
 **limit** | **optional.Int32**| Maximum number of suggestions to return | 
 **waitMs** | **optional.Int32**| Maximum amount of time, in milliseconds, to wait for responses If elapses, return the results compiled up to that point | [default to 1000]

### Return type

[**[]DomainSuggestion**](DomainSuggestion.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Tlds**
> []TldSummary Tlds(ctx, )
Retrieves a list of TLDs supported and enabled for sale

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TldSummary**](TldSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TransferIn**
> DomainPurchaseResponse TransferIn(ctx, domain, body, optional)
Purchase and start or restart transfer process

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain to transfer in | 
  **body** | [**DomainTransferIn**](DomainTransferIn.md)| Details for domain transfer purchase | 
 **optional** | ***TransferInOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TransferInOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| The Shopper to whom the domain should be transfered | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, domain, body, optional)
Update details for the specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose details are to be updated | 
  **body** | [**DomainUpdate**](DomainUpdate.md)| Changes to apply to existing Domain | 
 **optional** | ***UpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| Shopper for whom Domain is to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContacts**
> UpdateContacts(ctx, domain, contacts, optional)
Update domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose Contacts are to be updated. | 
  **contacts** | [**DomainContacts**](DomainContacts.md)| Changes to apply to existing Contacts | 
 **optional** | ***UpdateContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xShopperId** | **optional.String**| Shopper for whom domain contacts are to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Validate**
> Validate(ctx, body)
Validate the request body using the Domain Purchase Schema for the specified TLD

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainPurchase**](DomainPurchase.md)| An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyEmail**
> VerifyEmail(ctx, domain, optional)
Re-send Contact E-mail Verification for specified Domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **domain** | **string**| Domain whose Contact E-mail should be verified. | 
 **optional** | ***VerifyEmailOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VerifyEmailOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **optional.String**| Shopper for whom domain contact e-mail should be verified. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/xml, text/xml
 - **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

