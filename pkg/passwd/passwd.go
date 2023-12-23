package passwd

import (
	"fmt"

	"github.com/Tom5521/EduTrack/pkg/conf"
	"github.com/ncruces/zenity"
	"golang.org/x/crypto/bcrypt"
)

type Hash string

type Password string

func (p Password) ToHash() (Hash, error) {
	h, err := bcrypt.GenerateFromPassword(p.ToByte(), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return Hash(h), nil
}
func (p Password) ToByte() []byte {
	return []byte(p)
}

func (h Hash) ToByte() []byte {
	return []byte(h)
}
func (h Hash) Compare(p Password) error {
	return bcrypt.CompareHashAndPassword(h.ToByte(), p.ToByte())
}

func GetByte(b GetByter) []byte {
	return b.ToByte()
}

type GetByter interface {
	ToByte() []byte
}

func askpwd() Password {
	_, p, err := zenity.Password()
	if err != nil {
		panic(err)
	}
	return Password(p)
}

func AskPwd() Password {
	c := conf.Config
	if c.Password.Enabled && c.Password.Hash == "" {
		p := askpwd()
		pHash, err := p.ToHash()
		if err != nil {
			errWin(err.Error())
		}
		c.Password.Hash = string(pHash)
		c.Update()
		return p
	}
	p := askpwd()
	return p
}

func errWin(text string) {
	err := zenity.Error(text)
	if err != nil {
		fmt.Println(err)
	}
}
