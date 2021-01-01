---
layout: docs
title: Usage
navbarsection: usage
---

How to Use Ebiten UI
====================

Ebiten UI is a fairly complex and flexible user interface engine. This page explains the concepts used in it.

Retained Mode
-------------

Ebiten UI adopts the "[retained mode]" model for user interfaces. This means that calls into the Ebiten UI API,
such as constructing a Button, do not cause actual rendering to happen. Constructing UI widgets is rather
declarative: The UI is constructed to contain a number of containers, widgets, and layouts, and Ebiten UI itself
causes the actual rendering calls when appropriate.

Widget Hierarchy
----------------

The entire user interface is constructed as a hierarchy in Ebiten UI:

- At the top level, there is the [UI] type. This mostly contains a reference to the root container, as well as
  references to a tool tip renderer, a drag & drop renderer, and any floating windows. The UI is used to render
  the entire user interface. It is also responsible for delivering events throughout the user interface.
- [Container] is the main type to group things together, for example a row of buttons. Containers may contain
  any number of widgets as well as other Containers. They are responsible for layouting their child widgets.
- At the lowest level, there are widgets such as [Button] or [Checkbox].

Widget
------

The [Widget] type is special in Ebiten UI: Since Go only uses composition rather than inheritance, every widget
implementation (such as Button) "has" a Widget. The Widget type is responsible for handling basic widget tasks
such as remembering position, or firing cursor enter/exit and click events. It also contains the widget's
layout data if necessary (see below.)

Layouter
--------

Widgets are usually not layouted by hand in Ebiten UI. Instead, they are grouped together as child widgets in a
Container, and the Container's Layouter is responsible for layouting the widgets.

There are a few Layouter implementations in Ebiten UI:

- [AnchorLayout] can only position exactly one widget, which it anchors to either a corner or edge of the Container.
  It can optionally stretch the widget along any of the directions.
- [GridLayout] can layout any number of widgets in a grid. It can position widgets differently for each grid cell,
  and it can also stretch them.
- [RowLayout] layouts any number of widgets in a single row or column. It also can position widgets differently
  and stretch them if desired.

To configure a single widget to be layouted differently from its siblings, an optional "layout data" can be set
on the widget. The type of layout data depends on the layout implementation being used. For example,
AnchorLayout requires AnchorLayoutData to be used.

NineSlice
---------

While Ebiten uses a basic image type to draw images onto the screen, this is not sufficient for many widgets.
For example, a Button's image needs to stretch or shrink depending on the Button's text, without distorting the image.
For this reason, Ebiten UI adopts the [9-slice scaling] technique.

[NineSlice] is basically a 3x3 grid of image tiles: The corner tiles are drawn as-is, while the edge tiles and the
center tile are stretched:

![9-slice scaling](https://upload.wikimedia.org/wikipedia/commons/thumb/7/7a/Traditional_scaling_vs_9-slice_scaling.svg/320px-Traditional_scaling_vs_9-slice_scaling.svg.png)

*Top: Traditional scaling, corners are distorted. Bottom: 9-slice scaling, corners aren't distorted.
(Image: Wikipedia)*

Options
-------

When constructing widgets or layouters, most of them support a number of options to configure their rendering or
behavior. For example, Button has options to configure the actual images used for rendering, the button's text,
font face, text color, padding, and so on.

As an example, a Button may be constructed like this:

~~~go
button := widget.NewButton(
  // specify the images to use
  widget.ButtonOpts.Image(buttonImage),

  // specify the button's text, the font face, and the color
  widget.ButtonOpts.Text("Hello, World!", face, &widget.ButtonTextColor{
    Idle: color.RGBA{0xdf, 0xf4, 0xff, 0xff},
  }),

  // specify that the button's text needs some padding for correct display
  widget.ButtonOpts.TextPadding(widget.Insets{
    Left:  30,
    Right: 30,
  }),

  // ... click handler, etc. ...
)
~~~

Some of the options are the same for each widget implementation, such as specifying a layout data.
In this case, the options allow for specifying widget options:

~~~go
button := widget.NewButton(
  // ... other button options ...

  // set general widget options
  widget.ButtonOpts.WidgetOpts(
    // instruct the container's anchor layout to center the button
    // both horizontally and vertically
    widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
      HorizontalPosition: widget.AnchorLayoutPositionCenter,
      VerticalPosition:   widget.AnchorLayoutPositionCenter,
    }),
  ),
)
~~~

Depending on the widget implementation, some of the options are required to be specified (such as Button's images),
while others are optional. The order of the options usually does not matter. Some options may be specified multiple times.



[9-slice scaling]: https://en.wikipedia.org/wiki/9-slice_scaling
[AnchorLayout]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#AnchorLayout
[Button]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Button
[Checkbox]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Checkbox
[Container]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Container
[GridLayout]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#GridLayout
[Layouter]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Layouter
[NineSlice]: https://pkg.go.dev/github.com/blizzy78/ebitenui/image#NineSlice
[retained mode]: https://en.wikipedia.org/wiki/Retained_mode
[RowLayout]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#RowLayout
[UI]: https://pkg.go.dev/github.com/blizzy78/ebitenui#UI
[Widget]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Widget
