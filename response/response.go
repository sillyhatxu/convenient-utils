package response

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

type Entity struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"message"`
	Extra interface{} `json:"extra"`
}

const (
	Success string = "SUCCESS"

	BusinessError string = "BUSINESS_ERROR"

	Error string = "SERVER_ERROR"

	ParamsValidateError string = "PARAMS_VALIDATE_ERROR"

	NotFoundError string = "NOT_FOUND_ERROR"

	ImproperOperationError string = "IMPROPER_OPERATION_ERROR"

	UnauthorizedError string = "UNAUTHORIZED_ERROR"
)

var (
	SuccessMessage                = ResMessage{MsgZH: "服务器异常", MsgEN: "SUCCESS"}
	ErrorMessage                  = ResMessage{MsgZH: "服务器异常", MsgEN: "Server Error"}
	ParamsValidateErrorMessage    = ResMessage{MsgZH: "系统验证错误", MsgEN: "System Validate Error"}
	NotFoundErrorMessage          = ResMessage{MsgZH: "服务器异常，未找到相关事件", MsgEN: "Server Error. Not found"}
	ImproperOperationErrorMessage = ResMessage{MsgZH: "操作不合法", MsgEN: "This operation is not appropriate"}
	UnauthorizedErrorMessage      = ResMessage{MsgZH: "您无权访问此页面", MsgEN: "You are not authorized to access this page"}
)

type ResMessage struct {
	MsgZH string
	MsgEN string
}

func (rm ResMessage) GetResponseMessage(languageTag language.Tag) string {
	switch languageTag {
	case language.English:
		return rm.MsgEN
	case language.Chinese:
		return rm.MsgZH
	default:
		return rm.MsgEN
	}
}

func ServerSuccess(data interface{}, extra interface{}) *Entity {
	return &Entity{
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
	Code  string        `json:"code"`
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

func HTMLSuccess(data interface{}, extra interface{}) gin.H {
	return gin.H{
		"code":    Success,
		"data":    data,
		"message": SuccessMessage.GetResponseMessage(language.English),
		"extra":   extra,
	}
}
