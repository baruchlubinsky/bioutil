package bioutil

import (
	"bufio"
	"io"
	"os"
)

type Alignment struct {
	HeadLine []byte
	SeqLine  []byte
}

type AlignmentFunc func(alignment Alignment) (interface{}, error)

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
		result, err := function(Alignment{headLine, seqLine})
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
