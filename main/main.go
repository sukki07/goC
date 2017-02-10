package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"
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
	//DiameterOfTree()
	//CreateBst()
	//BuyHatke()
	//BuyHatke2()
	//Zauba1()
	//Zauba2()
	//Swap()
	//ComradesUsingHash()
	//ComradesUsingTree()
	//MonkAndTreeCounting()
	//MancuAndColouredTree()
	//MirrorsAndTrees()
	/*
	event1 := &event{duration:5,title:"ev1"}
	event2 := &event{duration:5,title:"ev2"}
	s := SimpleScheduler{}
	c :=  s.Schedule([]IEvent{event1,event2})
	fmt.Println(&c)
	*/
	//hello()
	//Oz()
	//treebo()
	//hansel()
}



func hansel() {
	f,_ := os.Open("mapping.txt")
	scn := bufio.NewScanner(f)
	scn.Split(bufio.ScanLines)
	lines := make([]string,0)
	for scn.Scan(){
		lines = append(lines,scn.Text())
	}
	fmt.Println(lines)
	fileMap := make(map[string]map[string]string)
	for i:=0;i<len(lines);{
		m := isThisLineAFunction(lines[i],"int ")
		if m==false{
			//it is a class
			originalName := strings.Split(lines[i],"->")[0]
			obfuscatedName := strings.Split(lines[i],"->")[1]
			fileMap[originalName] = make(map[string]string)
			for j:=i+1;;j++{
				fm := isThisLineAFunction(lines[j],"int ")
				if fm==true{
					origFname,obsFuncName := parseFunction(lines[j])
					fileMap[originalName][origFname] = obfuscatedName + obsFuncName
				}else{
					i = j + 1
					break
				}
				if j == (len(lines) -1) {
					i = j+1
					break
				}
			}
		}
	}

	//get obfuscated name

	cn,fn := splitName("")
	on := fileMap[cn][fn]
	fmt.Println(on)

}

func splitName(line string) (string,string){
	return "",""
}


func parseFunction(line string) (string,string){
	return "",""
}

func isThisLineAFunction( line string,match string) bool{
	return true
}

func treebo() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	tc,_ := strconv.Atoi(scn.Text())
	for i:=0;i<tc;i++ {
		scn.Scan()
		n,_ := strconv.Atoi(scn.Text())
		totalSum := (n * (n+1) )/ 2
		if totalSum %2 != 0{
			fmt.Println("No")
		}else{
			fmt.Println("Yes")
		}	}


}
func Oz() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	nodes,_ := strconv.Atoi(scn.Text())
	nodeMap := make(map[int]int)
	for i:=1;i <= nodes ;i ++ {
		nodeMap[i] = 0
	}
	for i:=1;i<= nodes;i++{
		scn.Scan()
		outgoingForI,_ := strconv.Atoi(scn.Text())
		nodeMap[i] = outgoingForI
	}

	nodemaplist := make(map[int][]int)
	for k,v := range nodeMap {
		if _,ok := nodemaplist[k];!ok{
			nodemaplist[k]	 = []int{}
		}
		visit := make(map[int]bool)
		fillList(nodemaplist,nodeMap,k,v,visit)
	}
	fmt.Println(nodemaplist)


/*
	scn.Scan()
	queries,_ := strconv.Atoi(scn.Text())
	for j:=0;j<queries;j++{
		scn.Scan()
		typ,_ := strconv.Atoi(scn.Text())
		scn.Scan()
		val,_ := strconv.Atoi(scn.Text())
		switch  typ{
		case 1:
			findStoppage(nodeMap,val)
		case 2:
			nodeMap[val] = 0
		}
		//fmt.Println(typ,val)
	}
	//fmt.Println(nodeMap)
	*/
}

