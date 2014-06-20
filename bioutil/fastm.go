package bioutil

import(
	"os"
	"bufio"
	"io"
)

type Alignment struct {
	HeadLine []byte
	SeqLine []byte
}

type AlignmentFunc func(Alignment) (interface{}, error)

func ScanFastmFile(inputPath string, f AlignmentFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	input := bufio.NewReader(inputFile)
	result := make([]interface{}, 0)
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, _ = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		a := Alignment{headLine, seqLine}
		r, _ := f(a)
		result = append(result, r)
	}
	inputFile.Close()
	return result, err
}