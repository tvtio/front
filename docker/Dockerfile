FROM ubuntu

RUN apt-get update && \
	apt-get install -y software-properties-common wget tmux git \
	&& add-apt-repository ppa:neovim-ppa/unstable \
	&& apt-get update \
	&& apt-get install neovim

RUN mkdir /installs
# Install Golang1.5.1
RUN cd /installs && \
wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz

# Vundle: Vim plugin manager
RUN git clone https://github.com/gmarik/Vundle.vim.git /root/.vim/bundle/Vundle.vim

ENV GOPATH /go
ENV GOBIN /go/bin
ENV PATH /usr/local/go/bin:/go/bin:$PATH
ENV LANG C.UTF-8
ENV HOME /root
WORKDIR /go/src

RUN go version | grep go1.5

ADD entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]

CMD ["vim"]
