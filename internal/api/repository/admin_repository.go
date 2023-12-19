package repository

import "gorm.io/gorm"

type AdminRepository interface {
	Login(user string, password string) (string, error)
}

type AdminRepositoryImpl struct {
	Db *gorm.DB
}

// Login implements AdminRepository.
func (*AdminRepositoryImpl) Login(user string, password string) (string, error) {
	panic("unimplemented")
}

func NewAdminRepositoryImpl(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{Db: db}
}
