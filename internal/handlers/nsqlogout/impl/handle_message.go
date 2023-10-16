package impl

import (
	"context"
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"vehicle-log-grpc-api/internal/usecase/logs"
)

func (h handler) HandleMessage(message *nsq.Message) error {
	var req logs.MessageLogOut
	err := json.Unmarshal(message.Body, &req)
	if err != nil {
		return err
	}
	return h.logUc.LogOutProcess(context.Background(), req)
}
