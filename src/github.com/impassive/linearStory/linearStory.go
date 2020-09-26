package main

import "fmt"

//storyPage text, nextStoryPage
type storyPage struct {
	text     string
	nextPage *storyPage
}

//type-specific func - receiver type for func
func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addtoEnd(text string) {
	pageToAdd := &storyPage{text, nil}
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = pageToAdd
}

func (page *storyPage) addAfter(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete

func main() {
	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addtoEnd("You are alone, and you need to find the sacred helmet before the bad guys do")
	page1.addtoEnd("You see a troll ahead")

	page1.addAfter("Testing AddAfter")
	page1.playStory()

	//functions - has a return value, may execute something
	//procedures - has no return value, just execute something
	//methods - functions attached to a struct/object/etc
}
