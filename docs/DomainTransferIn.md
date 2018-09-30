# DomainTransferIn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthCode** | **string** | Authorization code from registrar for transferring a domain | [default to null]
**Consent** | [***Consent**](Consent.md) | Required agreements can be retrieved via the GET ./domains/agreements endpoint | [default to null]
**Period** | **int32** | Can be more than 1 but no more than 10 years total including current registration length | [optional] [default to null]
**Privacy** | **bool** | Whether or not privacy has been requested | [optional] [default to null]
**RenewAuto** | **bool** | Whether or not the domain should be configured to automatically renew | [optional] [default to null]
**ContactAdmin** | [***Contact**](Contact.md) | The contact to use for the domain admin contact. Depending on the tld of the domain being transferred, this field may be required. | [optional] [default to null]
**ContactBilling** | [***Contact**](Contact.md) | The contact to use for the domain billing contact. Depending on the tld of the domain being transferred, this field may be required. | [optional] [default to null]
**ContactRegistrant** | [***Contact**](Contact.md) | The contact to use for the domain registrant contact. Depending on the tld of the domain being transferred, this field may be required. | [optional] [default to null]
**ContactTech** | [***Contact**](Contact.md) | The contact to use for the domain tech contact. Depending on the tld of the domain being transferred, this field may be required. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