func fillList(nml map[int][]int,m map[int]int,fixedVal int,nodeval int,visited map[int]bool) {
	if nodeval == 0{
		return
	}
	if _,ok := visited[nodeval];ok{
		return
	}
	visited[nodeval] = true
	nml[fixedVal] = append(nml[fixedVal],nodeval)
	nodeval = m[nodeval]
	fmt.Println("gc")
	fillList(nml,m,fixedVal,nodeval,visited)
}

func findStoppage(m map[int]int,iniNode int)  {
	visited := make(map[int]bool)
	explore1(m,visited,iniNode)

}

func explore1(m map[int]int,visited map[int]bool,index int) {
	if _,ok := visited[index];ok{
		fmt.Println("LOOP")
		return
	}
	if m[index] == 0{
		fmt.Println(index)
	}else{
		visited[index] = true
		index = m[index]
		explore1(m,visited,index)
	}
}


func hello() {
	fmt.Println("hi")
}

type IEvent interface {
	Duration() int
	Title() string
	String() string
}

type event struct {
	duration int
	title string
}
func (e *event) String() string{
	return e.title
}

func (e *event) Duration() int{
	return e.duration
}

func (e *event) Title() string{
	return e.title
}

type ISession interface {
	Add(IEvent) error
	Remove(IEvent) error
	String() string
}

type DefaultSession struct {
	events []IEvent
	length int
}
func (ds *DefaultSession) Add(e IEvent) error{
	ds.events = append(ds.events,e)
	return nil
}
func (ds *DefaultSession) Remove(e IEvent) error{
	return nil
}

func (ds *DefaultSession) String() string{
	slice := make([]string,0)
	for _,e:= range ds.events{
		slice =append(slice,e.String())
	}
	return strings.Join(slice,"")
}


type MorningSession struct {
	*DefaultSession
}

type EveningSesson struct {
	*DefaultSession
}

type ITrack interface {
	GetSessions() []ISession
}

type DefaultTrack struct{
	sessions []ISession
}

func (dt *DefaultTrack) GetSessions() []ISession {
	return dt.sessions
}


func NewDefaultTrack() ITrack{
	sessions := make([]ISession,2)
	ms := &MorningSession{&DefaultSession{length:180,events:make([]IEvent,0)}}
	es := &EveningSesson{&DefaultSession{length:240,events:make([]IEvent,0)}}
	sessions[0] = ms
	sessions[1] = es
	return &DefaultTrack{sessions:sessions}
}

type Conference struct {
	tracks []ITrack
}
func (c *Conference) String() string{
	slice := make([]string,0)
	for _,t := range c.tracks {
		for _,s := range t.GetSessions() {
			slice = append(slice,s.String())
		}
	}
	return strings.Join(slice,"")
}

type Scheduler interface {
	Schedule([]IEvent) Conference
}

type SimpleScheduler struct {
}

func (s *SimpleScheduler) Schedule(events []IEvent) Conference{
	tracks:= make([]ITrack,0)
	t := NewDefaultTrack()
	tracks = append(tracks,t)
	for _,e := range events{
		if s.addEvent(e,t) {
		}else{
			t = NewDefaultTrack()
			tracks = append(tracks,t)
		}
	}
	return Conference{tracks:tracks}
}

func (s *SimpleScheduler) addEvent(e IEvent,t ITrack) bool{
	sess :=t.GetSessions()
	for _,s := range sess {
		err := s.Add(e)
		if err!=nil{
			continue
		}else{
			return true
		}
	}
	return false
}




func MirrorsAndTrees() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	tc,_ := strconv.Atoi(scn.Text())
	for i:=0;i<tc;i++{
		scn.Scan()
		nodes,_ := strconv.Atoi(scn.Text())
		tree := make([]int,nodes+1)
		for j:=1;j<=nodes;j++{
			scn.Scan()
			n,_ := strconv.Atoi(scn.Text())
			tree[j] = n

		}
		processMirror(tree)
		fmt.Println()
	}
}

