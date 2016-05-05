FROM alpine
MAINTAINER Sebastian Vera <u.verainf@gmail.com>
COPY tiresias tiresias
EXPOSE 3000
ENTRYPOINT ["/tiresias"]
