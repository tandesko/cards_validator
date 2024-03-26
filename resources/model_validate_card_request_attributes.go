package resources

type ValidateCardRequestAttributes struct {
	// Number of credit card
	Cardnumber string `json:"cardnumber"`
	// End month of credit card
	Month int32 `json:"month"`
	// End year of credit card
	Year int32 `json:"year"`
}
