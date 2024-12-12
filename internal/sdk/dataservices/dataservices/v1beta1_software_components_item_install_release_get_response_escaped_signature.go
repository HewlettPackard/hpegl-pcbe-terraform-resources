package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature the metadata for files within a Software Release.
type V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of a file in a Software Release.
    filename *string
    // The SHA-256 checksum of a file in a Software Release in hexadecimal.
    sha256Checksum *string
    // The sizeInBytes property
    sizeInBytes *int64
}
// NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature instantiates a new V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature and sets the default values.
func NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature()(*V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) {
    m := &V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["sizeInBytes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInBytes(val)
        }
        return nil
    }
    return res
}
// GetFilename gets the filename property value. The name of a file in a Software Release.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) GetFilename()(*string) {
    return m.filename
}
// GetSha256Checksum gets the sha256Checksum property value. The SHA-256 checksum of a file in a Software Release in hexadecimal.
// returns a *string when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) GetSha256Checksum()(*string) {
    return m.sha256Checksum
}
// GetSizeInBytes gets the sizeInBytes property value. The sizeInBytes property
// returns a *int64 when successful
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) GetSizeInBytes()(*int64) {
    return m.sizeInBytes
}
// Serialize serializes information the current object
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("filename", m.GetFilename())
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
        err := writer.WriteInt64Value("sizeInBytes", m.GetSizeInBytes())
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
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFilename sets the filename property value. The name of a file in a Software Release.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) SetFilename(value *string)() {
    m.filename = value
}
// SetSha256Checksum sets the sha256Checksum property value. The SHA-256 checksum of a file in a Software Release in hexadecimal.
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) SetSha256Checksum(value *string)() {
    m.sha256Checksum = value
}
// SetSizeInBytes sets the sizeInBytes property value. The sizeInBytes property
func (m *V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signature) SetSizeInBytes(value *int64)() {
    m.sizeInBytes = value
}
type V1beta1SoftwareComponentsItemInstallReleaseGetResponse_signatureable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilename()(*string)
    GetSha256Checksum()(*string)
    GetSizeInBytes()(*int64)
    SetFilename(value *string)()
    SetSha256Checksum(value *string)()
    SetSizeInBytes(value *int64)()
}
