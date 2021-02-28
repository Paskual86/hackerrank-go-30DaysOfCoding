#  Struct Instead of Classes
*This part was extracted from this site:* https://golangbot.com/structs-instead-of-classes/

*Only I took the best parts that I think that I need.*

- The *Go Language* do not support constructor. If the zero value of a type is not usable, it is the job of the programmer to unexport the type to prevent access from other packages and also provide the function named **NewT(parameters)**. **It is a convetion in Go to name a function that creates a value of type T. If the package defines only one type, then it's a convention in Go to name this function just New(parameters) instead of NewT(parameters)**


Example:

```Go
package employee

import (  
    "fmt"
)

type Employee struct {  
    FirstName   string
    LastName    string
    TotalLeaves int
    LeavesTaken int
}

func (e Employee) LeavesRemaining() {  
    fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
```

_How to call or execute:_

```go
package main

import "oop/employee"

func main() {  
    e := employee.Employee {
        FirstName: "Sam",
        LastName: "Adolf",
        TotalLeaves: 30,
        LeavesTaken: 20,
    }
    e.LeavesRemaining()
}
```

Now if we have that the struct is not accessible (*declare as **private***) we need to change the name of the struct *Employee*:

```go
type employee struct {  
    firstName   string
    lastName    string
    totalLeaves int
    leavesTaken int
}

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {  
    e := employee {firstName, lastName, totalLeave, leavesTaken}
    return e
}
```

*Main.Go*

```go
package main  

import "oop/employee"

func main() {  
    e := employee.New("Sam", "Adolf", 30, 20)
    e.LeavesRemaining()
}
```