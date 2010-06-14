package x


type VALUE        uint32
type BYTE         uint8
type INT8         int8
type INT16        int8
type INT32        int32
type CARD8        uint8
type CARD16       uint16
type CARD32       uint32
type ID           CARD32
type BITMASK      CARD32
type WINDOW       ID
type PIXMAP       ID
type CURSOR       ID
type FONT         ID
type GCONTEXT     ID
type COLORMAP     ID
type DRAWABLE     ID
type FONTABLE     ID
type ATOM         ID
type VISUALID     ID
type TIMESTAMP    CARD32


type BITGRAVITY int
//             {Forget, Static, NorthWest, North, NorthEast, West, Center,
// BITGRAVITY
//              East, SouthWest, South, SouthEast}
//             {Unmap, Static, NorthWest, North, NorthEast, West, Center,
type WINGRAVITY int
//              East, SouthWest, South, SouthEast}
//             {True, False}
type BOOL int
type EVENT int

/*           {KeyPress, KeyRelease, OwnerGrabButton, ButtonPress,

              ButtonRelease, EnterWindow, LeaveWindow, PointerMotion,
              PointerMotionHint, Button1Motion, Button2Motion,
              Button3Motion, Button4Motion, Button5Motion, ButtonMotion,
              Exposure, VisibilityChange, StructureNotify, ResizeRedirect,
              SubstructureNotify, SubstructureRedirect, FocusChange,
              PropertyChange, ColormapChange, KeymapState}
             {ButtonPress, ButtonRelease, EnterWindow, LeaveWindow,
*/             
type POINTEREVENT int
/*
              PointerMotion, PointerMotionHint, Button1Motion,
              Button2Motion, Button3Motion, Button4Motion, Button5Motion,
              ButtonMotion, KeymapState}
             {KeyPress, KeyRelease, ButtonPress, ButtonRelease,
*/
type DEVICEEVENT int
/*
              PointerMotion, Button1Motion, Button2Motion, Button3Motion,
              Button4Motion, Button5Motion, ButtonMotion}
*/
type KEYSYM       ID
type KEYCODE      CARD8
type BUTTON       CARD8
//             {Shift, Lock, Control, Mod1, Mod2, Mod3, Mod4, Mod5}
type KEYMASK      int 
//             {Button1, Button2, Button3, Button4, Button5}
type BUTMASK      int 
type KEYBUTMASK   int
type STRING8      []CARD8
type CHAR2B       struct { byte1 CARD8 ; byte2 CARD8 }
type STRING16     []CHAR2B
type POINT        struct { x INT16 ; y INT16 } 
type RECTANGLE    struct { x INT16 ; y INT16 ; width CARD16;  height CARD16 }

type ARC          struct {
  x       INT16  
  y       INT16  
  width   CARD16  
  height  CARD16 
  angle1  INT16 
  angle2  INT16
}  



// constants for bitgravity and or wingravity
const ( 
  Forget    =   iota // 0
  NorthWest // 1 
  North
  NorthEast
  West
  Center
  East
  SouthWest
  South
  SouthEast
  Static  // untyil 10
)

// for wingravity  
const (
  Unmap   = 0
)
  
// For BOOL
const (
  False = BOOL(0)
  True  = BOOL(1)
)

type SETofEVENT uint32

// for SETofEVENT
const (
  KeyPress            = SETofEVENT(0x00000001)
  KeyRelease          = SETofEVENT(0x00000002)
  ButtonPress         = SETofEVENT(0x00000004)
  ButtonRelease       = SETofEVENT(0x00000008)
  EnterWindow         = SETofEVENT(0x00000010)
  LeaveWindow         = SETofEVENT(0x00000020)
  PointerMotion       = SETofEVENT(0x00000040)
  PointerMotionHint   = SETofEVENT(0x00000080)
  Button1Motion       = SETofEVENT(0x00000100)
  Button2Motion       = SETofEVENT(0x00000200)
  Button3Motion       = SETofEVENT(0x00000400)
  Button4Motion       = SETofEVENT(0x00000800)
  Button5Motion       = SETofEVENT(0x00001000)
  ButtonMotion        = SETofEVENT(0x00002000)
  KeymapState         = SETofEVENT(0x00004000)
  Exposure            = SETofEVENT(0x00008000)
  VisibilityChange    = SETofEVENT(0x00010000)
  StructureNotify     = SETofEVENT(0x00020000)
  ResizeRedirect      = SETofEVENT(0x00040000)
  SubstructureNotify  = SETofEVENT(0x00080000)
  SubstructureRedirect= SETofEVENT(0x00100000)
  FocusChange         = SETofEVENT(0x00200000)
  PropertyChange      = SETofEVENT(0x00400000)
  ColormapChange      = SETofEVENT(0x00800000)
  OwnerGrabButton     = SETofEVENT(0x01000000)
  UnusedEventMask     = SETofEVENT(0xFE000000)
  UnusedPointerEventMask  = SETofEVENT(0xFFFF8003)
  UnusedDeviceEventMask   = SETofEVENT(0xFFFF8003)
)
  
