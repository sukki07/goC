package main

import (
	"strings"
	"fmt"
	"sort"
	"bufio"
	"os"
	"strconv"
	"math"
)

type mountain struct {
	L int64
	R int64
}

func MonkAndSpecialInteger(){
	//input
	//func to calculate max sum of subarray of given size func(input,n)
	//binary search comparision with max
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	noOfints,_ := strconv.ParseInt(s.Text(),10,64)
	s.Scan()
	x,_:= strconv.ParseInt(s.Text(),10,64)
	input := make([]int64,noOfints)
	var i int64
	for i=0;i<noOfints;i++{
		s.Scan()
		input[i],_=strconv.ParseInt(s.Text(),10,64)
	}
	//MaxSubArraySum(input,2)
	o:=ModifiedBinarySearch(input,0,int64(len(input) - 1),x)
	fmt.Println(o)

}

func ModifiedBinarySearch(input []int64,start int64,end int64,x int64) int64 {
	if start <= end {
		mid := (start + end) / 2
		maxSubArraySumOfLengthMid := MaxSubArraySum(input, mid+1)
		switch {
		case maxSubArraySumOfLengthMid == x:
			return mid + 1
		case maxSubArraySumOfLengthMid < x:
			start = mid + 1
		case maxSubArraySumOfLengthMid > x:
			end = mid - 1
		}
		return ModifiedBinarySearch(input, start, end, x)
	}
	temp := MaxSubArraySum(input,start+1)
	if temp > x {
		return start  + 1 - 1
	}else{
		return start + 1
	}
}


func MaxSubArraySum(input []int64,subArraySize int64) int64{
	var start int64
	end := start + subArraySize - 1
	var firstSubArraySum int64
	for i:= start;i <= end;i++{
		firstSubArraySum+=input[i]
	}

	start++
	end++
	cs := firstSubArraySum
	ms:=firstSubArraySum
	for end < int64(len(input)) {
		elem := cs - input[start - 1] + input[end]
		if elem >= ms {
			ms = elem
		}
		cs = elem
		start++
		end++
	}
	return ms
}

func MonkAndMountains() {
	//read input
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	noOfMountains,_ := strconv.ParseInt(s.Text(),10,64)
	s.Scan()
	noOfQueries,_ := strconv.ParseInt(s.Text(),10,64)
	mountains := make([]mountain,noOfMountains)
	var i int64
	for i=0;i< noOfMountains ; i++{
		s.Scan()
		l,_ := strconv.ParseInt(s.Text(),10,64)
		s.Scan()
		r,_ := strconv.ParseInt(s.Text(),10,64)
		mountains[i] =  mountain{L:l,R:r}
	}

	queries := make([]int64,noOfQueries)
	for i=0;i < noOfQueries;i++{
		s.Scan()
		q,_ := strconv.ParseInt(s.Text(),10,64)
		queries[i] = q
	}

	cumulativeHeightsContributed := make([]int64,noOfMountains)
	var heightsContributed int64
	for index,mountain := range mountains {
		heightsContributed += mountain.R - mountain.L + 1
		cumulativeHeightsContributed[index] = heightsContributed
	}
	//fmt.Println(cumulativeHeightsContributed)
	for _,q := range queries {
		//res := BinarySearch(cumulativeHeightsContributed,q)
		index := BinarySearchInt64(cumulativeHeightsContributed,q,0,noOfMountains)
		var diff int64
		var hAtQ int64
		switch  {
		case cumulativeHeightsContributed[index] == q :
			hAtQ = mountains[index].R
		case cumulativeHeightsContributed[index] < q:
			diff = q - cumulativeHeightsContributed[index]
			hAtQ =  mountains[index + 1].L + diff - 1
		case cumulativeHeightsContributed[index] > q:
			if index > 0 {
				diff = q - cumulativeHeightsContributed[index - 1]
				hAtQ = mountains[index].L + diff - 1
			}else{
				hAtQ =  mountains[0].L + q - 1
			}
		}
		fmt.Println(hAtQ)
	}





	//keep counting values read , stop at x counts
}

func TestingInput() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	bytes := s.Bytes()
	fmt.Println(bytes)


}

type circularSearchPair struct{
	X int
	Y int
	R int //distance from origin, rounded off to next integer
}


type circularSortColl struct {
	base []*circularSearchPair
}

func (c circularSortColl) Less(i,j int) bool{
	if c.base[i].R <= c.base[j].R {
		return true
	}
	return false
}
func (c circularSortColl)Swap(i,j int) {
	temp := c.base[i]
	c.base[i] = c.base[j]
	c.base[j] = temp
}

