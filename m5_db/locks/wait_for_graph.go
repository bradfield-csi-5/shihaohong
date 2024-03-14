package main

import "fmt"

// Adapted from https://www.makeuseof.com/golang-graph-data-structure/
type WaitForGraph struct {
	nodes map[int]*Node
}

type Node struct {
	OutNeighbors []*NodeEdge // outgoing edges
	InNeighbors  []*NodeEdge // incoming edges
}

type NodeEdge struct {
	node           *Node
	associatedLock Lock
}

func NewWaitForGraph() *WaitForGraph {
	return &WaitForGraph{
		nodes: make(map[int]*Node),
	}
}

func (g *WaitForGraph) hasCycles(node int, visited map[int]bool) bool {
	for i := range g.nodes[node].OutNeighbors {
		if visited[i] {
			return true
		} else {
			visited[i] = true
			return g.hasCycles(node, visited)
		}
	}
	return false
}

func (g *WaitForGraph) AddNode(nodeID int) {
	if _, exists := g.nodes[nodeID]; !exists {
		newNode := &Node{
			OutNeighbors: []*NodeEdge{},
			InNeighbors:  []*NodeEdge{},
		}
		g.nodes[nodeID] = newNode
		fmt.Println("New node added to graph")
	} else {
		fmt.Println("Node already exists!")
	}
}

func (g *WaitForGraph) AddEdge(fromNodeID, toNodeID int, lock Lock) {
	node1 := g.nodes[fromNodeID]
	node2 := g.nodes[toNodeID]
	nodeEdge1 := &NodeEdge{
		node:           node1,
		associatedLock: lock,
	}
	nodeEdge2 := &NodeEdge{
		node:           node2,
		associatedLock: lock,
	}
	node1.OutNeighbors = append(node1.OutNeighbors, nodeEdge2)
	node2.InNeighbors = append(node2.InNeighbors, nodeEdge1)
	fmt.Println("New edge added to graph")
}

func (g *WaitForGraph) RemoveEdge(fromNode, toNode *Node) {
	index := -1
	for i, n := range fromNode.OutNeighbors {
		if n.node == toNode {
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
		if n.node == toNode {
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
