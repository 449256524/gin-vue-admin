package request

type FsLogin struct {
	Origin string `json:"origin"`
	Code string `json:"code"`
}