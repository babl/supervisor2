package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/larskluge/babl-server/kafka"
	pb "github.com/larskluge/babl/protobuf/messages"
)

func (s *server) BroadcastCancelRequest(module string, rid uint64) error {
	cr := pb.CancelRequest{RequestId: rid}
	req := pb.Meta{
		Cancel: &cr,
	}
	msg, err := proto.Marshal(&req)
	if err != nil {
		return err
	}

	topic := ModuleNameToTopic(module, true)
	kafka.SendMessage(s.kafkaProducer, "", topic, &msg) // TODO return err
	return nil
}
