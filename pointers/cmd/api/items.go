package api

import (
	"encoding/json"
	"go-playground/pointers/pkg/data"
	"net/http"
	"strconv"
)

func HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid price"))
		return
	}

	item, err := data.Store.Create(data.Item{
		Name:  r.FormValue("name"),
		Price: price,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(item)
	w.Write(response)
}

func HandleList(w http.ResponseWriter, r *http.Request) {
	items := data.Store.List()
	response, _ := json.Marshal(items)

	w.Write(response)
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item := data.Store.Read(name)
	if item == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response, _ := json.Marshal(item)
	w.Write(response)
}

func HandleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid price"))
		return
	}

	item, err := data.Store.Update(data.Item{
		Name:  r.FormValue("name"),
		Price: price,
	})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	response, _ := json.Marshal(item)
	w.Write(response)
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := data.Store.Delete(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
