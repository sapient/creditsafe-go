package creditsafe

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CompanyReportResponse struct {
	CorrelationID string `json:"correlationId"`
	Report        Report `json:"report"`
	CompanyID     string `json:"companyId"`
	DateOfOrder   string `json:"dateOfOrder"`
	Language      string `json:"language"`
	UserID        string `json:"userId"`
}

type Report struct {
	CompanyID                string                    `json:"companyId"`
	Language                 string                    `json:"language"`
	CompanySummary           CompanySummary            `json:"companySummary"`
	CompanyIdentification    CompanyIdentification     `json:"companyIdentification"`
	CreditScore              CreditScore               `json:"creditScore"`
	ContactInformation       ContactInformation        `json:"contactInformation"`
	ShareCapitalStructure    ShareCapitalStructure     `json:"shareCapitalStructure"`
	Directors                Directors                 `json:"directors"`
	OtherInformation         OtherInformation          `json:"otherInformation"`
	FinancialStatements      []FinancialStatement      `json:"financialStatements"`
	LocalFinancialStatements []LocalFinancialStatement `json:"localFinancialStatements"`
	PaymentData              PaymentData               `json:"paymentData"`
	NegativeInformation      NegativeInformation       `json:"negativeInformation"`
	AdditionalInformation    AdditionalInformation     `json:"additionalInformation"`
}

type AdditionalInformation struct {
	Miscellaneous                 Miscellaneous                 `json:"miscellaneous"`
	CompanyHistory                []CompanyHistory              `json:"companyHistory"`
	MortgageSummary               MortgageSummary               `json:"mortgageSummary"`
	Commentaries                  []Commentary                  `json:"commentaries"`
	PersonsWithSignificantControl PersonsWithSignificantControl `json:"personsWithSignificantControl"`
	RatingHistory                 []RatingHistory               `json:"ratingHistory"`
	CreditLimitHistory            []CreditLimitHistory          `json:"creditLimitHistory"`
	LandRegistry                  LandRegistry                  `json:"landRegistry"`
	EnquiriesTrend                map[string]int64              `json:"enquiriesTrend"`
	ScoreTiers                    []ScoreTier                   `json:"scoreTiers"`
}

type Commentary struct {
	CommentaryText     string `json:"commentaryText"`
	PositiveNegative   string `json:"positiveNegative"`
	PositiveOrNegative string `json:"positiveOrNegative"`
}

type CompanyHistory struct {
	Date        string `json:"date"`
	Description string `json:"description"`
}

type CreditLimitHistory struct {
	Date         string                         `json:"date"`
	CompanyValue LatestShareholdersEquityFigure `json:"companyValue"`
}

type LatestShareholdersEquityFigure struct {
	Currency Currency `json:"currency"`
	Value    int64    `json:"value"`
}

type LandRegistry struct {
	RegisteredLandAndProperty int64 `json:"registeredLandAndProperty"`
}

type Miscellaneous struct {
	FPS                  bool   `json:"fps"`
	FilingDateOfAccounts string `json:"filingDateOfAccounts"`
	OddsOfFailure        string `json:"oddsOfFailure"`
	AccountsDueDate      string `json:"accountsDueDate"`
}

type MortgageSummary struct {
	Outstanding int64 `json:"outstanding"`
	Satisfied   int64 `json:"satisfied"`
}

type PersonsWithSignificantControl struct {
	ActivePSC []ActivePSC `json:"activePSC"`
	CeasedPSC []CeasedPSC `json:"ceasedPSC"`
}

type ActivePSC struct {
	Title              string           `json:"title"`
	Name               string           `json:"name"`
	PersonType         string           `json:"personType"`
	Address            ActivePSCAddress `json:"address"`
	Country            string           `json:"country"`
	DateOfBirth        string           `json:"dateOfBirth"`
	Kind               string           `json:"kind"`
	CountryOfResidence string           `json:"countryOfResidence"`
	Nationality        string           `json:"nationality"`
	NatureOfControl    string           `json:"natureOfControl"`
	NotifiedOn         string           `json:"notifiedOn"`
	InsertDate         string           `json:"insertDate"`
}

