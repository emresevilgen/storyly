package tracing

type AuditInfo struct {
	CorrelationId string `json:"correlationId"`
	AgentName     string `json:"agentName"`
	ExecutorUser  string `json:"executorUser"`
}
