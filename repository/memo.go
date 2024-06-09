package repository

import (
	"personal_page/model"

	"gorm.io/gorm"
)

type MemoRepo struct {
	db *gorm.DB
}

// CompleteWithId implements MemoRepository.
func (m *MemoRepo) CompleteWithId(id uint) error {
	return m.db.Model(&model.Memo{}).Where("id = ?", id).UpdateColumn("completed", true).Error
}

// ListMemo implements MemoRepository.
func (m *MemoRepo) ListMemo() ([]*model.Memo, error) {
	var list []*model.Memo
	err := m.db.Find(&list).Error
	return list, err
}

func NewMemoRepo(db *gorm.DB) *MemoRepo {
	err := db.AutoMigrate(&model.Memo{})
	if err != nil {
		panic(err)
	}
	return &MemoRepo{
		db: db,
	}
}

// CreateMemo implements MemoRepository.
func (m *MemoRepo) CreateMemo(content string) error {
	return m.db.Create(&model.Memo{Content: content}).Error
}

// DeleteMemoById implements MemoRepository.
func (m *MemoRepo) DeleteMemoById(id uint) error {
	return m.db.Delete(&model.Memo{}, "id = ?", id).Error
}

// GetMemoById implements MemoRepository.
func (m *MemoRepo) GetMemoById(id uint) (*model.Memo, error) {
	memo := &model.Memo{}
	err := m.db.Take(memo, "id = ?", id).Error
	return memo, err
}

var _ MemoRepository = new(MemoRepo)
