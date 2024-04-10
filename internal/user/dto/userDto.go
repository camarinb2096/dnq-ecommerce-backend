package userDto

type CreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Prefix   string `json:"prefix"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}
