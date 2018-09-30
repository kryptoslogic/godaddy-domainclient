# ErrorLimit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Short identifier for the error, suitable for indicating the specific error within client code | [default to null]
**Fields** | [**[]ErrorField**](ErrorField.md) | List of the specific fields, and the errors found with their contents | [optional] [default to null]
**Message** | **string** | Human-readable, English description of the error | [optional] [default to null]
**RetryAfterSec** | **int32** | Number of seconds to wait before attempting a similar request | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


