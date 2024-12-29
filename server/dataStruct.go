package main

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Patient struct {
	CrNo   string `json:"CrNo"`
	Name   string `json:"name"`
	Age    int    `json:"age,string"` 
	Gender string `json:"gender"`
}

type Investigation struct {
	Name   string `json:"Name"`
	Result string `json:"Result"`
}

type Allergy struct {
	Name     string `json:"Name"`
	Reaction string `json:"Reaction"`
}

type Medication struct {
	Name 	  string `json:"Name"`
	Dose      string `json:"Dose"`
	Frequency string `json:"Frequency"`
}

type Vital struct {
    Height       string `json:"Height"`
    Weight       string `json:"Weight"`
    BP           string `json:"BP"`
    Pulse        string `json:"Pulse"`
    Temperature  string `json:"Temperature"`
}

type PatientCentric struct {
    Age            string `json:"Age"` // Keep as string to match the JSON
    Department     string `json:"Department"`
    Attended       bool   `json:"Attended"`
    Name           string `json:"Name"`
    Gender         string `json:"Gender"`
    ChiefComplaint string `json:"Chief Complaint"`
    History        string `json:"History"`
    Diagnosis      string `json:"Diagnosis"`
    Treatment      string `json:"Treatment"`
    FollowUp       string `json:"Follow-up"`
    Investigations []Investigation `json:"Investigations"`
    Allergies      []Allergy       `json:"Allergies"`
    Medications    []Medication    `json:"Medications"`
    Vitals         []Vital         `json:"Vitals"`
}