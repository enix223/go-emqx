package emqx

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ClientConfig client config
type ClientConfig struct {
	// EMQX API base url, defaul to http://localhost:8080
	BaseURL string
	// EMQX Application ID
	AppID string
	// EMQX Application Secret
	AppSecret string
	// EMQX client timeout
	Timeout time.Duration
}

// APIClient EMQX RESTFul API client
type APIClient struct {
	// BaseURL emqx RESTFul address
	BaseURL    string
	httpClient *http.Client
	appID      string
	appSecret  string
	token      string
}

// NewAPIClient create client
func NewAPIClient(c ClientConfig) Client {
	a := &APIClient{
		httpClient: http.DefaultClient,
		appID:      c.AppID,
		appSecret:  c.AppSecret,
	}

	if c.BaseURL == "" {
		c.BaseURL = "http://localhost:8080"
	}
	a.BaseURL = c.BaseURL

	if c.Timeout == 0 {
		c.Timeout = time.Second * 5
	}
	a.httpClient.Timeout = c.Timeout

	a.updateToken(c.AppID, c.AppSecret)

	return a
}

// UpdateToken update token with appID and appSecret
func (a *APIClient) updateToken(appID, appSecret string) {
	str := fmt.Sprintf("%s:%s", appID, appSecret)
	a.token = fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(str)))
}

