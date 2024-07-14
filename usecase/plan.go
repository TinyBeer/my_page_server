package usecase

import (
	"context"
	"strconv"

	"personal_page/domain"
)

type PlanUsecase struct {
	plan domain.PlanRepository
}

// List implements domain.PlanUsecase.
func (p *PlanUsecase) List(ctx context.Context) ([]*domain.UcPlan, error) {
	plans, err := p.plan.List(ctx)
	if err != nil {
		return nil, err
	}
	list := make([]*domain.UcPlan, 0, len(plans))
	for _, v := range plans {
		list = append(list, &domain.UcPlan{
			ID:       strconv.Itoa(int(v.ID)),
			Type:     v.Type,
			Content:  v.Content,
			Times:    v.Times,
			Duration: v.Duration,
		})
	}
	return list, nil
}

// Create implements domain.PlanUsecase.
func (p *PlanUsecase) Create(ctx context.Context, plan *domain.UcPlan) error {
	return p.plan.Create(ctx, &domain.RepoPlan{
		Type:     plan.Type,
		Content:  plan.Content,
		Times:    plan.Times,
		Duration: plan.Duration,
	})
}

// DeleteByID implements domain.PlanUsecase.
func (p *PlanUsecase) DeleteByID(ctx context.Context, id string) error {
	iid, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	return p.plan.DeleteByID(ctx, uint(iid))
}

func NewPlanUsecase(plan domain.PlanRepository) domain.PlanUsecase {
	return &PlanUsecase{
		plan: plan,
	}
}
