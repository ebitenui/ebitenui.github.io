+++
title = "Button"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-button"></i> '
weight = 1
+++

A simple button with text that can be set to a background color or image for each state and a callback to react to events.

<!--more-->

![preview](examples/wid_but_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Widget options

###### Text padding

Responsible for setting text offset from the center of the button.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/wid_but_tex_pad_lef.txt" lang="go" id="this" >}}
![lef](examples/wid_but_tex_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/wid_but_tex_pad_rig.txt" lang="go" id="this" >}}
![rig](examples/wid_but_tex_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/wid_but_tex_pad_top.txt" lang="go" id="this" >}}
![top](examples/wid_but_tex_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/wid_but_tex_pad_bot.txt" lang="go" id="this" >}}
![bot](examples/wid_but_tex_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text position

Responsible for setting text aligment position.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="StartxStart" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxsta.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_staxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxStart" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxsta.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_cenxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxStart" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxsta.txt" lang="go" id="this" >}}
![hei](examples/wid_but_tex_pos_endxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxCenter" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxcen.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_staxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}


{{% tab title="CenterxCenter" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxcen.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_cenxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxCenter" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxcen.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_endxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxEnd" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxend.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_staxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_staxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}} 

{{% tab title="CenterxEnd" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxend.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_cenxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_cenxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}} 

{{% tab title="EndxEnd" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxend.txt" lang="go" id="this" >}}
![wid](examples/wid_but_tex_pos_endxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_pos_endxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}


{{< /tabs >}}

###### Text Label

Responsible for setting multiline text on a button.

{{< tabs title="Label:" style="transparent" color="#131a22ff" >}}

{{% tab title="Login" %}}
{{< code src="assets/examples/wid_but_lab_log.txt" lang="go" id="this" >}}
![log](examples/wid_but_lab_log.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_lab_log.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Signup" %}}
{{< code src="assets/examples/wid_but_lab_sig.txt" lang="go" id="this" >}}
![horz](examples/wid_but_lab_sig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_lab_sig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Process BBCode

Responsible for processing text with BBCode. 

The only tag currently supported is `[color=][/color]` (without the `#`).

{{< tabs title="Enabled:" style="transparent" color="#131a22ff" >}}

{{% tab title="True" %}}
{{< code src="assets/examples/wid_but_tex_bbc_on.txt" lang="go" id="this" >}}
![on](examples/wid_but_tex_bbc_on.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_bbc_on.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="False" %}}
{{< code src="assets/examples/wid_but_tex_bbc_off.txt" lang="go" id="this" >}}
![off](examples/wid_but_tex_bbc_off.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_bbc_off.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Face

Responsible for setting the font family and font size.

{{< tabs title="Font:" style="transparent" color="#131a22ff" >}}

{{% tab title="Basicfont-13" %}}
{{< code src="assets/examples/wid_but_fon_b13.txt" lang="go" id="this" >}}
![b13](examples/wid_but_fon_b13.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_b13.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Inconsolata-16" %}}
{{< code src="assets/examples/wid_but_fon_i16.txt" lang="go" id="this" >}}
![i16](examples/wid_but_fon_i16.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_fon_i16.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Text Color

Responsible for setting the text color.

{{< tabs title="Color:" style="transparent" color="#131a22ff" >}}

{{% tab title="Idle" %}}
{{< code src="assets/examples/wid_but_tex_col_idl.txt" lang="go" id="this" >}}
![idle](examples/wid_but_tex_col_idl.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_col_idl.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Hover" %}}
{{< code src="assets/examples/wid_but_tex_col_hov.txt" lang="go" id="this" >}}
![hover](examples/wid_but_tex_col_hov.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_col_hov.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Pressed" %}}
{{< code src="assets/examples/wid_but_tex_col_pre.txt" lang="go" id="this" >}}
![pressed](examples/wid_but_tex_col_pre.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_tex_col_pre.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Image

Responsible for the background of the button; it can be filled with color or image tiles.

{{< tabs title="Image:" style="transparent" color="#131a22ff" >}}

{{% tab title="PillxIdle" %}}
{{< code src="assets/examples/wid_but_img_pillxidl.txt" lang="go" id="this" >}}
![pillxidl](examples/wid_but_img_pillxidl.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_pillxidl.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="PillxHover" %}}
{{< code src="assets/examples/wid_but_img_pillxhov.txt" lang="go" id="this" >}}
![pillxhov](examples/wid_but_img_pillxhov.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_pillxhov.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="PillxPressed" %}}
{{< code src="assets/examples/wid_but_img_pillxpre.txt" lang="go" id="this" >}}
![pillxpre](examples/wid_but_img_pillxpre.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_pillxpre.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="GemxIdle" %}}
{{< code src="assets/examples/wid_but_img_gemxidl.txt" lang="go" id="this" >}}
![gemxidl](examples/wid_but_img_gemxidl.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_gemxidl.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="GemxHover" %}}
{{< code src="assets/examples/wid_but_img_gemxhov.txt" lang="go" id="this" >}}
![gemxhov](examples/wid_but_img_gemxhov.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_gemxhov.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="GemxPressed" %}}
{{< code src="assets/examples/wid_but_img_gemxpre.txt" lang="go" id="this" >}}
![gemxpre](examples/wid_but_img_gemxpre.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_gemxpre.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="ColxIdle" %}}
{{< code src="assets/examples/wid_but_img_colxidl.txt" lang="go" id="this" >}}
![colxidl](examples/wid_but_img_colxidl.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_colxidl.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="ColxHover" %}}
{{< code src="assets/examples/wid_but_img_colxhov.txt" lang="go" id="this" >}}
![colxhov](examples/wid_but_img_colxhov.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_colxhov.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="ColxPressed" %}}
{{< code src="assets/examples/wid_but_img_colxpre.txt" lang="go" id="this" >}}
![colxpre](examples/wid_but_img_colxpre.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_img_colxpre.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}} 

###### Ignore Transparent Pixels

Responsible for ignoring transparent pixels when the button is clicked.

{{< tabs title="Enabled:" style="transparent" color="#131a22ff" >}}

{{% tab title="True" %}}
{{< code src="assets/examples/wid_but_ign_pix_on.txt" lang="go" id="this" >}}
![on](examples/wid_but_ign_pix_on.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_ign_pix_on.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="False" %}}
{{< code src="assets/examples/wid_but_ign_pix_off.txt" lang="go" id="this" >}}
![off](examples/wid_but_ign_pix_off.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/wid_but_ign_pix_off.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

### Widget handlers

###### Clicked Handler

Triggered when a release is performed anywhere after a press, its the default in most cases.

```go
button := widget.NewButton(
    widget.ButtonOpts.ClickedHandler(
        func(args *widget.ButtonClickedEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Button Clicked\n")
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Print(b.String())
        },
    ),
)
```

###### Pressed Handler

Triggered when a press is performed inside a widget, useful for long presses with different timings.

```go
button := widget.NewButton(
    widget.ButtonOpts.PressedHandler(
        func(args *widget.ButtonPressedEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Button Pressed\n")
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Print(b.String())
        },
    ),
)
```

###### Released Handler

Triggered when a release is performed inside a widget, useful for drag-n-drop things from one widget to another.

```go
button := widget.NewButton(
    widget.ButtonOpts.ReleasedHandler(
        func(args *widget.ButtonReleasedEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Button Released\n")
            fmt.Fprintf(&b, "> Inside: %v\n", args.Inside)
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Print(b.String())
        },
    ),
)
```

###### Cursor Entered Handler

Triggered when the cursor enters the button area.

```go
button := widget.NewButton(
    widget.ButtonOpts.CursorEnteredHandler(
        func(args *widget.ButtonHoverEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Cursor Entered\n")
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Fprintf(&b, "> Diff: %d, %d\n", args.DiffX, args.DiffY)
            fmt.Print(b.String())
        },
    ),
)
```

###### Cursor Moved Handler

Triggered when the cursor moves within the button area between entered and exited events.

```go
button := widget.NewButton(
    widget.ButtonOpts.CursorMovedHandler(
        func(args *widget.ButtonHoverEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Cursor Moved\n")
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Fprintf(&b, "> Diff: %d, %d\n", args.DiffX, args.DiffY)
            fmt.Print(b.String())
        },
    ),
)
```

###### Cursor Exited Handler

Triggered when the cursor exits the button area.

```go
button := widget.NewButton(
    widget.ButtonOpts.CursorExitedHandler(
        func(args *widget.ButtonHoverEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Cursor Exited\n")
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Fprintf(&b, "> Diff: %d, %d\n", args.DiffX, args.DiffY)
            fmt.Print(b.String())
        },
    ),
)
```

###### State Changed Handler

Triggered when the button's state changes like disabled, hovered, pressed, etc.

```go
button := widget.NewButton(
    widget.ButtonOpts.StateChangedHandler(
        func(args *widget.ButtonChangedEventArgs) {
            var b strings.Builder
            fmt.Fprintf(&b, "Button State Changed\n")
            fmt.Fprintf(&b, "> State: %v\n", args.State)
            fmt.Fprintf(&b, "> Offset: %d, %d\n", args.OffsetX, args.OffsetY)
            fmt.Print(b.String())
        },
    ),
)
```
