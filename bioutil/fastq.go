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

type SequenceFunc func(read Read) (interface{}, error)

func ScanFastqFile(inputPath string, function SequenceFunc) ([]interface{}, error) {
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
    sepLine, _ := input.ReadBytes('\n')
    qualLine, _ := input.ReadBytes('\n')
    result, err := function(Read{headLine, seqLine, sepLine, qualLine})
    if err != nil {
      return nil, err
    }
    results = append(results, result)
  }
  return results, nil
}
