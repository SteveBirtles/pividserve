package main

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"strconv"
)

type Category struct {
	Title string	`json:"title"`
	Videos []string	`json:"videos"`
}

var  (
	films []Category
	videos map[int]string
)


func pathTail(path string) int {
	pathBits := strings.Split(path, "/")
	number, _ := strconv.Atoi(pathBits[len(pathBits)-1])
	return number
}

func serve(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Serve")

	jsonFile, err := os.Open("videos.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &films)

	key := 0

	if w == nil {
		videos = make(map[int]string)
	}

	html := `<html>
    <head>
        <title>MCU Trailers</title>
        <script>
            function get(id) {
                const Http = new XMLHttpRequest();
                const url='/play/' + id;
                Http.open('GET', url);
                Http.send();
            }
        </script>
        <style>
            ul {
                list-style-type: none;
                margin: 0;
                padding: 0;
            }
            .button {
                background-color: #4CAF50;
                border: none;
                color: white;
                padding: 15px 32px;
                margin: 5px;
                text-align: center;
                text-decoration: none;
                display: inline-block;
                font-size: 16px;
            }
        </style>
    </head>
    <body>
        <h1>MCU Trailers</h1>
        <ul>`

	for _, film := range films {
		html += fmt.Sprintf("            <li><h3>%s</h3>\n", film.Title)
		html += fmt.Sprint("                <div>")
		for _, trailer := range film.Videos {
			key++
			if w == nil {
				videos[key] = trailer
			}
			html += fmt.Sprintf("<button onclick='get(%d)' class='button'>%s</button>", key, trailer)
		}
		html += fmt.Sprintln("</div>")
		html += fmt.Sprintln("            </li>")
	}
	html += fmt.Sprintln("        </ul>")
	html += fmt.Sprintln("    </body>")
	html += fmt.Sprintln("</html>")

	if w != nil {
		fmt.Fprint(w, html)
	}

}

func play(w http.ResponseWriter, r *http.Request) {

	id := pathTail(r.URL.Path)

	fmt.Println("Play:", id, videos[id])

	if id < 0 || id >= len(videos) {

		/*cmd := exec.Command("./start.sh", "videos/" + videos[id])
		err := cmd.Start()

		if err != nil {
			fmt.Printf("%s", err)
		}*/

	}

	w.WriteHeader(http.StatusOK)

}


func stop(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Stop")

	/*cmd := exec.Command("./stop.sh")
	err := cmd.Start()

	if err != nil {
		fmt.Printf("%s", err)
	}*/


	w.WriteHeader(http.StatusOK)
}

func main() {

	serve(nil, nil)	//setup map

	http.HandleFunc("/index.html", serve)
	http.HandleFunc("/play/", play)
	http.HandleFunc("/stop", stop)

	err := http.ListenAndServe(":8081", http.DefaultServeMux)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
