# 🔗 Distributed File System (DFS) using gRPC

A simplified yet robust **Distributed File System** written in Go that supports uploading, downloading, and replicating `.mp4` files across multiple nodes. This system simulates a fault-tolerant architecture with master coordination, heartbeat checks, and gRPC-based communication.

## 📐 Architecture

This DFS follows a centralized coordination model with three key components:

- **Master Tracker Node**:
  - Maintains a lookup table with: `filename`, `data node`, `file path`, `isAlive`.
  - Handles client requests for upload/download.
  - Coordinates data replication and monitors node health via heartbeats.

- **Data Keeper Nodes**:
  - Store actual file data.
  - Send heartbeat messages to the master.
  - Accept file uploads from clients and replicate files on command.

- **Client**:
  - Uploads `.mp4` files to the system.
  - Downloads files from available data nodes.
  - Interacts only with the Master Tracker initially.

## 🔌 Communication

- Communication is handled via **gRPC**.

## ❤️ Heartbeats

- Every **1 second**, each Data Keeper sends a heartbeat to the Master Tracker.
- Master updates its lookup table based on live status.
- If a Data Keeper becomes unresponsive, it's marked as dead in the table.

## ⬆️ Upload Protocol

1. Client requests an upload slot from the Master Tracker.
2. Master returns a port of an active Data Keeper.
3. Client uploads the `.mp4` file directly to the Data Keeper over TCP.
4. Data Keeper notifies the Master after upload completion.
5. Master updates the lookup table and sends a success response to the client.
6. Master initiates **replication** to 2 additional nodes.

## 📁 Replication Logic

- Every **10 seconds**, the Master scans its file records.
- Ensures each file is stored on **at least 3 live nodes**.
- If replication is needed:
  - Master chooses a source and target node.
  - Notifies both to transfer the file.

## ⬇️ Download Protocol

1. Client requests a file from the Master Tracker.
2. Master replies with a list of IPs/ports of nodes holding the file.
3. Client downloads the file in parallel as chunks from each of the nodes that have the file.

## ✅ Features

- 📡 gRPC-based communication
- 🧠 Centralized master with smart replication
- ♻️ Automatic recovery via periodic replication checks
- 🔄 Heartbeat-based node monitoring
- 🚀 Multi-threaded handling for concurrent requests
- 🔐 Fault tolerance with replicated storage

## 🛠️ Technologies

- Go (Golang)
- gRPC
- Concurrency via Goroutines and Mutex locks

### Contributors 

<table align="center" >
  <tr>
      <td align="center"><a href="https://github.com/SH8664"><img src="https://avatars.githubusercontent.com/u/113303945?v=4" width="150px;" alt=""/><br /><sub><b>Sara Bisheer</b></sub></a><br /></td>
      <td align="center"><a href="https://github.com/rawanMostafa08"><img src="https://avatars.githubusercontent.com/u/97397431?v=4" width="150px;" alt=""/><br /><sub><b>Rawan Mostafa</b></sub></a><br /></td>
      <td align="center"><a href="https://github.com//mennamohamed0207"><img src="https://avatars.githubusercontent.com/u/90017398?v=4" width="150px;" alt=""/><br /><sub><b>Menna Mohammed</b></sub></a><br /></td>
      <td align="center"><a href="https://github.com/fatmaebrahim"><img src="https://avatars.githubusercontent.com/u/113191710?v=4" width="150;" alt=""/><br /><sub><b>Fatma Ebrahim</b></sub></a><br /></td>
  </tr>
</table>
