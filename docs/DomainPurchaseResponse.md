# DomainPurchaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** | Currency in which the &#x60;total&#x60; is listed | [optional] [default to "USD"]
**ItemCount** | **int32** | Number items included in the order | 
**OrderId** | **int32** | Unique identifier of the order processed to purchase the domain | 
**Total** | **int32** | Total cost of the domain and any selected add-ons | 

## Methods

### NewDomainPurchaseResponse

`func NewDomainPurchaseResponse(itemCount int32, orderId int32, total int32, ) *DomainPurchaseResponse`

NewDomainPurchaseResponse instantiates a new DomainPurchaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainPurchaseResponseWithDefaults

`func NewDomainPurchaseResponseWithDefaults() *DomainPurchaseResponse`

NewDomainPurchaseResponseWithDefaults instantiates a new DomainPurchaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *DomainPurchaseResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DomainPurchaseResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DomainPurchaseResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *DomainPurchaseResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetItemCount

`func (o *DomainPurchaseResponse) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *DomainPurchaseResponse) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *DomainPurchaseResponse) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.


### GetOrderId

`func (o *DomainPurchaseResponse) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *DomainPurchaseResponse) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *DomainPurchaseResponse) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetTotal

`func (o *DomainPurchaseResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DomainPurchaseResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DomainPurchaseResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


