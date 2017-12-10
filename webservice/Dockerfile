FROM alpine:3.6

# Install ca-certificates
RUN apk -U add ca-certificates

# Create the directory where the application will reside
RUN mkdir -p /opt/webservice
RUN mkdir -p /etc/webservice

# Copy the application files
ADD webservice /opt/webservice/webservice

# Set the working directory to the app directory
WORKDIR /opt/webservice

# Expose the application on port 8080
EXPOSE 8080

# Set the entry point of the container to the application executable
ENTRYPOINT ["/opt/webservice/webservice"]
