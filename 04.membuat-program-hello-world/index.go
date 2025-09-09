// # Main Function
// Go-Lang, itu mirip seperti bahasa pemograman C / C++, dimana perlu ada yang namanya main function
// Main function adalah sebuah fungsi yang akan dijalankan ketika program berjalan
// Untuk membuat function, kita bisa menggunakan kata kunci func
// Main function harus terdapat didalam main package
// Titik koma di Golang, tidaklah wajib, artinya kita bisa menambahkan titik koma atau tidak, diakhir kode program kita (mirip seperti bahasa pemograman Javascript, tidak wajib tapi boleh ditambahkan).
// Pada bahasa pemograman Go-Lang, masih berlaku aturan terkait case sensitive, artinya huruf besar dan huruf kecil itu berbeda


// Println
// Untuk menulis tulisan, kita perlu melakukan import module fmt (singkatan dari format) terlebih dahulu
// Mirip seperti bahasa pemograman Java, kita perlu mengimport package terlebih dahulu sebelum menggunakannya
// Materi terkait import, akan kita bahas di bagian tersendiri


// Bagaimana cara menjalankannya?
// Pada bahasa pemograman Go / Go-Lang, kita bisa melakukan build terlebih dahulu menggunakan perintah go build <nama_file>.go
// Ketika perintah build dijalankan, maka dia akan mencari main function nya di file manapun (pada contoh studi kasus ini misalnya di file helloworld.go), selanjutnya dia akan coba mengcompile nya menjadi sebuah file executable (.exe) yang bisa langsung dijalankan
// selanjutnya kita bisa menjalankan file executable tersebut melalui terminal / command prompt dengan mengetikkan nama file nya saja, misal untuk studi kasus kali ini adalah "belajar-golang-dasar" lalu tekan enter


// Menjalankan tanpa kompilasi
// Selain cara diatas, kita juga bisa menjalankan program Go-Lang tanpa perlu melakukan kompilasi terlebih dahulu
// Cara seperti ini cocok untuk kita yang masih dalam tahap belajar, atau pada saat proses development program
// Caranya adalah dengan menggunakan perintah go run <nama_file>.go