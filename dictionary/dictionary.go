package dictionary

import "fmt"

func Dictionary() {

	dict := map[string]string{"book": "kitap", "table": "masa", "dictionary": "sözlük"}
	for k := range dict {
		fmt.Println(k,":", dict[k])
	}

}