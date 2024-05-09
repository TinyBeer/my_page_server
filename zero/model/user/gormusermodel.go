package user

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GormUserModel = (*customGormUserModel)(nil)

type (
	// GormUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGormUserModel.
	GormUserModel interface {
		gormUserModel
		withSession(session sqlx.Session) GormUserModel
	}

	customGormUserModel struct {
		*defaultGormUserModel
	}
)

// NewGormUserModel returns a model for the database table.
func NewGormUserModel(conn sqlx.SqlConn) GormUserModel {
	return &customGormUserModel{
		defaultGormUserModel: newGormUserModel(conn),
	}
}

func (m *customGormUserModel) withSession(session sqlx.Session) GormUserModel {
	return NewGormUserModel(sqlx.NewSqlConnFromSession(session))
}
