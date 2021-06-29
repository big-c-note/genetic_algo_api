FROM python:3.8-slim-buster

WORKDIR /app
COPY . /app
COPY . /app

RUN cd /app && pip3 install -r requirements.txt && pip3 install -e .
RUN rm -rf *egg* && rm -rf backend_api/__pycache__

CMD ["python3", "/app/backend_api/api.py"]
