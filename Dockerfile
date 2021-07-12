FROM golang:1.16.5-buster
RUN go get -u github.com/beego/bee
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_USER app
ENV APP_HOME /go/src/todoapp
USER $APP_USER
WORKDIR $APP_HOME
EXPOSE 8010
CMD ["bee", "run"]