func (c circularSortColl)Len() int {
	return len(c.base)
}


/*
func MonkAndCircularDistance() {
	//r := bufio.NewReader(os.Stdin)
	//take input
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	noOfpairs,_ := strconv.Atoi(s.Text())
	var circularSearchPairs []*circularSearchPair
	circularSearchPairs = make([]*circularSearchPair,noOfpairs)

	for i:=0;i < noOfpairs;i++ {
		s.Scan()
		xCoord,_ := strconv.Atoi(s.Text())
		s.Scan()
		yCoord,_ := strconv.Atoi(s.Text())
		circularSearchPairs[i] =  &circularSearchPair{X:xCoord,Y:yCoord,R:0}
	}

	//fmt.Println(circularSearchPairs)
	s.Scan()

	noOfQueries,_ := strconv.Atoi( s.Text())
	queries := make([]int,noOfQueries)
	for i:=0;i < noOfQueries;i++ {
		s.Scan()
		queries[i],_ =  strconv.Atoi(s.Text())
	}
	//fmt.Println(queries)
	//compute distance from origin and sort the collection on it
	querySlice := make([]int64,noOfpairs)
	for i:=0;i<len(circularSearchPairs);i++{
		r :=  distanceFromOrigin(circularSearchPairs[i].X,circularSearchPairs[i].Y)
		querySlice[i] = r
		circularSearchPairs[i].R = r
	}
	sort.Ints(querySlice)
	//coll := circularSortColl{base:circularSearchPairs}
	//sort.Sort(coll)
	for _,q:= range queries{
		res := BinarySearch(querySlice,q)
		noOfCoordsWithinq :=  0
		switch res.value {
		case AfterEnd:
			noOfCoordsWithinq = len(querySlice)
		case BeforeBeg:
			noOfCoordsWithinq = 0
		case WithinPresent:
			noOfCoordsWithinq =  res.endIndex + 1
		case WithinNotPresent:
			noOfCoordsWithinq =  (res.endIndex - 1) + 1
		}
		fmt.Println(noOfCoordsWithinq)
	}






	// make pair struct , with distance from origin as a field
	//store pairs in a array
	// sort the array on distance from origin ,
	// round off the distance to next integer value
	//query the array with given radius
}
*/


func distanceFromOrigin(x int,y int) int{
	distance := math.Sqrt(float64(x*x+y*y))
	nextInt := math.Ceil(distance)
	return int(nextInt)
}

func actualSearch(input []int64, fsv float64, start int, end int) int {
	if start >= end {
		return start
	}
	mid := (start + end) / 2
	if fsv < float64(input[mid]) {
		end = mid - 1
	} else {
		start = mid + 1
	}
	return actualSearch(input, fsv, start, end)
}

type BsearchResult struct {
	value    BsearchValue //after end , before beg , within_present  , within_notpresent
	begIndex int          //if within_present then beg , if within_notpresent then first index of just next greater element
	endIndex int          // if within then end ,if within_notpresent then first index of just next greater element
}
type BsearchValue int

const (
	AfterEnd         BsearchValue = iota // 0
	BeforeBeg                            // 1
	WithinPresent                        // 2
	WithinNotPresent                     // 3
)

func BinarySearchInt64(input[]int64, key int64,start int64,end int64) int64 {
	if start <= end {
		mid := (start + end)/2
		switch  {
		case input[mid]  == key :
			return mid
		case input[mid] > key:
			end = mid - 1
		case input[mid] < key:
			start = mid + 1
		}
		return BinarySearchInt64(input,key,start,end)
	}
	return start
}






//outside end
//before beg
//present start first index
//present end last index
//not present start,end - index of just next element [7,7,7,9,9,9,9] searching for 8 will return index of first 9
func BinarySearch(input []int64, sv int64) *BsearchResult {
	var fsv float64
	var index int
	var res *BsearchResult
	l := len(input)
	//determing if in limits or not
	switch {
	case sv < input[0]:
		return &BsearchResult{BeforeBeg, 0, 0}
	case sv > input[l-1]:
		return &BsearchResult{AfterEnd, 0, 0}
	}

	//determine if present or not
	fsv = float64(sv) - 0.1
	index = actualSearch(input, fsv, 0, len(input)-1)
	begIndex := 0
	endIndex := 0
	var val BsearchValue
	switch {
	case input[index] == sv:
		val = WithinPresent
		begIndex = index
	case input[index] < sv && index+1 < len(input) && input[index+1] == sv:
		val = WithinPresent
		begIndex = index + 1
	case input[index] > sv:
		val = WithinNotPresent
		begIndex = index
		endIndex = index
	case input[index] < sv && index+1 < len(input) && input[index+1] > sv:
		val = WithinNotPresent
		begIndex = index + 1
		endIndex = index + 1
	}

	//if not present return if present find end too
	switch val {
	case WithinNotPresent:
		res = &BsearchResult{value: WithinNotPresent, begIndex: begIndex, endIndex: endIndex}
	case WithinPresent:
		fsv = float64(sv) + 0.1
		index = actualSearch(input, fsv, 0, len(input)-1)
		switch {
		case input[index] == sv:
			endIndex = index
		case input[index] > sv && index-1 >= 0 && input[index-1] == sv:
			endIndex = index - 1
		}
		res = &BsearchResult{value: WithinPresent, begIndex: begIndex, endIndex: endIndex}
	}
	return res
}

