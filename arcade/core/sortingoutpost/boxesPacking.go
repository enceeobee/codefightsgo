package sortingoutpost

import (
	"fmt"
)

type box struct {
	l int
	w int
	h int
}

func boxesPacking(length []int, width []int, height []int) bool {
	if len(length) == 1 {
		return true
	}

	boxes := []box{}

	for i := range length {
		boxes = append(boxes, box{length[i], width[i], height[i]})
	}

	fmt.Println("boxes", boxes)

	// Take a box, and search the rest of the boxes for one that fits.
	// Recursive?

	return boxFits(boxes[0], boxes[1:])
}

func boxFits(b box, boxes []box) bool {
	if len(boxes) == 0 {
		return true
	}

	// Find if there's a box that fits
	// If no boxes fit, return false
	// Else, remove that box and run the func again

	/*
		Example: {2 4 2} {4 3 5}
		Turning {2 4 2} to {2 2 4} would work

		How to turn: flip on one of 3 axes
		x-axis -> turn on side = l remains, w becomes h, h becomes w
		y-axis -> turn on top = w remains, l becomes h, h becomes l
		z-axis -> turn sideways = h remains, l becomes w, w becomes l
	*/

	var doesFit bool
	for i, bx := range boxes {
		doesFit = fits(bx, b)

		if !doesFit {

			// Flip and try again
			// TODO
		}

		if doesFit {
			// Remove bx from boxes
			boxes = append(boxes[:i], boxes[i+1:]...)
			return boxFits(bx, boxes)
		}

		return false
	}

	return false
}

func fits(b1 box, b2 box) bool {
	return b1.l < b2.l && b1.w < b2.w && b1.h < b2.h
}
