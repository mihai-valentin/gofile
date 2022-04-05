package contracts

type FileRepositoryInterface interface {
	StoreFile(fileData FileUploadDataInterface) (FileEntityInterface, error)
	FindByUuid(uuid string) (FileEntityInterface, error)
	DeleteByUuid(uuid string) error
}
