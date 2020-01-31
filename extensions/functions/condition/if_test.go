package condition

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var c = &fnIf{}

func TestStaticFunc_Eq(t *testing.T) {
	final1, _ := c.Eval(true, "tibcoTrue", "tibcoFalse")
	assert.Equal(t, "tibcoTrue", final1)

	final2, _ := c.Eval(false, "tibcoTrue", "tibcoFalse")
	assert.Equal(t, "tibcoFalse", final2)

	final3, _ := c.Eval(true, 10, "tibcoFalse")
	assert.Equal(t, 10, final2)

	final3, _ := c.Eval(false, 10, "tibcoFalse")
	assert.Equal(t, "tibcoFalse", final2)
}