type query struct {
	tp  int
	num int
}

func MonkAndSearch() {
	//fmt.Println(time.Now())
	scn := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 100000*10)
	scn.Buffer(buf, 100000*10)
	scn.Scan()
	l, _ := strconv.Atoi(scn.Text())
	l++
	scn.Scan()
	slice := strings.Split(scn.Text(), " ")
	sliceint := make([]int, len(slice))
	for index, val := range slice {
		sliceint[index], _ = strconv.Atoi(val)
	}
	scn.Scan()
	q, _ := strconv.Atoi(scn.Text())
	sort.Ints(sliceint)
	qs := make([]query, q)
	for i := 0; i < q; i++ {
		tp := 0
		num := 0
		scn.Scan()
		s := scn.Text()
		tp, _ = strconv.Atoi(strings.Split(s, " ")[0])
		num, _ = strconv.Atoi(strings.Split(s, " ")[1])
		t := query{tp: tp, num: num}
		qs[i] = t
	}
	w := bufio.NewWriter(os.Stdout)
	processSearch(qs, sliceint, w)
}

func processSearch(qs []query, slice []int, w *bufio.Writer) {
	//fmt.Println(time.Now())
	for _, t := range qs {
		bsearch(t.num, slice, t, w)
	}
	w.Flush()
	//fmt.Println(time.Now())
}

func work(sv float32, orig int, slice []int) int {
	end := len(slice) - 1
	start := 0
	mid := (start + end) / 2
	for {
		if start == end {
			break
		} else if end == start-1 {
			break
		}
		if float32(slice[mid]) > sv {
			end = mid - 1
			mid = (start + end) / 2
		} else {
			start = mid + 1
			mid = (start + end) / 2
		}
	}
	return start
}

func gthanx(sv int, slice []int) int {
	searchValue := float32(sv) + 0.1
	pos := work(searchValue, sv, slice)
	gthanx := 0
	if slice[pos] <= sv {
		gthanx = (len(slice) - 1) - (pos + 1) + 1
	} else {
		gthanx = (len(slice) - 1) - (pos) + 1
	}
	return gthanx
}

func gthanequaltox(sv int, slice []int) int {
	searchValue := float32(sv) - 0.1
	pos := work(searchValue, sv, slice)
	gthanx := 0
	if slice[pos] >= sv {
		gthanx = (len(slice) - 1) - (pos) + 1
	} else {
		gthanx = (len(slice) - 1) - (pos + 1) + 1
	}
	return gthanx
}

func bsearch(sv int, slice []int, t query, w *bufio.Writer) {
	p := 0
	if t.tp == 0 {
		if sv > slice[len(slice)-1] {
			p = 0
		} else if sv < slice[0] {
			p = len(slice)
		} else {
			p = gthanequaltox(sv, slice)
		}
	} else {
		if sv > slice[len(slice)-1] {
			p = 0
		} else if sv < slice[0] {
			p = len(slice)
		} else {
			p = gthanx(sv, slice)
		}
	}
	w.WriteString(strconv.Itoa(p))
	w.WriteString("\n")
}

func MonkTakesAWalk() {
	tc := 0
	fmt.Scanf("%d", &tc)
	for i := 0; i < tc; i++ {
		input := ""
		fmt.Scanf("%s", &input)
		count := processWalk(input)
		fmt.Println(count)
	}
}

func processWalk(input string) int {
	count := 0
	slice := strings.Split(input, "")
	for _, val := range slice {
		switch val {
		case "a":
			count++
		case "e":
			count++
		case "i":
			count++
		case "o":
			count++
		case "u":
			count++
		case "A":
			count++
		case "E":
			count++
		case "I":
			count++
		case "O":
			count++
		case "U":
			count++
		}
	}
	return count
}

