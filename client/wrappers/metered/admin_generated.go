// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package metered

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/metered.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/client/admin"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/common/types"
)

// adminClient implements admin.Client interface instrumented with retries
type adminClient struct {
	client        admin.Client
	metricsClient metrics.Client
}

// NewAdminClient creates a new instance of adminClient with retry policy
func NewAdminClient(client admin.Client, metricsClient metrics.Client) admin.Client {
	return &adminClient{
		client:        client,
		metricsClient: metricsClient,
	}
}

func (c *adminClient) AddSearchAttribute(ctx context.Context, ap1 *types.AddSearchAttributeRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientAddSearchAttributeScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientAddSearchAttributeScope, metrics.CadenceClientLatency)
	err = c.client.AddSearchAttribute(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientAddSearchAttributeScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) CloseShard(ctx context.Context, cp1 *types.CloseShardRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientCloseShardScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientCloseShardScope, metrics.CadenceClientLatency)
	err = c.client.CloseShard(ctx, cp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientCloseShardScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) CountDLQMessages(ctx context.Context, cp1 *types.CountDLQMessagesRequest, p1 ...yarpc.CallOption) (cp2 *types.CountDLQMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientCountDLQMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientCountDLQMessagesScope, metrics.CadenceClientLatency)
	cp2, err = c.client.CountDLQMessages(ctx, cp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientCountDLQMessagesScope, metrics.CadenceClientFailures)
	}
	return cp2, err
}

func (c *adminClient) DeleteWorkflow(ctx context.Context, ap1 *types.AdminDeleteWorkflowRequest, p1 ...yarpc.CallOption) (ap2 *types.AdminDeleteWorkflowResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDeleteWorkflowScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDeleteWorkflowScope, metrics.CadenceClientLatency)
	ap2, err = c.client.DeleteWorkflow(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDeleteWorkflowScope, metrics.CadenceClientFailures)
	}
	return ap2, err
}

func (c *adminClient) DescribeCluster(ctx context.Context, p1 ...yarpc.CallOption) (dp1 *types.DescribeClusterResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDescribeClusterScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDescribeClusterScope, metrics.CadenceClientLatency)
	dp1, err = c.client.DescribeCluster(ctx, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDescribeClusterScope, metrics.CadenceClientFailures)
	}
	return dp1, err
}

func (c *adminClient) DescribeHistoryHost(ctx context.Context, dp1 *types.DescribeHistoryHostRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeHistoryHostResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDescribeHistoryHostScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDescribeHistoryHostScope, metrics.CadenceClientLatency)
	dp2, err = c.client.DescribeHistoryHost(ctx, dp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDescribeHistoryHostScope, metrics.CadenceClientFailures)
	}
	return dp2, err
}

func (c *adminClient) DescribeQueue(ctx context.Context, dp1 *types.DescribeQueueRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeQueueResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDescribeQueueScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDescribeQueueScope, metrics.CadenceClientLatency)
	dp2, err = c.client.DescribeQueue(ctx, dp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDescribeQueueScope, metrics.CadenceClientFailures)
	}
	return dp2, err
}

func (c *adminClient) DescribeShardDistribution(ctx context.Context, dp1 *types.DescribeShardDistributionRequest, p1 ...yarpc.CallOption) (dp2 *types.DescribeShardDistributionResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDescribeShardDistributionScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDescribeShardDistributionScope, metrics.CadenceClientLatency)
	dp2, err = c.client.DescribeShardDistribution(ctx, dp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDescribeShardDistributionScope, metrics.CadenceClientFailures)
	}
	return dp2, err
}

func (c *adminClient) DescribeWorkflowExecution(ctx context.Context, ap1 *types.AdminDescribeWorkflowExecutionRequest, p1 ...yarpc.CallOption) (ap2 *types.AdminDescribeWorkflowExecutionResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientDescribeWorkflowExecutionScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientDescribeWorkflowExecutionScope, metrics.CadenceClientLatency)
	ap2, err = c.client.DescribeWorkflowExecution(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientDescribeWorkflowExecutionScope, metrics.CadenceClientFailures)
	}
	return ap2, err
}

