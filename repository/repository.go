package repository

import "gorm.io/gorm"

type Repositoy[T any] struct {
	DB *gorm.DB
}
