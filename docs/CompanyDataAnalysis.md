# CompanyDataAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**SharePrice** | Pointer to **float64** |  | [optional] 
**MarketCap** | Pointer to **float64** |  | [optional] 
**IntrinsicDiscount** | Pointer to **float64** |  | [optional] 
**Pe** | Pointer to **float64** |  | [optional] 
**Pb** | Pointer to **float64** |  | [optional] 
**Peg** | Pointer to **float64** |  | [optional] 
**Roe** | Pointer to **float64** |  | [optional] 
**Roa** | Pointer to **float64** |  | [optional] 
**Eps** | Pointer to **float64** |  | [optional] 
**DebtEquity** | Pointer to **float64** |  | [optional] 
**AnalystCount** | Pointer to **float64** |  | [optional] 
**Dividend** | Pointer to [**CompanyAnalysisDividend**](CompanyAnalysisDividend.md) |  | [optional] 
**Future** | Pointer to [**CompanyAnalysisFuture**](CompanyAnalysisFuture.md) |  | [optional] 
**Past** | Pointer to [**CompanyAnalysisPast**](CompanyAnalysisPast.md) |  | [optional] 
**Extended** | Pointer to [**CompanyAnalysisExtendedData**](CompanyAnalysisExtendedData.md) |  | [optional] 

## Methods

### NewCompanyDataAnalysis

`func NewCompanyDataAnalysis() *CompanyDataAnalysis`

NewCompanyDataAnalysis instantiates a new CompanyDataAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyDataAnalysisWithDefaults

`func NewCompanyDataAnalysisWithDefaults() *CompanyDataAnalysis`

NewCompanyDataAnalysisWithDefaults instantiates a new CompanyDataAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyDataAnalysis) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyDataAnalysis) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyDataAnalysis) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyDataAnalysis) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSharePrice

`func (o *CompanyDataAnalysis) GetSharePrice() float64`

GetSharePrice returns the SharePrice field if non-nil, zero value otherwise.

### GetSharePriceOk

`func (o *CompanyDataAnalysis) GetSharePriceOk() (*float64, bool)`

GetSharePriceOk returns a tuple with the SharePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePrice

`func (o *CompanyDataAnalysis) SetSharePrice(v float64)`

SetSharePrice sets SharePrice field to given value.

### HasSharePrice

`func (o *CompanyDataAnalysis) HasSharePrice() bool`

HasSharePrice returns a boolean if a field has been set.

### GetMarketCap

`func (o *CompanyDataAnalysis) GetMarketCap() float64`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *CompanyDataAnalysis) GetMarketCapOk() (*float64, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *CompanyDataAnalysis) SetMarketCap(v float64)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *CompanyDataAnalysis) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetIntrinsicDiscount

`func (o *CompanyDataAnalysis) GetIntrinsicDiscount() float64`

GetIntrinsicDiscount returns the IntrinsicDiscount field if non-nil, zero value otherwise.

### GetIntrinsicDiscountOk

`func (o *CompanyDataAnalysis) GetIntrinsicDiscountOk() (*float64, bool)`

GetIntrinsicDiscountOk returns a tuple with the IntrinsicDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrinsicDiscount

`func (o *CompanyDataAnalysis) SetIntrinsicDiscount(v float64)`

SetIntrinsicDiscount sets IntrinsicDiscount field to given value.

### HasIntrinsicDiscount

`func (o *CompanyDataAnalysis) HasIntrinsicDiscount() bool`

HasIntrinsicDiscount returns a boolean if a field has been set.

### GetPe

`func (o *CompanyDataAnalysis) GetPe() float64`

GetPe returns the Pe field if non-nil, zero value otherwise.

### GetPeOk

`func (o *CompanyDataAnalysis) GetPeOk() (*float64, bool)`

GetPeOk returns a tuple with the Pe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPe

`func (o *CompanyDataAnalysis) SetPe(v float64)`

SetPe sets Pe field to given value.

### HasPe

`func (o *CompanyDataAnalysis) HasPe() bool`

HasPe returns a boolean if a field has been set.

### GetPb

`func (o *CompanyDataAnalysis) GetPb() float64`

GetPb returns the Pb field if non-nil, zero value otherwise.

### GetPbOk

`func (o *CompanyDataAnalysis) GetPbOk() (*float64, bool)`

GetPbOk returns a tuple with the Pb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPb

`func (o *CompanyDataAnalysis) SetPb(v float64)`

SetPb sets Pb field to given value.

### HasPb

`func (o *CompanyDataAnalysis) HasPb() bool`

HasPb returns a boolean if a field has been set.

### GetPeg

`func (o *CompanyDataAnalysis) GetPeg() float64`

GetPeg returns the Peg field if non-nil, zero value otherwise.

### GetPegOk

`func (o *CompanyDataAnalysis) GetPegOk() (*float64, bool)`

GetPegOk returns a tuple with the Peg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeg

`func (o *CompanyDataAnalysis) SetPeg(v float64)`

SetPeg sets Peg field to given value.

### HasPeg

`func (o *CompanyDataAnalysis) HasPeg() bool`

HasPeg returns a boolean if a field has been set.

### GetRoe

