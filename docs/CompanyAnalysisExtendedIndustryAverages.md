# CompanyAnalysisExtendedIndustryAverages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndustryId** | Pointer to **int64** |  | [optional] 
**CountryIso** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ValueScore** | Pointer to **float64** |  | [optional] 
**DividendsScore** | Pointer to **float64** |  | [optional] 
**FuturePerformanceScore** | Pointer to **float64** |  | [optional] 
**HealthScore** | Pointer to **float64** |  | [optional] 
**PastPerformanceScore** | Pointer to **float64** |  | [optional] 
**TotalScore** | Pointer to **float64** |  | [optional] 
**SharePrice** | Pointer to **float64** |  | [optional] 
**MarketCap** | Pointer to **float64** |  | [optional] 
**IntrinsicDiscount** | Pointer to **float64** |  | [optional] 
**AnalystCount** | Pointer to **int32** |  | [optional] 
**PE** | Pointer to **float64** |  | [optional] 
**PB** | Pointer to **float64** |  | [optional] 
**PEG** | Pointer to **float64** |  | [optional] 
**FutureOneYearGrowth** | Pointer to **float64** |  | [optional] 
**FutureThreeYearGrowth** | Pointer to **float64** |  | [optional] 
**HistoricalDividendYield** | Pointer to **float64** |  | [optional] 
**FutureOneYearROE** | Pointer to **NullableFloat64** |  | [optional] 
**FutureThreeYearROE** | Pointer to **NullableFloat64** |  | [optional] 
**PastOneYearGrowth** | Pointer to **float64** |  | [optional] 
**PastFiveYearGrowth** | Pointer to **float64** |  | [optional] 
**ROE** | Pointer to **float64** |  | [optional] 
**ROA** | Pointer to **float64** |  | [optional] 
**DividendYield** | Pointer to **float64** |  | [optional] 
**FutureDividendYield** | Pointer to **float64** |  | [optional] 
**PayoutRatio** | Pointer to **NullableFloat64** |  | [optional] 
**EPS** | Pointer to **float64** |  | [optional] 
**InsiderBuying** | Pointer to **float64** |  | [optional] 
**DebtEquity** | Pointer to **float64** |  | [optional] 
**LeveredBeta** | Pointer to **float64** |  | [optional] 
**UnleveredBeta** | Pointer to **float64** |  | [optional] 
**TotalBaseCount** | Pointer to **int32** |  | [optional] 
**ProfitableCount** | Pointer to **int32** |  | [optional] 
**AnalystCoverageCount** | Pointer to **int32** |  | [optional] 
**DividendCount** | Pointer to **int32** |  | [optional] 
**BetaCount** | Pointer to **int32** |  | [optional] 
**EarningsPerShareGrowthAnnual** | Pointer to **float64** |  | [optional] 
**NetIncomeGrowthAnnual** | Pointer to **float64** |  | [optional] 
**CashOpsGrowthAnnual** | Pointer to **float64** |  | [optional] 
**RevenueGrowthAnnual** | Pointer to **float64** |  | [optional] 
**LeveredBetaMedian** | Pointer to **float64** |  | [optional] 
**BaseSource** | Pointer to **string** |  | [optional] 
**ProfitableSource** | Pointer to **string** |  | [optional] 
**AnalystSource** | Pointer to **string** |  | [optional] 
**DividendSource** | Pointer to **string** |  | [optional] 
**BetaSource** | Pointer to **string** |  | [optional] 
**All** | Pointer to [**CompanyAnalysisExtendedIndustryAveragesAll**](CompanyAnalysisExtendedIndustryAveragesAll.md) |  | [optional] 

## Methods

### NewCompanyAnalysisExtendedIndustryAverages

`func NewCompanyAnalysisExtendedIndustryAverages() *CompanyAnalysisExtendedIndustryAverages`

NewCompanyAnalysisExtendedIndustryAverages instantiates a new CompanyAnalysisExtendedIndustryAverages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAnalysisExtendedIndustryAveragesWithDefaults

`func NewCompanyAnalysisExtendedIndustryAveragesWithDefaults() *CompanyAnalysisExtendedIndustryAverages`

