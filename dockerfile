FROM golang:1.19-alpine as build-base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go test --tags=unit -v ./...

RUN go build -o ./out/assessment .

# ====================

FROM alpine:3.16.2
COPY --from=build-base /app/out/assessment /app/assessment

CMD ["/app/assessment"]

# ====================

# ***Use
# docker build -t devops:assessment
# docker run devops:assessment

# ***multi stage build Docker image
#    1. go to directory `go-simple-multi-stage-build`
#    2. try to run `docker build -t go-simple-multi-stage:latest .`
#    3. check image size `docker images go-simple-multi-stage`
#    4. try to run container `docker run go-simple-multi-stage`

# ====================

# ***แก้error
# หาข้อมูลเรื่อง=   สอนติดตั้ง Rancher Desktop บน windows
# หาข้อมูลเรื่องแก้error=   error during connect: This error may indicate that the docker daemon is not running.: Get "http://%2F%2F.%2Fpipe%2Fdocker_engine/v1.24/version": open //./pipe/docker_engine: The system cannot find the file specified.
# 1.แก้เป็น golang เวอร์ชั่นของ windows, 1.ทำตามขั้นตอน เรื่องรองรับการ build ด้วย docker โดยใช้ multi-stage build
# 1.แก้ที่ version, อัพเดต docker เป็นของ windows (แก้error)
# 2.integration test สามารถใช้เทคนิค docker-compose testing sandbox (devops)

# ***เรื่องที่ต้องทำของ docker (DevOps)
# 1. ต้องแก้เป็น golang เวอร์ชั่นของ windows หรืออาจจะ ต้องไปทำตามขั้นตอนเริ่มต้นของที่คนอื่นสอน เรื่องรองรับการ build ด้วย docker โดยใช้ multi-stage build
# 1. และต้องไปแก้ที่ version หรืออาจจะ อัพเดต docker เป็นของ windows    = error during connect: This error may indicate that the docker daemon is not running.: Get "http://%2F%2F.%2Fpipe%2Fdocker_engine/v1.24/version": open //./pipe/docker_engine: The system cannot find the file specified.
# 2. integration test สามารถใช้เทคนิค docker-compose testing sandbox ที่เรียนในรายวิชา devops สำหรับรัน postgres ได้เพื่อการทดสอบ