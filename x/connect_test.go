package x

import "testing"
import "fmt"

func TestConnect(t * testing.T) {
  fmt.Println("OK!")
  c := ConnectLocal() ; defer c.Close() 
  fmt.Println(c)
  c.Authenticate()
  fmt.Println(c)


}





