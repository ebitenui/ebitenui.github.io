package assets

import (
	"embed"
	"fmt"
	"image"
	"io/fs"
	"regexp"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

//go:embed data
var data embed.FS

var (
	Frame    = load(data, "data/frame.png")
	Output   = "assets/examples"
	Parent   = "gen"
	Width    = 480
	Height   = 320
	Pivot    = 28.
	ColorR   = colornames.Indianred
	ColorY   = colornames.Goldenrod
	ColorB   = colornames.Steelblue
	ColorG   = colornames.Mediumseagreen
	ColorD   = colornames.Darkslategray
	ColorL   = colornames.Gainsboro
	ColorN   = colornames.Black
	ColorMap = map[string]string{
		"assets.ColorR": "colornames.Indianred",
		"assets.ColorY": "colornames.Goldenrod",
		"assets.ColorB": "colornames.Steelblue",
		"assets.ColorG": "colornames.Mediumseagreen",
		"assets.ColorD": "colornames.Darkslategray",
		"assets.ColorL": "colornames.Gainsboro",
		"assets.ColorN": "colornames.Black",
	}
	ColorPackage         = `"golang.org/x/image/colornames"`
	RegPackage           = regexp.MustCompile(`package .+`)
	RegAssets            = regexp.MustCompile(`(?m)^\s*"\S*assets\S*"\s*\n`)
	RegImports           = regexp.MustCompile(`import \(([\s\S]*?)\n\)`)
	RegNewline           = regexp.MustCompile(`\n+`)
	RegColor             = regexp.MustCompile(`\b(assets.Color[RYBGDLN])\b`)
	RegComment           = regexp.MustCompile(`//.*?`)
	RegWhitespaceOnly    = regexp.MustCompile("(?m)^[ \t]+$")
	RegLeadingWhitespace = regexp.MustCompile("(?m)(^[ \t]*)(?:[^ \t\n])")
	PartMain             = dedent(fmt.Sprintf(`
		func main() {
			ebiten.SetWindowSize(%d, %d)
			ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
			if err := ebiten.RunGame(NewGame()); err != nil {
				panic(err)
			}
		}
	`, Width, Height))
)

func load(fs fs.FS, path string) *ebiten.Image {
	f, err := fs.Open(path)
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func dedent(text string) string {
	var margin string
	text = RegWhitespaceOnly.ReplaceAllString(text, "")
	indents := RegLeadingWhitespace.FindAllStringSubmatch(text, -1)
	for i, indent := range indents {
		if i == 0 {
			margin = indent[1]
		} else if strings.HasPrefix(indent[1], margin) {
			continue
		} else if strings.HasPrefix(margin, indent[1]) {
			margin = indent[1]
		} else {
			margin = ""
			break
		}
	}
	if margin != "" {
		text = regexp.MustCompile("(?m)^"+margin).ReplaceAllString(text, "")
	}
	text = strings.TrimPrefix(text, "\n")
	return text
}
