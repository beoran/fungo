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

func (w * Widget) Each(block func(Any)){  
  for child := range w.children.Iter() {
      block(child.(*Widget))
  }
}

func (w * Widget) SelfAndEachChild(block func(Any)) {
  /*
  result := make(chan *Widget)
  return result 
  */
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












































