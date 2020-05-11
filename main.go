package main
import "fmt"

func main() {

	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	fmt.Println(title, "published by:", publisher, "drawn by:", artist, "and written by:", writer)
	fmt.Println(year, "it was released with pages:", pageNumber)
	fmt.Println("Grade achieved: ", grade)
	fmt.Println()

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe PaperClips"
	year = 2003
	pageNumber = 160
	grade = 9.0
	fmt.Println(title, "published by:", publisher, "drawn by:", artist, "and written by:", writer)
	fmt.Println(year, "it was released with pages:", pageNumber)
	fmt.Println("Grade achieved: ", grade)
}