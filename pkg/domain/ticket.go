package domain

type Ticket struct {
	ID           uint `json:"id" gorm:"primaryKey;unique"`
	Entries      string
	TambolaSetID uint // Foreign key reference
}

type TambolaSet struct {
	ID      uint     `json:"id" gorm:"primaryKey;unique"`
	Tickets []Ticket `gorm:"foreignKey:TambolaSetID"` // Specify the foreign key relationship
}

// Ticket model
