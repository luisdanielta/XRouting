FROM public.ecr.aws/lambda/python:3.12
RUN pip3 install --no-cache-dir awscli

WORKDIR /var/task
COPY adapters core ports utils ./
COPY main.py pyproject.toml requirements.txt setup.py README.md ./

RUN pip install --no-cache-dir -r requirements.txt
RUN chmod +x main.py

CMD ["main.py"]
