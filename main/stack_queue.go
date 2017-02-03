package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"container/list"
	"errors"
)

func TestB(k int64, l int64, h int64, s []int64) (int64, error) {
	var val int64
	switch {
	case h-l == -1 && h+1 < int64(len(s)):
		val = h+1
	case h-l == -1:
		return 0,errors.New("no key greater")
	default:
		midIntervalIndex := (h + l) / 2
		mid := s[midIntervalIndex]
		switch {
		case mid > k:
			h = midIntervalIndex - 1
		case mid <= k:
			l = midIntervalIndex + 1
		}
		return TestB(k, l, h, s)

	}
	return val,nil
}
type queryType struct {
	typ int64
	k   int64
	h   int64
}

func MonkAndOrderOfPhoenix() {
	//take input
	//find smallest in first row
	//query 2nd to last row for next greatest number (if there is a ith non empty  row)
	//next greates number via binary search of last occurence of number + 1 index
	//findNextGreatestElementInSlice(k,s)

	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	noOfRows, _ := strconv.ParseInt(scn.Text(), 10, 64)
	var i int64
	var j int64
	var rows [][]int64
	rows = make([][]int64, noOfRows)
	minStackForFirstRow := make([]int64, 0)
	for i = 0; i < noOfRows; i++ {
		scn.Scan()
		sizeOfRow, _ := strconv.ParseInt(scn.Text(), 10, 64)
		row := make([]int64, sizeOfRow)
		for j = 0; j < sizeOfRow; j++ {
			scn.Scan()
			elem, _ := strconv.ParseInt(scn.Text(), 10, 64)
			row[j] = elem
			if i == 0 {
				if j == 0 {
					minStackForFirstRow = append(minStackForFirstRow, elem)
				} else {
					if elem <= minStackForFirstRow[len(minStackForFirstRow) - 1] {
						minStackForFirstRow = append(minStackForFirstRow, elem)
					}
				}
			}
		}
		rows[i] = row
	}
	scn.Scan()
	noOfQueries, _ := strconv.ParseInt(scn.Text(), 10, 64)
	queries := make([]queryType, noOfQueries)
	for i = 0; i < noOfQueries; i++ {
		scn.Scan()
		qVal, _ := strconv.ParseInt(scn.Text(), 10, 64)
		switch qVal {
		case 0:
			scn.Scan()
			k, _ := strconv.ParseInt(scn.Text(), 10, 64)
			queries[i] = queryType{typ: 0, k: k, h: 0}
		case 1:
			scn.Scan()
			k, _ := strconv.ParseInt(scn.Text(), 10, 64)
			scn.Scan()
			h, _ := strconv.ParseInt(scn.Text(), 10, 64)
			queries[i] = queryType{typ: 1, k: k, h: h}
		case 2:
			queries[i] = queryType{typ: 2, k: 0, h: 0}
		}
	}
	for _, q := range queries {
		switch q.typ {
		case 0:
			rowNumber := q.k //fighter to be removed from this row
			if rowNumber == 1 {
				firstRow := rows[0]
				lastNum := firstRow[len(firstRow) - 1]
				if lastNum == minStackForFirstRow[len(minStackForFirstRow) - 1] {
					minStackForFirstRow = minStackForFirstRow[:(len(minStackForFirstRow) - 1)]
				}
				rows[0] = firstRow[:(len(firstRow) - 1)]
			}else{
				indexNumber := rowNumber - 1
				currentRow := rows[indexNumber]
				rowAfterRemoval := currentRow[:len(currentRow)-1]
				rows[indexNumber] = rowAfterRemoval
			}

		case 1:
			rowNumber := q.k
			if rowNumber == 1 {
				if q.h <= minStackForFirstRow[len(minStackForFirstRow) - 1] {
					minStackForFirstRow = append(minStackForFirstRow,q.h)
				}
				rows[0] = append(rows[0],q.h)
			}else {
				indexNumber := rowNumber - 1
				currentRow := rows[indexNumber]
				rowAfterAddtion := append(currentRow, q.h)
				rows[indexNumber] = rowAfterAddtion
			}

		case 2:
			var elem int64
			//min := findMinInFirstRow(rows[0])
			min := minStackForFirstRow[len(minStackForFirstRow) - 1]
			elem = min
			possible := true
			for _, r := range rows[1:] {
				indexOfNextGreatest, err := TestB(elem, 0, int64(len(r)-1), r)
				if err != nil {
					fmt.Println("NO")
					possible = false
					break
				}
				elem = r[indexOfNextGreatest]
			}
			if possible == true {
				fmt.Println("YES")
			}

		}
	}

	/*
		s := []int64{3,4,4,4,4,4,4,4,4,4,4,4,4,5,6,7,8,9}
		ans,err := FindNextGreatestToKey(4,0,int64(len(s) - 1),s)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(ans)
	*/
}

func findMinInFirstRow(s []int64) int64 {
	var min int64
	min = s[0]
	for _, val := range s {
		if val <= min {
			min = val
		}
	}
	return min

}

























