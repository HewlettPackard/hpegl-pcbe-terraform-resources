package poll

import (
	"context"
	"time"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async/dataservices"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// AsyncOperation polls an asynchronous operation until it completes.
// It returns the response of the async operation.
func AsyncOperation(
	ctx context.Context,
	client client.PCBeClient,
	OperationID string,
	diagsP *diag.Diagnostics,
) dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable {
	maxPolls := int(client.Config.MaxPolls)
	pollWaitTime := time.Duration(client.Config.PollInterval) * time.Second

	asyncClient, err := client.NewAsyncClient(ctx)
	if err != nil {
		(*diagsP).AddError(
			"error polling async operation "+OperationID,
			err.Error(),
		)

		return nil
	}

	var opResp dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable
	arc := dataservices.
		V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration{}

	for count := 1; ; count++ {
		opResp, err = asyncClient.DataServices().
			V1beta1().
			AsyncOperations().
			ById(OperationID).
			Get(ctx, &arc)
		if err != nil {
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				err.Error(),
			)

			return nil
		}

		if opResp == nil {
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				"nil op response",
			)

			return nil
		}

		opRespState := opResp.GetState()
		if opRespState == nil {
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				"operation has nil state",
			)

			return nil
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "FAILED" {
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				"operation state FAILED",
			)

			return nil
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "SUCCEEDED" {
			break
		}

		if count == maxPolls {
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				"max polls exceeded",
			)

			return nil
		}

		time.Sleep(pollWaitTime)
	}

	return opResp
}
