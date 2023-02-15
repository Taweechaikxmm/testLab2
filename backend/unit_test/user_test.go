package main

import (
	"testing"

	"github.com/Taweechaikxmm/testLab2/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	g := NewGomegaWithT(t)
	u := entity.User{
		Address: "",
	}
	ok, err := govalidator.ValidateStruct(u)
	g.Expect(ok).NotTo(BeTrue())
	g.Expect(err).NotTo(BeNil())
	g.Expect(err.Error()).To(Equal("test"))
}
