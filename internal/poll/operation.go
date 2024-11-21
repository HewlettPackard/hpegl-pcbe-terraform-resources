package poll

import (
	"context"
	"errors"
	"time"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async/dataservices"
	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type AsyncOperationType struct {
	ctx                    context.Context
	client                 client.PCBeClient
	operationID            string
	associatedResourceType string
	opResp                 dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable
}

func New(
	ctx context.Context,
	client client.PCBeClient,
	OperationID string,
	associatedResourceType string,
) AsyncOperationType {
	a := AsyncOperationType{}
	a.client = client
	a.operationID = OperationID
	a.associatedResourceType = associatedResourceType

	return a
}

func (a AsyncOperationType) GetAssociatedResourceUri() (string, error) {
	if a.opResp == nil {
		msg := "GetAssociatedResourceUri requies a non-nil opResp"

		return "", errors.New(msg)
	}

	if len(a.opResp.GetAssociatedResources()) != 1 {
		msg := "GetAssociatedResourceUri requires exactly one associatedResource"

		return "", errors.New(msg)

	}

	if a.opResp.GetAssociatedResources()[0].GetTypeEscaped() == nil {
		msg := "GetAssociatedResourceUri requires a non-nil associatedResource type"

		return "", errors.New(msg)
	}

	typeEscaped := *(a.opResp.GetAssociatedResources()[0].GetTypeEscaped())
	if typeEscaped != a.associatedResourceType {
		msg := "GetAssociatedResourceUri type mismatch (" + typeEscaped +
			" != " + a.associatedResourceType + ")"

		return "", errors.New(msg)
	}

	if a.opResp.GetAssociatedResources()[0].GetResourceUri() == nil {
		msg := "GetAssociatedResourceUri requires a non-nil associatedResource uri"

		return "", errors.New(msg)
	}

	return *(a.opResp.GetAssociatedResources()[0].GetResourceUri()), nil
}

// Poll waits for an asynchronous operation to complete
func (a *AsyncOperationType) Poll() error {
	maxPolls := int(a.client.Config.MaxPolls)
	pollWaitTime := time.Duration(a.client.Config.PollInterval) * time.Second

	asyncClient, err := a.client.NewAsyncClient(a.ctx)
	if err != nil {
		msg := "error polling async operation " + a.operationID + ": " +
			err.Error()

		return errors.New(msg)
	}

	var opResp dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable
	arc := dataservices.
		V1beta1AsyncOperationsAsyncOperationsItemRequestBuilderGetRequestConfiguration{}

	for count := 1; ; count++ {
		opResp, err = asyncClient.DataServices().
			V1beta1().
			AsyncOperations().
			ById(a.operationID).
			Get(a.ctx, &arc)

		if err != nil {
			msg := "error polling async operation " + a.operationID + ": " + err.Error()

			return errors.New(msg)
		}

		if opResp == nil {
			msg := "error polling async operation " + a.operationID + ": " +
				"nil op response"

			return errors.New(msg)
		}

		opRespState := opResp.GetState()
		if opRespState == nil {
			msg := "error polling async operation " + a.operationID + ": " +
				"operation has nil state"

			return errors.New(msg)
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "FAILED" {
			msg := "error polling async operation " + a.operationID + "operation state FAILED"

			return errors.New(msg)
		}

		// TODO: (API) Use enum not string for state when FF-28181 is fixed
		if *opRespState == "SUCCEEDED" {
			a.opResp = opResp

			break
		}

		if count == maxPolls {
			msg := "error polling async operation " + a.operationID + ": " +
				"max polls exceeded"

			return errors.New(msg)
		}

		time.Sleep(pollWaitTime)
	}

	return nil
}

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

func GetAssociatedResourceUri(
	opResp dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable,
	resourceType string, // enum
) (string, error) {

	if opResp == nil {
		msg := "GetAssociatedResourceUri requies a non-nil opResp"

		return "", errors.New(msg)
	}

	if len(opResp.GetAssociatedResources()) != 1 {
		msg := "GetAssociatedResourceUri requires exactly one associatedResource"

		return "", errors.New(msg)

	}

	if opResp.GetAssociatedResources()[0].GetTypeEscaped() == nil {
		msg := "GetAssociatedResourceUri requires a non-nil associatedResource type"

		return "", errors.New(msg)
	}

	typeEscaped := *(opResp.GetAssociatedResources()[0].GetTypeEscaped())
	if typeEscaped != resourceType {
		msg := "GetAssociatedResourceUri type mismatch (" + typeEscaped +
			" != " + resourceType + ")"

		return "", errors.New(msg)
	}

	if opResp.GetAssociatedResources()[0].GetResourceUri() == nil {
		msg := "GetAssociatedResourceUri requires a non-nil associatedResource uri"

		return "", errors.New(msg)
	}

	return *(opResp.GetAssociatedResources()[0].GetResourceUri()), nil
}
