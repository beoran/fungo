// Wrappers for OpenGL
// We won't wrap the 2.1 API, since a lot of it has become obsoleted 
// by 3.2 and onwards. 
// Only the API that is available in 2.x, but still supported 
// in the latest versions of openGL will be wrapped.
package gl

//#include <GL/gl.h>
import "C"


// Base GL types
type GLenum     C.GLenum
type GLboolean  C.GLboolean
type GLbitfield C.GLbitfield
type GLbyte     C.GLbyte
type GLshort    C.GLshort
type GLint      C.GLint
type GLsizei    C.GLsizei
type GLubyte    C.GLubyte
type GLushort   C.GLushort
type GLuint     C.GLuint
type GLfloat    C.GLfloat
type GLclampf   C.GLclampf
type GLdouble   C.GLdouble
type GLclampd   C.GLclampd
// typedef void GLvoid;




















