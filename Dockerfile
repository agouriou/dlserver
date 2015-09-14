FROM busybox

ADD dlserver /usr/local/dlserver/dlserver

CMD ["/usr/local/dlserver/dlserver"]
