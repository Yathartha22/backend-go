package backend

type GetInfoRequest struct {
	Id uint32
}

type GetInfoResponse struct {
	Info string `json:"info"`
}