type ActivePSCAddress struct {
	SimpleValue string  `json:"simpleValue"`
	Street      string  `json:"street"`
	HouseNumber *string `json:"houseNumber,omitempty"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postalCode"`
	Province    *string `json:"province,omitempty"`
}

type CeasedPSC struct {
	Kind       string    `json:"kind"`
	NotifiedOn string    `json:"notifiedOn"`
	CeasedOn   string    `json:"ceasedOn"`
	Statement  Statement `json:"statement"`
	InsertDate string    `json:"insertDate"`
}

type Statement struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type RatingHistory struct {
	Date              string      `json:"date"`
	CompanyValue      int64       `json:"companyValue"`
	RatingDescription Description `json:"ratingDescription"`
}

type ScoreTier struct {
	ActivityCode         string  `json:"activityCode"`
	NumberOfActivityCode int64   `json:"numberOfActivityCode"`
	Ranking              Ranking `json:"ranking"`
}

type Ranking struct {
	Code        int64  `json:"code"`
	Description string `json:"description"`
}

type CompanyIdentification struct {
	BasicInformation        BasicInformation         `json:"basicInformation"`
	ActivityClassifications []ActivityClassification `json:"activityClassifications"`
}

type ActivityClassification struct {
	Classification string      `json:"classification"`
	Activities     []Statement `json:"activities"`
}

type BasicInformation struct {
	BusinessName              string            `json:"businessName"`
	RegisteredCompanyName     string            `json:"registeredCompanyName"`
	CompanyRegistrationNumber string            `json:"companyRegistrationNumber"`
	GgsID                     string            `json:"ggsId"`
	Country                   string            `json:"country"`
	VatRegistrationNumber     string            `json:"vatRegistrationNumber"`
	CompanyRegistrationDate   string            `json:"companyRegistrationDate"`
	LegalForm                 LegalForm         `json:"legalForm"`
	CompanyStatus             CompanyStatus     `json:"companyStatus"`
	PrincipalActivity         PrincipalActivity `json:"principalActivity"`
	ContactAddress            Address           `json:"contactAddress"`
}

type CompanyStatus struct {
	Status         string `json:"status"`
	ProviderStatus string `json:"providerStatus"`
	Description    string `json:"description"`
}

type Address struct {
	Type                  string  `json:"type"`
	SimpleValue           string  `json:"simpleValue"`
	Street                string  `json:"street"`
	HouseNumber           string  `json:"houseNumber"`
	City                  string  `json:"city"`
	PostalCode            string  `json:"postalCode"`
	Province              string  `json:"province"`
	Telephone             string  `json:"telephone"`
	DirectMarketingOptOut bool    `json:"directMarketingOptOut"`
	Country               *string `json:"country,omitempty"`
}

type LegalForm struct {
	Description string `json:"description"`
}

type PrincipalActivity struct {
	Classification string `json:"classification"`
}

type CompanySummary struct {
	BusinessName                   string                         `json:"businessName"`
	Country                        string                         `json:"country"`
	CompanyNumber                  string                         `json:"companyNumber"`
	CompanyRegistrationNumber      string                         `json:"companyRegistrationNumber"`
	GgsID                          string                         `json:"ggsId"`
	MainActivity                   MainActivity                   `json:"mainActivity"`
	CompanyStatus                  CompanyStatus                  `json:"companyStatus"`
	LatestTurnoverFigure           LatestShareholdersEquityFigure `json:"latestTurnoverFigure"`
	LatestShareholdersEquityFigure LatestShareholdersEquityFigure `json:"latestShareholdersEquityFigure"`
	CreditRating                   CreditRating                   `json:"creditRating"`
}

type CreditRating struct {
	CommonValue         string        `json:"commonValue"`
	CommonDescription   Description   `json:"commonDescription"`
	CreditLimit         CreditLimit   `json:"creditLimit"`
	ProviderValue       ProviderValue `json:"providerValue"`
	ProviderDescription Description   `json:"providerDescription"`
	Pod                 *float64      `json:"pod,omitempty"`
}

