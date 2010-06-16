package x

import "net"
import "fmt"
import "os"
import "io"
import "bytes"
// import "bufio"

/*
Connection Setup
For TCP connections, displays on a given host are numbered starting from 0, and the server for
display N listens and accepts connections on port 6000 + N.
*/
func PortForDisplay(display int) int {
  return 6000 + display
}

// Type for the bythe order of the connection
type ConnectionByteorder   uint8
const ConnectionSetupLSB = ConnectionByteorder(0x6c)
const ConnectionSetupMSB = ConnectionByteorder(0x42)

// Information sent by the client at connection setup
type ConnectionInformation struct { 
  Byteorder   ConnectionByteorder // ConnectionSetupLSB or ConnectionSetupMSB
  Unused      PADDING1 
  Major       CARD16              // protocol major version
  Minor       CARD16              // protocol minor version
  NameLength  CARD16
  DataLength  CARD16
  Unused2     uint16
  Name        STRING8
  Padding1    STRING8
  Data        STRING8
  Padding2    STRING8
}

func (s STRING8) Bytes() ([]byte) {
  str := (string)(s)
  return ([]byte)(str)
}

func (s STRING8) Len() int {
  return len(s)
}

func (s STRING8) LenC8() CARD8 {
  return CARD8(s.Len())
}

func (s STRING8) LenC16() CARD16 {
  return CARD16(s.Len())
}

func (s STRING8) LenC32() CARD32 {
  return CARD32(s.Len())
}


func (i * ConnectionInformation) ToX(c io.Writer) (os.Error)  {
  i.NameLength = i.Name.LenC16()
  i.DataLength = i.Data.LenC16()   
  err := Pack(c, i.Byteorder, i.Unused, i.Major, i.Minor, i.NameLength, 
              i.DataLength, i.NameLength, i.Unused2, 
              i.Name.Bytes(), i.Data.Bytes())
  return err
} 


type  ConnectionReplyCode           uint8
const ConnectionFailedCode        = ConnectionReplyCode(0)
const ConnectionSuccessCode       = ConnectionReplyCode(1)
const ConnectionAuthenticateCode  = ConnectionReplyCode(2)

type GenericReply struct { 
  Code          byte
  Size          byte
  Sequence      CARD16
  Length        CARD32
  Padding1  [24]PADDING 
  // The padding is sometimes used as data in specific replies.
  Data        []byte
  // len(data) == Size * 4, often padded at the end
}


type ConnectionGenericReply struct {
  Code          ConnectionReplyCode
  ReasonLength  uint8
  Major         CARD16              // protocol major version
  Minor         CARD16              // protocol minor version
  DataLength    uint16
}

type ConnectionFailedReply struct {
  ConnectionGenericReply
  Reason        STRING8
  Padding       STRING8
}

type ConnectionAuthenticateReply ConnectionFailedReply


type  ImageByteOrderCode uint8
const ImageLSB          = ImageByteOrderCode(0)
const ImageMSB          = ImageByteOrderCode(1)

type  BitmapByteOrderCode uint8
const BitmapLSB         = BitmapByteOrderCode(0)
const BitmapMSB         = BitmapByteOrderCode(1)



type ConnectionOKReply struct {
  ConnectionGenericReply
  Release                   CARD32
  ResourceIDBase            CARD32
  ResourceIDMask            CARD32
  MotionBufferSize          CARD32
  VendorLength              CARD32
  MaximumRequestLength      CARD16
  NumberOfScreens           CARD8
  NumberOfFormats           CARD8
  ImageByteOrder            ImageByteOrderCode
  BitmapByteOrder           BitmapByteOrderCode
  BitmapFormatScanLineUnit  CARD8
  BitmapFormatScanLinePad   CARD8
  MinKeycode                KEYCODE
  MaxKeyCode                KEYCODE
  Unused                    uint32
  Vendor                    STRING8
  PaddingVendor             STRING8
  PixmapFormats             []FORMAT
  Roots                     []SCREEN  
}

// Reads in padding and discards it. Padding size is based on size. 
func ReadPadding(r io.Reader, size int) (os.Error) {
  padsize := size % 4
  if padsize == 0 { return nil }
  buf := make([]byte, padsize)
  _, err := r.Read(buf)
  return err
}  
  
  

// reads in a string8 with the given length.   
func STRING8FromX(r io.Reader, size int) (STRING8, os.Error) {
  buf := make([]byte, size)
  _ , err := r.Read(buf)
  err  = ReadPadding(r, size) 
  return (STRING8)(buf), err
}

