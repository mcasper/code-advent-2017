package challenge

import (
	"fmt"
	"math"
)

type Layer struct {
	min         int
	max         int
	sideLength  int
	number      int
	totalLength int
}

func DayThree() {
	// input := 23
	// input := 1024
	input := 368078
	layers := make([]Layer, 0)
	initialLayer := Layer{min: 1, max: 1, sideLength: 1, number: 0, totalLength: 1}
	layers = append(layers, initialLayer)
	var targetLayer Layer

OuterLoop:
	for {
		var layer Layer
		lastLayer := layers[len(layers)-1]

		newLayerLength := lastLayer.sideLength + 2
		numCount := (newLayerLength * 4) - 4

		newRange := []int{(lastLayer.max + 1), (lastLayer.max + numCount)}

		layer = Layer{min: newRange[0], max: newRange[1], sideLength: newLayerLength, number: lastLayer.number + 1, totalLength: numCount}

		layers = append(layers, layer)

		for i := newRange[0]; i < newRange[1]; i++ {
			if input == i {
				targetLayer = layer
				break OuterLoop
			}
		}
	}

	targetGrid := make([]int, 0)
	for i := targetLayer.min; i <= targetLayer.max; i++ {
		targetGrid = append(targetGrid, i)
	}

	if targetLayer.totalLength != len(targetGrid) {
		fmt.Println("Not equal!")
	}

	var targetIndex int
	for i := 0; i < len(targetGrid); i++ {
		if targetGrid[i] == input {
			targetIndex = i
		}
	}

	var position int
	if targetIndex > targetLayer.sideLength {
		offset := int(math.Floor(float64(targetIndex) / float64(targetLayer.sideLength)))
		position = (targetIndex % targetLayer.sideLength) - 1 - offset
	} else {
		position = targetIndex
	}

	offset := int(math.Floor(float64(targetIndex) / float64(targetLayer.sideLength)))
	middle := (targetLayer.sideLength / 2) + (offset * targetLayer.sideLength) - 1

	movesToMiddle := int(math.Abs(float64(targetIndex)-float64(middle))) + 1

	fmt.Println("targetIndex:", targetIndex)
	fmt.Println("position:", position)
	fmt.Println("middle:", middle)
	fmt.Println("sidelength:", targetLayer.sideLength)
	fmt.Println("layer:", targetLayer.number)
	fmt.Println("Number of moves is:", (movesToMiddle + targetLayer.number))
	fmt.Println(movesToMiddle, targetLayer.number)
	fmt.Println(targetLayer)
}
