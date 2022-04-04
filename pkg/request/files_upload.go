package request

import "mime/multipart"

type FilesUploadRequest struct {
	Disk      string                  `form:"disk" binding:"required"`
	OwnerSign string                  `form:"owner_sign" binding:"required"`
	FormFiles []*multipart.FileHeader `form:"files[]" binding:"required"`
	Presets   []uint                  `form:"presets[]"`
}

func (r *FilesUploadRequest) GetFormFiles() []*multipart.FileHeader {
	return r.FormFiles
}

func (r *FilesUploadRequest) GetPresets() []uint {
	return r.Presets
}

func (r *FilesUploadRequest) GetDisk() string {
	return r.Disk
}

func (r *FilesUploadRequest) GetOwnerSign() string {
	return r.OwnerSign
}
