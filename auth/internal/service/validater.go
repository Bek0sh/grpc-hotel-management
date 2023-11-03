package service

import (
	"github.com/Bek0sh/hotel-management-auth/pkg/pb"
	"net/http"
	"net/mail"
	"strings"

	"github.com/Bek0sh/hotel-management-auth/internal/myerrors"
	"github.com/Bek0sh/hotel-management-auth/pkg/utils"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

func validateRegisterUser(req *pb.RegisterRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	name := strings.Fields(req.GetFullName())[0]
	surname := strings.Fields(req.GetFullName())[1]
	if err := utils.ValidateUsername(name); err != nil {
		violations = append(violations, myerrors.FieldViolation("name", err))
	}
	if err := utils.ValidateUsername(surname); err != nil {
		violations = append(violations, myerrors.FieldViolation("surname", err))
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, myerrors.FieldViolation("password", err))
	}
	if err := utils.ValidatePassword(req.GetConfirmPassword()); err != nil {
		violations = append(violations, myerrors.FieldViolation("confirm_password", err))
	}
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		violations = append(violations, myerrors.FieldViolation("email", err))
	}

	return violations
}

func validateSignIn(req *pb.SignInRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if _, err := mail.ParseAddress(req.GetEmail()); err != nil {
		violations = append(violations, myerrors.FieldViolation("email", err))
	}
	if err := utils.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, myerrors.FieldViolation("password", err))
	}
	return
}

func InvalidArgumentError(violations []*errdetails.BadRequest_FieldViolation) error {
	badrequest := &errdetails.BadRequest{
		FieldViolations: violations,
	}

	statusInvaild := status.New(http.StatusBadRequest, "invalid parametres")

	statusDetails, err := statusInvaild.WithDetails(badrequest)
	if err != nil {
		return statusInvaild.Err()
	}

	return statusDetails.Err()
}
