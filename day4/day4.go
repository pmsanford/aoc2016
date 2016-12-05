package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type EncryptedName struct {
	name     []string
	sector   int
	checksum string
	count    LetterCount
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
	return EncryptedName{comps[:len(comps)-1], sector, checksum[:len(checksum)-1], nil}
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

func (mp *LetterCount) get_sorted_letters() string {
	var out string
	for len(*mp) > 0 {
		rn, _ := mp.pop_highest()
		out += string(rn)
	}

	return out
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
	sorted := mp.get_sorted_letters()
	return sorted[:5]
}

func read_input(filename string, ch chan string) {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
}

func parse_lines(ch chan string, och chan EncryptedName) {
	for str := range ch {
		och <- parse_line(str)
	}
	close(och)
}

func do_counts(ch chan EncryptedName, och chan EncryptedName) {
	for enm := range ch {
		enm.count = count_letters(enm.name)
		och <- enm
	}
	close(och)
}

func (enm *EncryptedName) decrypt() []string {
	shift := rune(enm.sector % 26)
	decrypted := make([]string, 0)
	for _, word := range enm.name {
		var newword string
		for _, rn := range word {
			newchar := rn + shift
			if newchar > 122 {
				newchar = 96 + (newchar - 122)
			}
			newword += string(newchar)
		}
		decrypted = append(decrypted, newword)
	}
	return decrypted
}

func (enm *EncryptedName) is_valid() bool {
	if enm.count == nil {
		enm.count = count_letters(enm.name)
	}
	chk := create_checksum(enm.count)
	return chk == enm.checksum
}

func main() {
	ipt_ch := make(chan string)
	enm_ch := make(chan EncryptedName)
	ctd_ch := make(chan EncryptedName)
	go read_input("input.txt", ipt_ch)
	go parse_lines(ipt_ch, enm_ch)
	go do_counts(enm_ch, ctd_ch)
	for enm := range ctd_ch {
		if enm.is_valid() {
			fmt.Printf("%d: %s\n", enm.sector, enm.decrypt())
		}
	}
}
