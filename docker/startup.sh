#!/bin/bash


ollama pull phi3 &&\
ollama create homesecurity -f ./Modelfile
