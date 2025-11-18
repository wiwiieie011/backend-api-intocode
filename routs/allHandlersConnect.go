package routs

import (
	"wiwieie011/controllers"

	"github.com/gin-gonic/gin"
)

func StudentsRout(c *gin.Engine) {
	s:=  c.Group("/students")
	{
	 s.GET("/",controllers.GetAllStudents)
	 s.GET("/:id", controllers.GetStudentsByID)
	 s.GET("/g_id", controllers.GetStudentsByGroupID)
	 s.GET("/p", controllers.GetPaidStudents)
	 s.GET("/st_s",controllers.GetStudentsStudyStatus)
	 s.POST("/", controllers.CreateStudent)
	 s.PATCH("/:id", controllers.PathStudent)
	 s.DELETE("/:id", controllers.DeleteStudentByID)
	}

	g:= c.Group("/group")
	{
		g.GET("/", controllers.GetGroups)
		g.GET("/c_w" ,controllers.GetGroupByWeek)
		g.GET("/f", controllers.GetFinishedGroups)
		g.GET("/:id", controllers.GetGroupByID)
		g.GET("/:id/stats/offer", controllers.GetOfferStats)
		g.GET("/finish",controllers.GetGroupByWeek)
		g.POST("/", controllers.CreateGroup)
		g.PATCH("/:id", controllers.UpdateGroupByID)
		g.DELETE("/:id", controllers.DeleteGroupByID)
	}

	n:= c.Group("/notes")
	{
		n.GET("/", controllers.GetNotes)	
		n.GET("/st_note/:id", controllers.GetNotesStudentsByID)	
		n.GET("/:id", controllers.GetNotesByID)
		n.POST("/", controllers.CreateNotes)
		n.PATCH("/:id", controllers.UpdateNotes)
		n.DELETE("/:id", controllers.DeleteNotes)	
	}
}

