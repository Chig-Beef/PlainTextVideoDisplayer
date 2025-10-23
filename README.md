# Plain Text Video Viewer
After working with some PPM files, which are plain text images, I decided it
was only natural to create a plain text video format. This format is useful for
programmers to create little animations from their programs to show off what
they've made. For example, there is `examples/gol.ptv`, which is a rendered
output from a game of life simulation. This is to show how easy it is to create
a video that you can view, and send to others to view.

## Format
The format starts with a few initialising instructions.

First, the width command sets the width of the image.
```
WIDTH <int>
WIDTH 100
```

Next, the height command sets the height of the image.
```
HEIGHT <int>
HEIGHT 100
```

Those two are compulsory to any video.

Then, the delay command describes how long between each frame fires (in seconds).
```
DELAY <float>
DELAY 0.2
```
This example above means each frame lasts for 0.2s. At 60fps, this is 12 frames.

To compress information, we can use the color command to create new colors.
```
COLOR <name> <int> <int> <int>
COLOR B 0 0 0
```
The name should be a small number of characters, as this allows for the best
compression. The three integers correspond to RGB values in the color.

These compressed values can then be used as such.
```
<name>
B
```
And this will put this pixel in the video.

Otherwise, full values can be used.
```
<int> <int> <int>
0 0 255
```
And it will also put this pixel in the video.

## Placing Pixels
All pixels are stored in a one dimensional fashion in the file format, but
are then translated into a three dimensional video (width, height, frames).
The order of pixels starts at the top left of the first frame, moving to the
right. It then moves down row after row, filling the frame, in the order of
reading a book. Once a frame is full, it will create a new frame, and start
filling that in a similar fashion.
