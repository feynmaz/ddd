package aggregate

import (
	"errors"

	"github.com/feynmaz/ddd/entity"
	"github.com/feynmaz/ddd/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("customer has to have a valid name")
)

type Customer struct {
	// Person is the root entity of the customer
	// which means person.ID is the main identifier for the Customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer is a factory to create a new cutomer aggregates
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{
			ID: id,
		}
	}
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{
			Name: name,
		}
	}
}


