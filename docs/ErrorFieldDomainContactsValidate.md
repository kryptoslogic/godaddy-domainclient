# ErrorFieldDomainContactsValidate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | [default to null]
**Domains** | **[]string** | An array of domain names the error is for. If tlds are specified in the request, &#x60;domains&#x60; will contain tlds. For example, if &#x60;domains&#x60; in request is [\&quot;test1.com\&quot;, \&quot;test2.uk\&quot;, \&quot;net\&quot;], and the field is invalid for com and net, then one of the &#x60;fields&#x60; in response will have [\&quot;test1.com\&quot;, \&quot;net\&quot;] as &#x60;domains&#x60; | [default to null]
**Message** | **string** | Human-readable, English description of the problem with the contents of the field | [optional] [default to null]
**Path** | **string** | 1) JSONPath referring to the field within the data containing an error&lt;br/&gt;or&lt;br/&gt;2) JSONPath referring to an object containing an error | [default to null]
**PathRelated** | **string** | JSONPath referring to the field on the object referenced by &#x60;path&#x60; containing an error | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


