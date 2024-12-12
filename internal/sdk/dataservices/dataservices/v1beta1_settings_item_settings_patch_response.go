package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SettingsItemSettingsPatchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URI for console screen that displays this resource
    consoleUri *string
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Value of the setting
    currentValue *string
    // The customer application identifier
    customerId *string
    // Description of the setting
    description *string
    // Name of the application to be displayed in UI
    externalApplicationName *string
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // Time when this setting was last updated for this account. RFC 3339
    lastUpdatedAt *string
    // UserId of the user who last updated this setting for this account
    lastUpdatedBy *string
    // A system specified name for the resource.
    name *string
    // Next possible value of the setting, which is updated by the workflow
    nextValue *string
    // A JSON array (stored as string) containing possible values of this setting
    possibleValues *string
    // The self reference for this resource.
    resourceUri *string
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewV1beta1SettingsItemSettingsPatchResponse instantiates a new V1beta1SettingsItemSettingsPatchResponse and sets the default values.
func NewV1beta1SettingsItemSettingsPatchResponse()(*V1beta1SettingsItemSettingsPatchResponse) {
    m := &V1beta1SettingsItemSettingsPatchResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SettingsItemSettingsPatchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SettingsItemSettingsPatchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SettingsItemSettingsPatchResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConsoleUri gets the consoleUri property value. The URI for console screen that displays this resource
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetConsoleUri()(*string) {
    return m.consoleUri
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCurrentValue gets the currentValue property value. Value of the setting
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetCurrentValue()(*string) {
    return m.currentValue
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDescription gets the description property value. Description of the setting
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetDescription()(*string) {
    return m.description
}
// GetExternalApplicationName gets the externalApplicationName property value. Name of the application to be displayed in UI
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetExternalApplicationName()(*string) {
    return m.externalApplicationName
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["currentValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentValue(val)
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
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["externalApplicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalApplicationName(val)
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
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["lastUpdatedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedAt(val)
        }
        return nil
    }
    res["lastUpdatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedBy(val)
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
    res["nextValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextValue(val)
        }
        return nil
    }
    res["possibleValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPossibleValues(val)
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
func (m *V1beta1SettingsItemSettingsPatchResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetLastUpdatedAt gets the lastUpdatedAt property value. Time when this setting was last updated for this account. RFC 3339
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetLastUpdatedAt()(*string) {
    return m.lastUpdatedAt
}
// GetLastUpdatedBy gets the lastUpdatedBy property value. UserId of the user who last updated this setting for this account
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetLastUpdatedBy()(*string) {
    return m.lastUpdatedBy
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetName()(*string) {
    return m.name
}
// GetNextValue gets the nextValue property value. Next possible value of the setting, which is updated by the workflow
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetNextValue()(*string) {
    return m.nextValue
}
// GetPossibleValues gets the possibleValues property value. A JSON array (stored as string) containing possible values of this setting
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetPossibleValues()(*string) {
    return m.possibleValues
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SettingsItemSettingsPatchResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// Serialize serializes information the current object
func (m *V1beta1SettingsItemSettingsPatchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("consoleUri", m.GetConsoleUri())
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
        err := writer.WriteStringValue("currentValue", m.GetCurrentValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalApplicationName", m.GetExternalApplicationName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastUpdatedAt", m.GetLastUpdatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastUpdatedBy", m.GetLastUpdatedBy())
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
        err := writer.WriteStringValue("nextValue", m.GetNextValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("possibleValues", m.GetPossibleValues())
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
func (m *V1beta1SettingsItemSettingsPatchResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConsoleUri sets the consoleUri property value. The URI for console screen that displays this resource
func (m *V1beta1SettingsItemSettingsPatchResponse) SetConsoleUri(value *string)() {
    m.consoleUri = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SettingsItemSettingsPatchResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCurrentValue sets the currentValue property value. Value of the setting
func (m *V1beta1SettingsItemSettingsPatchResponse) SetCurrentValue(value *string)() {
    m.currentValue = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SettingsItemSettingsPatchResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDescription sets the description property value. Description of the setting
func (m *V1beta1SettingsItemSettingsPatchResponse) SetDescription(value *string)() {
    m.description = value
}
// SetExternalApplicationName sets the externalApplicationName property value. Name of the application to be displayed in UI
func (m *V1beta1SettingsItemSettingsPatchResponse) SetExternalApplicationName(value *string)() {
    m.externalApplicationName = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SettingsItemSettingsPatchResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SettingsItemSettingsPatchResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetLastUpdatedAt sets the lastUpdatedAt property value. Time when this setting was last updated for this account. RFC 3339
func (m *V1beta1SettingsItemSettingsPatchResponse) SetLastUpdatedAt(value *string)() {
    m.lastUpdatedAt = value
}
// SetLastUpdatedBy sets the lastUpdatedBy property value. UserId of the user who last updated this setting for this account
func (m *V1beta1SettingsItemSettingsPatchResponse) SetLastUpdatedBy(value *string)() {
    m.lastUpdatedBy = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SettingsItemSettingsPatchResponse) SetName(value *string)() {
    m.name = value
}
// SetNextValue sets the nextValue property value. Next possible value of the setting, which is updated by the workflow
func (m *V1beta1SettingsItemSettingsPatchResponse) SetNextValue(value *string)() {
    m.nextValue = value
}
// SetPossibleValues sets the possibleValues property value. A JSON array (stored as string) containing possible values of this setting
func (m *V1beta1SettingsItemSettingsPatchResponse) SetPossibleValues(value *string)() {
    m.possibleValues = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SettingsItemSettingsPatchResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SettingsItemSettingsPatchResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SettingsItemSettingsPatchResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
type V1beta1SettingsItemSettingsPatchResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConsoleUri()(*string)
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCurrentValue()(*string)
    GetCustomerId()(*string)
    GetDescription()(*string)
    GetExternalApplicationName()(*string)
    GetGeneration()(*int64)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetLastUpdatedAt()(*string)
    GetLastUpdatedBy()(*string)
    GetName()(*string)
    GetNextValue()(*string)
    GetPossibleValues()(*string)
    GetResourceUri()(*string)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    SetConsoleUri(value *string)()
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCurrentValue(value *string)()
    SetCustomerId(value *string)()
    SetDescription(value *string)()
    SetExternalApplicationName(value *string)()
    SetGeneration(value *int64)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetLastUpdatedAt(value *string)()
    SetLastUpdatedBy(value *string)()
    SetName(value *string)()
    SetNextValue(value *string)()
    SetPossibleValues(value *string)()
    SetResourceUri(value *string)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
}
