package handlers

import (
	"github.com/tandesko/cards_validator/internal/service/api/helpers"
	"github.com/tandesko/cards_validator/internal/service/api/requests"
	"github.com/tandesko/cards_validator/internal/service/core"
	"github.com/tandesko/cards_validator/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)

	validationResultResponse := resources.ValidationResultResponse{
		Data: resources.ValidationResult{
			Key: resources.Key{
				ID:   "1",
				Type: "validate-card-response",
			},
			Attributes: resources.ValidationResultAttributes{
				Valid: false,
			},
		},
	}

	request, err := requests.NewValidateCardRequest(r)
	if err != nil {
		log.WithError(err).Info("wrong request")
		embedError(err, "wrong request", &validationResultResponse)
		ape.Render(w, validationResultResponse)
		return
	}

	validResult, err := core.ValidateAllMethods(request.Data.Attributes, log)
	if err != nil {
		embedError(err, "", &validationResultResponse)
		ape.Render(w, validationResultResponse)
		return
	}

	validationResultResponse.Data.Attributes.Valid = validResult
	ape.Render(w, validationResultResponse)
}

func embedError(err error, errorMsg string, response *resources.ValidationResultResponse) {
	response.Data.Attributes.Error = &resources.ErrorDetails{
		ErrorCode:    400,
		ErrorMessage: errors.Wrap(err, errorMsg).Error(),
	}
}
