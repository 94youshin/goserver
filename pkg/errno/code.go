package errno

var (
	//OK represents a successful request.
	OK = &Errno{Code: 0, Message: "OK"}

	// InternalServerError presents all unknown server-side errors.
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}

	// ErrPageNotFound represents a route not matched error.
	ErrPageNotFound = &Errno{Code: 10003, Message: "Page not found."}
)
