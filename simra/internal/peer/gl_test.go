package peer

import "testing"

func TestAppendNode(t *testing.T) {
	glpeer := &GLPeer{
		znodes: make([]*ZNode, 0),
	}
	zn1 := &ZNode{}
	zn2 := &ZNode{}
	zn3 := &ZNode{}
	glpeer.AppendNode(zn1)
	glpeer.AppendNode(zn2)
	glpeer.AppendNode(zn3)
	if len(glpeer.znodes) != 3 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(glpeer.znodes), 3)
	}
}

func TestRemoveNode(t *testing.T) {
	glpeer := &GLPeer{
		znodes: make([]*ZNode, 0),
	}
	zn1 := &ZNode{}
	zn2 := &ZNode{}
	zn3 := &ZNode{}
	glpeer.AppendNode(zn1)
	glpeer.AppendNode(zn2)
	glpeer.AppendNode(zn3)

	if len(glpeer.znodes) != 3 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(glpeer.znodes), 3)
	}

	glpeer.RemoveNode(zn1)
	glpeer.RemoveNode(zn2)
	glpeer.RemoveNode(zn3)

	if len(glpeer.znodes) != 0 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(glpeer.znodes), 0)
	}
}
