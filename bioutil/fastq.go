package bioutil

import(
	"os"
	"bufio"
	"io"
)

type Read struct {
	HeadLine []byte
	SeqLine []byte
	SepLine []byte
	QualLine []byte
}

type ReadFunc func(Read) (interface{}, error)

func ScanFastqFile(inputPath string, f ReadFunc) ([]interface{}, error) {
	inputFile, err := os.Open(inputPath)
	input := bufio.NewReader(inputFile)
	result := make([]interface{}, 0)
	for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, _ = input.ReadBytes('\n') {
		seqLine, _ := input.ReadBytes('\n')
		sepLine, _ := input.ReadBytes('\n')
		qualLine, _ := input.ReadBytes('\n')
		read := Read{headLine, seqLine, sepLine, qualLine}
		r, _ := f(read)
		result = append(result, r)
	}
	inputFile.Close()
	return result, err
}