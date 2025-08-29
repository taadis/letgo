.PHONY: venv
venv:
	python -m venv .venv && source .venv/bin/activate

.PHONY: install
install:
	pip install -r requirements.txt

.PHONY: uninstall
uninstall:
	pip uninstall -r requirements.txt -y