type CreditLimit struct {
	Currency Currency `json:"currency"`
	Value    string   `json:"value"`
}

type ProviderValue struct {
	MaxValue string `json:"maxValue"`
	MinValue string `json:"minValue"`
	Value    string `json:"value"`
}

type MainActivity struct {
	Code           string `json:"code"`
	Description    string `json:"description"`
	Classification string `json:"classification"`
}

type ContactInformation struct {
	MainAddress    Address        `json:"mainAddress"`
	OtherAddresses []OtherAddress `json:"otherAddresses"`
	Websites       []string       `json:"websites"`
}

type OtherAddress struct {
	Type                  string `json:"type"`
	SimpleValue           string `json:"simpleValue"`
	PostalCode            string `json:"postalCode"`
	Telephone             string `json:"telephone"`
	DirectMarketingOptOut bool   `json:"directMarketingOptOut"`
	Country               string `json:"country"`
}

type CreditScore struct {
	CurrentCreditRating    CreditRating                   `json:"currentCreditRating"`
	CurrentContractLimit   LatestShareholdersEquityFigure `json:"currentContractLimit"`
	PreviousCreditRating   CreditRating                   `json:"previousCreditRating"`
	LatestRatingChangeDate string                         `json:"latestRatingChangeDate"`
}

type Directors struct {
	CurrentDirectors  []CurrentDirector  `json:"currentDirectors"`
	PreviousDirectors []PreviousDirector `json:"previousDirectors"`
}

type CurrentDirector struct {
	ID             string                        `json:"id"`
	IDType         string                        `json:"idType"`
	Name           string                        `json:"name"`
	Title          string                        `json:"title"`
	FirstName      string                        `json:"firstName"`
	MiddleName     *string                       `json:"middleName,omitempty"`
	Surname        string                        `json:"surname"`
	Address        CurrentDirectorAddress        `json:"address"`
	Gender         string                        `json:"gender"`
	DateOfBirth    string                        `json:"dateOfBirth"`
	Nationality    string                        `json:"nationality"`
	DirectorType   string                        `json:"directorType"`
	Positions      []Position                    `json:"positions"`
	AdditionalData CurrentDirectorAdditionalData `json:"additionalData"`
}

type CurrentDirectorAdditionalData struct {
	PresentAppointments   int64  `json:"presentAppointments"`
	Disqualified          bool   `json:"disqualified"`
	DisqualifiedException bool   `json:"disqualifiedException"`
	Occupation            string `json:"occupation"`
}