NewCompanyAnalysisExtendedIndustryAveragesWithDefaults instantiates a new CompanyAnalysisExtendedIndustryAverages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndustryId

`func (o *CompanyAnalysisExtendedIndustryAverages) GetIndustryId() int64`

GetIndustryId returns the IndustryId field if non-nil, zero value otherwise.

### GetIndustryIdOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetIndustryIdOk() (*int64, bool)`

GetIndustryIdOk returns a tuple with the IndustryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustryId

`func (o *CompanyAnalysisExtendedIndustryAverages) SetIndustryId(v int64)`

SetIndustryId sets IndustryId field to given value.

### HasIndustryId

`func (o *CompanyAnalysisExtendedIndustryAverages) HasIndustryId() bool`

HasIndustryId returns a boolean if a field has been set.

### GetCountryIso

`func (o *CompanyAnalysisExtendedIndustryAverages) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *CompanyAnalysisExtendedIndustryAverages) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *CompanyAnalysisExtendedIndustryAverages) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### GetName

`func (o *CompanyAnalysisExtendedIndustryAverages) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyAnalysisExtendedIndustryAverages) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyAnalysisExtendedIndustryAverages) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValueScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetValueScore() float64`

GetValueScore returns the ValueScore field if non-nil, zero value otherwise.

### GetValueScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetValueScoreOk() (*float64, bool)`

GetValueScoreOk returns a tuple with the ValueScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetValueScore(v float64)`

SetValueScore sets ValueScore field to given value.

### HasValueScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasValueScore() bool`

HasValueScore returns a boolean if a field has been set.

### GetDividendsScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendsScore() float64`

GetDividendsScore returns the DividendsScore field if non-nil, zero value otherwise.

### GetDividendsScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendsScoreOk() (*float64, bool)`

GetDividendsScoreOk returns a tuple with the DividendsScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendsScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetDividendsScore(v float64)`

SetDividendsScore sets DividendsScore field to given value.

### HasDividendsScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasDividendsScore() bool`

HasDividendsScore returns a boolean if a field has been set.

### GetFuturePerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFuturePerformanceScore() float64`

GetFuturePerformanceScore returns the FuturePerformanceScore field if non-nil, zero value otherwise.

### GetFuturePerformanceScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFuturePerformanceScoreOk() (*float64, bool)`

GetFuturePerformanceScoreOk returns a tuple with the FuturePerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuturePerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFuturePerformanceScore(v float64)`

SetFuturePerformanceScore sets FuturePerformanceScore field to given value.

### HasFuturePerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFuturePerformanceScore() bool`

HasFuturePerformanceScore returns a boolean if a field has been set.

### GetHealthScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetHealthScore() float64`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetHealthScoreOk() (*float64, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetHealthScore(v float64)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetPastPerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastPerformanceScore() float64`

GetPastPerformanceScore returns the PastPerformanceScore field if non-nil, zero value otherwise.

### GetPastPerformanceScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastPerformanceScoreOk() (*float64, bool)`

GetPastPerformanceScoreOk returns a tuple with the PastPerformanceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastPerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPastPerformanceScore(v float64)`

SetPastPerformanceScore sets PastPerformanceScore field to given value.

### HasPastPerformanceScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPastPerformanceScore() bool`

HasPastPerformanceScore returns a boolean if a field has been set.

### GetTotalScore

`func (o *CompanyAnalysisExtendedIndustryAverages) GetTotalScore() float64`

GetTotalScore returns the TotalScore field if non-nil, zero value otherwise.

### GetTotalScoreOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetTotalScoreOk() (*float64, bool)`

GetTotalScoreOk returns a tuple with the TotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScore

`func (o *CompanyAnalysisExtendedIndustryAverages) SetTotalScore(v float64)`

SetTotalScore sets TotalScore field to given value.

### HasTotalScore

`func (o *CompanyAnalysisExtendedIndustryAverages) HasTotalScore() bool`

HasTotalScore returns a boolean if a field has been set.

### GetSharePrice

`func (o *CompanyAnalysisExtendedIndustryAverages) GetSharePrice() float64`

