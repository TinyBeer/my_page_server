package mysql

import (
	"context"
	"errors"

	"personal_page/domain"

	"gorm.io/gorm"
)

type PlanRepo struct {
	db *gorm.DB
}

// List implements domain.PlanRepository.
func (p *PlanRepo) List(ctx context.Context) ([]*domain.RepoPlan, error) {
	var list []*domain.RepoPlan
	err := p.db.WithContext(ctx).Find(&list).Error
	return list, err
}

// Create implements domain.PlanRepository.
func (p *PlanRepo) Create(ctx context.Context, plan *domain.RepoPlan) error {
	return p.db.WithContext(ctx).Create(plan).Error
}

// DeleteByID implements domain.PlanRepository.
func (p *PlanRepo) DeleteByID(ctx context.Context, id uint) error {
	return p.db.WithContext(ctx).Delete(&domain.RepoPlan{Model: gorm.Model{ID: id}}).Error
}

// UpdateByID implements domain.PlanRepository.
func (p *PlanRepo) UpdateByID(ctx context.Context, plan *domain.RepoPlan) error {
	if plan.ID == 0 {
		return errors.New("invalid id")
	}
	return p.db.WithContext(ctx).Updates(plan).Error
}

func NewPlanRepository(db *gorm.DB) domain.PlanRepository {
	repo := &PlanRepo{
		db: db,
	}
	err := repo.db.AutoMigrate(new(domain.RepoPlan))
	if err != nil {
		panic(err)
	}
	return repo
}
