//
// Gut is a package of Go utility functions that wrap around standard. 
// Go functionality to make it even easier to use.
//

package gut

import "os"
import "bytes"

// The default umask that Fopen uses when calling os.Open
var FopenDefaultUmask = 0666


// Parses the mode string and returns the fitting 
// os.Constants
func ParseMode(mode string) int {
  // read only by default
  var option int = os.O_RDONLY
  
  // If the string isn't valid, return read only
  if len(mode) == 0 { return option }
  
  if mode[0] == 'w' {   
    option = os.O_WRONLY | os.O_CREAT 
    // | os.O_TRUNC
  }
  
  if mode[0] == 'a' {
    option = os.O_APPEND | os.O_CREAT
  }
  
  // in all other cases, we assume the user wanted O_RDONLY
  if len(mode) == 1 { return option }
  
  // If we get here, there is more in the mode string
  if mode[1] == '+' {
    option = os.O_RDWR
    if mode[0] == 'a' {
      option = os.O_RDWR | os.O_CREAT | os.O_APPEND
    }
    
    if mode[0] == 'w' {
      option = os.O_RDWR | os.O_CREAT | os.O_TRUNC
    }
  }
  
  return option
}  


// Fopen opens a file conveniently, 
// using only the file name and a C-like, but enhanced fopen mode string.
// This mode string should start with r, w, or a 
// for read, write , or append mode. Otionally this is followed by + 
// to make access read/write
// w will truncate and create the file, a will append and create the file. 
func Fopen(filename, mode string) (*os.File, os.Error) {
  opt := ParseMode(mode)
  return os.Open(filename, opt, FopenDefaultUmask)
}
  
// Freadall opens the file named filename for reading and reads all of it  
// into a []byte which it will return. Also returns OS.error 
// if something went wrong  
 
func FreadAll(filename string) ([]byte, os.Error) {
  const BUFSIZE = 1024
  result    := make([]byte, 0)
  buf       := make([]byte, BUFSIZE)
  var sub []byte 
  file, err := Fopen(filename, "r")
  if err != nil  { return nil, err }
  defer file.Close(); // ensure file gets closed on return.
  for { // ever     
    count, err2 := file.Read(buf)    
    sub          = buf[0:count]
    // get right slice to append to result
    // append bytes to result
    if count > 0 { 
      result = bytes.Add(result, sub)
    }
    // if EOF, return the result. 
    if err2 == os.EOF { return result, nil }
    // on any other error return nil 
    if err2 != nil    { return nil   , err }    
  }
  // should not get here
  // file.Read(buf)  
  return nil, nil;
}


// Returns the file name of the home directory
func HomeDir() string {  
  return os.Getenv("HOME")
}  












