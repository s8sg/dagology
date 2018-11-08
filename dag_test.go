package dagology

import (
	"testing"
)

func TestMakeDotFile(t *testing.T) {
	d := New()
	d.AddVertex("s0", 0)
	d.AddVertex("s1", 1)
	d.AddVertex("s2", 2)
	d.AddVertex("s3", 3)
	d.AddVertex("s4", 3)

	if d.AddEdge("s0", "s1") != nil {
		t.Fatalf("Failed")
	}
	if d.AddEdge("s1", "s2") != nil {
		t.Fatalf("Failed")
	}
	if d.AddEdge("s0", "s2") != nil {
		t.Fatalf("Failed")
	}
	if d.AddEdge("s2", "s4") != nil {
		t.Fatalf("Failed")
	}
	if d.AddEdge("s0", "s3") != nil {
		t.Fatalf("Failed")
	}
	if d.AddEdge("s3", "s4") != nil {
		t.Fatalf("Failed")
	}
	//	if d.AddEdge("s0", "s4") != nil {
	//	t.Fatalf("Failed")
	//      }
	if d.AddEdge("s4", "s0") == nil {
		t.Fatalf("Failed")
	}

	d.MakeDotGraph("test.dot")
	t.Logf("dot -o test.png -T png test.dot")
}
