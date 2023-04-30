package solution_101_150

import (
	"fmt"
	"github.com/utmhikari/go-leetcode"
)

//133. 克隆图
//给你无向连通图中一个节点的引用，请你返回该图的深拷贝（克隆）。
//
//图中的每个节点都包含它的值 val（int） 和其邻居的列表（list[Node]）。
//
//class Node {
//public int val;
//public List<Node> neighbors;
//}


var nodesMap133 map[int]*main.Node
var nodes133 map[int]*main.Node
var startNode *main.Node = nil

func getNeighbors133(node *main.Node) []int {
	ret := make([]int, 0)
	if node == nil {
		return ret
	}
	for _, n := range node.Neighbors {
		ret = append(ret, n.Val)
	}
	return ret
}

func visit133(node *main.Node) {
	_, ok := nodesMap133[node.Val]
	if ok {
		return
	}

	fmt.Printf("%d, %v+\n", node.Val, getNeighbors133(node))

	nodesMap133[node.Val] = node
	newNode := &main.Node{Val: node.Val}
	if startNode == nil {
		startNode = newNode
	}
	nodes133[node.Val] = newNode
	for _, neighbor := range node.Neighbors {
		visit133(neighbor)
	}
}

func cloneGraph(node *main.Node) *main.Node {
	if node == nil {
		return nil
	}
	nodesMap133 = make(map[int]*main.Node)
	nodes133 = make(map[int]*main.Node)
	startNode = nil
	visit133(node)

	for i, n := range nodes133 {
		oldNode, ok := nodesMap133[i]
		n.Neighbors = make([]*main.Node, 0)
		if ok {
			for _, neighbor := range oldNode.Neighbors {
				neighborNode, nOk := nodes133[neighbor.Val]
				if nOk {
					n.Neighbors = append(n.Neighbors, neighborNode)
				}
			}

			fmt.Printf("%d, %d, %v+\n", i, n.Val, getNeighbors133(n))
		}
	}

	return startNode
}
