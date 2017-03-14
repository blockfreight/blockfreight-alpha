package bf_tx

import (
	"testing"
	"reflect"
)

func TestSetBF_TX(t *testing.T) {
	t.Log("Test on SetBF_TX function")
	var prot BF_TX
	bf_tx := SetBF_TX("../.././files/bf_tx_example.json")
	if reflect.TypeOf(bf_tx) != reflect.TypeOf(prot) {
    	t.Error("Error on type of result of SetBF_TX")
	}
}