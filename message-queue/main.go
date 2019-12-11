package main

import (
	"github.com/siadat/ipc"
	"fmt"
	"log"
	"time"
)

func main() {

	mykey, err := ipc.Ftok("/dev/null", 42)
	if err != nil {
		panic(fmt.Sprintf("Failed to generate key: %s\n", err))
	} else {
		fmt.Printf("Generate key %d\n", mykey)
	}

	qid, err := ipc.Msgget(mykey, ipc.IPC_CREAT|0600)
	if err != nil {
		panic(fmt.Sprintf("Failed to create ipc key %d: %s\n", mykey, err))
	} else {
		fmt.Printf("Create ipc queue id %d\n", qid)
	}

	defer func() {
		err = ipc.Msgctl(qid, ipc.IPC_RMID)
		if err != nil {
			log.Fatal(err)
		}
	}()

	input := []byte{0x18, 0x2d, 0x44, 0x00, 0xfb, 0x21, 0x09, 0x40}

	msg := ipc.Msgbuf{Mtype: 12, Mtext: input}
	err = ipc.Msgsnd(qid, &msg, 0)
	if err != nil {
		panic(fmt.Sprintf("Failed to send message to ipc id %d: %s\n", qid, err))
	} else {
		fmt.Printf("Message %v send to ipc id %d\n", input, qid)
	}

	fmt.Println("sleep for 5 seconds, to check if iPC is created run ipcs -q")
	time.Sleep(5 * time.Second)

	qbuf := &ipc.Msgbuf{Mtype: 12}

	err = ipc.Msgrcv(qid, qbuf, 0)

	if err != nil {
		panic(fmt.Sprintf("Failed to receive message to ipc id %d: %s\n", qid, err))
	} else {
		fmt.Printf("Message %v receive to ipc id %d\n", qbuf.Mtext, qid)
	}
}
