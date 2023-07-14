# GRAPHQL

## Review
Dengan graphQL secara tidak langsung kita bisa melakukan query data ke server/data source yang kita siapkan. Berbeda dengan restAPI atau protokol yang lain dimana data sudah disediakan pada handler-handler tertentu dan client mengambil data dari server. Pada graphQL kita dapat mengirimkan query data sesuai kebutuhan kita. Intinya pada graphQL kita dapat menyesuaikan data yang ingin kita ambil sehingga proses pengambilan data akan bergantung pada client.

## Alur Kerja
Karena graphQL bergantung pada query yang dikirimkan oleh client, maka persiapan query sesuai bisnis proses adalah hal yang perlu diperhatikan disini. Sehingga alur kerjanya akan menjadi seperti ini:
1. Server menyediakan data dan mengekspos route untuk handler graphQL berjalan
2. Client mengirimkan query menyesuaikan data yang diperlukan
3. Server mengirimkan data sesuai query

## Alur Pembuatan (pada bahasa Go)
1. Buat skema graphQL dengan ekstensi file *.graphqls
2. Siapkan file konfigurasi untuk gqlgen dengan nama gqlgen.yml
3. Compile skema menggunakan [gqlgen](https://github.com/99designs/gqlgen)
4. Setelah compile seharusnya dihasilkan 3 file untuk setiap skema yaitu: 
   1. Model : berisi struktur data sesuai yang kita buat pada file skema 
   2. Exec : berisi executable function untuk menginisialisasi graphQL
   3. Resolver : berisi handler untuk setiap mutasi dan graph yang ada pada file skema (file ini yang akan kita otak atik)
5. Modifikasi setiap handler pada file resolver sesuai dengan bisnis proses
6. Tambahkan default server yang ada pada file Exec pada http server kita untuk menjalankan graphQL