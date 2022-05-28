package config

import (
	"context"
	"github.com/google/uuid"
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

type httpConfig struct {
	Port, Host, ContextPath string
	ProfEnabled             bool
}

//ReadHttpConfig initialize the HTTPConfig
func ReadHttpConfig() *httpConfig {
	var h httpConfig
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

func AddCtxAndRequestIDIfPresent(ctx context.Context) any {
	if reqID := ctx.Value(RequestIDKey); reqID != nil {
		return reqID
	}
	return nil
}
