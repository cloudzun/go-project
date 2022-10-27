FROM registry.cn-beijing.aliyuncs.com/dotbalo/alpine-glibc:alpine-3.9

COPY conf/  ./conf 

COPY ./go-project ./ 

ENTRYPOINT [ "./go-project"] 
