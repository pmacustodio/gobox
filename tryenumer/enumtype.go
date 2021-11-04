package tryenumer

//go:generate enumer -type EnumType -json
type EnumType uint
type EnumType2 uint

const (
	Red EnumType = iota + 1
	Green
	Blue
)
