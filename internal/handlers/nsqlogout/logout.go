package nsqlogout

import "github.com/nsqio/go-nsq"

type Handler interface {
	HandleMessage(message *nsq.Message) error
}
