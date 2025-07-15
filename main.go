package main

import (
	"bytes"
	"image/color"
	"log"

	"golang.org/x/image/font/gofont/goregular"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type game struct {
	ui *ebitenui.UI
}

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten UI canvas scaling")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	rootContainer := widget.NewContainer()

	ui := &ebitenui.UI{
		Container: rootContainer,
	}

	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	var fontFace text.Face = &text.GoTextFace{
		Source: s,
		Size:   32,
	}

	helloWorldLabel := widget.NewText(
		widget.TextOpts.Text("Hello World!", fontFace, color.White),
	)

	rootContainer.AddChild(helloWorldLabel)

	game := game{
		ui: ui,
	}

	err = ebiten.RunGame(&game)
	if err != nil {
		log.Print(err)
	}
}

func (g *game) Update() error {
	g.ui.Update()

	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)
}

func (g *game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
