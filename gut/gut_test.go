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

func TestPack(t * testing.T) {
  var i32 int32 = 12345678
  bb := bytes.NewBufferString("")  
  Pack(bb, i32)
  b  := bb.Bytes()
  fmt.Printf("%x %x %x %x\n", b[0], b[1], b[2], b[3])
}
