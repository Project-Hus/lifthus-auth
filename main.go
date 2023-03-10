package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

// @title Project-Hus auth server
// @version 0.0.0
// @description This is Project-Hus's root authentication server containing each user's UUID, which is unique for all hus services.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url lifthus531@gmail.com
// @contact.email lifthus531@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host lifthus.com
// @BasePath /auth
func main() {
	// set .env
	err := godotenv.Load() // now you can use os.Getenv("VAR_NAME")
	if err != nil {
		log.Fatalf("[F] loading .env file failed : %s", err)
	}

	// connecting to hus_auth_db with ent
	/*
		client, err := db.ConnectToHusAuth()
		if err != nil {
			log.Fatal("%w", err)
		}
		defer client.Close()

		// Run the auto migration tool.
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("[F] creating schema resources failed : %v", err)
		}
	*/

	// subdomains
	hosts := map[string]*Host{}

	//  Create echo web server instance and set CORS headers
	e := echo.New()
	//e.Use(middleware.SetHusCorsHeaders)

	// authApi, which controls auth all over the services
	//authApi := auth.NewAuthApiController(client)
	//hosts["localhost:9090"] = &Host{Echo: authApi} // gonna use auth.cloudhus.com later

	// get requset and process by its subdomain
	e.Any("/*", func(c echo.Context) (err error) {
		req, res := c.Request(), c.Response()
		host, ok := hosts[req.Host] // if the host is not registered, it will be nil.
		if !ok {
			return c.NoContent(http.StatusNotFound)
		} else {
			host.Echo.ServeHTTP(res, req)
		}
		return err
	})

	// provide api docs with swagger 2.0
	//e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Run the server
	e.Logger.Fatal(e.Start(":9090"))
}

type Host struct {
	Echo *echo.Echo
}
