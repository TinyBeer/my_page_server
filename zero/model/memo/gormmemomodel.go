package memo

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GormMemoModel = (*customGormMemoModel)(nil)

type (
	// GormMemoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGormMemoModel.
	GormMemoModel interface {
		gormMemoModel
		withSession(session sqlx.Session) GormMemoModel
	}

	customGormMemoModel struct {
		*defaultGormMemoModel
	}
)

// NewGormMemoModel returns a model for the database table.
func NewGormMemoModel(conn sqlx.SqlConn) GormMemoModel {
	return &customGormMemoModel{
		defaultGormMemoModel: newGormMemoModel(conn),
	}
}

func (m *customGormMemoModel) withSession(session sqlx.Session) GormMemoModel {
	return NewGormMemoModel(sqlx.NewSqlConnFromSession(session))
}
