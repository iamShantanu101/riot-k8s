# Build Step
FROM golang:1.10-alpine AS build

# Build
ARG build
ARG version
RUN apk add git
RUN go get github.com/iamShantanu101/riot-k8s/cmd
RUN mkdir -p $GOPATH/src/build/
COPY . $GOPATH/src/build/
WORKDIR $GOPATH/src/build
RUN CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${version} -X main.Build=${build}" -o riot
RUN cp riot /

# Final Step
FROM alpine
ARG KUBECTLVERSION=v1.10.5
# Base packages
RUN apk update
RUN apk upgrade
RUN apk add curl tar gzip
RUN rm -rf /var/cache/apk/*
RUN curl -L -o /usr/bin/kubectl https://storage.googleapis.com/kubernetes-release/release/${KUBECTLVERSION}/bin/linux/amd64/kubectl && \
  chmod +x /usr/bin/kubectl
# Copy binary from build step
COPY --from=build /riot /home/riot
RUN mkdir -p /home/manifests
COPY ./manifests /home/manifests
# Define the ENTRYPOINT
WORKDIR /home
RUN ls -ltr /home
ENTRYPOINT ["./riot"]
