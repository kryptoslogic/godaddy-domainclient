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



## Available

> DomainAvailableResponse Available(ctx).Domain(domain).CheckType(checkType).ForTransfer(forTransfer).Execute()

Determine whether or not the specified domain is available for purchase

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain name whose availability is to be checked
    checkType := "checkType_example" // string | Optimize for time ('FAST') or accuracy ('FULL') (optional) (default to "FAST")
    forTransfer := true // bool | Whether or not to include domains available for transfer. If set to True, checkType is ignored (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Available(context.Background()).Domain(domain).CheckType(checkType).ForTransfer(forTransfer).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Available``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Available`: DomainAvailableResponse
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Available`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domain** | **string** | Domain name whose availability is to be checked | 
 **checkType** | **string** | Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [default to &quot;FAST&quot;]
 **forTransfer** | **bool** | Whether or not to include domains available for transfer. If set to True, checkType is ignored | [default to false]

### Return type

[**DomainAvailableResponse**](DomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AvailableBulk

> DomainAvailableBulk AvailableBulk(ctx).Domains(domains).CheckType(checkType).Execute()

Determine whether or not the specified domains are available for purchase

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domains := []string{"Property_example"} // []string | Domain names for which to check availability
    checkType := "checkType_example" // string | Optimize for time ('FAST') or accuracy ('FULL') (optional) (default to "FAST")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.AvailableBulk(context.Background()).Domains(domains).CheckType(checkType).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.AvailableBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AvailableBulk`: DomainAvailableBulk
    fmt.Fprintf(os.Stdout, "Response from `V1Api.AvailableBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAvailableBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domains** | **[]string** | Domain names for which to check availability | 
 **checkType** | **string** | Optimize for time (&#39;FAST&#39;) or accuracy (&#39;FULL&#39;) | [default to &quot;FAST&quot;]

### Return type

[**DomainAvailableBulk**](DomainAvailableBulk.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Cancel

> Cancel(ctx, domain).Execute()

Cancel a purchased domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain to cancel

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Cancel(context.Background(), domain).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Cancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelPrivacy

> CancelPrivacy(ctx, domain).XShopperId(xShopperId).Execute()

Submit a privacy cancellation request for the given domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose privacy is to be cancelled
    xShopperId := "xShopperId_example" // string | Shopper ID of the owner of the domain (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.CancelPrivacy(context.Background(), domain).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.CancelPrivacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose privacy is to be cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPrivacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **string** | Shopper ID of the owner of the domain | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ContactsValidate

> ContactsValidate(ctx).Body(body).XPrivateLabelId(xPrivateLabelId).MarketId(marketId).Execute()

Validate the request body using the Domain Contact Validation Schema for specified domains.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDomainsContactsBulk([]string{"Domains_example"}) // DomainsContactsBulk | An instance document expected for domains contacts validation
    xPrivateLabelId := int32(56) // int32 | PrivateLabelId to operate as, if different from JWT (optional) (default to 1)
    marketId := "marketId_example" // string | MarketId in which the request is being made, and for which responses should be localized (optional) (default to "en-US")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.ContactsValidate(context.Background()).Body(body).XPrivateLabelId(xPrivateLabelId).MarketId(marketId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.ContactsValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactsValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainsContactsBulk**](DomainsContactsBulk.md) | An instance document expected for domains contacts validation | 
 **xPrivateLabelId** | **int32** | PrivateLabelId to operate as, if different from JWT | [default to 1]
 **marketId** | **string** | MarketId in which the request is being made, and for which responses should be localized | [default to &quot;en-US&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> DomainDetail Get(ctx, domain).XShopperId(xShopperId).Execute()

Retrieve details for the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain name whose details are to be retrieved
    xShopperId := "xShopperId_example" // string | Shopper ID expected to own the specified domain (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Get(context.Background(), domain).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: DomainDetail
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name whose details are to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **string** | Shopper ID expected to own the specified domain | 

### Return type

[**DomainDetail**](DomainDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgreement

> []LegalAgreement GetAgreement(ctx).Tlds(tlds).Privacy(privacy).XMarketId(xMarketId).ForTransfer(forTransfer).Execute()

Retrieve the legal agreement(s) required to purchase the specified TLD and add-ons

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tlds := []string{"Inner_example"} // []string | list of TLDs whose legal agreements are to be retrieved
    privacy := true // bool | Whether or not privacy has been requested
    xMarketId := "xMarketId_example" // string | Unique identifier of the Market used to retrieve/translate Legal Agreements (optional) (default to "en-US")
    forTransfer := true // bool | Whether or not domain tranfer has been requested (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.GetAgreement(context.Background()).Tlds(tlds).Privacy(privacy).XMarketId(xMarketId).ForTransfer(forTransfer).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.GetAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgreement`: []LegalAgreement
    fmt.Fprintf(os.Stdout, "Response from `V1Api.GetAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tlds** | **[]string** | list of TLDs whose legal agreements are to be retrieved | 
 **privacy** | **bool** | Whether or not privacy has been requested | 
 **xMarketId** | **string** | Unique identifier of the Market used to retrieve/translate Legal Agreements | [default to &quot;en-US&quot;]
 **forTransfer** | **bool** | Whether or not domain tranfer has been requested | 

### Return type

[**[]LegalAgreement**](LegalAgreement.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> []DomainSummary List(ctx).XShopperId(xShopperId).Statuses(statuses).StatusGroups(statusGroups).Limit(limit).Marker(marker).Includes(includes).ModifiedDate(modifiedDate).Execute()

Retrieve a list of Domains for the specified Shopper

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xShopperId := "xShopperId_example" // string | Shopper ID whose domains are to be retrieved (optional)
    statuses := []string{"Statuses_example"} // []string | Only include results with `status` value in the specified set (optional)
    statusGroups := []string{"StatusGroups_example"} // []string | Only include results with `status` value in any of the specified groups (optional)
    limit := int32(56) // int32 | Maximum number of domains to return (optional)
    marker := "marker_example" // string | Marker Domain to use as the offset in results (optional)
    includes := []string{"Includes_example"} // []string | Optional details to be included in the response (optional)
    modifiedDate := "modifiedDate_example" // string | Only include results that have been modified since the specified date (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.List(context.Background()).XShopperId(xShopperId).Statuses(statuses).StatusGroups(statusGroups).Limit(limit).Marker(marker).Includes(includes).ModifiedDate(modifiedDate).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: []DomainSummary
    fmt.Fprintf(os.Stdout, "Response from `V1Api.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **string** | Shopper ID whose domains are to be retrieved | 
 **statuses** | **[]string** | Only include results with &#x60;status&#x60; value in the specified set | 
 **statusGroups** | **[]string** | Only include results with &#x60;status&#x60; value in any of the specified groups | 
 **limit** | **int32** | Maximum number of domains to return | 
 **marker** | **string** | Marker Domain to use as the offset in results | 
 **includes** | **[]string** | Optional details to be included in the response | 
 **modifiedDate** | **string** | Only include results that have been modified since the specified date | 

### Return type

[**[]DomainSummary**](DomainSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Purchase

> DomainPurchaseResponse Purchase(ctx).Body(body).XShopperId(xShopperId).Execute()

Purchase and register the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDomainPurchase(*openapiclient.NewConsent("AgreedAt_example", "AgreedBy_example", []string{"AgreementKeys_example"}), "Domain_example") // DomainPurchase | An instance document expected to match the JSON schema returned by `./schema/{tld}`
    xShopperId := "xShopperId_example" // string | The Shopper for whom the domain should be purchased (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Purchase(context.Background()).Body(body).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Purchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Purchase`: DomainPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Purchase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainPurchase**](DomainPurchase.md) | An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 
 **xShopperId** | **string** | The Shopper for whom the domain should be purchased | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurchasePrivacy

> DomainPurchaseResponse PurchasePrivacy(ctx, domain).Body(body).XShopperId(xShopperId).Execute()

Purchase privacy for a specified domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain for which to purchase privacy
    body := *openapiclient.NewPrivacyPurchase(*openapiclient.NewConsent("AgreedAt_example", "AgreedBy_example", []string{"AgreementKeys_example"})) // PrivacyPurchase | Options for purchasing privacy
    xShopperId := "xShopperId_example" // string | Shopper ID of the owner of the domain (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.PurchasePrivacy(context.Background(), domain).Body(body).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.PurchasePrivacy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PurchasePrivacy`: DomainPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `V1Api.PurchasePrivacy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain for which to purchase privacy | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurchasePrivacyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PrivacyPurchase**](PrivacyPurchase.md) | Options for purchasing privacy | 
 **xShopperId** | **string** | Shopper ID of the owner of the domain | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordAdd

> RecordAdd(ctx, domain).Records(records).XShopperId(xShopperId).Execute()

Add the specified DNS Records to the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose DNS Records are to be augmented
    records := []openapiclient.DNSRecord{*openapiclient.NewDNSRecord("Data_example", "Name_example", "Type_example")} // []DNSRecord | DNS Records to add to whatever currently exists
    xShopperId := "xShopperId_example" // string | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.RecordAdd(context.Background(), domain).Records(records).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.RecordAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose DNS Records are to be augmented | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **records** | [**[]DNSRecord**](DNSRecord.md) | DNS Records to add to whatever currently exists | 
 **xShopperId** | **string** | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordGet

> []DNSRecord RecordGet(ctx, domain, type_, name).XShopperId(xShopperId).Offset(offset).Limit(limit).Execute()

Retrieve DNS Records for the specified Domain, optionally with the specified Type and/or Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose DNS Records are to be retrieved
    type_ := "type__example" // string | DNS Record Type for which DNS Records are to be retrieved
    name := "name_example" // string | DNS Record Name for which DNS Records are to be retrieved
    xShopperId := "xShopperId_example" // string | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)
    offset := int32(56) // int32 | Number of results to skip for pagination (optional)
    limit := int32(56) // int32 | Maximum number of items to return (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.RecordGet(context.Background(), domain, type_, name).XShopperId(xShopperId).Offset(offset).Limit(limit).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.RecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecordGet`: []DNSRecord
    fmt.Fprintf(os.Stdout, "Response from `V1Api.RecordGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose DNS Records are to be retrieved | 
**type_** | **string** | DNS Record Type for which DNS Records are to be retrieved | 
**name** | **string** | DNS Record Name for which DNS Records are to be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xShopperId** | **string** | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 
 **offset** | **int32** | Number of results to skip for pagination | 
 **limit** | **int32** | Maximum number of items to return | 

### Return type

[**[]DNSRecord**](DNSRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordReplace

> RecordReplace(ctx, domain).Records(records).XShopperId(xShopperId).Execute()

Replace all DNS Records for the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose DNS Records are to be replaced
    records := []openapiclient.DNSRecord{*openapiclient.NewDNSRecord("Data_example", "Name_example", "Type_example")} // []DNSRecord | DNS Records to replace whatever currently exists
    xShopperId := "xShopperId_example" // string | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.RecordReplace(context.Background(), domain).Records(records).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.RecordReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose DNS Records are to be replaced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **records** | [**[]DNSRecord**](DNSRecord.md) | DNS Records to replace whatever currently exists | 
 **xShopperId** | **string** | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordReplaceType

> RecordReplaceType(ctx, domain, type_).Records(records).XShopperId(xShopperId).Execute()

Replace all DNS Records for the specified Domain with the specified Type

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose DNS Records are to be replaced
    type_ := "type__example" // string | DNS Record Type for which DNS Records are to be replaced
    records := []openapiclient.DNSRecordCreateType{*openapiclient.NewDNSRecordCreateType("Data_example", "Name_example")} // []DNSRecordCreateType | DNS Records to replace whatever currently exists
    xShopperId := "xShopperId_example" // string | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.RecordReplaceType(context.Background(), domain, type_).Records(records).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.RecordReplaceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose DNS Records are to be replaced | 
**type_** | **string** | DNS Record Type for which DNS Records are to be replaced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordReplaceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **records** | [**[]DNSRecordCreateType**](DNSRecordCreateType.md) | DNS Records to replace whatever currently exists | 
 **xShopperId** | **string** | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordReplaceTypeName

> RecordReplaceTypeName(ctx, domain, type_, name).Records(records).XShopperId(xShopperId).Execute()

Replace all DNS Records for the specified Domain with the specified Type and Name

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose DNS Records are to be replaced
    type_ := "type__example" // string | DNS Record Type for which DNS Records are to be replaced
    name := "name_example" // string | DNS Record Name for which DNS Records are to be replaced
    records := []openapiclient.DNSRecordCreateTypeName{*openapiclient.NewDNSRecordCreateTypeName("Data_example")} // []DNSRecordCreateTypeName | DNS Records to replace whatever currently exists
    xShopperId := "xShopperId_example" // string | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.RecordReplaceTypeName(context.Background(), domain, type_, name).Records(records).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.RecordReplaceTypeName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose DNS Records are to be replaced | 
**type_** | **string** | DNS Record Type for which DNS Records are to be replaced | 
**name** | **string** | DNS Record Name for which DNS Records are to be replaced | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecordReplaceTypeNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **records** | [**[]DNSRecordCreateTypeName**](DNSRecordCreateTypeName.md) | DNS Records to replace whatever currently exists | 
 **xShopperId** | **string** | Shopper ID which owns the domain. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Renew

> DomainPurchaseResponse Renew(ctx, domain).XShopperId(xShopperId).Body(body).Execute()

Renew the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain to renew
    xShopperId := "xShopperId_example" // string | Shopper for whom Domain is to be renewed. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)
    body := *openapiclient.NewDomainRenew() // DomainRenew | Options for renewing existing Domain (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Renew(context.Background(), domain).XShopperId(xShopperId).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Renew``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Renew`: DomainPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Renew`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain to renew | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **string** | Shopper for whom Domain is to be renewed. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 
 **body** | [**DomainRenew**](DomainRenew.md) | Options for renewing existing Domain | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Schema

> JsonSchema Schema(ctx, tld).Execute()

Retrieve the schema to be submitted when registering a Domain for the specified TLD

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tld := "tld_example" // string | The Top-Level Domain whose schema should be retrieved

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Schema(context.Background(), tld).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Schema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Schema`: JsonSchema
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Schema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | The Top-Level Domain whose schema should be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JsonSchema**](JsonSchema.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Suggest

> []DomainSuggestion Suggest(ctx).XShopperId(xShopperId).Query(query).Country(country).City(city).Sources(sources).Tlds(tlds).LengthMax(lengthMax).LengthMin(lengthMin).Limit(limit).WaitMs(waitMs).Execute()

Suggest alternate Domain names based on a seed Domain, a set of keywords, or the shopper's purchase history

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xShopperId := "xShopperId_example" // string | Shopper ID for which the suggestions are being generated (optional)
    query := "query_example" // string | Domain name or set of keywords for which alternative domain names will be suggested (optional)
    country := "country_example" // string | Two-letter ISO country code to be used as a hint for target region<br/><br/> NOTE: These are sample values, there are many <a href=\"http://www.iso.org/iso/country_codes.htm\">more</a> (optional)
    city := "city_example" // string | Name of city to be used as a hint for target region (optional)
    sources := []string{"Sources_example"} // []string | Sources to be queried<br/><br/><ul> <li><strong>CC_TLD</strong> - Varies the TLD using Country Codes</li> <li><strong>EXTENSION</strong> - Varies the TLD</li> <li><strong>KEYWORD_SPIN</strong> - Identifies keywords and then rotates each one</li> <li><strong>PREMIUM</strong> - Includes variations with premium prices</li></ul> (optional)
    tlds := []string{"Inner_example"} // []string | Top-level domains to be included in suggestions<br/><br/> NOTE: These are sample values, there are many <a href=\"http://www.godaddy.com/tlds/gtld.aspx#domain_search_form\">more</a> (optional)
    lengthMax := int32(56) // int32 | Maximum length of second-level domain (optional)
    lengthMin := int32(56) // int32 | Minimum length of second-level domain (optional)
    limit := int32(56) // int32 | Maximum number of suggestions to return (optional)
    waitMs := int32(56) // int32 | Maximum amount of time, in milliseconds, to wait for responses If elapses, return the results compiled up to that point (optional) (default to 1000)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Suggest(context.Background()).XShopperId(xShopperId).Query(query).Country(country).City(city).Sources(sources).Tlds(tlds).LengthMax(lengthMax).LengthMin(lengthMin).Limit(limit).WaitMs(waitMs).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Suggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Suggest`: []DomainSuggestion
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Suggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xShopperId** | **string** | Shopper ID for which the suggestions are being generated | 
 **query** | **string** | Domain name or set of keywords for which alternative domain names will be suggested | 
 **country** | **string** | Two-letter ISO country code to be used as a hint for target region&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.iso.org/iso/country_codes.htm\&quot;&gt;more&lt;/a&gt; | 
 **city** | **string** | Name of city to be used as a hint for target region | 
 **sources** | **[]string** | Sources to be queried&lt;br/&gt;&lt;br/&gt;&lt;ul&gt; &lt;li&gt;&lt;strong&gt;CC_TLD&lt;/strong&gt; - Varies the TLD using Country Codes&lt;/li&gt; &lt;li&gt;&lt;strong&gt;EXTENSION&lt;/strong&gt; - Varies the TLD&lt;/li&gt; &lt;li&gt;&lt;strong&gt;KEYWORD_SPIN&lt;/strong&gt; - Identifies keywords and then rotates each one&lt;/li&gt; &lt;li&gt;&lt;strong&gt;PREMIUM&lt;/strong&gt; - Includes variations with premium prices&lt;/li&gt;&lt;/ul&gt; | 
 **tlds** | **[]string** | Top-level domains to be included in suggestions&lt;br/&gt;&lt;br/&gt; NOTE: These are sample values, there are many &lt;a href&#x3D;\&quot;http://www.godaddy.com/tlds/gtld.aspx#domain_search_form\&quot;&gt;more&lt;/a&gt; | 
 **lengthMax** | **int32** | Maximum length of second-level domain | 
 **lengthMin** | **int32** | Minimum length of second-level domain | 
 **limit** | **int32** | Maximum number of suggestions to return | 
 **waitMs** | **int32** | Maximum amount of time, in milliseconds, to wait for responses If elapses, return the results compiled up to that point | [default to 1000]

### Return type

[**[]DomainSuggestion**](DomainSuggestion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Tlds

> []TldSummary Tlds(ctx).Execute()

Retrieves a list of TLDs supported and enabled for sale

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Tlds(context.Background()).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Tlds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Tlds`: []TldSummary
    fmt.Fprintf(os.Stdout, "Response from `V1Api.Tlds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTldsRequest struct via the builder pattern


### Return type

[**[]TldSummary**](TldSummary.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferIn

> DomainPurchaseResponse TransferIn(ctx, domain).Body(body).XShopperId(xShopperId).Execute()

Purchase and start or restart transfer process

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain to transfer in
    body := *openapiclient.NewDomainTransferIn("AuthCode_example", *openapiclient.NewConsent("AgreedAt_example", "AgreedBy_example", []string{"AgreementKeys_example"})) // DomainTransferIn | Details for domain transfer purchase
    xShopperId := "xShopperId_example" // string | The Shopper to whom the domain should be transfered (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.TransferIn(context.Background(), domain).Body(body).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.TransferIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferIn`: DomainPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `V1Api.TransferIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain to transfer in | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransferInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DomainTransferIn**](DomainTransferIn.md) | Details for domain transfer purchase | 
 **xShopperId** | **string** | The Shopper to whom the domain should be transfered | 

### Return type

[**DomainPurchaseResponse**](DomainPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Update(ctx, domain).Body(body).XShopperId(xShopperId).Execute()

Update details for the specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose details are to be updated
    body := *openapiclient.NewDomainUpdate() // DomainUpdate | Changes to apply to existing Domain
    xShopperId := "xShopperId_example" // string | Shopper for whom Domain is to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Update(context.Background(), domain).Body(body).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose details are to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DomainUpdate**](DomainUpdate.md) | Changes to apply to existing Domain | 
 **xShopperId** | **string** | Shopper for whom Domain is to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContacts

> UpdateContacts(ctx, domain).Contacts(contacts).XShopperId(xShopperId).Execute()

Update domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose Contacts are to be updated.
    contacts := *openapiclient.NewDomainContacts(*openapiclient.NewContact(*openapiclient.NewAddress("Address1_example", "City_example", "Country_example", "PostalCode_example", "State_example"), "Email_example", "NameFirst_example", "NameLast_example", "Phone_example")) // DomainContacts | Changes to apply to existing Contacts
    xShopperId := "xShopperId_example" // string | Shopper for whom domain contacts are to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.UpdateContacts(context.Background(), domain).Contacts(contacts).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.UpdateContacts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose Contacts are to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contacts** | [**DomainContacts**](DomainContacts.md) | Changes to apply to existing Contacts | 
 **xShopperId** | **string** | Shopper for whom domain contacts are to be updated. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Validate

> Validate(ctx).Body(body).Execute()

Validate the request body using the Domain Purchase Schema for the specified TLD

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewDomainPurchase(*openapiclient.NewConsent("AgreedAt_example", "AgreedBy_example", []string{"AgreementKeys_example"}), "Domain_example") // DomainPurchase | An instance document expected to match the JSON schema returned by `./schema/{tld}`

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.Validate(context.Background()).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DomainPurchase**](DomainPurchase.md) | An instance document expected to match the JSON schema returned by &#x60;./schema/{tld}&#x60; | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml, text/xml
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmail

> VerifyEmail(ctx, domain).XShopperId(xShopperId).Execute()

Re-send Contact E-mail Verification for specified Domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    domain := "domain_example" // string | Domain whose Contact E-mail should be verified.
    xShopperId := "xShopperId_example" // string | Shopper for whom domain contact e-mail should be verified. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you're a Reseller, but purchased a Domain via http://www.godaddy.com (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V1Api.VerifyEmail(context.Background(), domain).XShopperId(xShopperId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V1Api.VerifyEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain whose Contact E-mail should be verified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xShopperId** | **string** | Shopper for whom domain contact e-mail should be verified. NOTE: This is only required if you are a Reseller managing a domain purchased outside the scope of your reseller account. For instance, if you&#39;re a Reseller, but purchased a Domain via http://www.godaddy.com | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/javascript, application/xml, text/javascript, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

