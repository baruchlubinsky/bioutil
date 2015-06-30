package dnaio

import (
	"bufio"
	// "fmt"
	. "github.com/onsi/gomega"
	"io"
	"testing"
	"code.hyraxbio.co.za/bioutil"
)

func b(in string) []byte {
	return []byte(in)
}

func testRead() FastqRead {
	return FastqRead{b("Header"),
		b("ACTGACTGACTG"),
		b("+"),
		b("IIJJKKLLMMNN"),
	}
}

func testData() []byte {
	return b("Header\nACTGACTGACTG\n+\nIIJJKKLLMMNN\n")
}

func TestInputReading(t *testing.T) {
	RegisterTestingT(t)
	reader, writer := io.Pipe()
	input := bufio.NewReader(reader)
	c := make(chan bioutil.Read)
	go ScanFastqChan(input, c)
	writer.Write(testData())
	read := <-c
	seq := testRead().Sequence()
	Ω(read.Sequence()).Should(Equal(seq))
}

func TestDataWriting(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	data := testData()
	Ω(read.Data()).Should(Equal(data))
}

func TestTrimLeft(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	read = read.TrimLeft(2).(FastqRead)
	Ω(read.sequence).Should(Equal(b("TGACTGACTG")))
	Ω(read.quality).Should(Equal(b("JJKKLLMMNN")))
}

func TestTrimRight(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	read = read.TrimRight(3).(FastqRead)
	Ω(read.sequence).Should(Equal(b("ACTGACTGA")))
	Ω(read.quality).Should(Equal(b("IIJJKKLLM")))
}

func TestAppendHeader(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	read = read.AppendHeader("Test").(FastqRead)
	Ω(read.header).Should(Equal(b("Header\tTest")))
}

func TestMutability(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	seq := read.Sequence()
	seq[0] = 'X'
	Ω(seq).ShouldNot(Equal(read.Sequence()))
}

func TestTrimAndWrite(t *testing.T) {
	RegisterTestingT(t)
	read := testRead()
	read = read.TrimLeft(2).(FastqRead)
	Ω(read.Data()).Should(Equal(b("Header\nTGACTGACTG\n+\nJJKKLLMMNN\n")))
}
