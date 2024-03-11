package main

import "fmt"

// Adapted from https://www.makeuseof.com/golang-graph-data-structure/
type WaitForGraph struct {
	nodes map[int]*Node
}

type Node struct {
	OutNeighbors []*Node // outgoing edges
	InNeighbors  []*Node // incoming edges
}

func NewWaitForGraph() *WaitForGraph {
	return &WaitForGraph{
		nodes: make(map[int]*Node),
	}
}

func (g *WaitForGraph) AddNode(nodeID int) {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := &Node{
			OutNeighbors: []*Node{},
			InNeighbors:  []*Node{},
		}
		g.nodes[nodeID] = newNode
		fmt.Println("New node added to graph")
	} else {
		fmt.Println("Node already exists!")
	}
}

func (g *WaitForGraph) AddEdge(fromNodeID, toNodeID int) {
	node1 := g.nodes[fromNodeID]
	node2 := g.nodes[toNodeID]
	node1.OutNeighbors = append(node1.OutNeighbors, node2)
	node2.InNeighbors = append(node2.InNeighbors, node1)
	fmt.Println("New edge added to graph")
}

func (g *WaitForGraph) RemoveEdge(fromNode, toNode *Node) {
	index := -1
	for i, n := range fromNode.OutNeighbors {
		if n == toNode {
			index = i
			break
		}
	}
	if index != -1 {
		fromNode.OutNeighbors =
			append(fromNode.OutNeighbors[:index], fromNode.OutNeighbors[index+1:]...)
		fmt.Println("fromNode Edge Removed")
	}

	for i, n := range toNode.InNeighbors {
		if n == toNode {
			index = i
			break
		}
	}
	if index != -1 {
		toNode.InNeighbors =
			append(toNode.InNeighbors[:index], toNode.InNeighbors[index+1:]...)
		fmt.Println("toNode Edge Removed")
	}
}
