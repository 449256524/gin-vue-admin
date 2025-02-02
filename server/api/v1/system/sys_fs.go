package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
)

type FsApi struct{}

func (f *FsApi) FsLogin (c *gin.Context) {
	var l systemReq.FsLogin
	if err := c.ShouldBind(&l); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	orifin := l.Origin
	code := l.Code
	fsService.Login(orifin, code)
	response.Ok(c)
}