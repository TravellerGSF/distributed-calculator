package tests

import (
	"testing"

	"github.com/TravellerGSF/distributed-calculator/internal/agent"
	"github.com/TravellerGSF/distributed-calculator/internal/models"
)

func TestEvaluateTask(t *testing.T) {
	task := models.Task{Arg1: 2, Arg2: 3, Operation: "+"}
	result := agent.EvaluateTask(task)
	if result != 5 {
		t.Errorf("Ожидалось 5, получено %v", result)
	}
}
