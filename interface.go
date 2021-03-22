package edlib

type Source interface {
	String(i int) string
	Len() int
}
