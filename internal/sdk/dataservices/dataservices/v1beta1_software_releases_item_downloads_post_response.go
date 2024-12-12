package dataservices

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareReleasesItemDownloadsPostResponse the values needed to download a file within a Software Release.
type V1beta1SoftwareReleasesItemDownloadsPostResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The URL needed to download the file.
    downloadUrl *string
    // The expiresAt property
    expiresAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of a file in a Software Release.
    filename *string
    // The ID of the resource.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The SHA-256 checksum of a file in a Software Release in hexadecimal.
    sha256Checksum *string
    // The type of resource.
    typeEscaped *string
}
// NewV1beta1SoftwareReleasesItemDownloadsPostResponse instantiates a new V1beta1SoftwareReleasesItemDownloadsPostResponse and sets the default values.
func NewV1beta1SoftwareReleasesItemDownloadsPostResponse()(*V1beta1SoftwareReleasesItemDownloadsPostResponse) {
    m := &V1beta1SoftwareReleasesItemDownloadsPostResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SoftwareReleasesItemDownloadsPostResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareReleasesItemDownloadsPostResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareReleasesItemDownloadsPostResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDownloadUrl gets the downloadUrl property value. The URL needed to download the file.
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetDownloadUrl()(*string) {
    return m.downloadUrl
}
// GetExpiresAt gets the expiresAt property value. The expiresAt property
// returns a *Time when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expiresAt
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["downloadUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDownloadUrl(val)
        }
        return nil
    }
    res["expiresAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiresAt(val)
        }
        return nil
    }
    res["filename"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilename(val)
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
    res["sha256Checksum"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256Checksum(val)
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
    return res
}
// GetFilename gets the filename property value. The name of a file in a Software Release.
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetFilename()(*string) {
    return m.filename
}
// GetId gets the id property value. The ID of the resource.
// returns a *UUID when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetSha256Checksum gets the sha256Checksum property value. The SHA-256 checksum of a file in a Software Release in hexadecimal.
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetSha256Checksum()(*string) {
    return m.sha256Checksum
}
// GetTypeEscaped gets the type property value. The type of resource.
// returns a *string when successful
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) GetTypeEscaped()(*string) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("downloadUrl", m.GetDownloadUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("expiresAt", m.GetExpiresAt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filename", m.GetFilename())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256Checksum", m.GetSha256Checksum())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDownloadUrl sets the downloadUrl property value. The URL needed to download the file.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetDownloadUrl(value *string)() {
    m.downloadUrl = value
}
// SetExpiresAt sets the expiresAt property value. The expiresAt property
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expiresAt = value
}
// SetFilename sets the filename property value. The name of a file in a Software Release.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetFilename(value *string)() {
    m.filename = value
}
// SetId sets the id property value. The ID of the resource.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetSha256Checksum sets the sha256Checksum property value. The SHA-256 checksum of a file in a Software Release in hexadecimal.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetSha256Checksum(value *string)() {
    m.sha256Checksum = value
}
// SetTypeEscaped sets the type property value. The type of resource.
func (m *V1beta1SoftwareReleasesItemDownloadsPostResponse) SetTypeEscaped(value *string)() {
    m.typeEscaped = value
}
type V1beta1SoftwareReleasesItemDownloadsPostResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDownloadUrl()(*string)
    GetExpiresAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFilename()(*string)
    GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetSha256Checksum()(*string)
    GetTypeEscaped()(*string)
    SetDownloadUrl(value *string)()
    SetExpiresAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFilename(value *string)()
    SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetSha256Checksum(value *string)()
    SetTypeEscaped(value *string)()
}
