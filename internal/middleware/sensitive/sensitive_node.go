package sensitive

import (
	"bufio"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/text/gregex"
	"os"
	"strings"
	"sync"
)

const (
	DEFAULT_INITIAL_CAPACITY = 131072
)

var (
	instance    *SensitiveNode
	stringPoint *StringPointer
	once        sync.Once
	nodes       []*SensitiveNode
)

// 词库数据
type SensitiveNode struct {
	HeadTwoCharMix int
	Words          *garray.Array
	Next           *SensitiveNode
}

// 词库数据node构造
func (this *SensitiveNode) SensitiveNode(headTwoCharMix int) *SensitiveNode {
	obj := &SensitiveNode{
		HeadTwoCharMix: headTwoCharMix,
		Words:          garray.New(true),
	}
	return obj
}

// 词库数据node构造
func (this *SensitiveNode) SensitiveNodeByParent(headTwoCharMix int, parent *SensitiveNode) *SensitiveNode {
	obj := &SensitiveNode{
		HeadTwoCharMix: headTwoCharMix,
		Words:          garray.New(true),
	}
	parent.Next = obj
	return obj
}

// 装载词库单例
func (this *SensitiveNode) GetInstance() []*SensitiveNode {
	nodesList := make([]*SensitiveNode, DEFAULT_INITIAL_CAPACITY)
	once.Do(func() {
		file, err := os.Open("resource/sensitive_words_lines.txt")
		Check(err)
		defer file.Close()
		reader := bufio.NewReader(file)
		line, _, err := reader.ReadLine()
		Check(err)
		this.put(string(line), nodesList)
	})
	return nodesList
}

// 词库装入数组
func (this *SensitiveNode) put(word string, nodes []*SensitiveNode) bool {
	// 长度小于2的不加入
	if word == "" || len(strings.TrimSpace(word)) < 2 {
		return false
	}
	// 两个字符的不考虑
	if len(word) == 2 && gregex.IsMatchString("\\w\\w", word) {
		return false
	}
	point := stringPoint.StringPointer(strings.TrimSpace(word))
	// 计算头两个字符的hash
	hash := point.nextTwoCharHash(0)
	// 计算头两个字符的mix表示（mix相同，两个字符相同）
	mix := point.nextTwoCharMix(0)
	// 转为在hash桶中的位置
	index := int(hash) & (len(nodes) - 1)
	// 从桶里拿第一个节点
	node := nodes[index]
	if node == nil {
		// 如果没有节点，则生成一个节点数据
		node = this.SensitiveNode(int(mix))
		//添加词
		node.Words.Append(point)
		//放在桶里
		nodes[index] = node
	} else {
		// 如果已经有节点（1个或多个），找到正确的节点
		for ; node != nil; node = node.Next {
			// 匹配节点
			if node.HeadTwoCharMix == int(mix) {
				node.Words.Append(point)
				return true
			}
			// 如果匹配到最后仍然不成功，则追加一个节点
			if node.Next == nil {
				this.SensitiveNodeByParent(int(mix), node).Words.Append(point)
				return true
			}
		}
	}
	get, _ := nodes[index].Words.Get(0)
	fmt.Printf("值是：%c", get, "原值:", get)
	fmt.Printf("值是：%c %c", 25098, 36793)
	return true
}
