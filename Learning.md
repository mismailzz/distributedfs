```golang

package main 

import (
	"fmt"
)

type Car struct {
	Model string 
}

type Driver struct {
    Person string
}

type Gallery struct {
	location string
	car Car
	tempChan chan string 
    driver *Driver
}

func (c *Car) CarModel () string{
	return c.Model
}

func NewGallery() *Gallery {
	return &Gallery{
		location: "warsaw, Poland",
		// No Car Initialization 
        // No Chan Initialization - Due to which tempChan Channel is Nil
        // Driver is not intialized and as its a pointer - its Nil
	}
}
func main(){
	g := NewGallery()
	g.car.Model = "HelloWorld" // assigned value here (its also mean we intialized it without constructor)
	fmt.Println(g.car.Model) // This will work
    // g.tempChan <- "Hi Baby" // It will fail because we didnt intialize it and directly using the Nil Channel - PANIC will happen - Whole main program fail
	go func(){ // here main program will not fail but panic will happen here too which you will not see on the screen, as it will happen inside the separate go routine
		g.tempChan <- "Hi Baby"
	}()

    g.driver.Person := "James" // will definitly fail, as driver inside Gallery is not intialized = Nil pointer Exception

}

// Lesson: 
// 1. Please initialize the attributes of the objects before to use them 
// 2. Go routine crash doesnt cause whole program crash
// 3. Nil Pointer exception due to missing intialization 
```