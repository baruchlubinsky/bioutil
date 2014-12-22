package bioutil

import (
	"bufio"
	"io"
	"os"
)

func ScanPairsFileChan(inputPath string, out chan Alignment) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	input := bufio.NewReader(inputFile)
	go func() {
		// Skip 3 lines, refernce sequence
		input.ReadBytes('\n')
		input.ReadBytes('\n')
		input.ReadBytes('\n')
		for headLine, err := input.ReadBytes('\n'); err != io.EOF; headLine, err = input.ReadBytes('\n') {
			seqLine, _ := input.ReadBytes('\n')
			out <- Alignment{headLine, seqLine}
			// Skip 4 lines, refernce sequence
			input.ReadBytes('\n')
			input.ReadBytes('\n')
			input.ReadBytes('\n')
			input.ReadBytes('\n')
		}
		close(out)
		inputFile.Close()
	}()
	return nil
}
