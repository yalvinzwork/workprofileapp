FROM ubuntu:26.04

COPY workprofileapp /app

EXPOSE 9000/tcp

CMD ["/app"]