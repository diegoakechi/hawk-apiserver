package main

import (
	"os"
	"testing"
	"github.com/ClusterLabs/hawk-apiserver/cib"
	"github.com/stretchr/testify/assert"
)

func TestAsyncCib(t *testing.T) {
	testFile := "./test/old-cib.xml"
	testVersion := "0:86:125"
	os.Setenv("CIB_file", testFile)
	c := cib.AsyncCib{}
	c.Start()
	s := c.Wait(1, "")
	assert.Equal(t, s, testVersion, "Expected version triple of test/old-cib.xml")
	s2 := c.Get()
	assert.NotEmpty(t, s2)
	v := c.Version()
	assert.Equal(t, v.String(), testVersion, "Expected known CIB version")
}
