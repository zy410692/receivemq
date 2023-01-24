package main

import (
	"context"
	"fmt"
	"log"
	"receivemq/Lib"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/streadway/amqp"
)

func main() {
	// 建立连接
	conn, err := amqp.Dial("amqp://admin:123456@192.168.31.227:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel: %v", err)
	}
	defer ch.Close()

	// 声明消息队列
	q, err := ch.QueueDeclare(
		"newuser", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	// 消费消息
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
	}
	forever := make(chan bool)
	// 处理消息
	go func() {
		for msg := range msgs {
			processMessage(msg.Body)
		}
	}()

	// 阻塞
	<-forever
}

func processMessage(msg []byte) {

	jsonObject, err := simplejson.NewJson(msg)
	if err != nil {
		fmt.Println(err)
	}
	//workflow_id := jsonObject.Get("id").MustInt()
	//云主机创建工单
	workflow_stack := jsonObject.Get("form_data").GetIndex(0).Get("input_1612250282000_96996").MustString()
	//生产上线工单
	workflow_online := jsonObject.Get("from_data").GetIndex(0).Get("file_1611724316000_4915").MustString()
	//sonar扫描工单
	workflow_giturl := jsonObject.Get("form_data").GetIndex(0).Get("input_1615165427000_3196").MustString()
	workflow_branch := jsonObject.Get("form_data").GetIndex(0).Get("input_1615165498000_68047").MustString()
	//新建账号工单

	//判断工单下面的数量

	vpn_array := jsonObject.Get("form_data").GetIndex(0).Get("subform_1662442035000_24908").MustArray()
	fmt.Println(len(vpn_array))

	if len(vpn_array) > 0 {
		//执行python脚本生成excel,从minio下载文件
	}

	//获取工单ID,可通过数据库查看工单类型

	//jsonObject.Get("form_data").GetIndex(0).Get("")

	if workflow_stack != "" {
		fmt.Println("云主机创建流程 -- TODO")
		//fmt.Println(workflow_stack)

	}

	if workflow_online != "" {
		fmt.Println("生产上线工单 -- TODO")
		fmt.Println("生产工单上传的附件:\n")
		fmt.Println(workflow_online)
	}

	if workflow_giturl != "" {
		fmt.Println("sonar扫描工单 -- ")
		fmt.Println(workflow_giturl)
		fmt.Println(workflow_branch)
		conn := Lib.GetConn()
		params := map[string]string{"gitlab_url": strings.TrimSpace(workflow_giturl), "gitlab_branch": strings.TrimSpace(workflow_branch)}
		res, err := conn.BuildJob(context.Background(), "sonar-base", params)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}

}
