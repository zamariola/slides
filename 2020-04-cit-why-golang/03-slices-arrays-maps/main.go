package main

func main() {

	//Arrays
	var myArray [2]string
	myArray[0] = "firstValue"
	myArray[1] = "secondValue"

	//Slices
	var mySlice []string
	mySlice = append(mySlice, "firstValue")
	mySlice = append(mySlice, "secondValue")

	myOtherSlice := make([]string, 10)
	myOtherSlice = append(myOtherSlice, "firstValue")

	//Maps
	myMap := make(map[string]string)
	myMap["firstKey"] = "firstValue"

}
