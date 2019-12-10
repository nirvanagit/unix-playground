package main

import (
	"log"
	"syscall"

	"github.com/siadat/ipc"
)

func main() {
	key, err := ipc.Ftok("/dev/null", 42)
	if err != nil {
		panic(err)
	}

	qid, err := ipc.Msgget(key, ipc.IPC_CREAT|ipc.IPC_EXCL|0600)
	if err == syscall.EEXIST {
		log.Fatalf("queue(key=0x%x) exists", key)
	}
	if err != nil {
		log.Fatal(err)
	}

	msg := &ipc.Msgbuf{Mtype: 12, Mtext: []byte("message")}
	err = ipc.Msgsnd(qid, msg, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully sent a message")
}