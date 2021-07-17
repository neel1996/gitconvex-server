#!/bin/bash

mockgen -source=../git/middleware/repository.go -destination=../mocks/mock_repository.go -package=mocks
mockgen -source=../git/middleware/walk.go -destination=../mocks/mock_walk.go -package=mocks
mockgen -source=../git/middleware/reference.go -destination=../mocks/mock_reference.go -package=mocks
mockgen -source=../git/middleware/index.go -destination=../mocks/mock_index.go -package=mocks