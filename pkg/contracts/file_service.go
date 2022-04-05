package contracts

type FileServiceInterface interface {
	UploadFiles(filesUploadData []FileUploadDataInterface) (map[string]string, error)
	GetFile(uuid string, ownerSign string) (string, error)
	DeleteFile(uuid string, ownerSign string) error
}
