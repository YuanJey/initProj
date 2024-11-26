package main

const (
	dockerfile = "FROM golang:latest AS builder\nENV CGO_ENABLED 0\nENV GOPROXY https://goproxy.cn,direct\nWORKDIR /app\nCOPY . .\nRUN go mod tidy\nWORKDIR /app/cmd/main\nRUN go build -o main .\n\n\nFROM alpine:latest\nRUN apk add --no-cache tzdata\nENV TZ=Asia/Shanghai\nWORKDIR /root/\nCOPY --from=builder /app/cmd/main/main .\nCOPY --from=builder /app/config/config.yaml .\nCOPY --from=builder /app/start.sh .\nRUN chmod +x start.sh\nEXPOSE 8000\nCMD [\"sh\", \"start.sh\"]"
	startSH    = "#!/bin/sh\n./main &\nLOG_DIR=\"./logs\"\nCURRENT_DATE=$(date +%Y-%m-%d)\nLOG_FILE=\"$LOG_DIR/log.all.$CURRENT_DATE\"\nmkdir -p \"$LOG_DIR\"\ntouch \"$LOG_FILE\"\nwhile true; do\n    NEW_DATE=$(date +%Y-%m-%d)\n    if [ \"$NEW_DATE\" != \"$CURRENT_DATE\" ]; then\n        CURRENT_DATE=\"$NEW_DATE\"\n        LOG_FILE=\"$LOG_DIR/log.all.$CURRENT_DATE\"\n        touch \"$LOG_FILE\"\n    fi\n    if [ ! -f \"$LOG_DIR/log.all.$CURRENT_DATE\" ]; then\n        touch \"$LOG_DIR/log.all.$CURRENT_DATE\"\n    fi\n    if ! pgrep -f \"tail -f $LOG_FILE\" > /dev/null; then\n        tail -f \"$LOG_FILE\" &\n    fi\n    sleep 60\ndone"
	mainFile   = "package main\nimport (\n\t\"github.com/gin-gonic/gin\"\n)\nfunc main() {\n\tr := gin.New()\n\tr.Use(gin.Recovery())\n\taddress := \"0.0.0.0:8000\"\n\terr := r.Run(address)\n\tif err != nil {\n\t\tpanic(\"api start failed \" + err.Error())\n\t}\n}"
)