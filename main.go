package main

import (
	"fmt"
)

func main() {
	s, t := "anagram", "nagaram"
	fmt.Println(isAnagram(s, t))

}

//func isAnagram(s string, t string) bool {
//
//	if len(s) != len(t) {
//		return false
//	}
//	smap := make(map[int32]int)
//	for _, v := range s {
//		smap[v]++
//	}
//
//	tmap := make(map[int32]int)
//	for _, v := range t {
//		tmap[v]++
//	}
//	return reflect.DeepEqual(smap, tmap)
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a := make([]int, 26)

	for i, char := range s {
		a[char-97]++
		a[t[i]-97]--
	}
	for _, c := range a {
		if c != 0 {
			return false
		}
	}
	return true
}
