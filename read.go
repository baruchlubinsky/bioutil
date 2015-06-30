package bioutil

import (
	"bufio"
)

type Read interface{
	Sequence() []Nucleotide
	Quality() []byte
	Header() []byte
	Length() int
	Base(i int) byte
	QualityScore(i int) byte
	AppendHeader(word string) Read
	Data() []byte
	TrimLeft(n int) Read
	TrimRight(n int) Read
}

type ReadReader func(input *bufio.Reader, out chan Read)
