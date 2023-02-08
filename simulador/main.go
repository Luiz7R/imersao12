package main

import (
	"fmt"
	kafka2 "github.com/luiz7r/imersao12/application/kafka"
	"github.com/luiz7r/imersao12/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)	

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error on loading .env file")
	}
}

func main() {
	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	for {
		_= 1
	}
}

