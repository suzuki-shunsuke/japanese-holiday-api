FROM scratch
COPY ./server /workspace/
COPY ./config/ /workspace/config/
COPY ./data /workspace/data/
WORKDIR /workspace
CMD ["/workspace/server"]
