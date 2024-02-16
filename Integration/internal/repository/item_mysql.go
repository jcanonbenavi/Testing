package repository

import (
	"database/sql"
	"testing/integration/internal"
)

type ItemMySQL struct {
	db *sql.DB
}

func NewItemMySQL(db *sql.DB) *ItemMySQL {
	return &ItemMySQL{db: db}
}

func (m *ItemMySQL) FindById(id int) (i internal.Item, err error) {
	row := m.db.QueryRow("SELECT id, name, description, price FROM items WHERE id = ?", id)
	err = row.Scan(&i.ID, &i.Name, &i.Description, &i.Price)
	if err != nil {
		switch {
		case err == sql.ErrNoRows:
			err = internal.ErrRepositoryNotFound
		default:
			err = internal.ErrRepositoryInternal
		}
		return
	}
	return
}
