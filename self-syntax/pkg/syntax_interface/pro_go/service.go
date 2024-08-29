package main

import "fmt"

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
}

// type method
func (s *Service) printServiceDetails() {
	fmt.Println("Service Name:", s.description, "Service Price:", s.monthlyFee*float64(s.durationMonths))
}

// Methods implementing interface

func (s Service) getName() string {
	return s.description
}

func (s Service) getCost(recur bool) float64 {
	if recur {
		return s.monthlyFee * float64(s.durationMonths)
	}
	return s.monthlyFee
}
