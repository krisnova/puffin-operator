package stub

import (
	"context"

	"github.com/kris-nova/puffin-operator/pkg/apis/kubernetes/v1alpha1"

	"fmt"

	"github.com/kris-nova/puffin-operator/puffin"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
)

func NewHandler() sdk.Handler {
	return &Handler{
		puffinsPrinted: make(map[string]bool),
	}
}

type Handler struct {
	// Fill me

	// Cache
	puffinsPrinted map[string]bool
}

func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {

	// Nova write down what all we should do!
	//

	//select {
	//case <-ctx.Done():
	//default:
	//
	//}

	switch o := event.Object.(type) {
	case *v1alpha1.Puffin:
		// Enter another package context
		fmt.Printf("We have a puffin: %s\n", o.Name)
		if ok := h.puffinsPrinted[o.Name]; !ok {
			puffin.PrintPuffin()
			h.puffinsPrinted[o.Name] = true
			o.Status.Status = fmt.Sprintf("Printed %s\n", o.Name)
			sdk.Update(o)
		} else {
			o.Status.Status = fmt.Sprintf("Skipped %s\n", o.Name)
			sdk.Update(o)
		}
	case *v1alpha1.PuffinList:
		fmt.Printf("We have a list of puffins: %s\n", o.Kind)
	}
	return nil
}
