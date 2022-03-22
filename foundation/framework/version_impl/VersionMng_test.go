package version_impl

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	var v = false

	vmng, err := NewVersionMng("v2.2.2")
	if err != nil {
		fmt.Println(err)
		return
	}
	vmng.SetClientVersion("version2.2.2")

	vmng.CheckKingdomVersion("v2.2.1")
	fmt.Println(v)
	vmng.CheckKingdomVersion("v2.2.2")
	fmt.Println(v)
	vmng.CheckKingdomVersion("v2.2.2.3")
	fmt.Println(v)
	vmng.CheckKingdomVersion("v2.2.4.2")
	fmt.Println(v)

}
