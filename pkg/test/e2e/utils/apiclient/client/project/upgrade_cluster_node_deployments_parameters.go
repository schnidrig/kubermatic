// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// NewUpgradeClusterNodeDeploymentsParams creates a new UpgradeClusterNodeDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpgradeClusterNodeDeploymentsParams() *UpgradeClusterNodeDeploymentsParams {
	return &UpgradeClusterNodeDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeClusterNodeDeploymentsParamsWithTimeout creates a new UpgradeClusterNodeDeploymentsParams object
// with the ability to set a timeout on a request.
func NewUpgradeClusterNodeDeploymentsParamsWithTimeout(timeout time.Duration) *UpgradeClusterNodeDeploymentsParams {
	return &UpgradeClusterNodeDeploymentsParams{
		timeout: timeout,
	}
}

// NewUpgradeClusterNodeDeploymentsParamsWithContext creates a new UpgradeClusterNodeDeploymentsParams object
// with the ability to set a context for a request.
func NewUpgradeClusterNodeDeploymentsParamsWithContext(ctx context.Context) *UpgradeClusterNodeDeploymentsParams {
	return &UpgradeClusterNodeDeploymentsParams{
		Context: ctx,
	}
}

// NewUpgradeClusterNodeDeploymentsParamsWithHTTPClient creates a new UpgradeClusterNodeDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpgradeClusterNodeDeploymentsParamsWithHTTPClient(client *http.Client) *UpgradeClusterNodeDeploymentsParams {
	return &UpgradeClusterNodeDeploymentsParams{
		HTTPClient: client,
	}
}

/*
UpgradeClusterNodeDeploymentsParams contains all the parameters to send to the API endpoint

	for the upgrade cluster node deployments operation.

	Typically these are written to a http.Request.
*/
type UpgradeClusterNodeDeploymentsParams struct {

	// Body.
	Body *models.MasterVersion

	// ClusterID.
	ClusterID string

	// Dc.
	DC string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upgrade cluster node deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeClusterNodeDeploymentsParams) WithDefaults() *UpgradeClusterNodeDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upgrade cluster node deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpgradeClusterNodeDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithTimeout(timeout time.Duration) *UpgradeClusterNodeDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithContext(ctx context.Context) *UpgradeClusterNodeDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithHTTPClient(client *http.Client) *UpgradeClusterNodeDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithBody(body *models.MasterVersion) *UpgradeClusterNodeDeploymentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetBody(body *models.MasterVersion) {
	o.Body = body
}

// WithClusterID adds the clusterID to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithClusterID(clusterID string) *UpgradeClusterNodeDeploymentsParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDC adds the dc to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithDC(dc string) *UpgradeClusterNodeDeploymentsParams {
	o.SetDC(dc)
	return o
}

// SetDC adds the dc to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetDC(dc string) {
	o.DC = dc
}

// WithProjectID adds the projectID to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) WithProjectID(projectID string) *UpgradeClusterNodeDeploymentsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the upgrade cluster node deployments params
func (o *UpgradeClusterNodeDeploymentsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeClusterNodeDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.DC); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
