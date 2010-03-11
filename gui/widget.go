package gui
 
import "container/list"


type Widget struct {
  BasicObject
  parent   * Widget
  children * list.List
}


func (w * Widget) Add(child * Widget) (*Widget) {
  child.parent = w
  w.children.PushBack(child)
  return w
}

func (w * Widget) Each() (<-chan *Widget) {
  out := make(chan *Widget)
  itr := func() {
    for child := range w.children.Iter() {
      out <- child.(*Widget)
    }
  }
  go itr()
  return out 
}

func (w * Widget) SelfAndEachChild() (chan *Widget) {
  result := make(chan *Widget)
  return result 
}

func (w * Widget) Send(msg Message, args...) (Object) {
  return nil
}

func (w * Widget) Active() (bool) {
  return true
}

func (w * Widget) Ignore(message Message) (bool) {
  return false
}












































