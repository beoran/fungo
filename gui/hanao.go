// Gui library. 
package gui

import "fungo/sdl"
import "fungo/draw"
//import "math"

// Time mouse button must be down to generate a click
const CLICK_TIME    = 0.5
// Time keyboard button must be down to generate a repeat
const REPEAT_TIME   = 0.5 
// Time between repeated keys
const REPEAT_SPEED  = 0.05 
// Enable workaround for utf scancode bug on Ubuntu
const SCANCODE_BUG_WORKAROUND = true

var IMAGE_DIR = "data"

//  mod stores the current state of the keyboard modifiers as explained in SDL_GetModState.
// 
// The unicode field is only used when UNICODE translation is enabled with SDL_EnableUNICODE. If unicode is non-zero then this is the UNICODE character corresponding to the keypress. If the high 9 bits of the character are 0, then this maps to the equivalent ASCII character:
// 
// char ch;
// if ( (keysym.unicode & 0xFF80) == 0 ) {
//   ch = keysym.unicode & 0x7F;
// }
// else {
//   printf("An International Character.\n");
// }

// Miniature smalltalk/ruby inspired object system. 
// Messages that can be sent to objects 
type Message int

// The object interface
type Object interface {
  Send(m Message, args... Object) (Object)
  DefineMethod(messsage Message, method Method)
  GetMethod(mes Message) (Method)
} 

// Methods 
type Method func(o Object, args ...Object) (Object)

// Mapping from messages to methdos
type MessageMap map[Message] Method

// Sends a message to an Object
func Send(o Object, m Message, args ...Object) (Object) {
  action := o.GetMethod(m)
  if action == nil { return nil }
  return (action)(o, args)
} 


// Concrete implementation of a basic object
type BasicObject struct { 
  MessageMap  
} 

// constructs a basic object 
func NewObject() (Object) {
  m := make(map[Message] Method)
  return &BasicObject{m}
}

// Sends a message to a basic object
func (o * BasicObject) Send(m Message, args ...Object) (Object) {
  action := o.GetMethod(m)
  if action == nil { return nil }
  return (action)(o, args)
} 
        
    
// Sends add message to a basic object
func (o * BasicObject) DefineMethod(mes Message, met Method) {
  o.MessageMap[mes] = met
} 
    
// Gets the mehod for the message ro nil if nou found  add message to a basic object
func (o * BasicObject) GetMethod(mes Message) (Method) {  
  action, ok := o.MessageMap[mes]
  if !ok { return nil }
  return action
} 

    
type Event struct {
  *sdl.Event
  Object
} 
  
    
    
func LoadImage(fname string) (* sdl.Surface) {
  // fname = Fimyfi.join(self.image_dir, *names)
  return sdl.LoadFastSurface(fname, true) 
}

type Style struct {
  
}

type Keyboard struct {
  
}

type Mouse struct {
  
}

type Joysticks struct {
  
}

const (
  Active  = Message(sdl.ACTIVEEVENT)
  KeyDown = Message(sdl.KEYUP)
  KeyUp   = Message(sdl.KEYDOWN)
  Quit    = Message(sdl.QUIT)
)


// Hanao is the main GUI widget manager.
type Hanao struct { 
  Object
  * draw.Surface // Screen
  Style          // Default style duplicated for all widgets
  Mouse          // Mouse information
  Keyboard       // Keyboard information
  Joysticks      // Array of joystick information
  main    *Widget    // Main widget, which contains all child widgets
  hovered []*Widget  // The widgets we are hovering over, if any.
  clicked *Widget    // The widget we are clicking on, if any.
  dragged *Widget    // The widget we are currently dragging, if any.
  focused *Widget    // The widget that has the current input focus.
  pressed *Widget    // The widget that the mouse is pressed on if any.
  done   bool       // Is the system requesting quit?
  active bool       // Is the system active?  
  focuscursor * sdl.Surface // cursors  
}
  
var ACTIVE_HANAO * Hanao = nil;  
  
func NewHanao(screen * sdl.Surface) (*Hanao) {  
  res := &Hanao{}
  res.Init(screen)
  ACTIVE_HANAO = res
  return res
}

