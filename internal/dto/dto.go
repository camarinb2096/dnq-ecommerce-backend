package dtos

type RequestParams struct {
	ProductId   string `json:"productId" form:"productId"`
	ProductName string `json:"productName" form:"productName"`
	Page        string `json:"page" form:"page"`
	PageSize    string `json:"pageSize" form:"pageSize"`
}

type Response struct {
	Message  string    `json:"message"`
	Page     int       `json:"page"`
	PageSize int       `json:"pageSize"`
	Total    int       `json:"total"`
	Data     []Product `json:"data"`
}

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
}

type CreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Prefix   string `json:"prefix"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
