
package main

import (
    "errors"
    "fmt"
)

type DiyError struct {
    err error
}

func (de *DiyError) Error() error {
    return de.err
}


func RaiseError() *DiyError {
    de := &DiyError{err: errors.New("Raise error for test.")}
    return de
}

func TestError() (string,error){
    de := RaiseError()
    return "",de.err
}    

func TestSuccess() (string,error){
    msg,de := "Test success.",nil
    
    return msg,de.err
}    


func main () {
    msg1,err1 := TestError()
    if err1 != nil {
        fmt.Println(err1)
    }
    fmt.Println(msg1)

    msg2,err2 := TestSuccess()
    if err2 != nil {
        fmt.Println(err2)
    }
    fmt.Println(msg2)

}
