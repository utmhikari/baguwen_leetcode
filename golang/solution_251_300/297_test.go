package solution_251_300

import (
	"bytes"
	"fmt"
	"github.com/utmhikari/go-leetcode"
	"strconv"
	"strings"
)

//297. 二叉树的序列化与反序列化
//
//序列化是将一个数据结构或者对象转换为连续的比特位的操作，
//进而可以将转换后的数据存储在一个文件或者内存中，
//同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。
//
//请设计一个算法来实现二叉树的序列化与反序列化。
//这里不限定你的序列 / 反序列化算法执行逻辑，
//你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。
//
//提示: 输入输出格式与 LeetCode 目前使用的方式一致，
//详情请参阅LeetCode 序列化二叉树的格式。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */



type Codec struct {
	buf bytes.Buffer

	isGeneratingOrder bool

	preorderCounter map[int]int
	preorderRefCache map[*main.TreeNode]string

	cntSep string
	dataSep string
	orderSep string
}

func Constructor() Codec {
	return Codec{
		buf: bytes.Buffer{},

		isGeneratingOrder: false,

		preorderCounter: make(map[int]int),
		preorderRefCache: make(map[*main.TreeNode]string),

		cntSep: "#",
		dataSep: ",",
		orderSep: "|",
	}
}


func (c *Codec) reset() {
	c.buf.Reset()
	c.isGeneratingOrder = false
	c.preorderCounter = make(map[int]int)
	c.preorderRefCache = make(map[*main.TreeNode]string)
}


func (c *Codec) getValAndCnt(s string) (int, int) {
	splits := strings.Split(s, c.cntSep)
	if len(splits) != 2 {
		return 0, -1
	}

	val, vErr := strconv.Atoi(splits[0])
	if vErr != nil {
		return 0, -1
	}

	cnt, cErr := strconv.Atoi(splits[1])
	if cErr != nil {
		return 0, -1
	}

	return val, cnt
}

func (c *Codec) solve297(preOrder []string, inOrder []string, p1 int, p2 int, i1 int, i2 int) *main.TreeNode {
	if p2 < p1 {
		return nil
	}
	if p2 == p1 {
		v, _ := c.getValAndCnt(preOrder[p1])
		return &main.TreeNode{Val: v}
	}

	i := i1
	for i <= i2 {
		if inOrder[i] == preOrder[p1] {
			break
		}
		i++
	}

	v, _ := c.getValAndCnt(preOrder[p1])
	root := &main.TreeNode{Val: v}
	root.Left = c.solve297(preOrder, inOrder, p1 + 1, p1 + (i - i1), i1, i - 1)
	root.Right = c.solve297(preOrder, inOrder, p1 + (i + 1 - i1), p2, i + 1, i2)
	return root
}


func (c *Codec) buildTree297(preorder []string, inorder []string) *main.TreeNode {
	lp, li := len(preorder), len(inorder)
	if lp != li || lp == 0 {
		return nil
	}
	if lp == 1 {
		v, _ := c.getValAndCnt(preorder[0])
		return &main.TreeNode{Val: v}
	}
	return c.solve297(preorder, inorder, 0, lp - 1, 0, li - 1)
}


func (c *Codec) genDataSep() {
	if c.isGeneratingOrder {
		c.buf.WriteString(c.dataSep)
	}
	c.isGeneratingOrder = true
}

func (c *Codec) genDataValOfPreOrder(root *main.TreeNode) {
	if root == nil {
		panic("genDataValOfPreOrder: nil root")
	}

	val := root.Val
	var curCnt int
	cnt, ok := c.preorderCounter[val]
	if !ok {
		curCnt = 1
	} else {
		curCnt = cnt + 1
	}
	c.preorderCounter[val] = curCnt

	key := fmt.Sprintf("%d%s%d", val, c.cntSep, cnt)
	c.preorderRefCache[root] = key
	c.buf.WriteString(key)
}


func (c *Codec) genPreOrder(root *main.TreeNode) {
	if root == nil {
		return
	}

	c.genDataSep()
	c.genDataValOfPreOrder(root)
	c.genPreOrder(root.Left)
	c.genPreOrder(root.Right)
}

func (c *Codec) genDataValOfInOrder(root *main.TreeNode) {
	if root == nil {
		panic("genDataValOfInOrder: nil root")
	}

	key, ok := c.preorderRefCache[root]
	if !ok {
		panic(fmt.Sprintf("genDataValOfInOrder: cannot find cached ref, val: %d", root.Val))
	}

	c.buf.WriteString(key)
}

func (c *Codec) genInOrder(root *main.TreeNode) {
	if root == nil {
		return
	}

	c.genInOrder(root.Left)
	c.genDataSep()
	c.genDataValOfInOrder(root)
	c.genInOrder(root.Right)
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *main.TreeNode) string {
	c.reset()

	// gen preorder
	c.genPreOrder(root)

	// reset data
	c.preorderCounter = make(map[int]int)
	c.isGeneratingOrder = false

	// gen order separator
	c.buf.WriteString(c.orderSep)

	// gen inorder
	c.genInOrder(root)

	// reset data
	c.preorderRefCache = make(map[*main.TreeNode]string)

	// get serialized str
	s := c.buf.String()
	c.reset()

	fmt.Printf("serialized: %s\n", s)
	return s
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *main.TreeNode {
	c.reset()

	// validity check
	if len(data) == 0 || data == c.orderSep {
		return nil
	}

	orders := strings.Split(data, c.orderSep)
	if len(orders) != 2 {
		return nil
	}

	preorder, inorder := strings.Split(orders[0], c.dataSep), strings.Split(orders[1], c.dataSep)
	if len(preorder) != len(inorder) {
		return nil
	}

	fmt.Printf("preorder: %s, inorder: %s\n", preorder, inorder)
	return c.buildTree297(preorder, inorder)
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