type SETofKEYBUTMASK uint16
  
// SETofKEYBUTMASK  
const (  
    Shift   = SETofKEYBUTMASK(0x0001)
    Lock    = SETofKEYBUTMASK(0x0002)
    Control = SETofKEYBUTMASK(0x0004)
    Mod1    = SETofKEYBUTMASK(0x0008)
    Mod2    = SETofKEYBUTMASK(0x0010)
    Mod3    = SETofKEYBUTMASK(0x0020)
    Mod4    = SETofKEYBUTMASK(0x0040)
    Mod5    = SETofKEYBUTMASK(0x0080)
    Button1 = SETofKEYBUTMASK(0x0100)
    Button2 = SETofKEYBUTMASK(0x0200)
    Button3 = SETofKEYBUTMASK(0x0400)
    Button4 = SETofKEYBUTMASK(0x0800)
    Button5 = SETofKEYBUTMASK(0x1000)
    UnusedKeyButMask  = SETofKEYBUTMASK(0xE000)
    UnusedKeyMask     = SETofKEYBUTMASK(0xFF00)
)

type Family CARD8
 
// FAMILY cnstants
const (
  Internet           = Family(0)
  DECnet             = Family(1) 
  Chaos              = Family(2)
  ServerInterpreted  = Family(5)
  InternetV6         = Family(6)
) 
 
type HOST struct { 
  Family  Family
  Unused  uint8
  Length  uint16
  Address []BYTE
  Padding []BYTE
}

type STR struct { 
  Length  uint8
  Name    STRING8
}

type GenericError struct {
  Error     uint8     // Always 0
  Code      uint8     // Error code
  Sequence  CARD16    // Sequence ID of erroneous request
  Bad       CARD32    // What exactly was bad about the request
  Minor     CARD16    // Minor opcode
  Opcode    CARD8     // Major Opcode
  Unused    [21]uint8  // Unused padding (for extensions)
}

// Error code constants 
const ( 
  RequestErrorCode        = 1
  ValueErrorCode          = 2
  WindowErrorCode         = 3
  PixmapErrorCode         = 4
  AtomErrorCode           = 5
  CursorErrorCode         = 6
  FontErrorCode           = 7
  MatchErrorCode          = 8
  DrawableErrorCode       = 9
  AccessErrorCode         = 10
  AllocErrorCode          = 11
  ColormapErrorCode       = 12
  GContexErrorCode        = 13
  IDChoiceErrorCode       = 14
  NameErrorCode           = 15
  LengthErrorCode         = 16
  ImplementationErrorCode = 17
)

