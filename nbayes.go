package main

import (
	"os"
	"io"
	"fmt"
	"bufio"
	"bytes"
	"strings"

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

func ReadInTestData( testString string ) [][]int {
	linesFromTest, err := readLines( testString )
	fmt.Printf("Printing the err : %v \n", err )
	numVariables, numRows := ReturnFileDetails( linesFromTest )
	testData := make( [][]int, int(numRows))

	for i := 2; i < int(numRows) + 2; i ++  {
		currentRow := linesFromTest[i]
		testData[i - 2] = make( []int,  int(numVariables) + 1)
		tokenizedString := strings.FieldsFunc( currentRow, delimiter )
		for j := 0; j <= int(numVariables); j ++ {
			indexValue, tErr := strconv.ParseInt( tokenizedString[j], 0, 32 )
			fmt.Printf("Printing the err : %v \n", tErr )
			testData[i - 2][j] =  int(indexValue)
		}
	}
	return testData
}

func ReadInTrainData( trainString string ) ([]IncrementTable) {

	// =======> NumVariables & NumRows  
	linesFromTrain, err := readLines( trainString )
	fmt.Printf("linesFromTrain is %v, error is %v \n", linesFromTrain, err )
	numVariables, numRows := ReturnFileDetails( linesFromTrain )
	fmt.Printf("numVariables is %v, numRows is %v \n", numVariables, numRows )

	// =======> Read in the values into IncrementTables array
	trainData := make([]IncrementTable, numVariables)
	fmt.Printf("trainData is %v \n", trainData )

	for i := 2; i < int(numRows) + 2; i ++  {
		currentRow := linesFromTrain[i]
		tokenizedString := strings.FieldsFunc( currentRow, delimiter )
		yValue, yErr := strconv.ParseInt(tokenizedString[numVariables], 0, 32)
		fmt.Printf("yValue is %v, yErr is %v \n", yValue, yErr )
		for j := 0; j < int(numVariables); j ++ {
			xValue, xErr := strconv.ParseInt(tokenizedString[j], 0, 32)
			fmt.Printf("xValue is %v, xErr is %v \n", xValue, xErr )
			IncrementInputVariable( &trainData, j,  xValue, yValue )
		}
	}
	return trainData
}

func delimiter( r rune ) bool {
	return r == ':' || r == ' '
}

func IncrementInputVariable( trainData *[]IncrementTable, index int, xVal, yVal int64 ) {
	if xVal == 1 && yVal == 1 {
		(*trainData)[index].XTYT ++
	} else if xVal == 1 && yVal == 0 {
		(*trainData)[index].XTYF ++
	} else if xVal == 0 && yVal == 1 {
		(*trainData)[index].XFYT ++
	} else {
		(*trainData)[index].XFYF ++
	}
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
		// ====> Waits till the surrounding function returns
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

func OutputResults( numberCorrect, numberSolved int ) {
	fmt.Printf("The number correct is %v \n", numberCorrect )
	fmt.Printf("The number solved is %v \n", numberSolved )

	var accuracy float64 = float64(numberSolved) / numberCorrect
	fmt.Printf("The accuracy is %v \n", accuracy )
}

func CalculateProb( table []IncrementTable, testElement, index, yIs0 ) {
	if testElement == 0 && yIs0 == true {
		return float64(table[j].XFYT) / (table[j].XFYT + table[j].XTYT)
	} else if testElement == 0 && yIs0 == false  {
		return float64(table[j].XFYF) / (table[j].XFYF + table[j].XTYF)
	} else if testElement == 1 && yIs0 == true {
		return float64(table[j].XTYT) / (table[j].XFYT + table[j].XTYT)
	} else {
		return float64(table[j].XTYF) / (table[j].XFYF + table[j].XTYF)
	}
}

func ProbabilityOfY( table []IncrementTable, yIsO ) float64 {
	var totalNumber, respNumber int = 0
	if yIs0 {
		for i := 0; i < len(table); i ++ {
			totalNumber += table[i].XFYF
			totalNumber += table[i].XTYF
			totalNumber += table[i].XFYT
			totalNumber += table[i].XTYT

			respNumber += table[i].XFYF
			respNumber += table[i].XTYF
		}
	} else {
		for i := 0; i < len(table); i ++ {
			totalNumber += table[i].XFYF
			totalNumber += table[i].XTYF
			totalNumber += table[i].XFYT
			totalNumber += table[i].XTYT

			respNumber += table[i].XFYT
			respNumber += table[i].XTYT
		}
	}
	return float64(respNumber) / totalNumber
}

func AnalyzeData( trainData []IncrementTable, testData [][]int ) {
	var numberTested, numberCorrect int = 0
	for i := 0; i < len(testData); i ++ {
		var y0Value, y1Value float64 = 1
		for j := 0; j < len(testData[i]) - 1; j ++ {
			y0Value *= CalculateProb( IncrementTable, testData[i][j], j, true )
			y1Value *= CalculateProb( IncrementTable, testData[i][j], j, false )
		}
		y0Value *=  ProbabilityOfY( IncrementTable, true )
		y1Value *=  ProbabilityOfY( IncrementTable, false )
		var finalyValue float64
		if y0Value > y1Value {
			finalyValue = 0
		} else {
			finalyValue = 1
		}
		if finalyValue == testData[i][len(testData) - 1] {
			numberCorrect ++
		}
		numberTested ++
	}
	OutputResults( numberCorrect, numberTested )
}

/*========> Design of the Input 

	Example:
		"go run nbayes.go <train-file> <test-file>"

===============================>*/
func main() {
	trainData := ReadInTrainData( os.Args[1] )
	fmt.Printf("trainData returned is %#v \n", trainData )
	testData := ReadInTestData( os.Args[2] )
	fmt.Printf("testData returned is %#v \n", testData )

	AnalyzeData( trainData, testData )
}
