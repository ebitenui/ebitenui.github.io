// START root
package lay_gri_pos_staxsta

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
	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(assets.ColorL),
		),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(5),
			widget.GridLayoutOpts.DefaultStretch(true, true),
		)),
	)
	colors := []color.Color{
		assets.ColorR,
		assets.ColorY,
		assets.ColorB,
		assets.ColorG,
		assets.ColorD,
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			// START this
			child := widget.NewContainer(
				widget.ContainerOpts.BackgroundImage(
					image.NewNineSliceColor(colors[(x+y)%len(colors)]),
				),
				widget.ContainerOpts.WidgetOpts(
					widget.WidgetOpts.LayoutData(widget.GridLayoutData{
						MaxWidth:           32,
						MaxHeight:          32,
						HorizontalPosition: widget.GridLayoutPositionStart,
						VerticalPosition:   widget.GridLayoutPositionStart,
					}),
					widget.WidgetOpts.MinSize(64, 64),
				),
				widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
			)
			// END this
			root.AddChild(child)
		}
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
