package enumerate

type Enumerable[R any] interface {
	GetLen() int
	Next() R
	HasNext() bool
}
