package peer

import (
	"fmt"
	"image"
	"log"
	"time"

	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/exp/app/debug"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
	"golang.org/x/mobile/gl"
)

var self *Peer

var startTime = time.Now()

type TouchListener interface {
	OnTouch(x, y float32)
}

type PeerSprite struct {
	TouchListener
	W float32
	H float32
	X float32
	Y float32
	R float32
}

type peerSpriteContainer struct {
	peerSprite *PeerSprite
	node       *sprite.Node
}

type screenSize struct {
	width  float32
	height float32
	scale  float32
}

type Peer struct {
	glctx                gl.Context
	images               *glutil.Images
	fps                  *debug.FPS
	eng                  sprite.Engine
	scene                *sprite.Node
	sz                   size.Event
	peerSpriteContainers []*peerSpriteContainer
	touchListeners       []*TouchListener
	desiredScreenSize    screenSize
}

func GetInstance() *Peer {
	if self == nil {
		self = &Peer{}
	}
	return self
}

func (self *Peer) Initialize(in_glctx gl.Context) {
	self.glctx = in_glctx

	// transparency of png
	self.glctx.Enable(gl.BLEND)
	self.glctx.BlendEquation(gl.FUNC_ADD)
	self.glctx.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	onStart(self.glctx)
}

func onStart(glctx gl.Context) {
	self.images = glutil.NewImages(glctx)
	self.fps = debug.NewFPS(self.images)
	loadScene()
}

func loadScene() {
	if self.eng != nil {
		self.eng.Release()
	}
	self.eng = glsprite.Engine(self.images)
	self.scene = &sprite.Node{}
	self.eng.Register(self.scene)
	self.eng.SetTransform(self.scene, f32.Affine{
		{1, 0, 0},
		{0, 1, 0},
	})
}

func newNode() *sprite.Node {
	n := &sprite.Node{}
	self.eng.Register(n)
	self.scene.AppendChild(n)
	return n
}

func (self *Peer) calcScale() {
	h := self.desiredScreenSize.height
	w := self.desiredScreenSize.width

	if h/float32(self.sz.HeightPt) > w/float32(self.sz.WidthPt) {
		self.desiredScreenSize.scale = float32(self.sz.HeightPt) / h
		fmt.Println("scale = ", self.desiredScreenSize.scale)
	} else {
		self.desiredScreenSize.scale = float32(self.sz.WidthPt) / w
		fmt.Println("scale = ", self.desiredScreenSize.scale)
	}
}

func (self *Peer) SetDesiredScreenSize(w, h float32) {
	self.desiredScreenSize.height = h
	self.desiredScreenSize.width = w
	self.calcScale()
}

func (self *Peer) LoadTexture(assetName string, rect image.Rectangle) sprite.SubTex {

	a, err := asset.Open(assetName)
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	img, _, err := image.Decode(a)
	if err != nil {
		log.Fatal(err)
	}
	t, err := self.eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}

	return sprite.SubTex{t, rect}
}

func (self *Peer) SetScreenSize(in_sz size.Event) {
	self.sz = in_sz
	fmt.Println(in_sz)
	self.calcScale()
}

func (self *Peer) GetScreenSize() size.Event {
	return self.sz
}

func (self *Peer) Finalize() {
	self.eng.Release()
	self.fps.Release()
	self.images.Release()
	self.glctx = nil
}

func (self *Peer) Update() {
	if self.glctx == nil {
		return
	}
	self.glctx.ClearColor(1, 1, 1, 1) // white background
	self.glctx.Clear(gl.COLOR_BUFFER_BIT)
	now := clock.Time(time.Since(startTime) * 60 / time.Second)

	self.apply()

	self.eng.Render(self.scene, now, self.sz)
	self.fps.Draw(self.sz)
}

func (self *Peer) AddSprite(ps *PeerSprite, subTex sprite.SubTex) {
	fmt.Println("[IN] Peer.AddSprite()")
	var psc peerSpriteContainer
	psc.peerSprite = ps
	psc.node = newNode()
	self.peerSpriteContainers = append(self.peerSpriteContainers, &psc)
	self.eng.SetSubTex(psc.node, subTex)
}

func (self *Peer) Reset() {
	self.peerSpriteContainers = nil
	self.touchListeners = nil
	loadScene()
}

func (self *Peer) apply() {

	peerSpriteContainers := self.peerSpriteContainers

	for i := range peerSpriteContainers {
		psc := peerSpriteContainers[i]
		if psc.peerSprite == nil {
			continue
		}

		affine := &f32.Affine{
			{1, 0, 0},
			{0, 1, 0},
		}
		affine.Translate(affine,
			psc.peerSprite.X*self.desiredScreenSize.scale-psc.peerSprite.W/2*self.desiredScreenSize.scale,
			psc.peerSprite.Y*self.desiredScreenSize.scale-psc.peerSprite.H/2*self.desiredScreenSize.scale)
		//fmt.Println("x, y = ", psc.peerSprite.X, psc.peerSprite.Y)
		if psc.peerSprite.R != 0 {
			affine.Translate(affine, 0.5, 0.5)
			affine.Rotate(affine, psc.peerSprite.R)
			affine.Translate(affine, -0.5, -0.5)
		}
		//affine.Scale(affine, self.desiredScreenSize.width*self.desiredScreenSize.scale,
		//	self.desiredScreenSize.height*self.desiredScreenSize.scale)
		affine.Scale(affine, psc.peerSprite.W*self.desiredScreenSize.scale,
			psc.peerSprite.H*self.desiredScreenSize.scale)
		self.eng.SetTransform(psc.node, *affine)
	}
}

func (self *Peer) AddTouchListener(listener TouchListener) {
	self.touchListeners = append(self.touchListeners, &listener)
}

func (self *Peer) OnTouch(x, y float32) {
	for i := range self.touchListeners {
		listener := self.touchListeners[i]
		if listener == nil {
			fmt.Println("listener is nil!")
			continue
		}

		(*listener).OnTouch(x*self.desiredScreenSize.scale, y*self.desiredScreenSize.scale)
	}
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }
