## Handling Unix Signals

In golang, signal notification works by sending os.Signal values on a channel.

```
➜  signals git:(master) ✗ go run main.go
awaiting signal
^C
interrupt
exiting
➜  signals git:(master) ✗
```

### References
- [gobyexample](https://gobyexample.com/signals)
