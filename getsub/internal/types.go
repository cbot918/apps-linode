package internal

type GetSubRequest struct {
	Email    string `json:"email"`
	Name     string `json:"Name"`
	Password string `json:"password"`
}

type GetSubResponse struct {
	Email string `json:"email"`
	Name  string `json:"Name"`
}
