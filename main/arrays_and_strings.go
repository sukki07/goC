package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printSlice(input []string) {
	for index, val := range input {
		if index != len(input)-1 {
			fmt.Print(val)
			fmt.Print(" ")
		} else {
			fmt.Print(val)
		}
	}
}

func timeVersion() {
	slices := ReadStringWrapper()
	testcases := slices[0]
	numTestCases, _ := strconv.Atoi(testcases)
	for i := 1; i <= numTestCases; i++ {
		metaDataTestCaseIndex := 2*i - 1
		dataTestCaseIndex := 2 * i
		metaData := strings.Split(slices[metaDataTestCaseIndex], " ")
		//lengthOfInput := metaData[0]
		roatationTimes := metaData[1]
		numRoatationTimes, _ := strconv.Atoi(roatationTimes)
		input := strings.Split(slices[dataTestCaseIndex], " ")
		for r := 1; r <= numRoatationTimes; r++ {
			rotateBy1(input)
		}
		printSlice(input)
		fmt.Println()
	}
}

func spaceVersion() {
	slices := ReadStringWrapper()
	testcases := slices[0]
	numTestCases, _ := strconv.Atoi(testcases)
	for i := 1; i <= numTestCases; i++ {
		metaDataTestCaseIndex := 2*i - 1
		dataTestCaseIndex := 2 * i
		metaData := strings.Split(slices[metaDataTestCaseIndex], " ")
		roatationTimes := metaData[1]
		numRoatationTimes, _ := strconv.Atoi(roatationTimes)
		input := strings.Split(slices[dataTestCaseIndex], " ")
		shiftAheadVal := numRoatationTimes % len(input)
		newSlice := rotateByShiftAhead(input, shiftAheadVal)

		printSlice(newSlice)
		fmt.Println()
	}
}

func findCycle(start int, input []string, shiftAheadVal int) int {
	l := len(input)
	previousValue := input[start]
	startCycleIndex := start
	count := 0
	for {
		futureIndex := (startCycleIndex + shiftAheadVal) % l
		temp := input[futureIndex]
		input[futureIndex] = previousValue
		count++
		//fmt.Printf("index %v will be copied to %v", startCycleIndex, futureIndex)
		//fmt.Println()
		if futureIndex == start {
			break
		}
		startCycleIndex = futureIndex
		previousValue = temp
	}
	return count
}

func rotateByShiftAhead(input []string, shiftAheadVal int) []string {
	l := len(input)
	start := 0
	//end := (start + shiftAheadVal) % l
	count := 0
	for { //parent loop to start cycles from 0 to index where 0 goes , after than no cycles to avoid rep
		count += findCycle(start, input, shiftAheadVal)
		start++
		if count >= l { //even fewer cycles can
			break
		}
	}
	return input
}

/*
for i := 0; i < len(input); i++ {
newIndex := (i + shiftAheadVal) % len(input)
newSlice[newIndex] = input[i]
}
return newSlice
*/

func MonkAndRoatation() {
	//timeVersion()  //time version
	spaceVersion() //space version
}

func rotateBy1(numSlice []string) {
	temp := numSlice[len(numSlice)-1]
	for i := len(numSlice) - 1; i >= 1; i-- {
		numSlice[i] = numSlice[i-1]
	}
	numSlice[0] = temp
}

func MonkAndGoodString() {
	inputString := ReadStringWrapper()[0]
	chars := strings.Split(inputString, "")
	index := 0
	maxGoodStringCount := 0
	for index < len(chars) {
		found, beg := findBeginnigOfGoodString(index, chars)
		if found {
			end := findEndOfGoodString(beg, chars)
			index = end + 1
			count := end - beg + 1
			if count > maxGoodStringCount {
				maxGoodStringCount = count
			}
			//fmt.Println(chars[beg:index]) //index goes out of bounds ,but :index does not include print value at index   only before
		} else {
			break //if not found means string has ended already
		}
	}
	fmt.Println(maxGoodStringCount)
}

func findEndOfGoodString(beg int, chars []string) int {
	i := beg
	for ; i < len(chars); i++ {
		if isVowel(chars[i]) {
			continue
		}
		break
	}
	return i - 1
}

func isVowel(char string) bool {
	switch char {
	case "a":
		return true
	case "e":
		return true
	case "i":
		return true
	case "o":
		return true
	case "u":
		return true
	}
	return false
}

func findBeginnigOfGoodString(location int, chars []string) (bool, int) {
	i := location
	for ; i < len(chars); i++ {
		if isVowel(chars[i]) {
			return true, i
		}
	}
	return false, 0
}

func printTestCase(sizeIndex int, slices []string) int {
	size, _ := strconv.Atoi(slices[sizeIndex])
	startIndex := sizeIndex + 1
	endIndex := sizeIndex + size
	processTestCase(startIndex, endIndex, slices)
	return endIndex + 1
}

func convert(slice []string) []int {
	temp := []int{}
	for _, char := range slice {
		val, _ := strconv.Atoi(char)
		temp = append(temp, val)
	}
	return temp
}

func traverse(mainSlice [][]int, x int, y int) int {
	l := len(mainSlice[0])
	val := mainSlice[x][y]
	count := 0
	for i := x; i < l; i++ {
		for j := y; j < l; j++ {
			if val > mainSlice[i][j] {
				count++
			}
		}
	}
	return count
}
func findInversionPairs(mainSlice [][]int) {
	length := len(mainSlice[0])
	count := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			count += traverse(mainSlice, i, j)
		}
	}
	fmt.Println(count)
}
func processTestCase(start int, end int, slices []string) {
	mainSlice := [][]int{}
	for i := start; i <= end; i++ {
		individualSlice := strings.Split(slices[i], " ")
		numSlice := convert(individualSlice)
		mainSlice = append(mainSlice, numSlice)
	}
	findInversionPairs(mainSlice)
}

