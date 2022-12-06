package handler

type storagePlaceResponseDto struct {
	StoragePlaceId uint   `json:"storagePlaceId"`
	Title          string `json:"title,omitempty"`
	Code           string `json:"code"`
}
