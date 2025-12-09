// SPDX-License-Identifier: BSD-3-Clause

package metrics

import (
	"os"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus/collectors"
)

const (
	envNamespace = "METRICS_NAMESPACE"

	envWithoutProcessMetrics   = "METRICS_WITHOUT_PROCESS"
	envWithoutCPUMetrics       = "METRICS_WITHOUT_CPU"
	envWithoutGCMetrics        = "METRICS_WITHOUT_GC"
	envWithoutMemoryMetrics    = "METRICS_WITHOUT_MEMORY"
	envWithoutSchedulerMetrics = "METRICS_WITHOUT_SCHEDULER"
	envWithoutSyncMetrics      = "METRICS_WITHOUT_SYNC"
)

func init() {
	if v, ok := os.LookupEnv(envNamespace); ok {
		namespace = v
	}

	if withProcessMetrics() {
		MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	}

	var rules []collectors.GoRuntimeMetricsRule
	if withCPUMetrics() {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/cpu/classes/.*`)})
	}
	if withGCMetrics() {
		rules = append(rules, collectors.MetricsGC)
	}
	if withMemoryMetrics() {
		rules = append(rules, collectors.MetricsMemory)
	}
	if withSchedulerMetrics() {
		rules = append(rules, collectors.MetricsScheduler)
	}
	if withSyncMetrics() {
		rules = append(rules, collectors.GoRuntimeMetricsRule{Matcher: regexp.MustCompile(`^/sync/.*`)})
	}
	if len(rules) > 0 {
		MustRegister(collectors.NewGoCollector(collectors.WithGoCollectorRuntimeMetrics(rules...)))
	}
}

func withProcessMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutProcessMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withCPUMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutCPUMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withGCMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutGCMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withMemoryMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutMemoryMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withSchedulerMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutSchedulerMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}

func withSyncMetrics() bool {
	if v, ok := os.LookupEnv(envWithoutSyncMetrics); ok {
		without, _ := strconv.ParseBool(v)
		return !without
	}
	return true
}
