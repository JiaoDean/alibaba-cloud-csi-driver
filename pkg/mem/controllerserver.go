/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mem

import (
	"errors"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"golang.org/x/net/context"
)

type controllerServer struct {
	*csicommon.DefaultControllerServer
}

// newControllerServer creates a controllerServer object
func newControllerServer(d *csicommon.CSIDriver) *controllerServer {
	return &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
	}
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		return nil, err
	}
	if len(req.Name) == 0 {
		errMsg := fmt.Sprintf("Volume Name cannot be empty")
		log.Errorf(log.TypeMEM, log.StatusInternalError, errMsg)
		return nil, errors.New(errMsg)
	}
	if req.VolumeCapabilities == nil {
		errMsg := fmt.Sprintf("Volume Capabilities cannot be empty")
		log.Errorf(log.TypeMEM, log.StatusInternalError, errMsg)
		return nil, errors.New(errMsg)
	}

	volumeID := req.GetName()
	response := &csi.CreateVolumeResponse{}

	response = &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      volumeID,
			CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
			VolumeContext: req.GetParameters(),
		},
	}

	log.Infof(log.TypeMEM, log.StatusOK, fmt.Sprintf("Create volumeID %s is successfully, size:%s", volumeID, string(req.GetCapacityRange().GetRequiredBytes())))
	return response, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof(log.TypeMEM, log.StatusOK, fmt.Sprintf("Delete volumeID %s is successfully", req.GetVolumeId()))
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof(log.TypeMEM, log.StatusOK, fmt.Sprintf("ControllerUnpublish volumeID %s is successfully", req.GetVolumeId()))
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof(log.TypeMEM, log.StatusOK, fmt.Sprintf("ControllerPublish volumeID %s is successfully", req.GetVolumeId()))
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof(log.TypeMEM, log.StatusOK, fmt.Sprintf("ControllerExpand volumeID %s is successfully", req.GetVolumeId()))
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}
