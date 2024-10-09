FROM python:3.11-alpine

WORKDIR /

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
COPY . .

EXPOSE 8080
CMD ["python3", "-m", "app"]
