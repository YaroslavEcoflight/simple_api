package dto

type BookCreateRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

type BookUpdateRequest struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
