package goradixtree

import "fmt"


// 2019.03.10

type kind uint8


type children []*Node

type handlerFunc func(string) string

type Node struct {
	Kind     kind
	Lable    byte
	Preffix  string
	Children children
	ppath    string
	Handler  handlerFunc
}


func newNode(k kind, pre string, ppath string) *Node {
	return &Node{
		Kind:    k,
		Lable:   pre[0],
		Preffix: pre,
		ppath:   ppath,
	}
}


//添加子节点
func (node *Node) addChild(n *Node) {
	node.Children = append(node.Children, n)
}

func (node *Node) findChild(l byte, k kind) *Node {
	for _, c := range node.Children {
		if c.Lable == l && c.Kind == k {
			return c
		}
	}
	return nil
}


func (node *Node) findChildWithLabel(l byte) *Node {
	for _, c := range node.Children {
		if c.Lable == l {
			return c
		}
	}
	return nil
}

func (node *Node) findChildByKind(k kind) *Node {
	for _, c := range node.Children {
		if c.Kind == k {
			return c
		}
	}
	return nil
}


func (node *Node) addHandler(handler handlerFunc) {
	node.Handler = handler
}

func (node *Node) getHandler() handlerFunc {
	return node.Handler
}

type RadixTree struct {
	tree *Node
}

func (r *RadixTree) getNode() *Node {
	return r.tree
}

func newRadixTree() *RadixTree {
	return &RadixTree{
		tree: &Node{

		},
	}
}


func (r *RadixTree) insert(k kind, path string, handler handlerFunc) *RadixTree {
	if path == "" {
		panic("path cant be empty")
	}

	if path[0] != '/' {
		path = "/" + path
	}
	ppath := path
	cn := r.tree
	search := path
	for {

		sl := len(search)
		pl := len(cn.Preffix)
		max := pl
		l := 0

		if sl < max {
			max = sl
		}

		for ; l < max && search[l] == cn.Preffix[l]; l++ {
			// pass
		}

		if l == 0 { //root node
			cn.Lable = search[0]
			cn.Preffix = search
			if handler != nil {
				cn.Handler = handler
			}
			fmt.Println("ok")
		} else if l < pl {
			n := newNode(cn.Kind, cn.Preffix[:l], ppath)
			n.Children = cn.Children
			cn.Preffix = search[:l]
			cn.addChild(n)
			if l == sl {
				cn.Handler = handler
				cn.addHandler(handler)
			} else {
				n = newNode(k, search[:l], ppath)
				n.addHandler(handler)
				cn.addChild(n)
			}
		} else if l < sl {
			search = search[l:]
			c := cn.findChildWithLabel(search[0])
			if c != nil {
				cn = c
				continue
			}
			n := newNode(k, search, ppath)
			n.addHandler(handler)
			cn.addChild(n)
		} else {
			if handler != nil {
				cn.addHandler(handler)
				cn.ppath = path
			}
		}
		break

	}
	return r

}

func (r *RadixTree) findNode(path string) *Node {
	search := path
	cn := r.tree
	for {
		if search == "" {
			fmt.Println("path is nil")
			return nil
		}

		l := 0
		sl := len(search)
		pl := len(cn.Preffix)

		max := pl

		if sl < max {
			max = sl
		}

		for ; l < max && search[l] == cn.Preffix[l]; l++ {

		}

		if l == pl {
			search = search[l:]
		}

		if search == "" {
			break
		}
		if child := cn.findChildWithLabel(search[0]); child != nil {
			cn = child
			continue
		} else {
			fmt.Println("can not find it")
			return nil
		}
	}

	return cn

}










