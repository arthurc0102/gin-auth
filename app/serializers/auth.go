package serializers

import (
	"github.com/arthurc0102/gin-auth/app/utils"
)

// Login serializers
type Login struct {
	Username string `json:"username" binding:"required,max=50,notspace"`
	Password string `json:"password" binding:"required,max=60,notspace"`
}

// Register serializers
type Register struct {
	Username   string `json:"username" binding:"required,max=50,notspace"`
	Password01 string `json:"password01" binding:"required,max=60,notspace,eqfield=Password02"`
	Password02 string `json:"password02" binding:"required,max=60,notspace"`
	FirstName  string `json:"firstName" binding:"required,max=50,notspace"`
	LastName   string `json:"lastName" binding:"required,max=50,notspace"`
	Gender     string `json:"gender" binding:"required,max=1,notspace,choices=F M"`
	Age        uint   `json:"age" binding:"required,min=0"`
}

// Password hash
func (serializer *Register) Password() string {
	return utils.HashPassword(serializer.Password01)
}

// RefreshJWT serializers
type RefreshJWT struct {
	Token string `json:"token" binding:"required,notspace"`
}
