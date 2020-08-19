package main

type Address struct {
	City  NullString `json:"city,omitempty"`
	State NullString `json:"state,omitempty"`
}
