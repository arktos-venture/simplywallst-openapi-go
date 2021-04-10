# CompanyAnalysisExtendedAnalysisHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetIncome** | Pointer to **float64** |  | [optional] 
**NetInterestExpense** | Pointer to **float64** |  | [optional] 
**CashOperating** | Pointer to **float64** |  | [optional] 
**TotalAssets** | Pointer to **float64** |  | [optional] 
**Ppe** | Pointer to **float64** |  | [optional] 
**Inventory** | Pointer to **float64** |  | [optional] 
**CashStInvestments** | Pointer to **float64** |  | [optional] 
**Receivables** | Pointer to **float64** |  | [optional] 
**CurrentAssets** | Pointer to **float64** |  | [optional] 
**TotalEquity** | Pointer to **float64** |  | [optional] 
**TotalDebt** | Pointer to **float64** |  | [optional] 
**LongTermDebt** | Pointer to **float64** |  | [optional] 
**LongTermLiab** | Pointer to **float64** |  | [optional] 
**AccountsPayable** | Pointer to **float64** |  | [optional] 
**TotalCurrentLiab** | Pointer to **float64** |  | [optional] 
**TotalLiabEquity** | Pointer to **float64** |  | [optional] 
**TotalDebtEquity** | Pointer to **float64** |  | [optional] 
**TotalInventory** | Pointer to **float64** |  | [optional] 
**DebtToEquityRatio** | Pointer to **float64** |  | [optional] 
**DebtToEquityRatioPast** | Pointer to **float64** |  | [optional] 
**AssetsEquityRatio** | Pointer to **float64** |  | [optional] 
**FixedToTotalAssets** | Pointer to **float64** |  | [optional] 
**CurrentAssetsToTotalDebt** | Pointer to **float64** |  | [optional] 
**OperatingCashFlowToTotalDebt** | Pointer to **float64** |  | [optional] 
**NetInterestCover** | Pointer to **float64** |  | [optional] 
**CurrentSolvencyRatio** | Pointer to **float64** |  | [optional] 
**CurrentAssetsToLongTermLiab** | Pointer to **float64** |  | [optional] 
**OperatingExpenses** | Pointer to **float64** |  | [optional] 
**OperatingExpensesGrowthAnnual** | Pointer to **float64** |  | [optional] 
**OperatingExpensesStableYears** | Pointer to **float64** |  | [optional] 
**OperatingExpensesGrowthYears** | Pointer to **float64** |  | [optional] 
**Median2yrNetIncome** | Pointer to **float64** |  | [optional] 
**Capex** | Pointer to **float64** |  | [optional] 
**CapexGrowth1y** | Pointer to **float64** |  | [optional] 
**CapexGrowthAnnual** | Pointer to **float64** |  | [optional] 
**NetDebt** | Pointer to **float64** |  | [optional] 
**ManagementRateReturn** | Pointer to **float64** |  | [optional] 
**BookValuePerShare** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlow** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlow1y** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlow2y** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashGrowthAnnual** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlowGrowthAnnual** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlowGrowthYears** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlowStableYears** | Pointer to **float64** |  | [optional] 
**CashFromInvesting** | Pointer to **float64** |  | [optional] 
**NetDebtToEbitda** | Pointer to **float64** |  | [optional] 
**CapitalisationPercent** | Pointer to **float64** |  | [optional] 
**CapitalisationPercent1y** | Pointer to **float64** |  | [optional] 
**CashOperatingGrowth1y** | Pointer to **float64** |  | [optional] 
**CashFromInvesting1y** | Pointer to **float64** |  | [optional] 
**InventoryGrowth1y** | Pointer to **float64** |  | [optional] 
**NetOperatingAssets** | Pointer to **float64** |  | [optional] 
**NetOperatingAssets1y** | Pointer to **float64** |  | [optional] 
**AggregateAccruals** | Pointer to **float64** |  | [optional] 
**AccountsReceivableGrowth1y** | Pointer to **float64** |  | [optional] 
**AccountsReceivablePercent** | Pointer to **float64** |  | [optional] 
**LongTermAssets** | Pointer to **float64** |  | [optional] 
**LastBalanceSheetUpdate** | Pointer to **float64** |  | [optional] 
**CurrentPortionLeaseLiabilities** | Pointer to **float64** |  | [optional] 
**LongTermPortionLeaseLiabilities** | Pointer to **float64** |  | [optional] 
**TotalLeaseLiabilities** | Pointer to **float64** |  | [optional] 
**RestrictedCash** | Pointer to **float64** |  | [optional] 
**RestrictedCashRatio** | Pointer to **float64** |  | [optional] 
**LeveredFreeCashFlowBreakEvenYears** | Pointer to **float64** |  | [optional] 

## Methods

### NewCompanyAnalysisExtendedAnalysisHealth

`func NewCompanyAnalysisExtendedAnalysisHealth() *CompanyAnalysisExtendedAnalysisHealth`

NewCompanyAnalysisExtendedAnalysisHealth instantiates a new CompanyAnalysisExtendedAnalysisHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyAnalysisExtendedAnalysisHealthWithDefaults

`func NewCompanyAnalysisExtendedAnalysisHealthWithDefaults() *CompanyAnalysisExtendedAnalysisHealth`

NewCompanyAnalysisExtendedAnalysisHealthWithDefaults instantiates a new CompanyAnalysisExtendedAnalysisHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetIncome() float64`

GetNetIncome returns the NetIncome field if non-nil, zero value otherwise.

### GetNetIncomeOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetIncomeOk() (*float64, bool)`

