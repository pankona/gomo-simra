package peer

import (
	"sync"
	"testing"
	"time"

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

func (m *mockGLer) AppendNode(n *sprite.Node) {
	// nop
}

func (m *mockGLer) RemoveNode(n *sprite.Node) {
	// nop
}

func mapLen(m *sync.Map) int {
	count := 0
	m.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

func TestAddAndRemoveSprite(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	s1 := &Sprite{}
	s2 := &Sprite{}
	_ = sc.AddSprite(s1, nil, nil)
	_ = sc.AddSprite(s2, nil, nil)
	m := sc.GetSpriteNodePairs()
	if mapLen(m) != 2 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 2)
	}

	// RemoveSprite marks sprites as "not in use",
	// length of spriteContainer will not be changed
	sc.RemoveSprite(s1)
	sc.RemoveSprite(s2)
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}

	// if there're "not in use" sprite in spriteContainer,
	// AddSprite will reuse them. length of spriteContainer will not be changed until the number of
	// sprites don't reach to its capacity.
	_ = sc.AddSprite(s1, nil, nil)
	_ = sc.AddSprite(s2, nil, nil)
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 2 {
		t.Fatalf("unexpected result. [got] %d [want] %d", mapLen(m), 2)
	}

	// if there's not "not in use" sprite in spriteContainer,
	// length of spriteContainer will be extended.
	s3 := &Sprite{}
	_ = sc.AddSprite(s3, nil, nil)
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 3 {
		t.Fatalf("unexpected result. [got] %d [want] %d", mapLen(m), 3)
	}
}

func TestAddSpriteDuplicate(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	s1 := &Sprite{}
	err := sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	m := sc.GetSpriteNodePairs()
	if mapLen(m) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}

	// if specified sprite is already added, it will be ignored.
	err = sc.AddSprite(s1, nil, nil)
	if err == nil {
		t.Fatalf("unexpected behaviour. AddSprite should return error for duplicated adding")
	}
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}
}

func TestRemoveSpriteDuplicate(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	s1 := &Sprite{}
	err := sc.AddSprite(s1, nil, nil)
	if err != nil {
		t.Fatalf("failed add Sprite. err: %s", err.Error())
	}
	m := sc.GetSpriteNodePairs()
	if mapLen(m) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}

	sc.RemoveSprite(s1)
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}

	// if specified sprite is already removed, it will be ignored.
	sc.RemoveSprite(s1)
	m = sc.GetSpriteNodePairs()
	if mapLen(m) != 1 {
		t.Errorf("unexpected result. [got] %d [want] %d", mapLen(m), 0)
	}
}

func TestRemoveSprites(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	for i := 0; i < 10; i++ {
		err := sc.AddSprite(&Sprite{}, nil, nil)
		if err != nil {
			t.Fatalf(err.Error())
		}
	}
	if mapLen(&sc.spriteNodePairs) != 10 {
		t.Fatalf("unexpected map length")
	}

	sc.RemoveSprites()
	if mapLen(&sc.spriteNodePairs) != 0 {
		t.Fatalf("unexpected map length")
	}
}

func TestReplaceTexture(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	s := &Sprite{}
	err := sc.AddSprite(s, nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	tex := &Texture{}
	sc.ReplaceTexture(s, tex)
}

type listener struct {
	touchBegin func(x, y float32)
	touchMove  func(x, y float32)
	touchEnd   func(x, y float32)
}

func (l *listener) OnTouchBegin(x, y float32) {
	l.touchBegin(x, y)
}
func (l *listener) OnTouchMove(x, y float32) {
	l.touchMove(x, y)
}
func (l *listener) OnTouchEnd(x, y float32) {
	l.touchEnd(x, y)
}

func TestTouchEvent(t *testing.T) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}

	s := &Sprite{}
	err := sc.AddSprite(s, nil, nil)
	if err != nil {
		t.Fatalf(err.Error())
	}
	var wg sync.WaitGroup
	s.AddTouchListener(&listener{
		touchBegin: func(x, y float32) {
			wg.Done()
		},
		touchMove: func(x, y float32) {
			wg.Done()
		},
		touchEnd: func(x, y float32) {
			wg.Done()
		},
	})
	wg.Add(3)
	sc.OnTouchBegin(0, 0)
	sc.OnTouchMove(0, 0)
	sc.OnTouchEnd(0, 0)

	waitChan := make(chan struct{})
	go func() {
		wg.Wait()
		waitChan <- struct{}{}
	}()

	select {
	case <-waitChan:
		// success. nop.
	case <-time.After(3 * time.Second):
		t.Error("touch event didn't fired as expected")
	}
}

func BenchmarkAddSprite(b *testing.B) {
	sc := &SpriteContainer{}
	sc.gl = &mockGLer{}
	for i := 0; i < b.N; i++ {
		err := sc.AddSprite(&Sprite{}, nil, nil)
		if err != nil {
			b.Fatalf(err.Error())
		}
	}
}
