package main

// Creating a two page website application
import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// In other for a function to respond to a request from a web browser;
// It has to handle two parameters;
// A response writer called (w http.ResponseWriter, r *http.Request)
// and a request r *http.Request

// Home is the  home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")

}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	// %d is the placeholder for integer
	// _ means to ignore
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))

}

// addValues adds rwo integers and returns sum
func addValues(x, y int) int {

	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err :=divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}
    // %f the placeholder for float
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))

}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)






	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
