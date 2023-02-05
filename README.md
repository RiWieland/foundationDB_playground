## Overview
The repository shows how to store information on underlying binaries in a foundationDB and uses transactions to make consistent changes on that metadata.

## Description
Key-Value stored are often used for storing metadata for underlying binary objects. 
Having a seperate Metastore in a key-value database has several advantages, for example it enables more efficient queries and allow for an addional information layer (e.g. consider the case where you have a video in mp4 format and you want to tack events in the video, this could be done in a seperate key-value store)

## Data Model
