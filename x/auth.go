//
// Support for X authorisation 
// Currently only supports MIT-MAGIC-COOKIE
//
package x

import "fungo/gut"
import "bufio"
import "io"
import "os"


// Finds the local hostname if the hostname given is localhost or empty 
func GetAuthorityHostname(hostname string) (string, os.Error) {
  if len(hostname) == 0 || hostname == "localhost" {
    newhostname, err := os.Hostname()
    if err != nil {
      return "", err
    }
    return newhostname, nil
  }
  return hostname, nil
}


// Finds the authority file, opens it, and returns it.  
func GetAuthorityFile() (*os.File, os.Error) {
  filename := os.Getenv("XAUTHOTITY")
  if len(filename) == 0 { 
    home := gut.HomeDir()
    if len(home) == 0 {
      err := os.NewError("Xauthority not found: $XAUTHORITY, $HOME not set.")
      return nil, err
    }
    filename = gut.JoinDir(home, ".Xauthority")
  }
  file, err := gut.Fopen(filename, "r")
  if err != nil { return nil, err }
  return file, nil
}

type AuthorityInfo struct {
    Family      uint16
    Address     string
    Display     string
    Name        string
    Data        []byte
}

func ReadSize16BEBytes(buf * bufio.Reader) ([]byte, os.Error) {
  var size uint16
  err    := gut.UnpackBE(buf, size)
  if err != nil { return nil, err }
  result := make([]byte, size)
  _, err  = io.ReadFull(buf, result)
  if err != nil { return nil, err }
  return result, nil
}

func ReadSize16BEString(buf * bufio.Reader) (string, os.Error) {
  bytes, err := ReadSize16BEBytes(buf)
  if err != nil { return "", err }
  return (string)(bytes), nil   
}

// Reads in a single record of authority information from a bufio.Reader
func ReadAuthorityInfo(buf * bufio.Reader) (*AuthorityInfo, os.Error) {
  info  := &AuthorityInfo{}
  err   := gut.UnpackBE(buf, info.Family)
  if err != nil {  return nil, err }
  // reads sized string
  info.Address, err = ReadSize16BEString(buf)
  if err != nil {  return nil, err }
  // reads 0 terminated string
  info.Display, err = ReadSize16BEString(buf)
  if err != nil { return nil, err  }
  info.Name, err    = ReadSize16BEString(buf)
  if err != nil { return nil, err  }
  info.Data, err    = ReadSize16BEBytes(buf)
  if err != nil { return nil, err  }
  return info, nil
}   


// ReadAuthority reads the X authority file for the DISPLAY.
// If hostname == "" or hostname == "localhost",
// readAuthority uses the system's hostname (as returned by os.Hostname) instead.
// Returns a pointer to an AuthorityInfo struct
func ReadAuthority(hostname, display string) (* AuthorityInfo, os.Error) {
  // As per /usr/include/X11/Xauth.h.
  const familyLocal = 256
  
  hostname, err1   := GetAuthorityHostname(hostname)
  if err1 != nil { return nil, err1 }
  authfile, err2   := GetAuthorityFile()
  if err2 != nil { return nil, err2 }
  defer authfile.Close()

  buf := bufio.NewReader(authfile)

  for {
    info, err := ReadAuthorityInfo(buf)    
    if err != nil { return nil, err }
    if info.Family  == familyLocal && 
       info.Address == hostname && 
       info.Display == display {
      return info, nil
    }
  }
  return nil, nil 
}













