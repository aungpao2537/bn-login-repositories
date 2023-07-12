package main

import (

	// remove this before deploy

	"net/http"
	"os"

	ds "bn-login-repositories/domain/datasources"
	repo "bn-login-repositories/domain/repositories"
	"bn-login-repositories/src/gateways"

	// "bn-marketplace-repositories/src/gateways"
	sv "bn-login-repositories/src/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// // remove this before deploy ###################
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// /// ############################################

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	mongodb := ds.NewMongoDB(10)
	// redisdb := ds.NewRedisConnection()

	// firebase := repo.NewFirebaseRepository(mongodb)
	users := repo.NewUserRepository(mongodb)
	// login_redis := repo.NewRedisRepository(redisdb)

	// sv1 := sv.IFireBaseService(firebase)
	sv1 := sv.NewLoginService(users)

	gateways.NewHTTPGateway(e, sv1)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
