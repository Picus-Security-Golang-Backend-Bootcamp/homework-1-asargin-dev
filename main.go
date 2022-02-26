package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args

	books := []string{
		"A Tale of Two Cities",
		"The Hobbit",
		"The Little Prince",
		"The Alchemist",
		"Harry Potter and the Chamber of Secrets",
	}

	if len(args) > 1 {

		switch args[1] {

		case "list":

			list_book(books)

		case "search":
			// search argümanı sonrası girilececek değerleri Join ederek string bir arama querysi oluşturarak
			// kitaplarla beraber search_book fonksiyonumuza parametre olarak gönderiyoruz.
			search_book(strings.Join(args[2:], " "), books)

		}

	} else {

		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", "Kitaplık")

	}

}

func list_book(books []string) {

	if len(books) == 0 {

		fmt.Println("Kitaplıkta kitap bulunmamaktadır.")

	} else {

		for _, book := range books {

			fmt.Println(book)

		}

	}

}

func search_book(search_params string, books []string) {

	if len(search_params) == 0 {

		fmt.Println("Lütfen aramak istediğiniz kitabın adını yazınız.")

	} else {
		// Sonuçlar liste şeklinde istenildiğinden, İşlemler sonrası eşleşen sonuçları 'result' slice'ının içine append edeceğiz.
		result := []string{}

		for _, book := range books {
			//Büyük-küçük harf duyarlılığını kaldırmak için hem kitapları hem de arama parametrelerini lowercase'e çeviriyoruz.
			if strings.Contains(strings.ToLower(book), strings.ToLower(search_params)) {
				result = append(result, book)
			}
		}

		if len(result) == 0 {
			fmt.Println("Aradığınız kriterde kitap bulunamamıştır.")
		} else {
			fmt.Println(result)
		}

	}
}