// reads in the non-generic part of the ConnectionOKReply
func (i * ConnectionOKReply) FromX(r io.Reader) (os.Error) {
  err := Unpack(r, i.Release        , i.ResourceIDBase, 
                   i.ResourceIDMask , i.MotionBufferSize, 
                   i.VendorLength   , i.MaximumRequestLength,
                   i.NumberOfScreens, i.NumberOfFormats,
                   i.ImageByteOrder , i.BitmapFormatScanLineUnit,
                   i.BitmapFormatScanLineUnit,
                   i.BitmapFormatScanLinePad,
                   i.MinKeycode,
                   i.MaxKeyCode,
                   i.Unused)
  i.Vendor, err = STRING8FromX(r, int(i.VendorLength))
  return err
}

// Pad a length to align l on a bytes.
func PadLengthTo(l int, a int ) int { 
  return (l + (a-1)) & ^(a-1)
}

// Pad a length to align on 4 bytes
func PadLength4(l int) int {
  return PadLengthTo(l, 4)
} 


// The real McCoy, that is, the real connection to the X server
type Connection struct {
  net.Conn            // The connection. Inherit Read and Write from it. 
  os.Error            // Any errors in the connection
  Display string      // the display string 
  Host    string      // Host name
  // bufio.Writer        // Write buffer 
}

func ConnectLocal() (*Connection) {
  c := &Connection{}
  c.Conn, c.Error = net.Dial("unix", "",  "/tmp/.X11-unix/X0");
  c.Display       = "0"
  c.Host          = "localhost"
  return c
}

// String format
func (c * Connection) String() string {
  return fmt.Sprintf("Conection: %s : %s : %s : %s", c.Conn, c.Error, c.Display, c.Host)
}

// Sets this connection to error status
func (c * Connection) Fail(message string) (*Connection) {
  c.Error = os.NewError(message)
  return c
}
 

// Returns true if the connection is OK, false if not.
func (c * Connection) OK() (bool) {
  return c.Error != nil
}

// Returns false if the connection is OK, true if not.
func (c * Connection) Failed() (bool) {
  // Also works for nil connections
  if c == nil { return true }  
  return c.Error != nil
}

// Closes the connection.
func (c * Connection) Close() (*Connection) {
  if c.Failed() { return c }
  c.Error = c.Conn.Close();
  return c
}

// Authenticate the connection
func (c * Connection) Authenticate() (*Connection) {
  if c.Failed() { return c }
  // Get authentication data
  println("readauthority")
  auth, err := ReadAuthority(c.Host, c.Display)
  if err != nil { return c.Fail(err.String()) }
  
  // Assume that the authentication protocol is "MIT-MAGIC-COOKIE-1".
  if auth.Name != "MIT-MAGIC-COOKIE-1" || len(auth.Data) != 16 {
    return c.Fail("unsupported auth protocol" + auth.Name)
  }
  
  info := &ConnectionInformation{}
  info.Byteorder  = ConnectionSetupLSB
  info.Major      = 11
  info.Minor      = 0
  info.Name       = (STRING8)(auth.Name)
  info.Data       = (STRING8)(auth.Data)
  // set up connection information
  println(info.Name)
  println(info.Data)
  println("write connection info")
  
  c.Error         = info.ToX(c)
  println(c.Error)
  if c.Error != nil { return c.Fail("Send ConnectionInfo Failed.") } 
  // send it to X
  reply := &ConnectionGenericReply{}
  Unpack(c, reply)
  
  if reply.Major != 11 || reply.Minor != 0 {
    return c.Fail(fmt.Sprintf("x protocol version mismatch: %d.%d", 
    reply.Major, reply.Minor))
  }
  
  fmt.Println(reply, reply.Code)
  if reply.Code != 1 { 
    return c.Fail("Cionnection failed");  
  } 
  
  okreply := &ConnectionOKReply{}
  okreply.ConnectionGenericReply = *reply
  okreply.FromX(c)
  fmt.Println(okreply.Vendor, okreply)
  
  /*
  buf = make([]byte, int(dataLen)*4+8, int(dataLen)*4+8)
  copy(buf, head)
  if _, err = io.ReadFull(c.conn, buf[8:]); err != nil {
    return nil, err
  }

  if code == 0 {
    reason := buf[8 : 8+reasonLen]
    return nil, os.NewError(fmt.Sprintf("x protocol authentication refused: %s", string(reason)))
  }

  getSetupInfo(buf, &c.Setup)

  if c.defaultScreen >= len(c.Setup.Roots) {
    c.defaultScreen = 0
  }

  c.nextId = Id(c.Setup.ResourceIdBase)
  c.nextCookie = 1
  c.replies = make(map[Cookie][]byte)
  c.events = queue{make([][]byte, 100), 0, 0}
  */
  return c

}





type Sendable interface {
  // Send the data to a butes.Buffer which will be send to X.
  ToX(bytes.Buffer) 
} 

type Receiveable interface {
  // Allows itself to be filled on from a byte array with the data 
  // received from X
  FromX([]byte) 
} 









