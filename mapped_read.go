package bioutil

import(
	"bufio"
)

type MappedRead interface {
	Data() []byte
	Mutations() Mutations
	AlignmentScore() float64
}

type MappedReadReader func(input *bufio.Reader, out chan MappedRead)
