FROM golang:latest

RUN mkdir -p /go/src/GoWebFull/7-scalable_GoAPI_with_Kubernetes

WORKDIR /go/src/GoWebFull/7-scalable_GoAPI_with_Kubernetes

COPY . /go/src/GoWebFull/7-scalable_GoAPI_with_Kubernetes

RUN go install 7-scalable_GoAPI_with_Kubernetes

CMD /go/bin/GoWebFull/7-scalable_GoAPI_with_Kubernetes

EXPOSE 8080