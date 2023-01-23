# FairShare-Blockchain-Enabled-Fair-Accountable-and-Secure-Data-Sharing-for-Industrial-IoT
This repo contains the implementation code for the paper "FairShare: Blockchain Enabled Fair Accountable and Secure Data Sharing for Industrial IoT"

This would only work on an Ethereum private network. For this setup, we have deployed two blockchain peer nodes on the network. We have launched three terminals each of which represents the fog node, cloud and client respectively. Respective algorithms are triggered from the corresponding terminals to simulate the entire framework. Each of the terminals create a separate channel with the underlying blockchain network. 

The "deploy.go" file is used to initialise the entire system.
