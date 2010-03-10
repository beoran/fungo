package sdl

// #include <SDL.h>
import "C"
// import "unsafe"

// General keyboard/mouse state definitions 
const (
  RELEASED 	= C.SDL_RELEASED
  PRESSED	= C.SDL_PRESSED
)

// Event enumerations 
type SDL_Events uint8

const (
  // Unused (do not remove)
  NOEVENT		= iota
  // Application loses/gains visibility 
  ACTIVEEVENT   	
  // Keys pressed 
  KEYDOWN		
  // Keys released
  KEYUP			
  // Mouse moved 
  MOUSEMOTION		
  // Mouse button pressed 
  MOUSEBUTTONDOWN	
  // Mouse button released
  MOUSEBUTTONUP		
  // Joystick axis motion
  JOYAXISMOTION		
  // Joystick trackball motion 
  JOYBALLMOTION		
  // Joystick hat position change 
  JOYHATMOTION		
  // Joystick button pressed 
  JOYBUTTONDOWN		
  // Joystick button released        
  JOYBUTTONUP		
  // User-requested quit 
  QUIT			
  // System specific event 
  SYSWMEVENT		
  // User resized video mode   
  VIDEORESIZE		
  // Screen needs to be redrawn 
  VIDEOEXPOSE		
  // Events SDL_USEREVENT through SDL_MAXEVENTS-1 are for your use 
  USEREVENT		= 24
  // This last event is only for bounding internal arrays
  // It is the number of bits in the event mask datatype -- Uint3    
  NUMEVENTS		= 32
)

type SDL_EventMasks uint32

func EVENTMASK(x SDL_Events) (SDL_EventMasks) { 
  return SDL_EventMasks((1<<(x)))
}  

var (
  ACTIVEEVENTMASK	= EVENTMASK(ACTIVEEVENT)
  KEYDOWNMASK		= EVENTMASK(KEYDOWN)
  KEYUPMASK		= EVENTMASK(KEYUP)
  KEYEVENTMASK		= EVENTMASK(KEYDOWN) | EVENTMASK(KEYUP)
  MOUSEMOTIONMASK	= EVENTMASK(MOUSEMOTION)
  MOUSEBUTTONDOWNMASK	= EVENTMASK(MOUSEBUTTONDOWN)
  MOUSEBUTTONUPMASK	= EVENTMASK(MOUSEBUTTONUP)
  MOUSEEVENTMASK	= EVENTMASK(MOUSEMOTION)|
			  EVENTMASK(MOUSEBUTTONDOWN)|
	                  EVENTMASK(MOUSEBUTTONUP)
  JOYAXISMOTIONMASK	= EVENTMASK(JOYAXISMOTION)
  JOYBALLMOTIONMASK	= EVENTMASK(JOYBALLMOTION)
  JOYHATMOTIONMASK	= EVENTMASK(JOYHATMOTION)
  JOYBUTTONDOWNMASK	= EVENTMASK(JOYBUTTONDOWN)
  JOYBUTTONUPMASK	= EVENTMASK(JOYBUTTONUP)
  JOYEVENTMASK		= EVENTMASK(JOYAXISMOTION)|
	                  EVENTMASK(JOYBALLMOTION)|
	                  EVENTMASK(JOYHATMOTION)|
	                  EVENTMASK(JOYBUTTONDOWN)|
	                  EVENTMASK(JOYBUTTONUP)
  VIDEORESIZEMASK	= EVENTMASK(VIDEORESIZE)
  VIDEOEXPOSEMASK	= EVENTMASK(VIDEOEXPOSE)
  QUITMASK		= EVENTMASK(QUIT)
  SYSWMEVENTMASK	= EVENTMASK(SYSWMEVENT)
)
const ALLEVENTS		= SDL_EventMasks(0xFFFFFFFF)

// Predefined event masks 
// Not imported due to limitations in cgo

