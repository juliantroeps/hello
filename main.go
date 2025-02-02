package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello 👋"))
	})

	http.HandleFunc("/pod", func(w http.ResponseWriter, r *http.Request) {
		podName := os.Getenv("POD_NAME")
		if podName == "" {
			podName = "unknown"
		}

		podNamespace := os.Getenv("POD_NAMESPACE")
		if podNamespace == "" {
			podNamespace = "unknown"
		}

		podCpuLimit := os.Getenv("CPU_LIMIT")
		if podCpuLimit == "" {
			podCpuLimit = "unknown"
		}

		podMemLimit := os.Getenv("MEMORY_LIMIT")
		if podMemLimit == "" {
			podMemLimit = "unknown"
		}

		responseObject := map[string]string{
			"name":      podName,
			"namespace": podNamespace,
			"cpuLimit":  podCpuLimit,
			"memLimit":  podMemLimit,
		}

		jsonResponse, err := json.Marshal(responseObject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	})

	log.Println("server started on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