func processMirror(tree []int) {
	rmap := make(map[int]int)
	lmap := make(map[int]int)
	visitNode(1,1,rmap,lmap,tree,"r")
	visitNode(1,1,rmap,lmap,tree,"l")

}

func visitNode(index int,height int,mr map[int]int,ml map[int]int,tree []int,d string) {
	if index >= len(tree){
		return
	}
	if tree[index] == 0 {
		return
	}
	if d=="r" {
		if _, ok := mr[height]; !ok {
			mr[height] =index
			fmt.Println(tree[index])
		}
		visitNode(2*(index)+1, height+1, mr,ml ,tree,d)
		visitNode(2*index, height+1, mr,ml, tree,d)

	}else {
		if _, ok := ml[height]; !ok {
			ml[height] =index
			if mr[height] != index {
				fmt.Println(tree[index])
			}
		}
		visitNode(2*index, height+1, mr,ml, tree,d)
		visitNode(2*(index)+1, height+1, mr,ml, tree,d)
	}
}

type cnode struct {
	number int
	color int
	ancestor int
	subs []*cnode
}


func MancuAndColouredTree() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	 scn.Scan()
	totalNodes,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	_,_ = strconv.Atoi(scn.Text())
	nodeMap := make(map[int]*cnode)
	for i:=1;i<=totalNodes;i++{
		nodeMap[i] = &cnode{number:i,color:0,subs:make([]*cnode,0),ancestor:-1}
	}
	for c:=2;c<=totalNodes;c++{
		scn.Scan()
		parent,_ := strconv.Atoi(scn.Text())
		nodeMap[parent].subs = append(nodeMap[parent].subs,nodeMap[c])
	}

	for c:=1;c<=totalNodes;c++{
		scn.Scan()
		cls,_ := strconv.Atoi(scn.Text())
		nodeMap[c].color = cls
	}
	isThisNodeColorAncestor(nodeMap[1])
	for k:=1;k<=totalNodes;k++{
		fmt.Printf("%d ",nodeMap[k].ancestor)
	}

}
type cp struct {
	number int
	colour int
}

func isThisNodeColorAncestor(n *cnode) []*cnode{
	if n==nil{
		return []*cnode{}
	}
	below := make([]*cnode,0)
	for _,c := range n.subs{
		below=append(below,isThisNodeColorAncestor(c)...)
	}
	//check if current node is ancestor of all the belows
	for _,b := range below{
		if b.ancestor == -1 &&  n.color == b.color{
			b.ancestor = n.number
		}
	}
	below = append(below,n)
	return  below
}

type tcNode struct {
	val int
	left *tcNode
	right *tcNode
}

func (n *tcNode) print(node *tcNode) {
	if node != nil {
		node.print(node.left)
		fmt.Println(node.val)
		node.print(node.right)
	}
}

func MonkAndTreeCounting() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	numNodes,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	Kval,_ := strconv.Atoi(scn.Text())
	nodePmap := make(map[int]*tcNode)
	for i:=1;i<=numNodes;i++{
		scn.Scan()
		v,_ := strconv.Atoi(scn.Text())
		nodePmap[i] = &tcNode{val:v,left:nil,right:nil}
	}
	for i:=2;i<= numNodes;i++{
		scn.Scan()
		parent,_ := strconv.Atoi(scn.Text())
		p := nodePmap[parent]
		c := nodePmap[i]
		if p.left == nil {
			p.left = c
		}else {
			p.right = c
		}
	}
	//nodePmap[1].print(nodePmap[1])
	var c int
	tcSubNodes(nodePmap[1],Kval,&c)
	fmt.Println(c)
}

func tcSubNodes(n *tcNode,Kval int,c *int) []int{
	if n==nil{
		return []int{}
	}
	ls := tcSubNodes(n.left,Kval,c)
	rs := tcSubNodes(n.right,Kval,c)
	cs := append(ls,rs...)
	pairs := findJK(cs)
	for _,p := range pairs{
		if p.a + p.b + n.val >= Kval {
			*c+=1
		}
	}
	cs  = append(cs,n.val)
	return cs
}

