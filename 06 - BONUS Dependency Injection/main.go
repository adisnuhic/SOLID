package main

import (
	"net/http"

	"fmt"

	"io/ioutil"
)

/*
	Definition:
	The literal meaning is to inject dependency. Dependency is just another object
	that your class needs in order to function. You should not instantiate dependency in your
	class, instead take it as a constructor parameter. It decouples your class/struct
	construction from construction of your dependency.

	Dependency Injection as a form of inversion of control. Inversion of Control (IoC) means that objects do not create
	other objects on which they rely to do their work. Instead, they get the objects that they need from an outside source.
*/

/*--------------------------------------------------------*/
/*		 				WRONG WAY	          		    */
/*--------------------------------------------------------*/
type Print struct{}

func (Print) PrintData() {
	// The wrong thing with this is because we are using http.Client inside the function.
	// Since http.Client is going to make real requests, our test would have to rely on a valid internet connection
	// Most of the machines will have that ability but what if this is a DB connection?
	// We need to be able to mock these connections and our code is not currently mockable
	client := http.Client{}
	response, err := client.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	if response == nil {
		fmt.Println("Received empty response")
		return
	}

	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		panic(errRead)
	}

	fmt.Println(string(body))

}

/*--------------------------------------------------------*/
/*		 				GOOD WAY	          		    */
/*--------------------------------------------------------*/
// http.Client will implicitly satisfy our IHttpClient interface because they both have the same signature - Get methods
type IHttpClient interface {
	Get(string) (*http.Response, error)
}

type PrintGood struct{}

// We can easily mock IHttpClient.
// So our PrintDataGood do not control cretion of http Client, it is controlled outside of the function (IoC)
func (PrintGood) PrintDataGood(client IHttpClient, url string) {
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	if response == nil {
		fmt.Println("Received empty response")
	}

	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		panic(errRead)
	}

	fmt.Println(string(body))
}

/*
	Notice: Here in function PrintDataGood i am using interface instead of http.Client struct so we
	satisfy Dependency Inversion principle.. But i could pass http.Client also and it will still be Dependency Injection
*/

func main() {
	// Running bad example
	p := &Print{}
	p.PrintData()
	fmt.Println("------------------------------------------------------------------------")

	// Running good example
	p2 := &PrintGood{}
	client := http.Client{}
	p2.PrintDataGood(&client, "http://google.com")

}
