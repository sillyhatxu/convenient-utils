package response

import (
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
)

type ResponseEntity struct {
	Code  ResponseCode `json:"code"`
	Data  interface{}  `json:"data"`
	Msg   string       `json:"message"`
	Extra interface{}  `json:"extra"`
}

type ResponseCode int

const (
	Success ResponseCode = 1

	BusinessError ResponseCode = 0

	Error ResponseCode = -1

	ParamsValidateError ResponseCode = -2

	NotFoundError ResponseCode = -3

	ImproperOperationError ResponseCode = -4

	UnauthorizedError ResponseCode = -5
)

var (
	SuccessMessage                = ResponseMessage{MsgZH: "服务器异常", MsgEN: "SUCCESS"}
	ErrorMessage                  = ResponseMessage{MsgZH: "服务器异常", MsgEN: "Server Error"}
	ParamsValidateErrorMessage    = ResponseMessage{MsgZH: "系统验证错误", MsgEN: "System Validate Error"}
	NotFoundErrorMessage          = ResponseMessage{MsgZH: "服务器异常，未找到相关事件", MsgEN: "Server Error. Not found"}
	ImproperOperationErrorMessage = ResponseMessage{MsgZH: "操作不合法", MsgEN: "This operation is not appropriate"}
	UnauthorizedErrorMessage      = ResponseMessage{MsgZH: "您无权访问此页面", MsgEN: "You are not authorized to access this page"}
)

type ResponseMessage struct {
	MsgZH string
	MsgEN string
}

func (rm ResponseMessage) GetResponseMessage(languageTag language.Tag) string {
	switch languageTag {
	case language.English:
		return rm.MsgEN
	case language.Chinese:
		return rm.MsgZH
	default:
		return rm.MsgEN
	}
}

func ServerSuccess(data interface{}, extra interface{}) *ResponseEntity {
	return &ResponseEntity{
		Code:  Success,
		Data:  data,
		Msg:   SuccessMessage.GetResponseMessage(language.English),
		Extra: extra,
	}
}

func ServerError(data interface{}, msg string, extra interface{}) error {
	if msg == "" {
		msg = ErrorMessage.GetResponseMessage(language.English)
	}
	return &SillyHatError{
		Code:  Error,
		Data:  data,
		Msg:   msg,
		Extra: extra,
	}
}

func ServerParamsValidateError(data interface{}, extra interface{}) error {
	return &SillyHatError{
		Code:  ParamsValidateError,
		Data:  data,
		Msg:   ParamsValidateErrorMessage.GetResponseMessage(language.English),
		Extra: extra,
	}
}

func ServerNotFoundError(data interface{}, extra interface{}) error {
	return &SillyHatError{
		Code:  NotFoundError,
		Data:  data,
		Msg:   NotFoundErrorMessage.GetResponseMessage(language.English),
		Extra: extra,
	}
}

func ServerImproperOperationError(data interface{}, extra interface{}) error {
	return &SillyHatError{
		Code:  ImproperOperationError,
		Data:  data,
		Msg:   ImproperOperationErrorMessage.GetResponseMessage(language.English),
		Extra: extra,
	}
}

func ServerUnauthorizedError(data interface{}, extra interface{}) error {
	return &SillyHatError{
		Code:  UnauthorizedError,
		Data:  data,
		Msg:   UnauthorizedErrorMessage.GetResponseMessage(language.English),
		Extra: extra,
	}
}

type SillyHatError struct {
	Code  ResponseCode  `json:"code"`
	Data  interface{}   `json:"data"`
	Msg   string        `json:"message"`
	Extra interface{}   `json:"extra"`
	Args  []interface{} `json:"-"`
}

func (re *SillyHatError) Error() string {
	reJSON, err := json.Marshal(re)
	if err != nil {
		return fmt.Sprintf(`{"code": %v, "data": "", "message": "%v","extra": "SillyHatError to json error."}`, re.Code, re.Msg)
	}
	if re.Args != nil && len(re.Args) > 0 {
		return fmt.Sprintf(string(reJSON), re.Args)
	} else {
		return string(reJSON)
	}
}
