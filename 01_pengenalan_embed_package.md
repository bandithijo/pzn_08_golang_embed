# Pengenalan Embed Package

- Sejak Go versi 1.16, terdapat package baru dengan nama embed
- Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan isi file nya dalam variable
- <https://pkg.go.dev/embed>


# Cara Embed File

- Untuk melakukan embed file ke variable, kita bisa mengimport package embed terlebih dahulu
- Selanjutnya kita bisa tambahkan komentar //go:embed diikuti dengan nama filenya, di atas variable yang kita tuju
- Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara otomatis ketika kode Go dicompile
- Variable yang dituju tidak bisa disimpan di dalam function
