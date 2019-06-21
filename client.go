package emqx

// Client EMQX API client
type Client interface {
	// List all API describe
	// GET api/v3/
	ListAllAPI() (*ListAPIResponseV3, error)

	// List all Cluster
	// GET api/v3/brokers/
	ListCluster() (*ListClusterResponseV3, error)

	// Retrieve Info of a Node
	// GET api/v3/brokers/${node}
	GetNodeInfo(node string) (*NodeInfoResponseV3, error)

	// List Statistics of All Nodes in the Cluster
	// GET api/v3/nodes/
	ListNodeStats() (*ListNodeStatResponseV3, error)

	// Retrieve Statistics of a Specific Node
	// GET api/v3/nodes/${node}
	GetNodeStat(node string) (*NodeStatResponseV3, error)

	// List all Connections in the Cluster
	// GET api/v3/connections/
	ListClusterConnections() (*ListClusterConnectionsResponseV3, error)

	// List all Connections in the node
	// GET api/v3/nodes/${node}/connections/
	ListNodeConnections(node string) (*ListNodeConnectionsResponseV3, error)

	// Retrieve a Connection in the Cluster¶
	// GET api/v3/connections/${clientid}
	GetClusterConnection(clientid string) (*ClusterConnectionResponseV3, error)

	// Retrieve a Connection on a node
	// GET api/v3/connections/${clientid}
	GetNodeConnection(clientid string) (*NodeConnectionResponseV3, error)

	// List all Sessions in the Cluster
	// GET api/v3/sessions/
	ListClusterSessions() (*ListClusterSessionsResponseV3, error)

	// Retrieve a Session in the Cluster
	// GET api/v3/sessions/${clientid}
	GetClusterSession(clientid string) (*GetClusterSessionResponseV3, error)

	// List all Sessions on a Node
	// GET api/v3/nodes/${node}/sessions/
	ListNodeSession(node string) (*ListNodeSessionsResponseV3, error)

	// Retrieve a Session on a Node
	// GET api/v3/nodes/${node}/sessions/${clientid}
	GetNodeSession(node, clientid string) (*GetNodeSessionResponseV3, error)

	// List all Subscriptions in the Cluster
	// GET api/v3/subscriptions/
	ListClusterSubscriptions() (*ListClusterSubscriptionsResponseV3, error)

	// List Subscriptions of a Connection in the Cluster
	// GET api/v3/subscriptions/${clientid}
	ListClusterConnSubscriptions(clientid string) (*ListClusterConnSubscriptionsResponseV3, error)

	// List all Subscriptions in a node
	// GET api/v3/nodes/${node}/subscriptions/
	ListNodeSubscriptions(node string) (*ListNodeSubscriptionsResponseV3, error)

	// List Subscriptions of a Client on a node
	// GET api/v3/nodes/${node}/subscriptions/${clientid}
	ListNodeClientSubscriptions(node, clientid string) (*ListNodeClientSubscriptionsResponseV3, error)

	// List all Routes in the Cluster
	// GET api/v3/routes/
	ListRoutes() (*ListRoutesResponseV3, error)

	// GetTopicRoutesResponseV3 - Retrieve a Route of Topic in the Cluster
	// GET api/v3/routes/${topic}
	GetTopicRoute(topic string) (*GetTopicRoutesResponseV3, error)

	// Publish message request
	// POST api/v3/mqtt/publish
	PublishMessage(req *PublishMessageRequestV3) (*NoContentResponse, error)

	// create subscription
	// POST api/v3/mqtt/subscribe
	CreateSubscription(req *CreateSubscriptionRequestV3) (*NoContentResponse, error)

	// unsubscribe
	// POST api/v3/mqtt/unsubscribe
	Unsubscribe(req *UnSubscribeRequestV3) (*NoContentResponse, error)

	// ListClusterPlugins List all Plugins of Cluster
	// GET api/v3/plugins/
	ListClusterPlugins() (*ListClusterPluginResponseV3, error)

	// ListNodePlugins List all plugins in a node
	// GET api/v3/nodes/${node}/plugins/
	ListNodePlugins(node string) (*ListNodePluginResponseV3, error)

	// StartNodePlugins Start a plugin
	// PUT api/v3/nodes/${node}/plugins/${plugin}/load
	StartNodePlugins(node, plugin string) (*NoContentResponse, error)

	// StopNodePlugins stop a plugin
	// PUT api/v3/nodes/${node}/plugins/${plugin}/unload
	StopNodePlugins(node, plugin string) (*NoContentResponse, error)

	// ListClusterListeners List all listeners of Cluster
	// GET api/v3/listeners/
	ListClusterListeners() (*ListClusterListenersResponseV3, error)

	// ListNodeListeners List all listeners in a node
	// GET api/v3/nodes/${node}/plugins/
	ListNodeListeners(node string) (*ListNodeListenerResponseV3, error)

	// ListClusterMetrics List all metrics of Cluster
	// GET api/v3/metrics/
	ListClusterMetrics() (*ListClusterMetricsResponseV3, error)

	// GetNodeMetrics get all metrics in a node
	// GET api/v3/nodes/${node}/metrics/
	GetNodeMetrics(node string) (*GetNodeMetricsResponseV3, error)
}
