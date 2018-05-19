package conns

//defer session.Close()
import (
	"github.com/gocql/gocql"
	"ms/sun/shared/helper"
)

var CassndraSession *gocql.Session

func init() {
	connectToCassandra()
}

func connectToCassandra() {
	//cluster := gocql.NewCluster("192.168.1.250")
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "sunc"
	cluster.Consistency = gocql.One
	var err error
	CassndraSession, err = cluster.CreateSession()
	helper.NoErrDebug(err)
}
