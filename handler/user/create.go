package user

import (
	. "apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.",lager.Data{"X-Request-Id": util.GetReqID(c)})
	// 从HTTP消息体获取参数（用户名和密码）
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.参数校验
	if err := u.Validate(); err != nil{
		SendResponse(c,errno.ErrValidation,nil)
		return
	}

	// Encrypt the user password.加密密码
	if err := u.Encrypt(); err != nil{
		SendResponse(c,errno.ErrEncrypt,nil)
		return
	}

	// Insert the user to the database.在数据库总添加数据记录
	if err := u.Create(); err != nil{
		SendResponse(c,errno.ErrDatabase,nil)
		return
	}
	// 返回结果（这里是用户名）
	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information
	SendResponse(c,nil,rsp)

}

// 参数校验
func (r *CreateRequest) checkParam() error{
	if r.Username == ""{
		return errno.New(errno.ErrValidation,nil).Add("username is empty.")
	}

	if r.Password == ""{
		return errno.New(errno.ErrValidation,nil).Add("password is empty")
	}

	return nil
}
