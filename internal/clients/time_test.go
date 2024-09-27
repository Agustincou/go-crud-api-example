package clients

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSystemNow(t *testing.T) {

	systemTime := NewSystemTime()

	time1 := systemTime.Now()

	time.Sleep(1 * time.Microsecond)

	time2 := systemTime.Now()

	assert.True(t, time2.After(time1))

}
