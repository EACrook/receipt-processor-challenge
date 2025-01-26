package models

type Points struct {
	ID string `json:"id" binding:"required"`
	Points int `json:"points" binding:"required"`
}
