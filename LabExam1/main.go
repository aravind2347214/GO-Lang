package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type post struct {
	postId          string
	content         string
	likes           int64
	shares          int64
	popolarityScore float32
}

var numPeople int

var numPost int = 2

var postList []post

// func postExists(postId string) bool {
// 	for i := 0; i < numPost; i++ {
// 		fmt.Printf("%s --- %s", postId, postList[i].postId)
// 		if postList[i].postId == postId {
// 			return true
// 		}
// 	}
// 	return false
// }

func getPostData() {
	postList = make([]post, numPost)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < numPost; i++ {
		// reenter:
		fmt.Printf("\nEnter Details of Post %d", i+1)
		fmt.Println("\nEnter unique Post ID")
		fmt.Scan(&postList[i].postId)

		// if !postExists(postList[i].postId) {
		// 	goto reenter
		// }

		fmt.Println("Enter Post Content")
		fmt.Scanln(&postList[i].content)
		postList[i].content, _ = reader.ReadString('\n')

		fmt.Println("Enter Post Likes")
		fmt.Scan(&postList[i].likes)
		fmt.Println("Enter Post Shares")
		fmt.Scan(&postList[i].shares)
		calculatePopularityScore(i)
	}
}

func displayAllPosts() {
	for i := 0; i < numPost; i++ {
		fmt.Printf("Post ID : %s\n", postList[i].postId)
		fmt.Printf("Post Content : %s\n", postList[i].content)
		fmt.Printf("Post Likes : %d\n", postList[i].likes)
		fmt.Printf("Post Shares : %d\n", postList[i].shares)
		fmt.Printf("Popularity Score : %d", postList[i].shares)
		fmt.Print("\n--------------------------------------\n")
	}
}

func showPostBasedOnThreshold() {
	var threshold int64
	fmt.Println("Enter Threshold to show the posts: ")
	fmt.Scan(&threshold)
	fmt.Println("Post Above threshold: ")

	for i := 0; i < numPost; i++ {
		if postList[i].likes+postList[i].shares >= threshold {
			fmt.Printf("Post ID : %s\n", postList[i].postId)
			fmt.Printf("Post Content : %s\n", postList[i].content)
			fmt.Printf("Post Likes : %d\n", postList[i].likes)
			fmt.Printf("Post Shares : %d\n", postList[i].shares)
			fmt.Printf("Popularity Score : %d", postList[i].shares)
			fmt.Print("\n--------------------------------------\n")
		}
	}
}

func searchPost() {
	var phrase string
	present := false
	fmt.Println("Search for a keyword or phrase in Content")
	fmt.Scan(&phrase)
	for i := 0; i < numPost; i++ {
		if strings.Contains(postList[i].content, phrase) {
			present = true
			fmt.Printf("Post ID : %s\n", postList[i].postId)
			fmt.Printf("Post Content : %s\n", postList[i].content)
			fmt.Printf("Post Likes : %d\n", postList[i].likes)
			fmt.Printf("Post Shares : %d\n", postList[i].shares)
			fmt.Printf("Popularity Score : %d\n", postList[i].shares)
			fmt.Print("\n--------------------------------------")
		}
	}

	if !present {
		fmt.Printf("Such a post Does not exist witht the Phrase :'%s'", phrase)
	}
}

func calculatePopularityScore(i int) {
	postList[i].popolarityScore = float32(postList[i].likes) + float32(postList[i].shares)
	postList[i].popolarityScore = postList[i].popolarityScore / float32(numPeople)
}

func main() {
	rerunChoice := 'y'
	var menuChoice int
	fmt.Println("Enter Number of People")
	fmt.Scan(&numPeople)

	getPostData()
	for rerunChoice == 'y' {
		fmt.Println("1.Display All Posts")
		fmt.Println("2.Filter Posts by Threshold")
		fmt.Println("3.Search by Phrase")
		fmt.Println("4.Exit")
		fmt.Scan(&menuChoice)

		switch {
		case menuChoice == 1:
			displayAllPosts()
		case menuChoice == 2:
			showPostBasedOnThreshold()
		case menuChoice == 3:
			searchPost()
		case menuChoice == 4:
			return
		}

		fmt.Print("\nDo you want  to continue? (y/n): ")
		fmt.Scan(&rerunChoice)
		if rerunChoice != 'y' {
			break
		}
	}
}
