package edlib

type Source interface {
	String(i int) string
	Len() int
}
type Match struct {
	// The matched string.
	Str string
	// The index of the matched string in the supplied slice.
	Index int
	// Score used to rank matches
	Score float32
}
