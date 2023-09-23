package models

type BillerCategory struct {
	ID                     string `json:"Id"`
	BillerID               string `json:"billerId"`
	BillerName             string `json:"billerName"`
	BillerAliasName        string `json:"billerAliasName"`
	BillerCategoryName     string `json:"billerCategoryName"`
	BillerSubCategoryName  string `json:"billerSubCategoryName"`
	BillerMode             string `json:"billerMode"`
	BillerAcceptsAdhoc     string `json:"billerAcceptsAdhoc"`
	ParentBiller           string `json:"parentBiller"`
	ParentBillerID         string `json:"parentBillerId"`
	BillerCoverage         string `json:"billerCoverage"`
	FetchRequirement       string `json:"fetchRequirement"`
	PaymentAmountExactness string `json:"paymentAmountExactness"`
	SupportBillValidation  string `json:"supportBillValidation"`
	BillerEffctvFrom       string `json:"billerEffctvFrom"`
	BillerPymtModes        []struct {
		PaymentMode string `json:"paymentMode"`
		MaxLimit    string `json:"maxLimit"`
		MinLimit    string `json:"minLimit"`
	} `json:"billerPymtModes"`
	BillerPymtChnls []struct {
		PaymentChannel string `json:"paymentChannel"`
		MaxLimit       string `json:"maxLimit"`
		MinLimit       string `json:"minLimit"`
	} `json:"billerPymtChnls"`
	CustomerParams []struct {
		ParamName  string `json:"paramName"`
		DataType   string `json:"dataType"`
		Optional   string `json:"optional"`
		MinLength  string `json:"minLength"`
		MaxLength  string `json:"maxLength"`
		Regex      string `json:"regex"`
		Visibility string `json:"visibility"`
	} `json:"customerParams"`
	ReponseParams []struct {
		AmtBreakupList []struct {
			AmtBreakup string `json:"amtBreakup"`
		} `json:"amtBreakupList"`
	} `json:"reponseParams"`
	AdditonalInfo       []interface{} `json:"additonalInfo"`
	IntFeeConf          []interface{} `json:"intFeeConf"`
	IntChngFee          []interface{} `json:"intChngFee"`
	Status              string        `json:"status"`
	BillerResponseType  string        `json:"billerResponseType"`
	CustomerParamGroups struct {
		Group []interface{} `json:"group"`
	} `json:"customerParamGroups"`
	BillerPlanResponseParams struct {
		Params        []interface{} `json:"params"`
		AmountOptions []struct {
			AmtBreakupList []interface{} `json:"amtBreakupList"`
		} `json:"amountOptions"`
	} `json:"billerPlanResponseParams"`
	AdditonalInfoPayment []interface{} `json:"additonalInfoPayment"`
	PlanAdditionalInfo   []interface{} `json:"planAdditionalInfo"`
	PlanMdmRequirement   string        `json:"planMdmRequirement"`
}

// type PaymentMode struct {
// 	PaymentMode string `json:"paymentMode"`
// 	MaxLimit    string `json:"maxLimit"`
// 	MinLimit    string `json:"minLimit"`
// }

// type PaymentChannel struct {
// 	PaymentChannel string `json:"paymentChannel"`
// 	MaxLimit       string `json:"maxLimit"`
// 	MinLimit       string `json:"minLimit"`
// }

// type CustomerParams struct {
// 	ParamName  string `json:"paramName"`
// 	DataType   string `json:"dataType"`
// 	Optional   string `json:"optional"`
// 	MinLength  string `json:"minLength"`
// 	MaxLength  string `json:"maxLength"`
// 	Regex      string `json:"regex"`
// 	Visibility string `json:"visibility"`
// }

// type ResponseParams struct {
// 	AmtBreakupList []string `json:"amtBreakupList"`
// }

// // type BillerRequest struct {
// // 	BillerId               string           `json:"billerId"`
// // 	BillerName             string           `json:"billerName"`
// // 	BillerAliasName        string           `json:"billerAliasName"`
// // 	BillerCategoryName     string           `json:"billerCategoryName"`
// // 	BillerSubCategoryName  string           `json:"billerSubCategoryName"`
// // 	BillerMode             string           `json:"billerMode"`
// // 	BillerAcceptsAdhoc     string           `json:"billerAcceptsAdhoc"`
// // 	ParentBiller           string           `json:"parentBiller"`
// // 	ParentBillerId         string           `json:"parentBillerId"`
// // 	BillerCoverage         string           `json:"billerCoverage"`
// // 	FetchRequirement       string           `json:"fetchRequirement"`
// // 	PaymentAmountExactness string           `json:"paymentAmountExactness"`
// // 	SupportBillValidation  string           `json:"supportBillValidation"`
// // 	BillerEffctvFrom       string           `json:"billerEffctvFrom"`
// // 	BillerPymtModes        []PaymentMode    `json:"billerPymtModes"`
// // 	BillerPymtChnls        []PaymentChannel `json:"billerPymtChnls"`
// // 	CustomerParams         []CustomerParams `json:"customerParams"`
// // 	ResponseParams         ResponseParams   `json:"reponseParams"`
// // 	AdditonalInfo          []string         `json:"additonalInfo"`
// // 	IntFeeConf             []string         `json:"intFeeConf"`
// // 	IntChngFee             []string         `json:"intChnNullString
// // 	Status                 string           `json:"status"`
// // 	BillerResponseType     string           `json:"billerResponseType"`
// // 	CustomerParamGroups    struct {
// // 		Group []interface{} `json:"group"`
// // 	} `json:"customerParamGroups"`
// // 	BillerPlanResponseParams struct {
// // 		Params        []interface{} `json:"params"`
// // 		AmountOptions []struct {
// // 			AmtBreakupList []interface{} `json:"amtBreakupList"`
// // 		} `json:"amountOptions"`
// // 	} `json:"billerPlanResponseParams"`
// // 	AdditonalInfoPayment []string `json:"additonalInfoPayment"`
// // 	PlanAdditionalInfo   []string `json:"planAdditionalInfo"`
// // 	PlanMdmRequirement   string   `json:"planMdmRequirement"`
// // }
