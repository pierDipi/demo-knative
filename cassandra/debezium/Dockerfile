FROM docker.io/ubuntu as downloader

WORKDIR /opt

RUN apt update && apt install -y wget && \
    wget https://repo1.maven.org/maven2/io/debezium/debezium-connector-cassandra/1.5.0.Final/debezium-connector-cassandra-1.5.0.Final-plugin.tar.gz -O debezium-connector-cassandra.tar.gz && \
    tar xzf debezium-connector-cassandra.tar.gz

FROM docker.io/strimzi/kafka:0.20.1-kafka-2.6.0
USER root:root
RUN mkdir -p /opt/kafka/plugins/debezium
COPY --from=downloader ./opt/debezium-connector-cassandra/ /opt/kafka/plugins/debezium/
USER 1001