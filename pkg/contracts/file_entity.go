package contracts

type FileEntityInterface interface {
	GetUuid() string
	GetName() string
	GetDisk() string
	GetPath() string
	GetOwnerSign() string
}
