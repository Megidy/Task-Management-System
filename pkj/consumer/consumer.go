package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
)

const (
	topic string = "Status_changes"
)

func main() {
	msgCnt := 0
	worker, err := ConnectConsumer([]string{"localhost:9092"})
	if err != nil {
		log.Fatal(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatal(err)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				log.Println(err)
			case msg := <-consumer.Messages():
				msgCnt++
				log.Printf("received message: Count : %d: | Topic: (%s) | Status(%s)\n", msgCnt, string(msg.Topic), string(msg.Value))
				Status := string(msg.Value)
				var changeStatus types.ChangeStatus

				err := json.Unmarshal(msg.Value, &changeStatus)
				if err != nil {
					log.Println(err, "failed to change status in db")
				}

				err = models.ChangeStatus(changeStatus)
				if err != nil {
					log.Println(err, "failed to change status in db")
				}
				var NewLog = types.Log{
					UserId: user.(*types.User).Id,
					TaskId: task.Id,
					Action: "Created new task ",
				}
				err = models.CreateLog(NewLog)
				if err != nil {
					log.Println()
					return
				}

				log.Printf("StatusChanged %s\n", Status)
			case <-sigchan:
				log.Println("interrups detected")
				doneCh <- struct{}{}
			}

		}
	}()
	<-doneCh
	log.Println("Processed ", msgCnt, " Status")

	err = worker.Close()
	if err != nil {
		log.Fatal(err)
	}

}
func ConnectConsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	return sarama.NewConsumer(brokers, config)

}
