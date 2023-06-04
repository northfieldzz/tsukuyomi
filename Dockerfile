FROM python:3.11-bullseye as base

RUN apt-get update && apt-get upgrade -y
RUN apt-get install -y ffmpeg

ENV POETRY_HOME=/opt/poetry
# poetryインストール
RUN curl -sSL curl -sSL https://install.python-poetry.org | python - && cd /usr/local/bin && ln -s /opt/poetry/bin/poetry && poetry config virtualenvs.create false

WORKDIR usr/local/src/app



FROM base as development

EXPOSE 8080



FROM base as production

COPY . .

RUN poetry install
EXPOSE 80
ENTRYPOINT poetry run python wsgi.py
