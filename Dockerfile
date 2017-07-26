FROM scratch
COPY ./server ./syukujitsu.csv ./config.toml /workspace/
WORKDIR /workspace
CMD ["/workspace/server"]
