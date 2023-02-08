package kafka

import (
	"encoding/json"
	"log"
	"os"
	"time"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/luiz7r/imersao12/application/route"
	"github.com/luiz7r/imersao12/infra/kafka"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}