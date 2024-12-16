# DT Datastore Services

This repository contains the source code of the services responsible for collecting and storing gravitational waves data.
Additional software tools for infrastructure testing are included. All Docker images can be found here:

https://hub.docker.com/repositories/romanoa77

## Data Buffer

The Data Buffer is a storage service written in Go. It stores gw data as JSON files and collects metadata about the infrastructure.

### 1. Requirements
Go 1.23.3 or more recent version.
### 2. Configuration
Environment variables are defined inside the Dockerfile. For a local installation the envsetup.sh file can be used.

### 3. Container Setup
Environment variables description.__


> ADMROOT

Metadata and log root directory location

ADMFNM  Data descriptor file name

DSST    Services descriptor file name

DATAROOT Data directory location

DATANM   Data root directory location(duplicate)

LOGROOT  Log directory location

LOGSTREAM Log file name

SERVURL Port where the service listens for connections

DTVOLROOT Data root directory location

WCHSZ Internal buffer size

### 4. Kubernetes Setup
### 5. Commands
### 6. Application Data and Filesystem
### 7. Endpoints Description
   
## Datastore Logic

### 1. Requirements
### 2. Configuration
### 3. Container Setup
### 4. Kubernetes Setup
### 5. Commands
### 6. Endpoint Description

## GlitchflowAPI

### 1. Requirements
### 2. Configuration
### 3. Container Setup
### 4. Kubernetes Setup
### 5. Commands
### 6. API Description

## GWclient

## Preproc
