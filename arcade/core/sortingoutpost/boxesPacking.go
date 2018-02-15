package sortingoutpost

import (
	"sort"
)

type box struct {
	l int
	w int
	h int
}

type boxSort []box

func (b boxSort) Len() int           { return len(b) }
func (b boxSort) Less(i, j int) bool { return volume(b[i]) < volume(b[j]) }
func (b boxSort) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func boxesPacking(length []int, width []int, height []int) bool {
	if len(length) == 1 {
		return true
	}

	boxes := []box{}

	for i := range length {
		boxes = append(boxes, box{length[i], width[i], height[i]})
	}

	// Sort by volume
	sort.Sort(boxSort(boxes))

	return boxFits(boxes[0], boxes[1:])
}

// Check if `b` fits in any of the boxes
func boxFits(b box, boxes []box) bool {
	if len(boxes) == 0 {
		return true
	}

	/*
		{l w h}
		Example: {2 4 2} {4 3 5}
		Turning {2 4 2} to {2 2 4} would work

		How to turn: flip on one of 3 axes
		x-axis -> turn on side = l remains, w becomes h, h becomes w
		y-axis -> turn on top = w remains, l becomes h, h becomes l
		z-axis -> turn sideways = h remains, l becomes w, w becomes l
	*/

	var doesFit bool
	for i, outer := range boxes {
		doesFit = fits(b, outer)

		// Flip x
		if !doesFit {
			doesFit = fits(box{b.l, b.h, b.w}, outer)
		}

		// Flip y
		if !doesFit {
			doesFit = fits(box{b.h, b.w, b.l}, outer)
		}

		// Flip z
		if !doesFit {
			doesFit = fits(box{b.w, b.l, b.h}, outer)
		}

		// Flip x & y
		if !doesFit {
			doesFit = fits(box{b.w, b.h, b.l}, outer)
		}

		// Flip x & z
		if !doesFit {
			doesFit = fits(box{b.h, b.l, b.w}, outer)
		}

		// Flip y & z
		if !doesFit {
			doesFit = fits(box{b.w, b.h, b.l}, outer)
		}

		if doesFit {
			// Remove outer from boxes
			boxes = append(boxes[:i], boxes[i+1:]...)
			return boxFits(outer, boxes)
		}

		return false
	}

	return false
}

func fits(inner box, outer box) bool {
	return inner.l < outer.l && inner.w < outer.w && inner.h < outer.h
}

func volume(b box) int {
	return b.l * b.w * b.h
}
