package patientcontroller

import (
	"net/http"
	"text/template"

	"github.com/raymondchr/go-mysql/entities"
	"github.com/raymondchr/go-mysql/models"
)

var patients = models.NewModelPatient()

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/patient/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/patient/add.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		r.ParseForm()

		var patient entities.Patient
		patient.FullName = r.Form.Get("name")
		patient.Id = r.Form.Get("id")
		patient.Sex = r.Form.Get("sex")
		patient.PoC = r.Form.Get("birthplace")
		patient.DoB = r.Form.Get("birthday")
		patient.Address = r.Form.Get("address")
		patient.Phone = r.Form.Get("phone")

		patients.Create(patient)
		data := map[string]interface{}{
			"message": "Data has been inserted",
		}

		temp, _ := template.ParseFiles("views/patient/add.html")
		temp.Execute(w, data)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}
