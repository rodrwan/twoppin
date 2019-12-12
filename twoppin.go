package twoppin

import "time"

type Model struct {
	ID string `json:"id" db:"id"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type Backpack struct {
	Model
	UserID string          `json:"user_id" db:"user_id"`
	Items  []*BackpackItem `json:"items" db:"items"`
}

type BackpackItem struct {
	Model
	Name string `json:"name" db:"name"`
	Qty  int    `json:"qty" db:"qty"`
}

type Lodging struct {
	Model
	Name      string    `json:"name" db:"name"`
	Contact   string    `json:"contact" db:"contact"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Address   string    `json:"address" db:"address"`
}

type Ticket struct {
	Model
	DepartureDate time.Time `json:"departure_date" db:"departure_date"`
	Price         int       `json:"price" db:"price"`
	Terminal      string    `json:"terminal" db:"terminal"`
}

type Trips struct {
	Model
	ReturnDate time.Time   `json:"return_date" db:"return_date"`
	Place      string      `json:"place" db:"place"`
	Tickets    []*Ticket   `json:"tickets" db:"tickets"`
	Lodgings   *Lodging    `json:"lodgings" db:"lodgings"`
	Backpacks  []*Backpack `json:"backpacks" db:"backpacks"`
}

type User struct {
	Model
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}
