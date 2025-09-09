// # Membuat Project
// project di Go-lang, biasanya disebut sebagai module
// Untuk membuat module, kita bisa menggunakan perintah berikut di folder tempat kita akan membuat module :
// go mod init nama-module
// Fitur Go-Lang Modules sebenarnya masih banyak, namun akan kita bahas di kelas khusus membahas tentang Go-Lang Modules


// File go.mod
// file go.mod ini berfungsi untuk mendefinisikan module yang kita buat, versi go yang digunakan, serta dependency yang digunakan dalam project kita.
// file ini akan otomatis dibuat ketika kita menjalankan perintah `go mod init <nama-module>` di terminal.
// secara default file ini akan berisi nama module dan versi go yang digunakan seperti yang bisa kita lihat di atas.