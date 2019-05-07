// Author: Hector DeJesus
// May 7, 2019
package main

import "fmt"

// In this go program I will learn
// composition
// embedded types
// inner-type promotion
// methods
// interfaces

type robot struct {
	firstName string
	lastName  string
}

type fantasticRobot struct {
	robot
	workArea   string
	bodyColor  string
	walkFreely bool
	bodyShape  string
}

// This is the speak method for the robot type
// somerobot is the "receiver" for the method
// Anything of type robot, will be able to use the speak() method.
func (r robot) speak() {
	fmt.Println("I'm an Artificial Inteligence Robot. My first name is", r.firstName, "and my last name is", r.lastName)
}

func (fr fantasticRobot) speak() {
	fmt.Println("Hey, I'm a fantastic robot. My name is", fr.firstName)
}

func main() {
	robot1 := robot{
		"htrom2",
		"weekie",
	}

	// When I want the robot to speak, I call the method "speak()" with the object of type robot.
	// In this case the object is robot1.
	robot1.speak()

	////////////////////////////////////////////

	fantasticRobot1 := fantasticRobot{
		robot{
			"Brim5",
			"blinto",
		},
		"museum",
		"oragepeel",
		false,
		"tall",
	}
	fantasticRobot1.speak()
}