// Predefined ATOMs 
const (
  ATOM_PRIMARY              = 1
  ATOM_SECONDARY            = 2
  ATOM_ARC                  = 3
  ATOM_ATOM                 = 4
  ATOM_BITMAP               = 5
  ATOM_CARDINAL             = 6
  ATOM_COLORMAP             = 7
  ATOM_CURSOR               = 8
  ATOM_CUT_BUFFER0          = 9
  ATOM_CUT_BUFFER1          = 10
  ATOM_CUT_BUFFER2          = 11
  ATOM_CUT_BUFFER3          = 12
  ATOM_CUT_BUFFER4          = 13
  ATOM_CUT_BUFFER5          = 14
  ATOM_CUT_BUFFER6          = 15
  ATOM_CUT_BUFFER7          = 16
  ATOM_DRAWABLE             = 17
  ATOM_FONT                 = 18
  ATOM_INTEGER              = 19
  ATOM_PIXMAP               = 20
  ATOM_POINT                = 21
  ATOM_RECTANGLE            = 22
  ATOM_RESOURCE_MANAGER     = 23
  ATOM_RGB_COLOR_MAP        = 24
  ATOM_RGB_BEST_MAP         = 25
  ATOM_RGB_BLUE_MAP         = 26
  ATOM_RGB_DEFAULT_MAP      = 27
  ATOM_RGB_GRAY_MAP         = 28
  ATOM_RGB_GREEN_MAP        = 29
  ATOM_RGB_RED_MAP          = 30
  ATOM_STRING               = 31
  ATOM_VISUALID             = 32
  ATOM_WINDOW               = 33
  ATOM_WM_COMMAND           = 34
  ATOM_WM_HINTS             = 35
  ATOM_WM_CLIENT_MACHINE    = 36
  ATOM_WM_ICON_NAME         = 37
  ATOM_WM_ICON_SIZE         = 38
  ATOM_WM_NAME              = 39
  ATOM_WM_NORMAL_HINTS      = 40
  ATOM_WM_SIZE_HINTS        = 41
  ATOM_WM_ZOOM_HINTS        = 42
  ATOM_MIN_SPACE            = 43
  ATOM_NORM_SPACE           = 44
  ATOM_MAX_SPACE            = 45
  ATOM_END_SPACE            = 46
  ATOM_SUPERSCRIPT_X        = 47
  ATOM_SUPERSCRIPT_Y        = 48
  ATOM_SUBSCRIPT_X          = 49
  ATOM_SUBSCRIPT_Y          = 50
  ATOM_UNDERLINE_POSITION   = 51
  ATOM_UNDERLINE_THICKNESS  = 52
  ATOM_STRIKEOUT_ASCENT     = 53
  ATOM_STRIKEOUT_DESCENT    = 54
  ATOM_ITALIC_ANGLE         = 55
  ATOM_X_HEIGHT             = 56
  ATOM_QUAD_WIDTH           = 57
  ATOM_WEIGHT               = 58
  ATOM_POINT_SIZE           = 59
  ATOM_RESOLUTION           = 60
  ATOM_COPYRIGHT            = 61
  ATOM_NOTICE               = 62
  ATOM_FONT_NAME            = 63  
  ATOM_FAMILY_NAME          = 64
  ATOM_FULL_NAME            = 65
  ATOM_CAP_HEIGHT           = 66
  ATOM_WM_CLASS             = 67
  ATOM_WM_TRANSIENT_FOR     = 68
)

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


type FORMAT struct { 
  Depth                     CARD8
  BitsPerPixel              CARD8
  ScanlinePad               CARD8
  Unused                 [5]CARD8 // 5 unused bytes
}

type ScreenBackingStores uint8
const ScreenBackingNever      = ScreenBackingStores(0)
const ScreenBackingWhenMapped = ScreenBackingStores(1)
const ScreenBackingAlways     = ScreenBackingStores(2)


type SCREEN struct {
  Root                      WINDOW
  Colormap                  COLORMAP
  White                     CARD32
  Black                     CARD32
  CurrentInputMasks         SETofEVENT
  WidthInPixels             CARD16
  HeightInPixels            CARD16
  WidthInMillimeters        CARD16  
  HeightInMillimeters       CARD16
  MinInstalledMaps          CARD16
  MaxInstalledMaps          CARD16
  RootVisual                VISUALID
  Backing                   ScreenBackingStores
  SaveUnders                BOOL
  RootDepth                 CARD8
  DepthLength               CARD8
  AllowedDepths           []DEPTH  
}


type DEPTH struct {
  Depth             CARD8
  Unused            uint8
  VisualLength      CARD16
  Unused2           uint32
  Visuals         []VISUALTYPE 
}

type VisualTypeClass uint8
const StaticGray      = VisualTypeClass(0)
const GrayScale       = VisualTypeClass(1)
const StaticColor     = VisualTypeClass(2)
const PseudoColor     = VisualTypeClass(3)
const TrueColor       = VisualTypeClass(4)
const DirectColor     = VisualTypeClass(5)


type VISUALTYPE {
  VisualID          VISUALID
  Class             VisualTypeClass
  BitsPerRgbValue   CARD8
  RedMask           CARD16
  GreenMask         CARD16
  BlueMask          CARD16
  Unused            uint16
}

type RequestOpcode  uint8

const ( 
  RequestOpcodeCreateWindow             = RequestOpcode(iota + 1) // 1
  RequestOpcodeChangeWindowAttributes   // 2 , .. etc ...
  RequestOpcodeGetWindowAttributes
  RequestOpcodeDestroyWindow
  RequestOpcodeDestroySubwindows
  RequestOpcodeChangeSaveSet
  RequestOpcodeReparentWindow
  RequestOpcodeMapWindow
  RequestOpcodeMapSubwindows
  RequestOpcodeUnmapWindow
  RequestOpcodeUnmapSubwindows
  RequestOpcodeConﬁgureWindow
 /* 
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
  RequestOpcode
*/
)

// struct for a generic request to the X server
type Request struct {
  Opcode RequestOpcode
  Minor  uint8
  Length uint16   
  Data   []uint8
}





