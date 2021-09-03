package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/devafatek/WasteLibrary"
	_ "github.com/lib/pq"
)

func initStart() {

	WasteLibrary.LogStr("Successfully connected!")
	time.Sleep(time.Second * 10)
}
func main() {

	initStart()
	bulkDbSet()
	configDbSet()
	sumDbSet()
	staticDbSet()

}

func bulkDbSet() {
	var bulkDbHost string = "waste-bulkdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	bulkDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		bulkDbHost, port, user, password, dbname)

	bulkDb, err := sql.Open("postgres", bulkDbInfo)
	WasteLibrary.LogErr(err)
	defer bulkDb.Close()

	err = bulkDb.Ping()
	WasteLibrary.LogErr(err)

	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS listenerdata ( 
			data_id serial PRIMARY KEY,
			app_type varchar(50) NOT NULL DEFAULT '',
			serial_number varchar(50) NOT NULL DEFAULT '',
			data_type varchar(50) NOT NULL DEFAULT '',
			data TEXT NOT NULL DEFAULT '',
			customer_id INT NOT NULL DEFAULT -1,
			  data_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`)
	_, err = bulkDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	bulkDb.Close()
}

func configDbSet() {
	var configDbHost string = "waste-configdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	configDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configDbHost, port, user, password, dbname)

	configDb, err := sql.Open("postgres", configDbInfo)
	WasteLibrary.LogErr(err)
	defer configDb.Close()

	err = configDb.Ping()
	WasteLibrary.LogErr(err)

	configDb.Close()
}

func sumDbSet() {
	var sumDbHost string = "waste-sumdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	sumDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sumDbHost, port, user, password, dbname)

	sumDb, err := sql.Open("postgres", sumDbInfo)
	WasteLibrary.LogErr(err)
	defer sumDb.Close()

	err = sumDb.Ping()
	WasteLibrary.LogErr(err)

	var createSQL string = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS redisdata ( 
			data_id serial PRIMARY KEY,
			hashkey TEXT NOT NULL DEFAULT '',
			subkey TEXT NOT NULL DEFAULT '', 
			keyvalue TEXT NOT NULL DEFAULT '',
			  create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			);`)
	_, err = sumDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	sumDb.Close()
}

func staticDbSet() {
	var sumDbHost string = "waste-staticdb-cluster-ip"
	var port int = 5432
	var user string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var dbname string = os.Getenv("POSTGRES_DB")
	sumDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		sumDbHost, port, user, password, dbname)

	sumDb, err := sql.Open("postgres", sumDbInfo)
	WasteLibrary.LogErr(err)
	defer sumDb.Close()

	err = sumDb.Ping()
	WasteLibrary.LogErr(err)
	var createSQL string

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS tags ( 
		tag_id serial PRIMARY KEY,
		app_type varchar(50) NOT NULL DEFAULT '',
		serial_number varchar(50) NOT NULL DEFAULT '',
		customer_id INT NOT NULL DEFAULT -1,
		epc varchar(50) NOT NULL DEFAULT '',
		uid varchar(50) NOT NULL DEFAULT '',
		container_no varchar(50) NOT NULL DEFAULT '',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		statu varchar(50) NOT NULL DEFAULT '0',
		image_statu varchar(50) NOT NULL DEFAULT '0',
		active varchar(50) NOT NULL DEFAULT '1',
		read_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		check_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = sumDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS devices ( 
		device_id serial PRIMARY KEY,
		device_type varchar(50) NOT NULL DEFAULT '',
		serial_number varchar(50) NOT NULL DEFAULT '',
		device_name varchar(50) NOT NULL DEFAULT '',
		customer_id INT NOT NULL DEFAULT -1,
		reader_app_status varchar(50) NOT NULL DEFAULT '0',
		reader_conn_status varchar(50) NOT NULL DEFAULT '0',
		reader_status varchar(50) NOT NULL DEFAULT '0',
		cam_app_status varchar(50) NOT NULL DEFAULT '0',
		cam_conn_status varchar(50) NOT NULL DEFAULT '0',
		cam_status varchar(50) NOT NULL DEFAULT '0',
		gps_app_status varchar(50) NOT NULL DEFAULT '0',
		gps_conn_status varchar(50) NOT NULL DEFAULT '0',
		gps_status varchar(50) NOT NULL DEFAULT '0',
		therm_app_status varchar(50) NOT NULL DEFAULT '0',
		transfer_app_status varchar(50) NOT NULL DEFAULT '0',
		alive_status varchar(50) NOT NULL DEFAULT '0',
		contact_status varchar(50) NOT NULL DEFAULT '0',
		therm varchar(50) NOT NULL DEFAULT '0',
		latitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		longitude NUMERIC(14, 11)  NOT NULL DEFAULT 0, 
		active varchar(50) NOT NULL DEFAULT '1',
		read_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		therm_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		gps_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		status_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = sumDb.Exec(createSQL)
	WasteLibrary.LogErr(err)

	createSQL = fmt.Sprintf(`CREATE TABLE IF NOT EXISTS customers ( 
		customer_id serial PRIMARY KEY,
		customer_name varchar(50) NOT NULL DEFAULT '',
		domain varchar(50) NOT NULL DEFAULT '',
		rfid_app varchar(50) NOT NULL DEFAULT '0',
		ult_app varchar(50) NOT NULL DEFAULT '0',
		recy_app varchar(50) NOT NULL DEFAULT '0',
		active varchar(50) NOT NULL DEFAULT '1',
		create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`)
	_, err = sumDb.Exec(createSQL)
	WasteLibrary.LogErr(err)
	sumDb.Close()
}