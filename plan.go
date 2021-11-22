package invdapi

import (
	"errors"
	"github.com/gabrielsoaressantos/invoiced-go/invdendpoint"
	"strings"
)

type Plan struct {
	*Connection
	*invdendpoint.Plan
}

type Plans []*Plan

func (c *Connection) NewPlan() *Plan {
	plan := new(invdendpoint.Plan)
	return &Plan{c, plan}
}

func (c *Plan) Create(plan *Plan) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint

	planResp := new(Plan)

	if plan == nil {
		return nil, errors.New("plan is nil")
	}

	// safe prune file data for creation
	invdPlanDataToCreate, err := SafePlanForCreation(plan.Plan)
	if err != nil {
		return nil, err
	}

	apiErr := c.create(endpoint, invdPlanDataToCreate, planResp)

	if apiErr != nil {
		return nil, apiErr
	}

	planResp.Connection = c.Connection

	return planResp, nil
}

func (c *Plan) Save() error {
	endpoint := invdendpoint.PlanEndpoint + "/" + c.Id

	planResp := new(Plan)

	planDataToUpdate, err := SafePlanForUpdating(c.Plan)
	if err != nil {
		return err
	}

	apiErr := c.update(endpoint, planDataToUpdate, planResp)

	if apiErr != nil {
		return apiErr
	}

	c.Plan = planResp.Plan

	return nil
}

func (c *Plan) Delete() error {
	endpoint := invdendpoint.PlanEndpoint + "/" + c.Id

	err := c.delete(endpoint)
	if err != nil {
		return err
	}

	return nil
}

func (c *Plan) Retrieve(id string) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint + "/" + id
	planEndpoint := new(invdendpoint.Plan)

	plan := &Plan{c.Connection, planEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Plan) RetrieveWithSubNumber(id string) (*Plan, error) {
	endpoint := invdendpoint.PlanEndpoint + "/" + id + "?include=num_subscriptions"

	planEndpoint := new(invdendpoint.Plan)

	plan := &Plan{c.Connection, planEndpoint}

	_, err := c.retrieveDataFromAPI(endpoint, plan)
	if err != nil {
		return nil, err
	}

	return plan, nil
}

func (c *Plan) ListAllSubNumber(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Plans, error) {
	endpoint := invdendpoint.PlanEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	if strings.Contains(endpoint, "?") {
		endpoint = endpoint + "&include=num_subscriptions"
	} else {
		endpoint = endpoint + "?include=num_subscriptions"
	}

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpPlans)

	if apiErr != nil {
		return nil, apiErr
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	for _, plan := range plans {
		plan.Connection = c.Connection
	}

	return plans, nil
}

func (c *Plan) ListAll(filter *invdendpoint.Filter, sort *invdendpoint.Sort) (Plans, error) {
	endpoint := invdendpoint.PlanEndpoint

	endpoint = addFilterAndSort(endpoint, filter, sort)

	plans := make(Plans, 0)

NEXT:
	tmpPlans := make(Plans, 0)

	endpoint, apiErr := c.retrieveDataFromAPI(endpoint, &tmpPlans)

	if apiErr != nil {
		return nil, apiErr
	}

	plans = append(plans, tmpPlans...)

	if endpoint != "" {
		goto NEXT
	}

	for _, plan := range plans {
		plan.Connection = c.Connection
	}

	return plans, nil
}

// SafeCustomerForCreation prunes plan data for just fields that can be used for creation of a plan
func SafePlanForCreation(plan *invdendpoint.Plan) (*invdendpoint.Plan, error) {
	if plan == nil {
		return nil, errors.New("plan is nil")
	}

	planData := new(invdendpoint.Plan)
	planData.Id = plan.Id
	planData.Object = plan.Object
	planData.Item = plan.Item
	planData.Name = plan.Name
	planData.Currency = plan.Currency
	planData.Amount = plan.Amount
	planData.PricingMode = plan.PricingMode
	planData.QuantityType = plan.QuantityType
	planData.Interval = plan.Interval
	planData.IntervalCount = plan.IntervalCount
	planData.Tiers = plan.Tiers
	planData.CreatedAt = plan.CreatedAt
	planData.Metadata = plan.Metadata

	return planData, nil
}

// SafeTaskForUpdating prunes plan data for just fields that can be used for updating of a plan
func SafePlanForUpdating(plan *invdendpoint.Plan) (*invdendpoint.Plan, error) {
	if plan == nil {
		return nil, errors.New("plan is nil")
	}

	planData := new(invdendpoint.Plan)

	planData.Name = plan.Name
	planData.Metadata = plan.Metadata

	return planData, nil
}
