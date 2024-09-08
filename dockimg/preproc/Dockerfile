# syntax=docker/dockerfile:1
FROM python:3.8.10-slim AS build

RUN python3 -m venv /opt/venv

ENV PATH=/opt/venv/bin:$PATH

COPY requirements.txt .

RUN --mount=type=cache,target=/root/.cache/pip pip install \
--no-cache-dir \
-r requirements.txt

RUN --mount=type=cache,target=/root/.cache/pip pip install \
--no-cache-dir \
-i https://test.pypi.org/simple/ distQ==0.0.4




FROM python:3.8.10-slim

ENV PATH=/opt/venv/bin:$PATH
COPY --from=build /opt/venv /opt/venv


CMD [ "sleep", "infinity" ]
