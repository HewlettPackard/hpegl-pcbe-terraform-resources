package privatecloudbusiness

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo iLO Network Information.
type V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The gateway property
    gateway *string
    // iLO Hostname
    iloHostname *string
    // iLO Management IP address
    iloIp *string
    // The network property
    network *string
    // List of iLO network interfaces
    networkInterfaces []V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable
    // The subnetMask property
    subnetMask *string
}
// NewV1beta1SystemsItemServersGetResponse_items_iloNetworkInfo instantiates a new V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo and sets the default values.
func NewV1beta1SystemsItemServersGetResponse_items_iloNetworkInfo()(*V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) {
    m := &V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1SystemsItemServersGetResponse_items_iloNetworkInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1SystemsItemServersGetResponse_items_iloNetworkInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1SystemsItemServersGetResponse_items_iloNetworkInfo(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gateway"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGateway(val)
        }
        return nil
    }
    res["iloHostname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloHostname(val)
        }
        return nil
    }
    res["iloIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIloIp(val)
        }
        return nil
    }
    res["network"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetwork(val)
        }
        return nil
    }
    res["networkInterfaces"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateV1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable)
                }
            }
            m.SetNetworkInterfaces(res)
        }
        return nil
    }
    res["subnetMask"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnetMask(val)
        }
        return nil
    }
    return res
}
// GetGateway gets the gateway property value. The gateway property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetGateway()(*string) {
    return m.gateway
}
// GetIloHostname gets the iloHostname property value. iLO Hostname
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetIloHostname()(*string) {
    return m.iloHostname
}
// GetIloIp gets the iloIp property value. iLO Management IP address
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetIloIp()(*string) {
    return m.iloIp
}
// GetNetwork gets the network property value. The network property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetNetwork()(*string) {
    return m.network
}
// GetNetworkInterfaces gets the networkInterfaces property value. List of iLO network interfaces
// returns a []V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetNetworkInterfaces()([]V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable) {
    return m.networkInterfaces
}
// GetSubnetMask gets the subnetMask property value. The subnetMask property
// returns a *string when successful
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) GetSubnetMask()(*string) {
    return m.subnetMask
}
// Serialize serializes information the current object
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("gateway", m.GetGateway())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloHostname", m.GetIloHostname())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iloIp", m.GetIloIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("network", m.GetNetwork())
        if err != nil {
            return err
        }
    }
    if m.GetNetworkInterfaces() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetNetworkInterfaces()))
        for i, v := range m.GetNetworkInterfaces() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("networkInterfaces", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnetMask", m.GetSubnetMask())
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
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGateway sets the gateway property value. The gateway property
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetGateway(value *string)() {
    m.gateway = value
}
// SetIloHostname sets the iloHostname property value. iLO Hostname
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetIloHostname(value *string)() {
    m.iloHostname = value
}
// SetIloIp sets the iloIp property value. iLO Management IP address
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetIloIp(value *string)() {
    m.iloIp = value
}
// SetNetwork sets the network property value. The network property
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetNetwork(value *string)() {
    m.network = value
}
// SetNetworkInterfaces sets the networkInterfaces property value. List of iLO network interfaces
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetNetworkInterfaces(value []V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable)() {
    m.networkInterfaces = value
}
// SetSubnetMask sets the subnetMask property value. The subnetMask property
func (m *V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo) SetSubnetMask(value *string)() {
    m.subnetMask = value
}
type V1beta1SystemsItemServersGetResponse_items_iloNetworkInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGateway()(*string)
    GetIloHostname()(*string)
    GetIloIp()(*string)
    GetNetwork()(*string)
    GetNetworkInterfaces()([]V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable)
    GetSubnetMask()(*string)
    SetGateway(value *string)()
    SetIloHostname(value *string)()
    SetIloIp(value *string)()
    SetNetwork(value *string)()
    SetNetworkInterfaces(value []V1beta1SystemsItemServersGetResponse_items_iloNetworkInfo_networkInterfacesable)()
    SetSubnetMask(value *string)()
}
