package repository

import (
	"database/sql"
	"errors"
	"exampleAPIs/model"
	"fmt"
	"log"
)

type RepositoryPort interface {
	PostRepositories(parametersInput model.ParametersInput) error
	PatchRepositories(parametersUpdate model.ParametersUpdate, sqlStatement string, placeHolders ...interface{}) error
	GetRepositories(paramerter1 string) (model.InfoResponse, error)
	DeleteRepositories(parameter1 string) error
	GetAllReopsitories() ([]model.InfoResponseAll, error)
}

type repositoryAdapter struct {
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) RepositoryPort {
	return &repositoryAdapter{db: db}
}

func (r *repositoryAdapter) PostRepositories(parametersInput model.ParametersInput) error {
	var exists1 string
	err := r.db.QueryRow("SELECT firstname FROM student WHERE firstname = $1", parametersInput.Firstname).Scan(&exists1)
	if err == nil {
		log.Println(err)
		return errors.New("student already exists")
	}
	if err != sql.ErrNoRows {
		log.Println(err)
		return errors.New("unexpected error")
	}
	_, err = r.db.Exec("INSERT INTO student (gender, firstname, lastname, birthdate, nationality, ethnicity) VALUES ($1, $2 ,$3, $4, $5, $6)", parametersInput.Gender, parametersInput.Firstname, parametersInput.Lastname, parametersInput.BirthDate, parametersInput.Nationality, parametersInput.Ethnicity)
	if err != nil {
		log.Println(err)
		return errors.New("failed to insert parameter(student)")
	}
	return nil
}

func (r *repositoryAdapter) PatchRepositories(parametersUpdate model.ParametersUpdate, sqlStatement string, placeholders ...interface{}) error {
	// var parameter1Exists string
	// err := r.db.QueryRow("SELECT parameter1 FROM exampleapis WHERE parameter1 = $1", parametersUpdate.Parameter1).Scan(&parameter1Exists)
	// if err != nil {
	// 	log.Println(err)
	// 	return errors.New("parameter1 does not match")
	// }
	// fmt.Println(parameter1Exists)
	result, err := r.db.Exec(sqlStatement, placeholders...)
	if err != nil {
		log.Println(err)
		return err
	}
	rowChangedPatch, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(rowChangedPatch)
	if rowChangedPatch == 0 {
		return errors.New("studentID(patch) does not match")
	}
	return nil
}

func (r *repositoryAdapter) GetRepositories(parameter1 string) (model.InfoResponse, error) {
	// var infoResponse []model.InfoResponse // incase many row
	var infoResponse model.InfoResponse
	err := r.db.QueryRow("SELECT * FROM student WHERE student_id = $1", parameter1).Scan(&infoResponse.StudentID, &infoResponse.Gender, &infoResponse.Firstname, &infoResponse.Lastname, &infoResponse.BirthDate, &infoResponse.Nationality, &infoResponse.Ethnicity)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No rows were found.")
			return model.InfoResponse{}, errors.New("studentID(get) does not match")
		}
		log.Println(err)
		return model.InfoResponse{}, err
	}
	return infoResponse, nil
}

func (r *repositoryAdapter) DeleteRepositories(parameter1 string) error {
	// var parameter1Exists string
	// err := r.db.QueryRow("SELECT parameter1 FROM exampleapis WHERE parameter1 = $1", parameter1).Scan(&parameter1Exists)
	// if err != nil {
	// 	log.Println(err)
	// 	return errors.New("parameter1 does not match")
	// }
	result, err := r.db.Exec("DELETE FROM student WHERE student_id = $1", parameter1)
	if err != nil {
		log.Println(err)
		return err
	}
	rowChangedGet, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println(rowChangedGet)
	if rowChangedGet == 0 {
		return errors.New("studentID(delete) does not match")
	}
	return nil
}

func (r *repositoryAdapter) GetAllReopsitories() ([]model.InfoResponseAll, error) {
	var infoResponses []model.InfoResponseAll
	rows, err := r.db.Query("SELECT * FROM student")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var infoResponse model.InfoResponseAll
		err := rows.Scan(&infoResponse.StudentID, &infoResponse.Gender, &infoResponse.Firstname, &infoResponse.Lastname, &infoResponse.BirthDate, &infoResponse.Nationality, &infoResponse.Ethnicity)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		infoResponses = append(infoResponses, infoResponse)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return infoResponses, nil
}
