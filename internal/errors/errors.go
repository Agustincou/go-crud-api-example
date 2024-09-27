package errors

import "net/http"

type API struct {
	HttpStatus int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
}

func (a *API) Error() string {
	return a.Message
}

var (
	Unimplemented    = &API{HttpStatus: http.StatusInternalServerError, Code: 999, Message: "unimplemented"}
	InvalidName      = &API{HttpStatus: http.StatusBadRequest, Code: 1000, Message: "invalid name value"}
	InvalidQuantity  = &API{HttpStatus: http.StatusBadRequest, Code: 1001, Message: "invalid quantity value"}
	InvalidPrice     = &API{HttpStatus: http.StatusBadRequest, Code: 1002, Message: "invalid price value"}
	InvalidHeader    = &API{HttpStatus: http.StatusBadRequest, Code: 1003, Message: "bad request"}
	MalformedRequest = &API{HttpStatus: http.StatusBadRequest, Code: 1004, Message: "malformed request"}
	MethodNotAllowed = &API{HttpStatus: http.StatusMethodNotAllowed, Code: 1005, Message: "method not allowed"}
	InternalAPIError = &API{HttpStatus: http.StatusInternalServerError, Code: 1006, Message: "internal API error"}
	SaveError        = &API{HttpStatus: http.StatusInternalServerError, Code: 1007, Message: "save error"}
	GetError         = &API{HttpStatus: http.StatusInternalServerError, Code: 1008, Message: "get error"}
	DeleteError      = &API{HttpStatus: http.StatusInternalServerError, Code: 1009, Message: "delete error"}
	UpdateError      = &API{HttpStatus: http.StatusInternalServerError, Code: 1010, Message: "update error"}
)
