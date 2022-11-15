package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "", "6379"),
		Password: "",
		DB:       2,
		PoolSize: 100,
	})
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	read := Read{
		User: "",
		Pass: "",
		Addr: "",
		Name: "taikang",
	}
	dbObj, err := dbConnect(read.User, read.Pass, read.Addr, read.Name)
	if err != nil {
		fmt.Println(err)
	}
	msgs, err := RedisClient.XRangeN(ctx, "ConsumerMVStream", "-", "+", 1000000).Result()
	if err != nil {
		fmt.Println(err)
	}

	for _, msg := range msgs {
		values := msg.Values
		templateId := fmt.Sprintf("%v", values["TemplateId"])
		orderId := fmt.Sprintf("%v", values["OrderId"])
		wxOpenId := fmt.Sprintf("%v", values["WxOpenId"])
		name := fmt.Sprintf("%v", values["Name"])
		blessing := fmt.Sprintf("%v", values["Blessing"])

		ossUrlListJson := fmt.Sprintf("%v", values["OSSUrlListJson"])
		anniversariesJson := fmt.Sprintf("%v", values["AnniversariesJson"])

		err := dbObj.Create(&TempOrderTemplate{
			TemplateId:        templateId,
			OrderId:           orderId,
			WXOpenId:          wxOpenId,
			Name:              name,
			Blessing:          blessing,
			OssUrlListJson:    ossUrlListJson,
			AnniversariesJson: anniversariesJson,
		}).Error
		if err != nil {
			fmt.Println(err)
		}
	}

	// b := make([]byte, 16)
	// _, err = rand.Read(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
	// 	b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	// fmt.Println("start success" + uuid)

	// consumer, err := stream.NewConsumer(&stream.ConsumerOptions{
	// 	RedisClient: RedisClient,
	// 	Stream:      "CustomerStream6",
	// 	GroupName:   "CustomerStreamGroup",
	// 	Consumer:    uuid,
	// 	Start:       "0",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// consumer.CreateGroupMkStream()
	// for {
	// 	ctx := context.Background()
	// 	entities, err := consumer.Poll(ctx)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//TODO
	// 	if entities != nil {
	// 		for i := 0; i < len(entities[0].Messages); i++ {
	// 			messageId := entities[0].Messages[i].ID
	// 			values := entities[0].Messages[i].Values
	// 			fmt.Println(values)
	// 			data := values["222"]
	// 			fmt.Println(fmt.Sprintf("%+v", data))
	// 			//ACK
	// 			consumer.Ack(ctx, messageId)
	// 		}
	// 	}
	// }
}

type TempOrderTemplate struct {
	Id                int64  `gorm:"primary_key;id" `
	TemplateId        string `gorm:"column:templateId"`        //订单id
	OrderId           string `gorm:"column:orderId"`           //订单id
	WXOpenId          string `gorm:"column:wxOpenId"`          // 模板ID
	Name              string `grom:"column:name"`              //名称
	Blessing          string `gorm:"column:blessing"`          // （原）图片地址
	OssUrlListJson    string `gorm:"column:ossUrlListJson"`    // 压缩包图片下载
	AnniversariesJson string `gorm:"column:anniversariesJson"` // 模板属性 json

}

func (*TempOrderTemplate) TableName() string {
	return "temp_20221031_order_template"
}

type Read struct {
	Addr string `toml:"addr"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Name string `toml:"name"`
}

func dbConnect(user, pass, addr, dbName string) (*gorm.DB, error) {
	config := &gorm.Config{}

	newDb, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pass,
		addr,
		dbName)), config)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
		return nil, err
	}
	sqldb, _ := newDb.DB()
	sqldb.SetMaxOpenConns(1000)
	sqldb.SetMaxIdleConns(100)
	sqldb.SetConnMaxIdleTime(1 * time.Hour)

	return newDb, nil
}
