---
layout: docs
title: Custom Widgets
navbarsection: custom-widgets
---

Implementing Custom Widgets
===========================

With Ebiten UI it is not too difficult to implement custom widgets. These can range from just grouping
together a bunch of buttons to implementing a whole new widget. This page describes the concepts
behind doing so.

Retained Mode vs Immediate Mode
-------------------------------

From the API user's perspective, Ebiten UI adopts the "[retained mode]" model for user interfaces.
Actual widget implementations on the other hand always work in "[immediate mode]" - this is however
not noticeable for the API user. The reason for immediate mode is that Ebiten will cause widgets
to draw their current state in each GPU frame.

Implementation Structure
------------------------

The basic structure of a widget's implementation is as follows:

1. Each implementation needs to "have" a [Widget]. For simple "inheriting" implementations, an already
   existing widget implementation may be useful, while for completely new implementations Widget itself
   can be used.

   Some examples:

   - A new button implementation that is a regular [Button], but behaves a bit differently: The new
     implementation will use a regular Button as its widget.
   - A row of buttons behaving together as a single widget: The widget implementation could use
     [Container] as its widget, with the Container having regular Buttons as children.
   - A completely new widget that draws original images onto the screen: This implementation will use
     Widget directly.

   Do note that the implementation's widget will always be used to render first, after that, the
   implementation is free to render additional things. If it is not desirable that the widget renders
   anything, you need to use Widget (because it does not actually render anything onto the screen.)

2. A constructor function that takes options. Options are constructed using functions, such as
   ButtonOpts.Image. A constructor function and options are not strictly necessary, but they allow
   for easy future expansion.

3. Implement interface [PreferredSizeLocateableWidget]:

   - PreferredSize returns the widget's preferred size. This may depend on text being rendered,
     for example.
   - SetLocation is used by Layouters to position the widget on the screen.
   - GetWidget returns the implementation's widget, see above.

4. (optional) If the implementation is to actually render anything onto the screen, implement interface
   [Renderer].

5. (optional) If the implementation fires any events on its own (for example, on a timed basis),
   implement interface [Renderer] and fire events in the Render method.

6. (optional) If it is necessary to block user input from reaching widgets rendered below the widget
   (if overlapping), implement interface [Layerer].

Input Layering
--------------

see [Layer]

_TODO_



[Button]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Button
[Container]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Container
[immediate mode]: https://en.wikipedia.org/wiki/Immediate_mode_(computer_graphics)
[Layer]: https://pkg.go.dev/github.com/blizzy78/ebitenui/input#Layer
[Layerer]: https://pkg.go.dev/github.com/blizzy78/ebitenui/input#Layerer
[PreferredSizeLocateableWidget]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#PreferredSizeLocateableWidget
[Renderer]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Renderer
[retained mode]: https://en.wikipedia.org/wiki/Retained_mode
[Widget]: https://pkg.go.dev/github.com/blizzy78/ebitenui/widget#Widget
