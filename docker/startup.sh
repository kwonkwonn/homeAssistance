#!/bin/bash

ollama serve&&\
ollama pull phi3 &&\
ollama create homesecurity -f ./Modelfile