package dto

type GetUserReq struct {
	ID string `json:"-" format:"uuid" uri:"id"`
}
