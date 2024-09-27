package clients

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

var UUIDv4Regex = regexp.MustCompile("")

func TestUUIDGenerator(t *testing.T) {

	idGen := NewUuidGenerator()

	assert.True(t, UUIDv4Regex.MatchString(idGen.GetID()))

}