func findJK(input []int)  []soldierPair {
	retval := make([]soldierPair,0)
	for i:=0;i<len(input);i++ {
	    for j:=i+1;j<len(input);j++{
		    sp := soldierPair{a:input[i],b:input[j]}
		    retval= append(retval,sp)
	    }
    }
	return retval
}



func ComradesUsingTree() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalTests,_ := strconv.Atoi(scn.Text())
	for i:=0;i<totalTests;i++ {
		scn.Scan()
		noOfSoldiers, _ := strconv.Atoi(scn.Text())
		childMap := make(map[int][]int)
		for j := 1; j <= noOfSoldiers; j++ {
			scn.Scan()
			parent,_ :=  strconv.Atoi(scn.Text())
			if _,ok := childMap[parent];!ok{
				childMap[parent] = make([]int,0)
			}
			childMap[parent] = append(childMap[parent],j)
		}
		process(childMap,noOfSoldiers)
	}
}

func printTree(n *soldierNode) {
	if n.subs==nil{
		return
	}
	for _,v := range n.subs{
		printTree(v)
	}
}

func process(childMap map[int][]int,noOfSoldiers int) {
	val := childMap[0][0]
	commander := &soldierNode{number:val,subs:nil}
	makeTree(childMap,commander)
	var c int
	x := countChildern(commander,&c)
	total:=  (noOfSoldiers*(noOfSoldiers-1))/2
	value := x+c
	fb := total - value
	fmt.Printf("%d %d",value,fb)
	fmt.Println()

}

func countChildern(sn *soldierNode,c *int) int{
	if sn.subs == nil{
		return 0
	}
	var children int
	var y int
	for _,snc:= range sn.subs{
		y = countChildern(snc,c)
		*c = *c+y
		children +=y
	}
	length := len(sn.subs)
	return children+length
}

type soldierNode struct {
	number int
	subs []*soldierNode
}

func makeTree(m map[int][]int,sn *soldierNode) {
	if len(m[sn.number]) == 0 {
		return
	}
	childernSlice := m[sn.number]
	subSlice := make([]*soldierNode,len(childernSlice))
	for i,v := range childernSlice{
		s := &soldierNode{number:v,subs:nil}
		subSlice[i] = s
	}
	sn.subs = subSlice
	for _,sn := range sn.subs{
		makeTree(m,sn)
	}
}



func ComradesUsingHash() {
//input
	//pairs
	//search and find parent
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalTests,_ := strconv.Atoi(scn.Text())
	for i:=0;i<totalTests;i++{
		scn.Scan()
		noOfSoldiers,_ := strconv.Atoi(scn.Text())
		soldierDescMap := make(map[int]int)
		soldierSlice := make([]int,noOfSoldiers)
		for j:=1;j<=noOfSoldiers;j++{
			scn.Scan()
			soldierDescMap[j],_ =  strconv.Atoi(scn.Text())
			soldierSlice[j - 1] = j
		}
		pairs := findPairs(soldierSlice)
		hs :=0
		fb :=0
		for _,p := range pairs{
			if goTillCommander(p.a,p.b,soldierDescMap) {
				hs+=1
			}else if goTillCommander(p.b,p.a,soldierDescMap) {
				hs+=1
			}else{
				fb+=1
			}
		}
		fmt.Printf("%d %d",hs,fb)
		fmt.Println()
	}
}

func goTillCommander(s1 int,s2 int,m map[int]int) bool{
	for {
		if m[s1] == s2 {
			return true
		}else if m[s1] == 0{
			return false
		}else{
			s1 = m[s1]
		}
	}
	return false
}


type soldierPair struct {
	a int
	b int
}

