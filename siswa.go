package services

//jika nilai <= 40 maka nilai = c, jika nilai >40 && nilai <=70 maka nilai b, jika nilai >70 nilai a

func HitungNilai(nilai int) string {
	if nilai <= 40 {
		return "Nilai Siswa : C"
	} else if nilai > 40 && nilai <= 70 {
		return "Nilai Siswa : B"
	} else {
		return "Nilai Siswa : A"
	}
}
