# ImageRepo
An image repository storing images in blob storage buckets. For now, this is built with a local storage filesystem for testing and development. The next step is to move to S3 buckets paired with an AWS session or an Azure Blob Storage

# Tech Stack
- Go
- MySQL

# Assumptions
This project acts as an internal service and assumes that Auth is handled in a separate service. 
-Assume https (everything is encrypted)
-Assume front end images are encoded in base64 


-"GET" image request will return an S3 address -> curl address for image