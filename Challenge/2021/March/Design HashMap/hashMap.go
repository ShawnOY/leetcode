package main

import "fmt"

type MyHashMap struct {
	hashMap map[int]int
}

// Constructor initialize your data structure here.
func Constructor() MyHashMap {
	myHashMap := MyHashMap{
		hashMap: make(map[int]int),
	}

	return myHashMap
}

// Put will add/update the key, value always be non-negative.
func (mym *MyHashMap) Put(key int, value int) {
	mym.hashMap[key] = value
}

// Get returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key
func (mym *MyHashMap) Get(key int) int {
	if value, exists := mym.hashMap[key]; exists {
		return value
	}

	return -1
}

// Remove removes the mapping of the specified value key if this map contains a mapping for the key
func (mym *MyHashMap) Remove(key int) {
	delete(mym.hashMap, key)
}

func main() {
	var hashMap = Constructor()

	hashMap.Put(1, 1)
	hashMap.Put(2, 2)
	fmt.Println(hashMap.Get(1)) // Returns 1
	fmt.Println(hashMap.Get(3)) // Returns -1 (not found)
	hashMap.Put(2, 1)           // Update the existing value
	fmt.Println(hashMap.Get(2)) // Returns 1
	hashMap.Remove(2)           // Remove the mapping for 2
	fmt.Println(hashMap.Get(2)) // Returns -1 (not found)
}
