package requests

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tandesko/cards_validator/resources"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"regexp"
)

type ValidateCardRequest struct {
	Data resources.ValidateCardRequest `json:"data"`
}

func NewValidateCardRequest(r *http.Request) (ValidateCardRequest, error) {
	var request ValidateCardRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return ValidateCardRequest{}, errors.Wrap(err, "failed to decode request")
	}

	return request, request.validate()
}

func (r ValidateCardRequest) validate() error {
	return validation.Errors{
		"/data/attributes/cardnumber": validation.Validate(
			&r.Data.Attributes.Cardnumber, validation.Required, validation.Match(regexp.MustCompile("^([1-9])[0-9]{11,18}$|^[1-9]{4}(?:-[0-9]{4}){3}$"))),
		"/data/attributes/month": validation.Validate(
			&r.Data.Attributes.Month, validation.Required, validation.Length(1, 2)),
		"/data/attributes/year": validation.Validate(
			&r.Data.Attributes.Year, validation.Required, validation.Length(2, 4)),
	}.Filter()
}
