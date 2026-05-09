#!/bin/bash
find . -name "*.sh" | sed 's:.*/::;s/\.sh$//'