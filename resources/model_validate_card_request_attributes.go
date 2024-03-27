package resources

type ValidateCardRequestAttributes struct {
	// Number of credit card
	Cardnumber string `json:"cardnumber"`
	// End month of credit card
	Month string `json:"month"`
	// End year of credit card
	Year string `json:"year"`
}
