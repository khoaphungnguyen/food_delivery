package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessReponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{
		Data:   data,
		Paging: paging,
		Filter: filter,
	}
}

func SimpleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessReponse(data, nil, nil)
}
