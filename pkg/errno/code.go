package errno

var (
	// Common errors
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	OK                  = &Errno{Code: 20000, Message: "OK"}
	ErrValidation       = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase         = &Errno{Code: 20002, Message: "Database error."}
	ErrToken            = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	// Image upload errors
	ErrUploadSaveImageFail  = &Errno{Code: 30001, Message: "保存图片失败"}
	ErrUploadCheckImageFail = &Errno{Code: 30002, Message: "检查图片失败"}
	ErrUploadCheckImageFormat = &Errno{Code: 30003, Message: "检查图片错误，图片格式或大小有问题"}
)