GetSharePrice returns the SharePrice field if non-nil, zero value otherwise.

### GetSharePriceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetSharePriceOk() (*float64, bool)`

GetSharePriceOk returns a tuple with the SharePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePrice

`func (o *CompanyAnalysisExtendedIndustryAverages) SetSharePrice(v float64)`

SetSharePrice sets SharePrice field to given value.

### HasSharePrice

`func (o *CompanyAnalysisExtendedIndustryAverages) HasSharePrice() bool`

HasSharePrice returns a boolean if a field has been set.

### GetMarketCap

`func (o *CompanyAnalysisExtendedIndustryAverages) GetMarketCap() float64`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetMarketCapOk() (*float64, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *CompanyAnalysisExtendedIndustryAverages) SetMarketCap(v float64)`

SetMarketCap sets MarketCap field to given value.

### HasMarketCap

`func (o *CompanyAnalysisExtendedIndustryAverages) HasMarketCap() bool`

HasMarketCap returns a boolean if a field has been set.

### GetIntrinsicDiscount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetIntrinsicDiscount() float64`

GetIntrinsicDiscount returns the IntrinsicDiscount field if non-nil, zero value otherwise.

### GetIntrinsicDiscountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetIntrinsicDiscountOk() (*float64, bool)`

GetIntrinsicDiscountOk returns a tuple with the IntrinsicDiscount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntrinsicDiscount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetIntrinsicDiscount(v float64)`

SetIntrinsicDiscount sets IntrinsicDiscount field to given value.

### HasIntrinsicDiscount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasIntrinsicDiscount() bool`

HasIntrinsicDiscount returns a boolean if a field has been set.

### GetAnalystCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystCount() int32`

GetAnalystCount returns the AnalystCount field if non-nil, zero value otherwise.

### GetAnalystCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystCountOk() (*int32, bool)`

GetAnalystCountOk returns a tuple with the AnalystCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetAnalystCount(v int32)`

SetAnalystCount sets AnalystCount field to given value.

### HasAnalystCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasAnalystCount() bool`

HasAnalystCount returns a boolean if a field has been set.

### GetPE

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPE() float64`

GetPE returns the PE field if non-nil, zero value otherwise.

### GetPEOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPEOk() (*float64, bool)`

GetPEOk returns a tuple with the PE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPE

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPE(v float64)`

SetPE sets PE field to given value.

### HasPE

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPE() bool`

HasPE returns a boolean if a field has been set.

### GetPB

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPB() float64`

GetPB returns the PB field if non-nil, zero value otherwise.

### GetPBOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPBOk() (*float64, bool)`

GetPBOk returns a tuple with the PB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPB

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPB(v float64)`

SetPB sets PB field to given value.

### HasPB

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPB() bool`

HasPB returns a boolean if a field has been set.

### GetPEG

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPEG() float64`

GetPEG returns the PEG field if non-nil, zero value otherwise.

### GetPEGOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPEGOk() (*float64, bool)`

GetPEGOk returns a tuple with the PEG field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPEG

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPEG(v float64)`

SetPEG sets PEG field to given value.

### HasPEG

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPEG() bool`

HasPEG returns a boolean if a field has been set.

### GetFutureOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureOneYearGrowth() float64`

GetFutureOneYearGrowth returns the FutureOneYearGrowth field if non-nil, zero value otherwise.

### GetFutureOneYearGrowthOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureOneYearGrowthOk() (*float64, bool)`

GetFutureOneYearGrowthOk returns a tuple with the FutureOneYearGrowth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureOneYearGrowth(v float64)`

SetFutureOneYearGrowth sets FutureOneYearGrowth field to given value.

### HasFutureOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFutureOneYearGrowth() bool`

HasFutureOneYearGrowth returns a boolean if a field has been set.

### GetFutureThreeYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureThreeYearGrowth() float64`

GetFutureThreeYearGrowth returns the FutureThreeYearGrowth field if non-nil, zero value otherwise.

### GetFutureThreeYearGrowthOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureThreeYearGrowthOk() (*float64, bool)`

