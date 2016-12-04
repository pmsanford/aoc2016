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

func main() {
	fmt.Println(parse_line("aaaa-bbb-z-y-x-123[abxyz]"))
}
