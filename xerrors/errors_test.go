package xerrors

import (
	"errors"
	"testing"
)

func TestNew(t *testing.T) {
	t.Log("Hello BuGuai !!! ")
	t.Log(New("", "", nil).IsNull(), errors.New("") == nil)

	t.Log(Msg("msg", "msgStr").Error())
	t.Log(Msg("msg", "msgStr").Messge())

	t.Log(Data("data", "dataString").Data())
	t.Log(Data("data", 123).Error())
}
