package faculty

import (
	"errors"
)

func GetFaculty(nim string) (string, error){
	if len(nim) != 8 {
		return "", errors.New("invalid nim")
	}

	code := nim[2:4]

	facultyMap := map[string]string{
		"11" : "Fakultas Teknik Informatika",
		"12" : "Fakultas Ekonomi",
		"13" : "Fakultas Ilmu Sosial dan Politik",
	}

	if faculty, exists := facultyMap[code]; exists{
		return faculty, nil
	}
	return "", nil
}