package stdlib

import (
	"bufio"
	"bytes"
	"fmt"

	"os"
	//"strings"
)

/////////| Testing a scanner on a custom text |/////////////

const scannertext = "this is \nan experimental \ntext"
const scannertext2 = `
this
is 
also an
experiment
`

func Experiment_scanner() {
	temp, _ := os.CreateTemp(".", "scannerfile")
	temp.Write([]byte(scannertext))
	temp.Seek(0, 0)
	scanner := bufio.NewScanner(temp)
	i := 1 //I'm using the counter here to check whether the lines are properly divided at newline, or whether the scanner counts it all as one line that is merely printed across multiple lines in the terminal.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(i, line)
		i++
	}
	r := bytes.NewBufferString(scannertext2)
	scanner2 := bufio.NewScanner(r)
	j := 1
	for scanner2.Scan() {
		line := scanner2.Text()
		fmt.Println(j, line)
		j++
	}
	temp.Close()
	os.Remove(temp.Name())
}

//////////| Testinc a scanner with a cursom token divider function |////////////////

const scannertext3 = "this|is a|custom divider|experiment| "

func Experiment_customSplit_scanner() {
	temp, _ := os.CreateTemp(".", "scannerfile")
	temp.Write([]byte(scannertext3))
	temp.Seek(0, 0)
	scanner := bufio.NewScanner(temp)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) { // I don't think this split function is particularly well done or efficient, but it seems to work.
		if atEOF == true {
			if len(data) > 0 {
				return len(data), data, nil
			} else if len(data) == 0 {
				return 0, nil, nil
			}
		}
		divisions := bytes.Split(data, []byte("|"))
		i := 0
		for len(divisions[i]) == 0 {
			i++
		}
		if len(divisions[i]) == 0 {

		}
		return len(divisions[i]) + 1, divisions[i], nil

	})

	//type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(i, line)
		i++
	}

	temp.Close()
	os.Remove(temp.Name())
}

// The stdlib line split function:
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) { //Alternate: ScanWords, which splits at each space. ScanBytes returns each byte indivitually
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\n'); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func dropCR(data []byte) []byte { //Removes \r from lines that contain it, such as lines of an http message
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
