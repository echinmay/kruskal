package main

import (
	"fmt"
	"container/heap"
	"github.com/echinmay/unionfind"
	"bufio"
	"flag"
	"strconv"
	"strings"
	"os"
	"log"
)

type Edge struct {
	source  int
	peer	int
	weight	int
}

type EdgesSorted []*Edge

func (ed EdgesSorted) Len() int { 
	return len(ed) 
}

func (ed EdgesSorted) Less(i, j int) bool {
	return ed[i].weight < ed[j].weight
}

func (ed EdgesSorted) Swap(i, j int) {
	ed[i], ed[j] = ed[j], ed[i]
}

func (ed *EdgesSorted) Push(x interface{}) {
	edge := x.(*Edge)
	*ed = append(*ed, edge)
}

func (ed *EdgesSorted) Pop() interface{} {
	old := *ed
	n := len(old)
	item := old[n-1]
	*ed = old[0:n-1]
	return item
}

func main() {

	var inputfile string

	// Get the input file if sent as command line argument or use default
	flag.StringVar(&inputfile, "input", "input.txt", "<NAME OF INPUT FILE>")
	flag.Parse()

	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	inputs := strings.Fields(line)

	numnodes, _ := strconv.Atoi(inputs[0])
	uf := unionfind.New(numnodes)

	edges := make(EdgesSorted, 0)

	heap.Init(&edges)

	for scanner.Scan() {
		line = scanner.Text()
		if len(line) == 0 {
			continue
		}
		inputs = strings.Fields(line)

		source, err := strconv.Atoi(inputs[0])
		if err != nil {
			log.Fatal("Wrong format")
		}
		destination, err := strconv.Atoi(inputs[1])
		if err != nil {
			log.Fatal("Wrong format")
		}
		weight, err := strconv.Atoi(inputs[2])
		if err != nil {
			log.Fatal("Wrong format")
		}
		heap.Push(&edges, &Edge{source, destination, weight})
	}
	
	MST := make([]*Edge, 0)
	for edges.Len() > 0 && uf.GetNumClusters() > 1 {
		edgeout := heap.Pop(&edges).(*Edge)	
		if uf.Connected(edgeout.source, edgeout.peer) != true {
			uf.Union(edgeout.source, edgeout.peer)
			MST = append(MST, edgeout)
		}
	}

	if uf.GetNumClusters() > 1 {
		fmt.Println("Could not get Minimum Spanning Tree")
	} else {
		for x := range MST {
			fmt.Println("Edges in MST")
			fmt.Println(*MST[x])
		}
	}

}