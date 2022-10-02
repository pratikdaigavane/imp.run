package re

import (
	"github.com/gocql/gocql"
	"log"
	"time"
)

var Session *gocql.Session
var Err error

func Connect() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10

	Session, Err = cluster.CreateSession()
	if Err != nil {
		log.Println(Err)
		return
	}
}

func Close() {
	Session.Close()
}
