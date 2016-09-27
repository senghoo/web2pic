FROM ubuntu:16.04
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        cutycapt xvfb \
        fonts-arphic-ukai fonts-arphic-uming ttf-wqy-zenhei \
        flashplugin-nonfree \
    && rm -rf /var/lib/apt/lists/*
ADD capture /usr/local/bin
RUN fc-cache -v
CMD ["/usr/local/bin"]
