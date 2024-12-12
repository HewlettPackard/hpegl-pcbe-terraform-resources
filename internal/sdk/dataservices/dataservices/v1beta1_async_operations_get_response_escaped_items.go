package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1AsyncOperationsGetResponse_items the operation resource provides information about progress of asynchronousrequest processing and is where associated resources can be found.
type V1beta1AsyncOperationsGetResponse_items struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A link to be displayed in the Operations UI. This can be used when a operation is paused to take the user to the console UI page with information on how to unpause the operation, or for more general information when the operation is in other states.
    additionalDetails V1beta1AsyncOperationsGetResponse_items_additionalDetailsable
    // Resources that are associated with the operation. These may be created by the operation or other resources that are involved in the operation.
    associatedResources []V1beta1AsyncOperationsGetResponse_items_associatedResourcesable
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // The displayed name for the operation.
    displayName *string
    // The time this operation completed.
    endedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The error response status of the operation.
    error V1beta1AsyncOperationsGetResponse_items_errorable
    // An estimate of how long the operation will run before completing.
    estimatedRunningDurationMinutes *int32
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // A list of groups associated with this operation.
    groups []V1beta1AsyncOperationsGetResponse_items_groupsable
    // Indicates this operation possesses child operations. The list of child operations for a given operation may be acquired by querying /data-services/v1beta1/async-operations?filter=parent/id+eq+'{id}' with pagination parameters added as needed to retrieve the full list.
    hasChildOperations *bool
    // The health status indicates if any errors or problems have been encountered during the operation. Expected values are OK, ERROR, WARNING, UNKNOWN, and UNSPECIFIED.
    healthStatus *string
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Time stamped messages that record the progress of the operation. The number of messages presented is limited in number and only the most recent messages are included. Some older log messages may therefore be removed to make room for newer messages.
    logMessages []V1beta1AsyncOperationsGetResponse_items_logMessagesable
    // A system specified name for the resource.
    name *string
    // The parent is the operation that initiated this sub-operation. If this operation is not a sub-operation this will be null.
    parent V1beta1AsyncOperationsGetResponse_items_parentable
    // A percentage representation of progress to completion.
    progressPercent *int32
    // Recommendations on how to fix failing operations.
    recommendations []V1beta1AsyncOperationsGetResponse_items_recommendationsable
    // The self reference for this resource.
    resourceUri *string
    // The root of the tree of operations. If this operation is not part of a tree this will be null.
    rootOperation V1beta1AsyncOperationsGetResponse_items_rootOperationable
    // List of services this operation belongs to, can be used to filter to specific services in the UI.
    services []string
    // The resource that was used to initiate the operation.
    sourceResourceUri *string
    // The time this operation was started.
    startedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A message to indicate the current state of the operation, for example the current step in a workflow.Expected values are INITIALIZED, RUNNING, FAILED, SUCCEEDED, TIMEDOUT, PAUSED, CANCELLED, and UNSPECIFIED.
    state *string
    // The count of the number of child Operations below this one, recursively.
    subtreeOperationCount *int32
    // This attribute suggests a suitable interval to use when polling for progress. Where specified this will be based on the frequency with which the operation is likely to be updated.
    suggestedPollingIntervalSeconds *int32
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The ID or email address of the user that initiated the operation.
    userId *string
}
// NewV1beta1AsyncOperationsGetResponse_items instantiates a new V1beta1AsyncOperationsGetResponse_items and sets the default values.
func NewV1beta1AsyncOperationsGetResponse_items()(*V1beta1AsyncOperationsGetResponse_items) {
    m := &V1beta1AsyncOperationsGetResponse_items{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1AsyncOperationsGetResponse_itemsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1AsyncOperationsGetResponse_itemsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1AsyncOperationsGetResponse_items(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdditionalDetails gets the additionalDetails property value. A link to be displayed in the Operations UI. This can be used when a operation is paused to take the user to the console UI page with information on how to unpause the operation, or for more general information when the operation is in other states.
// returns a V1beta1AsyncOperationsGetResponse_items_additionalDetailsable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetAdditionalDetails()(V1beta1AsyncOperationsGetResponse_items_additionalDetailsable) {
    return m.additionalDetails
}
// GetAssociatedResources gets the associatedResources property value. Resources that are associated with the operation. These may be created by the operation or other resources that are involved in the operation.
// returns a []V1beta1AsyncOperationsGetResponse_items_associatedResourcesable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetAssociatedResources()([]V1beta1AsyncOperationsGetResponse_items_associatedResourcesable) {
    return m.associatedResources
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetCustomerId()(*string) {
    return m.customerId
}
// GetDisplayName gets the displayName property value. The displayed name for the operation.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndedAt gets the endedAt property value. The time this operation completed.
// returns a *Time when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetEndedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endedAt
}
// GetError gets the error property value. The error response status of the operation.
// returns a V1beta1AsyncOperationsGetResponse_items_errorable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetError()(V1beta1AsyncOperationsGetResponse_items_errorable) {
    return m.error
}
// GetEstimatedRunningDurationMinutes gets the estimatedRunningDurationMinutes property value. An estimate of how long the operation will run before completing.
// returns a *int32 when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetEstimatedRunningDurationMinutes()(*int32) {
    return m.estimatedRunningDurationMinutes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["additionalDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1AsyncOperationsGetResponse_items_additionalDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalDetails(val.(V1beta1AsyncOperationsGetResponse_items_additionalDetailsable))
        }
        return nil
    }
    res["associatedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AsyncOperationsGetResponse_items_associatedResourcesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AsyncOperationsGetResponse_items_associatedResourcesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AsyncOperationsGetResponse_items_associatedResourcesable)
                }
            }
            m.SetAssociatedResources(res)
        }
        return nil
    }
    res["createdAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedAt(val)
        }
        return nil
    }
    res["customerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomerId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["endedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndedAt(val)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1AsyncOperationsGetResponse_items_errorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(V1beta1AsyncOperationsGetResponse_items_errorable))
        }
        return nil
    }
    res["estimatedRunningDurationMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEstimatedRunningDurationMinutes(val)
        }
        return nil
    }
    res["generation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeneration(val)
        }
        return nil
    }
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AsyncOperationsGetResponse_items_groupsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AsyncOperationsGetResponse_items_groupsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AsyncOperationsGetResponse_items_groupsable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    res["hasChildOperations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasChildOperations(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["logMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AsyncOperationsGetResponse_items_logMessagesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AsyncOperationsGetResponse_items_logMessagesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AsyncOperationsGetResponse_items_logMessagesable)
                }
            }
            m.SetLogMessages(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1AsyncOperationsGetResponse_items_parentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParent(val.(V1beta1AsyncOperationsGetResponse_items_parentable))
        }
        return nil
    }
    res["progressPercent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProgressPercent(val)
        }
        return nil
    }
    res["recommendations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1AsyncOperationsGetResponse_items_recommendationsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1AsyncOperationsGetResponse_items_recommendationsable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1AsyncOperationsGetResponse_items_recommendationsable)
                }
            }
            m.SetRecommendations(res)
        }
        return nil
    }
    res["resourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceUri(val)
        }
        return nil
    }
    res["rootOperation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1AsyncOperationsGetResponse_items_rootOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRootOperation(val.(V1beta1AsyncOperationsGetResponse_items_rootOperationable))
        }
        return nil
    }
    res["services"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetServices(res)
        }
        return nil
    }
    res["sourceResourceUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceResourceUri(val)
        }
        return nil
    }
    res["startedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartedAt(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["subtreeOperationCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubtreeOperationCount(val)
        }
        return nil
    }
    res["suggestedPollingIntervalSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestedPollingIntervalSeconds(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTypeEscaped(val)
        }
        return nil
    }
    res["updatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatedAt(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetGeneration()(*int64) {
    return m.generation
}
// GetGroups gets the groups property value. A list of groups associated with this operation.
// returns a []V1beta1AsyncOperationsGetResponse_items_groupsable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetGroups()([]V1beta1AsyncOperationsGetResponse_items_groupsable) {
    return m.groups
}
// GetHasChildOperations gets the hasChildOperations property value. Indicates this operation possesses child operations. The list of child operations for a given operation may be acquired by querying /data-services/v1beta1/async-operations?filter=parent/id+eq+'{id}' with pagination parameters added as needed to retrieve the full list.
// returns a *bool when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetHasChildOperations()(*bool) {
    return m.hasChildOperations
}
// GetHealthStatus gets the healthStatus property value. The health status indicates if any errors or problems have been encountered during the operation. Expected values are OK, ERROR, WARNING, UNKNOWN, and UNSPECIFIED.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetHealthStatus()(*string) {
    return m.healthStatus
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetLogMessages gets the logMessages property value. Time stamped messages that record the progress of the operation. The number of messages presented is limited in number and only the most recent messages are included. Some older log messages may therefore be removed to make room for newer messages.
// returns a []V1beta1AsyncOperationsGetResponse_items_logMessagesable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetLogMessages()([]V1beta1AsyncOperationsGetResponse_items_logMessagesable) {
    return m.logMessages
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetName()(*string) {
    return m.name
}
// GetParent gets the parent property value. The parent is the operation that initiated this sub-operation. If this operation is not a sub-operation this will be null.
// returns a V1beta1AsyncOperationsGetResponse_items_parentable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetParent()(V1beta1AsyncOperationsGetResponse_items_parentable) {
    return m.parent
}
// GetProgressPercent gets the progressPercent property value. A percentage representation of progress to completion.
// returns a *int32 when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetProgressPercent()(*int32) {
    return m.progressPercent
}
// GetRecommendations gets the recommendations property value. Recommendations on how to fix failing operations.
// returns a []V1beta1AsyncOperationsGetResponse_items_recommendationsable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetRecommendations()([]V1beta1AsyncOperationsGetResponse_items_recommendationsable) {
    return m.recommendations
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetRootOperation gets the rootOperation property value. The root of the tree of operations. If this operation is not part of a tree this will be null.
// returns a V1beta1AsyncOperationsGetResponse_items_rootOperationable when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetRootOperation()(V1beta1AsyncOperationsGetResponse_items_rootOperationable) {
    return m.rootOperation
}
// GetServices gets the services property value. List of services this operation belongs to, can be used to filter to specific services in the UI.
// returns a []string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetServices()([]string) {
    return m.services
}
// GetSourceResourceUri gets the sourceResourceUri property value. The resource that was used to initiate the operation.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetSourceResourceUri()(*string) {
    return m.sourceResourceUri
}
// GetStartedAt gets the startedAt property value. The time this operation was started.
// returns a *Time when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startedAt
}
// GetState gets the state property value. A message to indicate the current state of the operation, for example the current step in a workflow.Expected values are INITIALIZED, RUNNING, FAILED, SUCCEEDED, TIMEDOUT, PAUSED, CANCELLED, and UNSPECIFIED.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetState()(*string) {
    return m.state
}
// GetSubtreeOperationCount gets the subtreeOperationCount property value. The count of the number of child Operations below this one, recursively.
// returns a *int32 when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetSubtreeOperationCount()(*int32) {
    return m.subtreeOperationCount
}
// GetSuggestedPollingIntervalSeconds gets the suggestedPollingIntervalSeconds property value. This attribute suggests a suitable interval to use when polling for progress. Where specified this will be based on the frequency with which the operation is likely to be updated.
// returns a *int32 when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetSuggestedPollingIntervalSeconds()(*int32) {
    return m.suggestedPollingIntervalSeconds
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetUserId gets the userId property value. The ID or email address of the user that initiated the operation.
// returns a *string when successful
func (m *V1beta1AsyncOperationsGetResponse_items) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *V1beta1AsyncOperationsGetResponse_items) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("additionalDetails", m.GetAdditionalDetails())
        if err != nil {
            return err
        }
    }
    if m.GetAssociatedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssociatedResources()))
        for i, v := range m.GetAssociatedResources() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("associatedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endedAt", m.GetEndedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("estimatedRunningDurationMinutes", m.GetEstimatedRunningDurationMinutes())
        if err != nil {
            return err
        }
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasChildOperations", m.GetHasChildOperations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("healthStatus", m.GetHealthStatus())
        if err != nil {
            return err
        }
    }
    if m.GetLogMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLogMessages()))
        for i, v := range m.GetLogMessages() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("logMessages", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("progressPercent", m.GetProgressPercent())
        if err != nil {
            return err
        }
    }
    if m.GetRecommendations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecommendations()))
        for i, v := range m.GetRecommendations() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("recommendations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("rootOperation", m.GetRootOperation())
        if err != nil {
            return err
        }
    }
    if m.GetServices() != nil {
        err := writer.WriteCollectionOfStringValues("services", m.GetServices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sourceResourceUri", m.GetSourceResourceUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startedAt", m.GetStartedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("subtreeOperationCount", m.GetSubtreeOperationCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("suggestedPollingIntervalSeconds", m.GetSuggestedPollingIntervalSeconds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1AsyncOperationsGetResponse_items) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdditionalDetails sets the additionalDetails property value. A link to be displayed in the Operations UI. This can be used when a operation is paused to take the user to the console UI page with information on how to unpause the operation, or for more general information when the operation is in other states.
func (m *V1beta1AsyncOperationsGetResponse_items) SetAdditionalDetails(value V1beta1AsyncOperationsGetResponse_items_additionalDetailsable)() {
    m.additionalDetails = value
}
// SetAssociatedResources sets the associatedResources property value. Resources that are associated with the operation. These may be created by the operation or other resources that are involved in the operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetAssociatedResources(value []V1beta1AsyncOperationsGetResponse_items_associatedResourcesable)() {
    m.associatedResources = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1AsyncOperationsGetResponse_items) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1AsyncOperationsGetResponse_items) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDisplayName sets the displayName property value. The displayed name for the operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndedAt sets the endedAt property value. The time this operation completed.
func (m *V1beta1AsyncOperationsGetResponse_items) SetEndedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endedAt = value
}
// SetError sets the error property value. The error response status of the operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetError(value V1beta1AsyncOperationsGetResponse_items_errorable)() {
    m.error = value
}
// SetEstimatedRunningDurationMinutes sets the estimatedRunningDurationMinutes property value. An estimate of how long the operation will run before completing.
func (m *V1beta1AsyncOperationsGetResponse_items) SetEstimatedRunningDurationMinutes(value *int32)() {
    m.estimatedRunningDurationMinutes = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1AsyncOperationsGetResponse_items) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGroups sets the groups property value. A list of groups associated with this operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetGroups(value []V1beta1AsyncOperationsGetResponse_items_groupsable)() {
    m.groups = value
}
// SetHasChildOperations sets the hasChildOperations property value. Indicates this operation possesses child operations. The list of child operations for a given operation may be acquired by querying /data-services/v1beta1/async-operations?filter=parent/id+eq+'{id}' with pagination parameters added as needed to retrieve the full list.
func (m *V1beta1AsyncOperationsGetResponse_items) SetHasChildOperations(value *bool)() {
    m.hasChildOperations = value
}
// SetHealthStatus sets the healthStatus property value. The health status indicates if any errors or problems have been encountered during the operation. Expected values are OK, ERROR, WARNING, UNKNOWN, and UNSPECIFIED.
func (m *V1beta1AsyncOperationsGetResponse_items) SetHealthStatus(value *string)() {
    m.healthStatus = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1AsyncOperationsGetResponse_items) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetLogMessages sets the logMessages property value. Time stamped messages that record the progress of the operation. The number of messages presented is limited in number and only the most recent messages are included. Some older log messages may therefore be removed to make room for newer messages.
