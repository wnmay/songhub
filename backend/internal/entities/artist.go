package entities

type Artist struct {
	UserID uint `gorm:"primaryKey"`
	User   User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}