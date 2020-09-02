# ImageRepo
An image repository storing images in blob storage buckets. For now, this is built with a local storage filesystem for testing and development. The next step is to move to S3 buckets paired with an AWS session or an Azure Blob Storage

# Tech Stack
- Go
- MySQL

# Assumptions
This project acts as an internal service and assumes that Auth is handled in a separate service. It assumes use of https (everything is encrypted). It also assumes front end images are encoded in base64 

# How it Works
- Build the schema: mysql < schema.sql
- Build and run the application: go build && ./imageRepo server

# Next steps
- Complete unit tests & table driven testing
- Complete implementation of message queue