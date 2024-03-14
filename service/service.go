package service

import (
	"errors"
	"exampleAPIs/model"
	"exampleAPIs/repository"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ServicePort interface {
	PostServices(parametersInput model.ParametersInput) error
	PatchServices(parametersUpdate model.ParametersUpdate) error
	GetServices(parameter1 string) (model.InfoResponse, error)
	DeleteServices(parameter1 string) error
	GetAllServices() ([]model.InfoResponseAll, error)
}

type serviceAdapter struct {
	r repository.RepositoryPort
}

func NewServiceAdapter(r repository.RepositoryPort) ServicePort {
	return &serviceAdapter{r: r}
}

func (s *serviceAdapter) PostServices(parametersInput model.ParametersInput) error {
	if parametersInput == (model.ParametersInput{}) {
		return errors.New("parameters must not be empty(case1)")
	}
	if parametersInput.Gender == "" || parametersInput.Firstname == "" || parametersInput.Lastname == "" || parametersInput.BirthDate == "" || parametersInput.Nationality == "" || parametersInput.Ethnicity == "" {
		return errors.New("parameters must not be empty(case2)")
	}
	err := s.r.PostRepositories(parametersInput)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// func (s *serviceAdapter) PatchServices(parametersUpdate model.ParametersUpdate) error {
// 	sqlStatement := "UPDATE student SET"
// 	var placeholders []interface{}

// 	if parametersUpdate.Gender != nil {
// 		sqlStatement += " gender = $2,"
// 		placeholders = append(placeholders, *parametersUpdate.Gender)
// 	}
// 	if parametersUpdate.Firstname != nil {
// 		sqlStatement += " firstname = $3,"
// 		placeholders = append(placeholders, *parametersUpdate.Firstname)
// 	}
// 	if parametersUpdate.Lastname != nil {
// 		sqlStatement += " lastname = $4,"
// 		placeholders = append(placeholders, *parametersUpdate.Lastname)
// 	}
// 	if parametersUpdate.BirthDate != nil {
// 		sqlStatement += " birthdate = $5,"
// 		placeholders = append(placeholders, *parametersUpdate.BirthDate)
// 	}
// 	if parametersUpdate.Nationality != nil {
// 		sqlStatement += " nationality = $6,"
// 		placeholders = append(placeholders, *parametersUpdate.Nationality)
// 	}
// 	if parametersUpdate.Ethnicity != nil {
// 		sqlStatement += " ethnicity = $7,"
// 		placeholders = append(placeholders, *parametersUpdate.Ethnicity)
// 	}

// 	if len(placeholders) == 0 {
// 		return errors.New("no parameters to update")
// 	}
// 	sqlStatement = strings.TrimSuffix(sqlStatement, ",")
// 	sqlStatement += " WHERE student_id = $1;"
// 	fmt.Println(sqlStatement)

// 	placeholders = append([]interface{}{parametersUpdate.StudentID}, placeholders...)
// 	fmt.Println(placeholders)
// 	err := s.r.PatchRepositories(parametersUpdate, sqlStatement, placeholders...)
// 	if err != nil {
// 		log.Println(err)
// 		return err
// 	}
// 	return nil
// }

func (s *serviceAdapter) PatchServices(parametersUpdate model.ParametersUpdate) error {
	sqlStatement := "UPDATE student SET"
	var placeholders []interface{}

	counter := 1
	if parametersUpdate.Gender != nil {
		counter++
		sqlStatement += " gender = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.Gender)
	}
	if parametersUpdate.Firstname != nil {
		counter++
		sqlStatement += " firstname = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.Firstname)
	}
	if parametersUpdate.Lastname != nil {
		counter++
		sqlStatement += " lastname = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.Lastname)
	}
	if parametersUpdate.BirthDate != nil {
		counter++
		sqlStatement += " birthdate = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.BirthDate)
	}
	if parametersUpdate.Nationality != nil {
		counter++
		sqlStatement += " nationality = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.Nationality)
	}
	if parametersUpdate.Ethnicity != nil {
		counter++
		sqlStatement += " ethnicity = $" + strconv.Itoa(counter) + ","
		placeholders = append(placeholders, *parametersUpdate.Ethnicity)
	}

	if len(placeholders) == 0 {
		return errors.New("no parameters to update")
	}

	sqlStatement = strings.TrimSuffix(sqlStatement, ",")
	sqlStatement += " WHERE student_id = $1;"
	placeholders = append([]interface{}{parametersUpdate.StudentID}, placeholders...)
	fmt.Println(sqlStatement)
	fmt.Println(placeholders)
	err := s.r.PatchRepositories(parametersUpdate, sqlStatement, placeholders...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *serviceAdapter) GetServices(parameter1 string) (model.InfoResponse, error) {
	response, err := s.r.GetRepositories(parameter1)
	if err != nil {
		log.Println(err)
		return model.InfoResponse{}, err
	}
	return response, nil
}

func (s *serviceAdapter) DeleteServices(parameter1 string) error {
	err := s.r.DeleteRepositories(parameter1)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *serviceAdapter) GetAllServices() ([]model.InfoResponseAll, error) {
	response, err := s.r.GetAllReopsitories()
	if err != nil {
		log.Println(err)
		return response, err
	}
	return response, nil
}
