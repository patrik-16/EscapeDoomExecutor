FROM python:3.10

ARG FILE=app.py

COPY ${FILE} app.py

CMD ["python3", "app.py"]