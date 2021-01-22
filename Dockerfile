FROM golang:1.14

WORKDIR /terraform-provider-gorillastack

COPY . .

# Install Documentation generator
RUN go get github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["terraform-provider-gorillastack"]