package gut

import "testing"
import "fmt"
import "bytes"
// import "os"

// Tests for fopen
func TestFopen(t * testing.T) {
  file, err := Fopen("test1.tmp","w")
  if err != nil { t.Fatal("Could not open temp file for writing."); }
  file.Close();
  file, err  = Fopen("test1.tmp","r")
  if err != nil { t.Error("Could not open temp file for reading."); }
  file.Close();
}

// Tests for FreadAll
func TestFreadAll(t * testing.T) {
  cont, err2 := FreadAll("gut_test.go")  
  if err2 != nil { t.Fatal("Could not open and read file."); }
  if len(cont) == 0 { t.Error("File not read in corectly."); }
}

func TestHomeDir(t * testing.T) {
  dir := HomeDir()
  t.Log(dir)
  fmt.Println(dir)
}

func TestJoinDir(t * testing.T) {
  dir := JoinDir("foo/", "bar")
  if dir != "foo/bar" { t.Errorf("Dir names not joined correctly: %s", dir) }
}

type Record struct {
  Ui32   uint32
  Ui16   uint16
  Buf	 []byte
  Str	 string
}  
  

func TestPack(t * testing.T) {
  rec := &Record{12345, 0, nil, "hi!"}
  bb := bytes.NewBufferString("")  
  PackBE(bb, rec.Ui32)
  b  := bb.Bytes()
  fmt.Printf("%x %x %x %x\n", b[0], b[1], b[2], b[3])
  b1 := bytes.NewBuffer(b)
  err := UnpackBE(b1, &rec.Ui16)
  fmt.Printf("%s %d\n", err, rec.Ui16)
  b3 := bytes.NewBufferString("")  
  rec2 := &Record{123, 16, nil, "hi!"}
  rec2.Buf = make([]byte, rec2.Ui16)
  rec2.Buf[0] = 12
  rec2.Buf[1] = 34
  err = Pack(b3, rec2)
  fmt.Println(err, b3.Bytes())
  
  
}
