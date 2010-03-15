// Wrappers for OpenGL
// We won't wrap the 2.1 API, since a lot of it has become obsoleted 
// by 3.2 and onwards. 
// Only the API that is available in 2.x, but still supported 
// in the latest versions of openGL will be wrapped.
package gl

//#define GL_GLEXT_PROTOTYPES
//#include <GL/gl.h>
//#include <GL/glext.h>
import "C"
import "unsafe"


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
type Pointer    unsafe.Pointer;

// typedef void GLvoid;


func VertexAttrib1d(index uint, v1 float64) { 
  C.glVertexAttrib1d (C.GLuint(index), C.GLdouble(v1))
}

func VertexAttrib1dv(index uint, v1 * float64) { 
  C.glVertexAttrib1dv(C.GLuint(index), (*C.GLdouble)(Pointer(v1)))
}

func VertexAttrib1f(index uint, v1 float32) { 
  C.glVertexAttrib1f(C.GLuint(index), C.GLfloat(v1))
}

func VertexAttrib1fv(index uint, v1 * float32) { 
  C.glVertexAttrib1fv(C.GLuint(index), (*C.GLfloat)(Pointer(v1)))
}

func VertexAttrib1s(index uint, v1 int16) {
  C.glVertexAttrib1s( C.GLuint(index), C.GLshort(v1))
}

func VertexAttrib1sv(index uint, v1 * int16) { 
  C.glVertexAttrib1sv(C.GLuint(index), (*C.GLshort)(Pointer(v1)))
}


/*


func VertexAttrib1f() { C.glVertexAttrib1f (GLuint, GLfloat)
func VertexAttrib1fv() { C.glVertexAttrib1fv (GLuint, const GLfloat *)
func VertexAttrib1s() { C.glVertexAttrib1s (GLuint, GLshort)
func VertexAttrib1sv() { C.glVertexAttrib1sv (GLuint, const GLshort *)
func VertexAttrib2d() { C.glVertexAttrib2d (GLuint, GLdouble, GLdouble)
func VertexAttrib2dv() { C.glVertexAttrib2dv (GLuint, const GLdouble *)
func VertexAttrib2f() { C.glVertexAttrib2f (GLuint, GLfloat, GLfloat)
func VertexAttrib2fv() { C.glVertexAttrib2fv (GLuint, const GLfloat *)
func VertexAttrib2s() { C.glVertexAttrib2s (GLuint, GLshort, GLshort)
func VertexAttrib2sv() { C.glVertexAttrib2sv (GLuint, const GLshort *)
func VertexAttrib3d() { C.glVertexAttrib3d (GLuint, GLdouble, GLdouble, GLdouble)
func VertexAttrib3dv() { C.glVertexAttrib3dv (GLuint, const GLdouble *)
func VertexAttrib3f() { C.glVertexAttrib3f (GLuint, GLfloat, GLfloat, GLfloat)
func VertexAttrib3fv() { C.glVertexAttrib3fv (GLuint, const GLfloat *)
func VertexAttrib3s() { C.glVertexAttrib3s (GLuint, GLshort, GLshort, GLshort)
func VertexAttrib3sv() { C.glVertexAttrib3sv (GLuint, const GLshort *)
func VertexAttrib4Nbv() { C.glVertexAttrib4Nbv (GLuint, const GLbyte *)
func VertexAttrib4Niv() { C.glVertexAttrib4Niv (GLuint, const GLint *)
func VertexAttrib4Nsv() { C.glVertexAttrib4Nsv (GLuint, const GLshort *)
func VertexAttrib4Nub() { C.glVertexAttrib4Nub (GLuint, GLubyte, GLubyte, GLubyte, GLubyte)
func VertexAttrib4Nubv() { C.glVertexAttrib4Nubv (GLuint, const GLubyte *)
func VertexAttrib4Nuiv() { C.glVertexAttrib4Nuiv (GLuint, const GLuint *)
func VertexAttrib4Nusv() { C.glVertexAttrib4Nusv (GLuint, const GLushort *)
func VertexAttrib4bv() { C.glVertexAttrib4bv (GLuint, const GLbyte *)
func VertexAttrib4d() { C.glVertexAttrib4d (GLuint, GLdouble, GLdouble, GLdouble, GLdouble)
func VertexAttrib4dv() { C.glVertexAttrib4dv (GLuint, const GLdouble *)
func VertexAttrib4f() { C.glVertexAttrib4f (GLuint, GLfloat, GLfloat, GLfloat, GLfloat)
func VertexAttrib4fv() { C.glVertexAttrib4fv (GLuint, const GLfloat *)
func VertexAttrib4iv() { C.glVertexAttrib4iv (GLuint, const GLint *)
func VertexAttrib4s() { C.glVertexAttrib4s (GLuint, GLshort, GLshort, GLshort, GLshort)
func VertexAttrib4sv() { C.glVertexAttrib4sv (GLuint, const GLshort *)
func VertexAttrib4ubv() { C.glVertexAttrib4ubv (GLuint, const GLubyte *)
func VertexAttrib4uiv() { C.glVertexAttrib4uiv (GLuint, const GLuint *)
func VertexAttrib4usv() { C.glVertexAttrib4usv (GLuint, const GLushort *)
func VertexAttribPointer() { C.glVertexAttribPointer (GLuint, GLint, GLenum, GLboolean, GLsizei, const GLvoid *)
GLAPI void APIENTRY glDisableVertexAttribArray (GLuint);
GLAPI void APIENTRY glEnableVertexAttribArray (GLuint);

*/

/*
func VertexAttribPointer(index uint, size, kind int, 
     normalized bool, stride int, pointer Pointer) {   C.glVertexAttribPointer(C.GLuint(index), C.GLint(size), 
      C.GLenum(kind), C.GLboolean(normalized), C.GLsizei(stride), pointer)
}


*/


func Drawarrays(mode, first, count int) { 
  C.glDrawArrays(C.GLenum(mode), C.GLint(first), C.GLsizei(count))
}

func DrawElements(mode, count, kind int, indices Pointer) {
  C.glDrawElements(C.GLenum(mode), C.GLsizei(count),
                   C.GLenum(kind), unsafe.Pointer(indices))
}

/*
GLAPI void GLAPIENTRY glDrawRangeElements( GLenum mode, GLuint start,
  GLuint end, GLsizei count, GLenum type, const GLvoid *indices );

GLAPI void GLAPIENTRY glTexImage3D( GLenum target, GLint level,
                                      GLint internalFormat,
                                      GLsizei width, GLsizei height,
                                      GLsizei depth, GLint border,
                                      GLenum format, GLenum type,
                                      const GLvoid *pixels );

GLAPI void GLAPIENTRY glTexSubImage3D( GLenum target, GLint level,
                                         GLint xoffset, GLint yoffset,
                                         GLint zoffset, GLsizei width,
                                         GLsizei height, GLsizei depth,
                                         GLenum format,
                                         GLenum type, const GLvoid *pixels);

GLAPI void GLAPIENTRY glCopyTexSubImage3D( GLenum target, GLint level,
                                             GLint xoffset, GLint yoffset,
                                             GLint zoffset, GLint x,
                                             GLint y, GLsizei width,
                                             GLsizei height );

*/





















