// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018-2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

// This package provides a implementation of a ProtocolDriver interface.
//
package driver

import (
        "fmt"
        "sync"
        "time"

        dsModels "github.com/edgexfoundry/device-sdk-go/pkg/models"
        "github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
        "github.com/edgexfoundry/go-mod-core-contracts/models"
)

var once sync.Once
var driver *RandomDriver

type RandomDriver struct {
        lc            logger.LoggingClient
        asyncCh       chan<- *dsModels.AsyncValues
        randomDevices sync.Map
}

func NewProtocolDriver() dsModels.ProtocolDriver {
        once.Do(func() {
                driver = new(RandomDriver)
        })
        return driver
}

func (d *RandomDriver) DisconnectDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
        d.lc.Info(fmt.Sprintf("RandomDriver.DisconnectDevice: device-random driver is disconnecting to %s", deviceName))
        return nil
}

func (d *RandomDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *dsModels.AsyncValues, deviceCh chan<- []dsModels.DiscoveredDevice) error {
        d.lc = lc
        d.asyncCh = asyncCh
        return nil
}

func (d *RandomDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest) (res []*dsModels.CommandValue, err error) {
        rd := d.retrieveRandomDevice(deviceName)

        res = make([]*dsModels.CommandValue, len(reqs))
        now := time.Now().UnixNano()

        for i, req := range reqs {
                n := req.DeviceResourceName
                t := req.Type
                v, err := rd.value(t)
                if err != nil {
                        return nil, err
                }
                var cv *dsModels.CommandValue
                switch n {
                case "RandomTemperature":
                        cv, _ = dsModels.NewInt32Value(req.DeviceResourceName, now, int32(v))                }
                res[i] = cv
        }

        return res, nil
}

func (d *RandomDriver) retrieveRandomDevice(deviceName string) (rdv *randomDevice) {
        rd, ok := d.randomDevices.LoadOrStore(deviceName, newRandomDevice())
        if rdv, ok = rd.(*randomDevice); !ok {
                panic("The value in randomDevices has to be a reference of randomDevice")
        }
        return rdv
}

func (d *RandomDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []dsModels.CommandRequest,
        params []*dsModels.CommandValue) error {
        rd := d.retrieveRandomDevice(deviceName)

        for _, param := range params {
                switch param.DeviceResourceName {
                case "Min_Temperature":
                        v, err := param.Int32Value()
                        if err != nil {
                                return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
                        }
                        if v < defMinTemperature {
                                return fmt.Errorf("RandomDriver.HandleWriteCommands: minimum value %d of %T must be int between %d ~ %d", v, v, defMinTemperature, defMaxTemperature)
                }

                        rd.minTemperature = int64(v)
                case "Max_Temperature":
                        v, err := param.Int32Value()
                        if err != nil {
                                return fmt.Errorf("RandomDriver.HandleWriteCommands: %v", err)
                        }
                        if v > defMaxTemperature {
                                return fmt.Errorf("RandomDriver.HandleWriteCommands: maximum value %d of %T must be int between %d ~ %d", v, v, defMinTemperature, defMaxTemperature)
                }

                        rd.maxTemperature = int64(v)
                default:
                        return fmt.Errorf("RandomDriver.HandleWriteCommands: there is no matched device resource for %s", param.String())
                }
        }

        return nil
}

func (d *RandomDriver) Stop(force bool) error {
        d.lc.Info("RandomDriver.Stop: device-random driver is stopping...")
        return nil
}

func (d *RandomDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
        d.lc.Debug(fmt.Sprintf("a new Device is added: %s", deviceName))
        return nil
}

func (d *RandomDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {
        d.lc.Debug(fmt.Sprintf("Device %s is updated", deviceName))
        return nil
}

func (d *RandomDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {
        d.lc.Debug(fmt.Sprintf("Device %s is removed", deviceName))
        return nil
}