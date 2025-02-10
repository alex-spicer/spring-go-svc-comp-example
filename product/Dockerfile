FROM amazoncorretto:23-alpine-jdk AS builder
WORKDIR /app
COPY . .
RUN ./gradlew build -x test

FROM amazoncorretto:23-alpine
WORKDIR /app
COPY --from=builder /app/build/libs/*.jar app.jar
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
