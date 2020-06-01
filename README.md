# ImageRepo
An image repository storing images in blob storage buckets. For now, this is built with a local storage filesystem for testing and development. The next step is to move to S3 buckets paired with an AWS session or an Azure Blob Storage

# Tech Stack
- Go
- MySQL

# Assumptions
This project acts as a part of a microservice and assumes that Auth is handled in a separate service. 