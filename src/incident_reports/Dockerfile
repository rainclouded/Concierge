FROM python:3.11-alpine

WORKDIR /incident_reports_server

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
COPY . .

EXPOSE 8080
CMD ["python", "/incident_reports_server/controllers/incident_reports_controller.py"]
