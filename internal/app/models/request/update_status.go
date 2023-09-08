package request

type UpdateStatusRequest struct {
	Id     uint   `json:"id"`
	Status string `json:"status"`
}
