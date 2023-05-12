# GitLab Dashboard

Create a simple dashboard with favorites projects and users.

![Preview](https://github.com/rafaelbmateus/simple-gitlab-dashboard/assets/8009492/8ccdb864-977b-47ca-a5c3-f4ec62a358c3)

# Get Started

To run this project in your machine using docker you need:

1. Create a [personal token](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html) in GitLab
and change from `<YOUR_PERSONAL_TOKEN>` to your token.

2. Create `config.yaml` file:

```console
cp config-example.yaml config-yaml
```

3. Build docker image:

```console
docker build . -t gitlab-dashboard
```

4. Run docker container:

```console
docker run -ti \
  -e GITLAB_TOKEN=<YOUR_PERSONAL_TOKEN> \
  gitlab-dashboard go run cmd/main.go
```
