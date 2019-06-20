build:
  GOOS=linux go build -ldflags="-s -w" . && upx go-github-pr-files
