FROM ubuntu

RUN apt-get update && \
	apt-get install -y software-properties-common wget tmux \
	&& add-apt-repository ppa:neovim-ppa/unstable \
	&& apt-get update \
	&& apt-get install neovim

RUN mkdir /installs
# Install Golang1.5.1
RUN cd /installs && \
wget https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz && tar -C /usr/local -xzf go1.5.1.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH
ENV LANG C.UTF-8

CMD ["tmux", "-u2"]
