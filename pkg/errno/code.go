package errno

var (
	// Common errors
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	OK                  = &Errno{Code: 20000, Message: "OK"}
	ErrValidation       = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase         = &Errno{Code: 20002, Message: "Database error."}
	ErrToken            = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrNotFind 			= &Errno{Code: 40400, Message: "404,Not Find."}
	// user errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	// Image upload errors
	ErrUploadSaveImageFail    = &Errno{Code: 30001, Message: "保存图片失败"}
	ErrUploadCheckImageFail   = &Errno{Code: 30002, Message: "检查图片失败"}
	ErrUploadCheckImageFormat = &Errno{Code: 30003, Message: "检查图片错误，图片格式或大小有问题"}

	// Video errors
	ErrCreateFail = &Errno{Code: 40001, Message: "视频创建失败"}
	ErrUpdateFail = &Errno{Code: 40002, Message: "视频更新失败"}
	ErrShowFail = &Errno{Code: 40003, Message: "视频显示失败"}
	ErrDeleteFail = &Errno{Code: 40004, Message: "视频删除失败"}
)
