package x

import "net"
import "os"

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
  Unused      uint8 
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

type  ConnectionReplyCode           uint8
const ConnectionFailedCode        = ConnectionReplyCode(0)
const ConnectionSuccessCode       = ConnectionReplyCode(1)
const ConnectionAuthenticateCode  = ConnectionReplyCode(2)

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


// The real McCoy, that is, the real connection to the X server
type Connection struct {
  net.Conn // the connection
  os.Error // Any errors in the connection 
}

func ConnectLocal() (*Connection) {
  c := &Connection{}
  c.Conn, c.Error = net.Dial("unix", "",  "/tmp/.X11-unix/X0");
  return c
}

// Opens the connection by sending a 
// ConnectionInformation to the socket and reading the reply
 
 

// Returns true if the connection is OK, false if not.
func (c * Connection) OK() (bool) {
  return c.Error != nil
}

// Returns false if the connection is OK, true if not.
func (c * Connection) Fail() (bool) {
  return c.Error != nil
}

// Closes the connection.
func (c * Connection) Close() {
  if c.Fail() {
    return
  }   
  c.Error = c.Conn.Close();
}




















