package x


type ID      uint32
type WINDOW  ID
type PIXMAP  ID
type CURSOR       ID
type FONT         ID
type GCONTEXT     ID
type COLORMAP     ID
type DRAWABLE     ID
type FONTABLE     ID
type ATOM         ID
type VISUALID     ID
type VALUE        uint32
type BYTE         uint8
type INT8         int8
type INT16        int8
type INT32        int32
type CARD8        uint8
type CARD16       uint16
type CARD32       uint32
type TIMESTAMP    CARD32

             {Forget, Static, NorthWest, North, NorthEast, West, Center,
BITGRAVITY
              East, SouthWest, South, SouthEast}
             {Unmap, Static, NorthWest, North, NorthEast, West, Center,
WINGRAVITY
              East, SouthWest, South, SouthEast}
             {True, False}
BOOL
             {KeyPress, KeyRelease, OwnerGrabButton, ButtonPress,
EVENT
              ButtonRelease, EnterWindow, LeaveWindow, PointerMotion,
              PointerMotionHint, Button1Motion, Button2Motion,
              Button3Motion, Button4Motion, Button5Motion, ButtonMotion,
              Exposure, VisibilityChange, StructureNotify, ResizeRedirect,
              SubstructureNotify, SubstructureRedirect, FocusChange,
              PropertyChange, ColormapChange, KeymapState}
             {ButtonPress, ButtonRelease, EnterWindow, LeaveWindow,
POINTEREVENT
              PointerMotion, PointerMotionHint, Button1Motion,
              Button2Motion, Button3Motion, Button4Motion, Button5Motion,
              ButtonMotion, KeymapState}
             {KeyPress, KeyRelease, ButtonPress, ButtonRelease,
DEVICEEVENT
              PointerMotion, Button1Motion, Button2Motion, Button3Motion,
              Button4Motion, Button5Motion, ButtonMotion}
KEYSYM       ID
KEYCODE      CARD8
BUTTON       CARD8
             {Shift, Lock, Control, Mod1, Mod2, Mod3, Mod4, Mod5}
KEYMASK
             {Button1, Button2, Button3, Button4, Button5}
BUTMASK
KEYBUTMASK   KEYMASK or BUTMASK
                                      3
Name      Value
STRING8   LISTofCARD8
STRING16  LISTofCHAR2B
CHAR2B    [byte1, byte2: CARD8]
POINT     [x, y: INT16]
RECTANGLE [x, y: INT16,
           width, height: CARD16]
ARC       [x, y: INT16,
           width, height: CARD16,
           angle1, angle2: INT16]
          [family: {Internet, InternetV6, ServerInterpreted, DECnet, Chaos}
HOST
           address: LISTofBYTE]
