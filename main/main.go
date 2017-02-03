package main

import ()
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//InputOutput()
	//CountDigits()
	//CountOnesInBinaryRep()
	//NQueen()
	//-------------------
	//MonkAndWelcomeProblem()
	//MonkTeachesPalindrome()
	//MonkAndInversions()
	//MonkAndGoodString()
	//MonkAndRoatation()
	//-------------------
	//MonkAndModuloSorting()
	//MonkAndSuffixSort() //try without taking memory for suffixes - done
	//MonkAndNiceStrings() //try without comparing each string with all previous ones
	//MonkAndSortingAlgorithm() // tle for 1 lakh input , try doing without reversing slice twice , can radix sort be used in this , or the problem is printing , try buffered print
	//MonkBeingMonitor()
	//--------------------
	//MonkTakesAWalk()
	//MonkAndSearch()
	//TestingInput()
	//MonkAndCircularDistance() //use float64 instead of float32 , while finding end and start of number in binary search
	//MonkAndMountains()
	//MonkAndSpecialInteger()
	//-------------------------
	//MonkAndPhilosophersStone()
	//MonkAndChamberOfSecrets() //done by storing , original position in the node itself , later if possible try storing not the position but just the movements to find original position , though not much will be gained memory wise
	//MonkAndAzkaban() //pending due to don't know how to use stack tle for most of the cases. TODO
	//MonkAndGobletOfFire() //could have used array to store pointers , instead of four variables
	//get min from first row -either loop or maintain a stack of subsequent smaller entries
	//MonkAndOrderOfPhoenix()
	//-----------------------------
	//MonkAndSquareRoot()
	//-----------------Skipping Number Theory For Later
	//Checkpoint 1
	//MonkAndLuckyMinimum()
	//MonkAndOperations()
	//----------------------------Skipping checkpoint and doing topic wise each
	DiameterOfTree()
}

type Tnode struct {
	data int
	left *Tnode
	right *Tnode
}

type Tree struct {
	root *Tnode
}

func (t *Tree) Print() {
}


func DiameterOfTree() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalNodes,_ := strconv.Atoi( scn.Text())
	scn.Scan()
	rootValue,_ := strconv.Atoi(scn.Text())
	rootNode := Tnode{data:rootValue,left:nil,right:nil}
	for i:=0;i<totalNodes - 1;i++ {
		scn.Scan()
		direction := scn.Text()
		var pointerToParentWhereChildToBeInserted  *Tnode
		pointerToParentWhereChildToBeInserted = &rootNode
		slice := strings.Split(direction,"")
		for j:=0;j<len(slice) - 1;j++ {
			switch slice[j] {
			case "L":
				if pointerToParentWhereChildToBeInserted.left == nil{
					pointerToParentWhereChildToBeInserted.left = &Tnode{}
				}
				pointerToParentWhereChildToBeInserted = pointerToParentWhereChildToBeInserted.left
			case "R":
				if pointerToParentWhereChildToBeInserted.right == nil {
					pointerToParentWhereChildToBeInserted.right = &Tnode{}
				}
				pointerToParentWhereChildToBeInserted=pointerToParentWhereChildToBeInserted.right
			}
		}
		scn.Scan()
		val,_ := strconv.Atoi( scn.Text() )
		switch slice[len(slice) - 1] {
		case "L":
			if pointerToParentWhereChildToBeInserted.left == nil {
				pointerToParentWhereChildToBeInserted.left = &Tnode{data: val, left: nil, right: nil}
			}else{
				pointerToParentWhereChildToBeInserted.left.data = val
			}

		case "R":
			if pointerToParentWhereChildToBeInserted.right == nil {
				pointerToParentWhereChildToBeInserted.right = &Tnode{data: val, left: nil, right: nil}
			}else{
				pointerToParentWhereChildToBeInserted.right.data = val
			}

		}
	}
	tree := Tree{root:&rootNode}
	//tree.Print()
	findDiameter(tree)
}

func findDiameter(t Tree){
	maxd:=0
	height(t.root,&maxd)
	fmt.Println(maxd)
}

