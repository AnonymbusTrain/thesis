package models

import (
	"strconv"
	"time"

	"github.com/SanyaNooB/thesis.git/config"
	"github.com/go-faker/faker/v4"
)

type Person struct {
	Id        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
	Timestamp time.Time
}

type Message struct {
	Message string `json:"message"`
}

var applicationConfig = config.GetApplicationConfig()

func GetNumberOfPersons(numberOfPersons string) ([]*Person, error) {
	return getPersonsByQuery("SELECT id, firstName, lastName, birthDate, timestamp FROM " + applicationConfig.Postgres.Table + " LIMIT " + numberOfPersons)
}

func GetAllPersons() ([]*Person, error) {
	return getPersonsByQuery("SELECT id, firstName, lastName, birthDate, timestamp FROM " + applicationConfig.Postgres.Table)
}

func getPersonsByQuery(sqlQuery string) ([]*Person, error) {
	personRows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer personRows.Close()

	var persons []*Person
	for personRows.Next() {
		person := new(Person)
		if err := personRows.Scan(&person.Id, &person.FirstName, &person.LastName, &person.BirthDate, &person.Timestamp); err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	if err := personRows.Err(); err != nil {
		return nil, err
	}
	return persons, err
}

func GetPersonById(id string) (*Person, error) {
	// InitDB()
	var person Person
	row := db.QueryRow("SELECT id, firstName, lastName, birthDate, timestamp FROM "+applicationConfig.Postgres.Table+" WHERE id=$1",
		id)
	err := row.Scan(&person.Id, &person.FirstName, &person.LastName, &person.BirthDate, &person.Timestamp)

	if err != nil {
		return nil, err
	}
	return &person, nil
}

func InsertPersonToDb(person *Person) (*Message, error) {
	_, err := db.Exec("INSERT INTO "+applicationConfig.Postgres.Table+" (firstName, lastName, birthDate, timestamp) VALUES ($1, $2, $3, $4)", person.FirstName, person.LastName, person.BirthDate, time.Now().Format("2006-01-02T15:04:05.111"))
	if err != nil {
		return nil, err
	}
	return &Message{Message: "You have successfully inserted a person"}, nil
}

func GenerateNumberOfRandomPersonsToDB(numberOfPersonsToGenerate int64) (*Message, error) {
	transaction, err := db.Begin()
	if err != nil {
		return nil, err
	}

	var i int64
	for i = 0; i < numberOfPersonsToGenerate; i++ {
		_, err = transaction.Exec("INSERT INTO "+applicationConfig.Postgres.Table+" (firstName, lastName, birthDate, timestamp) VALUES ($1, $2, $3, $4)", faker.FirstName(), faker.LastName(), faker.Date(), time.Now().Format("2006-01-02T15:04:05.111"))
		if err != nil {
			return nil, err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return nil, err
	}
	return &Message{Message: "You have successfully generated " + strconv.FormatInt(numberOfPersonsToGenerate, 10) + " persons in the DB"}, nil
}

func UpdatePerson(person *Person) (*Message, error) {

	if _, err := db.Exec("UPDATE "+applicationConfig.Postgres.Table+" SET firstName = $1, lastName = $2, birthDate = $3 WHERE id = $4", person.FirstName, person.LastName, person.BirthDate, person.Id); err != nil {
		return nil, err
	}
	return &Message{"You habe successfully updated the person with the id: " + strconv.FormatInt(person.Id, 10)}, nil
}

func DeletePerson(id string) (*Message, error) {
	if _, err := db.Exec("DELETE FROM "+applicationConfig.Postgres.Table+" WHERE id = $1", id); err != nil {
		return nil, err
	}
	return &Message{"You habe successfully deleted the person with the id: " + id}, nil
}

func DropDBTable() (*Message, error) {

	if _, err := db.Exec("DROP TABLE " + applicationConfig.Postgres.Table); err != nil {
		return nil, err
	}
	return &Message{"You have successfully deleted the " + applicationConfig.Postgres.Table + " table"}, nil
}