func (m *V1beta1AsyncOperationsGetResponse_items) SetLogMessages(value []V1beta1AsyncOperationsGetResponse_items_logMessagesable)() {
    m.logMessages = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1AsyncOperationsGetResponse_items) SetName(value *string)() {
    m.name = value
}
// SetParent sets the parent property value. The parent is the operation that initiated this sub-operation. If this operation is not a sub-operation this will be null.
func (m *V1beta1AsyncOperationsGetResponse_items) SetParent(value V1beta1AsyncOperationsGetResponse_items_parentable)() {
    m.parent = value
}
// SetProgressPercent sets the progressPercent property value. A percentage representation of progress to completion.
func (m *V1beta1AsyncOperationsGetResponse_items) SetProgressPercent(value *int32)() {
    m.progressPercent = value
}
// SetRecommendations sets the recommendations property value. Recommendations on how to fix failing operations.
func (m *V1beta1AsyncOperationsGetResponse_items) SetRecommendations(value []V1beta1AsyncOperationsGetResponse_items_recommendationsable)() {
    m.recommendations = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1AsyncOperationsGetResponse_items) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetRootOperation sets the rootOperation property value. The root of the tree of operations. If this operation is not part of a tree this will be null.
func (m *V1beta1AsyncOperationsGetResponse_items) SetRootOperation(value V1beta1AsyncOperationsGetResponse_items_rootOperationable)() {
    m.rootOperation = value
}
// SetServices sets the services property value. List of services this operation belongs to, can be used to filter to specific services in the UI.
func (m *V1beta1AsyncOperationsGetResponse_items) SetServices(value []string)() {
    m.services = value
}
// SetSourceResourceUri sets the sourceResourceUri property value. The resource that was used to initiate the operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetSourceResourceUri(value *string)() {
    m.sourceResourceUri = value
}
// SetStartedAt sets the startedAt property value. The time this operation was started.
func (m *V1beta1AsyncOperationsGetResponse_items) SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startedAt = value
}
// SetState sets the state property value. A message to indicate the current state of the operation, for example the current step in a workflow.Expected values are INITIALIZED, RUNNING, FAILED, SUCCEEDED, TIMEDOUT, PAUSED, CANCELLED, and UNSPECIFIED.
func (m *V1beta1AsyncOperationsGetResponse_items) SetState(value *string)() {
    m.state = value
}
// SetSubtreeOperationCount sets the subtreeOperationCount property value. The count of the number of child Operations below this one, recursively.
func (m *V1beta1AsyncOperationsGetResponse_items) SetSubtreeOperationCount(value *int32)() {
    m.subtreeOperationCount = value
}
// SetSuggestedPollingIntervalSeconds sets the suggestedPollingIntervalSeconds property value. This attribute suggests a suitable interval to use when polling for progress. Where specified this will be based on the frequency with which the operation is likely to be updated.
func (m *V1beta1AsyncOperationsGetResponse_items) SetSuggestedPollingIntervalSeconds(value *int32)() {
    m.suggestedPollingIntervalSeconds = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1AsyncOperationsGetResponse_items) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1AsyncOperationsGetResponse_items) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetUserId sets the userId property value. The ID or email address of the user that initiated the operation.
