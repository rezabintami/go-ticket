module ticketing

// +heroku goVersion go1.16
go 1.16

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/google/uuid v1.3.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/midtrans/midtrans-go v1.2.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.8.1
	github.com/stretchr/testify v1.7.0
	go.mongodb.org/mongo-driver v1.7.1
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gorm.io/driver/mysql v1.1.1
	gorm.io/gorm v1.21.12
)
