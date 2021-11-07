FROM amd64/golang


WORKDIR /root

COPY . /root
RUN cd /root

RUN go build .

RUN ls .

CMD /root/onetimesecret