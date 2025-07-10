package repositories

import (
	"fmt"

	"github.com/AndersonBargas/rainstorm/v5"
)

type catalogRainstorm struct {
	dbInstance *rainstorm.DB
}

func NewCatalogsRainstorm(instance *rainstorm.DB) *catalogRainstorm {
	return &catalogRainstorm{
		dbInstance: instance,
	}
}

func (c *catalogRainstorm) FindByID(ID string) (*Catalog, error) {
	var catalog *Catalog
	err := c.dbInstance.One("ID", ID, catalog)
	if err != nil {
		return catalog, fmt.Errorf("Error finding the catalog by ID from the rainstorm database: %w", err)
	}
	return catalog, nil
}

func (c *catalogRainstorm) GetAll() ([]Catalog, error) {
	var catalogs []Catalog
	err := c.dbInstance.All(&catalogs)
	if err != nil {
		return catalogs, fmt.Errorf("Error getting all the catalogs from the rainstorm database: %w", err)
	}
	return catalogs, nil
}
