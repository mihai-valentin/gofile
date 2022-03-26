package entity

type FileOwner struct {
	Uuid string `form:"owner_uuid" json:"owner_uuid" binding:"required"`
	Type string `form:"owner_type" json:"owner_type" binding:"required"`
}
