package semaphores

import (
	"fmt"
    "golang.org/x/sync/semaphore"
	"context"
)

type employee struct {
	ID string
}

var (
    sem = semaphore.NewWeighted(int64(10))
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
	ctx := context.Background()
	for _, employee := range employeeList {
		sem.Acquire(ctx, 1) // equivalent to sem <- 1 (using channel approach)
		go func(){
			remoteDeleteEmployeeRPC(employee.ID)
			sem.Release(1) // equivalent to <- sem (using channel approach)
		}()
	}
}
