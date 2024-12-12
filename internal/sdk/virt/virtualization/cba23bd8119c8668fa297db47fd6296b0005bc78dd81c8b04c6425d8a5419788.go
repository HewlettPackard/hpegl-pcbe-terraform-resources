package virtualization

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup the name of the subnet group this RDS Instance is associated with
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name property
    name *string
}
// NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup instantiates a new V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup and sets the default values.
func NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup()(*V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) {
    m := &V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name property
func (m *V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroup) SetName(value *string)() {
    m.name = value
}
type V1beta1CspRdsInstancesItemCspRdsInstancesGetResponse_cspInfo_networkInfo_subnetGroupable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    SetName(value *string)()
}