func MonkBeingMonitor() {
	var totalTestCases int
	fmt.Scanf("%d", &totalTestCases)
	for i := 0; i < totalTestCases; i++ {
		var length int
		fmt.Scanf("%d", &length)
		sliceNum := make([]int, length)
		for j := 0; j < length; j++ {
			fmt.Scanf("%d", &sliceNum[j])
		}
		process(sliceNum)
	}
}

func process(slice []int) {
	sort.Ints(slice)
	mfp := MostFrequent(slice)
	lfp := LeastFrequent(slice)
	if mfp.Freq == lfp.Freq {
		fmt.Println("-1")
	} else {
		n := 0
		if mfp.Freq > lfp.Freq {
			n = mfp.Freq - lfp.Freq
		} else {
			n = lfp.Freq - mfp.Freq
		}
		fmt.Println(n)
	}
}

type pair struct {
	Val  int
	Freq int
}

func LeastFrequent(slice []int) pair {
	f := 0
	v := slice[0]

	oldPair := pair{Val: 0, Freq: 1000001}
	for index, val := range slice {
		if val == v && index != (len(slice)-1) {
			f++
			continue
		} else if val == v {
			f++
			cp := pair{Freq: f, Val: v}
			if cp.Freq < oldPair.Freq {
				oldPair = cp
			}
		} else {
			cp := pair{Freq: f, Val: v}
			if cp.Freq < oldPair.Freq {
				oldPair = cp
			}
			v = val
			f = 1
		}
	}
	return oldPair
}

func MostFrequent(slice []int) pair {
	f := 0
	v := slice[0]

	oldPair := pair{Val: 0, Freq: 0}
	for index, val := range slice {
		if val == v && index != (len(slice)-1) {
			f++
			continue
		} else if val == v {
			f++
			cp := pair{Freq: f, Val: v}
			if cp.Freq > oldPair.Freq {
				oldPair = cp
			}
		} else {
			cp := pair{Freq: f, Val: v}
			if cp.Freq > oldPair.Freq {
				oldPair = cp
			}
			v = val
			f = 1
		}
	}
	return oldPair
}


type MonkAlgoColl struct {
	base     []string
	chunkVal int
}

func (mac MonkAlgoColl) Len() int {
	return len(mac.base)
}

func ReverseSlice(s []string) []string {
	l := len(s)
	for x := 0; x < l/2; x++ {
		//fmt.Println(l)
		//fmt.Println(x)
		temp := s[x]
		s[x] = s[l-1-x]
		s[l-1-x] = temp
	}
	return s
}

func (mac MonkAlgoColl) getNumberAt(i int) int {
	startIndex := (5*mac.chunkVal - 4 - 1)
	endIndex := 5*mac.chunkVal - 1
	orig := strings.Split(mac.base[i], "")

	ReverseSlice(orig)

	//fmt.Println(orig)

	sliceOfNumbers := orig

	if endIndex < len(mac.base[i]) {
		num, _ := strconv.Atoi(strings.Join(ReverseSlice(sliceOfNumbers[startIndex:endIndex+1]), ""))
		return num
	}
	tempSlice := make([]string, 0)
	for t := startIndex; t < len(sliceOfNumbers); t++ {
		tempSlice = append(tempSlice, sliceOfNumbers[t])
	}
	for t := len(tempSlice); t < 5; t++ {
		tempSlice = append(tempSlice, "0")
	}
	num1, _ := strconv.Atoi(strings.Join(ReverseSlice(tempSlice), ""))
	return num1
}

func (mac MonkAlgoColl) Less(i, j int) bool {
	iNumber := mac.getNumberAt(i)
	jNumber := mac.getNumberAt(j)
	//fmt.Printf("comparing numbers with index at %d and %d as %d with %d", i, j, iNumber, jNumber)
	//fmt.Println()
	if iNumber < jNumber {
		return true
	} else if iNumber > jNumber {
		return false
	} else if i < j {
		return true
	}
	return false
}
func (mac MonkAlgoColl) Swap(i, j int) {
	temp := mac.base[i]
	mac.base[i] = mac.base[j]
	mac.base[j] = temp
}

func MonkAndSortingAlgorithm() {
	sliceOfNumbersAsStrings := takeInput4()
	max := 0
	for _, val := range sliceOfNumbersAsStrings {
		l := len(val)
		if l > max {
			max = l
		}
	}
	completeLoops := max / 5
	rem := max % 5
	loops := completeLoops
	if rem != 0 {
		loops = completeLoops + 1
	}

	coll := MonkAlgoColl{base: sliceOfNumbersAsStrings, chunkVal: 0}
	for i := 1; i <= loops; i++ {
		coll.chunkVal = i
		sort.Stable(coll)
		//fmt.Println(strings.Trim(fmt.Sprint(coll.base), "[]"))
	}

}

