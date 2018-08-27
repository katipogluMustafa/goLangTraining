package main

import "fmt"

func main(){

  attended := map[string]bool{
  	"Mustafa" : true,
  	"Ahmet" : false,
  }
  if attended["Mahmut"]{//Will be false if the person is not in the map
  	fmt.Println("Mahmut, was at the meeting")
  }
  if attended["Mustafa"]{
  	fmt.Println("Mustafa, was at the meeting")
  }
	/*
   * Sometimes you need to distinguish a missing entry from a zero value.Is there an entry for "Mahmut" or is that the empty string because it's not in the map att all? You can discriminate with a form of multiple assignment.
   */
	if val, ok := attended["Mahmut"]; ok {// you can discriminate the value and its existence
		fmt.Println(val)
	}else{
		fmt.Println("Unknown person")
	}

	if val, ok := attended["Mustafa"]; ok {// you can discriminate the value and its existence
		fmt.Println(val)
	}else{
		fmt.Println("Unknown person")
	}
}
/*
 * An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map.
 */
