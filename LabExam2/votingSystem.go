package main

import (
	"fmt"
)

type Candidate struct {
	name  string
	age   int
	votes int
}

type Voter struct {
	name    string
	age     int
	isVoted bool
}

var candidates = []Candidate{}
var voters = []Voter{}

func registerVoter() {
	var tempVoter Voter
	tempVoter.isVoted = false
	fmt.Println("Enter Voter Name: ")
	fmt.Scan(&tempVoter.name)
	fmt.Println("Enter Voter Age: ")
	fmt.Scan(&tempVoter.age)
	if tempVoter.age < 18 {
		fmt.Println("Not Eligible to Vote")
		return
	}

	var tempCanName string
	fmt.Println("Enter Candidate Name You want to vote for")
	fmt.Scan(&tempCanName)
	votingFlag := 0
	// loop through the candidate slice and find the candidate with the same name
	for i := 0; i < len(candidates); i++ {
		if tempCanName == candidates[i].name {
			fmt.Println("Candidate Found")
			candidates[i].votes++
			tempVoter.isVoted = true
			votingFlag = 1
			voters = append(voters, tempVoter)
			break
		}
	}
	if votingFlag == 0 {
		fmt.Printf("candidate %s could not be found.Sorry", tempCanName)
	}
}

func delelteVoter() {
	var tempVoterName string
	fmt.Println("Enter Name of Voter to delete: ")
	fmt.Scan(&tempVoterName)
	for i, voter := range voters {
		if tempVoterName == voter.name {
			voters = append(voters[:i], voters[i+1:]...)
			fmt.Printf("Voter '%s' deleted successfully\n", tempVoterName)
			return
		}
	}
	fmt.Printf("Voter '%s' not found\n", tempVoterName)
}

func deleteCandidate() {
	var tempCandidateName string
	fmt.Println("Enter Name of Candidate to delete: ")
	fmt.Scan(&tempCandidateName)
	for i, candidate := range candidates {
		if tempCandidateName == candidate.name {
			candidates = append(candidates[:i], candidates[i+1:]...)
			fmt.Printf("Candidate '%s' deleted successfully\n", tempCandidateName)
			return
		}
	}
	fmt.Printf("Candidate '%s' not found\n", tempCandidateName)
}

func displayAllVoters() {
	fmt.Println("----All Voters----")
	if len(voters) == 0 {
		fmt.Println("No Voters In system")
		return
	}
	for i := 0; i < len(voters); i++ {
		fmt.Printf("[Name:%s --- Age:%d --- Voted:%v ]", voters[i].name, voters[i].age, voters[i].isVoted)
		fmt.Println("")
	}
}

func displayAllCandidates() {
	fmt.Println("----All Candidates----")
	if len(candidates) == 0 {
		fmt.Println("No Candidates In system")
		return
	}
	for i := 0; i < len(candidates); i++ {
		fmt.Printf("[Name:%s --- Age:%d ]", candidates[i].name, candidates[i].age)
		fmt.Println("")
	}

}

func showVotingResult() {
	if len(candidates) == 0 {
		fmt.Println("No Candidates In system")
		return
	}
	// Initialize variables to keep track of the candidate with the maximum votes
	maxVotes := candidates[0].votes
	maxCandidate := candidates[0]

	// Iterate through the candidates to find the one with maximum votes
	for _, candidate := range candidates {
		if candidate.votes > maxVotes {
			maxVotes = candidate.votes
			maxCandidate = candidate
		}
	}

	fmt.Println("---Voting Results---")
	for i := 0; i < len(candidates); i++ {
		fmt.Printf("[Name:%s --- Age:%d --- Votes:%d ]", candidates[i].name, candidates[i].age, candidates[i].votes)
		fmt.Println("")
	}

	fmt.Printf("\n\nThe Winner of the Election is %s with %d votes", maxCandidate.name, maxCandidate.votes)
}

func main() {
	var numCan int
	fmt.Println("Enter  the number of Candidates:")
	fmt.Scan(&numCan)
	fmt.Println("Enter Details of Candidates")
	for i := 0; i < numCan; i++ {
		var temp Candidate
		temp.votes = 0
		fmt.Printf("Enter Name of Candidate %d :\n", i+1)
		fmt.Scan(&temp.name)
		fmt.Printf("Enter Age of Candidate %d :\n", i+1)
		fmt.Scan(&temp.age)
		if temp.age < 21 {
			fmt.Println("Cannot register as Candidate (age<21)")
			i--
			fmt.Println("Re enter Details of Candidate")
			continue
		}
		candidates = append(candidates, temp)
	}

	var choice int = 0

	for choice != 6 {
		fmt.Println("\n\nEnter Choice")
		fmt.Println("1.Register New Voter")
		fmt.Println("2.Display All Voters")
		fmt.Println("3.Delete Voter")
		fmt.Println("4.Delete Candidate")
		fmt.Println("5.Display All Candidates")
		fmt.Println("6.End Election")
		fmt.Scan(&choice)
		if choice == 1 {
			registerVoter()
		}
		if choice == 2 {
			displayAllVoters()
		}
		if choice == 3 {
			delelteVoter()
		}
		if choice == 4 {
			deleteCandidate()
		}
		if choice == 5 {
			displayAllCandidates()
		}

	}
	showVotingResult()
	fmt.Println("Thanku For Using Golang Voting System")
}
