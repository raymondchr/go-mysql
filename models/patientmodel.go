package models

import (
	"database/sql"
	"fmt"

	"github.com/raymondchr/go-mysql/config"
	"github.com/raymondchr/go-mysql/entities"
)

type ModelPatient struct {
	conn *sql.DB
}

func NewModelPatient() *ModelPatient {
	connection, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &ModelPatient{
		conn: connection,
	}
}

func (p *ModelPatient) Create(patient entities.Patient) bool {
	result, err := p.conn.Exec("insert into patient (full_name, user_id, sex, place_of_birth, date_of_birth, address, phone) values(?,?,?,?,?,?,?)",
		patient.FullName, patient.Id, patient.Sex, patient.PoC, patient.DoB, patient.Address, patient.Phone)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
