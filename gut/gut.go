//
// Gut is a package of Go utility functions that wrap around standard. 
// Go functionality to make it even easier to use.
//

package gut

import "os"
import "bytes"
import "strings"
import "encoding/binary"
import "io"

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
  
  
// ReadAll takes an os.File as it's argument and reads in all of it until 
// it signals to stop with an EOF in 
func ReadAll(file * os.File) ([]byte, os.Error) { 
  const BUFSIZE = 1024
  result    := make([]byte, 0)
  buf       := make([]byte, BUFSIZE)
  var sub []byte 
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
    if err2 != nil    { return nil   , err2 }    
  }
  // should not get here  
  return nil, os.NewError("Can't happen?!");
}

  
// Freadall opens the file named filename for reading and reads all of it  
// into a []byte which it will return. Also returns OS.error 
// if something went wrong  
 
func FreadAll(filename string) ([]byte, os.Error) {
  file, err := Fopen(filename, "r")
  if err != nil  { return nil, err }
  defer file.Close() // ensure file gets closed on return.
  // read all contents from the file.
  return ReadAll(file)
}


// Returns the file name of the home directory
func HomeDir() string {  
  return os.Getenv("HOME")
}  

// Path separator to use for JoinDir(...) 
var PathSeparator = "/"

// Joins parts of the file name together, by joining them with 
// PathSeparator is that is needed 
func JoinDir(prefix string, parts ... string) (string) {
  result := prefix
  for index, part := range parts {
    result += part 
    if !strings.HasSuffix(part, PathSeparator) {
      if index < (len(parts) - 1) { 
        // don't add for the last part either
        result += PathSeparator
      }
    }    
  }
  return result
}
 
// UnpackOrder unpacks a byte array to the arguments given in the varargs
func UnpackOrder(reader io.Reader, 
  order binary.ByteOrder, args... interface{}) (os.Error) {
  for _ , arg := range args {
    err := binary.Read(reader, order, arg)
    if err != nil { return err; }  
  } 
  return nil;
}
 
// Byte order that Unpack should use, Little endian by default
var UnpackByteOrder = binary.LittleEndian

// Unpack unpacks a byte array to the arguments given in the varargs
// uses binary.Read on all of the arguments to this  
func Unpack(reader io.Reader, args... interface{}) (os.Error) {
  return UnpackOrder(reader, UnpackByteOrder, args);
}

// Unpacks using Big Endian byte order
func UnpackBE(reader io.Reader, args... interface{}) (os.Error) {
  return UnpackOrder(reader, binary.BigEndian, args);
}

// Unpacks using Lyttle Endian byte order
func UnpackLE(reader io.Reader, args... interface{}) (os.Error) {
  return UnpackOrder(reader, binary.LittleEndian, args);
}

// PackOrder packs a byte array into the writer 
func PackOrder(writer io.Writer, 
  order binary.ByteOrder, args... interface{}) (os.Error) {
  for _ , arg := range args {
    err := binary.Write(writer, order, arg)
    if err != nil { return err; }  
  } 
  return nil;
}
 
// Byte order that Pack should use, Little endian by default
var PackByteOrder = binary.LittleEndian

// Pack packs a byte array to the arguments given in the varargs
// uses binary.Writeon all of the arguments to this  
func Pack(writer io.Writer, args... interface{}) (os.Error) {
  return PackOrder(writer, PackByteOrder, args);
}

// Packs using Big Endian byte order
func PackBE(writer io.Writer, args... interface{}) (os.Error) {
  return PackOrder(writer, binary.BigEndian, args);
}

// Unpacks using Lyttle Endian byte order
func PackLE(writer io.Writer, args... interface{}) (os.Error) {
  return PackOrder(writer, binary.LittleEndian, args);
}











