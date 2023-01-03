package models

type Info struct {
	Meta struct {
		Pagination struct {
			Total int `json:"total"`
			Pages int `json:"pages"`
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Links struct {
				Previous string `json:"previous"`
				Current  string `json:"current"`
				Next     string `json:"next"`
			} `json:"links"`
		} `json:"pagination"`
	} `json:"meta"`
	Data []Post `json:"data"`
}
