package main

import ("fmt"
        "strings")

func main() {
	songs := []string{"Glorious", 
	                                "All good things (Come to an end)",
					"No rest for the wicked",
					"Мальчик на девятке",
					"Myth",
					"Foggy mountain breakdown",
					"good 4 u",
				}
mySong := "Myth"
	for x := 1; x < len(songs); x++ {
	  song := songs[x]
	  condition := strings.Contains(strings.ToLower(song), strings.ToLower(mySong))
	  switch condition {
	  	case true: fmt.Println(song)
	  }
	};
}
