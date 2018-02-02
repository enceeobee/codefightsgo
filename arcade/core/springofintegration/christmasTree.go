package springofintegration

import (
	s "strings"
)

func christmasTree(levelNum, levelHeight int) []string {
	var crown, levels, foot, tree []string
	var offset, footFill, levelFill int
	var level string

	longRow := s.Repeat("*", 3+(levelHeight*2)+((levelNum-1)*2))
	lenLongRow := len(longRow)
	midpoint := (lenLongRow / 2)

	// Set up foot
	if levelHeight%2 == 1 {
		footFill = levelHeight
	} else {
		footFill = levelHeight + 1
	}

	offset = midpoint - (footFill / 2)
	footRow := s.Repeat(" ", offset) + s.Repeat("*", footFill)

	// Set crown
	crown = []string{
		s.Repeat(" ", midpoint) + "*",
		s.Repeat(" ", midpoint) + "*",
		s.Repeat(" ", midpoint-1) + "***",
	}

	for n := 0; n < levelNum; n++ {
		// Set levels
		levelFill = 5 + (2 * n)

		for h := 0; h < levelHeight; h++ {
			offset = midpoint - (levelFill / 2)
			level = s.Repeat(" ", offset) + s.Repeat("*", levelFill)
			levels = append(levels, level)

			levelFill += 2
		}

		foot = append(foot, footRow)
	}

	tree = append(tree, crown...)
	tree = append(tree, levels...)
	tree = append(tree, foot...)

	return tree
}
