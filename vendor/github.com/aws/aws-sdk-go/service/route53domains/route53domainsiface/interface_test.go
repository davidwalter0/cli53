// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package route53domainsiface_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/route53domains"
	"github.com/aws/aws-sdk-go/service/route53domains/route53domainsiface"
	"github.com/stretchr/testify/assert"
)

func TestInterface(t *testing.T) {
	assert.Implements(t, (*route53domainsiface.Route53DomainsAPI)(nil), route53domains.New(nil))
}