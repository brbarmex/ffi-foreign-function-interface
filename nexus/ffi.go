package main

/*
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"time"
	"unsafe"

	eventpb "fii/protos"

	"google.golang.org/protobuf/proto" // Pacote gerado pelo Protobuf
)

// Simula a criação de uma lista de eventos SQS e retorna como um []byte serializado
//
//export GetSQSEvents
func GetSQSEvents(output **C.char, outputLen *C.int) C.int {
	// Criar uma lista de eventos SQS
	events := []*eventpb.SQSMessage{
		{
			Id:            "1",
			MessageBody:   "Mensagem 1",
			ReceiptHandle: "handle1",
			QueueUrl:      "https://sqs.us-east-1.amazonaws.com/12345/queue1",
			Timestamp:     time.Now().Unix(),
		},
		{
			Id:            "2",
			MessageBody:   "Mensagem 2",
			ReceiptHandle: "handle2",
			QueueUrl:      "https://sqs.us-east-1.amazonaws.com/12345/queue2",
			Timestamp:     time.Now().Unix(),
		},
	}

	// Criar um objeto SQSMessageList com a lista de eventos
	eventsList := &eventpb.SQSMessageList{
		Messages: events,
	}

	// Serializar para []byte
	data, err := proto.Marshal(eventsList)
	if err != nil {
		fmt.Println("Erro ao serializar eventos:", err)
		return C.int(-1)
	}

	// Alocar memória para os dados e retornar para o C#
	cData := C.CBytes(data)
	*output = (*C.char)(cData)
	*outputLen = C.int(len(data))

	return C.int(0)
}

//export FreeMemory
func FreeMemory(ptr *C.char) {
	C.free(unsafe.Pointer(ptr))
}

func main() {}