func takeInput4() []string {
	var length int
	fmt.Scanf("%d", &length)
	/*
		scn := bufio.NewScanner(os.Stdin)
		scn.Split(bufio.ScanWords)
		slice := make([]string, length)
		i := 0
		for scn.Scan() {
			slice[i] = scn.Text()
			i++
		}
		return slice
	*/
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.Trim(s, "\n")
	sl := strings.Split(s, " ")
	//fmt.Println(sl)
	return sl

}

func MonkAndNiceStrings() {
	slice := takeInput3()
	for i := 0; i < len(slice); i++ {
		iNiceVal := 0
		for j := i - 1; j >= 0; j-- {
			istr := slice[i]
			jstr := slice[j]
			if strings.Compare(jstr, istr) == -1 {
				iNiceVal++
			}
		}
		fmt.Println(iNiceVal)
	}
}

func takeInput3() []string {
	var noOfStrings int
	var str string
	var slice []string
	fmt.Scanf("%d", &noOfStrings)
	slice = make([]string, noOfStrings)
	for i := 0; i < noOfStrings; i++ {
		fmt.Scanf("%s", &str)
		slice[i] = str
	}
	return slice
}

type suffixSortCollection struct {
	baseSliceChar []string
	sliceIndex    []int
}

func (ssc suffixSortCollection) Len() int {
	return len(ssc.sliceIndex)
}
func (ssc suffixSortCollection) Less(i, j int) bool {
	islice := ssc.baseSliceChar[i:len(ssc.baseSliceChar)]
	jslice := ssc.baseSliceChar[j:len(ssc.baseSliceChar)]
	istring := strings.Join(islice, "")
	jstring := strings.Join(jslice, "")
	result := strings.Compare(istring, jstring)
	if result == -1 {
		return true
	}
	return false
}
func (ssc suffixSortCollection) Swap(i, j int) {
	temp := ssc.sliceIndex[i]
	ssc.sliceIndex[i] = ssc.sliceIndex[j]
	ssc.sliceIndex[j] = temp
}

func MonkAndSuffixSort() {
	sliceOfChars, ksmallest := takeInput2()
	indexSlice := make([]int, len(sliceOfChars))
	for index, _ := range indexSlice {
		indexSlice[index] = index
	}
	sortColl := suffixSortCollection{baseSliceChar: sliceOfChars, sliceIndex: indexSlice}
	//fmt.Println(sortColl.sliceIndex)
	sort.Stable(sortColl)
	suffixIndex := sortColl.sliceIndex[ksmallest-1]
	//fmt.Println(sortColl.sliceIndex)
	sliceToBePrinted := sortColl.baseSliceChar[suffixIndex:len(sortColl.baseSliceChar)]
	stringToBePrinted := strings.Join(sliceToBePrinted, "")
	fmt.Println(stringToBePrinted)
}

func takeInput2() ([]string, int) {
	var input string
	var index int
	fmt.Scanf("%s %d", &input, &index)
	return strings.Split(input, ""), index
}

func takeInput1() ([]int, int) {
	var length int
	var numberToBeModuleWith int
	fmt.Scanf("%d%d\n", &length, &numberToBeModuleWith)
	var num int
	var inputSlice []int
	inputSlice = make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &num)
		inputSlice[i] = num
	}
	fmt.Scanf("\n")
	return inputSlice, numberToBeModuleWith
}

type moduloSortCollection struct {
	intSlice     []int
	moduloNumber int
}

func (ifs moduloSortCollection) Len() int {
	return len(ifs.intSlice)

}
func (ifs moduloSortCollection) Less(i, j int) bool {
	if ifs.intSlice[i]%ifs.moduloNumber < ifs.intSlice[j]%ifs.moduloNumber {
		return true
	} else if ifs.intSlice[i]%ifs.moduloNumber > ifs.intSlice[j]%ifs.moduloNumber {
		return false
	} else if i < j {
		return true
	}
	return false
}
func (ifs moduloSortCollection) Swap(i, j int) {
	temp := ifs.intSlice[i]
	ifs.intSlice[i] = ifs.intSlice[j]
	ifs.intSlice[j] = temp
}

func MonkAndModuloSorting() {
	sliceOfNumbers, moduloNumber := takeInput1()
	coll := moduloSortCollection{intSlice: make([]int, len(sliceOfNumbers)), moduloNumber: moduloNumber}
	for index, val := range sliceOfNumbers {
		coll.intSlice[index] = val
	}
	sort.Stable(coll)
	fmt.Println(strings.Trim(fmt.Sprint(coll.intSlice), "[]"))
}