GetFutureThreeYearGrowthOk returns a tuple with the FutureThreeYearGrowth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureThreeYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureThreeYearGrowth(v float64)`

SetFutureThreeYearGrowth sets FutureThreeYearGrowth field to given value.

### HasFutureThreeYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFutureThreeYearGrowth() bool`

HasFutureThreeYearGrowth returns a boolean if a field has been set.

### GetHistoricalDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) GetHistoricalDividendYield() float64`

GetHistoricalDividendYield returns the HistoricalDividendYield field if non-nil, zero value otherwise.

### GetHistoricalDividendYieldOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetHistoricalDividendYieldOk() (*float64, bool)`

GetHistoricalDividendYieldOk returns a tuple with the HistoricalDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoricalDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) SetHistoricalDividendYield(v float64)`

SetHistoricalDividendYield sets HistoricalDividendYield field to given value.

### HasHistoricalDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) HasHistoricalDividendYield() bool`

HasHistoricalDividendYield returns a boolean if a field has been set.

### GetFutureOneYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureOneYearROE() float64`

GetFutureOneYearROE returns the FutureOneYearROE field if non-nil, zero value otherwise.

### GetFutureOneYearROEOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureOneYearROEOk() (*float64, bool)`

GetFutureOneYearROEOk returns a tuple with the FutureOneYearROE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureOneYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureOneYearROE(v float64)`

SetFutureOneYearROE sets FutureOneYearROE field to given value.

### HasFutureOneYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFutureOneYearROE() bool`

HasFutureOneYearROE returns a boolean if a field has been set.

### SetFutureOneYearROENil

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureOneYearROENil(b bool)`

 SetFutureOneYearROENil sets the value for FutureOneYearROE to be an explicit nil

### UnsetFutureOneYearROE
`func (o *CompanyAnalysisExtendedIndustryAverages) UnsetFutureOneYearROE()`

UnsetFutureOneYearROE ensures that no value is present for FutureOneYearROE, not even an explicit nil
### GetFutureThreeYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureThreeYearROE() float64`

GetFutureThreeYearROE returns the FutureThreeYearROE field if non-nil, zero value otherwise.

### GetFutureThreeYearROEOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureThreeYearROEOk() (*float64, bool)`

GetFutureThreeYearROEOk returns a tuple with the FutureThreeYearROE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureThreeYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureThreeYearROE(v float64)`

SetFutureThreeYearROE sets FutureThreeYearROE field to given value.

### HasFutureThreeYearROE

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFutureThreeYearROE() bool`

HasFutureThreeYearROE returns a boolean if a field has been set.

### SetFutureThreeYearROENil

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureThreeYearROENil(b bool)`

 SetFutureThreeYearROENil sets the value for FutureThreeYearROE to be an explicit nil

### UnsetFutureThreeYearROE
`func (o *CompanyAnalysisExtendedIndustryAverages) UnsetFutureThreeYearROE()`

UnsetFutureThreeYearROE ensures that no value is present for FutureThreeYearROE, not even an explicit nil
### GetPastOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastOneYearGrowth() float64`

GetPastOneYearGrowth returns the PastOneYearGrowth field if non-nil, zero value otherwise.

### GetPastOneYearGrowthOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastOneYearGrowthOk() (*float64, bool)`

GetPastOneYearGrowthOk returns a tuple with the PastOneYearGrowth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPastOneYearGrowth(v float64)`

SetPastOneYearGrowth sets PastOneYearGrowth field to given value.

### HasPastOneYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPastOneYearGrowth() bool`

HasPastOneYearGrowth returns a boolean if a field has been set.

### GetPastFiveYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastFiveYearGrowth() float64`

GetPastFiveYearGrowth returns the PastFiveYearGrowth field if non-nil, zero value otherwise.

### GetPastFiveYearGrowthOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPastFiveYearGrowthOk() (*float64, bool)`

GetPastFiveYearGrowthOk returns a tuple with the PastFiveYearGrowth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPastFiveYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPastFiveYearGrowth(v float64)`

SetPastFiveYearGrowth sets PastFiveYearGrowth field to given value.

### HasPastFiveYearGrowth

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPastFiveYearGrowth() bool`

