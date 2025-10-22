package main

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Video struct {
	screenWidth  int
	screenHeight int
	delay        int
	frames       [][][]color.RGBA
}

func decodeVideo(data []byte) Video {
	video := Video{delay: 1}
	lines := strings.Split(string(data), "\r\n")
	if len(lines) == 1 {
		lines = strings.Split(string(data), "\n")
	}

	initialised := 2

	curFrameIndex := -1
	curRow := 0
	curCol := 0

	for i := range len(lines) {
		line := lines[i]
		spline := strings.Split(line, " ")

		if len(spline) == 0 {
			continue
		}

		if spline[0] == "" {
			continue
		}

		cmd := spline[0]

		switch cmd {
		case "WIDTH":
			width, err := strconv.Atoi(spline[1])
			if err != nil {
				fmt.Println("Invalid width")
				panic(err)
			}

			if width <= 0 {
				panic("Invalid width")
			}

			video.screenWidth = width
			initialised--

		case "HEIGHT":
			height, err := strconv.Atoi(spline[1])
			if err != nil {
				fmt.Println("Invalid height")
				panic(err)
			}

			if height <= 0 {
				panic("Invalid height")
			}

			video.screenHeight = height
			initialised--

		case "DELAY":
			delay, err := strconv.Atoi(spline[1])
			if err != nil {
				fmt.Println("Invalid delay")
				panic(err)
			}

			if delay <= 0 {
				panic("Invalid delay")
			}

			video.delay = delay

		default:
			if initialised != 0 {
				panic("Attempted to start writing frame data when uninitialised")
			}

			if len(spline) != 3 {
				panic("Invalid number of args for RGB pixel")
			}

			rawR, err := strconv.Atoi(spline[0])
			if err != nil {
				fmt.Println("Invalid red")
				panic(err)
			}
			if rawR < 0 || rawR >= 256 {
				panic("Invalid red")
			}

			rawG, err := strconv.Atoi(spline[1])
			if err != nil {
				fmt.Println("Invalid green")
				panic(err)
			}
			if rawG < 0 || rawG >= 256 {
				panic("Invalid green")
			}

			rawB, err := strconv.Atoi(spline[2])
			if err != nil {
				fmt.Println("Invalid blue")
				panic(err)
			}
			if rawB < 0 || rawB >= 256 {
				panic("Invalid blue")
			}

			r := byte(rawR)
			g := byte(rawG)
			b := byte(rawB)

			clr := color.RGBA{r, g, b, 255}

			curCol++
			if curCol == video.screenWidth {
				curCol = 0
				curRow++
			}

			// Starting a new frame?
			if curFrameIndex == -1 || curRow == video.screenHeight {
				curFrameIndex++
				curCol = 0
				curRow = 0
				video.frames = append(video.frames, make([][]color.RGBA, video.screenHeight))
				for r := range video.screenHeight {
					video.frames[curFrameIndex][r] = make([]color.RGBA, video.screenWidth)
				}
			}

			video.frames[curFrameIndex][curRow][curCol] = clr
		}
	}


	return video
}

func (video *Video) play() {
	ebiten.SetWindowSize(video.screenWidth, video.screenHeight)
	ebiten.SetWindowTitle("PTV Viewer")

	g := Game{
		video: video,
		delay: video.delay,
	}

	err := ebiten.RunGame(&g)
	if err != nil {
		panic(err)
	}
}

type Game struct {
	video *Video
	curFrame int
	delay    int
}

func (g *Game) Update() error {
	g.delay--
	if g.delay == 0 {
		g.curFrame++
		g.delay = g.video.delay
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.curFrame >= len(g.video.frames) {
		g.curFrame = len(g.video.frames)-1
	}
	frame := g.video.frames[g.curFrame]

	for row := range g.video.screenHeight {
		for col := range g.video.screenWidth {
			pixel := frame[row][col]
			screen.Set(col, row, pixel)
		}
	}
}

func (g *Game) Layout(int, int) (int, int) {
	return g.video.screenWidth, g.video.screenHeight
}
