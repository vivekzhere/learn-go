package conts

import (
	"fmt"
)

func ShowArray() {
	arr := []string{
		"postgres",
		"mongo",
		"hana",
		"redis",
	}
	printCont(arr)
	arr = append(arr, "rabbitmq")
	printCont(arr)
}

func ShowHash() {
	hash := map[string]string{
		"manager":   "pranav",
		"seniordev": "anit",
	}
	printCont(hash)

	if v, ok := hash["seniordev"]; !ok {
		hash["seniordev"] = "anit"
	} else {
		hash["seniordev"] = fmt.Sprintf("%s%s", v, "wa")
	}

	printCont(hash)
}

func printCont(cont interface{}) {
	switch cont.(type) {
	case []string:
		printArray(cont.([]string))
	case map[string]string:
		printHash(cont.(map[string]string))
	}
}

func printArray(arr []string) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printHash(hash map[string]string) {
	for k, v := range hash {
		fmt.Println(k, v)
	}
}