HasPastFiveYearGrowth returns a boolean if a field has been set.

### GetROE

`func (o *CompanyAnalysisExtendedIndustryAverages) GetROE() float64`

GetROE returns the ROE field if non-nil, zero value otherwise.

### GetROEOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetROEOk() (*float64, bool)`

GetROEOk returns a tuple with the ROE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROE

`func (o *CompanyAnalysisExtendedIndustryAverages) SetROE(v float64)`

SetROE sets ROE field to given value.

### HasROE

`func (o *CompanyAnalysisExtendedIndustryAverages) HasROE() bool`

HasROE returns a boolean if a field has been set.

### GetROA

`func (o *CompanyAnalysisExtendedIndustryAverages) GetROA() float64`

GetROA returns the ROA field if non-nil, zero value otherwise.

### GetROAOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetROAOk() (*float64, bool)`

GetROAOk returns a tuple with the ROA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetROA

`func (o *CompanyAnalysisExtendedIndustryAverages) SetROA(v float64)`

SetROA sets ROA field to given value.

### HasROA

`func (o *CompanyAnalysisExtendedIndustryAverages) HasROA() bool`

HasROA returns a boolean if a field has been set.

### GetDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendYield() float64`

GetDividendYield returns the DividendYield field if non-nil, zero value otherwise.

### GetDividendYieldOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendYieldOk() (*float64, bool)`

GetDividendYieldOk returns a tuple with the DividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) SetDividendYield(v float64)`

SetDividendYield sets DividendYield field to given value.

### HasDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) HasDividendYield() bool`

HasDividendYield returns a boolean if a field has been set.

### GetFutureDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureDividendYield() float64`

GetFutureDividendYield returns the FutureDividendYield field if non-nil, zero value otherwise.

### GetFutureDividendYieldOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetFutureDividendYieldOk() (*float64, bool)`

GetFutureDividendYieldOk returns a tuple with the FutureDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) SetFutureDividendYield(v float64)`

SetFutureDividendYield sets FutureDividendYield field to given value.

### HasFutureDividendYield

`func (o *CompanyAnalysisExtendedIndustryAverages) HasFutureDividendYield() bool`

HasFutureDividendYield returns a boolean if a field has been set.

### GetPayoutRatio

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPayoutRatio() float64`

GetPayoutRatio returns the PayoutRatio field if non-nil, zero value otherwise.

### GetPayoutRatioOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetPayoutRatioOk() (*float64, bool)`

GetPayoutRatioOk returns a tuple with the PayoutRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutRatio

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPayoutRatio(v float64)`

SetPayoutRatio sets PayoutRatio field to given value.

### HasPayoutRatio

`func (o *CompanyAnalysisExtendedIndustryAverages) HasPayoutRatio() bool`

HasPayoutRatio returns a boolean if a field has been set.

### SetPayoutRatioNil

`func (o *CompanyAnalysisExtendedIndustryAverages) SetPayoutRatioNil(b bool)`

 SetPayoutRatioNil sets the value for PayoutRatio to be an explicit nil

### UnsetPayoutRatio
`func (o *CompanyAnalysisExtendedIndustryAverages) UnsetPayoutRatio()`

UnsetPayoutRatio ensures that no value is present for PayoutRatio, not even an explicit nil
### GetEPS

`func (o *CompanyAnalysisExtendedIndustryAverages) GetEPS() float64`

GetEPS returns the EPS field if non-nil, zero value otherwise.

### GetEPSOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetEPSOk() (*float64, bool)`

GetEPSOk returns a tuple with the EPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEPS

`func (o *CompanyAnalysisExtendedIndustryAverages) SetEPS(v float64)`

SetEPS sets EPS field to given value.

### HasEPS

`func (o *CompanyAnalysisExtendedIndustryAverages) HasEPS() bool`

HasEPS returns a boolean if a field has been set.

### GetInsiderBuying

`func (o *CompanyAnalysisExtendedIndustryAverages) GetInsiderBuying() float64`

GetInsiderBuying returns the InsiderBuying field if non-nil, zero value otherwise.

### GetInsiderBuyingOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetInsiderBuyingOk() (*float64, bool)`

GetInsiderBuyingOk returns a tuple with the InsiderBuying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderBuying

`func (o *CompanyAnalysisExtendedIndustryAverages) SetInsiderBuying(v float64)`

SetInsiderBuying sets InsiderBuying field to given value.

### HasInsiderBuying

`func (o *CompanyAnalysisExtendedIndustryAverages) HasInsiderBuying() bool`

HasInsiderBuying returns a boolean if a field has been set.

### GetDebtEquity

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDebtEquity() float64`

GetDebtEquity returns the DebtEquity field if non-nil, zero value otherwise.

### GetDebtEquityOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDebtEquityOk() (*float64, bool)`

GetDebtEquityOk returns a tuple with the DebtEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtEquity

`func (o *CompanyAnalysisExtendedIndustryAverages) SetDebtEquity(v float64)`

SetDebtEquity sets DebtEquity field to given value.

### HasDebtEquity

`func (o *CompanyAnalysisExtendedIndustryAverages) HasDebtEquity() bool`

HasDebtEquity returns a boolean if a field has been set.

### GetLeveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) GetLeveredBeta() float64`

GetLeveredBeta returns the LeveredBeta field if non-nil, zero value otherwise.

### GetLeveredBetaOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetLeveredBetaOk() (*float64, bool)`

GetLeveredBetaOk returns a tuple with the LeveredBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) SetLeveredBeta(v float64)`

SetLeveredBeta sets LeveredBeta field to given value.

### HasLeveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) HasLeveredBeta() bool`

HasLeveredBeta returns a boolean if a field has been set.

### GetUnleveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) GetUnleveredBeta() float64`

GetUnleveredBeta returns the UnleveredBeta field if non-nil, zero value otherwise.

### GetUnleveredBetaOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetUnleveredBetaOk() (*float64, bool)`

GetUnleveredBetaOk returns a tuple with the UnleveredBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnleveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) SetUnleveredBeta(v float64)`

SetUnleveredBeta sets UnleveredBeta field to given value.

### HasUnleveredBeta

`func (o *CompanyAnalysisExtendedIndustryAverages) HasUnleveredBeta() bool`

HasUnleveredBeta returns a boolean if a field has been set.

### GetTotalBaseCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetTotalBaseCount() int32`

GetTotalBaseCount returns the TotalBaseCount field if non-nil, zero value otherwise.

### GetTotalBaseCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetTotalBaseCountOk() (*int32, bool)`

GetTotalBaseCountOk returns a tuple with the TotalBaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBaseCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetTotalBaseCount(v int32)`

SetTotalBaseCount sets TotalBaseCount field to given value.

### HasTotalBaseCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasTotalBaseCount() bool`

HasTotalBaseCount returns a boolean if a field has been set.

### GetProfitableCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetProfitableCount() int32`

GetProfitableCount returns the ProfitableCount field if non-nil, zero value otherwise.

### GetProfitableCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetProfitableCountOk() (*int32, bool)`

GetProfitableCountOk returns a tuple with the ProfitableCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitableCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetProfitableCount(v int32)`

SetProfitableCount sets ProfitableCount field to given value.

### HasProfitableCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasProfitableCount() bool`

HasProfitableCount returns a boolean if a field has been set.

### GetAnalystCoverageCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystCoverageCount() int32`

GetAnalystCoverageCount returns the AnalystCoverageCount field if non-nil, zero value otherwise.

### GetAnalystCoverageCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystCoverageCountOk() (*int32, bool)`

GetAnalystCoverageCountOk returns a tuple with the AnalystCoverageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystCoverageCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetAnalystCoverageCount(v int32)`

SetAnalystCoverageCount sets AnalystCoverageCount field to given value.

### HasAnalystCoverageCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasAnalystCoverageCount() bool`

HasAnalystCoverageCount returns a boolean if a field has been set.

### GetDividendCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendCount() int32`

GetDividendCount returns the DividendCount field if non-nil, zero value otherwise.

### GetDividendCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendCountOk() (*int32, bool)`

