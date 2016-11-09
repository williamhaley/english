#!/usr/bin/env bash

curl -v \
	http://localhost:8000/api/v1/words \
	-d '{
		"term": "cat",
		"definition": "a feline"
	}'

