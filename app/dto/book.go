package dto

type Book struct {
	Id     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

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
