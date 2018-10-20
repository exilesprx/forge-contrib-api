package values

// Email an email address
type Email struct {
	Address string `json:"email" bson:"email"`
}
