package validation

import (
	"regexp"

	"github.com/pkg/errors"
	kotsv1beta1 "github.com/replicatedhq/kots/kotskinds/apis/kots/v1beta1"
	configtypes "github.com/replicatedhq/kots/pkg/kotsadmconfig/types"
)

const (
	regexMatchError = "Value does not match regex"
)

type regexValidator struct {
	*kotsv1beta1.RegexValidator
}

func (v *regexValidator) Validate(input string) (*configtypes.ValidationError, error) {
	regex, err := regexp.Compile(v.Pattern)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to compile regex pattern")
	}

	matched := regex.MatchString(input)
	if !matched {
		if v.Message == "" {
			v.Message = regexMatchError
		}
		return &configtypes.ValidationError{
			Message: v.Message,
		}, nil
	}
	return nil, nil
}
