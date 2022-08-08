package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type FsApi struct{}

func (ta *FsApi) FsLogin (c *gin.Context) {
	response.Ok(c)
}