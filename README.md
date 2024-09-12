# Day 6: Penghitung Karakter dengan Golang

Selamat datang di repository saya! Pada hari keenam, saya membuat sebuah program sederhana dalam Golang yang menghitung jumlah karakter dalam sebuah kata atau kalimat. Program ini memeriksa input pengguna untuk memastikan bahwa hanya huruf dan spasi yang dimasukkan, tanpa angka atau karakter khusus. Mengingat laptop saya sedang diperbaiki, saya memilih untuk mengerjakan proyek sederhana yang tetap memperdalam pemahaman saya tentang pemrosesan input dan validasi data.

## Apa yang Dipelajari

Dalam proyek ini, saya mempelajari:

1. **Penggunaan Package `bufio` dan `os`**: Mempelajari cara membaca input pengguna melalui terminal menggunakan `bufio.NewReader`.
2. **String Manipulation**: Menggunakan fungsi `strings.TrimSpace` untuk menghapus spasi di awal dan akhir input.
3. **Validasi Input**: Memastikan bahwa input hanya berisi huruf dan spasi menggunakan fungsi `unicode.IsLetter` dan `unicode.IsSpace` dari package `unicode`.
4. **Penghitungan Karakter**: Menghitung jumlah karakter dalam input yang valid menggunakan fungsi `len`.

## Struktur Proyek
└── penghitung.go
