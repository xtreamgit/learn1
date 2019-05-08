// Author: Hector DeJesus
// May 7, 2019
// Reference: Goland Methods, Interfaces and Composition Video from Learn to Code
// https://greatcommons.com Todd McCleod
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
	fmt.Println("I'm an Artificial Intelligence Robot. My first name is", r.firstName, "and my last name is", r.lastName)
}

func (fr fantasticRobot) speak() {
	fmt.Println("Hey, I'm a fantastic robot. My name is", fr.firstName, ".",
		"I am", fr.bodyShape, "and you can recognize me by my bright", fr.bodyColor, "color.")
}

type nonhuman interface {
	speak()
}

func saySomething(h nonhuman) {
	h.speak()
}

func main() {
	robot1 := robot{
		"htrom2",
		"weekie",
	}

	// When I want the robot to speak, I call the method "speak()" with the object of type robot.
	// In this case the object is robot1.
	saySomething(robot1)

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
	saySomething(fantasticRobot1) // Statement using the interface.
	//fantasticRobot1.speak() // Original way
	//fantasticRobot1.robot.speak() //Specific way - this yields a the output from the robot method.
}

// Interesting issue: If I comment out the fantasticRobot speak method function, the statement fantastic.speak() still works.
// Also interesting is that even if I don't comment it out, I can write this statement fantasticRobot1.robot.speak()
// and it will also work.
