#!/bin/bash

mockgen -source=../git/interface/repository.go -destination=../mocks/mock_repository.go -package=mocks
mockgen -source=../git/interface/walk.go -destination=../mocks/mock_walk.go -package=mocks