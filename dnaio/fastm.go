package dnaio

import (
	"code.hyraxbio.co.za/bioutil"
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"bytes"
)

type FastmAlignment struct {
	headLine []byte
	seqLine  []byte
	factory *bioutil.MutationFactory
}

var AlignmentScoreRe *regexp.Regexp

func init() {
	AlignmentScoreRe = regexp.MustCompile("AlignmentScore: ([0-9.]+)\t")
}

func (a FastmAlignment) AlignmentScore() float64 {
	match := AlignmentScoreRe.FindStringSubmatch(string(a.headLine))
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

func (a FastmAlignment) Mutations() bioutil.Mutations {
	mutations := bytes.Split([]byte(","), a.seqLine)
	if a.factory == nil {
		a.factory = new(bioutil.MutationFactory)
	}
	res := make(bioutil.Mutations, len(mutations))
	for i, m := range mutations {
		res[i] = a.factory.Parse(string(m))
	}
	return res
}

func (a FastmAlignment) MutationsInRegion(region *bioutil.GeneRegion) bioutil.Mutations {
	mutations := bytes.Split([]byte(","), a.seqLine)
	if a.factory == nil {
		a.factory = new(bioutil.MutationFactory)
	}
	res := make(bioutil.Mutations, len(mutations))
	for i, m := range mutations {
		res[i] = a.factory.ParseInRegion(string(m), region)
	}
	return res
}

func (a FastmAlignment) String() string {
	return fmt.Sprintf("%s\n%s\n", a.headLine, a.seqLine)
}

func (a FastmAlignment) Data() []byte {
	return append(a.headLine, a.seqLine...)
}

func (a FastmAlignment) ReferenceSequence() bioutil.AminoAcids {
	if a.factory == nil {
		return bioutil.AminoAcids{}
	}
	return a.factory.Sequence
}

func ScanFastmChan(factory *bioutil.MutationFactory, input *bufio.Reader, out chan bioutil.MappedRead) {
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		out <- FastmAlignment{headLine, seqLine, factory}
	}
}


func ScanFastmFile(factory *bioutil.MutationFactory, inputPath string) (chan bioutil.MappedRead, error) {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	out := make(chan bioutil.MappedRead)
	input := bufio.NewReader(inputFile)
	go func() {
		ScanFastmChan(factory, input, out)
		close(out)
		inputFile.Close()
	}()
	return out, nil
}

func ScanFastmFileChan(factory *bioutil.MutationFactory, inputPath string, out chan bioutil.MappedRead, c bool) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)
	ScanFastmChan(factory, input, out)
	inputFile.Close()
	if c {
		close(out)
	}

	return nil
}
