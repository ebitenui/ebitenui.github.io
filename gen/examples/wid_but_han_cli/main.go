// START root
package wid_but_han_cli

import (
	"bytes"
	"fmt"
	"gen/assets"
	"image/color"
	"math"
	"strings"

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
			Pressed: Mix(assets.ColorL, assets.ColorN, 0.4),
		}),
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle:         DefaultNineSlice(assets.ColorD),
			Hover:        DefaultNineSlice(Mix(assets.ColorD, assets.ColorG, 0.4)),
			Disabled:     DefaultNineSlice(Mix(assets.ColorD, assets.ColorL, 0.8)),
			Pressed:      PressedNineSlice(Mix(assets.ColorD, assets.ColorN, 0.4)),
			PressedHover: PressedNineSlice(Mix(assets.ColorD, assets.ColorN, 0.4)),
		}),
		// SKIP-END this
		widget.ButtonOpts.ClickedHandler(
			func(args *widget.ButtonClickedEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "ButtonClicked\n")
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.PressedHandler(
			func(args *widget.ButtonPressedEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "Button Pressed\n")
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.ReleasedHandler(
			func(args *widget.ButtonReleasedEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "Button Released\n")
				fmt.Fprintf(&b, "> Inside: %t\n", args.Inside)
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.CursorEnteredHandler(
			func(args *widget.ButtonHoverEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "Cursor Entered\n")
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.CursorMovedHandler(
			func(args *widget.ButtonHoverEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "Cursor Moved\n")
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.CursorExitedHandler(
			func(args *widget.ButtonHoverEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "Cursor Exited\n")
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
		widget.ButtonOpts.StateChangedHandler(
			func(args *widget.ButtonChangedEventArgs) {
				var b strings.Builder
				fmt.Fprintf(&b, "State Changed\n")
				fmt.Fprintf(&b, "> State: %v\n", args.State)
				fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
				fmt.Print(b)
			},
		),
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

func DefaultNineSlice(base color.Color) *image.NineSlice {
	var size float32 = 64
	var tiles float32 = 16
	var radius float32 = 8

	tile := size / tiles
	facet := Mix(base, assets.ColorL, 0.2)

	img := ebiten.NewImage(int(size), int(size))
	path := RoundedRectPath(0, tile, size, size-tile, radius, radius, radius, radius)
	vector.DrawFilledPath(img, path, base, true, vector.FillRuleEvenOdd)
	path = RoundedRectPath(0, tile, size, size-tile*2, radius, radius, radius, radius)
	vector.DrawFilledPath(img, path, facet, true, vector.FillRuleEvenOdd)
	path = RoundedRectPath(tile, tile*2, size-tile*2, size-tile*4, radius, radius, radius, radius)
	vector.DrawFilledPath(img, path, base, true, vector.FillRuleEvenOdd)

	return image.NewNineSliceBorder(img, int(tile*4))
}

func PressedNineSlice(base color.Color) *image.NineSlice {
	var size float32 = 64
	var tiles float32 = 16
	var radius float32 = 8

	tile := size / tiles
	facet := Mix(base, assets.ColorL, 0.2)

	img := ebiten.NewImage(int(size), int(size))
	path := RoundedRectPath(0, 0, size, size, radius, radius, radius, radius)
	vector.DrawFilledPath(img, path, facet, true, vector.FillRuleEvenOdd)
	path = RoundedRectPath(tile, tile, size-tile*2, size-tile*2, radius, radius, radius, radius)
	vector.DrawFilledPath(img, path, base, true, vector.FillRuleEvenOdd)

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

func RoundedRectPath(x, y, w, h, tl, tr, br, bl float32) *vector.Path {
	path := &vector.Path{}

	path.Arc(x+w-tr, y+tr, tr, 3*math.Pi/2, 0, vector.Clockwise)
	path.LineTo(x+w, y+h-br)
	path.Arc(x+w-br, y+h-br, br, 0, math.Pi/2, vector.Clockwise)
	path.LineTo(x+bl, y+h)
	path.Arc(x+bl, y+h-bl, bl, math.Pi/2, math.Pi, vector.Clockwise)
	path.LineTo(x, y+tl)
	path.Arc(x+tl, y+tl, tl, math.Pi, 3*math.Pi/2, vector.Clockwise)
	path.Close()

	return path
}

// END root