GetNetIncomeOk returns a tuple with the NetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetIncome(v float64)`

SetNetIncome sets NetIncome field to given value.

### HasNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetIncome() bool`

HasNetIncome returns a boolean if a field has been set.

### GetNetInterestExpense

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetInterestExpense() float64`

GetNetInterestExpense returns the NetInterestExpense field if non-nil, zero value otherwise.

### GetNetInterestExpenseOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetInterestExpenseOk() (*float64, bool)`

GetNetInterestExpenseOk returns a tuple with the NetInterestExpense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterestExpense

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetInterestExpense(v float64)`

SetNetInterestExpense sets NetInterestExpense field to given value.

### HasNetInterestExpense

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetInterestExpense() bool`

HasNetInterestExpense returns a boolean if a field has been set.

### GetCashOperating

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashOperating() float64`

GetCashOperating returns the CashOperating field if non-nil, zero value otherwise.

### GetCashOperatingOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashOperatingOk() (*float64, bool)`

GetCashOperatingOk returns a tuple with the CashOperating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOperating

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCashOperating(v float64)`

SetCashOperating sets CashOperating field to given value.

### HasCashOperating

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCashOperating() bool`

HasCashOperating returns a boolean if a field has been set.

### GetTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalAssets() float64`

GetTotalAssets returns the TotalAssets field if non-nil, zero value otherwise.

### GetTotalAssetsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalAssetsOk() (*float64, bool)`

GetTotalAssetsOk returns a tuple with the TotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalAssets(v float64)`

SetTotalAssets sets TotalAssets field to given value.

### HasTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalAssets() bool`

HasTotalAssets returns a boolean if a field has been set.

### GetPpe

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetPpe() float64`

GetPpe returns the Ppe field if non-nil, zero value otherwise.

### GetPpeOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetPpeOk() (*float64, bool)`

GetPpeOk returns a tuple with the Ppe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpe

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetPpe(v float64)`

SetPpe sets Ppe field to given value.

### HasPpe

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasPpe() bool`

HasPpe returns a boolean if a field has been set.

### GetInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetInventory() float64`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetInventoryOk() (*float64, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetInventory(v float64)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetCashStInvestments

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashStInvestments() float64`

GetCashStInvestments returns the CashStInvestments field if non-nil, zero value otherwise.

### GetCashStInvestmentsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashStInvestmentsOk() (*float64, bool)`

GetCashStInvestmentsOk returns a tuple with the CashStInvestments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashStInvestments

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCashStInvestments(v float64)`

SetCashStInvestments sets CashStInvestments field to given value.

### HasCashStInvestments

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCashStInvestments() bool`

HasCashStInvestments returns a boolean if a field has been set.

### GetReceivables

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetReceivables() float64`

GetReceivables returns the Receivables field if non-nil, zero value otherwise.

### GetReceivablesOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetReceivablesOk() (*float64, bool)`

GetReceivablesOk returns a tuple with the Receivables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivables

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetReceivables(v float64)`

SetReceivables sets Receivables field to given value.

### HasReceivables

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasReceivables() bool`

HasReceivables returns a boolean if a field has been set.

### GetCurrentAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssets() float64`

GetCurrentAssets returns the CurrentAssets field if non-nil, zero value otherwise.

### GetCurrentAssetsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssetsOk() (*float64, bool)`

GetCurrentAssetsOk returns a tuple with the CurrentAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCurrentAssets(v float64)`

SetCurrentAssets sets CurrentAssets field to given value.

### HasCurrentAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCurrentAssets() bool`

HasCurrentAssets returns a boolean if a field has been set.

### GetTotalEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalEquity() float64`

GetTotalEquity returns the TotalEquity field if non-nil, zero value otherwise.

### GetTotalEquityOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalEquityOk() (*float64, bool)`

GetTotalEquityOk returns a tuple with the TotalEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalEquity(v float64)`

SetTotalEquity sets TotalEquity field to given value.

### HasTotalEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalEquity() bool`

HasTotalEquity returns a boolean if a field has been set.

### GetTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalDebt() float64`

GetTotalDebt returns the TotalDebt field if non-nil, zero value otherwise.

### GetTotalDebtOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalDebtOk() (*float64, bool)`

GetTotalDebtOk returns a tuple with the TotalDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalDebt(v float64)`

SetTotalDebt sets TotalDebt field to given value.

### HasTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalDebt() bool`

HasTotalDebt returns a boolean if a field has been set.

### GetLongTermDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermDebt() float64`

GetLongTermDebt returns the LongTermDebt field if non-nil, zero value otherwise.

### GetLongTermDebtOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermDebtOk() (*float64, bool)`

GetLongTermDebtOk returns a tuple with the LongTermDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLongTermDebt(v float64)`

SetLongTermDebt sets LongTermDebt field to given value.

### HasLongTermDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLongTermDebt() bool`

HasLongTermDebt returns a boolean if a field has been set.

### GetLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermLiab() float64`

GetLongTermLiab returns the LongTermLiab field if non-nil, zero value otherwise.

### GetLongTermLiabOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermLiabOk() (*float64, bool)`

GetLongTermLiabOk returns a tuple with the LongTermLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLongTermLiab(v float64)`

SetLongTermLiab sets LongTermLiab field to given value.

### HasLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLongTermLiab() bool`

HasLongTermLiab returns a boolean if a field has been set.

