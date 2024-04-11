package productsDto

type RequestParams struct {
	ProductId   string `json:"productId" form:"productId"`
	ProductName string `json:"productName" form:"productName"`
	Page        string `json:"page" form:"page"`
	PageSize    string `json:"pageSize" form:"pageSize"`
}

type Response struct {
	Message  string  `json:"message"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
	Total    int     `json:"total"`
	Data     Product `json:"data"`
}

type Product struct {
	ProductId   string  `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Image       string  `json:"image"`
	Quantity    int     `json:"quantity"`
}
