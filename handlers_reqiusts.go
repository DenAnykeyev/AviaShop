package main

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

// Функция обработки запроса на получение всех продуктов
func getAllProductsHandler(c echo.Context, db *sql.DB) error {
	products, err := getAllProductsFromDB(db)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при получении продуктов из базы данных")
	}
	return c.JSON(http.StatusOK, products)
}

// Функция обработки запроса на получение определенных продуктов
func getProductsHandler(c echo.Context, db *sql.DB) error {
	pageStr := c.QueryParam("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Некорректный параметр 'page'")
	}

	pageSize := 3

	products, err := getProductsFromDB(db, page, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при получении продуктов из базы данных")
	}

	return c.JSON(http.StatusOK, products)
}

// Функция обработки запроса на получение количества страницы
func getTotalPagesHandler(c echo.Context, db *sql.DB) error {
	pageSize := 3

	totalPages, err := getTotalPagesFromDB(db, pageSize)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при получении общего числа страниц из базы данных")
	}

	return c.JSON(http.StatusOK, totalPages)
}

// Функция обработки запроса на редактирование продукта
func editProductHandler(c echo.Context, db *sql.DB) error {
	indexStr := c.FormValue("index")
	newName := c.FormValue("name")
	newDescription := c.FormValue("description")
	newPriceStr := c.FormValue("price")

	newPrice, err := strconv.Atoi(newPriceStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Недопустимая цена")
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Недопустимый индекс")
	}

	if err := editProductInDB(db, index, newName, newDescription, newPrice); err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось обновить продукт в базе данных")
	}

	return c.String(http.StatusOK, "Продукт успешно обновлен")
}

// Функция обработки запроса на удаление продукта
func deleteProductHandler(c echo.Context, db *sql.DB) error {
	indexStr := c.FormValue("index")

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Недопустимый индекс")
	}

	if err := deleteProductInDB(db, index); err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось удалить продукт из базы данных")
	}

	return c.String(http.StatusOK, "Продукт успешно удален")
}

// Функция обработки запроса на добавление продукта
func addProductHandler(c echo.Context, db *sql.DB) error {
	var newProduct Product
	err := c.Bind(&newProduct)
	if err != nil {
		return c.String(http.StatusBadRequest, "Ошибка чтения данных запроса")
	}

	if err := addProductToDB(db, newProduct); err != nil {
		return c.String(http.StatusInternalServerError, "Не удалось добавить продукт в базу данных")
	}

	return c.String(http.StatusOK, "Продукт успешно добавлен в базу данных")
}

// Функция обработки запроса на регистрацию пользователя
func registerUserHandler(c echo.Context, db *sql.DB) error {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.String(http.StatusBadRequest, "Что-то не так...")
	}

	// Insert the new user into the database
	result, err := addUserToDB(db, newUser)
	if err != nil {
		if strings.Contains(err.Error(), "уже существует") {
			return c.String(http.StatusBadRequest, "Пользователь с таким именем уже существует!")
		}
		return c.String(http.StatusInternalServerError, "Ошибка при сохранении пользователя в базе данных")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка при получении ID пользователя")
	}

	sess, _ := session.Get("session", c)

	sess.Values["id"] = result
	sess.Values["name"] = newUser.Name
	sess.Values["rules"] = newUser.Rules

	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.String(http.StatusInternalServerError, "Ошибка сохранения сессии")
	}

	return c.String(http.StatusOK, "Вы успешно зарегистрированы!")
}

// Функция обработки запроса на проверку авторизации пользователя
func checkAuthUserHandler(c echo.Context, db *sql.DB) error {
	sess, _ := session.Get("session", c)
	name, ok1 := sess.Values["name"].(string)
	rules, ok2 := sess.Values["rules"].(string)

	type resp struct {
		IsLoggedIn bool   `json:"isLoggedIn"`
		Name       string `json:"name"`
		Rules      string `json:"rules"`
	}

	if ok1 && ok2 {
		return c.JSON(http.StatusOK, resp{IsLoggedIn: true, Name: name, Rules: rules})
	}
	return c.JSON(http.StatusOK, resp{IsLoggedIn: false})
}

// Функция обработки запроса на выход пользователя из аккаунта
func logoutUserHandler(c echo.Context, db *sql.DB) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{MaxAge: -1}
	sess.Save(c.Request(), c.Response())
	return c.String(http.StatusOK, "Вы успешно вышли из аккаунта")
}

// Функция обработки запроса на авторизацию пользователя
func loginUserHandler(c echo.Context, db *sql.DB) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Что-то пошло не так")
	}

	dbUser, err := loginUserFromDB(db, user.Name, user.Password)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}

	sess, _ := session.Get("session", c)

	sess.Values["id"] = dbUser.Id
	sess.Values["name"] = dbUser.Name
	sess.Values["rules"] = dbUser.Rules

	sess.Save(c.Request(), c.Response())

	return c.String(http.StatusOK, "Вход выполнен успешно")
}

/*
func addProductToBasketHandler(c echo.Context, db *sql.DB) error {
	// Разберите идентификатор товара из запроса
	var request struct {
		ProductID int `json:"productId"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Неверный запрос"})
	}

	// Проверьте, аутентифицирован ли пользователь (вы можете реализовать эту логику)
	// Например, вы можете получить информацию о пользователе из сессии

	// Вставьте товар в таблицу корзины
	_, err := db.Exec("INSERT INTO baskets (user_id, product_id) VALUES (?, ?)", userID, request.ProductID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Не удалось добавить товар в корзину"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Товар добавлен в корзину"})
}
*/
