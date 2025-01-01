package domain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidEmail          = errors.New("invalid email format")
	ErrInvalidPassword       = errors.New("invalid password format")
	ErrEmailAlreadyExists    = errors.New("email already exists")
	ErrInvalidAuthentication = errors.New("email and password doesn't match")
)

type UserFilter struct {
	ID    UserID
	Phone string
	Email string
}
type (
	UserID       uuid.UUID
	Phone        string
	NationalCode string
	Email        string
)

func (u UserID) ConvStr() string {
	return uuid.UUID(u).String()
}
func (p Phone) IsValid() bool {
	// Define the regex pattern for Iranian mobile numbers
	// Iranian mobile numbers start with +98 or 0, followed by 9 and 9 more digits
	regexPattern := `^(?:\+98|0)?9\d{9}$`

	// Compile the regex
	regex := regexp.MustCompile(regexPattern)

	// Validate the phone number using the regex
	return regex.MatchString(string(p))
}

func (nc NationalCode) IsValid() bool {
	// Check if the national code is a 10-digit number
	regex := regexp.MustCompile(`^\d{10}$`)
	if !regex.MatchString(string(nc)) {
		return false
	}

	// Check if all digits are the same (e.g., "1111111111")
	if isAllSameDigits(string(nc)) {
		return false
	}

	// Perform checksum validation
	return validateChecksum(string(nc))
}

// Helper function to check if all digits are the same
func isAllSameDigits(nc string) bool {
	for i := 1; i < len(nc); i++ {
		if nc[i] != nc[0] {
			return false
		}
	}
	return true
}

// Helper function to validate checksum
func validateChecksum(nc string) bool {
	// Convert the first 9 digits to integers
	sum := 0
	for i := 0; i < 9; i++ {
		digit, _ := strconv.Atoi(string(nc[i]))
		sum += digit * (10 - i)
	}

	// Calculate the remainder
	remainder := sum % 11

	// Get the control digit (the last digit)
	controlDigit, _ := strconv.Atoi(string(nc[9]))

	// Validate based on the rules
	return (remainder < 2 && remainder == controlDigit) || (remainder >= 2 && controlDigit == 11-remainder)
}
func (e Email) IsValid() bool {
	// Define a regex pattern for validating email addresses
	// This pattern checks for a basic structure of local@domain
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regex
	regex := regexp.MustCompile(regexPattern)

	// Validate the email using the regex
	return regex.MatchString(string(e))
}

type User struct {
	ID           uuid.UUID
	Email        Email
	Password     string
	IsSuperAdmin bool
	IsAdmin      bool
	//Roles        []Role
	IsBlocked bool
	CreatedAt time.Time
	DeletedAt time.Time
	UpdatedAt time.Time
}

func (u *User) SetPassword(password string) {
	u.Password = password
}

func (u *User) PasswordIsValid(pass string) bool {
	h := sha256.New()
	h.Write([]byte(pass))
	passSha256 := h.Sum(nil)
	return fmt.Sprintf("%x", passSha256) == u.Password
}

func ValidateEmail(email string) error {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	isMatched := emailRegex.MatchString(email)
	if !isMatched {
		return ErrInvalidEmail
	}
	return nil
}

func ValidatePasswordWithFeedback(password string) error {
	tests := []struct {
		pattern string
		message string
	}{
		{".{7,}", "Password must be at least 7 characters long"},
		{"[a-z]", "Password must contain at least one lowercase letter"},
		{"[A-Z]", "Password must contain at least one uppercase letter"},
		{"[0-9]", "Password must contain at least one digit"},
		{"[^\\d\\w]", "Password must contain at least one special character"},
	}

	var errMessages []string

	for _, test := range tests {
		match, _ := regexp.MatchString(test.pattern, password)
		if !match {
			errMessages = append(errMessages, test.message)
		}
	}

	if len(errMessages) > 0 {
		feedback := strings.Join(errMessages, "\n")
		return errors.Join(ErrInvalidPassword, fmt.Errorf(feedback))
	}

	return nil
}

func LowerCaseEmail(email string) string {
	return strings.ToLower(email)
}

//type Role struct {
//	ID          uint
//	Name        string `gorm:"uniqueIndex;not null"`
//	Description string
//	Permissions []Permission `gorm:"many2many:role_permissions;"`
//}
//
//type Permission struct {
//	ID          uint
//	Name        string `gorm:"uniqueIndex;not null"`
//	Description string
//}
//
//type RolePermission struct {
//	ID           uint
//	RoleID       uint
//	PermissionID uint
//}

func (u *User) Validate() error {

	if !u.Email.IsValid() {
		return errors.New("email is not valid")
	}

	return nil
}
