package controllers

import (
	"net/http"
	"strconv"
	"wiwieie011/config"
	"wiwieie011/models"

	"github.com/gin-gonic/gin"
)

func GetAllStudents(c *gin.Context) {
	var s []models.Student
	err := config.DB.Preload("Group").Find(&s).Error
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"studentsList": s})
}

func GetStudentsByID(c *gin.Context) {
	var s models.Student

	err := config.DB.Preload("Group").First(&s, c.Param("id")).Error
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, s)
}

func GetStudentsByGroupID(c *gin.Context) {
	groupID := c.Query("group_id")
	if groupID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "group_id is required"})
		return
	}

	var students []models.Student

	if err := config.DB.Where("group_id = ?", groupID).Preload("Group").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func GetPaidStudents(c *gin.Context) {
	status := c.Query("payment_status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment_status is required"})
		return
	}

	var students []models.Student

	if err := config.DB.Where("payment_status = ?", status).Preload("Group").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func GetStudentsStudyStatus(c *gin.Context) {
	status := c.Query("study_status")
	if status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "study_status is required"})
		return
	}

	var students []models.Student

	if err := config.DB.Where("study_status = ?", status).Preload("Group").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

func CreateStudent(c *gin.Context) {
	var inputStudent models.StudentInput

	if err := c.ShouldBindJSON(&inputStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if inputStudent.PaymentStatus != "paid" &&
		inputStudent.PaymentStatus != "unpaid" &&
		inputStudent.PaymentStatus != "partial" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment_status"})
		return
	}

	if inputStudent.StudyStatus != "learning" &&
		inputStudent.StudyStatus != "job_search" &&
		inputStudent.StudyStatus != "offer" &&
		inputStudent.StudyStatus != "working" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid study_status"})
		return
	}

	student := models.Student{
		FullName:      inputStudent.FullName,
		Email:         inputStudent.Email,
		Telegram:      inputStudent.Telegram,
		GroupID:       inputStudent.GroupID,
		TuitionTotal:  inputStudent.TuitionTotal,
		TuitionPaid:   inputStudent.TuitionPaid,
		PaymentStatus: inputStudent.PaymentStatus,
		StudyStatus:   inputStudent.StudyStatus,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"created": true})
}

func PathStudent(c *gin.Context) {
	var student models.Student

	if err := config.DB.First(&student, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updateStudent models.UpdateStudentInput
	if err := c.ShouldBindJSON(&updateStudent); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	if err := config.DB.Model(&student).Updates(updateStudent).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"updates": true})
}

func DeleteStudentByID(c *gin.Context) {
	var student models.Student

	if err := config.DB.Unscoped().Delete(&student, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"deleted": true})
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func GetGroups(c *gin.Context) {
	var groups []models.Group
	if err := config.DB.Preload("Students").Find(&groups).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"find": groups})
}

func GetGroupByWeek(c *gin.Context) {
	week := c.Query("current_week")
	if week == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "current_week is required"})
		return
	}

	var group []models.Group

	if err := config.DB.Where("current_week = ?", week).Find(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, group)
}

func GetFinishedGroups(c *gin.Context) {

	finished := c.Query("is_finished")
	if finished == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "is_finished is required"})
		return
	}

	var group []models.Group

	if err := config.DB.Where("is_finished = ?", finished).Find(&group).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, group)
}


func GetGroupswww(c *gin.Context) {
    finished := c.Query("finished")

    var groups []models.Group

    if finished == "true" {
        if err := config.DB.Where("is_finished = ? OR current_week >= total_weeks", true).Preload("Students").Find(&groups).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
    }

    c.JSON(http.StatusOK, groups)
}


func GetGroupByID(c *gin.Context) {
	var group models.Group
	if err := config.DB.Preload("Students").First(&group, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"find": group})
}

func GetOfferStats(c *gin.Context) {
	groupIDStr := c.Param("id")
	groupID, err := strconv.Atoi(groupIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid group id"})
		return
	}

	var students []models.Student
	if err := config.DB.Where("group_id = ?", groupID).Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	if len(students) == 0 {
		c.JSON(http.StatusOK, gin.H{"offer_percent": 0})
		return
	}

	offerCount := 0
	for _, s := range students {
		if s.StudyStatus == "offer" {
			offerCount++
		}
	}

	percent := float64(offerCount) * 100.0 / float64(len(students))

	c.JSON(http.StatusOK, gin.H{"group_id": groupID, "offer_percent": percent})
}

func CreateGroup(c *gin.Context) {
	var InputGroup models.InputGroup
	
	if err := c.ShouldBindJSON(&InputGroup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	mainList := models.Group{
		Title:       InputGroup.Title,
		CurrentWeek: InputGroup.CurrentWeek,
		TotalWeeks:  InputGroup.TotalWeeks,
		IsFinished:  InputGroup.IsFinished,
	}
	if err := config.DB.Create(&mainList).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"created": true})
}

func UpdateGroupByID(c *gin.Context) {
	var mainGroup models.Group
	if err := config.DB.First(&mainGroup, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var updateGroup models.UpdateGroup
	if err := c.ShouldBindJSON(&updateGroup); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
	}
	if err := config.DB.Model(&mainGroup).Where("id = ?", c.Param("id")).Updates(updateGroup).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"update": true})

}

func DeleteGroupByID(c *gin.Context) {
	var deleteGroup models.Group
	if err := config.DB.Unscoped().Delete(&deleteGroup, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": true})
}

//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

func GetNotes(c *gin.Context) {
	var noteList []models.Note
	if err := config.DB.Preload("Student").Find(&noteList).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"notes": noteList})
}

func GetNotesStudentsByID(c *gin.Context) {
	var student models.Student
	id := c.Param("id")

	if err := config.DB.Preload("Notes").First(&student, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, student)
}

func GetNotesByID(c *gin.Context) {
	var note models.Note
	if err := config.DB.First(&note, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, note)
}

func CreateNotes(c *gin.Context) {
	var inputNotes models.InputNote

	if err := c.ShouldBindJSON(&inputNotes); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	mainNotesList := models.Note{
		StudentID: inputNotes.StudentID,
		Author:    inputNotes.Author,
		Text:      inputNotes.Text,
	}

	if err := config.DB.Create(&mainNotesList).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"created": true})
}

func UpdateNotes(c *gin.Context) {
	var mainNote models.Note
	if err := config.DB.First(&mainNote, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var updateNote models.UpdateNote
	if err:= c.ShouldBindJSON(&updateNote); err !=nil{
			c.IndentedJSON(http.StatusBadRequest ,  gin.H{"error": err})
		}

	if err := config.DB.Model(&mainNote).Where("id = ?", c.Param("id")).Updates(updateNote).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"update": true})
}

func DeleteNotes(c *gin.Context) {
	var mainNote models.Note
	if err := config.DB.Delete(&mainNote, c.Param("id")).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"delete": true})
}
