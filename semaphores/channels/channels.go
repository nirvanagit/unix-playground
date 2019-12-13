package channels

import (
	"fmt"
)

type employee struct {
	ID string
}

var (
	sem = make(chan int, 10)
	employeeList = []employee{{ID: "first"}, {ID: "second"}}
)


// remoteDeleteEmployeeRPC will delete an employee over the network.
func remoteDeleteEmployeeRPC(id string) {
	// your code here
	fmt.Printf("delete for %s", id)
	return
}

// sem is a channel that will allow up to 10 concurrent operations.
func main() {
	for _, employee := range employeeList {
		sem <- 1
		go func(){
			remoteDeleteEmployeeRPC(employee.ID)
			<-sem
		}()
	}
}