package main

import (
	"fmt"
	"regexp"
)

type Searcher struct {
	items []string
}

func NewSearcher() *Searcher {
	items := []string{
		"ssh pi@192.168.0.11",
		"say 'She sells seashells by the seashore,\nThe shells she sells are seashells, I'm sure.\nSo if she sells seashells on the seashore,\nThen I'm sure she sells seashore shells. '",
		"ssh pi@192.168.0.11",
		"ssh 192.168.0.10",
		"brew upgrade",
		"brew update",
		"git clone https://gitlab.gnome.org/World/design/icon-library.git",
		"cd icon-library",
		"pkg-config --libs gtk-3",
		"cargo build",
		"cargo run",
		"rustfmt src/main.rs",
		"git clone git@github.com:ThePrimeagen/uhh.git",
		"cd uhh",
		"mkdir build",
		"cd build",
		"cmake -G Ninja ..",
		"git submodule update --init",
		"git status",
		"git add src CMakeLists.txt",
		"git commit -m 'Updated Uhh to created on stack and not on the heap'",
		"git remote add voldy git@github.com:voldyman/uhh.git",
		"git push voldy master",
		"git pull --rebase",
		"echo yah-yah",
		"echo is 7 a prime number",
	}
	return &Searcher{items}
}

func (s *Searcher) Search(text string) []string {
	results := []string{}
	for _, item := range s.items {
		if match, err := regexp.MatchString("(?i)"+text, item); match {
			results = append(results, item)
		} else if err != nil {
			fmt.Println("error", err)
		}
	}
	return results
}
