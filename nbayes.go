package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"bytes"

	// "text/scanner"
	"strconv"
)

/* ================> Design of the IncrementTable:
		All of the different options for the input/output
		variables.
   =============================================> */
type IncrementTable struct {
	XTYT, XTYF, XFYT, XFYF int
}

func ReturnFileDetails(pointerFileContents []string) (int64, int64) {
	numVariables, errVariables := strconv.ParseInt(pointerFileContents[0], 0, 32)
	fmt.Printf("numVariables is %v, errVariables is %v \n", numVariables, errVariables ) 
	numRows, errRows := strconv.ParseInt(pointerFileContents[1], 0, 32)
	fmt.Printf("numRows is %v, errRows is %v \n", numRows, errRows ) 
	return numVariables, numRows
}

func ReadInTrainData( trainString string ) {

	// ========> NumVariables & NumRows  
	linesFromTrain, err := readLines( trainString )
	fmt.Printf("linesFromTrain is %v, error is %v \n", linesFromTrain, err )
	numVariables, numRows := ReturnFileDetails( linesFromTrain )
	fmt.Printf("numVariables is %v, numRows is %v \n", numVariables, numRows )

	// =======> Read in the values into IncrementTables array
	trainData := make([]IncrementTable, numVariables)
	fmt.Printf("trainData is %v \n", trainData )

	for i := 0; i < numVariables; i ++  {
		var s scanner.Scanner
		s.Init(linesFromTrain[i])
		tok := s.Scan()
		// =============> Somehow get the 'Y' value
		yValue := 1;
		for tok != scanner.EOF {
			xValue, xErr :=  strconv.ParseInt(tok, 0, 32)
			IncrementInputVariable( &trainData, i, xValue, yValue )
			tok = s.Scan()
		}
	}
}

func IncrementInputVariable( trainData *[]IncrementTable, index, xVal, yVal int ) {
	if ( xVal == 1 && yVal == 1 ) {
		trainData[index].XTYT ++
	} else if ( xVal == 1 && yVal == 0 ) {
		trainData[index].XTYF ++
	} else if ( xVal == 0 && yVal == 1 ){
		trainData[index].XFYT ++
	} else ( xVal == 0 && yVal == 0 ){
		trainData[index].XFYF ++
	}
}

func ReadInTestData( testString string ) {
	linesFromTest, err := readLines( testString )
	fmt.Printf("linesFromTest is %v, err is %v \n", linesFromTest, err )

}

func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

/*========> Design of the Input 

	Example:
		"go run nbayes.go <train-file> <test-file>"

===============================>*/
func main() {

	// ReadBytes from bufio should play a big part
	ReadInTrainData( os.Args[1] )
	ReadInTestData( os.Args[2] )
	// TrainData( )
	// TestData( )
	fmt.Println("Just printing something.")
}
