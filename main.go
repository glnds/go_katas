package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

func logRequest(req *http.Request) {
	dump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))
	// log.Debug(string(dump))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	keys, ok := r.URL.Query()["q"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'q' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key := keys[0]
	clean := strings.Replace(key, ",", "", -1)
	fmt.Println(key)
	words := strings.Fields(clean)
	fmt.Println(words, len(words))
	if len(words) == 6 {
		i1, _ := strconv.Atoi(words[3])
		i2, _ := strconv.Atoi(words[5])
		s := strconv.Itoa(i1 + i2)
		fmt.Println(s)
		w.Write([]byte(s))
	} else if len(words) == 7 {
		i1, _ := strconv.Atoi(words[3])
		i2, _ := strconv.Atoi(words[6])
		s := strconv.Itoa(i1 * i2)
		fmt.Println(s)
		w.Write([]byte(s))
	} else if len(words) == 11 {
		i1, _ := strconv.Atoi(words[9])
		i2, _ := strconv.Atoi(words[10])
		var b string
		if i1 > i2 {
			b = strconv.Itoa(i1)
		} else {
			b = strconv.Itoa(i2)
		}
		fmt.Println(b)
		w.Write([]byte(b))
	} else if len(words) == 15 {
		i1, _ := strconv.Atoi(words[13])
		i2, _ := strconv.Atoi(words[14])
		s := ""
		if squareCheck(i1) && cubicCheck(i1) {
			s = strconv.Itoa(i1)
		} else if squareCheck(i2) && cubicCheck(i2) {
			s = strconv.Itoa(i2)
		}
		fmt.Println(s)
		w.Write([]byte(s))
	} else if len(words) == 17 {
		arr := words[13:]
		max := 0
		for _, value := range arr {
			i, _ := strconv.Atoi(value)
			if squareCheck(i) && cubicCheck(i) {
				max = i // found another smaller value, replace previous value in max
			}
		}
		s := strconv.Itoa(max)
		fmt.Println(s)
		if max == 0 {
			w.Write([]byte(""))
		} else {
			w.Write([]byte(s))
		}
	} else {
		arr := words[9:]
		max := 0
		for _, value := range arr {
			i, _ := strconv.Atoi(value)
			if i > max {
				max = i // found another smaller value, replace previous value in max
			}
		}
		s := strconv.Itoa(max)
		fmt.Println(s)
		w.Write([]byte(s))
	}
}

func squareCheck(a int) bool {
	var intRoot int = int(math.Sqrt(float64(a)))
	return (intRoot * intRoot) == a
}

func cubicCheck(a int) bool {
	var intRoot int = int(math.Cbrt(float64(a)))
	return (intRoot * intRoot * intRoot) == a
}

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
