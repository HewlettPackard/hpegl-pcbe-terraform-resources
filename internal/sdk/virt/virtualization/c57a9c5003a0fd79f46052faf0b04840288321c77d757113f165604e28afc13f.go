package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo image deprecation status properties on the image.
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The alternative option specified by the Publisher for this image when this image is deprecated.
    alternativeOption *string
    // The state of the image.
    imageState *string
    // Time when the image will be deprecated.
    scheduledDeprecationTime *string
}
// NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo instantiates a new V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo and sets the default values.
func NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo()(*V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) {
    m := &V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlternativeOption gets the alternativeOption property value. The alternative option specified by the Publisher for this image when this image is deprecated.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) GetAlternativeOption()(*string) {
    return m.alternativeOption
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alternativeOption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeOption(val)
        }
        return nil
    }
    res["imageState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageState(val)
        }
        return nil
    }
    res["scheduledDeprecationTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledDeprecationTime(val)
        }
        return nil
    }
    return res
}
// GetImageState gets the imageState property value. The state of the image.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) GetImageState()(*string) {
    return m.imageState
}
// GetScheduledDeprecationTime gets the scheduledDeprecationTime property value. Time when the image will be deprecated.
// returns a *string when successful
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) GetScheduledDeprecationTime()(*string) {
    return m.scheduledDeprecationTime
}
// Serialize serializes information the current object
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alternativeOption", m.GetAlternativeOption())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("imageState", m.GetImageState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("scheduledDeprecationTime", m.GetScheduledDeprecationTime())
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
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlternativeOption sets the alternativeOption property value. The alternative option specified by the Publisher for this image when this image is deprecated.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) SetAlternativeOption(value *string)() {
    m.alternativeOption = value
}
// SetImageState sets the imageState property value. The state of the image.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) SetImageState(value *string)() {
    m.imageState = value
}
// SetScheduledDeprecationTime sets the scheduledDeprecationTime property value. Time when the image will be deprecated.
func (m *V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfo) SetScheduledDeprecationTime(value *string)() {
    m.scheduledDeprecationTime = value
}
type V1beta1CspMachineImagesItemCspMachineImagesGetResponse_cspInfoMember2_imageDeprecationStatusInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlternativeOption()(*string)
    GetImageState()(*string)
    GetScheduledDeprecationTime()(*string)
    SetAlternativeOption(value *string)()
    SetImageState(value *string)()
    SetScheduledDeprecationTime(value *string)()
}
