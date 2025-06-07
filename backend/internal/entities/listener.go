package entities

type Listener struct {
	UserID uint `gorm:"primaryKey"`
	User   User `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
}