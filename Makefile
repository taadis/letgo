.PHONY: venv
venv:
	python -m venv .venv

.PHONY: venv-activate
venv-activate:
	. .venv/bin/activate

.PHONY: install
install:
	pip install -r requirements.txt

.PHONY: run
run:
	python main.py

.PHONY: test
test:
	python -m unittest tests/xxx.py
