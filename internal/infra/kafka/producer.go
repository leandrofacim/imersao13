package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{
		ConfigMap: configMap,
	}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)

	if err != nil {
		return err
	}

	// deliveryChan := make(chan ckafka.Event)

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{
			Topic: &topic,
			// Partition: ckafka.PartitionAny,
			Partition: 0,
		},
		Value: msg.([]byte),
		Key:   key,
	}

	err = producer.Produce(message, nil)

	if err != nil {
		return err
	}

	// event := <-deliveryChan
	// msgDelivery := event.(*ckafka.Message)

	// if msgDelivery.TopicPartition.Error != nil {
	// 	return msgDelivery.TopicPartition.Error
	// }

	return nil
	// return err
}
