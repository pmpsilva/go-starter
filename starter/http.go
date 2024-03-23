package starter

import (
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	hTTPHost            = "HTTP_HOST"
	hTTPPort            = "HTTP_PORT"
	httpDefaultHost     = "0.0.0.0"
	httpDefaultPort     = "8080"
	contextPath         = "CONTEXT_PATH"
	defaultContextPath  = "/"
	envPprofEnabled     = "NET_HTTP_PPROF_ENABLED"
	defaultPprofEnabled = false
	RequestIDKey        = "requestID"
)

type HttpConfig struct {
	Port, Host, ContextPath string
	ProfEnabled             bool
}

// ReadHttpConfig initialize the HTTPConfig
func ReadHttpConfig() *HttpConfig {
	var h HttpConfig
	h.Port = ReadEnvVarOrDefault(hTTPPort, httpDefaultPort)
	h.Host = ReadEnvVarOrDefault(hTTPHost, httpDefaultHost)
	h.ContextPath = ReadEnvVarOrDefault(contextPath, defaultContextPath)
	h.ProfEnabled = ReadEnvBoolVarOrDefault(envPprofEnabled, defaultPprofEnabled)

	return &h
}

func DeriveContextWithRequestId(ctx context.Context) context.Context {
	uuidWithHyphen := uuid.New()
	return context.WithValue(ctx, RequestIDKey, uuidWithHyphen)
}

func addCtxAndRequestIDIfPresent(ctx context.Context) any {
	if reqID := ctx.Value(RequestIDKey); reqID != nil {
		return reqID
	}
	return nil
}

func ZapFieldWithRequestIdFromCtx(ctx context.Context) zap.Field {
	return zap.Any("request_id", addCtxAndRequestIDIfPresent(ctx))
}
