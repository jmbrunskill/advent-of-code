package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Print the result
	fmt.Println(processInput(f, 6, 25))
}

func processInput(f io.Reader, maxX, maxY int) string {
	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)

	var img image
	img.sizeX = maxX
	img.sizeY = maxY

	layer := 0
	x := 0
	y := 0

	for s.Scan() {
		// fmt.Println(s.Text())
		d, err := strconv.Atoi(s.Text())
		if err != nil {
			log.Fatal(err)
		}

		// fmt.Printf("(%d,%d):%d\n", x, y, d)
		img.addPx(layer, xy{x, y}, d)
		y++
		if y >= img.sizeY {
			y = 0
			x++
		}
		if x >= img.sizeX {
			x = 0
			y = 0
			layer++
		}
	}

	// img.Print()
	img.Decode()

	return fmt.Sprintf("%v", 0)
}

type image struct {
	layers []map[xy]int
	sizeX  int
	sizeY  int
}

func (img *image) addPx(layer int, point xy, pix int) {
	if len(img.layers) <= layer {
		img.layers = append(img.layers, make(map[xy]int))
	}
	img.layers[layer][point] = pix
}

func (img *image) calcDigitThingy() int {
	leastZeroCount := 9999999
	leastZeroResult := -1
	layerZeroCount := 0
	layerOnesCount := 0
	layerTwosCount := 0
	for _, layer := range img.layers {
		// fmt.Printf("Layer:%d\n", i)

		for x := 0; x < img.sizeX; x++ {
			for y := 0; y < img.sizeY; y++ {
				if (layer[xy{x, y}] == 0) {
					layerZeroCount++
				}
				if (layer[xy{x, y}] == 1) {
					layerOnesCount++
				}
				if (layer[xy{x, y}] == 2) {
					layerTwosCount++
				}
			}
		}
		if layerZeroCount < leastZeroCount {
			leastZeroCount = layerZeroCount
			leastZeroResult = layerOnesCount * layerTwosCount
			// fmt.Println(i, layerZeroCount, layerOnesCount, layerTwosCount)
		}
		layerZeroCount = 0
		layerOnesCount = 0
		layerTwosCount = 0
	}

	return leastZeroResult

}

func (img *image) Print() {
	for i, layer := range img.layers {
		fmt.Printf("Layer:%d\n", i)

		for x := 0; x < img.sizeX; x++ {
			for y := 0; y < img.sizeY; y++ {
				fmt.Printf("%d", layer[xy{x, y}])
			}
			fmt.Println()
		}

	}
}

func (img *image) Decode() {
	for x := 0; x < img.sizeX; x++ {
		for y := 0; y < img.sizeY; y++ {

			for _, layer := range img.layers {
				if (layer[xy{x, y}] == 0) {
					fmt.Printf(" ")
					break
				} else if (layer[xy{x, y}] == 1) {
					fmt.Printf("#")
					break
				}
			}
		}
		fmt.Println()
	}

}

type xy struct {
	x int
	y int
}

func (p xy) String() string {
	return fmt.Sprintf("<x=%d, y=%d>", p.x, p.y)
}
