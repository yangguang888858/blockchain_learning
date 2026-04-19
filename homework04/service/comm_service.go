package service

import "gorm.io/gorm"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func Paginate(current, size uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if current <= 0 {
			current = 1
		}
		if size <= 0 {
			size = 10
		}
		offset := (current - 1) * size
		return db.Offset(int(offset)).Limit(int(size))
	}
}
