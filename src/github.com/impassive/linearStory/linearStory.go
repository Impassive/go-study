package main

import "fmt"

//storyPage text, nextStoryPage
type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage) {
	if page == nil {
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

//Homework
func addPage(page *storyPage, text string) storyPage {
	pageN := storyPage{text, nil}
	page.nextPage = &pageN
	return pageN
}

func deletePage(page *storyPage, prevPage *storyPage) {
	prevPage.nextPage = nil
	page = nil
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)

	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	page4 := addPage(&page3, "page 4")
	fmt.Println(&page4)

	deletePage(&page4, &page3)
	fmt.Println(&page4)

}
