package request

type FsLogin struct {
	Origin string `json:"origin" binding:"required"`
	Code string `json:"code" binding:"required"`
}