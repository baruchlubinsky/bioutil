package bioutil

import (
	"bufio"
	"fmt"
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
	match := AlignmentScoreRe.FindStringSubmatch(string(a.HeadLine))
	if len(match) > 0 {
		s := match[1]
		score, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return 0.0
		}
		return score
	} else {
		return -1.0
	}

}

func (a *Alignment) String() string {
	return fmt.Sprintf("%s\n%s\n", a.HeadLine, a.SeqLine)
}

func (a *Alignment) Data() []byte {
	return append(a.HeadLine, a.SeqLine...)
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

func ScanFastmFileChan(inputPath string, out chan Alignment) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)
	go func() {
		for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
			seqLine, _ := input.ReadBytes('\n')
			out <- Alignment{headLine, seqLine}
		}
		close(out)
		inputFile.Close()
	}()
	return nil
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
