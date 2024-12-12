package dataservices

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The credentials property
    credentials V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable
    // An IP address or FQDN to address the proxy server
    networkAddress *string
    // Port number of the proxy server
    port *int32
}
// NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy instantiates a new V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy and sets the default values.
func NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy()(*V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) {
    m := &V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// returns a map[string]any when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCredentials gets the credentials property value. The credentials property
// returns a V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) GetCredentials()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable) {
    return m.credentials
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["credentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateV1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCredentials(val.(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable))
        }
        return nil
    }
    res["networkAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkAddress(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    return res
}
// GetNetworkAddress gets the networkAddress property value. An IP address or FQDN to address the proxy server
// returns a *string when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) GetNetworkAddress()(*string) {
    return m.networkAddress
}
// GetPort gets the port property value. Port number of the proxy server
// returns a *int32 when successful
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) GetPort()(*int32) {
    return m.port
}
// Serialize serializes information the current object
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("credentials", m.GetCredentials())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkAddress", m.GetNetworkAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
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
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCredentials sets the credentials property value. The credentials property
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) SetCredentials(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable)() {
    m.credentials = value
}
// SetNetworkAddress sets the networkAddress property value. An IP address or FQDN to address the proxy server
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) SetNetworkAddress(value *string)() {
    m.networkAddress = value
}
// SetPort sets the port property value. Port number of the proxy server
func (m *V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy) SetPort(value *int32)() {
    m.port = value
}
type V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCredentials()(V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable)
    GetNetworkAddress()(*string)
    GetPort()(*int32)
    SetCredentials(value V1beta1DataServicesConnectorsGetResponse_items_interfaces_network_proxy_credentialsable)()
    SetNetworkAddress(value *string)()
    SetPort(value *int32)()
}
