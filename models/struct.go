package models

import (
	"gorm.io/gorm"
)

// Group — группа студентов
type Group struct {
	*gorm.Model
	Title       string `json:"title"`
	CurrentWeek int    `json:"current_week"`
	TotalWeeks  int    `json:"total_weeks"`
	IsFinished  bool   `json:"is_finished"`

	Students []Student `gorm:"foreignKey:GroupID" json:"students"` // один ко многим
}

// Student — студент
type Student struct {
	*gorm.Model
	FullName      string `json:"full_name"`
	Email         string `json:"email" gorm:"uniqindex"`
	Telegram      string `json:"telegram" gorm:"uniqindex"`
	GroupID       uint   `json:"group_id"`
	TuitionTotal  int    `json:"tuition_total"`
	TuitionPaid   int    `json:"tuition_paid"`
	PaymentStatus string `json:"payment_status"` // "paid", "unpaid", "partial"
	StudyStatus   string `json:"study_status"`   // "learning", "job_search", "offer", "working"

	Group Group  `gorm:"foreignKey:GroupID" json:"-"`   // многие к одному
	Notes []Note `gorm:"foreignKey:StudentID" json:"-"` // один ко многим
}

// Note — заметка
type Note struct {
	*gorm.Model
	StudentID uint   `json:"student_id"`
	Author    string `json:"author"`
	Text      string `json:"text"`

	Student Student `gorm:"foreignKey:StudentID" json:"student"` // многие к одному
}

type InputGroup struct {
	Title       string `json:"title" binding:"required"`
	CurrentWeek int    `json:"current_week" binding:"required"`
	TotalWeeks  int    `json:"total_weeks" binding:"required"`
	IsFinished  bool   `json:"is_finished"`
}

type UpdateGroup struct {
	Title       *string `json:"title"`
	CurrentWeek *int    `json:"current_week"`
	TotalWeeks  *int    `json:"total_weeks"`
	IsFinished  *bool   `json:"is_finished"`
}

type StudentInput struct {
	FullName      string `json:"full_name" binding:"required"`
	Email         string `json:"email"  binding:"required"`
	Telegram      string `json:"telegram"  binding:"required"`
	GroupID       uint   `json:"group_id"  binding:"required"`
	TuitionTotal  int    `json:"tuition_total"  binding:"required"`
	TuitionPaid   int    `json:"tuition_paid"`
	PaymentStatus string `json:"payment_status"  binding:"required"`
	StudyStatus   string `json:"study_status"  binding:"required"`
}

type UpdateStudentInput struct {
	FullName      *string `json:"full_name"`
	Email         *string `json:"email"`
	Telegram      *string `json:"telegram"`
	GroupID       *uint   `json:"group_id"`
	TuitionTotal  *int    `json:"tuition_total"`
	TuitionPaid   *int    `json:"tuition_paid"`
	PaymentStatus *string `json:"payment_status"`
	StudyStatus   *string `json:"study_status"`
}

type InputNote struct {
	StudentID uint   `json:"student_id" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Text      string `json:"text" binding:"required"`
}

type UpdateNote struct {
	StudentID *uint   `json:"student_id"`
	Author    *string `json:"author"`
	Text      *string `json:"text"`
}
