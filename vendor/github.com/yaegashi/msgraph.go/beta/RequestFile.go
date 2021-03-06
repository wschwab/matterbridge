// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// FileClassificationRequestObjectRequestBuilder is request builder for FileClassificationRequestObject
type FileClassificationRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns FileClassificationRequestObjectRequest
func (b *FileClassificationRequestObjectRequestBuilder) Request() *FileClassificationRequestObjectRequest {
	return &FileClassificationRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FileClassificationRequestObjectRequest is request for FileClassificationRequestObject
type FileClassificationRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Get(ctx context.Context) (resObj *FileClassificationRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Update(ctx context.Context, reqObj *FileClassificationRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for FileClassificationRequestObject
func (r *FileClassificationRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// FileSecurityProfileRequestBuilder is request builder for FileSecurityProfile
type FileSecurityProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns FileSecurityProfileRequest
func (b *FileSecurityProfileRequestBuilder) Request() *FileSecurityProfileRequest {
	return &FileSecurityProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// FileSecurityProfileRequest is request for FileSecurityProfile
type FileSecurityProfileRequest struct{ BaseRequest }

// Get performs GET request for FileSecurityProfile
func (r *FileSecurityProfileRequest) Get(ctx context.Context) (resObj *FileSecurityProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for FileSecurityProfile
func (r *FileSecurityProfileRequest) Update(ctx context.Context, reqObj *FileSecurityProfile) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for FileSecurityProfile
func (r *FileSecurityProfileRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
