package exceptions

const SuccessRuleCode string = "00"
const ParameterRuleCode string = "01"
const DataNotFoundRuleCode string = "02"
const InvalidRuleCode string = "03"
const ThirdPartyRuleException string = "04"

type ExceptionResponse struct {
	ResponseCode        string      `json:"responseCode"`
	ResponseDesc        string      `json:"responseDesc"`
	ResponseData        interface{} `json:"responseData,omitempty"`
	ResponseValidations []string    `json:"responseValidations,omitempty"`
}
