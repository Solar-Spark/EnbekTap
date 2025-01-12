package entities

type Vacancy struct {
	VacancyID   int64  `gorm:"primaryKey;autoIncrement"`
	Vacancy     string `gorm:"type:varchar(255);column:vacancy"`
	Salary      int    `gorm:"column:salary"`
	JobType     string `gorm:"type:enum('full-time', 'part-time');column:jobtype"`
	Description string `gorm:"type:varchar(255);column:description"`
}
