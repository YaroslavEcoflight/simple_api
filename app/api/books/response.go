package books

import "simple_api/app/model"

type Response struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func toResponse(b model.Book) Response {
	return Response{
		Id:     b.Id,
		Title:  b.Title,
		Author: b.Author,
		Rating: b.Rating,
	}
}
