package main

func partOne(input []byte) int {
	checksum := 0
	currentFileID := 0
	disk := make([]int, 0)

	// Populate the disk with file IDs and free spaces (which are represented as -1)
	for i, block := range input {
		length := int(block - '0')

		if i%2 == 0 { // The block represents the length of a file
			for j := 0; j < length; j++ {
				disk = append(disk, currentFileID)
			}
			currentFileID++
		} else { // The block represents the length of free space
			for j := 0; j < length; j++ {
				disk = append(disk, -1)
			}
		}
	}

	for i := 0; i < len(disk); i++ {
		// Trim the trailing free spaces
		for disk[len(disk)-1] == -1 {
			disk = disk[:len(disk)-1]
		}

		// Replace free space with the last file ID
		if disk[i] == -1 {
			disk[i] = disk[len(disk)-1]
			disk = disk[:len(disk)-1]
		}

		// Update the checksum
		checksum += disk[i] * i
	}

	return checksum
}
