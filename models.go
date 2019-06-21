package emqx

// NoContentResponse no content
type NoContentResponse struct {
	Code    int
	Message string
}

//
// Cluster & Node
//

// APIV3 api info
type APIV3 struct {
	Name   string
	Path   string
	Descr  string
	Method string
}

// ListAPIResponseV3 - List all API describe
// GET api/v3/
type ListAPIResponseV3 struct {
	Code int
	Data []APIV3
}

// ClusterV3 cluster info
type ClusterV3 struct {
	Datetime   string
	Node       string
	NodeStatus string
	OtpRelease string
	Sysdescr   string
	Uptime     string
	Version    string
}

// ListClusterResponseV3 - List all Cluster
// GET api/v3/brokers/
type ListClusterResponseV3 struct {
	Code int
	Data []ClusterV3
}

// NodeInfoResponseV3 -  Retrieve Info of a Node
// GET api/v3/brokers/${node}
type NodeInfoResponseV3 struct {
	Code int
	Data ClusterV3
}

// NodeStatV3 node stat
type NodeStatV3 struct {
	Connections      int
	Load1            string
	Load15           string
	Load5            string
	MaxFds           int
	MemoryTotal      string
	MemoryUsed       string
	Name             string
	Node             string
	NodeStatus       string
	OtpRelease       string
	ProcessAvailable int
	ProcessUsed      int
	Uptime           string
	Version          string
}

// ListNodeStatResponseV3 - List Statistics of All Nodes in the Cluster
// GET api/v3/nodes/
type ListNodeStatResponseV3 struct {
	Code int
	Data []NodeStatV3
}

// NodeStatResponseV3 - Retrieve Statistics of a Specific Node
// GET api/v3/nodes/${node}
type NodeStatResponseV3 struct {
	Code int
	Data NodeStatV3
}

//
// Connections
//

// ConnectionV3 connection info
type ConnectionV3 struct {
	CleanStart  bool   `json:"clean_start"`
	ClientID    string `json:"client_id"`
	ConnMod     string `json:"conn_mod"`
	ConnectedAT string `json:"connected_at"`
	HeapSize    int    `json:"heap_size"`
	IPAddress   string `json:"ipaddress"`
	IsBridge    bool   `json:"is_bridge"`
	KeepAlive   int    `json:"keepalive"`
	MailboxLen  int    `json:"mailbox_len"`
	Node        string `json:"node"`
	PeerCert    string `json:"peercert"`
	Port        int    `json:"port"`
	ProtoName   string `json:"proto_name"`
	ProtoVer    int    `json:"proto_ver"`
	RecvCnt     int    `json:"recv_cnt"`
	RecvMsg     int    `json:"recv_msg"`
	RecvOct     int    `json:"recv_oct"`
	RecvPkt     int    `json:"recv_pkt"`
	Reductions  int    `json:"reductions"`
	SendCnt     int    `json:"send_cnt"`
	SendMsg     int    `json:"send_msg"`
	SendOct     int    `json:"send_oct"`
	SendPend    int    `json:"send_pend"`
	SendPkt     int    `json:"send_pkt"`
	Username    string `json:"username"`
	Zone        string `json:"zone"`
}

// ListClusterConnectionsResponseV3 - List all Connections in the Cluster
// GET api/v3/connections/
type ListClusterConnectionsResponseV3 struct {
	Code int
	Data []ConnectionV3
}

// ListNodeConnectionsResponseV3 - List all Connections in the node
// GET api/v3/nodes/${node}/connections/
type ListNodeConnectionsResponseV3 struct {
	Code int
	Data []ConnectionV3
}

// ClusterConnectionResponseV3 - Retrieve a Connection in the Cluster¶
// GET api/v3/connections/${clientid}
type ClusterConnectionResponseV3 struct {
	Code int
	Data []ConnectionV3
}

// NodeConnectionResponseV3 - Retrieve a Connection on a node
// GET api/v3/connections/${clientid}
type NodeConnectionResponseV3 struct {
	Code int
	Data []ConnectionV3
}

//
// Sessions
//

