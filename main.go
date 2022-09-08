package main

import (
	"fmt"
	"reflect"
)

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	smap := make(map[int32]int)
	for _, v := range s {
		smap[v]++
	}

	tmap := make(map[int32]int)
	for _, v := range t {
		tmap[v]++
	}
	return reflect.DeepEqual(smap, tmap)
}
