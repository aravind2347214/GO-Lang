// You're tasked with creating a program for a highly sensitive government agency that
// tracks top-secret missions. The program must manage a collection of mission data,
// including mission names, objectives, and statuses. However, there is a catch: the
// program must be designed to self-destruct (trigger a panic) if anyone tries to access
// a mission labeled with a certain codeword. This requirement is in place to prevent
// unauthorized access to missions of extreme importance.
// Your task is to design and implement this program in Go, ensuring that:
// 1. Mission data is stored securely and can only be accessed through proper
// channels.
// 2. The program includes error handling mechanisms to gracefully handle any
// unexpected issues.
// 3. A panic is triggered if the user tries to access a mission labeled with the
// specified codeword, but the program should also include a recovery
// mechanism to prevent the entire program from crashing.
// 4. Defer statements are used appropriately to clean up resources and ensure
// the program runs smoothly under normal conditions.

package main

import (
	"fmt"
	"os"
)

// Mission represents a top-secret mission.
type Mission struct {
	Name      string
	Objective string
	Status    string
}

// Codeword to trigger self-destruction.
const codeword = "topsecret"

// AccessMission retrieves mission details.
func AccessMission(missionName, userCodeword string) (*Mission, error) {
	if userCodeword != codeword {
		return nil, fmt.Errorf("Unauthorized access: Incorrect codeword")
	}

	// Simulate fetching mission data from a secure source.
	// Replace this with actual data retrieval logic.
	mission := &Mission{
		Name:      missionName,
		Objective: "Classified",
		Status:    "Active",
	}

	return mission, nil
}

func main() {

	var ch string
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Self-destruct initiated! Program terminated.")
			// Perform any necessary cleanup here.
			os.Exit(1)
		}
	}()

	// Example usage:
	missionName := "Operation X"
	userCodeword := "topsecret"

	mission, err := AccessMission(missionName, userCodeword)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Mission: %s\nObjective: %s\nStatus: %s\n", mission.Name, mission.Objective, mission.Status)

	fmt.Print("Enter the code if you know : ")
	fmt.Scan(&ch)
	if ch == codeword {
		fmt.Println("\nAccess granted!")
	} else {
		panic("Unauthorized access detected!")
	}

}
