module go_http

go 1.16

require (
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.3.4
	gorm.io/gorm v1.23.6
)

replace MyGo_middleware => ./../MyGo_middleware

require (
	MyGo_middleware v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.8.1
)
