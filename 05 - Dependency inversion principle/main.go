package main

import "fmt"

/*
	Definition:
	High level modules should not depend on low level modules, both should depend on abstractions.
	Abstractions should not depend on details.  Details should depend upon abstractions.
*/

/*--------------------------------------------------------*/
/*	 BREAKING DEPENDENCY INVERSION PRINCIPLE	          */
/*--------------------------------------------------------*/
type Mailer struct{}

func (Mailer) SendMail() {
	fmt.Println("Sending email welcome message...")
}

// The wrong thing with this is because WelcomeMessage is depending on concrete Mailer struct
type WelcomeMessage struct {
	Mailer
}

/*--------------------------------------------------------*/
/*	 IMPLEMENTING DEPENDENCY INVERSION PRINCIPLE	    */
/*--------------------------------------------------------*/
type IMailer interface {
	SendMyMail()
}

type SmtpMailer struct{}

// IMailer implements IMailer interface
func (SmtpMailer) SendMyMail() {
	fmt.Println("Sending welcome email using SmtpMailer...")
}

type SlackMailer struct{}

// SlackMailer implements IMailer interface
func (SlackMailer) SendMyMail() {
	fmt.Println("Sending welcome email using SlackMailer...")
}

// WelcomeMessageGood is not depending on concrete struct (SmtpMailer, SlackMailer)
// It depends on abstraction (IMailer)..
// So basically we can add more mailing services that implements IMailer
// interface so there is no need to edit WelcomeMessageGood struct in any way
type WelcomeMessageGood struct {
	IMailer
}

func main() {
	// Running bad example
	msg := &WelcomeMessage{Mailer{}}
	msg.SendMail()

	// Running good example
	msg2 := &WelcomeMessageGood{&SmtpMailer{}}
	msg2.SendMyMail()

	msg3 := &WelcomeMessageGood{&SlackMailer{}}
	msg3.SendMyMail()
}
