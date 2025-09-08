.PHONY: venv
venv:
	python -m venv .venv

.PHONY: activate
activate:
	. .venv/bin/activate

.PHONY: install
install:
	pip install -r requirements.txt

.PHONY: uninstall
uninstall:
	pip uninstall -r requirements.txt -y

.PHONY: test
test:
	python -m unittest discover
