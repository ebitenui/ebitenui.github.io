+++
title = "Row"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-row"></i> '
weight = 2
+++

Row layout places all child containers in one row or column. It can be useful for creating lists of widgets.

<!--more-->

![preview](examples/lay_row_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Layout options

###### Direction

Responsible for whether child containers will follow each other in rows or columns.

{{< tabs title="Direction:" style="transparent" color="#131a22ff" >}}

{{% tab title="Vertical" %}}
{{< code src="assets/examples/lay_row_dir_ver.txt" lang="go" id="this" >}}
![vert](examples/lay_row_dir_ver.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_dir_ver.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Horizontal" %}}
{{< code src="assets/examples/lay_row_dir_hor.txt" lang="go" id="this" >}}
![horz](examples/lay_row_dir_hor.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_dir_hor.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Padding

Layout allows you to specify padding for all child elements but not the itself. Please note that its not possible to specify padding at the *end*, so the *direction* will changed accordingly.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/lay_row_pad_lef.txt" lang="go" id="this" >}}
![lef](examples/lay_row_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/lay_row_pad_rig.txt" lang="go" id="this" >}}
![rig](examples/lay_row_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/lay_row_pad_top.txt" lang="go" id="this" >}}
![top](examples/lay_row_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_row_pad_bot.txt" lang="go" id="this" >}}
![top](examples/lay_row_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Spacing

Layout allows you to specify padding for all child elements but not the itself.

{{< tabs title="Spacing:" style="transparent" color="#131a22ff" >}}

{{% tab title="0" %}}
{{< code src="assets/examples/lay_row_spa_0.txt" lang="go" id="this" >}}
![0](examples/lay_row_spa_0.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_spa_0.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="25" %}}
{{< code src="assets/examples/lay_row_spa_25.txt" lang="go" id="this" >}}
![25](examples/lay_row_spa_25.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_spa_25.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="50" %}}
{{< code src="assets/examples/lay_row_spa_50.txt" lang="go" id="this" >}}
![50](examples/lay_row_spa_50.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_spa_50.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="75" %}}
{{< code src="assets/examples/lay_row_spa_75.txt" lang="go" id="this" >}}
![75](examples/lay_row_spa_75.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_spa_75.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

### Layout data

###### Stretch

Responsible for stretching the element along the entire length of the opposite axis.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}

{{% tab title="False" %}}
{{< code src="assets/examples/lay_row_str_fal.txt" lang="go" id="this" >}}
![fal](examples/lay_row_str_fal.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_str_fal.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="True" %}}
{{< code src="assets/examples/lay_row_str_tru.txt" lang="go" id="this" >}}
![tru](examples/lay_row_str_tru.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_str_tru.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Position

Responsible for aligning the element along the opposite axis if it is not stretched.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="Start" %}}
{{< code src="assets/examples/lay_row_pos_sta.txt" lang="go" id="this" >}}
![sta](examples/lay_row_pos_sta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pos_sta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Center" %}}
{{< code src="assets/examples/lay_row_pos_cen.txt" lang="go" id="this" >}}
![sta](examples/lay_row_pos_cen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pos_cen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="End" %}}
{{< code src="assets/examples/lay_row_pos_end.txt" lang="go" id="this" >}}
![sta](examples/lay_row_pos_end.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_pos_end.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}


###### Max size

Responsible for the allowable size of the container.

{{< tabs title="Max:" style="transparent" color="#131a22ff" >}}

{{% tab title="None" %}}
{{< code src="assets/examples/lay_row_max_off.txt" lang="go" id="this" >}}
![wid](examples/lay_row_max_off.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_max_off.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Width" %}}
{{< code src="assets/examples/lay_row_max_wid.txt" lang="go" id="this" >}}
![wid](examples/lay_row_max_wid.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_max_wid.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Height" %}}
{{< code src="assets/examples/lay_row_max_hei.txt" lang="go" id="this" >}}
![hei](examples/lay_row_max_hei.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_max_hei.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Both" %}}
{{< code src="assets/examples/lay_row_max_all.txt" lang="go" id="this" >}}
![wid](examples/lay_row_max_all.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_row_max_all.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}
