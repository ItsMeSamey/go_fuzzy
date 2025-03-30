package common

type Number interface {
  int8 | int16 | int32 | int64 | float32 | float64 | int
}
func Abs[T Number](num T) T {
  if num < 0 { return -num }
  return num
}

type FloatType interface {
  float32 | float64
}

type StringLike interface {
  string | []byte
}

type ByteWriter struct {
  Buf []byte
  At  int
}
func (w *ByteWriter) Write(p []byte) (n int, err error) {
  copy(w.Buf[w.At:], p)
  if len(w.Buf) < w.At + len(p) {
    w.Buf = append(w.Buf, p[len(w.Buf)-w.At:]...)
  }
  w.At += len(p)
  return len(p), nil
}

