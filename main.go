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
	sqlStorage:=connection(cfg)
	db,err:=sqlStorage.Init()
	if err!=nil{
		log.Fatal(err)
	}
	store:=Newstore(db)
	server:=newserverAPI(":3001",*store)
	server.run()
}