# GitLab Dashboard

Create a simple dashboard with favorites projects and users.

![Preview](https://github.com/rafaelbmateus/simple-gitlab-dashboard/assets/8009492/8ccdb864-977b-47ca-a5c3-f4ec62a358c3)

# Get Started

To run this project in your machine using docker you need:

1. Build docker image

```console
docker build . -t gitlab-dashboard

or:
make build
```

```console
docker run -ti \
  -v $(PWD):/app \
  -e GITLAB_TOKEN=<YOUR_PERSONAL_TOKEN> \
  --network host \
  gitlab-dashboard \
  go run cmd/main.go
 
or:
make build run token=<YOUR_PERSONAL_TOKEN>
```

Change `<YOUR_PERSONAL_TOKEN>` to [your personal token created](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html).
