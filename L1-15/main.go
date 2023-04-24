package main

import (
	"bytes"
	"fmt"
	"os"
)

//в данном фрагменте используется глобальная перменная, которая может неожиданно измениться, что приведет к ошибкам.
//лучше исопользовать лоакальную переменную если это возможно.

//Данный фрагмент кода может привести к утечке памяти,
//так как переменная v создает огромную строку,
//а затем сохраняет только первые 100 символов в переменную justString, оставшаяся часть строки "повисает" и занимает память

/*
var justString string

	func someFunc() {
		v := createHugeString(1 << 10)
		justString = v[:100]
	}

	func main() {
		someFunc()
	}
*/
/////////////////////////// корректный пример реализации:
func main() {
	var justString []byte
	//строки в Go иммутабельны, то есть когда мы их изменяем, фактически создаём новую строку.
	//Поэтому, если мы хотим изменить строку, то лучше использовать []byte:
	justString = someFunc()[:100] //выводим первые 100 строк файла, при условии 1 строка = 1 байту
	// (в нашем случае строка из одного сивола)
	fmt.Println(string(justString))
}

func createHugeString(size int, f *os.File) (string, error) {
	buffer := bytes.Buffer{}

	for i := 0; i < size; i++ {
		buffer.WriteString("S")
		_, err := f.Write([]byte("S"))
		if err != nil {
			return "", err
		}
	}
	return buffer.String(), nil
}

func someFunc() []byte {
	f, err := os.Create("huge_string.txt") // для работы с очень большими строками, которые могут не поместиться в памяти,
	// стоит их сохранять ,например, в файл
	if err != nil {
		panic(err)
	}
	defer f.Close()
	str, err := createHugeString(1<<10, f) //создаем строку размером 1024 байта
	// (1 байт = 1 символ, если символ из английского языка)
	if err != nil {
		fmt.Println("Err by createHugeString:.", err)
	}
	fmt.Println("Длина полной строки:", len(str))
	return []byte(str)
}
