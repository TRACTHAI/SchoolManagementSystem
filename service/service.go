package service

import (
	"errors"
	"exampleAPIs/model"
	"exampleAPIs/repository"
	"fmt"
	"log"
	"strings"
)

type ServicePort interface {
	PostServices(parametersInput model.ParametersInput) error
	PatchServices(parametersUpdate model.ParametersUpdate) error
	GetServices(parameter1 string) (model.InfoResponse, error)
	DeleteServices(parameter1 string) error
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
	if parametersInput.Parameter1 == "" || parametersInput.Parameter2 == "" || parametersInput.Parameter3 == "" {
		return errors.New("parameters must not be empty(case2)")
	}
	err := s.r.PostRepositories(parametersInput)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *serviceAdapter) PatchServices(parametersUpdate model.ParametersUpdate) error {
	sqlStatement := "UPDATE exampleapis SET"
	var placeholders []interface{}

	if parametersUpdate.Parameter2 != nil {
		sqlStatement += " parameter2 = $2,"
		placeholders = append(placeholders, *parametersUpdate.Parameter2)
	}
	if parametersUpdate.Parameter3 != nil {
		sqlStatement += " parameter3 = $3,"
		placeholders = append(placeholders, *parametersUpdate.Parameter3)
	}

	if len(placeholders) == 0 {
		return errors.New("no parameters to update")
	}
	sqlStatement = strings.TrimSuffix(sqlStatement, ",")
	sqlStatement += " WHERE parameter1 = $1;"
	fmt.Println(sqlStatement)

	placeholders = append([]interface{}{parametersUpdate.Parameter1}, placeholders...)
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
