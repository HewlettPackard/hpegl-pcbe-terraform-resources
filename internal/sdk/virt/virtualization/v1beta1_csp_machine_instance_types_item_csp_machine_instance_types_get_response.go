package virtualization

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse details of a CSP machine instance type
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // URI for the console screen that displays this resource
    consoleUri *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The cspInfo property
    cspInfo V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable
    // The customer application identifier
    customerId *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *string
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo composed type wrapper for classes V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able, V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo struct {
    // Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able
    v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able
    // Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able
    v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able
}
// NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo instantiates a new V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo and sets the default values.
func NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo()(*V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) {
    m := &V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo{
    }
    return m
}
// CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    return result, nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    if m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1() != nil {
        return m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1().GetFieldDeserializers()
    } else if m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2() != nil {
        return m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2().GetFieldDeserializers()
    }
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) GetIsComposedType()(bool) {
    return true
}
// GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1 gets the V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 property value. Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able) {
    return m.v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1
}
// GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2 gets the V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 property value. Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able) {
    return m.v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1())
        if err != nil {
            return err
        }
    } else if m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2() != nil {
        err := writer.WriteObjectValue("", m.GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1 sets the V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 property value. Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able)() {
    m.v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1 = value
}
// SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2 sets the V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 property value. Composed type representation for type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfo) SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able)() {
    m.v1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2 = value
}
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able)
    GetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able)
    SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember1(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember1able)()
    SetV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseCspInfoMember2(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_cspInfoMember2able)()
}
// NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse instantiates a new V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse and sets the default values.
func NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse()(*V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) {
    m := &V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. URI for the console screen that displays this resource
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCspInfo gets the cspInfo property value. The cspInfo property
// returns a V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetCspInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable) {
    return m.cspInfo
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["consoleUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsoleUri(val)
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
    res["cspInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCspInfo(val.(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable))
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
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
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
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cspInfo", m.GetCspInfo())
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
        err := writer.WriteTimeValue("updatedAt", m.GetUpdatedAt())
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
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. URI for the console screen that displays this resource
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCspInfo sets the cspInfo property value. The cspInfo property
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetCspInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable)() {
    m.cspInfo = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCspInfo()(V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable)
    GetCustomerId()(*string)
    GetGeneration()(*int64)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetConsoleUri(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCspInfo(value V1beta1CspMachineInstanceTypesItemCspMachineInstanceTypesGetResponse_CspMachineInstanceTypesGetResponse_cspInfoable)()
    SetCustomerId(value *string)()
    SetGeneration(value *int64)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
