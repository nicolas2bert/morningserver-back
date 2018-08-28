package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openBrowserTab(w http.ResponseWriter, r *http.Request) {
	fmt.Println("START!!!")
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("b!!!", b)
	var cmd *exec.Cmd
	const url = "http://www.youtube.com"
	log.Printf("Opening youtube in %v environemnt", runtime.GOOS)
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		log.Fatalf("unsupported platform")
	}
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// https://www.chromium.org/developers/design-documents/process-models
func killChromium(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("cmd!!!", cmd)
	// if cmd == nil {
	// 	log.Println("Process does not exist")
	// 	return
	// }
	// err := cmd.Process.Kill()
	// fmt.Println("err!!!", err)
	// if err != nil {
	// 	log.Fatalf("Unable to kill process")
	// }
	var err error
	var cmd *exec.Cmd
	cmd = exec.Command("pkill", "chromium")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/openbrowsertab", openBrowserTab)
	http.HandleFunc("/kill", killChromium)
	log.Printf("Server is listing on port 5050")
	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		panic(err)
	}
}
