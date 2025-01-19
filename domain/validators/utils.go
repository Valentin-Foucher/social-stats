package validators

import (
	"errors"
	"fmt"

	"github.com/Valentin-Foucher/social-stats/common"
)

func validationError(errorMessage string) *common.AppError {
	return common.NewInvalidParametersError(errors.New(errorMessage), errorMessage)
}

func validationErrorf(errorPattern string, parameters ...any) *common.AppError {
	return validationError(fmt.Sprintf(errorPattern, parameters...))
}
