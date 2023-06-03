
FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/golang As builder

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

#step 1 build go cache
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download

#step 2 build binary project
WORKDIR /project
ADD . .
RUN ls
RUN go build main.go

FROM sunmi-docker-images-registry.cn-hangzhou.cr.aliyuncs.com/public/centos:7.8.2003
#run binary project
WORKDIR /app
COPY --from=builder /project/main .

# your project shell [project] [arg1] [arg2] ...
CMD [ "/app/main","api","start"]