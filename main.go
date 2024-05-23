package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)
func main(){
	cfg := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	sqlStorage:=connection(cfg)  // db.go
	db,err:=sqlStorage.Init()    // creating tables db.go
	if err!=nil{
		log.Fatal(err) // fatal terminates code if err and prints it.
	}
	store:=Newstore(db)          //db.go
	server:=newserverAPI(":3001",*store)   // api.go
	server.run()   // api.go
}