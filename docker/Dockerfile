FROM ollama/ollama



WORKDIR /app
COPY ./Modelfile ./
COPY ./startup.sh ./

RUN chmod +x ./startup.sh 
RUN ./startup.sh


EXPOSE 11434


CMD ["ollama","run", "homeSecurity"]