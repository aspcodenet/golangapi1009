FROM golang:1.23.2-alpine
# INSTALLERA GO
# komppilera vår kod -> EXE-fil
# kör EXE-fil

COPY . .
RUN go get -d -v
RUN go build -o /app/cmd/main

# EXPOSE 8080 

ENTRYPOINT [ "/app/cmd/main" ]

# innehåller intruktioner för att skapa en CONTAINER IMAGE
# 1. Ta en tom Windows
# 2. Ladda ner C-runtime
# 3. Ladda ner Java v8
# 4. Ta min zio.fil och kiopera in på c:\bla\bla
# 5. Kör c:\bla\bla\aaa.exe
