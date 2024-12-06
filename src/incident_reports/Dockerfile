FROM python:3.11-alpine

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
COPY . .

RUN chown -R appuser:appgroup /app
USER appuser

EXPOSE 8080
CMD ["python", "-m", "incident_reports_server.controllers.incident_reports_controller"]
