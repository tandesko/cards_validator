package resources

import "encoding/json"

type ValidateCardRequest struct {
	Key
	Attributes ValidateCardRequestAttributes `json:"attributes"`
}
type ValidateCardRequestResponse struct {
	Data     ValidateCardRequest `json:"data"`
	Included Included            `json:"included"`
}

type ValidateCardRequestListResponse struct {
	Data     []ValidateCardRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
	Meta     json.RawMessage       `json:"meta,omitempty"`
}

// MustValidateCardRequest - returns ValidateCardRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustValidateCardRequest(key Key) *ValidateCardRequest {
	var validateCardRequest ValidateCardRequest
	if c.tryFindEntry(key, &validateCardRequest) {
		return &validateCardRequest
	}
	return nil
}
