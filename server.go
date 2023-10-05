package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Id          int    `json:"id"`
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

	// Инциализация соединения с базой данных
	db, err := initDB()
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных")
		return
	}
	defer db.Close()

	// Сессии
	store := sessions.NewCookieStore([]byte("secret_key"))
	e.Use(session.Middleware(store))

	e.GET("/get_all_products", func(c echo.Context) error {
		return getAllProductsHandler(c, db)
	})

	e.GET("/get_products", func(c echo.Context) error {
		return getProductsHandler(c, db)
	})

	e.GET("/get_total_pages", func(c echo.Context) error {
		return getTotalPagesHandler(c, db)
	})

	e.POST("/edit_product", func(c echo.Context) error {
		return editProductHandler(c, db)
	})

	e.POST("/delete_product", func(c echo.Context) error {
		return deleteProductHandler(c, db)
	})

	e.POST("/add_product", func(c echo.Context) error {
		return addProductHandler(c, db)
	})

	e.POST("/register_user", func(c echo.Context) error {
		return registerUserHandler(c, db)
	})

	e.GET("/check_auth", func(c echo.Context) error {
		return checkAuthUserHandler(c, db)
	})

	e.POST("/logout_user", func(c echo.Context) error {
		return logoutUserHandler(c, db)
	})

	e.POST("/login_user", func(c echo.Context) error {
		return loginUserHandler(c, db)
	})

	e.POST("/add_product_in_basket", func(c echo.Context) error {
		return nil //addProductToBasketHandler(c, db)
	})

	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/shop")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func registerReqiuestHandlers() {

}
