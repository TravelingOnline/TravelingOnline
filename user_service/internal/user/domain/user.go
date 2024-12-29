package domain

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
	"user_service/internal/role/domain"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorWhileGeneratingHashPassword   = errors.New("faield to hash password")
	ErrorInvalidEmail                  = errors.New("invalid email format")
	ErrorInvalidNationalCodeDigits     = errors.New("national code must be exactly 10 digits")
	ErrorInvalidNationalCodeOnlyDigits = errors.New("national code must contain only digits")
	ErrorInvalidNationalCode           = errors.New("invalid national code")
	EmailRegex                         = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

type UserId uint

type User struct {
	ID           UserId
	FullName     string
	Email        string
	Password     string
	NationalCode string
	CreatedAt    time.Time
	DeletedAt    time.Time
	Role         domain.Role
	RoleId       domain.RoleId
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func ValidateNationalCode(code string) error {
	if len(code) != 10 {
		return ErrorInvalidNationalCodeDigits
	}

	if _, err := strconv.Atoi(code); err != nil {
		return ErrorInvalidNationalCodeOnlyDigits
	}

	if code == "0000000000" {
		return ErrorInvalidNationalCode
	}
	if code == "9876543210" {
		return ErrorInvalidNationalCode
	}
	digits := make([]int, 10)
	for i, c := range code {
		digits[i], _ = strconv.Atoi(string(c))
	}

	var B int
	for i := 0; i < 9; i++ {
		B += digits[i] * (10 - i)
	}

	C := B - (B/11)*11
	A := digits[9]

	if C == 0 && A == C {
		return nil
	}
	if C == 1 && A == 1 {
		return nil
	}
	if C > 1 && A == (11-C) {
		return nil
	}

	return ErrorInvalidNationalCode
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", ErrorWhileGeneratingHashPassword
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
func LowerCaseEmail(email string) string {
	return strings.ToLower(email)
}
func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)
	if !EmailRegex.MatchString(email) {
		return ErrorInvalidEmail
	}
	return nil
}

func ValidateUserRegisteration(user *User) error {
	err := ValidateEmail(user.Email)
	if err != nil {
		return err
	}

	err = ValidateNationalCode(user.NationalCode)
	if err != nil {
		return err
	}
	return nil

}