GetDividendCountOk returns a tuple with the DividendCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetDividendCount(v int32)`

SetDividendCount sets DividendCount field to given value.

### HasDividendCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasDividendCount() bool`

HasDividendCount returns a boolean if a field has been set.

### GetBetaCount

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBetaCount() int32`

GetBetaCount returns the BetaCount field if non-nil, zero value otherwise.

### GetBetaCountOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBetaCountOk() (*int32, bool)`

GetBetaCountOk returns a tuple with the BetaCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaCount

`func (o *CompanyAnalysisExtendedIndustryAverages) SetBetaCount(v int32)`

SetBetaCount sets BetaCount field to given value.

### HasBetaCount

`func (o *CompanyAnalysisExtendedIndustryAverages) HasBetaCount() bool`

HasBetaCount returns a boolean if a field has been set.

### GetEarningsPerShareGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) GetEarningsPerShareGrowthAnnual() float64`

GetEarningsPerShareGrowthAnnual returns the EarningsPerShareGrowthAnnual field if non-nil, zero value otherwise.

### GetEarningsPerShareGrowthAnnualOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetEarningsPerShareGrowthAnnualOk() (*float64, bool)`

GetEarningsPerShareGrowthAnnualOk returns a tuple with the EarningsPerShareGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningsPerShareGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) SetEarningsPerShareGrowthAnnual(v float64)`

SetEarningsPerShareGrowthAnnual sets EarningsPerShareGrowthAnnual field to given value.

### HasEarningsPerShareGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) HasEarningsPerShareGrowthAnnual() bool`

HasEarningsPerShareGrowthAnnual returns a boolean if a field has been set.

### GetNetIncomeGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) GetNetIncomeGrowthAnnual() float64`

GetNetIncomeGrowthAnnual returns the NetIncomeGrowthAnnual field if non-nil, zero value otherwise.

### GetNetIncomeGrowthAnnualOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetNetIncomeGrowthAnnualOk() (*float64, bool)`

GetNetIncomeGrowthAnnualOk returns a tuple with the NetIncomeGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncomeGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) SetNetIncomeGrowthAnnual(v float64)`

SetNetIncomeGrowthAnnual sets NetIncomeGrowthAnnual field to given value.

### HasNetIncomeGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) HasNetIncomeGrowthAnnual() bool`

HasNetIncomeGrowthAnnual returns a boolean if a field has been set.

### GetCashOpsGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) GetCashOpsGrowthAnnual() float64`

GetCashOpsGrowthAnnual returns the CashOpsGrowthAnnual field if non-nil, zero value otherwise.

### GetCashOpsGrowthAnnualOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetCashOpsGrowthAnnualOk() (*float64, bool)`

GetCashOpsGrowthAnnualOk returns a tuple with the CashOpsGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOpsGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) SetCashOpsGrowthAnnual(v float64)`

SetCashOpsGrowthAnnual sets CashOpsGrowthAnnual field to given value.

### HasCashOpsGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) HasCashOpsGrowthAnnual() bool`

HasCashOpsGrowthAnnual returns a boolean if a field has been set.

### GetRevenueGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) GetRevenueGrowthAnnual() float64`

GetRevenueGrowthAnnual returns the RevenueGrowthAnnual field if non-nil, zero value otherwise.

### GetRevenueGrowthAnnualOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetRevenueGrowthAnnualOk() (*float64, bool)`

GetRevenueGrowthAnnualOk returns a tuple with the RevenueGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) SetRevenueGrowthAnnual(v float64)`

SetRevenueGrowthAnnual sets RevenueGrowthAnnual field to given value.

### HasRevenueGrowthAnnual

`func (o *CompanyAnalysisExtendedIndustryAverages) HasRevenueGrowthAnnual() bool`

HasRevenueGrowthAnnual returns a boolean if a field has been set.

### GetLeveredBetaMedian

`func (o *CompanyAnalysisExtendedIndustryAverages) GetLeveredBetaMedian() float64`

GetLeveredBetaMedian returns the LeveredBetaMedian field if non-nil, zero value otherwise.

### GetLeveredBetaMedianOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetLeveredBetaMedianOk() (*float64, bool)`

