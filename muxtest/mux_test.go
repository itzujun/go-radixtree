package muxtest

import (
	"fmt"
	"github.com/itzujun/go-radixtree"
	"testing"
)

func TestMux(t *testing.T) {

	tree := goradixtree.NewRadixTree()

	tree.Insert(1, "/hello", func(name string) string {
			return "你好：" + name
		}).Insert(1, "/world", func(s string) string {
			return "欢迎来到：" + s
		})

	node := tree.FindNode("/hello")
	if node != nil {
		fmt.Println(node.Handler("马云"))
	}

	node = tree.FindNode("/world")
	if node != nil {
		fmt.Println(node.Handler("北京"))
	}

}