### GetAccountsPayable

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsPayable() float64`

GetAccountsPayable returns the AccountsPayable field if non-nil, zero value otherwise.

### GetAccountsPayableOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsPayableOk() (*float64, bool)`

GetAccountsPayableOk returns a tuple with the AccountsPayable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsPayable

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetAccountsPayable(v float64)`

SetAccountsPayable sets AccountsPayable field to given value.

### HasAccountsPayable

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasAccountsPayable() bool`

HasAccountsPayable returns a boolean if a field has been set.

### GetTotalCurrentLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalCurrentLiab() float64`

GetTotalCurrentLiab returns the TotalCurrentLiab field if non-nil, zero value otherwise.

### GetTotalCurrentLiabOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalCurrentLiabOk() (*float64, bool)`

GetTotalCurrentLiabOk returns a tuple with the TotalCurrentLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCurrentLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalCurrentLiab(v float64)`

SetTotalCurrentLiab sets TotalCurrentLiab field to given value.

### HasTotalCurrentLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalCurrentLiab() bool`

HasTotalCurrentLiab returns a boolean if a field has been set.

### GetTotalLiabEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalLiabEquity() float64`

GetTotalLiabEquity returns the TotalLiabEquity field if non-nil, zero value otherwise.

### GetTotalLiabEquityOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalLiabEquityOk() (*float64, bool)`

GetTotalLiabEquityOk returns a tuple with the TotalLiabEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLiabEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalLiabEquity(v float64)`

SetTotalLiabEquity sets TotalLiabEquity field to given value.

### HasTotalLiabEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalLiabEquity() bool`

HasTotalLiabEquity returns a boolean if a field has been set.

### GetTotalDebtEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalDebtEquity() float64`

GetTotalDebtEquity returns the TotalDebtEquity field if non-nil, zero value otherwise.

### GetTotalDebtEquityOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalDebtEquityOk() (*float64, bool)`

GetTotalDebtEquityOk returns a tuple with the TotalDebtEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDebtEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalDebtEquity(v float64)`

SetTotalDebtEquity sets TotalDebtEquity field to given value.

### HasTotalDebtEquity

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalDebtEquity() bool`

HasTotalDebtEquity returns a boolean if a field has been set.

### GetTotalInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalInventory() float64`

GetTotalInventory returns the TotalInventory field if non-nil, zero value otherwise.

### GetTotalInventoryOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalInventoryOk() (*float64, bool)`

GetTotalInventoryOk returns a tuple with the TotalInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalInventory(v float64)`

SetTotalInventory sets TotalInventory field to given value.

### HasTotalInventory

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalInventory() bool`

HasTotalInventory returns a boolean if a field has been set.

### GetDebtToEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetDebtToEquityRatio() float64`

GetDebtToEquityRatio returns the DebtToEquityRatio field if non-nil, zero value otherwise.

### GetDebtToEquityRatioOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetDebtToEquityRatioOk() (*float64, bool)`

GetDebtToEquityRatioOk returns a tuple with the DebtToEquityRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtToEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetDebtToEquityRatio(v float64)`

SetDebtToEquityRatio sets DebtToEquityRatio field to given value.

### HasDebtToEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasDebtToEquityRatio() bool`

HasDebtToEquityRatio returns a boolean if a field has been set.

### GetDebtToEquityRatioPast

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetDebtToEquityRatioPast() float64`

GetDebtToEquityRatioPast returns the DebtToEquityRatioPast field if non-nil, zero value otherwise.

### GetDebtToEquityRatioPastOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetDebtToEquityRatioPastOk() (*float64, bool)`

GetDebtToEquityRatioPastOk returns a tuple with the DebtToEquityRatioPast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebtToEquityRatioPast

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetDebtToEquityRatioPast(v float64)`

SetDebtToEquityRatioPast sets DebtToEquityRatioPast field to given value.

### HasDebtToEquityRatioPast

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasDebtToEquityRatioPast() bool`

HasDebtToEquityRatioPast returns a boolean if a field has been set.

### GetAssetsEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAssetsEquityRatio() float64`

GetAssetsEquityRatio returns the AssetsEquityRatio field if non-nil, zero value otherwise.

### GetAssetsEquityRatioOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAssetsEquityRatioOk() (*float64, bool)`

GetAssetsEquityRatioOk returns a tuple with the AssetsEquityRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetAssetsEquityRatio(v float64)`

SetAssetsEquityRatio sets AssetsEquityRatio field to given value.

### HasAssetsEquityRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasAssetsEquityRatio() bool`

HasAssetsEquityRatio returns a boolean if a field has been set.

### GetFixedToTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetFixedToTotalAssets() float64`

GetFixedToTotalAssets returns the FixedToTotalAssets field if non-nil, zero value otherwise.

### GetFixedToTotalAssetsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetFixedToTotalAssetsOk() (*float64, bool)`

GetFixedToTotalAssetsOk returns a tuple with the FixedToTotalAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedToTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetFixedToTotalAssets(v float64)`

SetFixedToTotalAssets sets FixedToTotalAssets field to given value.

### HasFixedToTotalAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasFixedToTotalAssets() bool`

HasFixedToTotalAssets returns a boolean if a field has been set.

### GetCurrentAssetsToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssetsToTotalDebt() float64`

GetCurrentAssetsToTotalDebt returns the CurrentAssetsToTotalDebt field if non-nil, zero value otherwise.

### GetCurrentAssetsToTotalDebtOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssetsToTotalDebtOk() (*float64, bool)`

GetCurrentAssetsToTotalDebtOk returns a tuple with the CurrentAssetsToTotalDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAssetsToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCurrentAssetsToTotalDebt(v float64)`

SetCurrentAssetsToTotalDebt sets CurrentAssetsToTotalDebt field to given value.

### HasCurrentAssetsToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCurrentAssetsToTotalDebt() bool`

HasCurrentAssetsToTotalDebt returns a boolean if a field has been set.

### GetOperatingCashFlowToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingCashFlowToTotalDebt() float64`

GetOperatingCashFlowToTotalDebt returns the OperatingCashFlowToTotalDebt field if non-nil, zero value otherwise.

### GetOperatingCashFlowToTotalDebtOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingCashFlowToTotalDebtOk() (*float64, bool)`

GetOperatingCashFlowToTotalDebtOk returns a tuple with the OperatingCashFlowToTotalDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingCashFlowToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetOperatingCashFlowToTotalDebt(v float64)`

SetOperatingCashFlowToTotalDebt sets OperatingCashFlowToTotalDebt field to given value.

### HasOperatingCashFlowToTotalDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasOperatingCashFlowToTotalDebt() bool`

HasOperatingCashFlowToTotalDebt returns a boolean if a field has been set.

### GetNetInterestCover

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetInterestCover() float64`

GetNetInterestCover returns the NetInterestCover field if non-nil, zero value otherwise.

### GetNetInterestCoverOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetInterestCoverOk() (*float64, bool)`

GetNetInterestCoverOk returns a tuple with the NetInterestCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetInterestCover

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetInterestCover(v float64)`

SetNetInterestCover sets NetInterestCover field to given value.

### HasNetInterestCover

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetInterestCover() bool`

HasNetInterestCover returns a boolean if a field has been set.

### GetCurrentSolvencyRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentSolvencyRatio() float64`

GetCurrentSolvencyRatio returns the CurrentSolvencyRatio field if non-nil, zero value otherwise.

### GetCurrentSolvencyRatioOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentSolvencyRatioOk() (*float64, bool)`

GetCurrentSolvencyRatioOk returns a tuple with the CurrentSolvencyRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSolvencyRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCurrentSolvencyRatio(v float64)`

SetCurrentSolvencyRatio sets CurrentSolvencyRatio field to given value.

### HasCurrentSolvencyRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCurrentSolvencyRatio() bool`

HasCurrentSolvencyRatio returns a boolean if a field has been set.

### GetCurrentAssetsToLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssetsToLongTermLiab() float64`

GetCurrentAssetsToLongTermLiab returns the CurrentAssetsToLongTermLiab field if non-nil, zero value otherwise.

### GetCurrentAssetsToLongTermLiabOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentAssetsToLongTermLiabOk() (*float64, bool)`

GetCurrentAssetsToLongTermLiabOk returns a tuple with the CurrentAssetsToLongTermLiab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAssetsToLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCurrentAssetsToLongTermLiab(v float64)`

SetCurrentAssetsToLongTermLiab sets CurrentAssetsToLongTermLiab field to given value.

### HasCurrentAssetsToLongTermLiab

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCurrentAssetsToLongTermLiab() bool`

HasCurrentAssetsToLongTermLiab returns a boolean if a field has been set.

### GetOperatingExpenses

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpenses() float64`

GetOperatingExpenses returns the OperatingExpenses field if non-nil, zero value otherwise.

### GetOperatingExpensesOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesOk() (*float64, bool)`

GetOperatingExpensesOk returns a tuple with the OperatingExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpenses

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetOperatingExpenses(v float64)`

SetOperatingExpenses sets OperatingExpenses field to given value.

### HasOperatingExpenses

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasOperatingExpenses() bool`

HasOperatingExpenses returns a boolean if a field has been set.

### GetOperatingExpensesGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesGrowthAnnual() float64`

GetOperatingExpensesGrowthAnnual returns the OperatingExpensesGrowthAnnual field if non-nil, zero value otherwise.

### GetOperatingExpensesGrowthAnnualOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesGrowthAnnualOk() (*float64, bool)`

GetOperatingExpensesGrowthAnnualOk returns a tuple with the OperatingExpensesGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpensesGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetOperatingExpensesGrowthAnnual(v float64)`

SetOperatingExpensesGrowthAnnual sets OperatingExpensesGrowthAnnual field to given value.

### HasOperatingExpensesGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasOperatingExpensesGrowthAnnual() bool`

HasOperatingExpensesGrowthAnnual returns a boolean if a field has been set.

### GetOperatingExpensesStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesStableYears() float64`

GetOperatingExpensesStableYears returns the OperatingExpensesStableYears field if non-nil, zero value otherwise.

### GetOperatingExpensesStableYearsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesStableYearsOk() (*float64, bool)`

GetOperatingExpensesStableYearsOk returns a tuple with the OperatingExpensesStableYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpensesStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetOperatingExpensesStableYears(v float64)`

SetOperatingExpensesStableYears sets OperatingExpensesStableYears field to given value.

### HasOperatingExpensesStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasOperatingExpensesStableYears() bool`

HasOperatingExpensesStableYears returns a boolean if a field has been set.

### GetOperatingExpensesGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesGrowthYears() float64`

GetOperatingExpensesGrowthYears returns the OperatingExpensesGrowthYears field if non-nil, zero value otherwise.

### GetOperatingExpensesGrowthYearsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetOperatingExpensesGrowthYearsOk() (*float64, bool)`

