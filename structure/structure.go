package structure

type User struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Response struct {
	Status  string `json: "status"`
	Data    User   `json: "data"`
	Message string `json: "data"`
}
