package domain

import (
	"context"

	"gorm.io/gorm"
)

type PlanType string

const (
	PlanType_Daily   PlanType = "daily"
	PlanType_Weekly  PlanType = "weekly"
	PlanType_Monthly PlanType = "monthly"
)

type Plan struct {
	ID       string   `json:"id,omitempty"`
	Type     PlanType `json:"type,omitempty"`
	Content  string   `json:"content,omitempty"`
	Times    int      `json:"times,omitempty"`
	Duration int64    `json:"duration,omitempty"`
}

type PlanListResp struct {
	BaseResp
	Data []*Plan
}

type PlanCreateReq struct {
	Type     PlanType `json:"type,omitempty"`
	Content  string   `json:"content,omitempty"`
	Times    int      `json:"times,omitempty"`
	Duration int64    `json:"duration,omitempty"`
}

type PlanDeleteReq struct {
	ID string `json:"id,omitempty"`
}

type UcPlan struct {
	ID       string
	Type     PlanType
	Content  string
	Times    int
	Duration int64
}

type PlanUsecase interface {
	List(ctx context.Context) ([]*UcPlan, error)
	Create(ctx context.Context, plan *UcPlan) error
	DeleteByID(ctx context.Context, id string) error
}

type RepoPlan struct {
	gorm.Model
	Type     PlanType
	Content  string
	Times    int
	Duration int64
}

type PlanRepository interface {
	List(ctx context.Context) ([]*RepoPlan, error)
	Create(ctx context.Context, plan *RepoPlan) error
	DeleteByID(ctx context.Context, id uint) error
	UpdateByID(ctx context.Context, plan *RepoPlan) error
}
