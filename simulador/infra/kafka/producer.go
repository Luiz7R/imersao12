package kafka

import ckafka "github.com/confluentinc/conFluent-kafka-go/kafka"

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap {
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group_id":			 os.Getenv("KafkaConsumerGroupID"),
	}
	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}
	return p
}

func NewKafkaConsumer(msgChan chan *ckafka.Message) *KafkaConsumer {
	return &KafkaConsumer {
		MsgChan: msgChan
	}
}

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:			[]byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}

	return nil
}