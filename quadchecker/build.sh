#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}╔══════════════════════════════════════╗${NC}"
echo -e "${BLUE}║       Building All Quad Programs      ║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════╝${NC}"
echo ""

# Print current directory structure
echo -e "${YELLOW}Current directory structure:${NC}"
find . -name "*.go" -type f | sort | sed 's/^/  /'
echo ""

# Initialize counters
success_count=0
fail_count=0

# Function to create go.mod if missing
create_gomod() {
    local dir="$1"
    local module_name="$2"
    
    if [ ! -f "$dir/go.mod" ]; then
        echo -e "  ${YELLOW}Creating go.mod for $module_name...${NC}"
        echo "module $module_name" > "$dir/go.mod"
        echo "go 1.21" >> "$dir/go.mod"
    fi
}

# Function to build a quad program
build_quad() {
    local quad_name="$1"
    local quad_path="quad/$quad_name"
    
    echo -e "${BLUE}▶ Building $quad_name...${NC}"
    
    # Check if directory exists
    if [ ! -d "$quad_path" ]; then
        echo -e "  ${RED}✗ Directory not found: $quad_path${NC}"
        ((fail_count++))
        return 1
    fi
    
    # Check if main.go exists
    if [ ! -f "$quad_path/main.go" ]; then
        echo -e "  ${RED}✗ main.go not found in $quad_path${NC}"
        ((fail_count++))
        return 1
    fi
    
    # Create go.mod if missing
    create_gomod "$quad_path" "$quad_name"
    
    # Build the quad
    cd "$quad_path"
    go build -o "../../$quad_name"
    local build_status=$?
    cd ../..
    
    if [ $build_status -eq 0 ]; then
        echo -e "  ${GREEN}✓ Successfully built $quad_name${NC}"
        chmod +x "$quad_name"
        ((success_count++))
        return 0
    else
        echo -e "  ${RED}✗ Failed to build $quad_name${NC}"
        echo -e "  ${YELLOW}Trying alternative build method...${NC}"
        
        # Try alternative build method
        cd "$quad_path"
        go build -o "$quad_name" && mv "$quad_name" ../../
        local alt_status=$?
        cd ../..
        
        if [ $alt_status -eq 0 ]; then
            echo -e "  ${GREEN}✓ Successfully built $quad_name (alternative method)${NC}"
            chmod +x "$quad_name"
            ((success_count++))
            return 0
        else
            ((fail_count++))
            return 1
        fi
    fi
}

# Function to build quadchecker
build_quadchecker() {
    echo -e "${BLUE}▶ Building quadchecker...${NC}"
    
    # Check if directory exists
    if [ ! -d "quadchecker" ]; then
        echo -e "  ${RED}✗ Directory not found: quadchecker${NC}"
        ((fail_count++))
        return 1
    fi
    
    # Check if main.go exists
    if [ ! -f "quadchecker/main.go" ]; then
        echo -e "  ${RED}✗ main.go not found in quadchecker${NC}"
        ((fail_count++))
        return 1
    fi
    
    # Create go.mod if missing
    if [ ! -f "quadchecker/go.mod" ]; then
        echo -e "  ${YELLOW}Creating go.mod for quadchecker...${NC}"
        echo "module quadchecker" > quadchecker/go.mod
        echo "go 1.21" >> quadchecker/go.mod
    fi
    
    # Build quadchecker
    cd quadchecker
    go build -o "../quadchecker"
    local build_status=$?
    cd ..
    
    if [ $build_status -eq 0 ]; then
        echo -e "  ${GREEN}✓ Successfully built quadchecker${NC}"
        chmod +x quadchecker
        ((success_count++))
        return 0
    else
        echo -e "  ${RED}✗ Failed to build quadchecker${NC}"
        ((fail_count++))
        return 1
    fi
}

# Build all quad programs
echo -e "${YELLOW}Building Quad Programs:${NC}"
for quad in quadA quadB quadC quadD quadE; do
    build_quad "$quad"
done

echo ""
echo -e "${YELLOW}Building QuadChecker:${NC}"
build_quadchecker

echo ""
echo -e "${BLUE}╔══════════════════════════════════════╗${NC}"
echo -e "${BLUE}║           Build Summary              ║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════╝${NC}"
echo ""

# List all built executables in root directory
echo -e "${YELLOW}Executables in root directory:${NC}"
for exec in quadA quadB quadC quadD quadE quadchecker; do
    if [ -f "$exec" ] && [ -x "$exec" ]; then
        size=$(stat -c%s "$exec" 2>/dev/null || stat -f%z "$exec" 2>/dev/null)
        echo -e "  ${GREEN}✓ $exec ($(numfmt --to=iec $size 2>/dev/null || echo "${size} bytes"))${NC}"
    elif [ -f "$exec" ]; then
        echo -e "  ${YELLOW}⚠ $exec (exists but not executable)${NC}"
        chmod +x "$exec" 2>/dev/null
    else
        echo -e "  ${RED}✗ $exec (not found)${NC}"
    fi
done

# Print summary
echo ""
if [ $fail_count -eq 0 ]; then
    echo -e "${GREEN}✅ All $success_count builds successful!${NC}"
else
    echo -e "${YELLOW}📊 Build Results:${NC}"
    echo -e "  ${GREEN}Successful: $success_count${NC}"
    echo -e "  ${RED}Failed: $fail_count${NC}"
fi

echo ""
echo -e "${YELLOW}📋 Quick Test Commands:${NC}"
echo "  ./quadA 3 3 | ./quadchecker"
echo "  ./quadC 1 1 | ./quadchecker"
echo "  echo -e 'o---o\\\\n|   |\\\\no---o' | ./quadchecker"
echo "  echo 'A' | ./quadchecker"

echo ""
echo -e "${GREEN}✨ Build process complete!${NC}"