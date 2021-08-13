package main

import (
	_moviesUsecase "ticketing/business/movies"
	_moviesController "ticketing/controllers/movies"
	_moviesRepo "ticketing/drivers/databases/movies"
	_movieDB "ticketing/drivers/thirdparties/moviedb"

	_userUsecase "ticketing/business/users"
	_userController "ticketing/controllers/users"
	_userRepo "ticketing/drivers/databases/users"

	_theaterUsecase "ticketing/business/theater"
	_theaterController "ticketing/controllers/theater"
	_theaterRepo "ticketing/drivers/databases/theater"

	_topupUsecase "ticketing/business/topup"
	_topupController "ticketing/controllers/topup"
	_topupRepo "ticketing/drivers/databases/topup"

	_dbDriver "ticketing/drivers/mysql"

	_middleware "ticketing/app/middleware"
	_routes "ticketing/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	// topupUserRepo := _userRepo.NewMySQLUserRepository(db)
	topupRepo := _topupRepo.NewMySQLTopUpRepository(db)
	topupUsecase := _topupUsecase.NewTopUpUsecase(topupRepo, timeoutContext, userRepo)
	topupCtrl := _topupController.NewTopUpController(topupUsecase)

	theaterRepo := _theaterRepo.NewMySQLTheaterRepository(db)
	theaterUsecase := _theaterUsecase.NewTheaterUsecase(theaterRepo, timeoutContext)
	theaterCtrl := _theaterController.NewTheaterController(theaterUsecase)

	MovieDBRepo := _movieDB.NewFetchMovies()
	moviesRepo := _moviesRepo.NewMySQLMoviesRepository(db)
	moviesUsecase := _moviesUsecase.NewMoviesUsecase(moviesRepo, timeoutContext, MovieDBRepo)
	moviesCtrl := _moviesController.NewMovieController(moviesUsecase)
	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
		TopUpController:   *topupCtrl,
		TheaterController: *theaterCtrl,
		MoviesController:  *moviesCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
