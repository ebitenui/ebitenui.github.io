+++
title = "Grid"
date = 2024-10-04T19:45:29+03:00
menuPre = '<i class="icon-grid"></i> '
weight = 3
+++

Grid layout puts all child containers in an even grid, wrapping each row based on the number of columns.

<!--more-->

![preview](examples/lay_gri_pre.png)

{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pre.txt" lang="go" id="root" >}}
{{% /expand %}}

### Layout options

###### Columns

Responsible for wrapping cells beyond a given number.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}

{{% tab title="1" %}}
{{< code src="assets/examples/lay_gri_col_1.txt" lang="go" id="this" >}}
![all](examples/lay_gri_col_1.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_col_1.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="2" %}}
{{< code src="assets/examples/lay_gri_col_2.txt" lang="go" id="this" >}}
![all](examples/lay_gri_col_2.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_col_2.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="3" %}}
{{< code src="assets/examples/lay_gri_col_3.txt" lang="go" id="this" >}}
![all](examples/lay_gri_col_3.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_col_3.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="4" %}}
{{< code src="assets/examples/lay_gri_col_4.txt" lang="go" id="this" >}}
![all](examples/lay_gri_col_4.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_col_4.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="5" %}}
{{< code src="assets/examples/lay_gri_col_5.txt" lang="go" id="this" >}}
![all](examples/lay_gri_col_5.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_col_5.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Stretch

Responsible for choosing cell coordinates across each axis that will *evenly occupy* the remaining space.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}

