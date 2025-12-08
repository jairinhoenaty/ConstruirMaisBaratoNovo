package job

import "time"

type Job struct {
	ID               uint       `gorm:"primaryKey" json:"id"`
	Title            string     `json:"title"`
	HiringType       string     `json:"hiring_type"`
	Salary           float64    `json:"salary"`
	SalaryType       string     `json:"salary_type"`
	Location         string     `json:"location"`
	Description      string     `json:"description"`
	Schedule         string     `json:"schedule"`
	Requirements     string     `json:"requirements"`
	Benefits         string     `json:"benefits"`
	ContactEmail     string     `json:"contact_email"`
	ContactPhone     string     `json:"contact_phone"`
	OpeningsQuantity int        `json:"openings_quantity"`
	Status           string     `json:"status"`
	ProfessionID     uint       `json:"profession_id"`
	CityID           uint       `json:"city_id"`
	PublishedAt      *time.Time `json:"published_at"`
	ExpiresAt        *time.Time `json:"expires_at"`
	Approved         bool       `json:"approved"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `gorm:"index" json:"deleted_at"`
}
