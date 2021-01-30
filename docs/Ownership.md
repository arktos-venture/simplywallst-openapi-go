# Ownership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RankSharesHeld** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**OwnershipOwner**](ownership_owner.md) |  | [optional] 
**HoldingDate** | Pointer to **string** |  | [optional] 
**SharesHeld** | Pointer to **int64** |  | [optional] 
**PercentOfPortfolio** | Pointer to **float32** |  | [optional] 
**PercentOfSharesOutstanding** | Pointer to **float32** |  | [optional] 
**SharesChanged** | Pointer to **float32** |  | [optional] 
**PercentSharesChanged** | Pointer to **float32** |  | [optional] 
**RankSharesBought** | Pointer to **int64** |  | [optional] 
**RankSharesSold** | Pointer to **float32** |  | [optional] 
**PeriodStartDate** | Pointer to **int64** |  | [optional] 
**PeriodEndDate** | Pointer to **int64** |  | [optional] 

## Methods

### NewOwnership

`func NewOwnership() *Ownership`

NewOwnership instantiates a new Ownership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOwnershipWithDefaults

`func NewOwnershipWithDefaults() *Ownership`

NewOwnershipWithDefaults instantiates a new Ownership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRankSharesHeld

`func (o *Ownership) GetRankSharesHeld() int64`

GetRankSharesHeld returns the RankSharesHeld field if non-nil, zero value otherwise.

### GetRankSharesHeldOk

`func (o *Ownership) GetRankSharesHeldOk() (*int64, bool)`

GetRankSharesHeldOk returns a tuple with the RankSharesHeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankSharesHeld

`func (o *Ownership) SetRankSharesHeld(v int64)`

SetRankSharesHeld sets RankSharesHeld field to given value.

### HasRankSharesHeld

`func (o *Ownership) HasRankSharesHeld() bool`

HasRankSharesHeld returns a boolean if a field has been set.

### GetOwner

