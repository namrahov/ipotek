FROM ictcontact/alpine:3.13-nonroot

WORKDIR /app
COPY --chown=appuser ms-retail-products-info ./
COPY --chown=appuser profiles/default.env ./profiles/
COPY --chown=appuser migrations/*.sql ./migrations/
COPY --chown=appuser templates/* ./templates/
COPY --chown=appuser ./infra/entrypoint-parent.sh .

RUN sudo setcap CAP_NET_BIND_SERVICE=+eip ms-retail-products-info

RUN chmod 500 entrypoint-parent.sh

EXPOSE 80

CMD [ "./entrypoint-parent.sh" ]
