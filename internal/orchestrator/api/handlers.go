package api

import (
	"encoding/json"
	"net/http"

	"github.com/TravellerGSF/distributed-calculator/internal/models"
	"github.com/TravellerGSF/distributed-calculator/internal/orchestrator/services"
)

func SubmitExpression(w http.ResponseWriter, r *http.Request) {
	var expr models.ExpressionRequest
	if err := json.NewDecoder(r.Body).Decode(&expr); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusUnprocessableEntity)
		return
	}

	id, err := services.AddExpression(expr.Expression)
	if err != nil {
		http.Error(w, "Ошибка при добавлении выражения", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetExpressions(w http.ResponseWriter, r *http.Request) {
	expressions := services.GetExpressions()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]models.ExpressionResponse{"expressions": expressions})
}

func GetExpressionByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID не указан", http.StatusBadRequest)
		return
	}

	expr, err := services.GetExpressionByID(id)
	if err != nil {
		if err == services.ErrExpressionNotFound {
			http.Error(w, "Выражение не найдено", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при получении выражения", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]models.ExpressionResponse{"expression": expr})
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	task, err := services.GetTask()
	if err != nil {
		if err == services.ErrNoTasks {
			http.Error(w, "Нет доступных задач", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при получении задачи", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]models.Task{"task": task})
}

func SubmitResult(w http.ResponseWriter, r *http.Request) {
	var result models.TaskResult
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusUnprocessableEntity)
		return
	}

	if err := services.SubmitResult(result.ID, result.Result); err != nil {
		if err == services.ErrTaskNotFound {
			http.Error(w, "Задача не найдена", http.StatusNotFound)
			return
		}
		http.Error(w, "Ошибка при сохранении результата", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
