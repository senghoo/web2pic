FROM ubuntu:16.04
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        cutycapt xvfb \
    && rm -rf /var/lib/apt/lists/*
CMD ["xvfb-run", "--server-args='-screen 0, 1024x768x24'", "cutycapt"]
