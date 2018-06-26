// +-------------------------------------------------------------------------
// | Copyright (C) 2016 Yunify, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------




package service

import (
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"

    "github.com/yunify/qingstor-sdk-go/config"
    "github.com/yunify/qingstor-sdk-go/request"
    "github.com/yunify/qingstor-sdk-go/request/data"
    "github.com/yunify/qingstor-sdk-go/request/errors"
)

var _ fmt.State
var _ io.Reader
var _ http.Header
var _ strings.Reader
var _ time.Time
var _ config.Config


    // Metadata presents metadata.
    type Metadata struct {
        Config     *config.Config
        Properties *Properties
    }

    // Metadata initializes a new metadata.
    func (s *Service) Metadata(bucketName string,zone string,) (*Metadata, error) {
                zone = strings.ToLower(zone)
            properties := &Properties{
            BucketName: &bucketName,
            Zone: &zone,
            }

        return &Metadata{Config: s.Config, Properties: properties}, nil
    }



    
    
    

    
    
    

    
    

    
    
    
    
    
    

    // DeleteMetadata does Delete a Metadata.
        // Documentation URL: https://docs.qingcloud.com/qingstor/api/common/metadata
    func (s *Metadata) DeleteMetadata() (*DeleteMetadataOutput, error) {
    r, x, err := s.DeleteMetadataRequest()
        
        if err != nil {
            return x, err
        }

        err = r.Send()
        if err != nil {
            return nil, err
        }

        requestID := r.HTTPResponse.Header.Get(http.CanonicalHeaderKey("X-QS-Request-ID"))
        x.RequestID = &requestID

        return x, err
    }

    // DeleteMetadataRequest creates request and output object of DeleteMetadata.
        func (s *Metadata) DeleteMetadataRequest() (*request.Request, *DeleteMetadataOutput, error) {
    
        
        
        
        
            properties := *s.Properties
        

        o := &data.Operation{
            Config:        s.Config,
                Properties:    &properties,
            APIName:       "DELETE Metadata",
            RequestMethod: "DELETE",
            RequestURI:    "/<bucket-name>?metadata",
            StatusCodes: []int{
                204, // Metadata deleted
                    },
        }

        x := &DeleteMetadataOutput{}
        r, err := request.New(o, nil, x)
        if err != nil {
            return nil, nil, err
        }

        return r, x, nil
    }

    

    // DeleteMetadataOutput presents output for DeleteMetadata.
    type DeleteMetadataOutput struct {
        StatusCode *int `location:"statusCode"`

        RequestID *string `location:"requestID"`
        

            

            
        
    }
    
    


    
    
    

    
    
    

    
    

    
    
    
    
    
    

    // GetMetadata does GET Metadata.
        // Documentation URL: https://docs.qingcloud.com/qingstor/api/common/metadata
    func (s *Metadata) GetMetadata() (*GetMetadataOutput, error) {
    r, x, err := s.GetMetadataRequest()
        
        if err != nil {
            return x, err
        }

        err = r.Send()
        if err != nil {
            return nil, err
        }

        requestID := r.HTTPResponse.Header.Get(http.CanonicalHeaderKey("X-QS-Request-ID"))
        x.RequestID = &requestID

        return x, err
    }

    // GetMetadataRequest creates request and output object of GetMetadata.
        func (s *Metadata) GetMetadataRequest() (*request.Request, *GetMetadataOutput, error) {
    
        
        
        
        
            properties := *s.Properties
        

        o := &data.Operation{
            Config:        s.Config,
                Properties:    &properties,
            APIName:       "GET Metadata",
            RequestMethod: "GET",
            RequestURI:    "/<bucket-name>?metadata",
            StatusCodes: []int{
                200, // OK
                    },
        }

        x := &GetMetadataOutput{}
        r, err := request.New(o, nil, x)
        if err != nil {
            return nil, nil, err
        }

        return r, x, nil
    }

    

    // GetMetadataOutput presents output for GetMetadata.
    type GetMetadataOutput struct {
        StatusCode *int `location:"statusCode"`

        RequestID *string `location:"requestID"`
        

            
                
                // caching mechanism
            CacheControl *string `json:"Cache-Control,omitempty" name:"Cache-Control" location:"elements"` 
    // the default filename when an object is downloaded
            ContentDisposition *string `json:"Content-Disposition,omitempty" name:"Content-Disposition" location:"elements"` 
    // the content encoding type of the object
            ContentEncoding *string `json:"Content-Encoding,omitempty" name:"Content-Encoding" location:"elements"` 
    // the expiration date and time
            Expires *string `json:"Expires,omitempty" name:"Expires" location:"elements"` 
    // Custom metadata
            XQSMeta* *string `json:"x-qs-meta-*,omitempty" name:"x-qs-meta-*" location:"elements"` 
    

            

            
        
    }
    
    


    
    
    

    
    
    

    
    

    
    
    
    
    
    

    // PutMetadata does Create a new Metadata.
        // Documentation URL: https://docs.qingcloud.com/qingstor/api/common/metadata
    func (s *Metadata) PutMetadata(input *PutMetadataInput) (*PutMetadataOutput, error) {
    r, x, err := s.PutMetadataRequest(input)
        
        if err != nil {
            return x, err
        }

        err = r.Send()
        if err != nil {
            return nil, err
        }

        requestID := r.HTTPResponse.Header.Get(http.CanonicalHeaderKey("X-QS-Request-ID"))
        x.RequestID = &requestID

        return x, err
    }

    // PutMetadataRequest creates request and output object of PutMetadata.
        func (s *Metadata) PutMetadataRequest(input *PutMetadataInput) (*request.Request, *PutMetadataOutput, error) {
    
            if input == nil {
                input = &PutMetadataInput{}
            }
        
        
        
        
        
            properties := *s.Properties
        

        o := &data.Operation{
            Config:        s.Config,
                Properties:    &properties,
            APIName:       "PUT Metadata",
            RequestMethod: "PUT",
            RequestURI:    "/<bucket-name>?metadata",
            StatusCodes: []int{
                201, // Metadata created
                    },
        }

        x := &PutMetadataOutput{}
        r, err := request.New(o, input, x)
        if err != nil {
            return nil, nil, err
        }

        return r, x, nil
    }

    
        // PutMetadataInput presents input for PutMetadata.
        type PutMetadataInput struct {
                // Specify the caching mechanism that the request and response follow
            CacheControl *string `json:"Cache-Control,omitempty" name:"Cache-Control" location:"headers"` 
    // Specify the default filename when an object is downloaded
            ContentDisposition *string `json:"Content-Disposition,omitempty" name:"Content-Disposition" location:"headers"` 
    // Specify the content encoding type of the object
            ContentEncoding *string `json:"Content-Encoding,omitempty" name:"Content-Encoding" location:"headers"` 
    // Respond to the expiration date and time
            Expires *string `json:"Expires,omitempty" name:"Expires" location:"headers"` 
    // Custom metadata
            XQSMeta* *string `json:"x-qs-meta-*,omitempty" name:"x-qs-meta-*" location:"headers"` 
    

            
        }

        // Validate validates the input for PutMetadata.
        func (v *PutMetadataInput) Validate() error {
            
    
    

    

            
    
    

    
        
            
            
                
                
                
            

            

            
        
    
        
            
            
                
                
                
            

            

            
        
    
        
            
            
                
                
                
            

            

            
        
    
        
            
            
                
                
                
            

            

            
        
    
        
            
            
                
                
                
            

            

            
        
    

            
    
    

    


            return nil
        }
    

    // PutMetadataOutput presents output for PutMetadata.
    type PutMetadataOutput struct {
        StatusCode *int `location:"statusCode"`

        RequestID *string `location:"requestID"`
        

            

            
        
    }
    
    



















