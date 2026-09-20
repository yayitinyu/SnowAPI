package relay_test

import (
	"testing"

	"github.com/QuantumNous/new-api/dto"
	"github.com/QuantumNous/new-api/relay/channel/baidu"
	"github.com/QuantumNous/new-api/relay/channel/cloudflare"
	"github.com/QuantumNous/new-api/relay/channel/cohere"
	"github.com/QuantumNous/new-api/relay/channel/dify"
	"github.com/QuantumNous/new-api/relay/channel/jina"
	"github.com/QuantumNous/new-api/relay/channel/mistral"
	"github.com/QuantumNous/new-api/relay/channel/mokaai"
	"github.com/QuantumNous/new-api/relay/channel/palm"
	"github.com/QuantumNous/new-api/relay/channel/tencent"
	"github.com/QuantumNous/new-api/relay/channel/xunfei"
	"github.com/QuantumNous/new-api/relay/channel/zhipu"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type claudeRequestConverter interface {
	ConvertClaudeRequest(*gin.Context, *relaycommon.RelayInfo, *dto.ClaudeRequest) (any, error)
}

func TestUnsupportedClaudeAdaptorsReturnErrors(t *testing.T) {
	tests := []struct {
		name    string
		adaptor claudeRequestConverter
	}{
		{name: "baidu", adaptor: &baidu.Adaptor{}},
		{name: "cloudflare", adaptor: &cloudflare.Adaptor{}},
		{name: "cohere", adaptor: &cohere.Adaptor{}},
		{name: "dify", adaptor: &dify.Adaptor{}},
		{name: "jina", adaptor: &jina.Adaptor{}},
		{name: "mistral", adaptor: &mistral.Adaptor{}},
		{name: "mokaai", adaptor: &mokaai.Adaptor{}},
		{name: "palm", adaptor: &palm.Adaptor{}},
		{name: "tencent", adaptor: &tencent.Adaptor{}},
		{name: "xunfei", adaptor: &xunfei.Adaptor{}},
		{name: "zhipu", adaptor: &zhipu.Adaptor{}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var converted any
			var err error
			require.NotPanics(t, func() {
				converted, err = test.adaptor.ConvertClaudeRequest(
					&gin.Context{},
					&relaycommon.RelayInfo{},
					&dto.ClaudeRequest{},
				)
			})
			assert.Nil(t, converted)
			require.Error(t, err)
		})
	}
}
