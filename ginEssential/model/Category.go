package model

type Category struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time   `json:"create_at" gorm:"type:timestamp"`
	UpdateAt  Time   `json:"updated_at" gorm:"type:timestamp"`
}
