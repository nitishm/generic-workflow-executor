package actions

import (
	"context"
	"log"
	"time"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Delete handles the deletion or cleanup action
// This could mean deleting a kubernetes resource, or removing a file from a filesystem, etc.

// Add your logic here
// Extend the function parameters to add more options and inputs for your action.
func Delete(ctx context.Context, cancel context.CancelFunc, clientSet client.Client, interval time.Duration, data string) error {
	defer cancel()

	// Add your logic here
	// Step 1: Implement the install logic
	log.Printf("Delete called with data: %s", data)

	// Add your logic here
	// Step 2: (optional) Poll for success condition

	return nil
}
