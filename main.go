package main

import (
	"flag"
	"myadmin/config"
	"myadmin/routes"
)

//	@title						My Admin Service Api
//	@version					1.0
//	@description				Api Service boilerplate dari golang gin dan gorm
//	@termsOfService				http://swagger.io/terms/
//	@contact.name				API Support
//	@contact.url				http://www.swagger.io/support
//	@contact.email				support@swagger.io
//	@securityDefinitions.apiKey	BearerAuth
//	@in							header
//	@name						Authorization
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//	@host						localhost:7000
//	@BasePath					/api/v1
//	@schemes					http
func main() {
	config.LoadEnv()
	db := config.InitialDB()

	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		config.InitCommands(db)
	} else {
		routes.ApiRouter(db)
	}
}
