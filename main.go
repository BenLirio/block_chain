package main

import (
	"fmt"
	"bufio"
	"net/http"
	"os"
	"io"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if _, err := os.Stat("static"); os.IsNotExist(err) {
		fmt.Print("Enter the port you downloaded from: ")
		scanner.Scan()
		fileUrl := fmt.Sprintf("http://localhost:%s/main", scanner.Text())
		resp, err := http.Get(fileUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		err = os.Mkdir("static", 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		out, err := os.Create("static/main")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()
		_, err = io.Copy(out, resp.Body)
	}
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	fmt.Print("Enter the port you would like to host: ")
	scanner.Scan()
	fmt.Println(fmt.Sprintf("http://localhost:%s", scanner.Text()))
	err := http.ListenAndServe(fmt.Sprintf(":%s", scanner.Text()), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
