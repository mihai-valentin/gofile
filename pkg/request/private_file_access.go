package request

type PrivateFileAccessRequest struct {
	OwnerGuid string `form:"owner_guid" binding:"required"`
	OwnerType string `form:"owner_type" binding:"required"`
}
