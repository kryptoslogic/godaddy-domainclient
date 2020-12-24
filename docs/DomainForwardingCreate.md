# DomainForwardingCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of fowarding to implement&lt;br/&gt;&lt;ul&gt;&lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;MASKED&lt;/strong&gt; - Prevents the forwarded domain or subdomain URL from displaying in the browser&#39;s address bar.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REDIRECT_PERMANENT*&lt;/strong&gt; - Redirects to the url you specified in the forwardTo field using a &#x60;301 Moved Permanently&#x60; HTTP response. The HTTP 301 response code tells user-agents (including search engines) that the location has permanently moved.&lt;/li&gt;&lt;li&gt;&lt;strong style&#x3D;&#39;margin-left: 12px;&#39;&gt;REDIRECT_TEMPORARY&lt;/strong&gt; - Redirects to the url you specified in the forwardTo field using a &#x60;302 Found&#x60; HTTP response. The HTTP 302 response code tells user-agents (including search engines) that the location has temporarily moved.&lt;/li&gt;&lt;/ul&gt; | [default to "REDIRECT_PERMANENT"]
**Url** | **string** | Forwards http(s) traffic to this destination url (ex. http://www.somedomain.com/) | 
**Mask** | Pointer to [**DomainForwardingMask**](DomainForwardingMask.md) |  | [optional] 

## Methods

### NewDomainForwardingCreate

`func NewDomainForwardingCreate(type_ string, url string, ) *DomainForwardingCreate`

NewDomainForwardingCreate instantiates a new DomainForwardingCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainForwardingCreateWithDefaults

`func NewDomainForwardingCreateWithDefaults() *DomainForwardingCreate`

NewDomainForwardingCreateWithDefaults instantiates a new DomainForwardingCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DomainForwardingCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DomainForwardingCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DomainForwardingCreate) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *DomainForwardingCreate) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DomainForwardingCreate) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DomainForwardingCreate) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMask

`func (o *DomainForwardingCreate) GetMask() DomainForwardingMask`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *DomainForwardingCreate) GetMaskOk() (*DomainForwardingMask, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *DomainForwardingCreate) SetMask(v DomainForwardingMask)`

SetMask sets Mask field to given value.

### HasMask

`func (o *DomainForwardingCreate) HasMask() bool`

HasMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


