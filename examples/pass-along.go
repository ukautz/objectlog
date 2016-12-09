package main

import (
	"fmt"
	"github.com/ukautz/objectlog"
)

type (

	// Brand is a demo parental object, decorated by ObjectLog
	Brand struct {
		*objectlog.ObjectLog
		Name string
	}

	// Car is a demo child object, which will inherit it's parent ObjectLog
	Car struct {
		*Brand
		*objectlog.ObjectLog
		Model string
	}
)

var (
	logger = objectlog.NewStandardLogger()
)

// NewBrand showcases creation of a parental object, which is decorated by ObjectLog
func NewBrand(name string) *Brand {
	return &Brand{
		ObjectLog: objectlog.NewObjectLog(logger).SetLogPrefix(fmt.Sprintf("Brand(%s): ", name)),
		Name:      name,
	}
}

// NewCar showcases how to inherit a parent ObjectLog
func (this *Brand) NewCar(model string) *Car {
	return &Car{
		ObjectLog: this.LogCloneObjectLog().SetLogPrefix(fmt.Sprintf("%sModel(%s): ", this.LogPrefix(), model)),
		Brand:     this,
		Model:     model,
	}
}

func main() {
	brand1 := NewBrand("DeLorean")
	car1 := brand1.NewCar("DMC-12")

	brand2 := NewBrand("Ferrari")
	car2 := brand2.NewCar("F-40")

	// writes "Brand(DeLorean): Model(DMC-12): Wrumm, Wrumm"
	car1.LogInfo("Wrumm, Wrumm")

	// writes "Brand(Ferrari): Model(F-40): Roarr"
	car2.LogInfo("Roarr")
}