// Application visibility event structure 
/*
typedef struct SDL_ActiveEvent {
	Uint8 type;	// SDL_ACTIVEEVENT 
	Uint8 gain;	// Whether given states were gained or lost (1/0) 
	Uint8 state;	// A mask of the focus states 
} SDL_ActiveEvent;

// Keyboard event structure 
typedef struct SDL_KeyboardEvent {
	Uint8 type;	// SDL_KEYDOWN or SDL_KEYUP 
	Uint8 which;	// The keyboard device index 
	Uint8 state;	// SDL_PRESSED or SDL_RELEASED 
	SDL_keysym keysym;
} SDL_KeyboardEvent;

// Mouse motion event structure 
typedef struct SDL_MouseMotionEvent {
	Uint8 type;	// SDL_MOUSEMOTION 
	Uint8 which;	// The mouse device index 
	Uint8 state;	// The current button state 
	Uint16 x, y;	// The X/Y coordinates of the mouse 
	Sint16 xrel;	// The relative motion in the X direction 
	Sint16 yrel;	// The relative motion in the Y direction 
} SDL_MouseMotionEvent;

// Mouse button event structure 
typedef struct SDL_MouseButtonEvent {
	Uint8 type;	// SDL_MOUSEBUTTONDOWN or SDL_MOUSEBUTTONUP 
	Uint8 which;	// The mouse device index 
	Uint8 button;	// The mouse button index 
	Uint8 state;	// SDL_PRESSED or SDL_RELEASED 
	Uint16 x, y;	// The X/Y coordinates of the mouse at press time 
} SDL_MouseButtonEvent;

// Joystick axis motion event structure 
typedef struct SDL_JoyAxisEvent {
	Uint8 type;	// SDL_JOYAXISMOTION 
	Uint8 which;	// The joystick device index 
	Uint8 axis;	// The joystick axis index 
	Sint16 value;	// The axis value (range: -32768 to 32767) 
} SDL_JoyAxisEvent;

// Joystick trackball motion event structure 
typedef struct SDL_JoyBallEvent {
	Uint8 type;	// SDL_JOYBALLMOTION 
	Uint8 which;	// The joystick device index 
	Uint8 ball;	// The joystick trackball index 
	Sint16 xrel;	// The relative motion in the X direction 
	Sint16 yrel;	// The relative motion in the Y direction 
} SDL_JoyBallEvent;

// Joystick hat position change event structure 
typedef struct SDL_JoyHatEvent {
	Uint8 type;	// SDL_JOYHATMOTION 
	Uint8 which;	// The joystick device index 
	Uint8 hat;	// The joystick hat index 
	Uint8 value;	// The hat position value:
			    SDL_HAT_LEFTUP   SDL_HAT_UP       SDL_HAT_RIGHTUP
			    SDL_HAT_LEFT     SDL_HAT_CENTERED SDL_HAT_RIGHT
			    SDL_HAT_LEFTDOWN SDL_HAT_DOWN     SDL_HAT_RIGHTDOWN
			   Note that zero means the POV is centered.
			
} SDL_JoyHatEvent;

// Joystick button event structure 
typedef struct SDL_JoyButtonEvent {
	Uint8 type;	// SDL_JOYBUTTONDOWN or SDL_JOYBUTTONUP 
	Uint8 which;	// The joystick device index 
	Uint8 button;	// The joystick button index 
	Uint8 state;	// SDL_PRESSED or SDL_RELEASED 
} SDL_JoyButtonEvent;

// The "window resized" event
   When you get this event, you are responsible for setting a new video
   mode with the new width and height.
 
typedef struct SDL_ResizeEvent {
	Uint8 type;	// SDL_VIDEORESIZE 
	int w;		// New width 
	int h;		// New height 
} SDL_ResizeEvent;

// The "screen redraw" event 
typedef struct SDL_ExposeEvent {
	Uint8 type;	// SDL_VIDEOEXPOSE 
} SDL_ExposeEvent;

// The "quit requested" event 
typedef struct SDL_QuitEvent {
	Uint8 type;	// SDL_QUIT 
} SDL_QuitEvent;

// A user-defined event type 
typedef struct SDL_UserEvent {
	Uint8 type;	// SDL_USEREVENT through SDL_NUMEVENTS-1 
	int code;	// User defined event code 
	void *data1;	// User defined data pointer 
	void *data2;	// User defined data pointer 
} SDL_UserEvent;

// If you want to use this event, you should include SDL_syswm.h 
struct SDL_SysWMmsg;
typedef struct SDL_SysWMmsg SDL_SysWMmsg;
typedef struct SDL_SysWMEvent {
	Uint8 type;
	SDL_SysWMmsg *msg;
} SDL_SysWMEvent;

// General event structure 
typedef union SDL_Event {
	Uint8 type;
	SDL_ActiveEvent active;
	SDL_KeyboardEvent key;
	SDL_MouseMotionEvent motion;
	SDL_MouseButtonEvent button;
	SDL_JoyAxisEvent jaxis;
	SDL_JoyBallEvent jball;
	SDL_JoyHatEvent jhat;
	SDL_JoyButtonEvent jbutton;
	SDL_ResizeEvent resize;
	SDL_ExposeEvent expose;
	SDL_QuitEvent quit;
	SDL_UserEvent user;
	SDL_SysWMEvent syswm;
} SDL_Event;
*/


