package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareComponentsItemInstallReleaseGetResponse a release of software.
type V1beta1SoftwareComponentsItemInstallReleaseGetResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdAt property
    createdAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The customer application identifier
    customerId *string
    // Whether the files within the Software Release can be downloaded.
    downloadable *bool
    // A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
    generation *int64
    // An identifier for the resource, usually a UUID.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // A system specified name for the resource.
    name *string
    // The self reference for this resource.
    resourceUri *string
    // The metadata for files within a Software Release.
    signature V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable
    // The metadata for files within a Software Release.
    software V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable
    // The Software Component a Software Release belongs to.
    softwareComponent V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable
    // The type of resource.
    typeEscaped *string
    // The updatedAt property
    updatedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The version of a Software Release.
    version *string
}
// NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse instantiates a new V1beta1SoftwareComponentsItemInstallReleaseGetResponse and sets the default values.
func NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse()(*V1beta1SoftwareComponentsItemInstallReleaseGetResponse) {
    m := &V1beta1SoftwareComponentsItemInstallReleaseGetResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedAt gets the createdAt property value. The createdAt property
// returns a *Time when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdAt
}
// GetCustomerId gets the customerId property value. The customer application identifier
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetCustomerId()(*string) {
    return m.customerId
}
// GetDownloadable gets the downloadable property value. Whether the files within the Software Release can be downloaded.
// returns a *bool when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetDownloadable()(*bool) {
    return m.downloadable
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["downloadable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadable(val)
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
    res["signature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignature(val.(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable))
        }
        return nil
    }
    res["software"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftware(val.(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable))
        }
        return nil
    }
    res["softwareComponent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSoftwareComponent(val.(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable))
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
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetGeneration gets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
// returns a *int64 when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetGeneration()(*int64) {
    return m.generation
}
// GetId gets the id property value. An identifier for the resource, usually a UUID.
// returns a *UUID when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetName gets the name property value. A system specified name for the resource.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetName()(*string) {
    return m.name
}
// GetResourceUri gets the resourceUri property value. The self reference for this resource.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetResourceUri()(*string) {
    return m.resourceUri
}
// GetSignature gets the signature property value. The metadata for files within a Software Release.
// returns a V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetSignature()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable) {
    return m.signature
}
// GetSoftware gets the software property value. The metadata for files within a Software Release.
// returns a V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetSoftware()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable) {
    return m.software
}
// GetSoftwareComponent gets the softwareComponent property value. The Software Component a Software Release belongs to.
// returns a V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetSoftwareComponent()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable) {
    return m.softwareComponent
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// GetUpdatedAt gets the updatedAt property value. The updatedAt property
// returns a *Time when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.updatedAt
}
// GetVersion gets the version property value. The version of a Software Release.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdAt", m.GetCreatedAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("downloadable", m.GetDownloadable())
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
        err := writer.WriteObjectValue("signature", m.GetSignature())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("software", m.GetSoftware())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("softwareComponent", m.GetSoftwareComponent())
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
        err := writer.WriteStringValue("version", m.GetVersion())
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
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedAt sets the createdAt property value. The createdAt property
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdAt = value
}
// SetCustomerId sets the customerId property value. The customer application identifier
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetCustomerId(value *string)() {
    m.customerId = value
}
// SetDownloadable sets the downloadable property value. Whether the files within the Software Release can be downloaded.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetDownloadable(value *bool)() {
    m.downloadable = value
}
// SetGeneration sets the generation property value. A monotonically increasing value. This value updates when the resource is updated and can be used as a short way to determine if a resource has changed or which of two different copies of a resource is more up to date.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetGeneration(value *int64)() {
    m.generation = value
}
// SetId sets the id property value. An identifier for the resource, usually a UUID.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetName sets the name property value. A system specified name for the resource.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetName(value *string)() {
    m.name = value
}
// SetResourceUri sets the resourceUri property value. The self reference for this resource.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetResourceUri(value *string)() {
    m.resourceUri = value
}
// SetSignature sets the signature property value. The metadata for files within a Software Release.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetSignature(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable)() {
    m.signature = value
}
// SetSoftware sets the software property value. The metadata for files within a Software Release.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetSoftware(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable)() {
    m.software = value
}
// SetSoftwareComponent sets the softwareComponent property value. The Software Component a Software Release belongs to.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetSoftwareComponent(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable)() {
    m.softwareComponent = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
// SetUpdatedAt sets the updatedAt property value. The updatedAt property
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.updatedAt = value
}
// SetVersion sets the version property value. The version of a Software Release.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SoftwareComponentsItemInstallReleaseGetResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCustomerId()(*string)
    GetDownloadable()(*bool)
    GetGeneration()(*int64)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetName()(*string)
    GetResourceUri()(*string)
    GetSignature()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable)
    GetSoftware()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable)
    GetSoftwareComponent()(V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable)
    GetTypeEscaped()(*string)
    GetUpdatedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetCreatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCustomerId(value *string)()
    SetDownloadable(value *bool)()
    SetGeneration(value *int64)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetName(value *string)()
    SetResourceUri(value *string)()
    SetSignature(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable)()
    SetSoftware(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareable)()
    SetSoftwareComponent(value V1beta1SoftwareComponentsItemInstallReleaseGetResponse_softwareComponentable)()
    SetTypeEscaped(value *string)()
    SetUpdatedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
