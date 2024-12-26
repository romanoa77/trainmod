
# DT Datastore Services

This repository contains the source code of the services responsible for storing gravitational waves data(gw) .
Additional software tools for infrastructure testing are included. The services has been designed to be executed inside a cluster running Kubernetes. Their Docker images can be found here:

https://hub.docker.com/repositories/romanoa77

## Data Buffer

The Data Buffer is a storage service written in Go. It stores gw data as JSON files and collects metadata about the infrastructure.

### 1. Requirements
Go 1.23.3 or more recent version.
### 2. Configuration
Environment variables are defined inside the Dockerfile. For a local installation the envsetup.sh file can be used.

### 3. Container Setup
Environment variables description. <br><br>


> ADMROOT

Directory containing metadata.The path is expected to be  mountpoint/directory,  mountpoint is where the kubernetes volume is mounted. 
 
>ADMFNM

 Name of the metadata file relating to the gw stored data

>DSST

Name of the metadata file relating to the datastore state

>DATAROOT

Directory containing gw data. The path is expected to be  mountpoint/directory

>DATANM

Directory containing gw data(duplicate)

>LOGROOT

Log directory location

>LOGSTREAM

Log file name

>SERVURL

Port where the service listens for connections

>DTVOLROOT

Directory containing gw data inside the volume

>WCHSZ



### 4. Kubernetes Setup

Several object configuration files are used: <br><br>

-state.yaml

Defines a statefulset primitive for the service.

-adml.yaml

Defines a volume storing application metadata and the log file.

-gwdataln.yaml

Defines the volume storing gw data.

### 5. Commands
To compile the source first install the required  go libraries, from the main directory

> go mod download

then give the build command

> go build -o databuff

### 6. Directories description

Go is similar to the C programming language so a main function is expected. The main directory contains the main function definition with the addition of copies of the application data. The class directory contains the classes definitions. Go does not provide a class data type like in C++. A program can follow the OOP paradigm using
some language features like the struct data type and modules.

### 7. Application data

Inside a container the application will have the following directory tree:

    .
    ├── application/
    │   └── databuff
    ├── datavar/
    │   └── data/
    └── appdata/
        ├── adm/
        │   ├── StatDesc.json
        │   └── Dsstat.json
        └── log/
             └── LogStream.json

- databuff is the application executable
- The datavar and appdata directories are mountpoints for their respective kubernetes volume
- The data directory contains stored gw data
- The adm directory contains metadata relating to the system
- The log directory contains the log file

Regarding the metadata:

-StatDesc.json contains the following fields:

  - n_itm: the number of files stored
  - buff_size: total number of bytes written

-Dsstat.json contains the following fields:

  - dstatus: datastore state. If the freeze signal has been sent from an external operator. It can assume the OPERATIONAL/FROZEN values
  - user: client identification string
  - token: identity token

The last fields are related to a rudimentary client identification mechanism. 

For a local installation the directory tree is the same as the one inside the Databuffcomp directory.
    
### 8. Endpoints Description

GET methods

>  /stat

The service will send the content of the StatDesc.json file.

>  /dumpLogF


The service will send the content of the log file in the JSON format

>  /dstat

The service will send the content of the Dsstat.json file.

>  /dumpF

The service will send a list of the stored files using the JSON format.



POST methods

>  /sendF

This method accept a JSON file containing gw data. The received data will be stored on disk.

>  /upddsc

The service will update the Dsstat.json file with the content of the received JSON file. This is the the method used by the Datastore logic to freeze the Datastore.

>  /cleanall

This method will unfreeze the datastore. All written data will be stored inside a different directory named as "HOUR MIN DAY". The date refers to when the request has been accepted. All metadata will refers to a datastore accepting data.





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