GetOperatingExpensesGrowthYearsOk returns a tuple with the OperatingExpensesGrowthYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingExpensesGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetOperatingExpensesGrowthYears(v float64)`

SetOperatingExpensesGrowthYears sets OperatingExpensesGrowthYears field to given value.

### HasOperatingExpensesGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasOperatingExpensesGrowthYears() bool`

HasOperatingExpensesGrowthYears returns a boolean if a field has been set.

### GetMedian2yrNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetMedian2yrNetIncome() float64`

GetMedian2yrNetIncome returns the Median2yrNetIncome field if non-nil, zero value otherwise.

### GetMedian2yrNetIncomeOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetMedian2yrNetIncomeOk() (*float64, bool)`

GetMedian2yrNetIncomeOk returns a tuple with the Median2yrNetIncome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedian2yrNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetMedian2yrNetIncome(v float64)`

SetMedian2yrNetIncome sets Median2yrNetIncome field to given value.

### HasMedian2yrNetIncome

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasMedian2yrNetIncome() bool`

HasMedian2yrNetIncome returns a boolean if a field has been set.

### GetCapex

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapex() float64`

GetCapex returns the Capex field if non-nil, zero value otherwise.

### GetCapexOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapexOk() (*float64, bool)`

GetCapexOk returns a tuple with the Capex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapex

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCapex(v float64)`

SetCapex sets Capex field to given value.

### HasCapex

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCapex() bool`

HasCapex returns a boolean if a field has been set.

### GetCapexGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapexGrowth1y() float64`

GetCapexGrowth1y returns the CapexGrowth1y field if non-nil, zero value otherwise.

### GetCapexGrowth1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapexGrowth1yOk() (*float64, bool)`

GetCapexGrowth1yOk returns a tuple with the CapexGrowth1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapexGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCapexGrowth1y(v float64)`

SetCapexGrowth1y sets CapexGrowth1y field to given value.

### HasCapexGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCapexGrowth1y() bool`

HasCapexGrowth1y returns a boolean if a field has been set.

### GetCapexGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapexGrowthAnnual() float64`

GetCapexGrowthAnnual returns the CapexGrowthAnnual field if non-nil, zero value otherwise.

### GetCapexGrowthAnnualOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapexGrowthAnnualOk() (*float64, bool)`

GetCapexGrowthAnnualOk returns a tuple with the CapexGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapexGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCapexGrowthAnnual(v float64)`

SetCapexGrowthAnnual sets CapexGrowthAnnual field to given value.

### HasCapexGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCapexGrowthAnnual() bool`

HasCapexGrowthAnnual returns a boolean if a field has been set.

### GetNetDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetDebt() float64`

GetNetDebt returns the NetDebt field if non-nil, zero value otherwise.

### GetNetDebtOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetDebtOk() (*float64, bool)`

GetNetDebtOk returns a tuple with the NetDebt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetDebt(v float64)`

SetNetDebt sets NetDebt field to given value.

### HasNetDebt

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetDebt() bool`

HasNetDebt returns a boolean if a field has been set.

### GetManagementRateReturn

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetManagementRateReturn() float64`

GetManagementRateReturn returns the ManagementRateReturn field if non-nil, zero value otherwise.

### GetManagementRateReturnOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetManagementRateReturnOk() (*float64, bool)`

GetManagementRateReturnOk returns a tuple with the ManagementRateReturn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementRateReturn

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetManagementRateReturn(v float64)`

SetManagementRateReturn sets ManagementRateReturn field to given value.

### HasManagementRateReturn

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasManagementRateReturn() bool`

HasManagementRateReturn returns a boolean if a field has been set.

### GetBookValuePerShare

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetBookValuePerShare() float64`

GetBookValuePerShare returns the BookValuePerShare field if non-nil, zero value otherwise.

### GetBookValuePerShareOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetBookValuePerShareOk() (*float64, bool)`

GetBookValuePerShareOk returns a tuple with the BookValuePerShare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookValuePerShare

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetBookValuePerShare(v float64)`

SetBookValuePerShare sets BookValuePerShare field to given value.

### HasBookValuePerShare

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasBookValuePerShare() bool`

HasBookValuePerShare returns a boolean if a field has been set.

### GetLeveredFreeCashFlow

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlow() float64`

GetLeveredFreeCashFlow returns the LeveredFreeCashFlow field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlowOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowOk() (*float64, bool)`

GetLeveredFreeCashFlowOk returns a tuple with the LeveredFreeCashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlow

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlow(v float64)`

SetLeveredFreeCashFlow sets LeveredFreeCashFlow field to given value.

### HasLeveredFreeCashFlow

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlow() bool`

HasLeveredFreeCashFlow returns a boolean if a field has been set.

### GetLeveredFreeCashFlow1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlow1y() float64`

GetLeveredFreeCashFlow1y returns the LeveredFreeCashFlow1y field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlow1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlow1yOk() (*float64, bool)`

GetLeveredFreeCashFlow1yOk returns a tuple with the LeveredFreeCashFlow1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlow1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlow1y(v float64)`

SetLeveredFreeCashFlow1y sets LeveredFreeCashFlow1y field to given value.

### HasLeveredFreeCashFlow1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlow1y() bool`

HasLeveredFreeCashFlow1y returns a boolean if a field has been set.

### GetLeveredFreeCashFlow2y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlow2y() float64`

GetLeveredFreeCashFlow2y returns the LeveredFreeCashFlow2y field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlow2yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlow2yOk() (*float64, bool)`

GetLeveredFreeCashFlow2yOk returns a tuple with the LeveredFreeCashFlow2y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlow2y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlow2y(v float64)`

