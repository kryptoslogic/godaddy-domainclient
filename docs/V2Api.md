# \V2Api

All URIs are relative to *https://api.ote-godaddy.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DomainsForwardsDelete**](V2Api.md#DomainsForwardsDelete) | **Delete** /v2/customers/{customerId}/domains/forwards/{fqdn} | Submit a forwarding cancellation request for the given fqdn
[**DomainsForwardsGet**](V2Api.md#DomainsForwardsGet) | **Get** /v2/customers/{customerId}/domains/forwards/{fqdn} | Retrieve the forwarding information for the given fqdn
[**DomainsForwardsPut**](V2Api.md#DomainsForwardsPut) | **Put** /v2/customers/{customerId}/domains/forwards/{fqdn} | Modify the forwarding information for the given fqdn



## DomainsForwardsDelete

> DomainsForwardsDelete(ctx, customerId, fqdn).Execute()

Submit a forwarding cancellation request for the given fqdn



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
    customerId := "customerId_example" // string | The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.
    fqdn := "fqdn_example" // string | The fully qualified domain name whose forwarding details are to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DomainsForwardsDelete(context.Background(), customerId, fqdn).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DomainsForwardsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
**fqdn** | **string** | The fully qualified domain name whose forwarding details are to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsForwardsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsForwardsGet

> []DomainForwarding DomainsForwardsGet(ctx, customerId, fqdn).IncludeSubs(includeSubs).Execute()

Retrieve the forwarding information for the given fqdn



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
    customerId := "customerId_example" // string | The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.
    fqdn := "fqdn_example" // string | The fully qualified domain name whose forwarding details are to be retrieved.
    includeSubs := true // bool | Optionally include all sub domains if the fqdn specified is a domain and not a sub domain. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DomainsForwardsGet(context.Background(), customerId, fqdn).IncludeSubs(includeSubs).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DomainsForwardsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DomainsForwardsGet`: []DomainForwarding
    fmt.Fprintf(os.Stdout, "Response from `V2Api.DomainsForwardsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
**fqdn** | **string** | The fully qualified domain name whose forwarding details are to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsForwardsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeSubs** | **bool** | Optionally include all sub domains if the fqdn specified is a domain and not a sub domain. | 

### Return type

[**[]DomainForwarding**](DomainForwarding.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DomainsForwardsPut

> DomainsForwardsPut(ctx, customerId, fqdn).Body(body).Execute()

Modify the forwarding information for the given fqdn



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
    customerId := "customerId_example" // string | The Customer identifier<br/> Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you're operating on behalf of; otherwise use your shopper id.
    fqdn := "fqdn_example" // string | The fully qualified domain name whose forwarding details are to be modified.
    body := *openapiclient.NewDomainForwardingCreate("Type_example", "Url_example") // DomainForwardingCreate | Domain forwarding rule to create or replace on the fqdn

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.V2Api.DomainsForwardsPut(context.Background(), customerId, fqdn).Body(body).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `V2Api.DomainsForwardsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customerId** | **string** | The Customer identifier&lt;br/&gt; Note: For API Resellers, performing actions on behalf of your customers, you need to specify the Subaccount you&#39;re operating on behalf of; otherwise use your shopper id. | 
**fqdn** | **string** | The fully qualified domain name whose forwarding details are to be modified. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDomainsForwardsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DomainForwardingCreate**](DomainForwardingCreate.md) | Domain forwarding rule to create or replace on the fqdn | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

