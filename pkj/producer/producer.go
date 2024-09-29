package producer

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"

	"github.com/Megidy/TaskManagmentSystem/pkj/models"
	"github.com/Megidy/TaskManagmentSystem/pkj/types"
	"github.com/Megidy/TaskManagmentSystem/pkj/utils"
	"github.com/gin-gonic/gin"
)

const (
	Topic string = "Status_changes"
)

func ChangeStatus(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		utils.HandleError(c, nil, "failed to retrieve users data", http.StatusUnauthorized)
		return
	}
	id := c.Param("taskId")
	taskId, err := strconv.Atoi(id)
	if err != nil {
		utils.HandleError(c, err, "failed to retrieve params", http.StatusBadRequest)
		return
	}
	ok, err = models.IsCreated(taskId, user.(*types.User).Id)
	if err != nil {
		utils.HandleError(c, err, "failed to get data from db", http.StatusInternalServerError)
		return
	}
	if !ok {
		utils.HandleError(c, nil, "no task found", http.StatusBadRequest)
		return
	}

	var ChangeStatus types.ChangeStatus
	ChangeStatus.TaskId = taskId
	ChangeStatus.UserId = user.(*types.User).Id

	err = c.ShouldBindJSON(&ChangeStatus)
	if err != nil {
		utils.HandleError(c, err, "failed to read body", http.StatusBadRequest)
		return
	}
	if ChangeStatus.Status == "done" {
		ok, err := models.HaveDependencies(user.(*types.User).Id, taskId)
		if err != nil {
			utils.HandleError(c, err, "failed to get dependencies from db", http.StatusInternalServerError)
			return
		}
		if ok {
			c.JSON(http.StatusConflict, gin.H{
				"error": "you cant complete task when you didnt complete other dependent tasks",
			})
			return
		}
	}

	statusInBytes, err := json.Marshal(ChangeStatus)
	if err != nil {
		utils.HandleError(c, err, "failed to marshal body ", http.StatusInternalServerError)
		return
	}
	err = PushStatusToQueue(Topic, statusInBytes)
	if err != nil {
		utils.HandleError(c, err, "failed to send message to broker ", http.StatusInternalServerError)
		return
	}
	var NewLog = types.Log{
		UserId: user.(*types.User).Id,
		TaskId: taskId,
		Action: "Updated tasks status",
	}
	err = models.CreateLog(NewLog)
	if err != nil {
		log.Println()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "status of user" + user.(*types.User).Username + "updated ",
	})

}
func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)

}

func PushStatusToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(&msg)
	if err != nil {

		return err
	}
	log.Printf("status is stored in topic(%s)/partition(%d)/offset(%d)\n",
		topic, partition, offset)
	return nil
}
