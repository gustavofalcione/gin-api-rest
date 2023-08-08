package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gustavofalcione/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Substitua pelo domínio do seu frontend
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	r.Use(cors.New(config))

	r.GET("/students", controllers.ShowAllStudents)
	r.POST("/students", controllers.CreateStudent)
	r.GET("/students/:id", controllers.FindStudentById)
	r.DELETE("/students/:id", controllers.Delete)
	r.PATCH("/students/:id", controllers.Update)
	r.GET("/students/cpf/:cpf", controllers.FindStudentByCpf)

	r.GET("/:name", controllers.Greetings)
	r.Run()
}