func MonkAndInversions() {
	slices := ReadStringWrapper()
	testcases, _ := strconv.Atoi(slices[0])
	sizeIndex := 1
	for i := 1; i <= testcases; i++ {
		sizeIndex = printTestCase(sizeIndex, slices)
	}
}

func checkPrintPalindrome(input string) {
	slice := strings.Split(input, "")
	lenght := len(slice)
	i := 0
	j := lenght - 1
	for i <= j {
		if slice[i] == slice[j] {
		} else {
			fmt.Println("NO")
			return
		}
		i++
		j--
	}
	if i == j+1 {
		fmt.Println("YES EVEN")
	} else {
		fmt.Println("YES ODD")
	}
}

func MonkTeachesPalindrome() {
	slices := ReadStringWrapper()
	for i := 1; i < len(slices); i++ {
		checkPrintPalindrome(slices[i])
	}
}

func MonkAndWelcomeProblem() {
	input := ReadStringWrapper()
	lengthString := input[0]
	length, _ := strconv.Atoi(lengthString)
	firstArray := strings.Split(input[1], " ")
	secondArray := strings.Split(input[2], " ")
	thirdSlice := []int{}
	for i := 0; i < length; i++ {
		fi, _ := strconv.Atoi(firstArray[i])
		si, _ := strconv.Atoi(secondArray[i])
		thirdSlice = append(thirdSlice, fi+si)
	}
	for i := 0; i < length; i++ {
		fmt.Print(thirdSlice[i])
		fmt.Print(" ")
	}

}

func NQueen() {
	input := ReadStringWrapper()[0]
	number, _ := strconv.Atoi(input)

	Board := [][]bool{}
	for i := 0; i < number; i++ {
		temp := []bool{}
		for j := 0; j < number; j++ {
			temp = append(temp, false)
		}
		Board = append(Board, temp)
	}
	if Place(Board, number, number, 0) {
		fmt.Println("YES")
		for i := 0; i < len(Board); i++ {
			for j := 0; j < len(Board[i]); j++ {
				if Board[i][j] == true {
					fmt.Print(1)
					fmt.Print(" ")
				} else {
					fmt.Print(0)
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("NO")
	}
}

func Place(Board [][]bool, queensRemaining int, maxQueenNumber int, x int) bool {
	if queensRemaining == 0 {
		return true
	}
	for i := x; i < maxQueenNumber; i++ {
		for j := 0; j < maxQueenNumber; j++ {
			if canBePlacedAt(Board, i, j, maxQueenNumber) {
				Board[i][j] = true
				if Place(Board, queensRemaining-1, maxQueenNumber, i) {
					return true
				}
				Board[i][j] = false
			}
		}
	}
	return false
}

func canBePlacedAt(Board [][]bool, x int, y int, maxQueenNumber int) bool {
	//check horizontal
	for j := 0; j < maxQueenNumber; j++ {
		if Board[x][j] == true {
			return false
		}
	}

	for i := 0; i < maxQueenNumber; i++ {
		if Board[i][y] == true {
			return false
		}
	}

	for i := 0; i < maxQueenNumber; i++ {
		for j := 0; j < maxQueenNumber; j++ {
			if i+j == x+y {
				if Board[i][j] == true {
					return false
				}
			}
			if i-j == x-y {
				if Board[i][j] == true {
					return false
				}
			}
		}
	}
	return true
}

func CountOnesInBinaryRep() {
	stringSlice := ReadStringWrapper()
	for i := 1; i < len(stringSlice); i++ {
		number, _ := strconv.Atoi(stringSlice[i])
		count := 0
		for number > 0 {
			number = number & (number - 1)
			count++
		}
		fmt.Println(count)
	}
}

func CountDigits() {
	stringSlice := ReadStringWrapper()
	input := stringSlice[0]
	var countArray [10]int
	for _, runeValue := range input {
		switch runeValue {
		case '0':
			countArray[0]++
		case '1':
			countArray[1]++
		case '2':
			countArray[2]++
		case '3':
			countArray[3]++
		case '4':
			countArray[4]++
		case '5':
			countArray[5]++
		case '6':
			countArray[6]++
		case '7':
			countArray[7]++
		case '8':
			countArray[8]++
		case '9':
			countArray[9]++
		}
	}
	for index, value := range countArray {
		fmt.Printf("%d %d", index, value)
		fmt.Println()
	}
}

func InputOutput() error {
	stringSlice := ReadStringWrapper()
	intValue, _ := strconv.Atoi(stringSlice[0])
	fmt.Println(intValue * 2)
	fmt.Println(stringSlice[1])
	return nil
}

func ReadStringWrapper() []string {
	//bufStdIN := bufio.NewReader(os.Stdin)
	stringSlice := []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stringSlice = append(stringSlice, scanner.Text())
	}

	/*
		for {
			line, err := bufStdIN.ReadString('\n')
			if err != nil && err == io.EOF {
				stringSlice = append(stringSlice, line)
				break
			} else {
				line = strings.Trim(line, "\n")
				stringSlice = append(stringSlice, line)
			}
		}
	*/
	return stringSlice
}
