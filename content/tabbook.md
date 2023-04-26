---
layout: docs
title: TabBooks
navbarsection: container
---

How to Use [TabBooks](https://pkg.go.dev/github.com/ebitenui/ebitenui/widget#TabBook)
====================

The Tabbook container is used to create a stack of tabs that can be switched between by the use

Tabs
------
A Tab is a single container in the stack of a tabbook. It is an extension of a basic container with the addition of a label.
This label is what is displayed in the tab. By default the tab container is set to the same size as the parent TabBook.

~~~go
 widget.NewTabBookTab("Tab",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{0, 255, 0, 0xff})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)
~~~

Required Options
------
* TabButtonImage - The TabBook widget requires two [ButtonImage](https://pkg.go.dev/github.com/ebitenui/ebitenui/widget#ButtonImage) objects which contains 4 separate NineSlice images each. There is an idle ButtonImage object and a selected ButtonImage object
    * Idle - Required - Shown when not active
    * Hover - Required - Shown when cursor is over the tab
    * Pressed - Required - Shown when the user clicks the tab
    * Disabled - Required if disabled - Shown if the tab has been disabled
* TabButtonText - The TabBook requires a font face, and a [ButtonTextColor](https://pkg.go.dev/github.com/ebitenui/ebitenui/widget#ButtonTextColor) object. The ButtonTextColor has two colors
    * Idle - Required - Used when button is active
    * Disabled - Required if disabled - Used when button is disabled
* Tabs - An array of Tabs (see above)

Common Options
------
* TabButtonOpts - This allows you to set options for the tab buttons such as minimum size or text padding.
* TabButtonSpacing - This allows you to specify how far apart the Buttons are spaced. Default: 0

Event Handlers
------
1. TabSelectedHandler - A callback that is fired when the user changes tabs with these [args](https://pkg.go.dev/github.com/ebitenui/ebitenui/widget#TabBookTabSelectedEventArgs)

Methods
------
* func (t *TabBook) SetTab(tab *TabBookTab) - Sets the current visible Tab
* func (t *TabBook) Tab() *TabBookTab - Gets the current visible Tab

Live Demo
-----------

<iframe src="/wasm_control/tabbook.html" height="400" width="400" title="Live Demo" scrolling="no"></iframe>


Code
-------
See full example here: [https://github.com/ebitenui/ebitenui/blob/master/_examples/widget_demos/tabbook/main.go](https://github.com/ebitenui/ebitenui/blob/master/_examples/widget_demos/tabbook/main.go)

~~~go
	// load images for button states: idle, hover, and pressed
	buttonImage, _ := loadButtonImage()
	selectedImage, _ := loadSelectedImage()

	// load text font
	face, _ := loadFont(20)

	// construct a new container that serves as the root of the UI hierarchy
	rootContainer := widget.NewContainer(
		// the container will use a plain color as its background
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{0x13, 0x1a, 0x22, 0xff})),

		// the container will use an anchor layout to layout its single child widget
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	// Create the first tab
	// A TabBookTab is a labelled container. The text here is what will show up in the tab button
	tabRed := widget.NewTabBookTab("Red Tab",
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.RGBA{255, 0, 0, 255})),
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	redBtn := widget.NewText(
		widget.TextOpts.Text("Red Tab Button", face, color.White),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	tabRed.AddChild(redBtn)

	tabBook := widget.NewTabBook(
		widget.TabBookOpts.TabButtonImage(buttonImage, selectedImage),
		widget.TabBookOpts.TabButtonText(face, &widget.ButtonTextColor{Idle: color.White}),
		widget.TabBookOpts.TabButtonSpacing(0),
		widget.TabBookOpts.ContainerOpts(
			widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				StretchHorizontal:  true,
				StretchVertical:    true,
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
			),
		),
		widget.TabBookOpts.TabButtonOpts(widget.StateButtonOpts.ButtonOpts(
			widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.MinSize(135, 0)),
		)),
		widget.TabBookOpts.Tabs(tabRed),
	)
	// add the tabBook as a child of the container
	rootContainer.AddChild(tabBook)

...

func loadButtonImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.RGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.RGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.RGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadSelectedImage() (*widget.ButtonImage, error) {
	idle := image.NewNineSliceColor(color.RGBA{R: 100, G: 100, B: 120, A: 255})

	hover := image.NewNineSliceColor(color.RGBA{R: 100, G: 100, B: 120, A: 255})

	pressed := image.NewNineSliceColor(color.RGBA{R: 100, G: 100, B: 120, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func loadFont(size float64) (font.Face, error) {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(ttfFont, &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}), nil
}

~~~