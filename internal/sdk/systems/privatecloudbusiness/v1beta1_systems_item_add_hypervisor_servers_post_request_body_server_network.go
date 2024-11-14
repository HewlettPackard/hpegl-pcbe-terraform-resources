package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // CPU family of the server. Applies only to Alletra dHCI.
    cpuFamily *string
    // Data IP's for the server data network. Applies only to Alletra dHCI.
    dataIpInfos []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable
    // Current IP address of the hypervisorServer. Applies only to Alletra dHCI
    esxIpAddress *string
    // Management IP for the ESX server. Applies only to Alletra dHCI.
    esxMgmtIpInfo V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable
    // Management IP for the server iLO. Applies only to Alletra dHCI.
    iloMgmtIpInfo V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable
    // Serial number of the server. Applies only to Alletra dHCI.
    serial *string
    // Hypervisor version. Applies only to Alletra dHCI.
    version *string
}
// NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork instantiates a new V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork and sets the default values.
func NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork()(*V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) {
    m := &V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCpuFamily gets the cpuFamily property value. CPU family of the server. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetCpuFamily()(*string) {
    return m.cpuFamily
}
// GetDataIpInfos gets the dataIpInfos property value. Data IP's for the server data network. Applies only to Alletra dHCI.
// returns a []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetDataIpInfos()([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable) {
    return m.dataIpInfos
}
// GetEsxIpAddress gets the esxIpAddress property value. Current IP address of the hypervisorServer. Applies only to Alletra dHCI
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetEsxIpAddress()(*string) {
    return m.esxIpAddress
}
// GetEsxMgmtIpInfo gets the esxMgmtIpInfo property value. Management IP for the ESX server. Applies only to Alletra dHCI.
// returns a V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetEsxMgmtIpInfo()(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable) {
    return m.esxMgmtIpInfo
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["cpuFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCpuFamily(val)
        }
        return nil
    }
    res["dataIpInfos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable)
                }
            }
            m.SetDataIpInfos(res)
        }
        return nil
    }
    res["esxIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsxIpAddress(val)
        }
        return nil
    }
    res["esxMgmtIpInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEsxMgmtIpInfo(val.(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable))
        }
        return nil
    }
    res["iloMgmtIpInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloMgmtIpInfo(val.(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable))
        }
        return nil
    }
    res["serial"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerial(val)
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
// GetIloMgmtIpInfo gets the iloMgmtIpInfo property value. Management IP for the server iLO. Applies only to Alletra dHCI.
// returns a V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetIloMgmtIpInfo()(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable) {
    return m.iloMgmtIpInfo
}
// GetSerial gets the serial property value. Serial number of the server. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetSerial()(*string) {
    return m.serial
}
// GetVersion gets the version property value. Hypervisor version. Applies only to Alletra dHCI.
// returns a *string when successful
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("cpuFamily", m.GetCpuFamily())
        if err != nil {
            return err
        }
    }
    if m.GetDataIpInfos() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDataIpInfos()))
        for i, v := range m.GetDataIpInfos() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("dataIpInfos", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("esxIpAddress", m.GetEsxIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("esxMgmtIpInfo", m.GetEsxMgmtIpInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("iloMgmtIpInfo", m.GetIloMgmtIpInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("serial", m.GetSerial())
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
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCpuFamily sets the cpuFamily property value. CPU family of the server. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetCpuFamily(value *string)() {
    m.cpuFamily = value
}
// SetDataIpInfos sets the dataIpInfos property value. Data IP's for the server data network. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetDataIpInfos(value []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable)() {
    m.dataIpInfos = value
}
// SetEsxIpAddress sets the esxIpAddress property value. Current IP address of the hypervisorServer. Applies only to Alletra dHCI
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetEsxIpAddress(value *string)() {
    m.esxIpAddress = value
}
// SetEsxMgmtIpInfo sets the esxMgmtIpInfo property value. Management IP for the ESX server. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetEsxMgmtIpInfo(value V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable)() {
    m.esxMgmtIpInfo = value
}
// SetIloMgmtIpInfo sets the iloMgmtIpInfo property value. Management IP for the server iLO. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetIloMgmtIpInfo(value V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable)() {
    m.iloMgmtIpInfo = value
}
// SetSerial sets the serial property value. Serial number of the server. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetSerial(value *string)() {
    m.serial = value
}
// SetVersion sets the version property value. Hypervisor version. Applies only to Alletra dHCI.
func (m *V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork) SetVersion(value *string)() {
    m.version = value
}
type V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetworkable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCpuFamily()(*string)
    GetDataIpInfos()([]V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable)
    GetEsxIpAddress()(*string)
    GetEsxMgmtIpInfo()(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable)
    GetIloMgmtIpInfo()(V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable)
    GetSerial()(*string)
    GetVersion()(*string)
    SetCpuFamily(value *string)()
    SetDataIpInfos(value []V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_dataIpInfosable)()
    SetEsxIpAddress(value *string)()
    SetEsxMgmtIpInfo(value V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_esxMgmtIpInfoable)()
    SetIloMgmtIpInfo(value V1beta1SystemsItemAddHypervisorServersPostRequestBody_serverNetwork_iloMgmtIpInfoable)()
    SetSerial(value *string)()
    SetVersion(value *string)()
}