GetLeveredBetaMedianOk returns a tuple with the LeveredBetaMedian field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredBetaMedian

`func (o *CompanyAnalysisExtendedIndustryAverages) SetLeveredBetaMedian(v float64)`

SetLeveredBetaMedian sets LeveredBetaMedian field to given value.

### HasLeveredBetaMedian

`func (o *CompanyAnalysisExtendedIndustryAverages) HasLeveredBetaMedian() bool`

HasLeveredBetaMedian returns a boolean if a field has been set.

### GetBaseSource

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBaseSource() string`

GetBaseSource returns the BaseSource field if non-nil, zero value otherwise.

### GetBaseSourceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBaseSourceOk() (*string, bool)`

GetBaseSourceOk returns a tuple with the BaseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSource

`func (o *CompanyAnalysisExtendedIndustryAverages) SetBaseSource(v string)`

SetBaseSource sets BaseSource field to given value.

### HasBaseSource

`func (o *CompanyAnalysisExtendedIndustryAverages) HasBaseSource() bool`

HasBaseSource returns a boolean if a field has been set.

### GetProfitableSource

`func (o *CompanyAnalysisExtendedIndustryAverages) GetProfitableSource() string`

GetProfitableSource returns the ProfitableSource field if non-nil, zero value otherwise.

### GetProfitableSourceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetProfitableSourceOk() (*string, bool)`

GetProfitableSourceOk returns a tuple with the ProfitableSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitableSource

`func (o *CompanyAnalysisExtendedIndustryAverages) SetProfitableSource(v string)`

SetProfitableSource sets ProfitableSource field to given value.

### HasProfitableSource

`func (o *CompanyAnalysisExtendedIndustryAverages) HasProfitableSource() bool`

HasProfitableSource returns a boolean if a field has been set.

### GetAnalystSource

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystSource() string`

GetAnalystSource returns the AnalystSource field if non-nil, zero value otherwise.

### GetAnalystSourceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAnalystSourceOk() (*string, bool)`

GetAnalystSourceOk returns a tuple with the AnalystSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalystSource

`func (o *CompanyAnalysisExtendedIndustryAverages) SetAnalystSource(v string)`

SetAnalystSource sets AnalystSource field to given value.

### HasAnalystSource

`func (o *CompanyAnalysisExtendedIndustryAverages) HasAnalystSource() bool`

HasAnalystSource returns a boolean if a field has been set.

### GetDividendSource

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendSource() string`

GetDividendSource returns the DividendSource field if non-nil, zero value otherwise.

### GetDividendSourceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetDividendSourceOk() (*string, bool)`

GetDividendSourceOk returns a tuple with the DividendSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendSource

`func (o *CompanyAnalysisExtendedIndustryAverages) SetDividendSource(v string)`

SetDividendSource sets DividendSource field to given value.

### HasDividendSource

`func (o *CompanyAnalysisExtendedIndustryAverages) HasDividendSource() bool`

HasDividendSource returns a boolean if a field has been set.

### GetBetaSource

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBetaSource() string`

GetBetaSource returns the BetaSource field if non-nil, zero value otherwise.

### GetBetaSourceOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetBetaSourceOk() (*string, bool)`

GetBetaSourceOk returns a tuple with the BetaSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaSource

`func (o *CompanyAnalysisExtendedIndustryAverages) SetBetaSource(v string)`

SetBetaSource sets BetaSource field to given value.

### HasBetaSource

`func (o *CompanyAnalysisExtendedIndustryAverages) HasBetaSource() bool`

HasBetaSource returns a boolean if a field has been set.

### GetAll

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAll() CompanyAnalysisExtendedIndustryAveragesAll`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CompanyAnalysisExtendedIndustryAverages) GetAllOk() (*CompanyAnalysisExtendedIndustryAveragesAll, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CompanyAnalysisExtendedIndustryAverages) SetAll(v CompanyAnalysisExtendedIndustryAveragesAll)`

SetAll sets All field to given value.

### HasAll

`func (o *CompanyAnalysisExtendedIndustryAverages) HasAll() bool`

HasAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


