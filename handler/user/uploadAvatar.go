package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"tongue/handler"
	"tongue/log"
	"tongue/pkg/errno"
	service2 "tongue/service"
	service "tongue/service/user"
	"tongue/util"
)

// UploadAvatar ... 上传头像
// @Summary Get Qiniuyun token
// @Description
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "token 用户令牌"
// @Success 200 {string} json {"Code":200,"Token":"token"}
// @Router /user/avatar [post]
func UploadAvatar(c *gin.Context) {
	log.Info("User getInfo function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	file, err := c.FormFile("avatar")
	email := c.MustGet("email").(string)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrGetFile, nil, err.Error(), handler.GetLine())
		return
	}
	err, url := service2.UploadFile(file)
	if err != nil || url == "" {
		handler.SendError(c, errno.ErrUploadFile, nil, err.Error(), handler.GetLine())
		return
	}
	err = service.UpdateAvatar(email, url)
	if err != nil {
		handler.SendError(c, errno.ErrDatabase, nil, err.Error(), handler.GetLine())
		return
	}
	handler.SendResponse(c, nil, "Upload avatar success!")
}