// SessionV3 session info
type SessionV3 struct {
	AwaitingRelLen     int    `json:"awaiting_rel_len"`
	Binding            string `json:"binding"`
	CleanStart         bool   `json:"clean_start"`
	ClientID           string `json:"client_id"`
	CreatedAt          string `json:"created_at"`
	DeliverMsg         int    `json:"deliver_msg"`
	EnqueueMsg         int    `json:"enqueue_msg"`
	ExpiryInterval     int    `json:"expiry_interval"`
	HeapSize           int    `json:"heap_size"`
	InflightLen        int    `json:"inflight_len"`
	MailboxLen         int    `json:"mailbox_len"`
	MaxAwaitingRel     int    `json:"max_awaiting_rel"`
	MaxInflight        int    `json:"max_inflight"`
	MaxMqueue          int    `json:"max_mqueue"`
	MaxSubscriptions   int    `json:"max_subscriptions"`
	MqueueDropped      int    `json:"mqueue_dropped"`
	MqueueLen          int    `json:"mqueue_len"`
	Node               string `json:"node"`
	Reductions         int    `json:"reductions"`
	SubscriptionsCount int    `json:"subscriptions_count"`
	Username           string `json:"username"`
}

// ListClusterSessionsResponseV3 - List all Sessions in the Cluster
// GET api/v3/sessions/
type ListClusterSessionsResponseV3 struct {
	Code int
	Data []SessionV3
}

// GetClusterSessionResponseV3 - Retrieve a Session in the Cluster
// GET api/v3/sessions/${clientid}
type GetClusterSessionResponseV3 struct {
	Code int
	Data []SessionV3
}

// ListNodeSessionsResponseV3 - List all Sessions on a Node
// GET api/v3/nodes/${node}/sessions/
type ListNodeSessionsResponseV3 struct {
	Code int
	Data []SessionV3
}

// GetNodeSessionResponseV3 - Retrieve a Session on a Node
// GET api/v3/nodes/${node}/sessions/${clientid}
type GetNodeSessionResponseV3 struct {
	Code int
	Data []SessionV3
}

//
// Subscriptions
//

// SubscriptionV3 subscription
type SubscriptionV3 struct {
	ClientID string `json:"client_id"`
	Node     string `json:"node"`
	Qos      string `json:"qos"`
	Topic    string `json:"topic"`
}

// ListClusterSubscriptionsResponseV3 - List all Subscriptions in the Cluster
// GET api/v3/subscriptions/
type ListClusterSubscriptionsResponseV3 struct {
	Code int
	Data []SubscriptionV3
}

// ListClusterConnSubscriptionsResponseV3 - List Subscriptions of a Connection in the Cluster
// GET api/v3/subscriptions/${clientid}
type ListClusterConnSubscriptionsResponseV3 struct {
	Code int
	Data []SubscriptionV3
}

// ListNodeSubscriptionsResponseV3 - List all Subscriptions in a node
// GET api/v3/nodes/${node}/subscriptions/
type ListNodeSubscriptionsResponseV3 struct {
	Code int
	Data []SubscriptionV3
}

// ListNodeClientSubscriptionsResponseV3 - List Subscriptions of a Client on a node
// GET api/v3/nodes/${node}/subscriptions/${clientid}
type ListNodeClientSubscriptionsResponseV3 struct {
	Code int
	Data []SubscriptionV3
}

//
// Routes
//

// RouteV3 route
type RouteV3 struct {
	Node  string `json:"node"`
	Topic string `json:"topic"`
}

// ListRoutesResponseV3 - List all Routes in the Cluster
// GET api/v3/routes/
type ListRoutesResponseV3 struct {
	Code int
	Data []RouteV3
}

// GetTopicRoutesResponseV3 - Retrieve a Route of Topic in the Cluster
// GET api/v3/routes/${topic}
type GetTopicRoutesResponseV3 struct {
	Code int
	Data []RouteV3
}

//
// Publish/Subscribe
//

// PublishMessageRequestV3 Publish message request
// POST api/v3/mqtt/publish
type PublishMessageRequestV3 struct {
	Topic    string `json:"topic"`
	Payload  string `json:"payload"`
	Qos      int    `json:"qos"`
	Retain   bool   `json:"retain"`
	ClientID string `json:"client_id"`
}

// CreateSubscriptionRequestV3 create subscription
// POST api/v3/mqtt/subscribe
type CreateSubscriptionRequestV3 struct {
	Topic    string `json:"topic"`
	Qos      int    `json:"qos"`
	ClientID string `json:"client_id"`
}

// UnSubscribeRequestV3 unsubscribe
// POST api/v3/mqtt/unsubscribe
type UnSubscribeRequestV3 struct {
	Topic    string `json:"topic"`
	ClientID string `json:"client_id"`
}

//
// Plugins
//

// PluginV3 plugin detail
type PluginV3 struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
}

// NodePluginV3 plugins info list of a node
type NodePluginV3 struct {
	Node    string     `json:"node"`
	Plugins []PluginV3 `json:"plugins"`
}

// ListClusterPluginResponseV3 list all plugins of cluster
// GET api/v3/plugins/
type ListClusterPluginResponseV3 struct {
	Code int
	Data []NodePluginV3
}