/*
Requests
CreateWindow
   1   1                            opcode
   1   CARD8                        depth
   2   8+n                          request length
   4   WINDOW                       wid
   4   WINDOW                       parent
   2   INT16                        x
   2   INT16                        y
   2   CARD16                       width
   2   CARD16                       height
   2   CARD16                       border-width
   2                                class
       0          CopyFromParent
       1          InputOutput
       2          InputOnly
   4   VISUALID                     visual
       0          CopyFromParent
   4   BITMASK                      value-mask (has n bits set to 1)
       #x00000001 background-pixmap
       #x00000002 background-pixel
       #x00000004 border-pixmap
       #x00000008 border-pixel
       #x00000010 bit-gravity
       #x00000020 win-gravity
       #x00000040 backing-store
       #x00000080 backing-planes
       #x00000100 backing-pixel
       #x00000200 override-redirect
       #x00000400 save-under
       #x00000800 event-mask
       #x00001000 do-not-propagate-mask
       #x00002000 colormap
       #x00004000 cursor
   4n LISTofVALUE                   value-list
 VALUEs
   4   PIXMAP                       background-pixmap
       0          None
       1          ParentRelative
   4   CARD32                       background-pixel
   4   PIXMAP                       border-pixmap
       0          CopyFromParent
   4   CARD32                       border-pixel
   1   BITGRAVITY                   bit-gravity
   1   WINGRAVITY                   win-gravity
   1                                backing-store
                                             120
        0                   NotUseful
        1                   WhenMapped
        2                   Always
   4    CARD32                                backing-planes
   4    CARD32                                backing-pixel
   1    BOOL                                  override-redirect
   1    BOOL                                  save-under
   4    SETofEVENT                            event-mask
   4    SETofDEVICEEVENT                      do-not-propagate-mask
   4    COLORMAP                              colormap
        0                   CopyFromParent
   4    CURSOR                                cursor
        0                   None
ChangeWindowAttributes
   1    2                                     opcode
   1                                          unused
   2    3+n                                   request length
   4    WINDOW                                window
   4    BITMASK                               value-mask (has n bits set to 1)
        encodings are the same as for CreateWindow
   4n LISTofVALUE                             value-list
        encodings are the same as for CreateWindow
GetWindowAttributes
   1    3                                     opcode
   1                                          unused
   2    2                                     request length
   4    WINDOW                                window
→
   1    1                                     Reply
   1                                          backing-store
        0                   NotUseful
        1                   WhenMapped
        2                   Always
   2    CARD16                                sequence number
   4    3                                     reply length
   4    VISUALID                              visual
   2                                          class
        1                   InputOutput
        2                   InputOnly
   1    BITGRAVITY                            bit-gravity
   1    WINGRAVITY                            win-gravity
   4    CARD32                                backing-planes
   4    CARD32                                backing-pixel
   1    BOOL                                  save-under
   1    BOOL                                  map-is-installed
   1                                          map-state
        0                   Unmapped
        1                   Unviewable
        2                   Viewable
   1    BOOL                                  override-redirect
   4    COLORMAP                              colormap
        0                   None
   4    SETofEVENT                            all-event-masks
   4    SETofEVENT                            your-event-mask
   2    SETofDEVICEEVENT                      do-not-propagate-mask
   2                                          unused
DestroyWindow
   1    4                                     opcode
   1                                          unused
                                                       121
   2    2                request length
   4    WINDOW           window
DestroySubwindows
   1    5                opcode
   1                     unused
   2    2                request length
   4    WINDOW           window
ChangeSaveSet
   1    6                opcode
   1                     mode
        0         Insert
        1         Delete
   2    2                request length
   4    WINDOW           window
ReparentWindow
   1    7                opcode
   1                     unused
   2    4                request length
   4    WINDOW           window
   4    WINDOW           parent
   2    INT16            x
   2    INT16            y
MapWindow
   1    8                opcode
   1                     unused
   2    2                request length
   4    WINDOW           window
MapSubwindows
   1    9                opcode
   1                     unused
   2    2                request length
   4    WINDOW           window
UnmapWindow
   1    10               opcode
   1                     unused
   2    2                request length
   4    WINDOW           window
UnmapSubwindows
   1    11               opcode
   1                     unused
   2    2                request length
   4    WINDOW           window
ConﬁgureWindow
   1    12               opcode
   1                     unused
   2    3+n              request length
   4    WINDOW           window
   2    BITMASK          value-mask (has n bits set to 1)
        #x0001    x
        #x0002    y
        #x0004    width
        #x0008    height
                                 122
         #x0010      border-width
         #x0020      sibling
         #x0040      stack-mode
   2                              unused
   4n    LISTofVALUE              value-list
 VALUEs
   2     INT16                    x
   2     INT16                    y
   2     CARD16                   width
   2     CARD16                   height
   2     CARD16                   border-width
   4     WINDOW                   sibling
   1                              stack-mode
         0           Above
         1           Below
         2           TopIf
         3           BottomIf
         4           Opposite
CirculateWindow
   1     13                       opcode
   1                              direction
         0           RaiseLowest
         1           LowerHighest
   2     2                        request length
   4     WINDOW                   window
GetGeometry
   1     14                       opcode
   1                              unused
   2     2                        request length
   4     DRAWABLE                 drawable
→
   1     1                        Reply
   1     CARD8                    depth
   2     CARD16                   sequence number
   4     0                        reply length
   4     WINDOW                   root
   2     INT16                    x
   2     INT16                    y
   2     CARD16                   width
   2     CARD16                   height
   2     CARD16                   border-width
   10                             unused
QueryTree
   1     15                       opcode
   1                              unused
   2     2                        request length
   4     WINDOW                   window
→
   1     1                        Reply
   1                              unused
   2     CARD16                   sequence number
   4     n                        reply length
   4     WINDOW                   root
   4     WINDOW                   parent
         0           None
   2     n                        number of WINDOWs in children
                                           123
    14                       unused
    4n  LISTofWINDOW         children
InternAtom
    1   16                   opcode
    1   BOOL                 only-if-exists
    2   2+(n+p)/4            request length
    2   n                    length of name
    2                        unused
    n   STRING8              name
    p                        unused, p=pad(n)
→
    1   1                    Reply
    1                        unused
    2   CARD16               sequence number
    4   0                    reply length
    4   ATOM                 atom
        0            None
    20                       unused
GetAtomName
    1   17                   opcode
    1                        unused
    2   2                    request length
    4   ATOM                 atom
→
    1   1                    Reply
    1                        unused
    2   CARD16               sequence number
    4   (n+p)/4              reply length
    2   n                    length of name
    22                       unused
    n   STRING8              name
    p                        unused, p=pad(n)
ChangeProperty
    1   18                   opcode
    1                        mode
        0            Replace
        1            Prepend
        2            Append
    2   6+(n+p)/4            request length
    4   WINDOW               window
    4   ATOM                 property
    4   ATOM                 type
    1   CARD8                format
    3                        unused
    4   CARD32               length of data in format units
                             (= n for format = 8)
                             (= n/2 for format = 16)
                             (= n/4 for format = 32)
    n   LISTofBYTE           data
                             (n is a multiple of 2 for format = 16)
                             (n is a multiple of 4 for format = 32)
    p                        unused, p=pad(n)
DeleteProperty
    1   19                   opcode
    1                        unused
    2   3                    request length
                                      124
    4    WINDOW                     window
    4    ATOM                       property
GetProperty
    1    20                         opcode
    1    BOOL                       delete
    2    6                          request length
    4    WINDOW                     window
    4    ATOM                       property
    4    ATOM                       type
         0          AnyPropertyType
    4    CARD32                     long-offset
    4    CARD32                     long-length
→
    1    1                          Reply
    1    CARD8                      format
    2    CARD16                     sequence number
    4    (n+p)/4                    reply length
    4    ATOM                       type
         0          None
    4    CARD32                     bytes-after
    4    CARD32                     length of value in format units
                                    (= 0 for format = 0)
                                    (= n for format = 8)
                                    (= n/2 for format = 16)
                                    (= n/4 for format = 32)
    12                              unused
    n    LISTofBYTE                 value
                                    (n is zero for format = 0)
                                    (n is a multiple of 2 for format = 16)
                                    (n is a multiple of 4 for format = 32)
    p                               unused, p=pad(n)
ListProperties
    1    21                         opcode
    1                               unused
    2    2                          request length
    4    WINDOW                     window
→
    1    1                          Reply
    1                               unused
    2    CARD16                     sequence number
    4    n                          reply length
    2    n                          number of ATOMs in atoms
    22                              unused
    4n   LISTofATOM                 atoms
SetSelectionOwner
    1    22                         opcode
    1                               unused
    2    4                          request length
    4    WINDOW                     owner
         0          None
    4    ATOM                       selection
    4    TIMESTAMP                  time
         0          CurrentTime
GetSelectionOwner
    1    23                         opcode
    1                               unused
                                             125
   2    2                                       request length
   4    ATOM                                    selection
→
   1    1                                       Reply
   1                                            unused
   2    CARD16                                  sequence number
   4    0                                       reply length
   4    WINDOW                                  owner
        0                  None
   20                                           unused
ConvertSelection
   1    24                                      opcode
   1                                            unused
   2    6                                       request length
   4    WINDOW                                  requestor
   4    ATOM                                    selection
   4    ATOM                                    target
   4    ATOM                                    property
        0                  None
   4    TIMESTAMP                               time
        0                  CurrentTime
SendEvent
   1    25                                      opcode
   1    BOOL                                    propagate
   2    11                                      request length
   4    WINDOW                                  destination
        0                  PointerWindow
        1                  InputFocus
   4    SETofEVENT                              event-mask
   32                                           event
        standard event format (see the Events section)
GrabPointer
   1    26                                      opcode
   1    BOOL                                    owner-events
   2    6                                       request length
   4    WINDOW                                  grab-window
   2    SETofPOINTEREVENT                       event-mask
   1                                            pointer-mode
        0                  Synchronous
        1                  Asynchronous
   1                                            keyboard-mode
        0                  Synchronous
        1                  Asynchronous
   4    WINDOW                                  conﬁne-to
        0                  None
   4    CURSOR                                  cursor
        0                  None
   4    TIMESTAMP                               time
        0                  CurrentTime
→
   1    1                                       Reply
   1                                            status
        0                  Success
        1                  AlreadyGrabbed
        2                  InvalidTime
        3                  NotViewable
        4                  Frozen
   2    CARD16                                  sequence number
                                                         126
   4   0                             reply length
   24                                unused
UngrabPointer
   1   27                            opcode
   1                                 unused
   2   2                             request length
   4   TIMESTAMP                     time
       0                CurrentTime
GrabButton
   1   28                            opcode
   1   BOOL                          owner-events
   2   6                             request length
   4   WINDOW                        grab-window
   2   SETofPOINTEREVENT             event-mask
   1                                 pointer-mode
       0                Synchronous
       1                Asynchronous
   1                                 keyboard-mode
       0                Synchronous
       1                Asynchronous
   4   WINDOW                        conﬁne-to
       0                None
   4   CURSOR                        cursor
       0                None
   1   BUTTON                        button
       0                AnyButton
   1                                 unused
   2   SETofKEYMASK                  modiﬁers
       #x8000           AnyModiﬁer
UngrabButton
   1   29                            opcode
   1   BUTTON                        button
       0                AnyButton
   2   3                             request length
   4   WINDOW                        grab-window
   2   SETofKEYMASK                  modiﬁers
       #x8000           AnyModiﬁer
   2                                 unused
ChangeActivePointerGrab
   1   30                            opcode
   1                                 unused
                                     request length
   2   4
   4   CURSOR                        cursor
       0                None
   4   TIMESTAMP                     time
       0                CurrentTime
   2   SETofPOINTEREVENT             event-mask
   2                                 unused
GrabKeyboard
   1   31                            opcode
   1   BOOL                          owner-events
   2   4                             request length
   4   WINDOW                        grab-window
   4   TIMESTAMP                     time
       0                CurrentTime
   1                                 pointer-mode
       0                Synchronous
                                             127
        1            Asynchronous
   1                                keyboard-mode
        0            Synchronous
        1            Asynchronous
   2                                unused
→
   1    1                           Reply
   1                                status
        0            Success
        1            AlreadyGrabbed
        2            InvalidTime
        3            NotViewable
        4            Frozen
   2    CARD16                      sequence number
   4    0                           reply length
   24                               unused
UngrabKeyboard
   1    32                          opcode
   1                                unused
   2    2                           request length
   4    TIMESTAMP                   time
        0            CurrentTime
GrabKey
   1    33                          opcode
   1    BOOL                        owner-events
   2    4                           request length
   4    WINDOW                      grab-window
   2    SETofKEYMASK                modiﬁers
        #x8000       AnyModiﬁer
   1    KEYCODE                     key
        0            AnyKey
   1                                pointer-mode
        0            Synchronous
        1            Asynchronous
   1                                keyboard-mode
        0            Synchronous
        1            Asynchronous
   3                                unused
UngrabKey
   1    34                          opcode
   1    KEYCODE                     key
        0            AnyKey
   2    3                           request length
   4    WINDOW                      grab-window
   2    SETofKEYMASK                modiﬁers
        #x8000       AnyModiﬁer
   2                                unused
AllowEvents
   1    35                          opcode
   1                                mode
        0            AsyncPointer
        1            SyncPointer
        2            ReplayPointer
        3            AsyncKeyboard
        4            SyncKeyboard
        5            ReplayKeyboard
        6            AsyncBoth
                                            128
        7             SyncBoth
   2    2                         request length
   4    TIMESTAMP                 time
        0             CurrentTime
GrabServer
   1    36                        opcode
   1                              unused
   2    1                         request length
UngrabServer
   1    37                        opcode
   1                              unused
   2    1                         request length
QueryPointer
   1    38                        opcode
   1                              unused
   2    2                         request length
   4    WINDOW                    window
→
   1    1                         Reply
   1    BOOL                      same-screen
   2    CARD16                    sequence number
   4    0                         reply length
   4    WINDOW                    root
   4    WINDOW                    child
        0             None
   2    INT16                     root-x
   2    INT16                     root-y
   2    INT16                     win-x
   2    INT16                     win-y
   2    SETofKEYBUTMASK           mask
   6                              unused
GetMotionEvents
   1    39                        opcode
   1                              unused
   2    4                         request length
   4    WINDOW                    window
   4    TIMESTAMP                 start
        0             CurrentTime
   4    TIMESTAMP                 stop
        0             CurrentTime
→
   1    1                         Reply
   1                              unused
   2    CARD16                    sequence number
   4    2n                        reply length
   4    n                         number of TIMECOORDs in events
   20                             unused
   8n   LISTofTIMECOORD           events
 TIMECOORD
   4    TIMESTAMP                 time
   2    INT16                     x
   2    INT16                     y
                                          129
TranslateCoordinates
   1     40                      opcode
   1                             unused
   2     4                       request length
   4     WINDOW                  src-window
   4     WINDOW                  dst-window
   2     INT16                   src-x
   2     INT16                   src-y
→
   1     1                       Reply
   1     BOOL                    same-screen
   2     CARD16                  sequence number
   4     0                       reply length
   4     WINDOW                  child
         0           None
   2     INT16                   dst-x
   2     INT16                   dst-y
   16                            unused
WarpPointer
   1     41                      opcode
   1                             unused
   2     6                       request length
   4     WINDOW                  src-window
         0           None
   4     WINDOW                  dst-window
         0           None
   2     INT16                   src-x
   2     INT16                   src-y
   2     CARD16                  src-width
   2     CARD16                  src-height
   2     INT16                   dst-x
   2     INT16                   dst-y
SetInputFocus
   1     42                      opcode
   1                             revert-to
         0           None
         1           PointerRoot
         2           Parent
   2     3                       request length
   4     WINDOW                  focus
         0           None
         1           PointerRoot
   4     TIMESTAMP               time
         0           CurrentTime
GetInputFocus
   1     43                      opcode
   1                             unused
   2     1                       request length
→
   1     1                       Reply
   1                             revert-to
         0           None
         1           PointerRoot
         2           Parent
   2     CARD16                  sequence number
   4     0                       reply length
   4     WINDOW                  focus
         0           None
                                          130
        1             PointerRoot
   20                             unused
QueryKeymap
   1    44                        opcode
   1                              unused
   2    1                         request length
→
   1    1                         Reply
   1                              unused
   2    CARD16                    sequence number
   4    2                         reply length
   32   LISTofCARD8               keys
OpenFont
   1    45                        opcode
   1                              unused
   2    3+(n+p)/4                 request length
   4    FONT                      ﬁd
   2    n                         length of name
   2                              unused
   n    STRING8                   name
   p                              unused, p=pad(n)
CloseFont
   1    46                        opcode
   1                              unused
   2    2                         request length
   4    FONT                      font
QueryFont
   1    47                        opcode
   1                              unused
   2    2                         request length
   4    FONTABLE                  font
→
   1    1                         Reply
   1                              unused
   2    CARD16                    sequence number
   4    7+2n+3m                   reply length
   12   CHARINFO                  min-bounds
   4                              unused
   12   CHARINFO                  max-bounds
   4                              unused
   2    CARD16                    min-char-or-byte2
   2    CARD16                    max-char-or-byte2
   2    CARD16                    default-char
   2    n                         number of FONTPROPs in properties
   1                              draw-direction
        0             LeftToRight
        1             RightToLeft
   1    CARD8                     min-byte1
   1    CARD8                     max-byte1
   1    BOOL                      all-chars-exist
   2    INT16                     font-ascent
   2    INT16                     font-descent
   4    m                         number of CHARINFOs in char-infos
   8n LISTofFONTPROP              properties
   12m LISTofCHARINFO             char-infos
                                          131
 FONTPROP
    4    ATOM                   name
    4    <32-bits>              value
 CHARINFO
    2    INT16                  left-side-bearing
    2    INT16                  right-side-bearing
    2    INT16                  character-width
    2    INT16                  ascent
    2    INT16                  descent
    2    CARD16                 attributes
QueryTextExtents
    1    48                     opcode
    1    BOOL                   odd length, True if p = 2
    2    2+(2n+p)/4             request length
    4    FONTABLE               font
    2n STRING16                 string
    p                           unused, p=pad(2n)
→
    1    1                      Reply
    1                           draw-direction
         0          LeftToRight
         1          RightToLeft
    2    CARD16                 sequence number
    4    0                      reply length
    2    INT16                  font-ascent
    2    INT16                  font-descent
    2    INT16                  overall-ascent
    2    INT16                  overall-descent
    4    INT32                  overall-width
    4    INT32                  overall-left
    4    INT32                  overall-right
    4                           unused
ListFonts
    1    49                     opcode
    1                           unused
    2    2+(n+p)/4              request length
    2    CARD16                 max-names
    2    n                      length of pattern
    n    STRING8                pattern
    p                           unused, p=pad(n)
→
    1    1                      Reply
    1                           unused
    2    CARD16                 sequence number
    4    (n+p)/4                reply length
    2    CARD16                 number of STRs in names
    22                          unused
    n    LISTofSTR              names
    p                           unused, p=pad(n)
ListFontsWithInfo
    1    50                     opcode
    1                           unused
    2    2+(n+p)/4              request length
    2    CARD16                 max-names
    2    n                      length of pattern
    n    STRING8                pattern
                                         132
   p                                       unused, p=pad(n)
→ (except for last in series)
   1      1                                Reply
   1      n                                length of name in bytes
   2      CARD16                           sequence number
   4      7+2m+(n+p)/4                     reply length
   12 CHARINFO                             min-bounds
   4                                       unused
   12 CHARINFO                             max-bounds
   4                                       unused
   2      CARD16                           min-char-or-byte2
   2      CARD16                           max-char-or-byte2
   2      CARD16                           default-char
   2      m                                number of FONTPROPs in properties
   1                                       draw-direction
          0                   LeftToRight
          1                   RightToLeft
   1      CARD8                            min-byte1
   1      CARD8                            max-byte1
   1      BOOL                             all-chars-exist
   2      INT16                            font-ascent
   2      INT16                            font-descent
   4      CARD32                           replies-hint
   8m LISTofFONTPROP                       properties
   n      STRING8                          name
   p                                       unused, p=pad(n)
 FONTPROP
   encodings are the same as for QueryFont
 CHARINFO
   encodings are the same as for QueryFont
→ (last in series)
   1      1                                Reply
   1      0                                last-reply indicator
   2      CARD16                           sequence number
   4      7                                reply length
   52                                      unused
SetFontPath
   1      51                               opcode
   1                                       unused
   2      2+(n+p)/4                        request length
   2      CARD16                           number of STRs in path
   2                                       unused
   n      LISTofSTR                        path
   p                                       unused, p=pad(n)
GetFontPath
   1      52                               opcode
   1                                       unused
   2      1                                request list
→
   1      1                                Reply
   1                                       unused
   2      CARD16                           sequence number
   4      (n+p)/4                          reply length
   2      CARD16                           number of STRs in path
   22                                      unused
   n      LISTofSTR                        path
                                                    133
   p                                     unused, p=pad(n)
CreatePixmap
   1    53                               opcode
   1    CARD8                            depth
   2    4                                request length
   4    PIXMAP                           pid
   4    DRAWABLE                         drawable
   2    CARD16                           width
   2    CARD16                           height
FreePixmap
   1    54                               opcode
   1                                     unused
   2    2                                request length
   4    PIXMAP                           pixmap
CreateGC
   1    55                               opcode
   1                                     unused
   2    4+n                              request length
   4    GCONTEXT                         cid
   4    DRAWABLE                         drawable
   4    BITMASK                          value-mask (has n bits set to 1)
        #x00000001 function
        #x00000002 plane-mask
        #x00000004 foreground
        #x00000008 background
        #x00000010 line-width
        #x00000020 line-style
        #x00000040 cap-style
        #x00000080 join-style
        #x00000100 ﬁll-style
        #x00000200 ﬁll-rule
        #x00000400 tile
        #x00000800 stipple
        #x00001000 tile-stipple-x-origin
        #x00002000 tile-stipple-y-origin
        #x00004000 font
        #x00008000 subwindow-mode
        #x00010000 graphics-exposures
        #x00020000 clip-x-origin
        #x00040000 clip-y-origin
        #x00080000 clip-mask
        #x00100000 dash-offset
        #x00200000 dashes
        #x00400000 arc-mode
   4n LISTofVALUE                        value-lis
*/




