{{% tab title="AllxAll" %}}
{{< code src="assets/examples/lay_gri_str_allxall.txt" lang="go" id="this" >}}
![all](examples/lay_gri_str_allxall.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_allxall.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Allx3" %}}
{{< code src="assets/examples/lay_gri_str_allx3.txt" lang="go" id="this" >}}
![row](examples/lay_gri_str_allx3.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_allx3.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="3xAll" %}}
{{< code src="assets/examples/lay_gri_str_3xall.txt" lang="go" id="this" >}}
![col](examples/lay_gri_str_3xall.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_3xall.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="1x1" %}}
{{< code src="assets/examples/lay_gri_str_1x1.txt" lang="go" id="this" >}}
![1x1](examples/lay_gri_str_1x1.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_1x1.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="3x3" %}}
{{< code src="assets/examples/lay_gri_str_3x3.txt" lang="go" id="this" >}}
![3x3](examples/lay_gri_str_3x3.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_3x3.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="5x5" %}}
{{< code src="assets/examples/lay_gri_str_5x5.txt" lang="go" id="this" >}}
![5x5](examples/lay_gri_str_5x5.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_5x5.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="2,3,4x3" %}}
{{< code src="assets/examples/lay_gri_str_234x3.txt" lang="go" id="this" >}}
![234x3](examples/lay_gri_str_234x3.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_234x3.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="3x2,3,4" %}}
{{< code src="assets/examples/lay_gri_str_3x234.txt" lang="go" id="this" >}}
![3x234](examples/lay_gri_str_3x234.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_3x234.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="2,3,4x2,3,4" %}}
{{< code src="assets/examples/lay_gri_str_234x234.txt" lang="go" id="this" >}}
![234x234](examples/lay_gri_str_234x234.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_234x234.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="AllxOff" %}}
{{< code src="assets/examples/lay_gri_str_allxoff.txt" lang="go" id="this" >}}
![off](examples/lay_gri_str_allxoff.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_allxoff.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="OffxAll" %}}
{{< code src="assets/examples/lay_gri_str_offxall.txt" lang="go" id="this" >}}
![off](examples/lay_gri_str_offxall.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_offxall.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="OffxOff" %}}
{{< code src="assets/examples/lay_gri_str_offxoff.txt" lang="go" id="this" >}}
![off](examples/lay_gri_str_offxoff.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_str_offxoff.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Default stretch

Responsible for selecting all cells along each axis. Can be used instead of the previous option to avoid listing all cells. It can also be overridden by a regular property.

{{< tabs title="Stretch:" style="transparent" color="#131a22ff" >}}

{{% tab title="TruexTrue" %}}
{{< code src="assets/examples/lay_gri_dtr_tru_tru.txt" lang="go" id="this" >}}
![all](examples/lay_gri_dtr_tru_tru.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_dtr_tru_tru.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="TruexFalse" %}}
{{< code src="assets/examples/lay_gri_dtr_tru_fal.txt" lang="go" id="this" >}}
![all](examples/lay_gri_dtr_tru_fal.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_dtr_tru_fal.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="FalsexTrue" %}}
{{< code src="assets/examples/lay_gri_dtr_fal_tru.txt" lang="go" id="this" >}}
![all](examples/lay_gri_dtr_fal_tru.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_dtr_fal_tru.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="FalsexFalse" %}}
{{< code src="assets/examples/lay_gri_dtr_fal_fal.txt" lang="go" id="this" >}}
![all](examples/lay_gri_dtr_fal_fal.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_dtr_fal_fal.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Padding

Responsible for the offset of the parent container.

{{< tabs title="Padding:" style="transparent" color="#131a22ff" >}}

{{% tab title="Left" %}}
{{< code src="assets/examples/lay_gri_pad_lef.txt" lang="go" id="this" >}}
![lef](examples/lay_gri_pad_lef.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pad_lef.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Right" %}}
{{< code src="assets/examples/lay_gri_pad_rig.txt" lang="go" id="this" >}}
![rig](examples/lay_gri_pad_rig.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pad_rig.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Top" %}}
{{< code src="assets/examples/lay_gri_pad_top.txt" lang="go" id="this" >}}
![top](examples/lay_gri_pad_top.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pad_top.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Bottom" %}}
{{< code src="assets/examples/lay_gri_pad_bot.txt" lang="go" id="this" >}}
![top](examples/lay_gri_pad_bot.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pad_bot.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Spacing

Responsible for the offset between child elements.

{{< tabs title="Spacing:" style="transparent" color="#131a22ff" >}}

{{% tab title="0x0" %}}
{{< code src="assets/examples/lay_gri_spa_0x0.txt" lang="go" id="this" >}}
![0](examples/lay_gri_spa_0x0.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_spa_0x0.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="0x25" %}}
{{< code src="assets/examples/lay_gri_spa_0x25.txt" lang="go" id="this" >}}
![25](examples/lay_gri_spa_0x25.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_spa_0x25.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="25x0" %}}
{{< code src="assets/examples/lay_gri_spa_25x0.txt" lang="go" id="this" >}}
![50](examples/lay_gri_spa_25x0.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_spa_25x0.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="25x25" %}}
{{< code src="assets/examples/lay_gri_spa_25x25.txt" lang="go" id="this" >}}
![50](examples/lay_gri_spa_25x25.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_spa_25x25.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

### Layout data

###### Max size

Responsible for the allowable size of the child containers.

{{< tabs title="Max:" style="transparent" color="#131a22ff" >}}

{{% tab title="None" %}}
{{< code src="assets/examples/lay_gri_max_off.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_max_off.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_max_off.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Width" %}}
{{< code src="assets/examples/lay_gri_max_wid.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_max_wid.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_max_wid.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Height" %}}
{{< code src="assets/examples/lay_gri_max_hei.txt" lang="go" id="this" >}}
![hei](examples/lay_gri_max_hei.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_max_hei.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="Both" %}}
{{< code src="assets/examples/lay_gri_max_all.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_max_all.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_max_all.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}

###### Position

Responsible for the allowable size of the child containers. It works when the size of the cells is so limited by the maximum size that it is smaller than the size of the stretched cells that there is free space between them.

{{< tabs title="Position:" style="transparent" color="#131a22ff" >}}

{{% tab title="StartxStart" %}}
{{< code src="assets/examples/lay_gri_pos_staxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxStart" %}}
{{< code src="assets/examples/lay_gri_pos_cenxsta.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_cenxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_cenxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxStart" %}}
{{< code src="assets/examples/lay_gri_pos_endxsta.txt" lang="go" id="this" >}}
![hei](examples/lay_gri_pos_endxsta.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_endxsta.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxCenter" %}}
{{< code src="assets/examples/lay_gri_pos_staxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxCenter" %}}
{{< code src="assets/examples/lay_gri_pos_cenxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_cenxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_cenxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxCenter" %}}
{{< code src="assets/examples/lay_gri_pos_endxcen.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_endxcen.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_endxcen.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="StartxEnd" %}}
{{< code src="assets/examples/lay_gri_pos_staxend.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_staxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_staxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="CenterxEnd" %}}
{{< code src="assets/examples/lay_gri_pos_cenxend.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_cenxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_cenxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{% tab title="EndxEnd" %}}
{{< code src="assets/examples/lay_gri_pos_endxend.txt" lang="go" id="this" >}}
![wid](examples/lay_gri_pos_endxend.png)
{{% expand title="Full example" %}}
{{< code src="assets/examples/lay_gri_pos_endxend.txt" lang="go" id="root" >}}
{{% /expand %}}
{{% /tab %}}

{{< /tabs >}}