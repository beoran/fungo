// Created by cgo - DO NOT EDIT
//line sdl.go:1
/*
Go Language wrappers around Open CV
*/
package opencv

// #include <SDL.h>


/*
import "unsafe"
import "fmt"
import "os"
import "runtime"
*/

// The available application states
const (
	APPMOUSEFOCUS	= SDL_APPMOUSEFOCUS	//0x01		/* The app has mouse coverage */
	APPINPUTFOCUS	= SDL_APPINPUTFOCUS	//0x02		/* The app has input focus */
	APPACTIVE	= SDL_APPACTIVE		//0x04		/* The application is active */
)

/* Function prototypes */
/*
 * This function returns the current state of the application, which is a
 * bitwise combination of SDL_APPMOUSEFOCUS, SDL_APPINPUTFOCUS, and
 * SDL_APPACTIVE.  If SDL_APPACTIVE is set, then the user is able to
 * see your application, otherwise it has been iconified or disabled.
 */

func GetAppState() uint8	{ return uint8(_C_SDL_GetAppState()) }
