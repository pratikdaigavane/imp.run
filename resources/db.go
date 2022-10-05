package re

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
	"time"
)

var Session gocqlx.Session
var Err error

func Connect() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "url_shortener"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = time.Second * 10

	Session, Err = gocqlx.WrapSession(cluster.CreateSession())
	if Err != nil {
		log.Println(Err)
		return
	}
}

func Close() {
	Session.Close()
}
