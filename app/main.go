package main

import (
	_moviesUsecase "ticketing/business/movies"
	_moviesController "ticketing/controllers/movies"
	_moviesRepo "ticketing/drivers/databases/movies"
	_movieDB "ticketing/drivers/thirdparties/moviedb"

	_userUsecase "ticketing/business/users"
	_userController "ticketing/controllers/users"
	_userRepo "ticketing/drivers/databases/users"

	_ticketsUsecase "ticketing/business/tickets"
	_ticketsController "ticketing/controllers/tickets"
	_ticketsRepo "ticketing/drivers/databases/tickets"

	_theaterUsecase "ticketing/business/theater"
	_theaterController "ticketing/controllers/theater"
	_theaterRepo "ticketing/drivers/databases/theater"

	_topupUsecase "ticketing/business/topup"
	_topupController "ticketing/controllers/topup"
	_topupRepo "ticketing/drivers/databases/topup"

	_config "ticketing/app/config"
	_dbMongoDriver "ticketing/drivers/mongodb"
	_dbMysqlDriver "ticketing/drivers/mysql"

	_middleware "ticketing/app/middleware"
	_routes "ticketing/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func main() {
	configApp := _config.GetConfig()
	mysqlConfigDB := _dbMysqlDriver.ConfigDB{
		DB_Username: configApp.Mysql.User,
		DB_Password: configApp.Mysql.Pass,
		DB_Host:     configApp.Mysql.Host,
		DB_Port:     configApp.Mysql.Port,
		DB_Database: configApp.Mysql.Name,
	}
	mongoConfigDB := _dbMongoDriver.ConfigDB{
		DB_Username: configApp.Mongo.User,
		DB_Password: configApp.Mongo.Pass,
		DB_Host:     configApp.Mongo.Host,
		DB_Port:     configApp.Mongo.Port,
		DB_Database: configApp.Mongo.Name,
	}
	mysql_db := mysqlConfigDB.InitialMysqlDB()
	_ = mongoConfigDB.InitMongoDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _userRepo.NewMySQLUserRepository(mysql_db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	topupRepo := _topupRepo.NewMySQLTopUpRepository(mysql_db)
	topupUsecase := _topupUsecase.NewTopUpUsecase(topupRepo, timeoutContext, userRepo)
	topupCtrl := _topupController.NewTopUpController(topupUsecase)

	theaterRepo := _theaterRepo.NewMySQLTheaterRepository(mysql_db)
	theaterUsecase := _theaterUsecase.NewTheaterUsecase(theaterRepo, timeoutContext)
	theaterCtrl := _theaterController.NewTheaterController(theaterUsecase)

	ticketsRepo := _ticketsRepo.NewMySQLTicketsRepository(mysql_db)
	ticketsUsecase := _ticketsUsecase.NewTicketsUsecase(ticketsRepo, userRepo, timeoutContext)
	ticketsCtrl := _ticketsController.NewTicketsController(ticketsUsecase)

	MovieDBRepo := _movieDB.NewFetchMovies()
	moviesRepo := _moviesRepo.NewMySQLMoviesRepository(mysql_db)
	moviesUsecase := _moviesUsecase.NewMoviesUsecase(moviesRepo, timeoutContext, MovieDBRepo)
	moviesCtrl := _moviesController.NewMovieController(moviesUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userCtrl,
		TopUpController:   *topupCtrl,
		TheaterController: *theaterCtrl,
		MoviesController:  *moviesCtrl,
		TicketsController: *ticketsCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
