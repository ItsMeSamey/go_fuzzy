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

