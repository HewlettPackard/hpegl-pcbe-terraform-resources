package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1StorageLocationsItemStorageLocationsGetResponse required properties of a storage location.
type V1beta1StorageLocationsItemStorageLocationsGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The list of features that support the storage location.
    capabilities []string
    // The city the storage location is in.
    city *string
    // The Cloud Service Provider (CSP) ID.
    cloudServiceProviderId *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // The geographical location of the storage location. Current geographies include: - North America- Europe- Asia Pacific- Middle East and Africa
    geography *string
    // An identifier for the resource.
    id *string
    // A system specified name for the resource.
    name *string
    // The 'self' reference for this resource.
    resourceUri *string
    // The Timezone of the location as defined by the local standard time (non summer time) offset from UTC.
    timezone *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1StorageLocationsItemStorageLocationsGetResponse instantiates a new V1beta1StorageLocationsItemStorageLocationsGetResponse and sets the default values.
func NewV1beta1StorageLocationsItemStorageLocationsGetResponse()(*V1beta1StorageLocationsItemStorageLocationsGetResponse) {
    m := &V1beta1StorageLocationsItemStorageLocationsGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1StorageLocationsItemStorageLocationsGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1StorageLocationsItemStorageLocationsGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1StorageLocationsItemStorageLocationsGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapabilities gets the capabilities property value. The list of features that support the storage location.
// returns a []string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetCapabilities()([]string) {
    return m.capabilities
}
// GetCity gets the city property value. The city the storage location is in.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetCity()(*string) {
    return m.city
}
// GetCloudServiceProviderId gets the cloudServiceProviderId property value. The Cloud Service Provider (CSP) ID.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetCloudServiceProviderId()(*string) {
    return m.cloudServiceProviderId
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
            m.SetCapabilities(res)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["cloudServiceProviderId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloudServiceProviderId(val)
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
    res["geography"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeography(val)
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
    res["timezone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimezone(val)
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
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetGeography gets the geography property value. The geographical location of the storage location. Current geographies include: - North America- Europe- Asia Pacific- Middle East and Africa
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetGeography()(*string) {
    return m.geography
}
// GetId gets the id property value. An identifier for the resource.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetId()(*string) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The 'self' reference for this resource.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTimezone gets the timezone property value. The Timezone of the location as defined by the local standard time (non summer time) offset from UTC.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetTimezone()(*string) {
    return m.timezone
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCapabilities() != nil {
        err := writer.WriteCollectionOfStringValues("capabilities", m.GetCapabilities())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("cloudServiceProviderId", m.GetCloudServiceProviderId())
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
        err := writer.WriteInt64Value("generation", m.GetGeneration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("geography", m.GetGeography())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
        err := writer.WriteStringValue("resourceUri", m.GetResourceUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timezone", m.GetTimezone())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetTypeEscaped())
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
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapabilities sets the capabilities property value. The list of features that support the storage location.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetCapabilities(value []string)() {
    m.capabilities = value
}
// SetCity sets the city property value. The city the storage location is in.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetCity(value *string)() {
    m.city = value
}
// SetCloudServiceProviderId sets the cloudServiceProviderId property value. The Cloud Service Provider (CSP) ID.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetCloudServiceProviderId(value *string)() {
    m.cloudServiceProviderId = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetGeography sets the geography property value. The geographical location of the storage location. Current geographies include: - North America- Europe- Asia Pacific- Middle East and Africa
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetGeography(value *string)() {
    m.geography = value
}
// SetId sets the id property value. An identifier for the resource.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetId(value *string)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The 'self' reference for this resource.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTimezone sets the timezone property value. The Timezone of the location as defined by the local standard time (non summer time) offset from UTC.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetTimezone(value *string)() {
    m.timezone = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1StorageLocationsItemStorageLocationsGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1StorageLocationsItemStorageLocationsGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapabilities()([]string)
    GetCity()(*string)
    GetCloudServiceProviderId()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGeneration()(*int64)
    GetGeography()(*string)
    GetId()(*string)
    GetName()(*string)
    GetResourceUri()(*string)
    GetTimezone()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetCapabilities(value []string)()
    SetCity(value *string)()
    SetCloudServiceProviderId(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGeneration(value *int64)()
    SetGeography(value *string)()
    SetId(value *string)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetTimezone(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
