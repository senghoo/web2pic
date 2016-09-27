FROM ubuntu:16.04
RUN echo deb http://archive.canonical.com/ubuntu xenial partner >> /etc/apt/sources.list && \
    apt-get update \
    && apt-get install -y --no-install-recommends \
        cutycapt xvfb \
        fonts-arphic-ukai fonts-arphic-uming ttf-wqy-zenhei \
         adobe-flashplugin\
    && rm -rf /var/lib/apt/lists/*
ADD capture /usr/local/bin
RUN fc-cache -v
CMD ["/usr/local/bin/capture"]
