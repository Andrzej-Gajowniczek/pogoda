package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type WeatherData struct {
	Stacja      string `json:"stacja"`
	Temperatura string `json:"temperatura"`
	Cisnienie   string `json:"cisnienie"`
	Zmierzono   string `json:"godzina_pomiaru"`
}

func main() {
	for {
		response, err := http.Get("https://danepubliczne.imgw.pl/api/data/synop")
		if err != nil {
			fmt.Println("Błąd podczas pobierania danych:", err)
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Błąd podczas odczytu odpowiedzi:", err)
			return
		}

		var data []WeatherData

		if err := json.Unmarshal(body, &data); err != nil {
			fmt.Println("Błąd podczas dekodowania JSON:", err)
			return
		}

		czytelnik := bufio.NewReader(os.Stdin)
		var znaleziono bool
		var miasto string
		//for {
		fmt.Print("\npodaj miasto:")
		miasto, _ = czytelnik.ReadString('\n')
		miasto = miasto[:len(miasto)-1]
		//fmt.Scanf("%v", &miasto)
		//fmt.Println("to jest to miasto:", miasto)
		znaleziono = false
		for _, entry := range data {
			//fmt.Println(entry.Stacja)
			if entry.Stacja == miasto {
				fmt.Printf("temperatura: %v\ncisnienie: %v\ngodzinia pomiaru:%v:00\n",
					entry.Temperatura,
					entry.Cisnienie,
					entry.Zmierzono)
				znaleziono = true
			}
		}
		if !znaleziono {
			fmt.Println("Brak danych na temat miasta:", miasto)
			fmt.Println("Wybierz z poniższych:")
			for _, nazwa := range data {
				fmt.Print(nazwa.Stacja, ",")
			}

		}
	}
}
