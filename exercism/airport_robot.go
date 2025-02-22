package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Italian implements
type Italian struct {
}

func (italian Italian) LanguageName() string {
	return "Italian"
}

func (italian Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese implements
type Portuguese struct {
}

func (port Portuguese) LanguageName() string {
	return "Portuguese"
}

func (port Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