// Pumps the event loop, gathering events from the input devices.
// This function updates the event queue and internal input device state.
// This should only be run in the thread that sets the video mode.
func PumpEvents() {
  C.SDL_PumpEvents()
}

type SDL_eventaction int

const ( 
  ADDEVENT 	= SDL_eventaction(C.SDL_ADDEVENT)
  PEEKEVENT 	= SDL_eventaction(C.SDL_PEEKEVENT)
  GETEVENT  	= SDL_eventaction(C.SDL_GETEVENT)
)
// Checks the event queue for messages and optionally returns them.
// If 'action' is ADDEVENT, up to 'numevents' events will be added to
// the back of the event queue.
// If 'action' is PEEKEVENT, up to 'numevents' events at the front
// of the event queue, matching 'mask', will be returned and will not
// be removed from the queue.
// If 'action' is GETEVENT, up to 'numevents' events at the front 
// of the event queue, matching 'mask', will be returned and will be
// removed from the queue.
// This function returns the number of events actually stored, or -1
// if there was an error.  This function is thread-safe.
func peepEvents( events * C.SDL_Event, numevents int, 
		 action SDL_eventaction, mask uint32) (int) {
  return int(C.SDL_PeepEvents(events, C.int(numevents), 
    C.SDL_eventaction(action), C.Uint32(mask)))
}


// Polls for currently pending events, and returns 1 if there are any pending
// events, or 0 if there are none available.  If 'event' is not NULL, the next
// event is removed from the queue and stored in that area.
func pollEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_PollEvent(event))
}
  

// Waits indefinitely for the next available event, returning 1, or 0 if there
// was an error while waiting for events.  If 'event' is not NULL, the next
// event is removed from the queue and stored in that area.
func waitEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_WaitEvent(event))
}
 
// Add an event to the event queue.
// This function returns 0 on success, or -1 if the event queue was full
// or there was some other error.
func pushEvent(event * C.SDL_Event) (int) { 
  return int(C.SDL_PushEvent(event))
}

// Event filtereing is not supported yet due to cg not supporting callbacks 
// well.
//
// This function sets up a filter to process all events before they
// change internal state and are posted to the internal event queue.
// The filter is protypted as:
// typedef int (SDLCALL *SDL_EventFilter)(const SDL_Event *event);
//
//   If the filter returns 1, then the event will be added to the internal queue.
//   If it returns 0, then the event will be dropped from the queue, but the 
//   internal state will still be updated.  This allows selective filtering of
//   dynamically arriving events.
// 
//   WARNING:  Be very careful of what you do in the event filter function, as 
//             it may run in a different thread!
// 
//   There is one caveat when dealing with the SDL_QUITEVENT event type.  The
//   event filter is only called when the window manager desires to close the
//   application window.  If the event filter returns 1, then the window will
//   be closed, otherwise the window will remain open if possible.
//   If the quit event is generated by an interrupt signal, it will bypass the
//   internal queue and be delivered to the application at the next event poll.
// 
// extern DECLSPEC void SDLCALL SDL_SetEventFilter(SDL_EventFilter filter);
//
//   Return the current event filter - can be used to "chain" filters.
//   If there is no event filter set, this function returns NULL.
// 
// extern DECLSPEC SDL_EventFilter SDLCALL SDL_GetEventFilter(void);


const (
  QUERY		= -1
  IGNORE	= 0
  DISABLE	= 0
  ENABLE	= 1
)
//
//   This function allows you to set the state of processing certain events.
//   If 'state' is set to SDL_IGNORE, that event will be automatically dropped
//   from the event queue and will not event be filtered.
//   If 'state' is set to SDL_ENABLE, that event will be processed normally.
//   If 'state' is set to SDL_QUERY, SDL_EventState() will return the 
//   current processing state of the specified event.
func EventState(kind uint8, state int) (uint8) {
  return uint8(C.SDL_EventState(C.Uint8(kind), C.int(state)))
}

// Returns true if a quit was requested.
func QuitRequested() (bool) { 
  C.SDL_PumpEvents()
  return i2b(int(C.SDL_PeepEvents(nil, 0, C.SDL_eventaction(PEEKEVENT), 		
	  C.Uint32(QUITMASK))))
}	

