package main

func partTwo(input []byte) int {
	checksum := 0

	type diskSection struct {
		pos int
		len int
	}

	files := make(map[int]diskSection)
	spaces := make([]diskSection, 0)

	currentFileID := 0
	currentDiskPos := 0

	for i, block := range input {
		length := int(block - '0')

		if i%2 == 0 { // The block represents the length of a file
			files[currentFileID] = diskSection{currentDiskPos, length}
			currentFileID++
		} else { // The block represents the length of free space
			if length > 0 {
				spaces = append(spaces, diskSection{currentDiskPos, length})
			}
		}

		currentDiskPos += length
	}

	// Iterate over files from the end of the disk
	for currentFileID > 0 {
		currentFileID--

		file := files[currentFileID]

		for i, space := range spaces {
			if space.pos > file.pos {
				spaces = spaces[:i]
				break
			}

			if space.len >= file.len {
				files[currentFileID] = diskSection{space.pos, file.len}

				if space.len == file.len {
					spaces = append(spaces[:i], spaces[i+1:]...)
				} else {
					spaces[i] = diskSection{space.pos + file.len, space.len - file.len}
				}

				break // File was moved
			}
		}
	}

	for id, file := range files {
		start := file.pos
		end := file.pos + file.len - 1

		// Calculate the sum using arithmetic series formula
		checksum += id * (file.len * (start + end) / 2)
	}

	return checksum
}
