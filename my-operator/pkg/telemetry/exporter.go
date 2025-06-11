package telemetry

import (
	"context"
	"fmt"
)

// TelemetryExporter defines the interface for exporting telemetry data, such as
// events, metrics, and traces. This abstraction allows the operator to integrate
// with different telemetry backends (e.g., OpenTelemetry collector, Prometheus).
type TelemetryExporter interface {
	// RecordEvent records a specific event that occurred within the operator.
	// Events are discrete occurrences, like "CRDProcessed" or "ErrorEncountered".
	// Attributes provide additional context to the event.
	RecordEvent(eventName string, attributes map[string]string)

	// TrackTrace starts a new trace span for monitoring the execution of a specific operation.
	// It takes a parent context (which might contain an existing trace) and returns a new
	// context with the new span, along with a function to end the span when the operation is complete.
	// This is a simplified representation; actual OpenTelemetry usage is more nuanced.
	TrackTrace(spanName string, parentCtx context.Context) (context.Context, func())
}

// DummyTelemetryExporter is a placeholder implementation of TelemetryExporter.
// It's used for testing or when a real telemetry backend is not configured.
// Its methods log that they were called but perform no actual OpenTelemetry operations.
type DummyTelemetryExporter struct{}

// NewDummyTelemetryExporter creates a new DummyTelemetryExporter.
func NewDummyTelemetryExporter() *DummyTelemetryExporter {
	return &DummyTelemetryExporter{}
}

// RecordEvent logs that an event recording was attempted.
// A real implementation would send this event to a telemetry backend.
func (e *DummyTelemetryExporter) RecordEvent(eventName string, attributes map[string]string) {
	fmt.Printf("TelemetryExporter (Dummy): RecordEvent called for event: '%s', with attributes: %v\n", eventName, attributes)
}

// TrackTrace logs that a trace tracking was attempted and returns a no-op closer function.
// A real implementation would create and manage an actual trace span using OpenTelemetry libraries.
func (e *DummyTelemetryExporter) TrackTrace(spanName string, parentCtx context.Context) (context.Context, func()) {
	fmt.Printf("TelemetryExporter (Dummy): TrackTrace called for span: '%s'\n", spanName)
	// Return the parent context and a no-op function for ending the span.
	return parentCtx, func() {
		fmt.Printf("TelemetryExporter (Dummy): Span '%s' ended.\n", spanName)
	}
}