//search for last occurence of given key if no key then just greatest than key
func FindNextGreatestToKey(k int64, l int64, h int64, s []int64) (int64, error) {
	intervalLength := h - l
	switch {
	case intervalLength == 0:
		switch {
		case s[h] > k:
			return h, nil
		case h+1 < int64(len(s)):
			return h + 1, nil
		default:
			return 0, errors.New("no elem > than key")
		}
	case intervalLength == 1:
		switch {
		case s[l] > k:
			return l, nil
		case s[h] > k:
			return h, nil
		case h+1 < int64(len(s)):
			return h + 1, nil
		default:
			return 0, errors.New("no elem > than key")
			//no k
		}
	default:
		midIntervalIndex := (h + l) / 2
		mid := s[midIntervalIndex]
		switch {
		case mid > k:
			h = midIntervalIndex - 1
		case mid <= k:
			l = midIntervalIndex + 1
		}
		return FindNextGreatestToKey(k, l, h, s)
	}
}

//based on first time interval length 1 or 0
//basically bsearch always has a base case with interval length 0 or 1 , ie. l=h or l+1=h ; in both cases we can sya things based on m i.e l - 1 or h +1
func BSearchFirstKey(k int64, l int64, h int64, s []int64) {
	fmt.Println()
	intervalLength := h - l
	switch {
	case intervalLength == 0:
		switch {
		case s[l] == k:
			fmt.Println("a")
			fmt.Println(l)
		case s[l+1] == k:
			fmt.Println("b")
			fmt.Println(l + 1)
		}
	case intervalLength == 1:
		switch {
		case s[l] == k:
			fmt.Println("c")
			fmt.Println(l)
		case s[h] == k:
			fmt.Println("d")
			fmt.Println(h)
		case s[h+1] == k:
			fmt.Println("e")
			fmt.Println(h + 1)
		}
	default:
		midIntervalIndex := (h + l) / 2
		mid := s[midIntervalIndex]
		switch {
		case mid >= k:
			h = midIntervalIndex - 1
		case mid < k:
			l = midIntervalIndex + 1
		}
		BSearchFirstKey(k, l, h, s)
	}
}

type gobletQuery struct {
	val        string
	school     int
	rollNumber int
}

type student struct {
	school     int
	rollNumber int
}

func MonkAndGobletOfFire() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalQueries, _ := strconv.Atoi(scn.Text())
	queries := make([]gobletQuery, totalQueries)
	for i := 0; i < totalQueries; i++ {
		scn.Scan()
		typeOfQuery := scn.Text()
		switch typeOfQuery {
		case "E":
			scn.Scan()
			schoolNumber, _ := strconv.Atoi(scn.Text())
			scn.Scan()
			rollNumber, _ := strconv.Atoi(scn.Text())
			queries[i] = gobletQuery{val: "E", school: schoolNumber, rollNumber: rollNumber}
		case "D":
			queries[i] = gobletQuery{val: "D", school: 0, rollNumber: 0}
		}
	}
	l := list.New()
	var lastStudentOfSchool1 *list.Element
	lastStudentOfSchool1 = nil
	var lastStudentOfSchool2 *list.Element
	lastStudentOfSchool2 = nil
	var lastStudentOfSchool3 *list.Element
	lastStudentOfSchool3 = nil
	var lastStudentOfSchool4 *list.Element
	lastStudentOfSchool4 = nil

	for _, q := range queries {
		switch q.val {
		case "E":
			s := student{school: q.school, rollNumber: q.rollNumber}
			switch q.school {
			case 1:
				if lastStudentOfSchool1 == nil {
					e := l.PushBack(s)
					lastStudentOfSchool1 = e
				} else {
					e := l.InsertAfter(s, lastStudentOfSchool1)
					lastStudentOfSchool1 = e

				}
			case 2:
				if lastStudentOfSchool2 == nil {
					e := l.PushBack(s)
					lastStudentOfSchool2 = e
				} else {
					e := l.InsertAfter(s, lastStudentOfSchool2)
					lastStudentOfSchool2 = e

				}
			case 3:
				if lastStudentOfSchool3 == nil {
					e := l.PushBack(s)
					lastStudentOfSchool3 = e
				} else {
					e := l.InsertAfter(s, lastStudentOfSchool3)
					lastStudentOfSchool3 = e

				}
			case 4:
				if lastStudentOfSchool4 == nil {
					e := l.PushBack(s)
					lastStudentOfSchool4 = e
				} else {
					e := l.InsertAfter(s, lastStudentOfSchool4)
					lastStudentOfSchool4 = e

				}
			}
		case "D":
			e := l.Front()
			if e != nil {
				fmt.Printf("%d %d", e.Value.(student).school, e.Value.(student).rollNumber)
				switch e {
				case lastStudentOfSchool1:
					lastStudentOfSchool1 = nil
				case lastStudentOfSchool2:
					lastStudentOfSchool2 = nil
				case lastStudentOfSchool3:
					lastStudentOfSchool3 = nil
				case lastStudentOfSchool4:
					lastStudentOfSchool4 = nil
				}
				fmt.Println()
				l.Remove(e)
			}
		}
		/*
			fmt.Printf("list->")
			for e := l.Front();e != nil ; e = e.Next() {
				fmt.Printf(" %d %d|", e.Value.(student).school, e.Value.(student).rollNumber)
			}
			fmt.Println()
		*/
	}
}

