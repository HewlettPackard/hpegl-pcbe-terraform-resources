package poll

import (
	"context"
	"time"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async/dataservices"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

// AsyncOperation polls an asynchronous operation until it completes.
// It returns the URI of the source resource of the operation.
func AsyncOperation(
	ctx context.Context,
	client client.PCBeClient,
	OperationID string,
	diagsP *diag.Diagnostics,
) *string {
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
			msg := "nil op response"
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				msg,
			)

			return nil
		}

		opRespState := opResp.GetState()
		if opRespState == nil {
			msg := "operation has nil state"
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				msg,
			)

			return nil
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "FAILED" {
			msg := "operation state FAILED"

			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				msg,
			)

			return nil
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "SUCCEEDED" {
			break
		}

		if count == maxPolls {
			msg := "max polls exceeded"
			(*diagsP).AddError(
				"error polling async operation "+OperationID,
				msg,
			)

			return nil
		}

		time.Sleep(pollWaitTime)
	}

	return opResp.GetSourceResourceUri()
}
