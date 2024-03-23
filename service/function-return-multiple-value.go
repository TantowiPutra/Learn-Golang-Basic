package service

// Coba" return string, slice, ama map
func PrintFunctionReturnMultipleValue() (string, []string, map[int]string) {
	firstString := "Tantowi"
	firstSlice := []string {
		1: "Slice1",
		2: "Slice2",
		3: "Slice3",
		4: "Slice4",
		5: "Slice5",
	}
	firstMap := map[int]string{
		1: "Map1",
		2: "Map2",
		3: "Map3",
		4: "Map4",
		5: "Map5",
	}

	return firstString, firstSlice, firstMap
}