func height(n *Tnode,cd *int) int{
	if n==nil{
		return 0
	}
	l := height(n.left,cd)
	r := height(n.right,cd)
	d := l+r+1
	if d>*cd{
		*cd = d
	}
	//fmt.Println(l,r)
	if l>r{
		return l+1
	}else{
		return r+1
	}
}






type A struct {
	a int
}
type B struct {
	b int
	*A
}
func (a A) inc(){
	a.a++
}

func TestEmbedding() {
	b := B{b:1,A:&A{a:2}}
	fmt.Println(b.a)
	b.inc()
	fmt.Println(b.a)
}

func MonkAndOperations() {
	scn:= bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	nVal,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	mVal,_ := strconv.Atoi(scn.Text())
	matrix := make([][]int,nVal)
	for i:=0;i<nVal;i++{
		row  := make([]int,mVal)
		for j:=0;j<mVal;j++{
			scn.Scan()
			e,_ :=  strconv.Atoi(scn.Text())
			row[j] = e
		}
		matrix[i] = row
	}
	vals := make([]int,4)
	for i:=0;i<4;i++ {
		scn.Scan()
		v,_ := strconv.Atoi( scn.Text())
		vals[i] = v
	}
	rowOpSum := 0
	for _,s := range matrix{
		maxRowSum := findRowMax(s,vals[0],vals[1])
		rowOpSum+=maxRowSum
	}
	colOpSum := 0
	for i:=0 ; i < mVal ; i++ {
		col := make([]int,nVal)
		for j:=0;j< nVal;j++ {
			col[j]	= matrix[j][i]
		}
		maxColSum := findRowMax(col,vals[2],vals[3])
		//fmt.Println(maxColSum)
		colOpSum+=maxColSum
	}
	if rowOpSum>=colOpSum{
		fmt.Println(rowOpSum)
	}else {
		fmt.Println(colOpSum)
	}
}

func findRowMax(s []int,v1 int,v2 int) int{
	noOpSum :=0
	for _,e := range s {
		noOpSum+=abs(e)
	}
	v1Sum := 0
	for _,e:= range s{
	v1Sum+=	abs(e+v1)
	}
	v2Sum := abs(v2 * len(s))
	sumSlice := make([]int,3)
	sumSlice[0] = noOpSum
	sumSlice[1] = v1Sum
	sumSlice[2] = v2Sum
	return findMaxInSlice(sumSlice)
}

func abs(e int) int{
	if e == 0 {
		return 0
	}
	if e < 0 {
		return -1 * e
	}
	return e

}

func findMaxInSlice(s []int) int {
	max := s[0]
	for _,e := range s {
		if e>max {
			max = e
		}
	}
	return max
}

func MonkAndLuckyMinimum() {
	scn:= bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	testCases,_ := strconv.Atoi(scn.Text())
	input := make([][]int,testCases)
	for i:=0;i< testCases;i++ {
		scn.Scan()
		lengthOfArray,_ := strconv.Atoi(scn.Text())
		slice := make([]int,lengthOfArray)
		for j:=0;j<lengthOfArray;j++ {
			scn.Scan()
			elem,_ := strconv.Atoi(scn.Text())
			slice[j] = elem
		}
		input[i] = slice
	}
	for _,s := range input {
		min:=findMinInSlice(s)
		freq := findFreqOfElem(s,min)
		switch freq % 2 {
		case 0:
			fmt.Println("Unlucky")
		default:
			fmt.Println("Lucky")
		}
	}

}

func findMinInSlice(s []int) int{
	min := s[0]
	for _,e := range s {
		if e < min {
			min = e
		}
	}
	return min
}

func findFreqOfElem(s []int,e int) int{
	count :=0
	for _,v := range s {
		if v == e {
			count++
		}
	}
	return count
}








func MonkAndSquareRoot() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	noOfTestCases,_ := strconv.Atoi(scn.Text())
	for i:=0;i<noOfTestCases ;i++  {
		scn.Scan()
		nVal,_ := strconv.Atoi(scn.Text())
		scn.Scan()
		mVal,_:=strconv.Atoi(scn.Text())
		findX(nVal,mVal)
	}
}
func findX(n int,m int) {
	for i:=0;;i++ {
		if ((i*i) % m )== n {
			fmt.Println(i)
		}
	}
}