SetLeveredFreeCashFlow2y sets LeveredFreeCashFlow2y field to given value.

### HasLeveredFreeCashFlow2y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlow2y() bool`

HasLeveredFreeCashFlow2y returns a boolean if a field has been set.

### GetLeveredFreeCashGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashGrowthAnnual() float64`

GetLeveredFreeCashGrowthAnnual returns the LeveredFreeCashGrowthAnnual field if non-nil, zero value otherwise.

### GetLeveredFreeCashGrowthAnnualOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashGrowthAnnualOk() (*float64, bool)`

GetLeveredFreeCashGrowthAnnualOk returns a tuple with the LeveredFreeCashGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashGrowthAnnual(v float64)`

SetLeveredFreeCashGrowthAnnual sets LeveredFreeCashGrowthAnnual field to given value.

### HasLeveredFreeCashGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashGrowthAnnual() bool`

HasLeveredFreeCashGrowthAnnual returns a boolean if a field has been set.

### GetLeveredFreeCashFlowGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowGrowthAnnual() float64`

GetLeveredFreeCashFlowGrowthAnnual returns the LeveredFreeCashFlowGrowthAnnual field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlowGrowthAnnualOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowGrowthAnnualOk() (*float64, bool)`

GetLeveredFreeCashFlowGrowthAnnualOk returns a tuple with the LeveredFreeCashFlowGrowthAnnual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlowGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlowGrowthAnnual(v float64)`

SetLeveredFreeCashFlowGrowthAnnual sets LeveredFreeCashFlowGrowthAnnual field to given value.

### HasLeveredFreeCashFlowGrowthAnnual

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlowGrowthAnnual() bool`

HasLeveredFreeCashFlowGrowthAnnual returns a boolean if a field has been set.

### GetLeveredFreeCashFlowGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowGrowthYears() float64`

GetLeveredFreeCashFlowGrowthYears returns the LeveredFreeCashFlowGrowthYears field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlowGrowthYearsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowGrowthYearsOk() (*float64, bool)`

GetLeveredFreeCashFlowGrowthYearsOk returns a tuple with the LeveredFreeCashFlowGrowthYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlowGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlowGrowthYears(v float64)`

SetLeveredFreeCashFlowGrowthYears sets LeveredFreeCashFlowGrowthYears field to given value.

### HasLeveredFreeCashFlowGrowthYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlowGrowthYears() bool`

HasLeveredFreeCashFlowGrowthYears returns a boolean if a field has been set.

### GetLeveredFreeCashFlowStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowStableYears() float64`

GetLeveredFreeCashFlowStableYears returns the LeveredFreeCashFlowStableYears field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlowStableYearsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowStableYearsOk() (*float64, bool)`

GetLeveredFreeCashFlowStableYearsOk returns a tuple with the LeveredFreeCashFlowStableYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlowStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlowStableYears(v float64)`

SetLeveredFreeCashFlowStableYears sets LeveredFreeCashFlowStableYears field to given value.

### HasLeveredFreeCashFlowStableYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlowStableYears() bool`

HasLeveredFreeCashFlowStableYears returns a boolean if a field has been set.

### GetCashFromInvesting

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashFromInvesting() float64`

GetCashFromInvesting returns the CashFromInvesting field if non-nil, zero value otherwise.

### GetCashFromInvestingOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashFromInvestingOk() (*float64, bool)`

GetCashFromInvestingOk returns a tuple with the CashFromInvesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromInvesting

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCashFromInvesting(v float64)`

SetCashFromInvesting sets CashFromInvesting field to given value.

### HasCashFromInvesting

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCashFromInvesting() bool`

HasCashFromInvesting returns a boolean if a field has been set.

### GetNetDebtToEbitda

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetDebtToEbitda() float64`

GetNetDebtToEbitda returns the NetDebtToEbitda field if non-nil, zero value otherwise.

### GetNetDebtToEbitdaOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetDebtToEbitdaOk() (*float64, bool)`

GetNetDebtToEbitdaOk returns a tuple with the NetDebtToEbitda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetDebtToEbitda

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetDebtToEbitda(v float64)`

SetNetDebtToEbitda sets NetDebtToEbitda field to given value.

### HasNetDebtToEbitda

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetDebtToEbitda() bool`

HasNetDebtToEbitda returns a boolean if a field has been set.

### GetCapitalisationPercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapitalisationPercent() float64`

GetCapitalisationPercent returns the CapitalisationPercent field if non-nil, zero value otherwise.

### GetCapitalisationPercentOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapitalisationPercentOk() (*float64, bool)`

GetCapitalisationPercentOk returns a tuple with the CapitalisationPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalisationPercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCapitalisationPercent(v float64)`

SetCapitalisationPercent sets CapitalisationPercent field to given value.

### HasCapitalisationPercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCapitalisationPercent() bool`

HasCapitalisationPercent returns a boolean if a field has been set.

### GetCapitalisationPercent1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapitalisationPercent1y() float64`

GetCapitalisationPercent1y returns the CapitalisationPercent1y field if non-nil, zero value otherwise.

### GetCapitalisationPercent1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCapitalisationPercent1yOk() (*float64, bool)`

GetCapitalisationPercent1yOk returns a tuple with the CapitalisationPercent1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapitalisationPercent1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCapitalisationPercent1y(v float64)`

SetCapitalisationPercent1y sets CapitalisationPercent1y field to given value.

### HasCapitalisationPercent1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCapitalisationPercent1y() bool`

