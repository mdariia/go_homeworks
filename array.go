package main

import "fmt"

func main() {
	var songs [7]string = [7]string{"Glorious", 
	                                "All good things (Come to an end)",
					"No rest for the wicked",
					"Мальчик на девятке",
					"Myth",
					"Foggy mountain breakdown",
					"good 4 u",
				}

/*
        songs := [7]string{"Glorious", ¬
  8 ▸                      "All good things (Come to an end)",¬
  7 ▸       ▸       ▸      "No rest for the wicked",¬
  6 ▸       ▸       ▸      "Мальчик на девятке",¬
  5 ▸       ▸       ▸      "Myth",¬
  4 ▸       ▸       ▸      "Foggy mountain breakdown",¬
  3 ▸       ▸       ▸      "good 4 u",¬
  2 ▸       ▸       ▸       ▸       }
*/

/*
        var songs [7]string; songs = [7]string{"Glorious", ¬
 17 ▸                                          "All good things (Come to an end)",¬
 16 ▸       ▸       ▸       ▸       ▸          "No rest for the wicked",¬
 15 ▸       ▸       ▸       ▸       ▸          "Мальчик на девятке",¬
 14 ▸       ▸       ▸       ▸       ▸          "Myth",¬
 13 ▸       ▸       ▸       ▸       ▸          "Foggy mountain breakdown",¬
 12 ▸       ▸       ▸       ▸       ▸          "good 4 u",¬
 11 ▸       ▸       ▸       ▸       }
*/
 fmt.Println(songs[0])
 fmt.Println(songs[len(songs)-1])
 fmt.Println(songs[3])

 songs[3] = "Spa"
 
 songsSlice := songs[1:4]
 songsSlice[1] = "Ghostkeeper"
 fmt.Println(songs)
 fmt.Println(songsSlice)
 
 //var media []string = []string{
 }
