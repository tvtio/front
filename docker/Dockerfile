FROM ubuntu

RUN apt-get update && \
	apt-get install -y software-properties-common tmux \
	&& add-apt-repository ppa:neovim-ppa/unstable \
	&& apt-get update \
	&& apt-get install neovim

ENV LANG C.UTF-8

CMD ["tmux", "-u2"]
