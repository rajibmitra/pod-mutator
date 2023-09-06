package main

import (
	"context"
	"encoding/json"
	"net/http"

	v1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func mutatePods(ctx context.Context, req *v1.AdmissionReview) *v1.AdmissionResponse {
	// Assuming the incoming request is of kind Pod
	pod := corev1.Pod{}
	if err := json.Unmarshal(req.Request.Object.Raw, &pod); err != nil {
		return &v1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	}

	// Mutate the Pod (add label "env=dev")
	if pod.Labels == nil {
		pod.Labels = make(map[string]string)
	}
	pod.Labels["env"] = "dev"

	patch, err := json.Marshal([]map[string]interface{}{
		{
			"op":    "add",
			"path":  "/metadata/labels/env",
			"value": "dev",
		},
	})
	if err != nil {
		return &v1.AdmissionResponse{
			Result: &metav1.Status{
				Message: err.Error(),
			},
		}
	}

	return &v1.AdmissionResponse{
		Allowed: true,
		Patch:   patch,
		PatchType: func() *v1.PatchType {
			pt := v1.PatchTypeJSONPatch
			return &pt
		}(),
	}
}

func handleMutatePod(w http.ResponseWriter, r *http.Request) {
	var admissionReview v1.AdmissionReview

	if err := json.NewDecoder(r.Body).Decode(&admissionReview); err != nil {
		http.Error(w, "could not decode request body", http.StatusBadRequest)
		return
	}

	admissionResponse := mutatePods(context.Background(), &admissionReview)
	admissionReview.Response = admissionResponse

	if err := json.NewEncoder(w).Encode(admissionReview); err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/mutate-pod", handleMutatePod)
	http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
}