func CurrentHanao() (*Hanao) {
  return ACTIVE_HANAO
}

func (h * Hanao) Done() (bool) {
  return h.done
}
    
    
// Polls the SDL state to update the GUI    
func (h * Hanao) Update() {
  for { 
    ev          := sdl.PollEvent()
    if ev == nil { break } 
    kind        := ev.Type    
    Send(h, Message(kind), Event{ev, NewObject()})
    // call the handler
  }
}
    
// Handles one event     
    
func (h * Hanao) Init(screen * sdl.Surface) {
  h.Surface        = draw.FromSDL(screen)
  // h.Style       = NewStyle()
  h.Object         = NewObject()
  h.done           = false
  h.active         = true 
  h.hovered        = make([]*Widget, 16)
  // The widgets we are hovering over, if any.
  h.pressed        = nil 
  // The widget we are pressing down on, if any.
  h.clicked        = nil // The widget we are clicking on, if any.
  h.dragged        = nil // The widget we are currently dragging, if any.
  h.focused        = nil // The widget that has the current input focus.
  h.focuscursor    = LoadImage("data/joystick_0.png")
  // Cursor for focusing
  // XXX: Set up handlers
  h.Object.DefineMethod(Quit, OnQuit)
}  

// Sends events to every widget interested in it
func (h * Hanao) sendToWidgets(message Message, args ...) { 
  for widget := range h.main.SelfAndEachChild() { 
    if ! widget.Active() { continue }
    if widget.Ignore(message) { continue }
    res := widget.Send(message, args)
    if res == nil { break }
  } 
}

// Event handlers

// On activation
func OnActive(h * Hanao, event * sdl.Event) {

}

// To unwrap the arguments.
func unwrapArgs(o Object, args ...Object) (*Hanao, *sdl.Event) {
      hanao   := o.(*Hanao)
      event   := args[0].(Event).Event
      return hanao, event
}

// Called when the system wants to shutdown
func  OnQuit(o Object, args ...Object) (Object) {
      h, event:= unwrapArgs(o, args) 
      println(event.Type)
      h.done   = true
      return h
}

// Called when key is pressed
func OnKeydown(h * Hanao, event * sdl.Event) {
  kevent := event.Keyboard()
  text   := "" // CleanupUnicode(kevent)
  // keyboard.press(event.sym, event.mod, text)          
  h.sendToWidgets(KeyDown, kevent.Keysym, kevent.Keysym.Mod, text)
}
    
// Called when key is released
func OnKeyup(h * Hanao, event * sdl.Event) {
  kevent := event.Keyboard()
  text   := "" // CleanupUnicode(kevent)
  // text    = cleanup_unicode(event)
  // state   = keyboard.state(event.sym)
  h.sendToWidgets(KeyUp, kevent.Keysym, kevent.Keysym.Mod, text)  
  // keyboard.release(event.sym)
}


// Called when the mouse moves
func OnMouseMotion(h * Hanao, event * sdl.Event) {
}
/*
    old_x , old_y     = @mouse.x , @mouse.y
      @mouse.move(event.x, event.y)
      old_hovered       = @hovered || []
      @hovered          = @main.all_under?(event.x, event.y)    
      new_hovered       = @hovered - old_hovered
      unhovered         = old_hovered - @hovered
      moving_over       = @hovered - new_hovered      
      unhovered.each    { |widget| widget.on_mouse_out(event.x, event.y, nil) }
      new_hovered.each  { |widget| widget.on_mouse_in(event.x, event.y , nil) }
      moving_over.each  { |widget| widget.on_mouse_move(event.x, event.y)     }
      if @hovered.member?(@pressed)
        //we're above the pressed widget. Drag it
        @pressed.on_mouse_drag(old_x, old_y, event.x, event.y)
      end
      @mouse.under      = @hovered
      // send_to_interested(:on_mouse_move_to, event.x, event.y, old_x, old_y)
*/

    
// Called when the mouse wheel is scrolled
func OnMouseScroll(h * Hanao, event * sdl.Event) {
}
/*
      widget = @hovered.first
      if (widget)
        widget.on_scroll(scroll)
      end
*/







