module github.com/draganaF/gled.io/UserService

go 1.18

require gorm.io/gorm v1.23.8

require github.com/go-sql-driver/mysql v1.6.0 // indirect

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/rs/cors v1.8.2
	gorm.io/driver/mysql v1.3.6
)