HasCapitalisationPercent1y returns a boolean if a field has been set.

### GetCashOperatingGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashOperatingGrowth1y() float64`

GetCashOperatingGrowth1y returns the CashOperatingGrowth1y field if non-nil, zero value otherwise.

### GetCashOperatingGrowth1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashOperatingGrowth1yOk() (*float64, bool)`

GetCashOperatingGrowth1yOk returns a tuple with the CashOperatingGrowth1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashOperatingGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCashOperatingGrowth1y(v float64)`

SetCashOperatingGrowth1y sets CashOperatingGrowth1y field to given value.

### HasCashOperatingGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCashOperatingGrowth1y() bool`

HasCashOperatingGrowth1y returns a boolean if a field has been set.

### GetCashFromInvesting1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashFromInvesting1y() float64`

GetCashFromInvesting1y returns the CashFromInvesting1y field if non-nil, zero value otherwise.

### GetCashFromInvesting1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCashFromInvesting1yOk() (*float64, bool)`

GetCashFromInvesting1yOk returns a tuple with the CashFromInvesting1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFromInvesting1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCashFromInvesting1y(v float64)`

SetCashFromInvesting1y sets CashFromInvesting1y field to given value.

### HasCashFromInvesting1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCashFromInvesting1y() bool`

HasCashFromInvesting1y returns a boolean if a field has been set.

### GetInventoryGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetInventoryGrowth1y() float64`

GetInventoryGrowth1y returns the InventoryGrowth1y field if non-nil, zero value otherwise.

### GetInventoryGrowth1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetInventoryGrowth1yOk() (*float64, bool)`

GetInventoryGrowth1yOk returns a tuple with the InventoryGrowth1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetInventoryGrowth1y(v float64)`

SetInventoryGrowth1y sets InventoryGrowth1y field to given value.

### HasInventoryGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasInventoryGrowth1y() bool`

HasInventoryGrowth1y returns a boolean if a field has been set.

### GetNetOperatingAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetOperatingAssets() float64`

GetNetOperatingAssets returns the NetOperatingAssets field if non-nil, zero value otherwise.

### GetNetOperatingAssetsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetOperatingAssetsOk() (*float64, bool)`

GetNetOperatingAssetsOk returns a tuple with the NetOperatingAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOperatingAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetOperatingAssets(v float64)`

SetNetOperatingAssets sets NetOperatingAssets field to given value.

### HasNetOperatingAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetOperatingAssets() bool`

HasNetOperatingAssets returns a boolean if a field has been set.

### GetNetOperatingAssets1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetOperatingAssets1y() float64`

GetNetOperatingAssets1y returns the NetOperatingAssets1y field if non-nil, zero value otherwise.

### GetNetOperatingAssets1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetNetOperatingAssets1yOk() (*float64, bool)`

GetNetOperatingAssets1yOk returns a tuple with the NetOperatingAssets1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOperatingAssets1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetNetOperatingAssets1y(v float64)`

SetNetOperatingAssets1y sets NetOperatingAssets1y field to given value.

### HasNetOperatingAssets1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasNetOperatingAssets1y() bool`

HasNetOperatingAssets1y returns a boolean if a field has been set.

### GetAggregateAccruals

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAggregateAccruals() float64`

GetAggregateAccruals returns the AggregateAccruals field if non-nil, zero value otherwise.

### GetAggregateAccrualsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAggregateAccrualsOk() (*float64, bool)`

GetAggregateAccrualsOk returns a tuple with the AggregateAccruals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateAccruals

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetAggregateAccruals(v float64)`

SetAggregateAccruals sets AggregateAccruals field to given value.

### HasAggregateAccruals

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasAggregateAccruals() bool`

HasAggregateAccruals returns a boolean if a field has been set.

### GetAccountsReceivableGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsReceivableGrowth1y() float64`

GetAccountsReceivableGrowth1y returns the AccountsReceivableGrowth1y field if non-nil, zero value otherwise.

### GetAccountsReceivableGrowth1yOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsReceivableGrowth1yOk() (*float64, bool)`

GetAccountsReceivableGrowth1yOk returns a tuple with the AccountsReceivableGrowth1y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsReceivableGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetAccountsReceivableGrowth1y(v float64)`

SetAccountsReceivableGrowth1y sets AccountsReceivableGrowth1y field to given value.

### HasAccountsReceivableGrowth1y

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasAccountsReceivableGrowth1y() bool`

HasAccountsReceivableGrowth1y returns a boolean if a field has been set.

### GetAccountsReceivablePercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsReceivablePercent() float64`

GetAccountsReceivablePercent returns the AccountsReceivablePercent field if non-nil, zero value otherwise.

### GetAccountsReceivablePercentOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetAccountsReceivablePercentOk() (*float64, bool)`

GetAccountsReceivablePercentOk returns a tuple with the AccountsReceivablePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsReceivablePercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetAccountsReceivablePercent(v float64)`

SetAccountsReceivablePercent sets AccountsReceivablePercent field to given value.

### HasAccountsReceivablePercent

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasAccountsReceivablePercent() bool`

HasAccountsReceivablePercent returns a boolean if a field has been set.

### GetLongTermAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermAssets() float64`

GetLongTermAssets returns the LongTermAssets field if non-nil, zero value otherwise.

### GetLongTermAssetsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermAssetsOk() (*float64, bool)`

