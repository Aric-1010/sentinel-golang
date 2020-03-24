package dubbo

import (
	"testing"

	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/protocol"
	"github.com/apache/dubbo-go/protocol/invocation"
	"github.com/stretchr/testify/assert"
)

func TestGetResourceName(t *testing.T) {
	url, err := common.NewURL("dubbo://127.0.0.1:20000/UserProvider?anyhost=true&" +
		"version=1.0.0&group=myGroup&" +
		"application=BDTService&category=providers&default.timeout=10000&dubbo=dubbo-provider-golang-1.0.0&" +
		"environment=dev&interface=com.ikurento.user.UserProvider&ip=192.168.56.1&methods=GetUser%2C&" +
		"module=dubbogo+user-info+server&org=ikurento.com&owner=ZX&pid=1447&revision=0.0.1&" +
		"side=provider&timeout=3000&timestamp=1556509797245&bean.name=UserProvider")
	assert.NoError(t, err)
	mockInvoker := protocol.NewBaseInvoker(url)
	methodResourceName := getResourceName(mockInvoker,
		invocation.NewRPCInvocation("hello", []interface{}{"OK"}, make(map[string]string)), "prefix_")
	assert.Equal(t, "prefix_com.ikurento.user.UserProvider:myGroup:1.0.0:hello()", methodResourceName)
}