`func (o *CompanyDataAnalysis) GetRoe() float64`

GetRoe returns the Roe field if non-nil, zero value otherwise.

### GetRoeOk

`func (o *CompanyDataAnalysis) GetRoeOk() (*float64, bool)`

GetRoeOk returns a tuple with the Roe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoe

`func (o *CompanyDataAnalysis) SetRoe(v float64)`

SetRoe sets Roe field to given value.

### HasRoe

`func (o *CompanyDataAnalysis) HasRoe() bool`

HasRoe returns a boolean if a field has been set.

### GetRoa

`func (o *CompanyDataAnalysis) GetRoa() float64`

GetRoa returns the Roa field if non-nil, zero value otherwise.

### GetRoaOk

`func (o *CompanyDataAnalysis) GetRoaOk() (*float64, bool)`

GetRoaOk returns a tuple with the Roa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoa

`func (o *CompanyDataAnalysis) SetRoa(v float64)`

SetRoa sets Roa field to given value.

### HasRoa

`func (o *CompanyDataAnalysis) HasRoa() bool`

HasRoa returns a boolean if a field has been set.

### GetEps

`func (o *CompanyDataAnalysis) GetEps() float64`

GetEps returns the Eps field if non-nil, zero value otherwise.

### GetEpsOk

`func (o *CompanyDataAnalysis) GetEpsOk() (*float64, bool)`

GetEpsOk returns a tuple with the Eps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEps

`func (o *CompanyDataAnalysis) SetEps(v float64)`

SetEps sets Eps field to given value.

### HasEps

`func (o *CompanyDataAnalysis) HasEps() bool`

HasEps returns a boolean if a field has been set.

### GetDebtEquity

`func (o *CompanyDataAnalysis) GetDebtEquity() float64`

GetDebtEquity returns the DebtEquity field if non-nil, zero value otherwise.

### GetDebtEquityOk

`func (o *CompanyDataAnalysis) GetDebtEquityOk() (*float64, bool)`

GetDebtEquityOk returns a tuple with the DebtEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtEquity

`func (o *CompanyDataAnalysis) SetDebtEquity(v float64)`

SetDebtEquity sets DebtEquity field to given value.

### HasDebtEquity

`func (o *CompanyDataAnalysis) HasDebtEquity() bool`

HasDebtEquity returns a boolean if a field has been set.

### GetAnalystCount

`func (o *CompanyDataAnalysis) GetAnalystCount() float64`

GetAnalystCount returns the AnalystCount field if non-nil, zero value otherwise.

### GetAnalystCountOk

`func (o *CompanyDataAnalysis) GetAnalystCountOk() (*float64, bool)`

GetAnalystCountOk returns a tuple with the AnalystCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystCount

`func (o *CompanyDataAnalysis) SetAnalystCount(v float64)`

SetAnalystCount sets AnalystCount field to given value.

### HasAnalystCount

`func (o *CompanyDataAnalysis) HasAnalystCount() bool`

HasAnalystCount returns a boolean if a field has been set.

### GetDividend

`func (o *CompanyDataAnalysis) GetDividend() CompanyAnalysisDividend`

GetDividend returns the Dividend field if non-nil, zero value otherwise.

### GetDividendOk

`func (o *CompanyDataAnalysis) GetDividendOk() (*CompanyAnalysisDividend, bool)`

GetDividendOk returns a tuple with the Dividend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividend

`func (o *CompanyDataAnalysis) SetDividend(v CompanyAnalysisDividend)`

SetDividend sets Dividend field to given value.

### HasDividend

`func (o *CompanyDataAnalysis) HasDividend() bool`

HasDividend returns a boolean if a field has been set.

### GetFuture

`func (o *CompanyDataAnalysis) GetFuture() CompanyAnalysisFuture`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *CompanyDataAnalysis) GetFutureOk() (*CompanyAnalysisFuture, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *CompanyDataAnalysis) SetFuture(v CompanyAnalysisFuture)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *CompanyDataAnalysis) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetPast

`func (o *CompanyDataAnalysis) GetPast() CompanyAnalysisPast`

GetPast returns the Past field if non-nil, zero value otherwise.

### GetPastOk

`func (o *CompanyDataAnalysis) GetPastOk() (*CompanyAnalysisPast, bool)`

GetPastOk returns a tuple with the Past field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPast

`func (o *CompanyDataAnalysis) SetPast(v CompanyAnalysisPast)`

SetPast sets Past field to given value.

### HasPast

`func (o *CompanyDataAnalysis) HasPast() bool`

HasPast returns a boolean if a field has been set.

### GetExtended

`func (o *CompanyDataAnalysis) GetExtended() CompanyAnalysisExtendedData`

GetExtended returns the Extended field if non-nil, zero value otherwise.

### GetExtendedOk

`func (o *CompanyDataAnalysis) GetExtendedOk() (*CompanyAnalysisExtendedData, bool)`

GetExtendedOk returns a tuple with the Extended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtended

`func (o *CompanyDataAnalysis) SetExtended(v CompanyAnalysisExtendedData)`

SetExtended sets Extended field to given value.

### HasExtended

`func (o *CompanyDataAnalysis) HasExtended() bool`

HasExtended returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


