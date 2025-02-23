package gofiber

import (
	config "github.com/max38/golang-clean-code-architecture/src/config"
	entities "github.com/max38/golang-clean-code-architecture/src/domain/entities"
	gofiberserver "github.com/max38/golang-clean-code-architecture/src/infrastructure/gofiber/server"
	postgresuserrepository "github.com/max38/golang-clean-code-architecture/src/interface/repositories/postgres/user"
	userusecase "github.com/max38/golang-clean-code-architecture/src/usecases/user"
)

// @title Training Kawaii Shop API
// @version 1.0
// @description This is a sample swagger for Training Kawaii Shop
// @termsOfService http://swagger.io/terms/
// @contact.name Sukhum Butrkam
// @contact.email sukhum_butrkam@hotmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	// Load environment variables from a configuration file
	config.Load(".env.dev")

	// Initialize repositories and use cases
	var userRepository = postgresuserrepository.UserRepository()
	var userUsecase = userusecase.UserUsecase(userRepository)

	// Create an application entity with the user usecase
	var applicationEntity = entities.ApplicationEntity{
		UserUsecase: userUsecase,
	}

	// Initialize the server and setup Swagger
	var server gofiberserver.IServer = gofiberserver.NewServer(applicationEntity)

	// Start the server
	server.Start()
}
