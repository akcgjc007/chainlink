package fluxmonitorv2

import (
	"math/big"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/flux_aggregator_wrapper"
	"github.com/smartcontractkit/chainlink/core/utils"
)

func (fm *FluxMonitor) ExportedPollIfEligible(threshold, absoluteThreshold float64) {
	fm.pollIfEligible(NewDeviationChecker(threshold, absoluteThreshold))
}

func (fm *FluxMonitor) ExportedSufficientFunds(state flux_aggregator_wrapper.OracleRoundState) bool {
	return fm.sufficientFunds(state)
}

func (fm *FluxMonitor) ExportedSufficientPayment(payment *big.Int) bool {
	return fm.sufficientPayment(payment)
}

func (fm *FluxMonitor) ExportedIsFlagLowered() (bool, error) {
	return fm.isFlagLowered()
}

func (fm *FluxMonitor) ExportedProcessLogs() {
	fm.processLogs()
}

func (fm *FluxMonitor) ExportedBacklog() *utils.BoundedPriorityQueue {
	return fm.backlog
}

func (fm *FluxMonitor) ExportedRoundState() {
	fm.roundState(0)
}

func (fm *FluxMonitor) ExportedRespondToNewRoundLog(log *flux_aggregator_wrapper.FluxAggregatorNewRound) {
	fm.respondToNewRoundLog(*log)
}
