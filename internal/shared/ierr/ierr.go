package ierr

import (
	"fmt"
)

type IErr struct {
	message string
	code    string
}

func (i IErr) Error() string {
	return fmt.Sprintf("[%s] message: %s", i.code, i.message)
}

var (
	ErrBadRequest     = IErr{message: "bad request", code: "400001"}
	ErrInternalServer = IErr{message: "internal server error", code: "500001"}
	ErrForbidden      = IErr{message: "forbidden request", code: "403001"}
	ErrUnAuthorized   = IErr{message: "unauthorized action", code: "401001"}
	ErrNotFound       = IErr{message: "not found item", code: "404001"}

	ErrTownCodeNotHaveParkingSlot = IErr{message: "town code not have parking slot", code: "404002"}
	ErrVehicleAlreadyInTown       = IErr{message: "vehicle already in town", code: "400002"}
	ErrVehicleNotFoundInTown      = IErr{message: "vehicle not found in town", code: "400003"}
)
