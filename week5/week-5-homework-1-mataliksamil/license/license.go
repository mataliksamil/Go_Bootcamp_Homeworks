package main

import (
	"fmt"
)

// interface just for the allocation operation
type LicenseAllocation interface {
	Allocate()
}

// software struct in order to differentiate different softwwares and licenses
type Software struct {
	Name string
}

// License struct includes software since it has to defined for a software product
type License struct {
	LicenseAvailable int
	SoftwareInstance Software
}

// real allocation function which reduces the license seats
func (l *License) Allocate() {
	l.LicenseAvailable = l.LicenseAvailable - 1
	fmt.Printf(" One %#v allocated, %#v License remained\n", l.SoftwareInstance.Name, l.LicenseAvailable)
}

// proxy license in order to interface to proxy allocate function
type LicenseProxy struct {
	license *License
}

// proxy allocation function in order to check if there are enough seat beforehand
func (l *LicenseProxy) Allocate() {
	if l.license.LicenseAvailable > 0 {
		l.license.Allocate()
	} else {
		// error handling
		fmt.Printf(" No license available for %v please wait or try another time\n", l.license.SoftwareInstance.Name)
	}
}

// creates and returns proxy license
func newLicenseProxy(license *License) *LicenseProxy {
	return &LicenseProxy{license}
}

// creates and returns real license (this is just for the sake of the simplicty)
func NewLicense(s string, i int) *LicenseProxy {
	L := License{i, Software{s}}
	return newLicenseProxy(&L)
}

// print for checking the state
func (l *LicenseProxy) PrintLicense() {
	fmt.Printf(" There are %#v License in %#v Software \n", l.license.LicenseAvailable, l.license.SoftwareInstance.Name)
}

func main() {
	fmt.Println(" Hello World !")
	license1 := NewLicense("samil soft", 12)
	license2 := NewLicense("akın soft", 2)

	license1.Allocate()
	license1.Allocate()
	license1.PrintLicense()
	license2.Allocate()
	license2.Allocate()
	license2.Allocate()
	license2.Allocate()
	license2.PrintLicense()

}