func (c *adminClient) GetCrossClusterTasks(ctx context.Context, gp1 *types.GetCrossClusterTasksRequest, p1 ...yarpc.CallOption) (gp2 *types.GetCrossClusterTasksResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetCrossClusterTasksScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetCrossClusterTasksScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetCrossClusterTasks(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetCrossClusterTasksScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) GetDLQReplicationMessages(ctx context.Context, gp1 *types.GetDLQReplicationMessagesRequest, p1 ...yarpc.CallOption) (gp2 *types.GetDLQReplicationMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetDLQReplicationMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetDLQReplicationMessagesScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetDLQReplicationMessages(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetDLQReplicationMessagesScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) GetDomainIsolationGroups(ctx context.Context, request *types.GetDomainIsolationGroupsRequest, opts ...yarpc.CallOption) (gp1 *types.GetDomainIsolationGroupsResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetDomainIsolationGroupsScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetDomainIsolationGroupsScope, metrics.CadenceClientLatency)
	gp1, err = c.client.GetDomainIsolationGroups(ctx, request, opts...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetDomainIsolationGroupsScope, metrics.CadenceClientFailures)
	}
	return gp1, err
}

func (c *adminClient) GetDomainReplicationMessages(ctx context.Context, gp1 *types.GetDomainReplicationMessagesRequest, p1 ...yarpc.CallOption) (gp2 *types.GetDomainReplicationMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetDomainReplicationMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetDomainReplicationMessagesScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetDomainReplicationMessages(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetDomainReplicationMessagesScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) GetDynamicConfig(ctx context.Context, gp1 *types.GetDynamicConfigRequest, p1 ...yarpc.CallOption) (gp2 *types.GetDynamicConfigResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetDynamicConfigScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetDynamicConfigScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetDynamicConfig(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetDynamicConfigScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) GetGlobalIsolationGroups(ctx context.Context, request *types.GetGlobalIsolationGroupsRequest, opts ...yarpc.CallOption) (gp1 *types.GetGlobalIsolationGroupsResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetGlobalIsolationGroupsScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetGlobalIsolationGroupsScope, metrics.CadenceClientLatency)
	gp1, err = c.client.GetGlobalIsolationGroups(ctx, request, opts...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetGlobalIsolationGroupsScope, metrics.CadenceClientFailures)
	}
	return gp1, err
}

func (c *adminClient) GetReplicationMessages(ctx context.Context, gp1 *types.GetReplicationMessagesRequest, p1 ...yarpc.CallOption) (gp2 *types.GetReplicationMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetReplicationMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetReplicationMessagesScope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetReplicationMessages(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetReplicationMessagesScope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) GetWorkflowExecutionRawHistoryV2(ctx context.Context, gp1 *types.GetWorkflowExecutionRawHistoryV2Request, p1 ...yarpc.CallOption) (gp2 *types.GetWorkflowExecutionRawHistoryV2Response, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientGetWorkflowExecutionRawHistoryV2Scope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientGetWorkflowExecutionRawHistoryV2Scope, metrics.CadenceClientLatency)
	gp2, err = c.client.GetWorkflowExecutionRawHistoryV2(ctx, gp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientGetWorkflowExecutionRawHistoryV2Scope, metrics.CadenceClientFailures)
	}
	return gp2, err
}

func (c *adminClient) ListDynamicConfig(ctx context.Context, lp1 *types.ListDynamicConfigRequest, p1 ...yarpc.CallOption) (lp2 *types.ListDynamicConfigResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientListDynamicConfigScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientListDynamicConfigScope, metrics.CadenceClientLatency)
	lp2, err = c.client.ListDynamicConfig(ctx, lp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientListDynamicConfigScope, metrics.CadenceClientFailures)
	}
	return lp2, err
}

func (c *adminClient) MaintainCorruptWorkflow(ctx context.Context, ap1 *types.AdminMaintainWorkflowRequest, p1 ...yarpc.CallOption) (ap2 *types.AdminMaintainWorkflowResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientMaintainCorruptWorkflowScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientMaintainCorruptWorkflowScope, metrics.CadenceClientLatency)
	ap2, err = c.client.MaintainCorruptWorkflow(ctx, ap1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientMaintainCorruptWorkflowScope, metrics.CadenceClientFailures)
	}
	return ap2, err
}

func (c *adminClient) MergeDLQMessages(ctx context.Context, mp1 *types.MergeDLQMessagesRequest, p1 ...yarpc.CallOption) (mp2 *types.MergeDLQMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientMergeDLQMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientMergeDLQMessagesScope, metrics.CadenceClientLatency)
	mp2, err = c.client.MergeDLQMessages(ctx, mp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientMergeDLQMessagesScope, metrics.CadenceClientFailures)
	}
	return mp2, err
}

