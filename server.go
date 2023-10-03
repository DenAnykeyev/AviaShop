package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	"io/ioutil"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	PhotoUrl    string `json:"photoUrl"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Rules    string `json:"rules"`
}

func main() {
	e := echo.New()

	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	file1, err := ioutil.ReadFile("products.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var products []Product
	err = json.Unmarshal(file1, &products)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	file2, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	var users []User
	err = json.Unmarshal(file2, &users)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return
	}

	e.GET("/get_all_products", func(c echo.Context) error {
		productsJSON, err := json.Marshal(products)
		if err != nil {
			fmt.Println("Ошибка кодирования JSON:", err)
			return err
		}

		return c.String(http.StatusOK, string(productsJSON))
	})

	e.GET("/get_products", func(c echo.Context) error {
		pageStr := c.QueryParam("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Некорректный параметр 'page'")
		}

		startIdx := (page - 1) * 3
		endIdx := startIdx + 3
		if startIdx >= len(products) {
			return c.String(http.StatusNotFound, "Страница не найдена")
		}
		if endIdx > len(products) {
			endIdx = len(products)
		}

		pageProducts := products[startIdx:endIdx]

		pageJSON, err := json.Marshal(pageProducts)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Ошибка преобразования данных в JSON")
		}

		return c.String(http.StatusOK, string(pageJSON))
	})

	e.GET("/get_total_pages", func(c echo.Context) error {
		totalPages := len(products) / 3
		if len(products)%3 != 0 {
			totalPages++
		}

		totalPagesJSON, err := json.Marshal(totalPages)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Ошибка преобразования данных в JSON")
		}

		return c.String(http.StatusOK, string(totalPagesJSON))
	})

	e.POST("/edit_product", func(c echo.Context) error {
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

		if index < 0 || index >= len(products) {
			return c.String(http.StatusNotFound, "Продукт не найден")
		}

		products[index].Name = newName
		products[index].Description = newDescription
		products[index].Price = newPrice

		if err := saveProductsToFile(products); err != nil {
			return c.String(http.StatusInternalServerError, "Не удалось сохранить данные")
		}

		return c.String(http.StatusOK, "Продукт успешно обновлен")
	})

	e.POST("/delete_product", func(c echo.Context) error {
		indexStr := c.FormValue("index")

		index, err := strconv.Atoi(indexStr)
		if err != nil {
			return c.String(http.StatusBadRequest, "Недопустимый индекс")
		}

		if index < 0 || index >= len(products) {
			return c.String(http.StatusNotFound, "Продукт не найден")
		}

		products = append(products[:index], products[index+1:]...)

		if err := saveProductsToFile(products); err != nil {
			return c.String(http.StatusInternalServerError, "Не удалось сохранить данные")
		}

		return c.String(http.StatusOK, "Продукт успешно удален")
	})

	e.POST("/add_product", func(c echo.Context) error {
		var newProduct Product
		err := c.Bind(&newProduct)
		if err != nil {
			return c.String(http.StatusBadRequest, "Ошибка чтения данных запроса")
		}

		products = append(products, newProduct)

		if err := saveProductsToFile(products); err != nil {
			return c.String(http.StatusInternalServerError, "Не удалось сохранить данные")
		}

		return c.String(http.StatusOK, "Продукт успешно добавлен")
	})

	store := sessions.NewCookieStore([]byte("secret_key"))

	e.Use(session.Middleware(store))

	e.POST("/register_user", func(c echo.Context) error {
		var newUser User
		if err := c.Bind(&newUser); err != nil {
			return c.String(http.StatusBadRequest, "Invalid registration data")
		}

		if userExists(newUser.Name, users) {
			return c.String(http.StatusBadRequest, "Пользователь с таким именем уже существует!")
		}

		users = append(users, newUser)

		if err := saveUserToFile(users); err != nil {
			return c.String(http.StatusInternalServerError, "Ошибка сохранения пользователя в базу данных")
		}

		sess, _ := session.Get("session", c)
		fmt.Print(c)

		sess.Values["name"] = newUser.Name
		sess.Values["rules"] = newUser.Rules

		if err := sess.Save(c.Request(), c.Response()); err != nil {
			return c.String(http.StatusInternalServerError, "Ошибка сохранения сессии")
		}

		return c.String(http.StatusOK, "Вы успешно зарегистированы!")
	})

	e.GET("/check_auth", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		name, ok1 := sess.Values["name"].(string)
		rules, ok2 := sess.Values["rules"].(string)

		if ok1 && ok2 {
			return c.JSON(http.StatusOK, map[string]interface{}{"isLoggedIn": true, "name": name, "rules": rules})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"isLoggedIn": false})
	})

	e.POST("/logout_user", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{MaxAge: -1}
		sess.Save(c.Request(), c.Response())
		return c.String(http.StatusOK, "Вы успешно вышли из аккаунта")
	})

	e.POST("/login_user", func(c echo.Context) error {
		var user User
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, "Invalid registration data")
		}

		username := user.Name
		password := user.Password

		var foundUser *User
		for _, user := range users {
			fmt.Println(user.Name)
			fmt.Println(username)
			fmt.Println(user.Name == username)
			if user.Name == username {
				foundUser = &user
				break
			}
		}

		if foundUser == nil {
			return c.String(http.StatusUnauthorized, "Неверное имя пользователя ")
		}

		if foundUser.Password != password {
			return c.String(http.StatusUnauthorized, "Неверный пароль")
		}

		sess, _ := session.Get("session", c)
		sess.Values["name"] = foundUser.Name
		sess.Values["rules"] = foundUser.Rules
		sess.Save(c.Request(), c.Response())

		return c.String(http.StatusOK, "Вход выполнен успешно")
	})

	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func saveProductsToFile(products []Product) error {
	productsJSON, err := json.Marshal(products)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("products.json", productsJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}

func userExists(name string, users []User) bool {
	for _, user := range users {
		if user.Name == name {
			return true
		}
	}
	return false
}

func saveUserToFile(users []User) error {
	usersJSON, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("users.json", usersJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
