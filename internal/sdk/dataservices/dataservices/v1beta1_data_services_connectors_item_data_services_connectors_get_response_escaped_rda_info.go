package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo this provides details of the RDA for Data Services Connector.
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Is RDA enabled.
    enabled *bool
    // Features support in RDA.
    features []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable
    // Is RDA connected to internet.
    internetConnected *bool
    // RDA midway URL.
    midway *string
    // name of DSC product class.
    productClass *string
    // name of DSC product line.
    productLine *string
    // Is remote access to RDA enabled.
    remoteAccessEnabled *bool
    // Reason for the RDA being in the current state. This will be empty when the RDA is in an OK state.
    stateReason *string
}
// NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo instantiates a new V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo and sets the default values.
func NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo()(*V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) {
    m := &V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnabled gets the enabled property value. Is RDA enabled.
// returns a *bool when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetEnabled()(*bool) {
    return m.enabled
}
// GetFeatures gets the features property value. Features support in RDA.
// returns a []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetFeatures()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable) {
    return m.features
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["features"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable)
                }
            }
            m.SetFeatures(res)
        }
        return nil
    }
    res["internetConnected"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetConnected(val)
        }
        return nil
    }
    res["midway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMidway(val)
        }
        return nil
    }
    res["productClass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductClass(val)
        }
        return nil
    }
    res["productLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductLine(val)
        }
        return nil
    }
    res["remoteAccessEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoteAccessEnabled(val)
        }
        return nil
    }
    res["stateReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateReason(val)
        }
        return nil
    }
    return res
}
// GetInternetConnected gets the internetConnected property value. Is RDA connected to internet.
// returns a *bool when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetInternetConnected()(*bool) {
    return m.internetConnected
}
// GetMidway gets the midway property value. RDA midway URL.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetMidway()(*string) {
    return m.midway
}
// GetProductClass gets the productClass property value. name of DSC product class.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetProductClass()(*string) {
    return m.productClass
}
// GetProductLine gets the productLine property value. name of DSC product line.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetProductLine()(*string) {
    return m.productLine
}
// GetRemoteAccessEnabled gets the remoteAccessEnabled property value. Is remote access to RDA enabled.
// returns a *bool when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetRemoteAccessEnabled()(*bool) {
    return m.remoteAccessEnabled
}
// GetStateReason gets the stateReason property value. Reason for the RDA being in the current state. This will be empty when the RDA is in an OK state.
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) GetStateReason()(*string) {
    return m.stateReason
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetFeatures() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFeatures()))
        for i, v := range m.GetFeatures() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("features", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("internetConnected", m.GetInternetConnected())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("midway", m.GetMidway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productClass", m.GetProductClass())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("productLine", m.GetProductLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("remoteAccessEnabled", m.GetRemoteAccessEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateReason", m.GetStateReason())
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
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnabled sets the enabled property value. Is RDA enabled.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetFeatures sets the features property value. Features support in RDA.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetFeatures(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable)() {
    m.features = value
}
// SetInternetConnected sets the internetConnected property value. Is RDA connected to internet.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetInternetConnected(value *bool)() {
    m.internetConnected = value
}
// SetMidway sets the midway property value. RDA midway URL.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetMidway(value *string)() {
    m.midway = value
}
// SetProductClass sets the productClass property value. name of DSC product class.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetProductClass(value *string)() {
    m.productClass = value
}
// SetProductLine sets the productLine property value. name of DSC product line.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetProductLine(value *string)() {
    m.productLine = value
}
// SetRemoteAccessEnabled sets the remoteAccessEnabled property value. Is remote access to RDA enabled.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetRemoteAccessEnabled(value *bool)() {
    m.remoteAccessEnabled = value
}
// SetStateReason sets the stateReason property value. Reason for the RDA being in the current state. This will be empty when the RDA is in an OK state.
func (m *V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo) SetStateReason(value *string)() {
    m.stateReason = value
}
type V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetFeatures()([]V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable)
    GetInternetConnected()(*bool)
    GetMidway()(*string)
    GetProductClass()(*string)
    GetProductLine()(*string)
    GetRemoteAccessEnabled()(*bool)
    GetStateReason()(*string)
    SetEnabled(value *bool)()
    SetFeatures(value []V1beta1DataServicesConnectorsItemDataServicesConnectorsGetResponse_rdaInfo_featuresable)()
    SetInternetConnected(value *bool)()
    SetMidway(value *string)()
    SetProductClass(value *string)()
    SetProductLine(value *string)()
    SetRemoteAccessEnabled(value *bool)()
    SetStateReason(value *string)()
}
