package maps

// Keys returns a slice of keys extracted from the input map.
func Keys[K comparable, E any](aMap map[K]E) []K {
	var keys []K
	for k := range aMap {
		keys = append(keys, k)
	}
	return keys
}
