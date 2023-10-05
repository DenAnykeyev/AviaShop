package main

import (
	"database/sql"
	"errors"
	"fmt"
)

func getAllProductsFromDB(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT id, name, description, price, photoUrl FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.PhotoUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func getProductsFromDB(db *sql.DB, page int, pageSize int) ([]Product, error) {
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("SELECT name, description, price, photoUrl FROM products LIMIT %d, %d", offset, pageSize)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Name, &product.Description, &product.Price, &product.PhotoUrl)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func getTotalPagesFromDB(db *sql.DB, pageSize int) (int, error) {
	var totalProducts int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&totalProducts)
	if err != nil {
		return 0, err
	}

	totalPages := totalProducts / pageSize
	if totalProducts%pageSize != 0 {
		totalPages++
	}

	return totalPages, nil
}

func editProductInDB(db *sql.DB, index int, newName string, newDescription string, newPrice int) error {
	if index < 0 {
		return errors.New("Недопустимый индекс")
	}

	var productID int
	err := db.QueryRow("SELECT id FROM products LIMIT 1 OFFSET ?", index).Scan(&productID)
	if err == sql.ErrNoRows {
		return errors.New("Продукт не найден")
	} else if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE products SET name=?, description=?, price=? WHERE id=?", newName, newDescription, newPrice, productID)
	if err != nil {
		return err
	}

	return nil
}

func deleteProductInDB(db *sql.DB, id int) error {
	// Удаляем запись с заданным id из таблицы
	query := "DELETE FROM products WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	// Обновляем значения столбца id для оставшихся записей в таблице
	_, err = db.Exec("SET @new_id := 0")
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE products SET id = (@new_id := @new_id + 1) ORDER BY id")
	if err != nil {
		return err
	}

	return nil
}

func addProductToDB(db *sql.DB, newProduct Product) error {
	_, err := db.Exec("INSERT INTO products (name, description, price, photoUrl) VALUES (?, ?, ?, ?)",
		newProduct.Name, newProduct.Description, newProduct.Price, newProduct.PhotoUrl)
	if err != nil {
		return err
	}

	return nil
}

func addUserToDB(db *sql.DB, newUser User) error {
	var existingUser User
	err := db.QueryRow("SELECT * FROM users WHERE name=?", newUser.Name).Scan(&existingUser.Name, &existingUser.Password, &existingUser.Rules)

	if err == nil {
		return fmt.Errorf("Пользователь с именем %s уже существует", newUser.Name)
	} else if err != sql.ErrNoRows {
		return err
	}

	_, err = db.Exec("INSERT INTO users (name, password, rules) VALUES (?, ?, ?)",
		newUser.Name, newUser.Password, newUser.Rules)
	if err != nil {
		return err
	}

	return nil
}

func loginUserFromDB(db *sql.DB, username, password string) (*User, error) {
	var user User
	err := db.QueryRow("SELECT name, rules FROM users WHERE name = ? AND password = ?", username, password).Scan(&user.Name, &user.Rules)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Неверное имя пользователя или пароль")
		}
		return nil, fmt.Errorf("Ошибка при выполнении запроса")
	}
	return &user, nil
}
