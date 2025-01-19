package validators

import (
	"github.com/Valentin-Foucher/social-stats/common"
	"github.com/Valentin-Foucher/social-stats/domain/providers"
)

func ValidateProvider(providerID string) *common.AppError {
	if _, err := providers.New(providerID, nil); err != nil {
		return validationErrorf("Provider %s is not supported", providerID)
	}
	return nil
}
