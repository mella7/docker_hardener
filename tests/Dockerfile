# Use a minimal base image
FROM ubuntu:20.04

# Set non-root user (modified by the tool)
USER nonrootuser

# Install essential packages only
RUN apt-get update && \
    apt-get install --no-install-recommends -y curl wget

# Set environment variables for security
ENV APP_ENV=production

# Clean up unnecessary files
RUN rm -rf /var/lib/apt/lists/*

