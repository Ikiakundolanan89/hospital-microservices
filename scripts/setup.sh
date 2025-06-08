#!/bin/bash

echo "🏥 Setting up Hospital Microservice - Patient Service"
echo "=================================================="

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed. Please install Go 1.21 or later.${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Go is installed${NC}"

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo -e "${YELLOW}Docker is not installed. You'll need Docker to run SQL Server locally.${NC}"
fi

# Create project structure
echo -e "\n${YELLOW}Creating project structure...${NC}"

# Create directories
directories=(
    "services/patient-service/cmd"
    "services/patient-service/internal/config"
    "services/patient-service/internal/domain"
    "services/patient-service/internal/repository"
    "services/patient-service/internal/service"
    "services/patient-service/internal/handler"
    "services/patient-service/internal/middleware"
    "services/patient-service/internal/dto"
    "services/patient-service/internal/database/migrations"
    "services/patient-service/pkg/validator"
    "services/patient-service/pkg/utils"
    "shared/jwt"
    "shared/logger"
    "shared/monitoring"
    "shared/errors"
    "kubernetes/patient-service"
    "scripts"
)

for dir in "${directories[@]}"; do
    mkdir -p "$dir"
    echo -e "${GREEN}✓ Created $dir${NC}"
done

# Copy .env.example to .env
if [ -f "services/patient-service/.env.example" ]; then
    cp services/patient-service/.env.example services/patient-service/.env
    echo -e "${GREEN}✓ Created .env file${NC}"
fi

# Install Go dependencies
echo -e "\n${YELLOW}Installing Go dependencies...${NC}"
cd services/patient-service
go mod init patient-service 2>/dev/null || true
go get -u github.com/gofiber/fiber/v2
go get -u github.com/denisenkom/go-mssqldb
go get -u github.com/joho/godotenv
go get -u github.com/golang-jwt/jwt/v5
go get -u github.com/google/uuid
go get -u github.com/go-playground/validator/v10
go get -u github.com/prometheus/client_golang
go get -u github.com/gofiber/adaptor/v2

echo -e "${GREEN}✓ Dependencies installed${NC}"

# Install development tools
echo -e "\n${YELLOW}Installing development tools...${NC}"
go install github.com/cosmtrek/air@latest
go install github.com/swaggo/swag/cmd/swag@latest
echo -e "${GREEN}✓ Development tools installed${NC}"

# Create air config for hot reload
cat > .air.toml << EOF
root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ./cmd/main.go"
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_error = true

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false
EOF

echo -e "${GREEN}✓ Air config created${NC}"

# Start SQL Server with Docker
echo -e "\n${YELLOW}Starting SQL Server with Docker...${NC}"
if command -v docker &> /dev/null; then
    docker run -e "ACCEPT_EULA=Y" -e "SA_PASSWORD=YourStrong@Passw0rd" \
        -p 1433:1433 --name hospital_sqlserver \
        -d mcr.microsoft.com/mssql/server:2019-latest
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓ SQL Server started${NC}"
        echo -e "${YELLOW}Waiting for SQL Server to be ready...${NC}"
        sleep 15
    else
        echo -e "${YELLOW}SQL Server container might already be running${NC}"
    fi
else
    echo -e "${RED}Docker not found. Please start SQL Server manually.${NC}"
fi

echo -e "\n${GREEN}🎉 Setup completed!${NC}"
echo -e "\n${YELLOW}Next steps:${NC}"
echo "1. cd services/patient-service"
echo "2. Update .env file with your configuration"
echo "3. Run 'make run' or 'air' for hot reload"
echo "4. Access the service at http://localhost:3001"
echo "5. Health check: http://localhost:3001/health"
echo "6. Metrics: http://localhost:3001/metrics"
echo -e "\n${YELLOW}For production:${NC}"
echo "- Update JWT_SECRET in .env"
echo "- Configure proper database credentials"
echo "- Set up proper monitoring with Prometheus/Grafana"
echo "- Configure Kubernetes deployments"
