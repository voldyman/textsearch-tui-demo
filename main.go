package main

import (
	"fmt"
	"log"

	"github.com/marcusolsson/tui-go"
)

func main() {
	searcher := NewSearcher()
	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	selected := -1
	suggestionsList := tui.NewList()

	suggestionBox := tui.NewVBox(suggestionsList)
	suggestionBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(inputBox, suggestionBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	ui, err := tui.New(chat)
	if err != nil {
		log.Fatal(err)
	}
	searcher.Search("abc")

	input.OnChanged(func(e *tui.Entry) {
		txt := e.Text()

		if len(txt) > 0 {
			results := searcher.Search(txt)
			suggestionsList.RemoveItems()
			if len(results) == 0 {
				selected = -1
				return
			}

			suggestionsList.AddItems(results...)
			selected = 0
			suggestionsList.SetSelected(selected)
		}
	})

	input.OnSubmit(func(e *tui.Entry) {
		cmd := ""
		if suggestionsList.Selected() >= 0 {
			cmd = suggestionsList.SelectedItem()
		}
		ui.Quit()
		fmt.Println(cmd)
	})

	ui.SetKeybinding("Esc", func() { ui.Quit() })
	ui.SetKeybinding("Down", func() {
		selected = wrap(selected+1, suggestionsList.Length())
		suggestionsList.SetSelected(selected)
	})
	ui.SetKeybinding("Up", func() {
		selected = wrap(selected-1, suggestionsList.Length())
		suggestionsList.SetSelected(selected)
	})

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func wrap(num, limit int) int {
	if num < 0 {
		return limit - 1
	}
	if num >= limit {
		return num % limit
	}
	return num
}
