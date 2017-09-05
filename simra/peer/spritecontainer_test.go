package peer

import (
	"testing"

	"golang.org/x/mobile/exp/sprite"
)

func TestGetSpriteContainer(t *testing.T) {
	sc := GetSpriteContainer()
	if sc == nil {
		t.Errorf("GetSpriteContainer returned nil. unexpected")
	}
}

type mockGLer struct {
	GLer
}

func (m *mockGLer) NewNode(fn arrangerFunc) *sprite.Node {
	return &sprite.Node{}
}

func (m *mockGLer) SetSubTex(n *sprite.Node, subTex *sprite.SubTex) {
	// nop
}

func TestAddSprite(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gler = &mockGLer{}

	err := sc.AddSprite(&Sprite{}, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}

	if len(sc.spriteNodePairs) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}

	err = sc.AddSprite(&Sprite{}, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	if len(sc.spriteNodePairs) != 2 {
		t.Errorf("unexpected result. [got] %d [want] %d", len(spriteContainer.spriteNodePairs), 0)
	}
}
