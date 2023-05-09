# AusCERT 2023 - Going back to the basics of API security

:warning: **This code is terribly insecure and is only provided as a demo purpose. Please do not use this in any production like system.**

## Local Testing

If you want to work with the API locally you will ideally need to have Docker Desktop installed (or the docker runtime) for
the easiest way to get started.

You will need to perform the following steps to run locally:

1. Ensure Docker Desktop is installed and running
2. cd to the api folder
```bash
cd api
```
3. Run Docker compose to run the containers
```bash
docker compose up
```
4. Test the environment is running:

```bash
curl http://localhost:9000/healthcheck -v
```

5. You should have a 200 response returned from the API

If you want to make changes to the code then the docker compose environment will auto reload when you save a file. 
If you need to completely remove the environment and start again run the following commands:

```bash
docker rm air
docker rm mongodb
docker compose up
```

## Shared Virtual Machines

There are some parts of the tutorial that will require the use of a VM with a public IP. I have created the following VMs
that everyone can share:

* 3.106.122.32
* 54.252.81.134
* 13.211.81.218
* 3.106.60.130
* 13.54.54.25
* 3.26.202.108
* 3.26.64.49
* 54.153.179.214
* 3.26.155.144
* 3.25.173.194

```bash
ssh ubuntu@<ip>
```
