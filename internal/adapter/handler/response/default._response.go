package response

type ErrorResponseDefault struct {
	Meta
}

type Meta struct {
	Status boolen `json:"status"`
	Message string `jsong:"message"`
}