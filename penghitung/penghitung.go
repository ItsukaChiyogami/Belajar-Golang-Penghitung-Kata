package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Masukkan  kata atau kalimat untuk dihitung (tidak boleh mengandung angka):")

	reader := bufio.NewReader(os.Stdin)
	jawaban, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Terjadi kesalahan saat membaca input.")
		return
	}

	jawaban = strings.TrimSpace(jawaban)

	if !isAlphaSpace(jawaban) {
		fmt.Println("Input tidak valid, hanya boleh berisi huruf dan spasi.")
		return
	}

	fmt.Println("Di bawah ini hitungan karakter dari kata yang Anda masukkan:")
	fmt.Println(len(jawaban))
}

func isAlphaSpace(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}
