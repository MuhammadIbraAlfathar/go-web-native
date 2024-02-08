package categorymodel

import (
	"database/sql"
	"github.com/MuhammadIbraAlfathar/go-web-native/entities"
)

func GetAll(db *sql.DB) []entities.Category {

	rows, err := db.Query("SELECT id, name from categories")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(
			&category.Id,
			&category.Name,
		)

		if err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories
}

func CreateCategory(db *sql.DB, category entities.Category) {
	_, err := db.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	if err != nil {
		panic(err)
	}
}

func DetailCategory(id int, db *sql.DB) entities.Category {
	row := db.QueryRow("SELECT id, name from categories where id = ?", id)

	var category entities.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		panic(err)
	}

	return category
}
