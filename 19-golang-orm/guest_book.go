package gorm

type GuestBook struct {
	ID        int64  `gorm:"column:id;primary_key;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Message   string `gorm:"column:message"`
	CreatedAt string `gorm:"column:created_at;autoCreatedTime"`
	UpdateAt  string `gorm:"column:updated_at;autoCreatedTime;autoUpdatedTime"`
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}
