package async

import (
	"context"
	"errors"
	"time"

	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/client"
	"github.com/HewlettPackard/hpegl-pcbe-terraform-resources/internal/sdk/async/dataservices"
)

type Operation struct {
	ctx                    context.Context
	client                 client.PCBeClient
	operationID            string
	associatedResourceType string
	opResp                 dataservices.V1beta1AsyncOperationsItemAsyncOperationsGetResponseable
}

// New returns a new Operation struct that can be used to poll
// asynchronous operations for any resource type.
func New(
	ctx context.Context,
	client client.PCBeClient,
	OperationID string,
	associatedResourceType string,
) Operation {
	a := Operation{}
	a.ctx = ctx
	a.client = client
	a.operationID = OperationID
	a.associatedResourceType = associatedResourceType

	return a
}

func (a Operation) GetSourceResourceURI() (string, error) {
	if a.opResp == nil {
		msg := "GetSourceResourceURI requies a non-nil opResp"

		return "", errors.New(msg)
	}

	if a.opResp.GetSourceResourceUri() == nil {
		msg := "GetSourceResourceURI requies a non-nil GetSourceResourceUri"

		return "", errors.New(msg)
	}

	return *a.opResp.GetSourceResourceUri(), nil
}

func (a Operation) GetAssociatedResourceURI() (string, error) {
	if a.opResp == nil {
		msg := "GetAssociatedResourceURI requies a non-nil opResp"

		return "", errors.New(msg)
	}

	if len(a.opResp.GetAssociatedResources()) != 1 {
		msg := "GetAssociatedResourceURI requires exactly one associatedResource"

		return "", errors.New(msg)

	}

	if a.opResp.GetAssociatedResources()[0].GetTypeEscaped() == nil {
		msg := "GetAssociatedResourceURI requires a non-nil associatedResource type"

		return "", errors.New(msg)
	}

	typeEscaped := *(a.opResp.GetAssociatedResources()[0].GetTypeEscaped())
	if typeEscaped != a.associatedResourceType {
		msg := "GetAssociatedResourceURI type mismatch (" + typeEscaped +
			" != " + a.associatedResourceType + ")"

		return "", errors.New(msg)
	}

	if a.opResp.GetAssociatedResources()[0].GetResourceUri() == nil {
		msg := "GetAssociatedResourceURI requires a non-nil associatedResource uri"

		return "", errors.New(msg)
	}

	return *(a.opResp.GetAssociatedResources()[0].GetResourceUri()), nil
}

// Poll waits for an asynchronous operation to complete
// It is not thread safe (but does not currently need to be)
// Typically the operation flow will be:
//
//	a := New(...)
//	a.Poll()
//	uri ... = a.GetAssociatedResourceURI()
func (a *Operation) Poll() error {
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
