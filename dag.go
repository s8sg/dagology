// Directed Acyclic Graph implementation in golang.
package dagology

import (
	"fmt"
	"github.com/funkygao/golib/str"
	"os"
)

var (
	ErrCyclic = fmt.Errorf("dag has cyclic dependency")
)

type Dag struct {
	nodes map[string]*Node
}

type Node struct {
	id  string
	val interface{}

	outdegree int
	indegree  int
	children  []*Node
	next      []*Node
	prev      []*Node
}

/*
{
   “order” :  {
     "0" : [
        “s0”
     ],
     "1": [
        {
        “order” :  {
           "0":  [
                “s1”
           ],
           "1": [
                “s2”
           ]
        }},
        “s3”
    ],
    "2": [
        “s4”
    ]
   }
}
*/
func New() *Dag {
	this := new(Dag)
	this.nodes = make(map[string]*Node)
	return this
}

func (this *Dag) AddVertex(id string, val interface{}) *Node {
	node := &Node{id: id, val: val}
	this.nodes[id] = node
	return node
}

func (this *Node) inSlice(list []*Node) bool {
	for _, b := range list {
		if b.id == this.id {
			return true
		}
	}
	return false
}

func (this *Dag) AddEdge(from, to string) error {
	fromNode := this.nodes[from]
	toNode := this.nodes[to]

	// Check if cyclic dependency
	if fromNode.inSlice(toNode.next) || toNode.inSlice(fromNode.prev) {
		return ErrCyclic
	}

	// Update next of fromnode and all prev nodes of fromnode
	fromNode.next = append(fromNode.next, toNode)
	fromNode.next = append(fromNode.next, toNode.next...)
	for _, b := range fromNode.prev {
		b.next = append(b.next, toNode)
		b.next = append(b.next, toNode.next...)
	}

	// Update prev of fromnode and all next nodes of fromnode
	toNode.prev = append(toNode.prev, fromNode)
	toNode.prev = append(toNode.prev, fromNode.prev...)
	for _, b := range toNode.next {
		b.prev = append(b.prev, fromNode)
		b.prev = append(b.prev, fromNode.prev...)
	}

	fromNode.children = append(fromNode.children, toNode)
	fromNode.outdegree++
	toNode.indegree++

	return nil
}

func (this *Dag) Node(id string) *Node {
	return this.nodes[id]
}

func findInBase(base [][]string, id string) int {
	for pos, list := range base {
		if list != nil {
			for _, b := range list {
				if b == id {
					return pos
				}
			}
		}
	}
	return -1
}

func (this *Dag) GetDependencyOrder() *DependencyOrder {
	var start string
	var end string

	base = make([][]string, len(list))

	/*
		visitReferenceMap := make(map[string]bool)
		// Find the start and end node and initialize the visit reference map
			for _, b := range list {
			if b.indegree == 0 {
				start = b.id
			}
			if b.outdegree == 0 {
				end = b.id
			}
			visitReferenceMap[b.name] = false
		}
	*/

	pos := 0
	for _, b := range list {
		if len(b.prev) == 0 {
			continue
		}
		// Check if b is added in base
		bp := findInBase(base, b)
		if bp == -1 {
			// add b into base
			base[0] = append(base[0], b)
			bp = 0
		}
		for _, d := range b.prev {
			dp := findInBase(base, d)
			if dp == nil {
				base[bp+1] = append(base[bp+1], d]
			} else {
				delete(base[dp], d)
				base[bp+1] = append(base[bp+1], d]
			}
		}
	}
}

func (this *Dag) MakeDotGraph(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sb := str.NewStringBuilder()
	sb.WriteString("digraph depgraph {\n\trankdir=LR;\n")
	for _, node := range this.nodes {
		node.dotGraph(sb)
	}
	sb.WriteString("}\n")
	file.WriteString(sb.String())
	return sb.String()
}

func (this *Node) dotGraph(sb *str.StringBuilder) {
	if len(this.children) == 0 {
		sb.WriteString(fmt.Sprintf("\t\"%s\";\n", this.id))
		return
	}

	for _, child := range this.children {
		sb.WriteString(fmt.Sprintf(`%s -> %s [label="%v"]`, this.id, child.id, this.val))
		sb.WriteString("\r\n")
	}
}

func (this *Node) Children() []*Node {
	return this.children
}
