# Hasil Embed Dicompile

- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca disimpan dalam binary file Go nya
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file Go sudah dicompile, kita tidak butuh lagi file external nya dan bahkan jika diubah file externalnya, isi variable nya tidak akan berubah lagi