func findPairs(input []int) []soldierPair{
	l := len(input)
	size := (l*(l-1)/2)
	sps := make([]soldierPair,size)
	count:=0
    for i:=0;i<len(input);i++ {
	    for j:=i+1;j<len(input);j++{
		    sp := soldierPair{a:input[i],b:input[j]}
		    sps[count] = sp
		    count+=1
	    }
    }
	return sps
}















































func Swap() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	tcs,_ :=strconv.Atoi(scn.Text())
	for i:=0;i<tcs;i++{
		scn.Scan()
		n,_ := strconv.Atoi(scn.Text())
		scn.Scan()
		k,_ := strconv.Atoi(scn.Text())
		r := happy(n,k)
		fmt.Println(r)
	}
}

func happy1(n int,k int) int{
	c := 1
	for i:=n;i>=(n-(k-1));i--{
		c*=(i)
	}
	return c
}

func happy(n int,k int) int{
	c := 1
	m:=1000000007
	for i:=n;i>=(n-(k-1));i--{
		c*=(i%m)
	}
	for j:=k;j>=1;j--{
	 c = c/j
	}
	c=c*(2%m)
	r := c%m
	return r
}



func Zauba1() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanLines)
	scn.Scan()
	totalTestCases,_ := strconv.Atoi(scn.Text())
	fmt.Println(totalTestCases)
	for i:=0;i<totalTestCases;i++{
		scn.Scan()
		noOfStrings,_ := strconv.Atoi(scn.Text())
		var c coll
		c = make(coll,noOfStrings)
		for j:=0;j<noOfStrings;j++{
			scn.Scan()
			s := scn.Text()
			c[j] = s
		}
		sort.Sort(c)
		c.Print()
	}
}

type coll []string
func (c coll) Print(){
	for _,v := range c{
		fmt.Println(v)
	}
}
func (c coll) Len() int{
	return len(c)
}

func (c coll) 	Less(i,j int) bool {
	si := c[i]
	sj := c[j]
	lsi := len(si)
	lsj := len(sj)
	var l int
	if lsi < lsj{
		l = lsi
	}else{
		l = lsj
	}

	for i:=0;i<l;i++{
		a:= strings.Split(si,"")[i]
		b:=strings.Split(sj,"")[i]
		la := strings.ToLower(a)
		lb := strings.ToLower(b)
		if la!=lb{
			if la == " " {
				return true
			}
			if lb == " " {
				return false
			}
			if la <lb {
				return true
			}else{
				return false
			}

		}
		if a == b {
			continue
		}
		if a == strings.ToLower(a) {
			return true
		}
		if b == strings.ToLower(b) {
			return false
		}
	}
	return true
}

func (c coll) Swap(i,j int) {
	c[i], c[j] = c[j], c[i]
}

func BuyHatke2() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalCities,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	totalSpecialCities,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	totalRoads,_ := strconv.Atoi(scn.Text())
	specialCityMap := make(map[int]bool)
	for i:=0;i<totalSpecialCities;i++ {
		scn.Scan()
		sc,_ := strconv.Atoi(scn.Text())
		specialCityMap[sc] = true
	}

	roadMap := make(map[int][]int)
	for k:=1;k<=totalCities;k++ {
		roadMap[k] = make([]int,0)
		roadMap[k] = append(roadMap[k],k)
	}
	for i:=0;i<totalRoads;i++ {
			scn.Scan()
			sc1,_:=strconv.Atoi(scn.Text())
			scn.Scan()
			sc2,_:=strconv.Atoi(scn.Text())
			roadMap[sc1] = append(roadMap[sc1],sc2)
			roadMap[sc2] = append(roadMap[sc2],sc1)
	}
	p:=0
	for city,_ := range roadMap{
		rm := make(map[int]bool)
		vm := make(map[int]bool)
		explore(city,roadMap,rm,vm)
		for k,_ := range rm {
			if specialCityMap[k]  == true{
				p+=1
			}
		}
	}
	fmt.Println(p)

}

