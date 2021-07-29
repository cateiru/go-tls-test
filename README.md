# クライアント証明書テスト

Dockerコンテナ内で、Golangのhttp.Getを使用したSSL通信を行うと`x509: certificate signed by unknown authority`エラーが発生する。

## Dockerビルド

- Go

    ```bash
    docker build -t sample .
    ```

- Python

    ```bash
    docker build -t sample -f Dockerfile.python .
    ```
