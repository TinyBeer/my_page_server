package model

type User struct {
	ID       string
	Name     string
	Password string
}

type Memo struct {
	ID        string
	Content   string
	Completed bool
}

// type Memo struct {
// 	Id        uint   `gorm:"size:10"`
// 	Content   string `gorm:"size:255;not null"`
// 	Completed bool   `gorm:"not null"`
// 	CreatedAt int64  `gorm:"autoCreateTime"`
// 	UpdatedAt int64  `gorm:"autoUpdateTime"`
// 	DeletedAt soft_delete.DeletedAt
// }

func (m *Memo) ToMemoItem() MemoItem {
	return MemoItem{
		ID:        m.ID,
		Content:   m.Content,
		Completed: m.Completed,
	}
}
