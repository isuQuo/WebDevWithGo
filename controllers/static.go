package controllers

import (
	"net/http"
)

// would we use generics here instead of an interface as a parameter?
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   string
	}{
		{
			Question: "What did you eat for breakfast?",
			Answer:   "Never you mind.",
		},
		{
			Question: "What time did you come over for dinner?",
			Answer:   "One day haha.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
