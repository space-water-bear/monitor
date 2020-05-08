package work

type createRequest struct {
	ID   uint64 `json:"id" form:"id"`
	Data string `json:"data" form:"data"`
}
