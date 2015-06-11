package bioutil

type Read interface{
	Sequence() []Nucleotide
	Quality() []byte
	Header() []byte
	Length() int
	Base(i int) byte
	QualityScore(i int) byte
	AppendHeader(word string)
	Data() []byte
	TrimLeft(n int)
	TrimRight(n int)
}