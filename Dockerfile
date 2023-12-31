# syntax=docker/dockerfile:1
# 使用するベースイメージ
FROM golang:latest

ENV TZ=Asia/Tokyo

# root以外のユーザーを作成
RUN useradd -m go_user && \
    apt-get update && \
    apt-get install -y sudo neovim && \
    echo "go_user ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/go_user && \
    chmod 0440 /etc/sudoers.d/go_user && \
    chsh -s /bin/bash go_user && \
    chown -R go_user:go_user /home/go_user

# 作成したユーザーに切り替え
USER go_user

# 作業ディレクトリを指定
WORKDIR /home/go_user

ENTRYPOINT [ "/bin/bash" ]
