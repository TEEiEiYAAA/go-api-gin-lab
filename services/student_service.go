package services

import (
	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) error {
	// เช็คก่อนว่ามีนักศึกษาคนนี้อยู่ไหม
	_, err := s.GetStudentByID(id)
	if err != nil {
		return err
	}
	// ส่งไปให้ Repository อัปเดตข้อมูล
	return s.Repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id string) error {
	// เช็คก่อนว่ามีนักศึกษาคนนี้อยู่ไหม
	_, err := s.GetStudentByID(id)
	if err != nil {
		return err
	}
	// ส่งไปให้ Repository ลบข้อมูล
	return s.Repo.Delete(id)
}
