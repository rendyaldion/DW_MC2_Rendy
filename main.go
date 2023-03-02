package main

import "fmt"

type Profile struct {
	Name string
	Health int
	Power int
	Exp int
}

func MakeProfile(characterName string) Profile {
	profile := Profile{}
	if characterName == "Sasuke" {
		profile.Name = characterName
		profile.Health = 200
		profile.Power = 100
		profile.Exp = 0
	}

	if characterName == "Goku" {
		profile.Name = characterName
		profile.Health = 400
		profile.Power = 300
		profile.Exp = 100
	}

	if characterName == "Naruto" {
		profile.Name = characterName
		profile.Health = 150
		profile.Power = 200
		profile.Exp = 50
	}

	return profile
}

func PowerUp(profile Profile, multiplier int) Profile {
	profile.Health = profile.Health * multiplier
	profile.Power = profile.Power * multiplier
	profile.Exp = profile.Exp * 2

	return profile
}

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name:", profile.Name)
	fmt.Println("Health:", profile.Health)
	fmt.Println("Power:", profile.Power)
	fmt.Println("Exp:", profile.Exp)
	
	fmt.Println("========== HEROES POWER UP ==========")

	powerUp := PowerUp(profile, 2)
	fmt.Println("Name:", powerUp.Name)
	fmt.Println("Health:", powerUp.Health)
	fmt.Println("Power:", powerUp.Power)
	fmt.Println("Exp:", powerUp.Exp)
}