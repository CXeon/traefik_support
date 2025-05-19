package http

import (
	"github.com/CXeon/micro_contrib/response"
	"github.com/gin-gonic/gin"
)

//统一认证接口
func (c *controller) forwardAuth(ctx *gin.Context) {

	response.ResponseSuccess(ctx, struct{}{})
}
