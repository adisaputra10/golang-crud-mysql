package main

import (
	"belajar/app"
	"belajar/config"

	"github.com/labstack/echo"
)

type M map[string]interface{}

func main() {

	// INIT ECHO
	//e := echo.New()
	// ROUTING

	//e.GET("/", welcome)
	config.ConnectDB()

	r := echo.New()

	go app.GetApiAdd(r)
	go app.GetApiUpdate(r)
	go app.GetApiDelete(r)
	go app.GetApi(r)
	//
	apiGroup := r.Group("/user")
	go app.UseSubGroup(apiGroup)

	// r.GET("/", app.Welcome)

	// cookieGroup := r.Group("/group")
	// cookieGroup.GET("/", app.Welcome)
	// cookieGroup.GET("/main", app.GroupApi)
	// r.Validator = &CustomValidator{validator: validator.New()}

	// r.GET("/", func(ctx echo.Context) error {
	// 	data := "Hello from /index"
	// 	return ctx.String(http.StatusOK, data)
	// })

	// r.POST("/users", func(ctx echo.Context) error {
	// 	u := new(User)
	// 	if err := ctx.Bind(u); err != nil {
	// 		res := M{"Message": "Not Found", "Counter": 2}
	// 		return ctx.JSON(http.StatusNotFound, res)
	// 	}
	// 	if err := ctx.Validate(u); err != nil {
	// 		res := M{"Message": "Not Found", "Counter": 1}
	// 		return ctx.JSON(http.StatusNotFound, res)
	// 	}
	// 	return ctx.JSON(http.StatusOK, true)

	// })

	// r.GET("/json", func(ctx echo.Context) error {
	// 	data := M{"Message": "Hello", "Counter": 2}
	// 	return ctx.JSON(http.StatusOK, data)
	// })

	//	r.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//	Format: "method=${method}, uri=${uri}, status=${status}\n",
	//	}))

	r.Logger.Fatal(r.Start(":9000"))
}
