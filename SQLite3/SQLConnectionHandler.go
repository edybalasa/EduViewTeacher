package SQLite3

import (
	"EduViewTeacher/networking"
	"database/sql"
	"log"
	"net"
)

func connectToSQLite3() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func checkForSQLite3Integrity() {
	db := connectToSQLite3()
	defer db.Close()
	db.Exec("CREATE TABLE IF NOT EXISTS connections(" +
		"ip VARCHAR(50) NOT NULL," +
		"hostname VARCHAR(50)) NOT NULL")
}

func searchHostnameInSQL(db *sql.DB, hostname string) bool {
	exists, err := db.Exec("SELECT hostname FROM connections WHERE HOSTNAME = %s", hostname)
	if err != nil {
		log.Fatal(err)
	}
	return exists == nil
}

func handleSQLite3Connection() *sql.DB {
	checkForSQLite3Integrity()
	db := connectToSQLite3()
	return db
}

func handleNonExistingHostname(db *sql.DB, ip *net.UDPAddr, hostname string) {
	db.Exec("insert into connections(ip, hostname) values(?, ?)", ip.String(), hostname)
	db.Close()
	sg := networking.SignalSender{}
	sg.SendPairConfirmationSignal(ip)
}

func HandleHostname(ip *net.UDPAddr, hostname string) {
	db := handleSQLite3Connection()
	defer db.Close()
	if searchHostnameInSQL(db, hostname) {
		db.Close()
		return
	} else {
		handleNonExistingHostname(db, ip, hostname)
	}
}