// ListNodePluginResponseV3 list all plugins of a node
// GET api/v3/nodes/${node}/plugins/
type ListNodePluginResponseV3 struct {
	Code int
	Data []PluginV3
}

//
// Listeners
//

// ShutdownCountV3 listener shutdown count
type ShutdownCountV3 struct {
	Closed int
	Kicked int
}

// ListenerV3 listener info
type ListenerV3 struct {
	Acceptors     int
	CurrentConns  int
	ListenOn      string
	MaxConns      int
	Protocol      string
	ShutdownCount ShutdownCountV3
}

// NodeListenersV3 listeners info list of a node
type NodeListenersV3 struct {
	Node      string       `json:"node"`
	Listeners []ListenerV3 `json:"listeners"`
}

// ListClusterListenersResponseV3 list all listeners of cluster
// GET api/v3/listeners/
type ListClusterListenersResponseV3 struct {
	Code int
	Data []NodeListenersV3
}

// ListNodeListenerResponseV3 list all listeners of a node
// GET api/v3/nodes/${node}/listeners
type ListNodeListenerResponseV3 struct {
	Code int
	Data []ListenerV3
}

//
// Metrics
//

// MetricsV3 metrics
type MetricsV3 struct {
	BytesReceived             int `json:"bytes/received"`
	PacketsPubrelSent         int `json:"packets/pubrel/sent"`
	PacketsPubcompMissed      int `json:"packets/pubcomp/missed"`
	PacketsSent               int `json:"packets/sent"`
	PacketsPubrelReceived     int `json:"packets/pubrel/received"`
	MessagesQos1Received      int `json:"messages/qos1/received"`
	PacketsPublishReceived    int `json:"packets/publish/received"`
	PacketsAuth               int `json:"packets/auth"`
	MessagesQos0Received      int `json:"messages/qos0/received"`
	PacketsPubcompReceived    int `json:"packets/pubcomp/received"`
	PacketsUnsuback           int `json:"packets/unsuback"`
	PacketsPubrecMissed       int `json:"packets/pubrec/missed"`
	MessagesQos1Sent          int `json:"messages/qos1/sent"`
	MessagesQos2Sent          int `json:"messages/qos2/sent"`
	BytesSent                 int `json:"bytes/sent"`
	MessagesReceived          int `json:"messages/received"`
	MessagesDropped           int `json:"messages/dropped"`
	MessagesQos2Received      int `json:"messages/qos2/received"`
	PacketsConnect            int `json:"packets/connect"`
	MessagesQos0Sent          int `json:"messages/qos0/sent"`
	PacketsDisconnectReceived int `json:"packets/disconnect/received"`
	PacketsPubrecSent         int `json:"packets/pubrec/sent"`
	PacketsPublishSent        int `json:"packets/publish/sent"`
	PacketsPubrecReceived     int `json:"packets/pubrec/received"`
	PacketsReceived           int `json:"packets/received"`
	PacketsUnsubscribe        int `json:"packets/unsubscribe"`
	PacketsSubscribe          int `json:"packets/subscribe"`
	PacketsDisconnectSent     int `json:"packets/disconnect/sent"`
	PacketsPingresp           int `json:"packets/pingresp"`
	MessagesQos2Dropped       int `json:"messages/qos2/dropped"`
	PacketsPubackMissed       int `json:"packets/puback/missed"`
	PacketsPingreq            int `json:"packets/pingreq"`
	PacketsConnack            int `json:"packets/connack"`
	PacketsPubrelMissed       int `json:"packets/pubrel/missed"`
	MessagesSent              int `json:"messages/sent"`
	PacketsSuback             int `json:"packets/suback"`
	MessagesRetained          int `json:"messages/retained"`
	PacketsPubackSent         int `json:"packets/puback/sent"`
	PacketsPubackReceived     int `json:"packets/puback/received"`
	MessagesQos2Expired       int `json:"messages/qos2/expired"`
	MessagesForward           int `json:"messages/forward"`
	MessagesExpired           int `json:"messages/expired"`
	PacketsPubcompS           int `json:"packets/pubcomp/s"`
}

// NodeMetricsV3 plugin info list of a node
type NodeMetricsV3 struct {
	Node    string      `json:"node"`
	Metrics []MetricsV3 `json:"metrics"`
}

// ListClusterMetricsResponseV3 list metrics of the cluster
// GET api/v3/metrics/
type ListClusterMetricsResponseV3 struct {
	Code int
	Data []NodeMetricsV3
}

// GetNodeMetricsResponseV3 get all metrics of a node
// GET api/v3/nodes/${node}/metrics
type GetNodeMetricsResponseV3 struct {
	Code int
	Data MetricsV3
}