type CurrentDirectorAddress struct {
	Type        string  `json:"type"`
	SimpleValue string  `json:"simpleValue"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	PostalCode  string  `json:"postalCode"`
	Province    *string `json:"province,omitempty"`
}

type Position struct {
	DateAppointed string `json:"dateAppointed"`
	PositionName  string `json:"positionName"`
}

type PreviousDirector struct {
	ID              string                         `json:"id"`
	IDType          string                         `json:"idType"`
	Name            string                         `json:"name"`
	Title           string                         `json:"title"`
	Address         ActivePSCAddress               `json:"address"`
	Gender          string                         `json:"gender"`
	DateOfBirth     string                         `json:"dateOfBirth"`
	Nationality     string                         `json:"nationality"`
	DirectorType    string                         `json:"directorType"`
	Positions       []Position                     `json:"positions"`
	AdditionalData  PreviousDirectorAdditionalData `json:"additionalData"`
	ResignationDate string                         `json:"resignationDate"`
}

type PreviousDirectorAdditionalData struct {
	PresentAppointments   int64  `json:"presentAppointments"`
	PreviousAppointments  int64  `json:"previousAppointments"`
	DissolvedAppointments int64  `json:"dissolvedAppointments"`
	Disqualified          bool   `json:"disqualified"`
	DisqualifiedException bool   `json:"disqualifiedException"`
	Occupation            string `json:"occupation"`
}

type FinancialStatement struct {
	Type                 string                            `json:"type"`
	YearEndDate          string                            `json:"yearEndDate"`
	NumberOfWeeks        int64                             `json:"numberOfWeeks"`
	Currency             Currency                          `json:"currency"`
	ConsolidatedAccounts bool                              `json:"consolidatedAccounts"`
	ProfitAndLoss        map[string]int64                  `json:"profitAndLoss"`
	BalanceSheet         map[string]int64                  `json:"balanceSheet"`
	OtherFinancials      FinancialStatementOtherFinancials `json:"otherFinancials"`
	Ratios               map[string]float64                `json:"ratios"`
}

type FinancialStatementOtherFinancials struct {
	ContingentLiabilities string `json:"contingentLiabilities"`
	WorkingCapital        int64  `json:"workingCapital"`
	NetWorth              int64  `json:"netWorth"`
}

type LocalFinancialStatement struct {
	Type                 Type                                    `json:"type"`
	YearEndDate          string                                  `json:"yearEndDate"`
	NumberOfWeeks        int64                                   `json:"numberOfWeeks"`
	Currency             Currency                                `json:"currency"`
	ConsolidatedAccounts bool                                    `json:"consolidatedAccounts"`
	AuditQualification   *string                                 `json:"auditQualification,omitempty"`
	ProfitAndLoss        map[string]int64                        `json:"profitAndLoss"`
	BalanceSheet         map[string]int64                        `json:"balanceSheet"`
	CashFlow             *CashFlow                               `json:"cashFlow,omitempty"`
	OtherFinancials      *LocalFinancialStatementOtherFinancials `json:"otherFinancials,omitempty"`
	Ratios               map[string]float64                      `json:"ratios,omitempty"`
	AccountsStatus       *string                                 `json:"accountsStatus,omitempty"`
}

type CashFlow struct {
	NetCashFlowBeforeFinancing int64 `json:"netCashFlowBeforeFinancing"`
}

type LocalFinancialStatementOtherFinancials struct {
	ContingentLiabilities bool  `json:"contingentLiabilities"`
	BankOverdraftAndLTL   int64 `json:"bankOverdraftAndLTL"`
	WorkingCapital        int64 `json:"workingCapital"`
	CapitalEmployed       int64 `json:"capitalEmployed"`
	NetWorth              int64 `json:"netWorth"`
}

type NegativeInformation struct {
	CcjSummary CcjSummary `json:"ccjSummary"`
}

type CcjSummary struct {
	ExactRegistered    int64    `json:"exactRegistered"`
	PossibleRegistered int64    `json:"possibleRegistered"`
	NumberOfExact      int64    `json:"numberOfExact"`
	NumberOfPossible   int64    `json:"numberOfPossible"`
	NumberOfSatisfied  int64    `json:"numberOfSatisfied"`
	NumberOfWrits      int64    `json:"numberOfWrits"`
	Currency           Currency `json:"currency"`
}

type OtherInformation struct {
	EmployeesInformation []EmployeesInformation `json:"employeesInformation"`
}

type EmployeesInformation struct {
	Year              int64  `json:"year"`
	NumberOfEmployees string `json:"numberOfEmployees"`
}

type PaymentData struct {
	PaymentsOnFile                         int64  `json:"paymentsOnFile"`
	PaymentsOnTime                         int64  `json:"paymentsOnTime"`
	PaymentsPaidLate                       int64  `json:"paymentsPaidLate"`
	PaymentsSentLegal                      int64  `json:"paymentsSentLegal"`
	PaymentsStillOwingLate                 int64  `json:"paymentsStillOwingLate"`
	PaymentsPaid0To30Days                  int64  `json:"paymentsPaid0to30Days"`
	HighestInvoiceValueOutstandingLate     int64  `json:"highestInvoiceValueOutstandingLate"`
	PaymentsPaid90DaysplusLate             int64  `json:"paymentsPaid90DaysplusLate"`
	TotalBalanceStillOwingLate             int64  `json:"totalBalanceStillOwingLate"`
	PaymentsPaid61To90Days                 int64  `json:"paymentsPaid61to90Days"`
	TotalBalanceStillOwing                 int64  `json:"totalBalanceStillOwing"`
	Payments31To60DaysLate                 int64  `json:"payments31to60DaysLate"`
	Payments61To90DaysLate                 int64  `json:"payments61to90DaysLate"`
	HighestInvoiceValueOutstanding         int64  `json:"highestInvoiceValueOutstanding"`
	PaymentsStillOwing                     int64  `json:"paymentsStillOwing"`
	PaymentsWithinTerms                    int64  `json:"paymentsWithinTerms"`
	Payments0To30Dayslate                  int64  `json:"payments0to30Dayslate"`
	AverageInvoiceValue                    int64  `json:"averageInvoiceValue"`
	PaymentsPaid31To60Days                 int64  `json:"paymentsPaid31to60Days"`
	PaymentsPaid90Daysplus                 int64  `json:"paymentsPaid90Daysplus"`
	TotalInvoiceValues                     int64  `json:"totalInvoiceValues"`
	PaymentTrend                           string `json:"paymentTrend"`
	IndustryDBT                            int64  `json:"industryDBT"`
	NumberOfInvoicesPaidWithinTerms        int64  `json:"numberOfInvoicesPaidWithinTerms"`
	NumberOfInvoicesPaid1To30Days          int64  `json:"numberOfInvoicesPaid1To30Days"`
	NumberOfInvoicesPaid31To60Days         int64  `json:"numberOfInvoicesPaid31To60Days"`
	NumberOfInvoicesPaid61To90Days         int64  `json:"numberOfInvoicesPaid61To90Days"`
	NumberOfInvoicesPaid91PlusDays         int64  `json:"numberOfInvoicesPaid91PlusDays"`
	NumberOfInvoicesOutstandingWithinTerms int64  `json:"numberOfInvoicesOutstandingWithinTerms"`
	NumberOfInvoicesOutstanding1To30Days   int64  `json:"numberOfInvoicesOutstanding1To30Days"`
	NumberOfInvoicesOutstanding31To60Days  int64  `json:"numberOfInvoicesOutstanding31To60Days"`
	NumberOfInvoicesOutstanding61To90Days  int64  `json:"numberOfInvoicesOutstanding61To90Days"`
	NumberOfInvoicesOutstanding91PlusDays  int64  `json:"numberOfInvoicesOutstanding91PlusDays"`
}

type ShareCapitalStructure struct {
	IssuedShareCapital   LatestShareholdersEquityFigure `json:"issuedShareCapital"`
	NumberOfSharesIssued int64                          `json:"numberOfSharesIssued"`
	ShareHolders         []ShareHolder                  `json:"shareHolders"`
}

type ShareHolder struct {
	Name                     string       `json:"name"`
	ShareholderType          string       `json:"shareholderType"`
	ShareType                string       `json:"shareType"`
	Currency                 Currency     `json:"currency"`
	TotalValueOfSharesOwned  int64        `json:"totalValueOfSharesOwned"`
	TotalNumberOfSharesOwned int64        `json:"totalNumberOfSharesOwned"`
	PercentSharesHeld        *float64     `json:"percentSharesHeld,omitempty"`
	ShareClasses             []ShareClass `json:"shareClasses"`
	ID                       *string      `json:"id,omitempty"`
}

type ShareClass struct {
	ShareType           string                   `json:"shareType"`
	Currency            Currency                 `json:"currency"`
	ValuePerShare       int64                    `json:"valuePerShare"`
	NumberOfSharesOwned int64                    `json:"numberOfSharesOwned"`
	ValueOfSharesOwned  int64                    `json:"valueOfSharesOwned"`
	AdditionalData      ShareClassAdditionalData `json:"additionalData"`
}

type ShareClassAdditionalData struct {
	PercentSharesHeld float64 `json:"percentSharesHeld"`
}

type Currency string

const (
	Gbp Currency = "GBP"
)

type Description string

const (
	FinancialStatementsTooOld Description = "Financial Statements too old"
	LowRisk                   Description = "Low Risk"
	ModerateRisk              Description = "Moderate Risk"
	VeryLowRisk               Description = "Very Low Risk"
)

type Type string

const (
	LocalFinancialsCSUK Type = "LocalFinancialsCSUK"
	LocalFinancialsGAAP Type = "LocalFinancialsGAAP"
)
