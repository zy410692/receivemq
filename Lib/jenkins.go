package Lib

import (
	"context"
	"log"

	"github.com/bndr/gojenkins"
)

var Conn *gojenkins.Jenkins

func init_no() {
	jenkins := gojenkins.CreateJenkins(nil, "http://x.x.x.x:xxxx/", "admin", "admin")
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	_, err := jenkins.Init(context.Background())

	if err != nil {
		log.Printf("连接Jenkins失败, %v\n", err)
		return
	}
	log.Println("Jenkins连接成功")

	Conn = jenkins
}

func GetConn() *gojenkins.Jenkins {
	return Conn
}