func MonkAndAzkaban() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	length, _ := strconv.Atoi(scn.Text())
	array := make([]int64, length)
	for i := 0; i < length; i++ {
		scn.Scan()
		array[i], _ = strconv.ParseInt(scn.Text(), 10, 64)
	}
	for i := 0; i < len(array); i++ {
		var x int
		var y int
		x = -1
		y = -1
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[i] {
				x = j + 1 //1 based indexing
				break
			}
		}
		for k := i + 1; k < len(array); k++ {
			if array[k] > array[i] {
				y = k + 1 //1 based indexing
				break
			}
		}
		num := x + y
		fmt.Printf("%d ", num)
	}
}

type spider struct {
	power         int
	originalIndex int
}

//save initial index in the node  version
func MonkAndChamberOfSecrets() {
	//input
	//loop till x times
	// each loop creates a temp slice
	// function to loop over temp slice find max , prints its original index , decrements and reenques
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalSpiders, _ := strconv.Atoi(scn.Text())
	scn.Scan()
	totalIterations, _ := strconv.Atoi(scn.Text())
	spiders := make([]spider, totalSpiders)
	for i := 0; i < totalSpiders; i++ {
		scn.Scan()
		p, _ := strconv.Atoi(scn.Text())
		spiders[i] = spider{power: p, originalIndex: i + 1} //i+1 since position starts from 1
	}
	for i := 0; i < totalIterations; i++ {
		var tempSlice []spider
		if len(spiders) <= totalIterations {
			tempSlice = spiders[:]
			spiders = []spider{}
		} else {
			tempSlice = spiders[:totalIterations]
			spiders = spiders[totalIterations:]
		}
		toBeEnquedBack := processIteration(tempSlice)
		spiders = append(spiders, toBeEnquedBack...)
	}

}

func processIteration(temp []spider) []spider {
	max := 0
	foundIndex := 0
	for index, s := range temp {
		if s.power > max {
			max = s.power
			foundIndex = index
		}
	}
	fmt.Printf("%d ", temp[foundIndex].originalIndex)
	newtemp := make([]spider, len(temp)-1)
	j := 0
	for i := 0; i < len(temp); i++ {
		if i != foundIndex {
			newPower := (temp[i].power - 1)
			if newPower < 0 {
				newPower = 0
			}
			newtemp[j] = spider{power: newPower, originalIndex: temp[i].originalIndex}
			j++
		}
	}
	return newtemp
}

func MonkAndPhilosophersStone() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalCoins, _ := strconv.ParseInt(scn.Text(), 10, 64)
	var i int64
	coins := make([]int64, totalCoins)
	for i = 0; i < totalCoins; i++ {
		scn.Scan()
		val := scn.Text()
		coins[i], _ = strconv.ParseInt(val, 10, 64)
	}
	scn.Scan()
	totalQueries, _ := strconv.ParseInt(scn.Text(), 10, 64)
	scn.Scan()
	finalWorth, _ := strconv.ParseInt(scn.Text(), 10, 64)
	queries := make([]string, totalQueries)
	ms := NewStackInt64OnSlice()
	for i = 0; i < totalQueries; i++ {
		scn.Scan()
		queries[i] = scn.Text()
	}
	i = 0
	for _, query := range queries {
		switch query {
		case "Harry":
			ms.push(coins[i])
			i++
		case "Remove":
			ms.pop()
		}
		if ms.currentSum() == finalWorth {
			fmt.Println(ms.len())
			return
		}
	}
	fmt.Println("-1")

	//s := NewStackInt64OnSlice()
}

type StackInt64OnSlice struct {
	storage []int64
	sum     int64
}

func NewStackInt64OnSlice() *StackInt64OnSlice {
	return &StackInt64OnSlice{storage: make([]int64, 0), sum: 0}

}

func (s *StackInt64OnSlice) pop() error {
	if len(s.storage) < 1 {
		return errors.New("stack is zero , can't pop")
	}
	elem := s.storage[len(s.storage)-1]
	s.sum -= elem
	s.storage = s.storage[:len(s.storage)-1]
	return nil
}
func (s *StackInt64OnSlice) push(elem int64) {
	s.storage = append(s.storage, elem)
	s.sum += elem

}
func (s *StackInt64OnSlice) len() int {
	return len(s.storage)
}

func (s *StackInt64OnSlice) currentSum() int64 {
	return s.sum
}

func (s *StackInt64OnSlice) print() {
	for _, val := range s.storage {
		fmt.Printf("%d,", val)
	}
	fmt.Println()
}
