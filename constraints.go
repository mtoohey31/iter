package iter

// TODO: replace with official constraints package's version once I can get it working

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type number interface {
	integer | ~float32 | ~float64
}

type ordered interface {
	number | ~string
}
