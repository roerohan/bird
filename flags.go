package main

import (
	"flag"
	"strings"

	"github.com/roerohan/bird/logger"
)

// BirdHome stores the home directory where bird is installed
var (
	urls     Urls
	wordlist string
	success  Success
)

// Urls is a type to store multiple
// target URLs to be bruteforced
type Urls []string

// Success stores the status codes
// which represent success
type Success []string

// Set is a function on Urls
// to implement flag.Value
func (u *Urls) Set(value string) error {
	*u = append(*u, value)
	return nil
}

// String is a function on Urls to
// implement flag.Value
func (u *Urls) String() string {
	return strings.Join(*u, " ")
}

// Set is a function on Success
// to implement flag.Value
func (s *Success) Set(value string) error {
	*s = append(*s, value)
	return nil
}

// String is a function on Success to
// implement flag.Value
func (s *Success) String() string {
	return strings.Join(*s, " ")
}

func init() {
	defer logger.Welcome()

	flag.Var(&urls, "u", "Target URL to be bruteforced [required]")
	flag.Var(&success, "s", "Status code for success, default 200")
	flag.StringVar(&wordlist, "w", "", "Path to wordlist [required]")
}
