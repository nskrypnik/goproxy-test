package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var times [10]uint64
		latencyTimes := ""
		for i := 1; i < 11; i++ {
			headerKey := "PROXY_" + fmt.Sprint(i)
			t, _ := strconv.ParseUint(r.Header.Get(headerKey), 10, 64)
			times[i-1] = t
		}

		for i := 9; i > 0; i-- {
			latencyTimes += " " + fmt.Sprint(times[i]-times[i-1])
		}
		// w.Header().Set("PROXY_1", r.Header.Get("PROXY_1"))
		// w.Header().Set("PROXY_2", r.Header.Get("PROXY_2"))
		// w.Header().Set("PROXY_3", r.Header.Get("PROXY_3"))
		// w.Header().Set("PROXY_4", r.Header.Get("PROXY_4"))
		// w.Header().Set("PROXY_5", r.Header.Get("PROXY_5"))
		fmt.Fprintf(w, "Latency times: %s\n", latencyTimes)
	})

	http.ListenAndServe(":1330", nil)
}
