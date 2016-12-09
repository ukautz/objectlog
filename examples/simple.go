package main

import (
	"fmt"
	"github.com/ukautz/objectlog"
)

var (
	// logger uses Go's built-in "log" package - any logging implementation would do (can be adapted)
	logger = objectlog.NewStandardLogger()
)

type (

	// Person is a demo struct to showcase transparent logging decoration
	Person struct {

		// ObjectLog used directly
		objectlog.ObjectLog

		// Name is an arbitrary demo attribute
		Name string
	}

	// Company is a demo struct to showcase logging decoration
	Company struct {

		// reference to ObjectLog
		*objectlog.ObjectLog

		// Name is an arbitrary demo attribute
		Name string
	}
)

// NewPerson constructor showcases how to decorate an object transparently
func NewPerson(name string) *Person {

	// generate log instance, to decorate newly created person
	log := objectlog.NewObjectLog(logger).SetLogPrefix(fmt.Sprintf("[Person: %s] ", name))

	// generate new Person with a logger
	return &Person{
		ObjectLog: *log,
		Name:      name,
	}
}

func main() {

	// create person, decorated with logging
	person := NewPerson("Mr. Foo")

	// log from person
	// 2016/12/20 20:50:59 [INFO] [Person: Mr. Foo] Hello! I am created
	person.LogInfo("Hello! I am created")

	// create company and init logging
	log := objectlog.NewObjectLog(logger).SetLogPrefix("[Company: ACME Inc] ")
	company := &Company{
		ObjectLog: log,
		Name:      "ACME Inc",
	}

	// log from company
	// 2016/12/20 20:50:59 [INFO] [Company: ACME Inc] We build token tokens.
	company.LogInfo("We build token tokens.")
}
