package bioutil

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
)

type Alignment struct {
	HeadLine []byte
	SeqLine  []byte
}

var AlignmentScoreRe *regexp.Regexp

func init() {
	AlignmentScoreRe = regexp.MustCompile("AlignmentScore: ([0-9.]+)\t")
}

func (a *Alignment) Score() float64 {
	s := AlignmentScoreRe.FindStringSubmatch(string(a.HeadLine))[1]
	score, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return score
}

type AlignmentFunc func(alignment *Alignment) (interface{}, error)

func ScanFastmFile(inputPath string, function AlignmentFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	input := bufio.NewReader(inputFile)
	defer func() {
		inputFile.Close()
	}()
	info, _ := inputFile.Stat()
	results := make([]interface{}, 0, info.Size()/400)
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		result, err := function(&Alignment{headLine, seqLine})
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

type AlignmentIndexFunc func(alignment *Alignment, index int) (interface{}, error)

func ScanFastmFileIndex(inputPath string, function AlignmentIndexFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	input := bufio.NewReader(inputFile)
	defer func() {
		inputFile.Close()
	}()
	info, _ := inputFile.Stat()
	results := make([]interface{}, 0, info.Size()/400)
	index := 0
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		result, err := function(&Alignment{headLine, seqLine}, index)
		index++
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
