package bioutil

import (
	"bufio"
	"io"
	"os"
)

// A read from the sequencer. The data is immutable and may only be trimmed using the provided functions.
type Read struct {
	header   []byte
	sequence []byte
	sep      []byte
	quality  []byte
}

type SequenceFunc func(read *Read) (interface{}, error)

type SequenceFilterFunc func(read *Read) bool

// Returns a copy of the sequence data, use Base() where possible
func (a *Read) Sequence() []byte {
	res := make([]byte, len(a.sequence))
	copy(res, a.sequence)
	return res
}

// Returns a copy of the quality data, use QualityScore() where possible
func (a *Read) Quality() []byte {
	res := make([]byte, len(a.quality))
	copy(res, a.quality)
	return res
}

func (a *Read) Header() []byte {
	res := make([]byte, len(a.header))
	copy(res, a.header)
	return res
}

// Returns the length of the sequence
func (a *Read) Length() int {
	return len(a.sequence)
}

func (a *Read) Base(i int) byte {
	return a.sequence[i]
}

func (a *Read) QualityScore(i int) byte {
	return a.quality[i]
}

func (a *Read) AppendHeader(word string) {
	a.header = append(a.header, []byte("\t"+word)...)
}

// Get the full read to write into a .fastq file
func (a *Read) Data() []byte {
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
func (r *Read) TrimLeft(n int) {
	r.sequence = r.sequence[n:]
	r.quality = r.quality[n:]
}

// Remove the last n bases from the Read
func (r *Read) TrimRight(n int) {
	r.sequence = r.sequence[0 : len(r.sequence)-n]
	r.quality = r.quality[0 : len(r.quality)-n]
}

// Read the contents of the provided .fastq file asynchronously, returns a channel
// that the data may be read from that is closed at EOF.
func ScanFastqFile(inputPath string) (chan Read, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	out := make(chan Read)
	input := bufio.NewReader(inputFile)
	go func() {
		ScanFastqChan(input, out)
		inputFile.Close()
		close(out)
	}()
	return out, nil
}

// Read the buffer onto the channel as .fastq reads.
func ScanFastqChan(input *bufio.Reader, out chan Read) {
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		// do not store new lines
		out <- Read{headLine[:len(headLine)-1], seqLine[:len(seqLine)-1], sepLine[:len(sepLine)-1], qualLine[:len(qualLine)-1]}
	}
}

// Read a .fastq file onto an existing channel. Optionally close the channel when complete.
func ScanFastqFileChan(inputPath string, out chan Read, c bool) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)

	ScanFastqChan(input, out)
	if c {
		inputFile.Close()
	}

	return nil
}

// func ScanFastqFilterFile(inputPath string, filter SequenceFilterFunc, function SequenceFunc) ([]interface{}, error) {
// 	inputFile, err := os.Open(inputPath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	input := bufio.NewReader(inputFile)
// 	defer func() {
// 		inputFile.Close()
// 	}()
// 	return ScanFastqFilter(input, filter, function)
// }

// func ScanFastqFilter(input *bufio.Reader, filter SequenceFilterFunc, function SequenceFunc) ([]interface{}, error) {
// 	results := make([]interface{}, 0)
// 	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
// 		seqLine, _ := input.ReadBytes('\n')
// 		sepLine, _ := input.ReadBytes('\n')
// 		qualLine, _ := input.ReadBytes('\n')
// 		read := &Read{headLine, seqLine, sepLine, qualLine}
// 		if filter(read) {
// 			result, err := function(read)
// 			if err != nil {
// 				return nil, err
// 			}
// 			results = append(results, result)
// 		}
// 	}
// 	return results, nil
// }
