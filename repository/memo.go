package repository

import (
	"personal_page/model"

	"gorm.io/gorm"
)

type MemoRepo struct {
	db *gorm.DB
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

// UpdateMemoById implements MemoRepository.
func (m *MemoRepo) UpdateMemoById(id uint, content string) error {
	return m.db.Model(&model.Memo{}).Where("id = ?", id).Update("content", content).Error
}

var _ MemoRepository = new(MemoRepo)
