package faculty_test

import (
	"UNIT_TEST/faculty"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get Faculty", func(){
	Describe("Getting Faculty from NIM", func(){
		It("should return an error an invalid NIM length", func(){
			nim := "1234567"
			facultyName, err := faculty.GetFaculty(nim)
			Expect(err). To(HaveOccurred())
			Expect(err.Error()).To(Equal("invalid nim"))
			Expect(facultyName).To(Equal(""))
		})

		It("should return the correct faculty for NIM 22111980", func(){
			nim := "22111980"
			facultyName, err := faculty.GetFaculty(nim)
			Expect(err). To(BeNil())
			Expect(facultyName).To(Equal("Fakultas Teknik Informatika"))
		})

		It("should return the correct faculty for NIM 22121980", func(){
			nim := "22121980"
			facultyName, err := faculty.GetFaculty(nim)
			Expect(err). To(BeNil())
			Expect(facultyName).To(Equal("Fakultas Ekonomi"))
		})

		It("should return the correct faculty for NIM 22131980", func(){
			nim := "22131980"
			facultyName, err := faculty.GetFaculty(nim)
			Expect(err). To(BeNil())
			Expect(facultyName).To(Equal("Fakultas Ilmu sosial dan Politik"))
		})

		It("should return an empty string for an unrecognized faculty code", func(){
			nim := "22141980"
			facultyName, err := faculty.GetFaculty(nim)
			Expect(err). To(BeNil())
			Expect(facultyName).To(Equal(""))
		})
	})
})
