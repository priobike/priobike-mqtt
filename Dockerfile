FROM emqx:5.5.1

COPY acl.conf /opt/emqx/etc/acl.conf
COPY emqx.conf /opt/emqx/etc/emqx.conf
