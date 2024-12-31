package main

import (
	"container/list"
	"container/ring"
	"flag"
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)
type People struct {
	Name string
	Age int
}

type UserSlice []People

func (val UserSlice) Len() int {
	return len(val)
}

func (val UserSlice) Less(i, j int) bool {
	return val[i].Age < val[j].Age
}

func (val UserSlice) Swap(i, j int) {
	val[i], val[j] = val[j], val[i]
}

func main() {
	// OS
	args := os.Args
	fmt.Println(args)

	// OS (HOSTNAME)
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error:", err.Error())
	}else {
		fmt.Println("Hostname:", hostname)
	}

	// OS (ENV)
	getUsername := os.Getenv("USERNAME")
	getPassword := os.Getenv("PASSWORD")

	fmt.Println(getUsername)
	fmt.Println(getPassword)

	// FLAG
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 100, "Put your number")

	flag.Parse()
	
	fmt.Println("Host:", *host)
	fmt.Println("Username:", *username)
	fmt.Println("Password:",*password)
	fmt.Println("Number:",*number)

	// Strings
	fmt.Println(strings.Contains("Muhammad Nabil Wafi", "Nabil"))
	fmt.Println(strings.Split("Muhammad Nabil Wafi", " "))
	fmt.Println(strings.ToLower("Muhammad Nabil Wafi"))
	fmt.Println(strings.ToUpper("Muhammad Nabil Wafi"))
	fmt.Println(strings.Trim("i     Muhammad Nabil Wafi      i", ""))
	fmt.Println(strings.ReplaceAll("Nabil Wafi Nabil", "Nabil", "Muhammad"))

	// StrConv
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(boolean)
	}

	integer, err := strconv.ParseInt("100000", 10, 64)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(integer)
	}

	val, err := strconv.Atoi("10000")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println(val)
	}

	val2 := strconv.Itoa(10000)
	fmt.Println(val2)

	// Math
	fmt.Println(math.Ceil(1.4))
	fmt.Println(math.Floor(1.5))
	fmt.Println(math.Round(1.5))
	fmt.Println(math.Max(1,5))
	fmt.Println(math.Min(1,5))

	// List
	data := list.New()
	data.PushBack("Muhammad")
	data.PushBack("Nabil")
	data.PushBack("Wafi")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Container
	var data2 *ring.Ring = ring.New(5)
	for i := 0; i < data2.Len(); i++ {
		data2.Value = "Value " + strconv.FormatInt(int64(i), 10)
		data2 = data2.Next()
	}

	fmt.Println(data2)
	data2.Do(func (val interface{}) {
		fmt.Println(val)
	})

	// SORT
	users := UserSlice{
		{"Nabil", 15},
		{"Ikhsan", 20},
		{"Jani", 10},
	}

	sort.Sort(users)
	fmt.Println(users)

	// TIME
	now := time.Now()
	fmt.Println(now.Day())
	fmt.Println(now.Year())
	fmt.Println(now.Month())

	// REFLECT
	sample := reflect.TypeOf("Nbail")
	fmt.Println(sample)

	// Regex
	var regex = regexp.MustCompile(`n([a-z])l`)
	fmt.Println(regex.MatchString("nabil"))
	fmt.Println(regex.MatchString("nobil"))
}