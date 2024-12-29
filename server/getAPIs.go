package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Department data API
func departments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("GET /departments")

	// Resolve the absolute path to the JSON file
	jsonFilePath, err := filepath.Abs("../RequiredData/Department.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	// Read the JSON file
	data, err := os.ReadFile(jsonFilePath)
	if err != nil {
		http.Error(w, "Error reading Department.json", http.StatusInternalServerError)
		return
	}

	// Parse the JSON file into a slice of Department structs
	var departments []Department
	err = json.Unmarshal(data, &departments)
	if err != nil {
		http.Error(w, "Failed to parse departments data", http.StatusInternalServerError)
		return
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(departments)
}

// Patient list API
func patients(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	log.Println("GET /patients")

	// Resolve the absolute path to the JSON file
	jsonFilePath, err := filepath.Abs("../RequiredData/Patients.json")
	if err != nil {
		http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
		return
	}

	// Get query parameters: DepartmentID and Status
	departmentID := r.URL.Query().Get("DepartmentID")
	status := r.URL.Query().Get("Status")

	// Validate the query parameters
	if departmentID == "" || status == "" {
		http.Error(w, "DepartmentID and Status are required", http.StatusBadRequest)
		return
	}

	// Open and read the JSON file
	file, err := os.Open(jsonFilePath)
	if err != nil {
		http.Error(w, "Error opening Patients.json", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Define a structure that can hold the array of maps
	var jsonData []map[string]map[string][]Patient 

	// Decode the entire JSON array
	err = json.NewDecoder(file).Decode(&jsonData)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
		return
	}

	// Since the JSON starts with an array, we extract the first element
	if len(jsonData) == 0 {
		http.Error(w, "No data found in the JSON", http.StatusInternalServerError)
		return
	}
	data := jsonData[0]

	// Check if the requested status exists (e.g., "Waiting" or "Attended")
	statusData, foundStatus := data[status]
	if !foundStatus {
		http.Error(w, "Invalid Status provided", http.StatusBadRequest)
		return
	}

	// Check if there are any patients under the requested status
	if len(statusData) == 0 {
		http.Error(w, "No patients found for the given Status", http.StatusNotFound)
		return
	}

	// Check if the requested department exists within the status
	if deptPatients, foundDept := statusData[departmentID]; foundDept {
		// Send the filtered patients as a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(deptPatients)
	} else {
		http.Error(w, "No patients found for the given DepartmentID", http.StatusNotFound)
	}
}

// Patient-centric data API
func patientCentric(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    log.Println("GET /patientCentric")

    // Resolve the absolute path to the JSON file
    jsonFilePath, err := filepath.Abs("../RequiredData/PatientCentricData.json")
    if err != nil {
        http.Error(w, "Failed to resolve file path", http.StatusInternalServerError)
        return
    }

    // Get crno query parameter
    crno := r.URL.Query().Get("CrNo")
    if crno == "" {
        http.Error(w, "CrNo is required", http.StatusBadRequest)
        return
    }

    // Open and read the JSON file
    file, err := os.Open(jsonFilePath)
    if err != nil {
        http.Error(w, "Error opening PatientCentricData.json", http.StatusInternalServerError)
        return
    }
    defer file.Close()

    // Decode the JSON into an array of maps
    var jsonData []map[string]PatientCentric
    err = json.NewDecoder(file).Decode(&jsonData)
    if err != nil {
        log.Printf("Error decoding JSON: %v", err)
        http.Error(w, "Error decoding JSON", http.StatusInternalServerError)
        return
    }

    // Search for the matching CrNo
    for _, record := range jsonData {
        if patientData, found := record[crno]; found {
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode(patientData)
            return
        }
    }

    http.Error(w, "No data found for the given CrNo", http.StatusNotFound)
}