GetLongTermAssetsOk returns a tuple with the LongTermAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLongTermAssets(v float64)`

SetLongTermAssets sets LongTermAssets field to given value.

### HasLongTermAssets

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLongTermAssets() bool`

HasLongTermAssets returns a boolean if a field has been set.

### GetLastBalanceSheetUpdate

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLastBalanceSheetUpdate() float64`

GetLastBalanceSheetUpdate returns the LastBalanceSheetUpdate field if non-nil, zero value otherwise.

### GetLastBalanceSheetUpdateOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLastBalanceSheetUpdateOk() (*float64, bool)`

GetLastBalanceSheetUpdateOk returns a tuple with the LastBalanceSheetUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBalanceSheetUpdate

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLastBalanceSheetUpdate(v float64)`

SetLastBalanceSheetUpdate sets LastBalanceSheetUpdate field to given value.

### HasLastBalanceSheetUpdate

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLastBalanceSheetUpdate() bool`

HasLastBalanceSheetUpdate returns a boolean if a field has been set.

### GetCurrentPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentPortionLeaseLiabilities() float64`

GetCurrentPortionLeaseLiabilities returns the CurrentPortionLeaseLiabilities field if non-nil, zero value otherwise.

### GetCurrentPortionLeaseLiabilitiesOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetCurrentPortionLeaseLiabilitiesOk() (*float64, bool)`

GetCurrentPortionLeaseLiabilitiesOk returns a tuple with the CurrentPortionLeaseLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetCurrentPortionLeaseLiabilities(v float64)`

SetCurrentPortionLeaseLiabilities sets CurrentPortionLeaseLiabilities field to given value.

### HasCurrentPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasCurrentPortionLeaseLiabilities() bool`

HasCurrentPortionLeaseLiabilities returns a boolean if a field has been set.

### GetLongTermPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermPortionLeaseLiabilities() float64`

GetLongTermPortionLeaseLiabilities returns the LongTermPortionLeaseLiabilities field if non-nil, zero value otherwise.

### GetLongTermPortionLeaseLiabilitiesOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLongTermPortionLeaseLiabilitiesOk() (*float64, bool)`

GetLongTermPortionLeaseLiabilitiesOk returns a tuple with the LongTermPortionLeaseLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongTermPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLongTermPortionLeaseLiabilities(v float64)`

SetLongTermPortionLeaseLiabilities sets LongTermPortionLeaseLiabilities field to given value.

### HasLongTermPortionLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLongTermPortionLeaseLiabilities() bool`

HasLongTermPortionLeaseLiabilities returns a boolean if a field has been set.

### GetTotalLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalLeaseLiabilities() float64`

GetTotalLeaseLiabilities returns the TotalLeaseLiabilities field if non-nil, zero value otherwise.

### GetTotalLeaseLiabilitiesOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetTotalLeaseLiabilitiesOk() (*float64, bool)`

GetTotalLeaseLiabilitiesOk returns a tuple with the TotalLeaseLiabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetTotalLeaseLiabilities(v float64)`

SetTotalLeaseLiabilities sets TotalLeaseLiabilities field to given value.

### HasTotalLeaseLiabilities

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasTotalLeaseLiabilities() bool`

HasTotalLeaseLiabilities returns a boolean if a field has been set.

### GetRestrictedCash

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetRestrictedCash() float64`

GetRestrictedCash returns the RestrictedCash field if non-nil, zero value otherwise.

### GetRestrictedCashOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetRestrictedCashOk() (*float64, bool)`

GetRestrictedCashOk returns a tuple with the RestrictedCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCash

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetRestrictedCash(v float64)`

SetRestrictedCash sets RestrictedCash field to given value.

### HasRestrictedCash

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasRestrictedCash() bool`

HasRestrictedCash returns a boolean if a field has been set.

### GetRestrictedCashRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetRestrictedCashRatio() float64`

GetRestrictedCashRatio returns the RestrictedCashRatio field if non-nil, zero value otherwise.

### GetRestrictedCashRatioOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetRestrictedCashRatioOk() (*float64, bool)`

GetRestrictedCashRatioOk returns a tuple with the RestrictedCashRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCashRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetRestrictedCashRatio(v float64)`

SetRestrictedCashRatio sets RestrictedCashRatio field to given value.

### HasRestrictedCashRatio

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasRestrictedCashRatio() bool`

HasRestrictedCashRatio returns a boolean if a field has been set.

### GetLeveredFreeCashFlowBreakEvenYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowBreakEvenYears() float64`

GetLeveredFreeCashFlowBreakEvenYears returns the LeveredFreeCashFlowBreakEvenYears field if non-nil, zero value otherwise.

### GetLeveredFreeCashFlowBreakEvenYearsOk

`func (o *CompanyAnalysisExtendedAnalysisHealth) GetLeveredFreeCashFlowBreakEvenYearsOk() (*float64, bool)`

GetLeveredFreeCashFlowBreakEvenYearsOk returns a tuple with the LeveredFreeCashFlowBreakEvenYears field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeveredFreeCashFlowBreakEvenYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) SetLeveredFreeCashFlowBreakEvenYears(v float64)`

SetLeveredFreeCashFlowBreakEvenYears sets LeveredFreeCashFlowBreakEvenYears field to given value.

### HasLeveredFreeCashFlowBreakEvenYears

`func (o *CompanyAnalysisExtendedAnalysisHealth) HasLeveredFreeCashFlowBreakEvenYears() bool`

HasLeveredFreeCashFlowBreakEvenYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


