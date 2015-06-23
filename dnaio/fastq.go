package dnaio

import (
	"code.hyraxbio.co.za/bioutil"
	"bufio"
	"io"
	"os"
)

// A read from the sequencer. 
type FastqRead struct {
	header   []byte
	sequence []byte
	sep      []byte
	quality  []byte
}

// Returns a copy of the sequence data, use Base() where possible
func (a FastqRead) Sequence() []bioutil.Nucleotide {
	res := make([]bioutil.Nucleotide, len(a.sequence))
	for i, b := range a.sequence {
		res[i] = bioutil.Nucleotide(b)
	}
	return res
}

// Returns a copy of the quality data, use QualityScore() where possible
func (a FastqRead) Quality() []byte {
	res := make([]byte, len(a.quality))
	copy(res, a.quality)
	return res
}

func (a FastqRead) Header() []byte {
	res := make([]byte, len(a.header))
	copy(res, a.header)
	return res
}

// Returns the length of the sequence
func (a FastqRead) Length() int {
	return len(a.sequence)
}

func (a FastqRead) Base(i int) byte {
	return a.sequence[i]
}

func (a FastqRead) QualityScore(i int) byte {
	return a.quality[i]
}

func (a FastqRead) AppendHeader(word string) bioutil.Read {
	a.header = append(a.header, []byte("\t"+word)...)
	return a
}

// Get the full read to write into a .fastq file
func (a FastqRead) Data() []byte {
	var temp []byte
	temp = append(a.header, '\n')
	temp = append(temp, a.sequence...)
	temp = append(temp, '\n')
	temp = append(temp, a.sep...)
	temp = append(temp, '\n')
	temp = append(temp, a.quality...)
	temp = append(temp, '\n')
	return temp
}

// Remove the first n bases from the Read
func (r FastqRead) TrimLeft(n int) bioutil.Read {
	r.sequence = r.sequence[n:]
	r.quality = r.quality[n:]
	return r
}

// Remove the last n bases from the Read
func (r FastqRead) TrimRight(n int) bioutil.Read {
	r.sequence = r.sequence[0 : len(r.sequence)-n]
	r.quality = r.quality[0 : len(r.quality)-n]
	return r
}

// Read the buffer onto the channel as .fastq reads.
func ScanFastqChan(input *bufio.Reader, out chan bioutil.Read) {
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		// do not store new lines
		out <- FastqRead{headLine[:len(headLine)-1], seqLine[:len(seqLine)-1], sepLine[:len(sepLine)-1], qualLine[:len(qualLine)-1]}
	}
}

// Read the contents of the provided .fastq file asynchronously, returns a channel
// that the data may be read from that is closed at EOF.
func ScanFastqFile(inputPath string) (chan bioutil.Read, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	out := make(chan bioutil.Read)
	input := bufio.NewReader(inputFile)
	go func() {
		ScanFastqChan(input, out)
		inputFile.Close()
		close(out)
	}()
	return out, nil
}

// Read a .fastq file onto an existing channel. Optionally close the channel when complete.
func ScanFastqFileChan(inputPath string, out chan bioutil.Read, c bool) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)
	ScanFastqChan(input, out)
	inputFile.Close()
	if c {
		close(out)
	}

	return nil
}

