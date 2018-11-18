package serializers

// User serializers
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Age       uint   `json:"age"`
}