func (c *adminClient) PurgeDLQMessages(ctx context.Context, pp1 *types.PurgeDLQMessagesRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientPurgeDLQMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientPurgeDLQMessagesScope, metrics.CadenceClientLatency)
	err = c.client.PurgeDLQMessages(ctx, pp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientPurgeDLQMessagesScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) ReadDLQMessages(ctx context.Context, rp1 *types.ReadDLQMessagesRequest, p1 ...yarpc.CallOption) (rp2 *types.ReadDLQMessagesResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientReadDLQMessagesScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientReadDLQMessagesScope, metrics.CadenceClientLatency)
	rp2, err = c.client.ReadDLQMessages(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientReadDLQMessagesScope, metrics.CadenceClientFailures)
	}
	return rp2, err
}

func (c *adminClient) ReapplyEvents(ctx context.Context, rp1 *types.ReapplyEventsRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientReapplyEventsScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientReapplyEventsScope, metrics.CadenceClientLatency)
	err = c.client.ReapplyEvents(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientReapplyEventsScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) RefreshWorkflowTasks(ctx context.Context, rp1 *types.RefreshWorkflowTasksRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientRefreshWorkflowTasksScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientRefreshWorkflowTasksScope, metrics.CadenceClientLatency)
	err = c.client.RefreshWorkflowTasks(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientRefreshWorkflowTasksScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) RemoveTask(ctx context.Context, rp1 *types.RemoveTaskRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientRemoveTaskScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientRemoveTaskScope, metrics.CadenceClientLatency)
	err = c.client.RemoveTask(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientRemoveTaskScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) ResendReplicationTasks(ctx context.Context, rp1 *types.ResendReplicationTasksRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientResendReplicationTasksScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientResendReplicationTasksScope, metrics.CadenceClientLatency)
	err = c.client.ResendReplicationTasks(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientResendReplicationTasksScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) ResetQueue(ctx context.Context, rp1 *types.ResetQueueRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientResetQueueScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientResetQueueScope, metrics.CadenceClientLatency)
	err = c.client.ResetQueue(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientResetQueueScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) RespondCrossClusterTasksCompleted(ctx context.Context, rp1 *types.RespondCrossClusterTasksCompletedRequest, p1 ...yarpc.CallOption) (rp2 *types.RespondCrossClusterTasksCompletedResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientRespondCrossClusterTasksCompletedScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientRespondCrossClusterTasksCompletedScope, metrics.CadenceClientLatency)
	rp2, err = c.client.RespondCrossClusterTasksCompleted(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientRespondCrossClusterTasksCompletedScope, metrics.CadenceClientFailures)
	}
	return rp2, err
}

func (c *adminClient) RestoreDynamicConfig(ctx context.Context, rp1 *types.RestoreDynamicConfigRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientRestoreDynamicConfigScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientRestoreDynamicConfigScope, metrics.CadenceClientLatency)
	err = c.client.RestoreDynamicConfig(ctx, rp1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientRestoreDynamicConfigScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) UpdateDomainIsolationGroups(ctx context.Context, request *types.UpdateDomainIsolationGroupsRequest, opts ...yarpc.CallOption) (up1 *types.UpdateDomainIsolationGroupsResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientUpdateDomainIsolationGroupsScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientUpdateDomainIsolationGroupsScope, metrics.CadenceClientLatency)
	up1, err = c.client.UpdateDomainIsolationGroups(ctx, request, opts...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientUpdateDomainIsolationGroupsScope, metrics.CadenceClientFailures)
	}
	return up1, err
}

func (c *adminClient) UpdateDynamicConfig(ctx context.Context, up1 *types.UpdateDynamicConfigRequest, p1 ...yarpc.CallOption) (err error) {
	c.metricsClient.IncCounter(metrics.AdminClientUpdateDynamicConfigScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientUpdateDynamicConfigScope, metrics.CadenceClientLatency)
	err = c.client.UpdateDynamicConfig(ctx, up1, p1...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientUpdateDynamicConfigScope, metrics.CadenceClientFailures)
	}
	return err
}

func (c *adminClient) UpdateGlobalIsolationGroups(ctx context.Context, request *types.UpdateGlobalIsolationGroupsRequest, opts ...yarpc.CallOption) (up1 *types.UpdateGlobalIsolationGroupsResponse, err error) {
	c.metricsClient.IncCounter(metrics.AdminClientUpdateGlobalIsolationGroupsScope, metrics.CadenceClientRequests)

	sw := c.metricsClient.StartTimer(metrics.AdminClientUpdateGlobalIsolationGroupsScope, metrics.CadenceClientLatency)
	up1, err = c.client.UpdateGlobalIsolationGroups(ctx, request, opts...)
	sw.Stop()

	if err != nil {
		c.metricsClient.IncCounter(metrics.AdminClientUpdateGlobalIsolationGroupsScope, metrics.CadenceClientFailures)
	}
	return up1, err
}