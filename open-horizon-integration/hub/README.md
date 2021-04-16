# Delivering EdgeX Foundry with Open Horizon

This is a set of instructions to demonstrate one method for using the [Open Horizon](https://github.com/open-horizon) project to deliver and monitor a simple deployment of [EdgeX Foundry](https://wiki.edgexfoundry.org).

- [Delivering EdgeX Foundry with Open Horizon](#delivering-edgex-foundry-with-open-horizon)
  - [Background](#background)
  - [Goals](#goals)
    - [Run the Services and Agent](#run-the-services-and-agent)
    - [Configure an Organization and Account](#configure-an-organization-and-account)
    - [Register a Service and Deployment Pattern, Policies](#register-a-service-and-deployment-pattern-policies)
    - [Register an Edge Node](#register-an-edge-node)
  - [How to install Open Horizon Management Hub and Agent](#how-to-install-open-horizon-management-hub-and-agent)
  - [Troubleshooting Guide](#troubleshooting-guide)

## Background

The EdgeX Foundry instance will contain a [Random Integer Device Service](https://docs.edgexfoundry.org/1.2/examples/Ch-ExamplesRandomDeviceService/) that will post a simple random event message every five seconds.

## Goals

### Run the Services and Agent

These instructions will show you how to configure and run a single instance of the Open Horizon Management Hub Services (Exchange, Switchboard, AgBots, Sync Service) within a Virtual Machine (VM) and the Open Horizon Agent (Anax) in another VM.  While the Agent could be deployed in the same logical tier as the Services, this example keeps it separate so that you can more easily see how to deploy multiple instances of the Agent connecting to a single deployment of the Services.  

### Configure an Organization and Account

You will configure an organization and an admin-enabled user account (into that organization) on the Horizon Exchange.  This will give you credentials that will be used in subsequent steps below.  By creating an admin account, it will allow you to try out all of the APIs.  By creating a new organization, it will allow you to see how Services, Deployment Patterns, and Policies are all limited to their owning organizations.

### Register a Service and Deployment Pattern, Policies

You will use the Horizon Agent to register the EdgeX micro-services as a single Service.  This will involve creating an asymmetric key pair and using it to sign your service definition while publishing it to the Exchange. Likewise, you will create a Deployment Pattern, publish it, and then register the pattern on the edge node to receive the specified Service.  Last, you can create Policies (Node, Business, and Service) and associate those in order to see the Service automatically deployed to the edge node.

### Register an Edge Node

You will register the device that is running the Agent as an edge node with the Exchange, then associate the Deployment Pattern with the device to trigger Service deployment, which will stand up the micro-services, monitor them, and keep them running.

## How to install Open Horizon Management Hub and Agent

1. [Prerequisites.md](01-prerequisites.md)
2. [Build and Run Open Horizon Management Hub Services](02-build-and-run-horizon.md) 
3. [Install the Open Horizon Agent](03-install-agent.md)
4. [Deploy Open Horizon Services](04-deploy-oh-services.md)
5. [View the Device Data on Agent VM](05-view-device-data.md)
6. [Exposing Open Horizon Agent API](06-expose-agent-api.md)
7. [Pushing New Image to Artifactory](07-push-image-artifactory.md)

## Troubleshooting Guide

[How to troubleshoot](troubleshooting-guide.md)