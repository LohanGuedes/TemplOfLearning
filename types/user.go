package types

type User struct {
	ID    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
