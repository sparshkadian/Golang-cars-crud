package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Car struct {
	Company           string `json:"company"`
	Model             string `json:"model"`
	Color             string `json:"color"`
	Price             int    `json:"price"`
	ManufacturingYear int    `json:"manufacturingYear"`
	Owner             `json:"owner"`
}

type Owner struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var cars []Car

const port string = ":8080"

func getAllCars(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(cars); err != nil {
		http.Error(w, "Failed to encode Cars Data", http.StatusInternalServerError)
	}
}

func addCar(w http.ResponseWriter, r *http.Request){
	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Error decoding car data", http.StatusInternalServerError)
	}
	cars = append(cars, car)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cars)
}

func updateCar(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	company := params["Company"]

	var updatedCar Car
	if err := json.NewDecoder(r.Body).Decode(&updatedCar); err != nil{
		http.Error(w, "Error decoding car data", http.StatusInternalServerError)
	}

	var found bool
	for i , car := range cars{
		if car.Company == company{
			cars[i] = updatedCar
			found = true
			break

		}
	}
	if !found {
		http.Error(w, "No car found of this company", http.StatusBadRequest)
        return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedCar)
}

func removeCar(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	company := params["Company"]
	
	var indexToRemove int = -1
	for i , car := range cars{
		if company == car.Company{
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
        http.Error(w, "No car found of this company", http.StatusBadRequest)
        return
    }

	cars = append(cars[:indexToRemove], cars[indexToRemove+1:]...)
	w.WriteHeader(http.StatusNoContent)
}

func main() {

	r := mux.NewRouter()

	cars = append(cars, Car{Company: "Honda", Model: "Civic", Color: "Blue", Price: 20000, ManufacturingYear: 2021, Owner: Owner{Name: "John", Age: 30}})
	cars = append(cars, Car{Company: "Toyota", Model: "Corolla", Color: "Red", Price: 22000, ManufacturingYear: 2020, Owner: Owner{Name: "Alice", Age: 35}})
	cars = append(cars, Car{Company: "BMW", Model: "X5", Color: "Black", Price: 50000, ManufacturingYear: 2019, Owner: Owner{Name: "Bob", Age: 40}})
	cars = append(cars, Car{Company: "Audi", Model: "A4", Color: "White", Price: 45000, ManufacturingYear: 2020, Owner: Owner{Name: "Carol", Age: 45}})
	cars = append(cars, Car{Company: "Ford", Model: "Mustang", Color: "Yellow", Price: 40000, ManufacturingYear: 2018, Owner: Owner{Name: "David", Age: 50}})

	fmt.Printf("Server Started at port %s\n", port)

	r.HandleFunc("/cars", getAllCars).Methods("GET")
	r.HandleFunc("/cars", addCar).Methods("POST")
	r.HandleFunc("/cars/{Company}", updateCar).Methods("PUT")
	r.HandleFunc("/cars/{Company}", removeCar).Methods("DELETE")

	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Fatal("Error Starting Server", err)
	}
}