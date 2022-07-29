# strcombiner: combines strings from a dictionary into a new string with a given length

## Introduction
Strcombiner is a Go library that combines strings from a dictionary into a new string which length is as closer to given value as possible.

## Basic Usage

```go
package main

import (
	"fmt"

	"github.com/GHopperMSK/strcombiner"
)

func main() {
    // prepare data
	dict := []string{"100", "101", "102", "109", "110", "6593", "7832", "7981", "8923", "9800", "9950", "9951", "9952", "923423", "7369292", "74234223", "7453321", "89322331"}
	for i := 0; i < len(dict); i++ {
		dict[i] = "user_id=" + dict[i] + ","
	}

	fmt.Print("Init dictionary: ")
	fmt.Printf("%+v\n", dict)
    // Init dictionary: [user_id=100, user_id=101, user_id=102, user_id=109, user_id=110, user_id=6593, user_id=7832, user_id=7981, user_id=8923, user_id=9800, user_id=9950, user_id=9951, user_id=9952, user_id=923423, user_id=7369292, user_id=74234223, user_id=7453321, user_id=89322331,]
	res := strcombiner.Combine(dict, 0, 0, 154)
	fmt.Print("Result after the first call: ")
	fmt.Printf("%+v\n", res)
    // Result after the first call: {Length:154 Words:map[0:user_id=100, 1:user_id=101, 2:user_id=102, 3:user_id=109, 4:user_id=110, 5:user_id=6593, 6:user_id=7832, 7:user_id=7981, 8:user_id=8923, 9:user_id=9800, 10:user_id=9950, 11:user_id=7369292,]}
	fmt.Print("Remove processed values from the dictionary: ")
	strcombiner.RemoveProcessed(&dict, &res)
	fmt.Printf("%+v\n", dict)
    // Remove processed values from the dictionary: [user_id=9951, user_id=9952, user_id=923423, user_id=74234223, user_id=7453321, user_id=89322331,]
	fmt.Println("")

	fmt.Print("Result after the second call: ")
	res = strcombiner.Combine(dict, 0, 0, 154)
	fmt.Printf("%+v\n", res)
    // Result after the second call: {Length:91 Words:map[0:user_id=9951, 1:user_id=9952, 2:user_id=923423, 3:user_id=74234223, 4:user_id=7453321, 5:user_id=89322331,]}
	fmt.Print("Remove processed values from the dictionary: ")
	strcombiner.RemoveProcessed(&dict, &res)
	fmt.Printf("%+v\n", dict)
    // Remove processed values from the dictionary: []
	fmt.Println("")

	fmt.Print("Result after the third call: ")
	res = strcombiner.Combine(dict, 0, 0, 154)
	fmt.Printf("%+v\n", res)
    // Result after the third call: {Length:-1 Words:map[]}
}
```

