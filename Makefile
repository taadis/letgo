.PHONY: venv
venv:
	python -m venv .venv && source .venv/bin/activate

.PHONY: install
install:
	pip install -r requirements.txt
