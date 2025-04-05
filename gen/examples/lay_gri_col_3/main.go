// START root
package lay_gri_col_3

import (
	"gen/assets"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ui *ebitenui.UI
}

func NewGame() *Game {
	// START this
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(assets.ColorL),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(3),
		)),
	)
	// END this
	colors := []color.Color{
		assets.ColorR,
		assets.ColorY,
		assets.ColorB,
		assets.ColorG,
		assets.ColorD,
	}
	for i := 0; i < 5; i++ {
		child := widget.NewContainer(
			widget.ContainerOpts.BackgroundImage(
				image.NewNineSliceColor(colors[i%len(colors)]),
			),
			widget.ContainerOpts.WidgetOpts(
				widget.WidgetOpts.LayoutData(widget.GridLayoutData{}),
				widget.WidgetOpts.MinSize(64, 64),
			),
			widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		)
		root.AddChild(child)
	}

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

// END root
