package models

// TestGroup model
type TestGroup struct {
	ID    uint    `json:"id" gorm:"id;primaryKey;autoIncrement"`
	Name  string  `json:"name" gorm:"name"`
	Tests []*Test `json:"tests" gorm:"foreignKey:TestGroupID"`
}
