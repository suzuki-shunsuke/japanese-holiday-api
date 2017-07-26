FROM scratch
COPY ./server ./syukujitsu.csv /workspace/
WORKDIR /workspace
CMD ["/workspace/server"]
