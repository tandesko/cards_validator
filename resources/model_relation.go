package resources

type Relation struct {
	Data  *Key   `json:"data,omitempty"`
	Links *Links `json:"links,omitempty"`
}
