# DomainAvailableError

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | [default to null]
**Domain** | **string** | Domain name | [default to null]
**Message** | **string** | Human-readable, English description of the error | [optional] [default to null]
**Path** | **string** | &lt;ul&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field containing an error&lt;/li&gt; &lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;OR&lt;/strong&gt; &lt;li style&#x3D;&#39;margin-left: 12px;&#39;&gt;JSONPath referring to a field that refers to an object containing an error, with more detail in &#x60;pathRelated&#x60;&lt;/li&gt; &lt;/ul&gt; | [default to null]
**Status** | **int32** | HTTP status code that would return for a single check | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


