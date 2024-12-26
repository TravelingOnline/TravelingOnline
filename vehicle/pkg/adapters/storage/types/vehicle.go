package types

type Vehicle struct {
	Id              string
	Unicode         string
	RequiredExperts int32
	Speed           int32
	RentPrice       int32
	IsActive        bool
	Type            string
	OwnerID         uint64 // Foreign key to Owner
	Owner           *Owner  `gorm:"foreignKey:OwnerID"` // Define the relationship
}

type Owner struct {
	Id        uint64 `gorm:"primaryKey"` // Primary key
	FirstName string
	LastName  string
	Email     string
}
