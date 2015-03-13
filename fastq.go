package bioutil

import (
	"bufio"
	"io"
	"os"
)

type Read struct {
	HeadLine []byte
	SeqLine  []byte
	SepLine  []byte
	QualLine []byte
}

type SequenceFunc func(read *Read) (interface{}, error)

type SequenceFilterFunc func(read *Read) bool

func (a *Read) Data() []byte {
	temp := append(a.HeadLine, a.SeqLine...)
	temp = append(temp, a.SepLine...)
	return append(temp, a.QualLine...)
}

func ScanFastqFile(inputPath string, function SequenceFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	input := bufio.NewReader(inputFile)
	defer func() {
		inputFile.Close()
	}()
	return ScanFastq(input, function)
}

func ScanFastq(input *bufio.Reader, function SequenceFunc) ([]interface{}, error) {
	results := make([]interface{}, 0)
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		result, err := function(&Read{headLine, seqLine, sepLine, qualLine})
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func ScanFastqChan(input *bufio.Reader, out chan Read) {
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		out <- Read{headLine, seqLine, sepLine, qualLine}
	}
	close(out)
}

func ScanFastqFileChan(inputPath string, out chan Read) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)
	go func() {
		for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
			seqLine, _ := input.ReadBytes('\n')
			sepLine, _ := input.ReadBytes('\n')
			qualLine, _ := input.ReadBytes('\n')
			out <- Read{headLine, seqLine, sepLine, qualLine}
		}
		close(out)
		inputFile.Close()
	}()
	return nil
}

func ScanFastqFilterFile(inputPath string, filter SequenceFilterFunc, function SequenceFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	input := bufio.NewReader(inputFile)
	defer func() {
		inputFile.Close()
	}()
	return ScanFastqFilter(input, filter, function)
}

func ScanFastqFilter(input *bufio.Reader, filter SequenceFilterFunc, function SequenceFunc) ([]interface{}, error) {
	results := make([]interface{}, 0)
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		read := &Read{headLine, seqLine, sepLine, qualLine}
		if filter(read) {
			result, err := function(read)
			if err != nil {
				return nil, err
			}
			results = append(results, result)
		}
	}
	return results, nil
}