// makeRequest makeRequest
func (a *APIClient) makeRequest(method, endpoint string, payload []byte, resp interface{}) error {
	url := fmt.Sprintf("%s/%s", a.BaseURL, endpoint)
	var body io.Reader
	if payload != nil {
		body = bytes.NewBuffer(payload)
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	request.Header = http.Header{
		"Authorization": []string{a.token},
		"Content-Type":  []string{"application/json"},
	}

	response, err := a.httpClient.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(content, resp)
	if err != nil {
		return errors.New(string(content))
	}

	return nil
}

// ListAllAPI ListAllAPI
// List all API describe
// GET api/v3/
func (a *APIClient) ListAllAPI() (*ListAPIResponseV3, error) {
	var resp ListAPIResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCluster ListCluster
// List all Cluster
// GET api/v3/brokers/
func (a *APIClient) ListCluster() (*ListClusterResponseV3, error) {
	var resp ListClusterResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/brokers/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeInfo GetNodeInfo
// Retrieve Info of a Node
// GET api/v3/brokers/${node}
func (a *APIClient) GetNodeInfo(node string) (*NodeInfoResponseV3, error) {
	var resp NodeInfoResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/brokers/%s", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeStats ListNodeStats
// List Statistics of All Nodes in the Cluster
// GET api/v3/nodes/
func (a *APIClient) ListNodeStats() (*ListNodeStatResponseV3, error) {
	var resp ListNodeStatResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/nodes/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeStat GetNodeStat
// Retrieve Statistics of a Specific Node
// GET api/v3/nodes/${node}
func (a *APIClient) GetNodeStat(node string) (*NodeStatResponseV3, error) {
	var resp NodeStatResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterConnections ListClusterConnections
// List all Connections in the Cluster
// GET api/v3/connections/
func (a *APIClient) ListClusterConnections() (*ListClusterConnectionsResponseV3, error) {
	var resp ListClusterConnectionsResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/connections/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeConnections ListNodeConnections
// List all Connections in the node
// GET api/v3/nodes/${node}/connections/
func (a *APIClient) ListNodeConnections(node string) (*ListNodeConnectionsResponseV3, error) {
	var resp ListNodeConnectionsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/connections", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterConnection GetClusterConnection
// Retrieve a Connection in the Cluster¶
// GET api/v3/connections/${clientid}
func (a *APIClient) GetClusterConnection(clientid string) (*ClusterConnectionResponseV3, error) {
	var resp ClusterConnectionResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/connections/%s", clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeConnection GetNodeConnection
// Retrieve a Connection on a node
// GET api/v3/connections/${clientid}
func (a *APIClient) GetNodeConnection(clientid string) (*NodeConnectionResponseV3, error) {
	var resp NodeConnectionResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/connections/%s", clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterSessions ListClusterSessions
// List all Sessions in the Cluster
// GET api/v3/sessions/
func (a *APIClient) ListClusterSessions() (*ListClusterSessionsResponseV3, error) {
	var resp ListClusterSessionsResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/sessions/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterSession GetClusterSession
// Retrieve a Session in the Cluster
// GET api/v3/sessions/${clientid}
func (a *APIClient) GetClusterSession(clientid string) (*GetClusterSessionResponseV3, error) {
	var resp GetClusterSessionResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/sessions/%s", clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeSession ListNodeSession
// List all Sessions on a Node
// GET api/v3/nodes/${node}/sessions/
func (a *APIClient) ListNodeSession(node string) (*ListNodeSessionsResponseV3, error) {
	var resp ListNodeSessionsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/sessions/", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeSession GetNodeSession
// Retrieve a Session on a Node
// GET api/v3/nodes/${node}/sessions/${clientid}
func (a *APIClient) GetNodeSession(node, clientid string) (*GetNodeSessionResponseV3, error) {
	var resp GetNodeSessionResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/sessions/%s", node, clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterSubscriptions ListClusterSubscriptions
// List all Subscriptions in the Cluster
// GET api/v3/subscriptions/
func (a *APIClient) ListClusterSubscriptions() (*ListClusterSubscriptionsResponseV3, error) {
	var resp ListClusterSubscriptionsResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/subscriptions/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterConnSubscriptions ListClusterConnSubscriptions
// List Subscriptions of a Connection in the Cluster
// GET api/v3/subscriptions/${clientid}
func (a *APIClient) ListClusterConnSubscriptions(clientid string) (*ListClusterConnSubscriptionsResponseV3, error) {
	var resp ListClusterConnSubscriptionsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/subscriptions/%s", clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeSubscriptions ListNodeSubscriptions
// List all Subscriptions in a node
// GET api/v3/nodes/${node}/subscriptions/
func (a *APIClient) ListNodeSubscriptions(node string) (*ListNodeSubscriptionsResponseV3, error) {
	var resp ListNodeSubscriptionsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/subscriptions/", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeClientSubscriptions ListNodeClientSubscriptions
// List Subscriptions of a Client on a node
// GET api/v3/nodes/${node}/subscriptions/${clientid}
func (a *APIClient) ListNodeClientSubscriptions(node, clientid string) (*ListNodeClientSubscriptionsResponseV3, error) {
	var resp ListNodeClientSubscriptionsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/subscriptions/%s", node, clientid), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRoutes ListRoutes
// List all Routes in the Cluster
// GET api/v3/routes/
func (a *APIClient) ListRoutes() (*ListRoutesResponseV3, error) {
	var resp ListRoutesResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/routes/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTopicRoute GetTopicRoute
// GetTopicRoutesResponseV3 - Retrieve a Route of Topic in the Cluster
// GET api/v3/routes/${topic}
func (a *APIClient) GetTopicRoute(topic string) (*GetTopicRoutesResponseV3, error) {
	var resp GetTopicRoutesResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/routes/%s", topic), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PublishMessage PublishMessage
// Publish message request
// POST api/v3/mqtt/publish
func (a *APIClient) PublishMessage(req *PublishMessageRequestV3) (*NoContentResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var resp NoContentResponse
	err = a.makeRequest(http.MethodPost, "api/v3/mqtt/publish", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSubscription CreateSubscription
// create subscription
// POST api/v3/mqtt/subscribe
func (a *APIClient) CreateSubscription(req *CreateSubscriptionRequestV3) (*NoContentResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var resp NoContentResponse
	err = a.makeRequest(http.MethodPost, "api/v3/mqtt/subscribe", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Unsubscribe Unsubscribe
// unsubscribe
// POST api/v3/mqtt/unsubscribe
func (a *APIClient) Unsubscribe(req *UnSubscribeRequestV3) (*NoContentResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var resp NoContentResponse
	err = a.makeRequest(http.MethodPost, "api/v3/mqtt/unsubscribe", payload, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//
// Plugins
//

// ListClusterPlugins ListClusterPlugins
// ListClusterPluginResponseV3 - List all Plugins of Cluster
// GET api/v3/plugins/
func (a *APIClient) ListClusterPlugins() (*ListClusterPluginResponseV3, error) {
	var resp ListClusterPluginResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/plugins/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodePlugins ListNodePlugins
// List all plugins in a node
// GET api/v3/nodes/${node}/plugins/
func (a *APIClient) ListNodePlugins(node string) (*ListNodePluginResponseV3, error) {
	var resp ListNodePluginResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/plugins/", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartNodePlugins Start a plugin
// PUT api/v3/nodes/${node}/plugins/${plugin}/load
func (a *APIClient) StartNodePlugins(node, plugin string) (*NoContentResponse, error) {
	var resp NoContentResponse
	err := a.makeRequest(http.MethodPut, fmt.Sprintf("api/v3/nodes/%s/plugins/%s/load", node, plugin), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopNodePlugins stop a plugin
// PUT api/v3/nodes/${node}/plugins/${plugin}/unload
func (a *APIClient) StopNodePlugins(node, plugin string) (*NoContentResponse, error) {
	var resp NoContentResponse
	err := a.makeRequest(http.MethodPut, fmt.Sprintf("api/v3/nodes/%s/plugins/%s/unload", node, plugin), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//
// Listeners
//

// ListClusterListeners List all listeners of Cluster
// GET api/v3/listeners/
func (a *APIClient) ListClusterListeners() (*ListClusterListenersResponseV3, error) {
	var resp ListClusterListenersResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/listeners/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeListeners List all listeners in a node
// GET api/v3/nodes/${node}/plugins/
func (a *APIClient) ListNodeListeners(node string) (*ListNodeListenerResponseV3, error) {
	var resp ListNodeListenerResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/listeners/", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

//
// Metrics
//

// ListClusterMetrics List all metrics of Cluster
// GET api/v3/metrics/
func (a *APIClient) ListClusterMetrics() (*ListClusterMetricsResponseV3, error) {
	var resp ListClusterMetricsResponseV3
	err := a.makeRequest(http.MethodGet, "api/v3/metrics/", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeMetrics get all metrics in a node
// GET api/v3/nodes/${node}/metrics/
func (a *APIClient) GetNodeMetrics(node string) (*GetNodeMetricsResponseV3, error) {
	var resp GetNodeMetricsResponseV3
	err := a.makeRequest(http.MethodGet, fmt.Sprintf("api/v3/nodes/%s/metrics/", node), nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
