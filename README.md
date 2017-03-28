# rancher-forwarder

To run this image:

docker run -e DST_IP="10.42.173.8" -e DST_PORT="8080" -p 93:9090 juandavidgc/rancher-forwarder