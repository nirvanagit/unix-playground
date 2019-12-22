## Two Programs
sender program sends a SIGTERM to the receiver program, where receiver program gracefully 
handles it


```
# cd into signal-receiver and run
➜  signal-receiver git:(master) ✗ go run main.go &
➜  signal-receiver git:(master) ✗ go run main.go &
[1] 21479
➜  signal-receiver git:(master) ✗ awaiting signal
```

From signal sender run
```
# cd into signal-sender
➜  signal-sender git:(master) ✗ RECEIVER_PID="21479" go run main.go
2019/12/21 21:47:35 sending interrupt to process: 21479
➜  signal-sender git:(master) ✗
```
