package handlers

import (
	"github.com/tandesko/cards_validator/internal/service/helpers"
	"github.com/tandesko/cards_validator/internal/service/requests"
	"github.com/tandesko/cards_validator/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"time"
)

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	log := helpers.Log(r)

	request, err := requests.NewValidateCardRequest(r)
	if err != nil {
		log.WithError(err).Info("wrong request")
		ape.Render(w, resources.ValidationResultResponse{
			Data: resources.ValidationResult{
				Key: resources.Key{
					ID:   "1",
					Type: "validate-card-response",
				},
				Attributes: resources.ValidationResultAttributes{
					Valid: false,
					Error: &resources.ErrorDetails{
						ErrorCode:    400,
						ErrorMessage: errors.Wrap(err, "wrong request").Error(),
					},
				},
			},
		})
		return
	}

	valid := false
	if int32(time.Now().Month()) < request.Data.Attributes.Month && int32(time.Now().Year()) <= request.Data.Attributes.Year {
		valid = true
	}

	ape.Render(w, resources.ValidationResultResponse{
		Data: resources.ValidationResult{
			Key: resources.Key{
				ID:   "1",
				Type: "validate-card-response",
			},
			Attributes: resources.ValidationResultAttributes{
				Valid: valid,
			},
		},
	})
}
