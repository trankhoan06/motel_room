package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Root       error  `json:"-"`
	Msg        string `json:"msg"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewAppError(root error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Root:       root,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}
func (e *AppError) RootErr() error {
	if err, ok := e.Root.(*AppError); ok {
		return err.RootErr()
	}
	return e.Root
}
func (e *AppError) Error() string {
	return e.RootErr().Error()
}
func NewFullErrorResponse(status int, root error, msg, log, key string) *AppError {
	return &AppError{
		status,
		root,
		msg,
		log,
		key,
	}
}
func NewAuthorize(root error, msg, log, key string) *AppError {
	return &AppError{
		http.StatusUnauthorized,
		root,
		msg,
		log,
		key,
	}
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"Something went wrong in the server", err.Error(), "ErrInternal")
}
func NewCustomErr(root error, msg, key string) *AppError {
	if root != nil {
		return NewAppError(root, msg, root.Error(), key)
	}
	return NewAppError(errors.New(msg), msg, msg, key)
}
func ErrDb(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}
func ErrInvalid(err error) *AppError {
	return NewCustomErr(err, "Invalid request", "ERRVALID")
}
func ErrEmailOfPass(err error) *AppError {
	return NewCustomErr(err, "email of pass invalid", "ERRVALID")
}
func ErrPass(err error) *AppError {
	return NewCustomErr(err, "pass invalid", "ERRVALID")
}
func ErrItem(err error) *AppError {
	return NewCustomErr(err, "item not found", "ERRITEM")
}
func ErrCart(err error) *AppError {
	return NewCustomErr(err, "cart haven't this item of this itetm has been deleted", "ERRITEM_CART")
}
func ErrUnauthorized(err error) *AppError {
	return NewFullErrorResponse(http.StatusUnauthorized, err, "Unauthorized", err.Error(), "ErrUnauthorized")
}
func ErrCommonDeleted(err error) *AppError {
	return NewCustomErr(err, "comment haven't this item of this comment has been deleted of no exist", "ERRCOMMENT")
}
func ErrUserUpdate(err error) *AppError {
	return NewCustomErr(err, "no permission", "ERRITEM_USER")
}
func ErrUneditedUpdate(err error) *AppError {
	return NewCustomErr(err, "no permission", "ERRITEM_USER")
}
func ErrOrder(err error) *AppError {
	return NewCustomErr(err, "you don't order this item", "ERRITEM_USER")
}
func ErrEmailNoExist(err error) *AppError {
	return NewCustomErr(err, "Email don't exist", "ERREMAIL_USER")
}
