package main

type Request struct {
	Data []struct {
		UserId      int64  `json:"user_id"`
		Name        string `json:"name"`
		DateOfBirth string `json:"date_of_birth"`
		CreatedOn   int64  `json:"created_on"`
	} `json:"data"`
}

type Response struct {
	Data []ResponseItem `json:"data"`
}

type ResponseItem struct {
	UserId      int64  `json:"user_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
	CreatedOn   string `json:"created_on"`
}