`func (o *Ownership) GetOwner() OwnershipOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Ownership) GetOwnerOk() (*OwnershipOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Ownership) SetOwner(v OwnershipOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Ownership) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetHoldingDate

`func (o *Ownership) GetHoldingDate() string`

GetHoldingDate returns the HoldingDate field if non-nil, zero value otherwise.

### GetHoldingDateOk

`func (o *Ownership) GetHoldingDateOk() (*string, bool)`

GetHoldingDateOk returns a tuple with the HoldingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldingDate

`func (o *Ownership) SetHoldingDate(v string)`

SetHoldingDate sets HoldingDate field to given value.

### HasHoldingDate

`func (o *Ownership) HasHoldingDate() bool`

HasHoldingDate returns a boolean if a field has been set.

### GetSharesHeld

`func (o *Ownership) GetSharesHeld() int64`

GetSharesHeld returns the SharesHeld field if non-nil, zero value otherwise.

### GetSharesHeldOk

`func (o *Ownership) GetSharesHeldOk() (*int64, bool)`

GetSharesHeldOk returns a tuple with the SharesHeld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesHeld

`func (o *Ownership) SetSharesHeld(v int64)`

SetSharesHeld sets SharesHeld field to given value.

### HasSharesHeld

`func (o *Ownership) HasSharesHeld() bool`

HasSharesHeld returns a boolean if a field has been set.

### GetPercentOfPortfolio

`func (o *Ownership) GetPercentOfPortfolio() float32`

GetPercentOfPortfolio returns the PercentOfPortfolio field if non-nil, zero value otherwise.

### GetPercentOfPortfolioOk

`func (o *Ownership) GetPercentOfPortfolioOk() (*float32, bool)`

GetPercentOfPortfolioOk returns a tuple with the PercentOfPortfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfPortfolio

`func (o *Ownership) SetPercentOfPortfolio(v float32)`

SetPercentOfPortfolio sets PercentOfPortfolio field to given value.

### HasPercentOfPortfolio

`func (o *Ownership) HasPercentOfPortfolio() bool`

HasPercentOfPortfolio returns a boolean if a field has been set.

### GetPercentOfSharesOutstanding

`func (o *Ownership) GetPercentOfSharesOutstanding() float32`

GetPercentOfSharesOutstanding returns the PercentOfSharesOutstanding field if non-nil, zero value otherwise.

### GetPercentOfSharesOutstandingOk

`func (o *Ownership) GetPercentOfSharesOutstandingOk() (*float32, bool)`

GetPercentOfSharesOutstandingOk returns a tuple with the PercentOfSharesOutstanding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentOfSharesOutstanding

`func (o *Ownership) SetPercentOfSharesOutstanding(v float32)`

SetPercentOfSharesOutstanding sets PercentOfSharesOutstanding field to given value.

### HasPercentOfSharesOutstanding

`func (o *Ownership) HasPercentOfSharesOutstanding() bool`

HasPercentOfSharesOutstanding returns a boolean if a field has been set.

### GetSharesChanged

`func (o *Ownership) GetSharesChanged() float32`

GetSharesChanged returns the SharesChanged field if non-nil, zero value otherwise.

### GetSharesChangedOk

`func (o *Ownership) GetSharesChangedOk() (*float32, bool)`

GetSharesChangedOk returns a tuple with the SharesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharesChanged

`func (o *Ownership) SetSharesChanged(v float32)`

SetSharesChanged sets SharesChanged field to given value.

### HasSharesChanged

`func (o *Ownership) HasSharesChanged() bool`

HasSharesChanged returns a boolean if a field has been set.

### GetPercentSharesChanged

`func (o *Ownership) GetPercentSharesChanged() float32`

GetPercentSharesChanged returns the PercentSharesChanged field if non-nil, zero value otherwise.

### GetPercentSharesChangedOk

`func (o *Ownership) GetPercentSharesChangedOk() (*float32, bool)`

GetPercentSharesChangedOk returns a tuple with the PercentSharesChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentSharesChanged

`func (o *Ownership) SetPercentSharesChanged(v float32)`

SetPercentSharesChanged sets PercentSharesChanged field to given value.

### HasPercentSharesChanged

`func (o *Ownership) HasPercentSharesChanged() bool`

HasPercentSharesChanged returns a boolean if a field has been set.

### GetRankSharesBought

`func (o *Ownership) GetRankSharesBought() int64`

GetRankSharesBought returns the RankSharesBought field if non-nil, zero value otherwise.

### GetRankSharesBoughtOk

`func (o *Ownership) GetRankSharesBoughtOk() (*int64, bool)`

GetRankSharesBoughtOk returns a tuple with the RankSharesBought field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankSharesBought

`func (o *Ownership) SetRankSharesBought(v int64)`

SetRankSharesBought sets RankSharesBought field to given value.

### HasRankSharesBought

`func (o *Ownership) HasRankSharesBought() bool`

HasRankSharesBought returns a boolean if a field has been set.

### GetRankSharesSold

`func (o *Ownership) GetRankSharesSold() float32`

GetRankSharesSold returns the RankSharesSold field if non-nil, zero value otherwise.

### GetRankSharesSoldOk

`func (o *Ownership) GetRankSharesSoldOk() (*float32, bool)`

GetRankSharesSoldOk returns a tuple with the RankSharesSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankSharesSold

`func (o *Ownership) SetRankSharesSold(v float32)`

SetRankSharesSold sets RankSharesSold field to given value.

### HasRankSharesSold

`func (o *Ownership) HasRankSharesSold() bool`

HasRankSharesSold returns a boolean if a field has been set.

### GetPeriodStartDate

`func (o *Ownership) GetPeriodStartDate() int64`

GetPeriodStartDate returns the PeriodStartDate field if non-nil, zero value otherwise.

### GetPeriodStartDateOk

`func (o *Ownership) GetPeriodStartDateOk() (*int64, bool)`

GetPeriodStartDateOk returns a tuple with the PeriodStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStartDate

`func (o *Ownership) SetPeriodStartDate(v int64)`

SetPeriodStartDate sets PeriodStartDate field to given value.

### HasPeriodStartDate

`func (o *Ownership) HasPeriodStartDate() bool`

HasPeriodStartDate returns a boolean if a field has been set.

### GetPeriodEndDate

`func (o *Ownership) GetPeriodEndDate() int64`

GetPeriodEndDate returns the PeriodEndDate field if non-nil, zero value otherwise.

### GetPeriodEndDateOk

`func (o *Ownership) GetPeriodEndDateOk() (*int64, bool)`

GetPeriodEndDateOk returns a tuple with the PeriodEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEndDate

`func (o *Ownership) SetPeriodEndDate(v int64)`

SetPeriodEndDate sets PeriodEndDate field to given value.

### HasPeriodEndDate

`func (o *Ownership) HasPeriodEndDate() bool`

HasPeriodEndDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


