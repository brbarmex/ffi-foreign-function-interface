package queue

/*
#include <stdlib.h>
#include <string.h>  // Para C.memcpy
#include <stddef.h>  // Para C.size_t
*/
import "C"

import (
	"fmt"
	"os"
	"time"
	"unsafe"

	eventpb "cloudsdk/protos"

	"google.golang.org/protobuf/proto"
)

//export GetSQSEvents
func GetSQSEvents(output **C.char, outputLen *C.int) C.int {
	events := getEvent()
	data, err := proto.Marshal(&events)
	if err != nil {
		fmt.Println("Erro ao serializar eventos:", err)
		return C.int(-1)
	}

	// Aloca memória para os dados e passa para o C#
	cData := C.CBytes(data)
	*output = (*C.char)(cData)
	*outputLen = C.int(len(data))

	return C.int(0)
}

//export FreeMemory
func FreeMemory(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func getEvent() eventpb.MessageQueueList {
	env := os.Getenv("cloudsdk_provider")
	if env == "azure" {
		return getEventFromEventBus()
	}

	//ever will to print aws as default, its a simple test
	return getEventFromSQS()
}

func getEventFromSQS() eventpb.MessageQueueList {
	return eventpb.MessageQueueList{
		Messages: []*eventpb.MessageQueue{
			{
				Id:            "1",
				MessageBody:   "Mensagem 1",
				ReceiptHandle: "handle1",
				QueueUrl:      "https://sqs.us-east-1.amazonaws.com/12345/queue1",
				Timestamp:     time.Now().Unix(),
				Provider:      "AWS",
			},
			{
				Id:            "2",
				MessageBody:   "Mensagem 2",
				ReceiptHandle: "handle2",
				QueueUrl:      "https://sqs.us-east-1.amazonaws.com/12345/queue2",
				Timestamp:     time.Now().Unix(),
				Provider:      "AWS",
			},
		},
	}
}

func getEventFromEventBus() eventpb.MessageQueueList {
	return eventpb.MessageQueueList{
		Messages: []*eventpb.MessageQueue{
			{
				Id:            "1",
				MessageBody:   "Mensagem 1",
				ReceiptHandle: "handle1",
				QueueUrl:      "https://event-bus-azure.com/12345/queue1",
				Timestamp:     time.Now().Unix(),
				Provider:      "Azure",
			},
			{
				Id:            "2",
				MessageBody:   "Mensagem 2",
				ReceiptHandle: "handle2",
				QueueUrl:      "https://event-bus-azure.com/12345/queue2",
				Timestamp:     time.Now().Unix(),
				Provider:      "Azure",
			},
		},
	}
}
