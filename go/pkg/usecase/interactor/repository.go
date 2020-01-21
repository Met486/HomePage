package interactor

import (
	"homepage/pkg/domain"
	"time"
)

// AccountRepository dbにつなぐ。実装はinterface > *_repositoryで
type AccountRepository interface {
	FindAccountByUserID(userID int) (domain.User, error)
	FindAccountByStudentID(studentID string) (user domain.User, err error)
	StoreAccount(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error
	UpdateAccount(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error
	DeleteAccount(userID int) error
}

// UserRepository dbにつなぐ。実装はinterface > *_repositoryで
type UserRepository interface {
	FindUsers() (domain.Users, error)
	FindUserByUserID(userID int) (domain.User, error)
	StoreUser(name, password, role, studentID, department, comment string, grade int, createdAt time.Time) error
	UpdateUser(userID int, name, password, role, studentID, department, comment string, grade int, updatedAt time.Time) error
	DeleteUser(userID int) error
}

// ActivityRepository dbにつなぐ。実装は interface > *_repository
type ActivityRepository interface {
	FindActivities() (domain.Activities, error)
	FindActivityByID(actID int) (domain.Activity, error)
	StoreActivity(date time.Time, act string, createdAt time.Time) (int, error)
	UpdateActivity(actID int, date time.Time, act string, updatedAt time.Time) error
	DeleteActivity(actID int) error
}

// SocietyRepository dbにつなぐ。実装は interface > *_repository
type SocietyRepository interface {
	FindSocieties() (domain.Societies, error)
	FindSocietyByID(socID int) (domain.Society, error)
	StoreSociety(title, author, society, award string, date, createdAt time.Time) (int, error)
	UpdateSociety(socID int, title, author, society, award string, date, updatedAt time.Time) error
	DeleteSociety(socID int) error
}
