// START root
package wid_but_img_gemxidl

import (
	"bytes"
	"gen/assets"
	"image/color"
	"math"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

type Game struct {
	ui *ebitenui.UI
}

func NewGame() *Game {
	// START this
	button := widget.NewButton(
		// SKIP-START this
		widget.ButtonOpts.TextLabel("Button"),
		widget.ButtonOpts.TextFace(DefaultFont()),
		widget.ButtonOpts.TextColor(&widget.ButtonTextColor{
			Idle:    assets.ColorL,
			Hover:   assets.ColorL,
			Pressed: assets.ColorL,
		}),
		// SKIP-END this
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:    DefaultGem(assets.ColorD),
			Hover:   DefaultGem(Mix(assets.ColorD, assets.ColorG, 0.4)),
			Pressed: DefaultGem(Mix(assets.ColorD, assets.ColorN, 0.4)),
		}),
		// SKIP-START this
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
			widget.WidgetOpts.MinSize(180, 48),
		),
		// SKIP-END this
	)
	// END this
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(assets.ColorL),
		),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
	root.AddChild(button)
	return &Game{
		ui: &ebitenui.UI{Container: root},
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *Game) Layout(w, h int) (int, int) {
	return w, h
}

func DefaultFont() text.Face {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		panic(err)
	}
	return &text.GoTextFace{
		Source: s,
		Size:   20,
	}
}

func DefaultGem(base color.Color) *image.NineSlice {
	var size float32 = 64
	var tiles float32 = 16
	var padding float32 = 4

	tile := (size - padding*2) / tiles
	fore := Mix(base, assets.ColorL, 0.4)
	back := Mix(base, assets.ColorN, 0.4)
	stroke := &vector.StrokeOptions{Width: 4, LineCap: vector.LineCapRound}

	img := ebiten.NewImage(int(size), int(size))
	top := func(path *vector.Path) {
		path.MoveTo(padding, padding+size-tile*2)
		path.LineTo(padding, padding+tile*2)
		path.LineTo(padding+tile*2, padding)
		path.LineTo(padding+size-tile*2, padding)
	}
	bottom := func(path *vector.Path) {
		path.MoveTo(padding+size-tile*2, padding)
		path.LineTo(padding+size-tile*2, padding+size-tile*4)
		path.LineTo(padding+size-tile*4, padding+size-tile*2)
		path.LineTo(padding, padding+size-tile*2)
	}
	path := &vector.Path{}
	top(path)
	bottom(path)
	vector.DrawFilledPath(img, path, base, true, vector.FillRuleEvenOdd)
	path.Reset()
	top(path)
	vector.StrokePath(img, path, fore, true, stroke)
	path.Reset()
	bottom(path)
	vector.StrokePath(img, path, back, true, stroke)

	return image.NewNineSliceBorder(img, int(tile*4))
}

func Mix(a, b color.Color, percent float64) color.Color {
	rgba := func(c color.Color) (r, g, b, a uint8) {
		r16, g16, b16, a16 := c.RGBA()
		return uint8(r16 >> 8), uint8(g16 >> 8), uint8(b16 >> 8), uint8(a16 >> 8)
	}
	lerp := func(x, y uint8) uint8 {
		return uint8(math.Round(float64(x) + percent*(float64(y)-float64(x))))
	}
	r1, g1, b1, a1 := rgba(a)
	r2, g2, b2, a2 := rgba(b)

	return color.RGBA{
		R: lerp(r1, r2),
		G: lerp(g1, g2),
		B: lerp(b1, b2),
		A: lerp(a1, a2),
	}
}

// END root
