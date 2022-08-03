package main

import "fmt"

func main() {
	//map1()
	//map2()
	//iterateMap()
	map3()
}

func map1() {
	m := make(map[int]string) // empty map
	fmt.Println(m, "length :", len(m))

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	fmt.Println(m)

	fmt.Println(m[1])
	fmt.Println(m[61])

	s := m[10]
	fmt.Println("City with plate # 10 is", s)

	s = m[24]
	fmt.Println("City with plate # 24 is", s) // there is no defined city with that plate so it returns empty str

	delete(m, 10)
	s = m[10]
	fmt.Println("City with plate # 10 is", s) // after the deletion plate 10 also returns empty str

	m[81] += " Dümdüz" // string addition operations valid for maps
	fmt.Println(m[81])
	fmt.Println(m)

	// Replace an existing value
	m[35] = "İzmirX"
	fmt.Println(m)
}

func map2() {
	m1 := map[int]string{
		1: "Adana",
		6: "Ankara",
		//6:  "XXX",
		10: "Balıkesir",
		16: "Bursa",
	}

	fmt.Println(m1)

	var m2 map[string]string // this is a nil map
	fmt.Println(m2)
	// storing to a nil map causes panic
	// we must allocate the map before we store in it
	//m2["greeting"] = "Selam" // Panic!

	// however if map would be declared like below
	var m3 = make(map[string]string) // this would be an empty map so
	m3["greeting"] = "Selam"         // this wouldnt cause panic due to storing in nil map issue
	fmt.Println(m3)

	var m4 = map[string]string{} // this is also declares an empty map, not nil map
	m4["greeting"] = "Selam"     // so again this expression will not cause Panic
	fmt.Println(m4)
	fmt.Println("*****************")

	m2 = map[string]string{ // this wont cause panic since allocation also happens
		"greeting": "Selam",
		"goodbye":  "Hoşçakal",
	}
	fmt.Println(m2)
	s1, ok1 := m2["goodbye"]
	fmt.Println("ok1 value : ", ok1) // ok1 0 true means there is the key value pair in the map
	if ok1 {
		fmt.Println(s1)
	} else {
		fmt.Println("No such entry exist!")
	}

	if _, ok2 := m2["goodbyex"]; !ok2 {
		fmt.Println("No such entry exist!")
	}
}

func iterateMap() {
	m := make(map[int]string)
	fmt.Println(m)

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	fmt.Println("\nIterating Through Keys")
	for plate := range m {
		fmt.Printf("%d\n", plate) // prints keys in random order
	}

	fmt.Println("Iterating Through Values")
	for _, city := range m { // dummy var _ stands for holding the keys
		fmt.Printf("%s\n", city)
	}

	fmt.Println("Iterating Through Map")
	for plate, city := range m {
		fmt.Printf("%d\t%s\n", plate, city)
	}
}

func map3() {
	m := make(map[int]string) // empty map
	fmt.Println(m)

	m[1] = "Adana"
	m[6] = "Ankara"
	m[10] = "Balıkesir"
	m[16] = "Bursa"
	m[30] = "Hakkari"
	m[34] = "İstanbul"
	m[34] = "İstanbul"
	m[35] = "İzmir"
	m[61] = "Trabzon"
	m[81] = "Düzce"

	plates := make([]int, 0, len(m)) // this is not a map just emptyslice
	fmt.Println(plates)
	plates = append(plates, 1)
	fmt.Println(plates)

}
