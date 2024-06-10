package mysql

import (
	"strconv"

	"personal_page/model"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Memo struct {
	ID        int    `gorm:"size:10"`
	Content   string `gorm:"size:255;not null"`
	Completed bool   `gorm:"not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
	DeletedAt soft_delete.DeletedAt
}

func (m Memo) toModelMemo() *model.Memo {
	return &model.Memo{
		ID:        strconv.Itoa(m.ID),
		Content:   m.Content,
		Completed: m.Completed,
	}
}

type MemoRepo struct {
	db *gorm.DB
}

// CompleteWithId implements MemoRepository.
func (m *MemoRepo) CompleteWithId(id string) error {
	return m.db.Model(&Memo{}).Where("id = ?", id).UpdateColumn("completed", true).Error
}

// ListMemo implements MemoRepository.
func (m *MemoRepo) ListMemo() ([]*model.Memo, error) {
	var ml []*Memo
	err := m.db.Find(&ml).Error
	if err != nil {
		return nil, err
	}

	mml := make([]*model.Memo, 0, len(ml))
	for _, item := range ml {
		mml = append(mml, item.toModelMemo())
	}
	return mml, nil
}

// CreateMemo implements MemoRepository.
func (m *MemoRepo) CreateMemo(content string) error {
	return m.db.Create(&Memo{Content: content}).Error
}

// DeleteMemoById implements MemoRepository.
func (m *MemoRepo) DeleteMemoById(id string) error {
	return m.db.Delete(&Memo{}, "id = ?", id).Error
}

func NewMemoRepo(db *gorm.DB) *MemoRepo {
	err := db.AutoMigrate(&Memo{})
	if err != nil {
		panic(err)
	}
	return &MemoRepo{
		db: db,
	}
}
