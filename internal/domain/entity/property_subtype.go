package entity

type PropertySubtype struct {
	ID          string       `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string       `gorm:"not null" json:"name"`        // snake_case code
	DisplayName string       `gorm:"not null" json:"displayName"` // Human readable
	Type        PropertyType `gorm:"type:varchar(20);not null;index" json:"type"`
	Active      bool         `gorm:"default:true" json:"active"`
}