func (m *V1beta1AsyncOperationsGetResponse_items) SetUserId(value *string)() {
    m.userId = value
}
type V1beta1AsyncOperationsGetResponse_itemsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDetails()(V1beta1AsyncOperationsGetResponse_items_additionalDetailsable)
    GetAssociatedResources()([]V1beta1AsyncOperationsGetResponse_items_associatedResourcesable)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDisplayName()(*string)
    GetEndedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetError()(V1beta1AsyncOperationsGetResponse_items_errorable)
    GetEstimatedRunningDurationMinutes()(*int32)
    GetGeneration()(*int64)
    GetGroups()([]V1beta1AsyncOperationsGetResponse_items_groupsable)
    GetHasChildOperations()(*bool)
    GetHealthStatus()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLogMessages()([]V1beta1AsyncOperationsGetResponse_items_logMessagesable)
    GetName()(*string)
    GetParent()(V1beta1AsyncOperationsGetResponse_items_parentable)
    GetProgressPercent()(*int32)
    GetRecommendations()([]V1beta1AsyncOperationsGetResponse_items_recommendationsable)
    GetResourceUri()(*string)
    GetRootOperation()(V1beta1AsyncOperationsGetResponse_items_rootOperationable)
    GetServices()([]string)
    GetSourceResourceUri()(*string)
    GetStartedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetState()(*string)
    GetSubtreeOperationCount()(*int32)
    GetSuggestedPollingIntervalSeconds()(*int32)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetUserId()(*string)
    SetAdditionalDetails(value V1beta1AsyncOperationsGetResponse_items_additionalDetailsable)()
    SetAssociatedResources(value []V1beta1AsyncOperationsGetResponse_items_associatedResourcesable)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDisplayName(value *string)()
    SetEndedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetError(value V1beta1AsyncOperationsGetResponse_items_errorable)()
    SetEstimatedRunningDurationMinutes(value *int32)()
    SetGeneration(value *int64)()
    SetGroups(value []V1beta1AsyncOperationsGetResponse_items_groupsable)()
    SetHasChildOperations(value *bool)()
    SetHealthStatus(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLogMessages(value []V1beta1AsyncOperationsGetResponse_items_logMessagesable)()
    SetName(value *string)()
    SetParent(value V1beta1AsyncOperationsGetResponse_items_parentable)()
    SetProgressPercent(value *int32)()
    SetRecommendations(value []V1beta1AsyncOperationsGetResponse_items_recommendationsable)()
    SetResourceUri(value *string)()
    SetRootOperation(value V1beta1AsyncOperationsGetResponse_items_rootOperationable)()
    SetServices(value []string)()
    SetSourceResourceUri(value *string)()
    SetStartedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetState(value *string)()
    SetSubtreeOperationCount(value *int32)()
    SetSuggestedPollingIntervalSeconds(value *int32)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetUserId(value *string)()
}
