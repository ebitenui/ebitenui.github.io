+++
title = "Anchor"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-anchor"></i> '
weight = 1
+++

Anchor layout allows the child containers to be anchored to a specific constraint.

<!--more-->

![preview](examples/lay_anc_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Layout options

###### Padding

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/lay_anc_pad_lef.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/lay_anc_pad_rig.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/lay_anc_pad_top.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_anc_pad_bot.txt" lang="go" id="this" >}}
![left](examples/lay_anc_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

### Layout data

###### Position

Responsible for aligning the child container along the horizontal and vertical axes.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="StartxStart" %}}
{{< code src="assets/examples/lay_anc_pos_staxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_staxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_staxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxStart" %}}
{{< code src="assets/examples/lay_anc_pos_cenxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_cenxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_cenxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxStart" %}}
{{< code src="assets/examples/lay_anc_pos_endxsta.txt" lang="go" id="this" >}}
![hei](examples/lay_anc_pos_endxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_endxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxCenter" %}}
{{< code src="assets/examples/lay_anc_pos_staxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_staxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_staxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxCenter" %}}
{{< code src="assets/examples/lay_anc_pos_cenxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_cenxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_cenxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxCenter" %}}
{{< code src="assets/examples/lay_anc_pos_endxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_endxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_endxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxEnd" %}}
{{< code src="assets/examples/lay_anc_pos_staxend.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_staxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_staxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxEnd" %}}
{{< code src="assets/examples/lay_anc_pos_cenxend.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_cenxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_cenxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxEnd" %}}
{{< code src="assets/examples/lay_anc_pos_endxend.txt" lang="go" id="this" >}}
![wid](examples/lay_anc_pos_endxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_anc_pos_endxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}
