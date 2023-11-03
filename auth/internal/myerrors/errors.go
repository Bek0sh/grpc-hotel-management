package myerrors

import "google.golang.org/genproto/googleapis/rpc/errdetails"

func FieldViolation(title string, err error) *errdetails.BadRequest_FieldViolation {
	return &errdetails.BadRequest_FieldViolation{
		Field:       title,
		Description: err.Error(),
	}
}
