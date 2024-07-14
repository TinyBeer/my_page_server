package mysql

import (
	"context"
	"errors"
	"log"

	"personal_page/domain"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// GetByName implements domain.UserRepository.
func (u *UserRepo) GetByName(ctx context.Context, username string) (*domain.RepoUser, error) {
	user := new(domain.RepoUser)
	err := u.db.Take(user, "name = ?", username).Error
	return user, err
}

// Check implements domain.UserRepository.
func (u *UserRepo) Check(ctx context.Context, user *domain.RepoUser) error {
	return u.db.WithContext(ctx).Take(user).Error
}

// Create implements domain.UserRepository.
func (u *UserRepo) Create(ctx context.Context, user *domain.RepoUser) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	admin := &domain.RepoUser{
		Name: "admin",
		// bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
		Password: "$2a$10$1iXYspdrld4iQ.W41m6iaOvbOgoyOncycJQ8pWRCdzyOWg8bsEMnq",
	}
	has := db.Migrator().HasTable(admin)
	if !has {
		err := errors.Join(db.AutoMigrate(admin), db.Create(admin).Error)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return &UserRepo{
		db: db,
	}
}
