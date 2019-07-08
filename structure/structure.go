package structure

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
