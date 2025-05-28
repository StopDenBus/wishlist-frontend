# Wish

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Product** | **string** |  | 
**Price** | **float32** |  | 
**Url** | **string** |  | 
**Priority** | [**PriorityEnum**](PriorityEnum.md) |  | 

## Methods

### NewWish

`func NewWish(id int32, product string, price float32, url string, priority PriorityEnum, ) *Wish`

NewWish instantiates a new Wish object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWishWithDefaults

`func NewWishWithDefaults() *Wish`

NewWishWithDefaults instantiates a new Wish object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Wish) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Wish) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Wish) SetId(v int32)`

SetId sets Id field to given value.


### GetProduct

`func (o *Wish) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Wish) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Wish) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetPrice

`func (o *Wish) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Wish) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Wish) SetPrice(v float32)`

SetPrice sets Price field to given value.


### GetUrl

`func (o *Wish) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Wish) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Wish) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPriority

`func (o *Wish) GetPriority() PriorityEnum`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Wish) GetPriorityOk() (*PriorityEnum, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Wish) SetPriority(v PriorityEnum)`

SetPriority sets Priority field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


