package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/student-api/models"
	"example.com/student-api/services"
)

type StudentHandler struct {
	Service *services.StudentService
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.Service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := h.Service.GetStudentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	//Challenge 3
	if student.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must not be empty"})
		return
	}
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must not be empty"})
		return
	}
	if student.GPA < 0.00 || student.GPA > 4.00 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gpa must be between 0.00 and 4.00"})
		return
	}

	if err := h.Service.CreateStudent(student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create student. ID might already exist."})
		return
	}

	c.JSON(http.StatusCreated, student)
}

// Challenge 1 and 3
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}
	//Challenge 3
	if student.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must not be empty"})
		return
	}
	if student.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name must not be empty"})
		return
	}
	if student.GPA < 0.00 || student.GPA > 4.00 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "gpa must be between 0.00 and 4.00"})
		return
	}

	if err := h.Service.UpdateStudent(id, student); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	student.Id = id
	c.JSON(http.StatusOK, student)
}

// Challenge 2
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteStudent(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
