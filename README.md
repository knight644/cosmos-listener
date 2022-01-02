# Saagemonitor
Utility tool to monitor the saage chain. It listens for new block creation events and for every new block created, it fetches monotoring info and reward info.

## Getting Started
Get the sourcecode from cloning the repo
`git clone https://gitlab.com/Live_Project/saage/saagesports/saagemonitor`

OR

Get the latest relese source code from https://gitlab.com/Live_Project/saage/saagesports/saagemonitor/-/releases/v0.0.1-DevNet

### Installing and Running the monitor
To install the application, navigate to the root directory of the sourcecode and run: 
```
go install
```

Once the application is installed, you can run it form anywhere using `saagemonitor` command

## Configure
Initialization parameters of the proxy server can be found in:
```
saagemonitor/configs/config.go
```

A list of `constant` parameters
|Name         |Type  |Description                                               |
|-------------|------|----------------------------------------------------------|
|ChainAddress |string|URI of the chain to listen to                             |
|ChainRPCPort |int   |RPC port of the chain (required to open websocket)        |
|ChainRESTPort|int   |REST port of the chain (required to fetch additional info)|
|ChainWSPath  |string|path to subscribe to websocket                            |
|ChainName    |string|name of the chain to subscribe to                         |

Change the *chain address* and the *chain name* to listen to the chain you desire.

## Architecture
The saagemonitor consists of the following modules
* **configs**: Configuration options for the monitor
* **record**: File writer functionalities
* **subscription**: Subscribing to blockchain events and their corresponding handling
* **main.go**: Entrypoint for saagemonitor

## Working
The saagemonitor, when run, initializes the file-writers and opens an websocket to listen to blockchain events
It listens for "NewBlock" events in oarticular. On receiving this event, it analyzes the block header and associated events to find the monitor parameters values. The floowing parameters are recorded for each new block that is created:
1. **Block height**
2. **Block time**
3. **Inflation**
4. **Inflation Change**
5. **Total supply of uSaage tokens**
6. **Total amount of uSaage tokens staked**
7. **Ratio of staked tokens**

This info is written to `monitorRecord.csv` file

It also records the rewards earned by validators for each block. The following info is recorded:
1. **Block height**
2. **Block time**
3. **Address of the proposer of the block**
4. **Reward earned by proposer for creating the block**
5. **Total reward earned by proposer**: block creation reward + validation reward + [bonus reward for including all other validator signatures]
6. **[validator address, validator reward]**: a list of validators followed by the rewards they earned

This info is written to `validatorRewardRecord.csv` file.
