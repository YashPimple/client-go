FROM ubuntu

COPY ./Client-go ./Client-go

ENTRYPOINT [ "./Client-go" ]



