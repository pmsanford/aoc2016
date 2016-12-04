package main

import (
	"fmt"
	"strings"
)

type EncryptedName struct {
	name     []string
	sector   int
	checksum string
}

type LetterCount map[rune]uint

func parse_line(name string) EncryptedName {
	comps := strings.Split(name, "-")
	sector_str := comps[len(comps)-1]
	var (
		sector   int
		checksum string
	)
	fmt.Sscanf(sector_str, "%d[%s]", &sector, &checksum)
	return EncryptedName{comps[:len(comps)-1], sector, checksum[:len(checksum)-1]}
}

func combine_strings(letters []string) string {
	var out string
	for _, str := range letters {
		out += str
	}
	return out
}

func (mp *LetterCount) delete(key rune) {
	_, ok := (*mp)[key]
	if ok {
		delete(*mp, key)
	}
}

func (mp *LetterCount) pop_highest() (rune, uint) {
	var highest_rune rune
	var highest_num uint
	highest_num = 0
	for rn, num := range *mp {
		if (num == highest_num && rn < highest_rune) ||
			(num > highest_num) {
			highest_rune = rn
			highest_num = num
		}
	}

	mp.delete(highest_rune)
	return highest_rune, highest_num
}

func (mp *LetterCount) add(rn rune, num uint) {
	(*mp)[rn] = num
}

func (mp *LetterCount) sort() {
	new_map := make(LetterCount)
	for len(*mp) > 0 {
		rn, num := mp.pop_highest()
		new_map.add(rn, num)
	}

	for rn, num := range new_map {
		mp.add(rn, num)
	}
}

func count_letters(letters []string) LetterCount {
	mp := make(LetterCount)
	for _, letter := range combine_strings(letters) {
		val, pres := mp[letter]
		if pres {
			mp[letter] = val + 1
		} else {
			mp[letter] = 1
		}
	}
	return mp
}

func create_checksum(mp LetterCount) string {
	var checksum string
	mp.sort()
	for rn := range mp {
		checksum += string(rn)
	}
	return checksum
}

func main() {
	name := parse_line("aaaa-bbb-z-y-x-123[abxyz]")
	count := count_letters(name.name)
	fmt.Println(create_checksum(count))
}
