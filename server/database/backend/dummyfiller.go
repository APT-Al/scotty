package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789 "
const tokenCharset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ipNumSet  = "1234567890."
const numSet  = "1234567890"
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {b[i] = charset[seededRand.Intn(len(charset))]}
	return string(b)
}
func String(length int) string {return StringWithCharset(length, charset)}
func dbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "pass"
	dbName := "dbname"
	// "myuser:password@tcp(127.0.0.1:3306)/mydb" formatıyla alt satırı editle
	whoWherePort:=dbUser+":"+dbPass+"@tcp(ip:port)/"+dbName
	db, err := sql.Open(dbDriver, whoWherePort)
	if err != nil {
		panic(err.Error())
	}
	return db
}
func randate() time.Time {
	min := time.Date(2010, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2021, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
func main() {
	db := dbConnect()
	//insForm, _ := db.Prepare("INSERT INTO Victim(malware_id, victim_ip,victim_local_ip,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status ) VALUES(?,?,?,?,?,?,?,?,?)")
	//_, _ = insForm.Exec(1, StringWithCharset(16, ipNumSet), StringWithCharset(16, ipNumSet), StringWithCharset(16, charset), StringWithCharset(10, charset), seededRand.Intn(4), StringWithCharset(10, charset), seededRand.Intn(1), seededRand.Intn(1))


	//for i:=0 ;i<180 ;i++  {
	//	insForm, err := db.Prepare("INSERT INTO AttackVector( attack_vector_type,embedded_file ) VALUES(?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, _ = insForm.Exec( "PBR", "/root/PBR/")
	//}


	//for i:=0 ;i<200 ;i++  {
	//	insForm, err := db.Prepare("INSERT INTO MalwareType( malware_type_id, malware_type,malware_version,malware_release_date,target_system,atack_vector_id ) VALUES(?,?,?,?,?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	var mal_type string
	//	var targetSystem string
	//	var ranran = rand.Intn(2)
	//	if ranran%2==0 {mal_type="offline"}else {mal_type = "online"}
	//	ranran = rand.Intn(2)
	//	if ranran%3==0 {targetSystem="linux"}else {targetSystem = "windows"}
	//	_, _ = insForm.Exec( i, mal_type,string(StringWithCharset(1, numSet)+"."+StringWithCharset(1, numSet)), randate(),targetSystem,rand.Intn(200))
	//}


	//for i:=0 ;i<200 ;i++  {
	//	insForm, err := db.Prepare("INSERT INTO Malware( malware_type_id,situation) VALUES(?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, _ = insForm.Exec( rand.Intn(100),StringWithCharset(15,charset))
	//}


	for i:=0 ;i<150 ;i++  {
		insForm, err := db.Prepare("INSERT INTO MalwareStatus ( id,create_date,infected_date,first_touch_with_cc,clean_date) VALUES(?,?,?,?,?)")
		if err != nil {
			panic(err)
		}
		_, _ = insForm.Exec( i, randate(),randate(),randate(),randate())
	}


	//for i:=0 ;i<100 ;i++  {
	//	insForm, err := db.Prepare("INSERT INTO Money (malware_id,  amount,status ) VALUES(?,?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, _ = insForm.Exec( rand.Intn(200) , rand.Intn(1000),rand.Intn(2))
	//}


	//for i:=0 ;i<200 ;i++  {
	//	insForm, err := db.Prepare("INSERT INTO Victim(id, victim_ip,victim_local_ip,computer_name,username,computer_ram,computer_cpu,computer_status,botnet_status ) VALUES(?,?,?,?,?,?,?,?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, _ = insForm.Exec(rand.Intn(200),StringWithCharset(16, ipNumSet),StringWithCharset(16, ipNumSet),StringWithCharset(25, charset),StringWithCharset(25, charset),rand.Intn(16256),
	//		"Intel "+StringWithCharset(10,charset) ,rand.Intn(1),rand.Intn(1))
	//}


	//for i:=0 ;i<100 ;i++  {
	//	proto := "TCP";if i%12==0 {proto ="UDP"}else {proto = "TCP"}
	//	insForm, err := db.Prepare("INSERT INTO Botnet(victim_id,port,protocol,token ) VALUES(?,?,?,?)")
	//	if err != nil {
	//		panic(err)
	//	}
	//	_, _ = insForm.Exec(rand.Intn(200),rand.Intn(65535),proto,StringWithCharset(40,tokenCharset))
	//}
	defer db.Close()
}
