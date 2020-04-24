# It's necessary to set this because some environments don't link sh -> bash.
SHELL := /bin/bash

include ./make/*.mk
.DEFAULT_GOAL := help