func explore(city int,roadMap map[int][]int,reachable map[int]bool,visited map[int]bool) {
	for _,c := range roadMap[city] {
		reachable[c] = true
		if _,ok := visited[c] ; !ok {
			visited[c] = true
			explore(c,roadMap,reachable,visited)
		}
	}
}


func buildMaps(city int,rm map[int]map[int]bool,em map[int]map[int]bool) {

}


func BuyHatke() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalStudents,_ := strconv.Atoi( scn.Text())
	scn.Scan()
	totalRows,_ := strconv.Atoi(scn.Text())
	scn.Scan()
	maxRowSize,_ := strconv.Atoi(scn.Text())
	prefRows := make([]int,totalStudents)
	for i:=0;i<totalStudents;i++ {
		scn.Scan()
		prefRowNumber,_ := strconv.Atoi(scn.Text())
		prefRows[i] = prefRowNumber
	}
	//fmt.Println(totalRows)
	//fmt.Println(maxRowSize)
	//fmt.Println(prefRows)
	rowHash := make(map[int]int)
	for i:=1;i<=totalRows;i++ {
		rowHash[i] = 0
	}
	//fmt.Println(rowHash)
	totalCount := 0
	resultCount := 0
	for _,pRow := range prefRows{
		totalCount+=1
		if totalCount > totalRows*maxRowSize{
			resultCount+=1
			continue
		}
		if rowHash[pRow] < maxRowSize {
			rowHash[pRow]+=1
			//resultCount+=1
		}else{
			resultCount+=1
			nextRow := pRow
			for {
				//fmt.Println(rowHash)
				//fmt.Println(nextRow)
				if rowHash[nextRow]  < maxRowSize{
					rowHash[nextRow]+=1
					break
				}else{
					nextRow+=1
				}
				if nextRow==(totalRows+1){
					nextRow=1
				}
			}
		}
	}
	fmt.Println(resultCount)
}


func CreateBst() {
	scn := bufio.NewScanner(os.Stdin)
	scn.Split(bufio.ScanWords)
	scn.Scan()
	totalNodes,_ := strconv.Atoi(scn.Text())
	var root *Tnode
	for i:=0;i<totalNodes;i++{
		scn.Scan()
		val,_:= strconv.Atoi( scn.Text())
		if i==0 {
			root = &Tnode{data:val,left:nil,right:nil}
			continue
		}
		insertValInTree(val,root)
	}
	scn.Scan()
	inOrderNodeVal,_ := strconv.Atoi(scn.Text())
	tree := &Tree{root:root}
	node := tree.GetNodeForVal(inOrderNodeVal)
	tree.Print(node)
}



func insertValInTree(val int,root *Tnode)  {
	pointerToParent := root
	for pointerToParent!=nil {
		switch  {
		case val > pointerToParent.data:
			if pointerToParent.right!=nil{
				pointerToParent = pointerToParent.right
			}else{
				pointerToParent.right = &Tnode{data:val,left:nil,right:nil}
				return
			}
		case val <= pointerToParent.data:
			if pointerToParent.left!=nil{
				pointerToParent = pointerToParent.left
			}else{
				pointerToParent.left = &Tnode{data:val,left:nil,right:nil}
				return
			}
		}
	}
	return
}






type Tnode struct {
	data int
	left *Tnode
	right *Tnode
}

type Tree struct {
	root *Tnode
}

func (t *Tree) Print(node *Tnode) {
	t.PreOrder(node)
}

func (t *Tree) GetNodeForVal(data int) *Tnode {
	parent := t.root
	var node  *Tnode
	for parent!=nil{
		switch  {
		case data > parent.data:
			parent = parent.right
		case data < parent.data:
			parent = parent.left
		case data == parent.data:
			node = parent
			return node
		}
	}
	return node
}


func (t *Tree) PreOrder(node *Tnode)  {
	if node!=nil{
		fmt.Printf("%d\n",node.data)
		t.PreOrder(node.left)
		t.PreOrder(node.right)
	}
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
