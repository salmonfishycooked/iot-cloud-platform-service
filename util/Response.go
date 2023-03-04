package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	SUCCESS         = 0          // 成功的 code
	ERROR           = -1         // 有错误的 code
	SUCCESS_MESSAGE = "success!" // 成功信息
	ERROR_MESSAGE   = "error!"   // 失败信息
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// ResponseCommon
// @Description: 通用 Response
// @param ctx 上下文
func ResponseCommon(code int, data interface{}, msg string, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Data:    data,
		Message: msg,
	})
}

// ResponseOK
// @Description: 一般成功响应
// @param ctx 上下文
// @param data 需要返回的数据
func ResponseOK(ctx *gin.Context, data interface{}) {
	ResponseCommon(SUCCESS, data, SUCCESS_MESSAGE, ctx)
}

// ResponseOKWithMsg
// @Description: 带信息的成功响应
// @param ctx 上下文
// @param msg 成功信息
func ResponseOKWithMsg(ctx *gin.Context, msg string) {
	ResponseCommon(SUCCESS, nil, msg, ctx)
}

// ResponseError
// @Description: 一般错误响应
// @param ctx 上下文
func ResponseError(ctx *gin.Context) {
	ResponseCommon(ERROR, nil, ERROR_MESSAGE, ctx)
}

// ResponseErrorWithMsg
// @Description: 带信息的错误响应
// @param ctx 上下文
// @param msg 错误信息
func ResponseErrorWithMsg(ctx *gin.Context, msg string) {
	ResponseCommon(ERROR, nil, msg, ctx)
}
