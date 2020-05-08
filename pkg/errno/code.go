package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	ErrValidation      = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase        = &Errno{Code: 20002, Message: "Database error."}
	ErrToken           = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token"}
	ErrErcordDuplicate = &Errno{Code: 20004, Message: "记录已存在"}
	ErrEncodeError     = &Errno{Code: 20005, Message: "解码失败"}
)
