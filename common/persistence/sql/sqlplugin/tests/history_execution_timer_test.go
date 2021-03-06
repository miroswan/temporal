package tests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"go.temporal.io/server/common/convert"
	"go.temporal.io/server/common/persistence/sql/sqlplugin"
	"go.temporal.io/server/common/primitives"
	"go.temporal.io/server/common/shuffle"
)

type (
	historyExecutionTimerSuite struct {
		suite.Suite
		*require.Assertions

		store sqlplugin.HistoryExecutionTimer
	}
)

const (
	testHistoryExecutionTimerID       = "random history timer ID"
	testHistoryExecutionTimerEncoding = "random encoding"
)

var (
	testHistoryExecutionTimerData = []byte("random history execution timer data")
)

func newHistoryExecutionTimerSuite(
	t *testing.T,
	store sqlplugin.HistoryExecutionTimer,
) *historyExecutionTimerSuite {
	return &historyExecutionTimerSuite{
		Assertions: require.New(t),
		store:      store,
	}
}

func (s *historyExecutionTimerSuite) SetupSuite() {

}

func (s *historyExecutionTimerSuite) TearDownSuite() {

}

func (s *historyExecutionTimerSuite) SetupTest() {
	s.Assertions = require.New(s.T())
}

func (s *historyExecutionTimerSuite) TearDownTest() {

}

func (s *historyExecutionTimerSuite) TestReplace_Single() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()
	timerID := shuffle.String(testHistoryExecutionTimerID)

	timer := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, timerID)
	result, err := s.store.ReplaceIntoTimerInfoMaps([]sqlplugin.TimerInfoMapsRow{timer})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))
}

func (s *historyExecutionTimerSuite) TestReplace_Multiple() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()

	timer1 := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, shuffle.String(testHistoryExecutionTimerID))
	timer2 := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, shuffle.String(testHistoryExecutionTimerID))
	result, err := s.store.ReplaceIntoTimerInfoMaps([]sqlplugin.TimerInfoMapsRow{timer1, timer2})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(2, int(rowsAffected))
}

func (s *historyExecutionTimerSuite) TestReplaceSelect_Single() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()

	timer := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, shuffle.String(testHistoryExecutionTimerID))
	result, err := s.store.ReplaceIntoTimerInfoMaps([]sqlplugin.TimerInfoMapsRow{timer})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
	}
	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	rowMap := map[string]sqlplugin.TimerInfoMapsRow{}
	for _, timer := range rows {
		rowMap[timer.TimerID] = timer
	}
	s.Equal(map[string]sqlplugin.TimerInfoMapsRow{
		timer.TimerID: timer,
	}, rowMap)
}

func (s *historyExecutionTimerSuite) TestReplaceSelect_Multiple() {
	numTimers := 20

	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()

	var timers []sqlplugin.TimerInfoMapsRow
	for i := 0; i < numTimers; i++ {
		timer := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, shuffle.String(testHistoryExecutionTimerID))
		timers = append(timers, timer)
	}
	result, err := s.store.ReplaceIntoTimerInfoMaps(timers)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(numTimers, int(rowsAffected))

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
	}
	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	timerMap := map[string]sqlplugin.TimerInfoMapsRow{}
	for _, timer := range timers {
		timerMap[timer.TimerID] = timer
	}
	rowMap := map[string]sqlplugin.TimerInfoMapsRow{}
	for _, timer := range rows {
		rowMap[timer.TimerID] = timer
	}
	s.Equal(timerMap, rowMap)
}

func (s *historyExecutionTimerSuite) TestDeleteSelect_Single() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()
	timerID := shuffle.String(testHistoryExecutionTimerID)

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
		TimerID:     convert.StringPtr(timerID),
	}
	result, err := s.store.DeleteFromTimerInfoMaps(filter)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(0, int(rowsAffected))

	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	s.Equal([]sqlplugin.TimerInfoMapsRow(nil), rows)
}

func (s *historyExecutionTimerSuite) TestDeleteSelect_Multiple() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
		TimerID:     nil,
	}
	result, err := s.store.DeleteFromTimerInfoMaps(filter)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(0, int(rowsAffected))

	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	s.Equal([]sqlplugin.TimerInfoMapsRow(nil), rows)
}

func (s *historyExecutionTimerSuite) TestReplaceDeleteSelect_Single() {
	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()
	timerID := shuffle.String(testHistoryExecutionTimerID)

	timer := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, timerID)
	result, err := s.store.ReplaceIntoTimerInfoMaps([]sqlplugin.TimerInfoMapsRow{timer})
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
		TimerID:     convert.StringPtr(timerID),
	}
	result, err = s.store.DeleteFromTimerInfoMaps(filter)
	s.NoError(err)
	rowsAffected, err = result.RowsAffected()
	s.NoError(err)
	s.Equal(1, int(rowsAffected))

	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	s.Equal([]sqlplugin.TimerInfoMapsRow(nil), rows)
}

func (s *historyExecutionTimerSuite) TestReplaceDeleteSelect_Multiple() {
	numTimers := 20

	shardID := rand.Int31()
	namespaceID := primitives.NewUUID()
	workflowID := shuffle.String(testHistoryExecutionWorkflowID)
	runID := primitives.NewUUID()

	var timers []sqlplugin.TimerInfoMapsRow
	for i := 0; i < numTimers; i++ {
		timer := s.newRandomExecutionTimerRow(shardID, namespaceID, workflowID, runID, shuffle.String(testHistoryExecutionTimerID))
		timers = append(timers, timer)
	}
	result, err := s.store.ReplaceIntoTimerInfoMaps(timers)
	s.NoError(err)
	rowsAffected, err := result.RowsAffected()
	s.NoError(err)
	s.Equal(numTimers, int(rowsAffected))

	filter := &sqlplugin.TimerInfoMapsFilter{
		ShardID:     shardID,
		NamespaceID: namespaceID,
		WorkflowID:  workflowID,
		RunID:       runID,
	}
	result, err = s.store.DeleteFromTimerInfoMaps(filter)
	s.NoError(err)
	rowsAffected, err = result.RowsAffected()
	s.NoError(err)
	s.Equal(numTimers, int(rowsAffected))

	rows, err := s.store.SelectFromTimerInfoMaps(filter)
	s.NoError(err)
	s.Equal([]sqlplugin.TimerInfoMapsRow(nil), rows)
}

func (s *historyExecutionTimerSuite) newRandomExecutionTimerRow(
	shardID int32,
	namespaceID primitives.UUID,
	workflowID string,
	runID primitives.UUID,
	timerID string,
) sqlplugin.TimerInfoMapsRow {
	return sqlplugin.TimerInfoMapsRow{
		ShardID:      shardID,
		NamespaceID:  namespaceID,
		WorkflowID:   workflowID,
		RunID:        runID,
		TimerID:      timerID,
		Data:         shuffle.Bytes(testHistoryExecutionTimerData),
		DataEncoding: testHistoryExecutionTimerEncoding,
	}
}
