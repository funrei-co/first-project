# 2020/10/14最新versionを取得
FROM golang:1.23rc1-bullseye
# ワーキングディレクトリの設定
WORKDIR /go
# ホストのファイルをコンテナの作業ディレクトリに移行
#ADD . /go/src/app

