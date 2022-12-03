package rucksack

import "fmt"

type Rucksack struct {
	Data  string
	Items []Item
}

type Item struct {
	Label    rune
	Priority int
}

func (i Item) String() string {
	return fmt.Sprintf("[%s (%v): %v]", string(i.Label), i.Label, i.Priority)
}

// GeneratePriorityMap generates a map of item labels (strings) to priorities (ints).
func GeneratePriorityMap() map[rune]int {
	priorityMap := map[rune]int{}

	// fmt.Println("Generating priority map...")

	// Generate the item priority map dynamically to prevent needing to hard code
	// individual priority values for every letter (both upper and lowercase).
	start := int('a')
	for c := 'a'; c <= 'z'; c++ {
		offset := int(c) - start // determine offset of character for priority

		// fmt.Println(string(c), offset+1, string(c-32), offset+26+1)
		priorityMap[c] = offset + 1
		priorityMap[c-32] = offset + 26 + 1
	}

	return priorityMap
}

func ParseRucksack(itemMap map[rune]int, in string) *Rucksack {
	var items []Item
	for _, c := range in {
		item := &Item{
			Label:    c,
			Priority: itemMap[c],
		}

		items = append(items, *item)
	}

	return &Rucksack{
		Data:  in,
		Items: items,
	}
}

// Determine the most common rune between the two compartments of the rucksack.
func (r *Rucksack) DetermineCommonType() rune {
	compartmentLength := len(r.Items) / 2
	first := r.Items[0:compartmentLength]
	second := r.Items[compartmentLength:len(r.Items)]

	// Loop through one of the halves and check each item of the second half to see
	// if there's a match. As soon as we have one, ensure we return it.
	for _, i1 := range first {
		for _, i2 := range second {
			if i1.Label == i2.Label {
				return i1.Label
			}
		}
	}

	// Otherwise, we return a '!' rune to signify there is no common item type.
	return '!'
}

// Find the common item type across a variable number of Rucksacks.
func FindCommonItem(sacks ...Rucksack) rune {
	sackCount := len(sacks)
	runeCounts := []map[rune]int{}

	for _, r := range sacks {
		runeCount := map[rune]int{}
		for _, i := range r.Items {
			if runeCount[i.Label] > 0 {
				runeCount[i.Label] = runeCount[i.Label] + 1
			} else {
				runeCount[i.Label] = 1
			}
		}

		runeCounts = append(runeCounts, runeCount)
	}

	// Finally, loop through the first sack items, and return as soon as we find
	// one contained within all of the sacks.
	first := runeCounts[0]
	for r := range first {
		// For each rune within the sack, check to see that every other sack
		// contains at least one instance.
		within := 1
		// fmt.Println("Checking rune against other sacks:", string(r))
		for y := 1; y < sackCount; y++ {
			// fmt.Print("Checking sack:", y, "\n")
			// fmt.Print(string(r), " ")

			if runeCounts[y][r] > 0 {
				within += 1
			}
		}

		if within >= sackCount {
			// fmt.Println("Breaking loop early, common type found!", string(r))
			return r
		}
	}

	return '!'
}
