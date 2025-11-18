package main

import (
	"os"
	"wiwieie011/config"
	"wiwieie011/routs"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	config.ConnectionDB()
	r:= gin.Default()
	routs.StudentsRout(r)
	s:= os.Getenv("PORT")
	r.Run(":"+s)
}